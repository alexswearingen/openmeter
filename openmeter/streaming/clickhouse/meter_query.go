package clickhouse

import (
	_ "embed"
	"fmt"
	"sort"
	"time"

	"github.com/huandu/go-sqlbuilder"

	"github.com/openmeterio/openmeter/openmeter/meter"
	"github.com/openmeterio/openmeter/pkg/slicesx"
)

type queryMeter struct {
	Database        string
	EventsTableName string
	Namespace       string
	Meter           meter.Meter
	Subject         []string
	FilterGroupBy   map[string][]string
	From            *time.Time
	To              *time.Time
	GroupBy         []string
	WindowSize      *meter.WindowSize
	WindowTimeZone  *time.Location
}

// from returns the from time for the query.
func (d *queryMeter) from() *time.Time {
	// If none of the from times are set, return nil
	if d.From == nil && d.Meter.EventFrom == nil {
		return nil
	}

	// If only the event from time is set, use it
	if d.From == nil && d.Meter.EventFrom != nil {
		return d.Meter.EventFrom
	}

	// If only the query from time is set, use it
	if d.From != nil && d.Meter.EventFrom == nil {
		return d.From
	}

	// If both the query from time and the event from time are set
	// use the query from time if it's after the event from time
	if d.From.After(*d.Meter.EventFrom) {
		return d.From
	}

	return d.Meter.EventFrom
}

// toCountRowSQL returns the SQL query for the estimated number of rows.
// This estimate is useful for query progress tracking.
// We only filter by columns that are in the ClickHouse table order.
func (d *queryMeter) toCountRowSQL() (string, []interface{}) {
	tableName := getTableName(d.Database, d.EventsTableName)
	getColumn := columnFactory(d.EventsTableName)
	timeColumn := getColumn("time")

	query := sqlbuilder.ClickHouse.NewSelectBuilder()
	query.Select("count() AS total")
	query.From(tableName)

	// The event table is ordered by namespace, type
	query.Where(query.Equal("namespace", d.Namespace))
	query.Where(query.Equal("type", d.Meter.EventType))

	if len(d.Subject) > 0 {
		mapFunc := func(subject string) string {
			return query.Equal(getColumn("subject"), subject)
		}

		query.Where(query.Or(slicesx.Map(d.Subject, mapFunc)...))
	}

	// The event table is partitioned by time
	from := d.from()

	if from != nil {
		query.Where(query.GreaterEqualThan(timeColumn, from.Unix()))
	}

	if d.To != nil {
		query.Where(query.LessEqualThan(timeColumn, d.To.Unix()))
	}

	sql, args := query.Build()
	return sql, args
}

func (d *queryMeter) toSQL() (string, []interface{}, error) {
	tableName := getTableName(d.Database, d.EventsTableName)
	getColumn := columnFactory(d.EventsTableName)
	timeColumn := getColumn("time")

	var selectColumns, groupByColumns, where []string

	// Select windows if any
	groupByWindowSize := d.WindowSize != nil

	tz := "UTC"
	if d.WindowTimeZone != nil {
		tz = d.WindowTimeZone.String()
	}

	if groupByWindowSize {
		switch *d.WindowSize {
		case meter.WindowSizeMinute:
			selectColumns = append(
				selectColumns,
				fmt.Sprintf("tumbleStart(%s, toIntervalMinute(1), '%s') AS windowstart", timeColumn, tz),
				fmt.Sprintf("tumbleEnd(%s, toIntervalMinute(1), '%s') AS windowend", timeColumn, tz),
			)

		case meter.WindowSizeHour:
			selectColumns = append(
				selectColumns,
				fmt.Sprintf("tumbleStart(%s, toIntervalHour(1), '%s') AS windowstart", timeColumn, tz),
				fmt.Sprintf("tumbleEnd(%s, toIntervalHour(1), '%s') AS windowend", timeColumn, tz),
			)

		case meter.WindowSizeDay:
			selectColumns = append(
				selectColumns,
				fmt.Sprintf("tumbleStart(%s, toIntervalDay(1), '%s') AS windowstart", timeColumn, tz),
				fmt.Sprintf("tumbleEnd(%s, toIntervalDay(1), '%s') AS windowend", timeColumn, tz),
			)

		default:
			return "", nil, fmt.Errorf("invalid window size type: %s", *d.WindowSize)
		}

		groupByColumns = append(groupByColumns, "windowstart", "windowend")
	} else {
		// TODO: remove this when we don't round to the nearest minute anymore
		// We round them to the nearest minute to ensure the result is the same as with
		// streaming connector using materialized views with per minute windows
		selectColumn := fmt.Sprintf("tumbleStart(min(%s), toIntervalMinute(1)) AS windowstart, tumbleEnd(max(%s), toIntervalMinute(1)) AS windowend", timeColumn, timeColumn)
		selectColumns = append(selectColumns, selectColumn)
	}

	// Select Value
	sqlAggregation := ""
	switch d.Meter.Aggregation {
	case meter.MeterAggregationSum:
		sqlAggregation = "sum"
	case meter.MeterAggregationAvg:
		sqlAggregation = "avg"
	case meter.MeterAggregationMin:
		sqlAggregation = "min"
	case meter.MeterAggregationMax:
		sqlAggregation = "max"
	case meter.MeterAggregationUniqueCount:
		sqlAggregation = "uniq"
	case meter.MeterAggregationCount:
		sqlAggregation = "count"
	default:
		return "", []interface{}{}, fmt.Errorf("invalid aggregation type: %s", d.Meter.Aggregation)
	}

	switch d.Meter.Aggregation {
	case meter.MeterAggregationCount:
		selectColumns = append(selectColumns, fmt.Sprintf("toFloat64(%s(*)) AS value", sqlAggregation))
	case meter.MeterAggregationUniqueCount:
		selectColumns = append(selectColumns, fmt.Sprintf("toFloat64(%s(JSON_VALUE(%s, '%s'))) AS value", sqlAggregation, getColumn("data"), sqlbuilder.Escape(*d.Meter.ValueProperty)))
	default:
		// JSON_VALUE returns an empty string if the JSON Path is not found. With toFloat64OrNull we convert it to NULL so the aggregation function can handle it properly.
		selectColumns = append(selectColumns, fmt.Sprintf("%s(toFloat64OrNull(JSON_VALUE(%s, '%s'))) AS value", sqlAggregation, getColumn("data"), sqlbuilder.Escape(*d.Meter.ValueProperty)))
	}

	for _, groupByKey := range d.GroupBy {
		// Subject is a special case as it's a top level column
		if groupByKey == "subject" {
			selectColumns = append(selectColumns, getColumn("subject"))
			groupByColumns = append(groupByColumns, "subject")
			continue
		}

		// Group by columns need to be parsed from the JSON data
		groupByColumn := sqlbuilder.Escape(groupByKey)
		groupByJSONPath := sqlbuilder.Escape(d.Meter.GroupBy[groupByKey])
		selectColumn := fmt.Sprintf("JSON_VALUE(%s, '%s') as %s", getColumn("data"), groupByJSONPath, groupByColumn)

		selectColumns = append(selectColumns, selectColumn)
		groupByColumns = append(groupByColumns, groupByColumn)
	}

	query := sqlbuilder.ClickHouse.NewSelectBuilder()
	query.Select(selectColumns...)
	query.From(tableName)
	query.Where(query.Equal(getColumn("namespace"), d.Namespace))
	query.Where(query.Equal(getColumn("type"), d.Meter.EventType))

	if len(d.Subject) > 0 {
		mapFunc := func(subject string) string {
			return query.Equal(getColumn("subject"), subject)
		}

		where = append(where, query.Or(slicesx.Map(d.Subject, mapFunc)...))
	}

	if len(d.FilterGroupBy) > 0 {
		// We sort the group by s to ensure the query is deterministic
		groupByKeys := make([]string, 0, len(d.FilterGroupBy))
		for k := range d.FilterGroupBy {
			groupByKeys = append(groupByKeys, k)
		}
		sort.Strings(groupByKeys)

		for _, groupByKey := range groupByKeys {
			if _, ok := d.Meter.GroupBy[groupByKey]; !ok {
				return "", nil, fmt.Errorf("meter does not have group by: %s", groupByKey)
			}

			groupByJSONPath := sqlbuilder.Escape(d.Meter.GroupBy[groupByKey])

			values := d.FilterGroupBy[groupByKey]
			if len(values) == 0 {
				return "", nil, fmt.Errorf("empty filter for group by: %s", groupByKey)
			}
			mapFunc := func(value string) string {
				column := fmt.Sprintf("JSON_VALUE(%s, '%s')", getColumn("data"), groupByJSONPath)

				// Subject is a special case
				if groupByKey == "subject" {
					column = "subject"
				}

				return fmt.Sprintf("%s = '%s'", column, sqlbuilder.Escape((value)))
			}

			where = append(where, query.Or(slicesx.Map(values, mapFunc)...))
		}
	}

	from := d.from()

	if from != nil {
		where = append(where, query.GreaterEqualThan(timeColumn, from.Unix()))
	}

	if d.To != nil {
		where = append(where, query.LessEqualThan(timeColumn, d.To.Unix()))
	}

	if len(where) > 0 {
		query.Where(where...)
	}

	query.GroupBy(groupByColumns...)

	if groupByWindowSize {
		query.OrderBy("windowstart")
	}

	sql, args := query.Build()
	return sql, args, nil
}

type listMeterSubjectsQuery struct {
	Database        string
	EventsTableName string
	Namespace       string
	Meter           meter.Meter
	From            *time.Time
	To              *time.Time
}

func (d listMeterSubjectsQuery) toSQL() (string, []interface{}) {
	tableName := getTableName(d.Database, d.EventsTableName)

	sb := sqlbuilder.ClickHouse.NewSelectBuilder()
	sb.Select("DISTINCT subject")
	sb.Where(sb.Equal("namespace", d.Namespace))
	sb.Where(sb.Equal("type", d.Meter.EventType))
	sb.From(tableName)
	sb.OrderBy("subject")

	if d.From != nil {
		sb.Where(sb.GreaterEqualThan("time", d.From.Unix()))
	}

	if d.To != nil {
		sb.Where(sb.LessEqualThan("time", d.To.Unix()))
	}

	sql, args := sb.Build()
	return sql, args
}

func columnFactory(alias string) func(string) string {
	return func(column string) string {
		return fmt.Sprintf("%s.%s", alias, column)
	}
}
