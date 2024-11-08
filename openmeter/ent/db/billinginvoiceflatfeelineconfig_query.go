// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/openmeter/ent/db/billinginvoiceflatfeelineconfig"
	"github.com/openmeterio/openmeter/openmeter/ent/db/predicate"
)

// BillingInvoiceFlatFeeLineConfigQuery is the builder for querying BillingInvoiceFlatFeeLineConfig entities.
type BillingInvoiceFlatFeeLineConfigQuery struct {
	config
	ctx        *QueryContext
	order      []billinginvoiceflatfeelineconfig.OrderOption
	inters     []Interceptor
	predicates []predicate.BillingInvoiceFlatFeeLineConfig
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BillingInvoiceFlatFeeLineConfigQuery builder.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) Where(ps ...predicate.BillingInvoiceFlatFeeLineConfig) *BillingInvoiceFlatFeeLineConfigQuery {
	bifflcq.predicates = append(bifflcq.predicates, ps...)
	return bifflcq
}

// Limit the number of records to be returned by this query.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) Limit(limit int) *BillingInvoiceFlatFeeLineConfigQuery {
	bifflcq.ctx.Limit = &limit
	return bifflcq
}

// Offset to start from.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) Offset(offset int) *BillingInvoiceFlatFeeLineConfigQuery {
	bifflcq.ctx.Offset = &offset
	return bifflcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) Unique(unique bool) *BillingInvoiceFlatFeeLineConfigQuery {
	bifflcq.ctx.Unique = &unique
	return bifflcq
}

// Order specifies how the records should be ordered.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) Order(o ...billinginvoiceflatfeelineconfig.OrderOption) *BillingInvoiceFlatFeeLineConfigQuery {
	bifflcq.order = append(bifflcq.order, o...)
	return bifflcq
}

// First returns the first BillingInvoiceFlatFeeLineConfig entity from the query.
// Returns a *NotFoundError when no BillingInvoiceFlatFeeLineConfig was found.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) First(ctx context.Context) (*BillingInvoiceFlatFeeLineConfig, error) {
	nodes, err := bifflcq.Limit(1).All(setContextOp(ctx, bifflcq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{billinginvoiceflatfeelineconfig.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) FirstX(ctx context.Context) *BillingInvoiceFlatFeeLineConfig {
	node, err := bifflcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BillingInvoiceFlatFeeLineConfig ID from the query.
// Returns a *NotFoundError when no BillingInvoiceFlatFeeLineConfig ID was found.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bifflcq.Limit(1).IDs(setContextOp(ctx, bifflcq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{billinginvoiceflatfeelineconfig.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) FirstIDX(ctx context.Context) string {
	id, err := bifflcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BillingInvoiceFlatFeeLineConfig entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BillingInvoiceFlatFeeLineConfig entity is found.
// Returns a *NotFoundError when no BillingInvoiceFlatFeeLineConfig entities are found.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) Only(ctx context.Context) (*BillingInvoiceFlatFeeLineConfig, error) {
	nodes, err := bifflcq.Limit(2).All(setContextOp(ctx, bifflcq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{billinginvoiceflatfeelineconfig.Label}
	default:
		return nil, &NotSingularError{billinginvoiceflatfeelineconfig.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) OnlyX(ctx context.Context) *BillingInvoiceFlatFeeLineConfig {
	node, err := bifflcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BillingInvoiceFlatFeeLineConfig ID in the query.
// Returns a *NotSingularError when more than one BillingInvoiceFlatFeeLineConfig ID is found.
// Returns a *NotFoundError when no entities are found.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bifflcq.Limit(2).IDs(setContextOp(ctx, bifflcq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{billinginvoiceflatfeelineconfig.Label}
	default:
		err = &NotSingularError{billinginvoiceflatfeelineconfig.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) OnlyIDX(ctx context.Context) string {
	id, err := bifflcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BillingInvoiceFlatFeeLineConfigs.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) All(ctx context.Context) ([]*BillingInvoiceFlatFeeLineConfig, error) {
	ctx = setContextOp(ctx, bifflcq.ctx, ent.OpQueryAll)
	if err := bifflcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BillingInvoiceFlatFeeLineConfig, *BillingInvoiceFlatFeeLineConfigQuery]()
	return withInterceptors[[]*BillingInvoiceFlatFeeLineConfig](ctx, bifflcq, qr, bifflcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) AllX(ctx context.Context) []*BillingInvoiceFlatFeeLineConfig {
	nodes, err := bifflcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BillingInvoiceFlatFeeLineConfig IDs.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) IDs(ctx context.Context) (ids []string, err error) {
	if bifflcq.ctx.Unique == nil && bifflcq.path != nil {
		bifflcq.Unique(true)
	}
	ctx = setContextOp(ctx, bifflcq.ctx, ent.OpQueryIDs)
	if err = bifflcq.Select(billinginvoiceflatfeelineconfig.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) IDsX(ctx context.Context) []string {
	ids, err := bifflcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bifflcq.ctx, ent.OpQueryCount)
	if err := bifflcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bifflcq, querierCount[*BillingInvoiceFlatFeeLineConfigQuery](), bifflcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) CountX(ctx context.Context) int {
	count, err := bifflcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bifflcq.ctx, ent.OpQueryExist)
	switch _, err := bifflcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) ExistX(ctx context.Context) bool {
	exist, err := bifflcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BillingInvoiceFlatFeeLineConfigQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) Clone() *BillingInvoiceFlatFeeLineConfigQuery {
	if bifflcq == nil {
		return nil
	}
	return &BillingInvoiceFlatFeeLineConfigQuery{
		config:     bifflcq.config,
		ctx:        bifflcq.ctx.Clone(),
		order:      append([]billinginvoiceflatfeelineconfig.OrderOption{}, bifflcq.order...),
		inters:     append([]Interceptor{}, bifflcq.inters...),
		predicates: append([]predicate.BillingInvoiceFlatFeeLineConfig{}, bifflcq.predicates...),
		// clone intermediate query.
		sql:  bifflcq.sql.Clone(),
		path: bifflcq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Namespace string `json:"namespace,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BillingInvoiceFlatFeeLineConfig.Query().
//		GroupBy(billinginvoiceflatfeelineconfig.FieldNamespace).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) GroupBy(field string, fields ...string) *BillingInvoiceFlatFeeLineConfigGroupBy {
	bifflcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BillingInvoiceFlatFeeLineConfigGroupBy{build: bifflcq}
	grbuild.flds = &bifflcq.ctx.Fields
	grbuild.label = billinginvoiceflatfeelineconfig.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Namespace string `json:"namespace,omitempty"`
//	}
//
//	client.BillingInvoiceFlatFeeLineConfig.Query().
//		Select(billinginvoiceflatfeelineconfig.FieldNamespace).
//		Scan(ctx, &v)
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) Select(fields ...string) *BillingInvoiceFlatFeeLineConfigSelect {
	bifflcq.ctx.Fields = append(bifflcq.ctx.Fields, fields...)
	sbuild := &BillingInvoiceFlatFeeLineConfigSelect{BillingInvoiceFlatFeeLineConfigQuery: bifflcq}
	sbuild.label = billinginvoiceflatfeelineconfig.Label
	sbuild.flds, sbuild.scan = &bifflcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BillingInvoiceFlatFeeLineConfigSelect configured with the given aggregations.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) Aggregate(fns ...AggregateFunc) *BillingInvoiceFlatFeeLineConfigSelect {
	return bifflcq.Select().Aggregate(fns...)
}

func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bifflcq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bifflcq); err != nil {
				return err
			}
		}
	}
	for _, f := range bifflcq.ctx.Fields {
		if !billinginvoiceflatfeelineconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if bifflcq.path != nil {
		prev, err := bifflcq.path(ctx)
		if err != nil {
			return err
		}
		bifflcq.sql = prev
	}
	return nil
}

func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BillingInvoiceFlatFeeLineConfig, error) {
	var (
		nodes = []*BillingInvoiceFlatFeeLineConfig{}
		_spec = bifflcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BillingInvoiceFlatFeeLineConfig).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BillingInvoiceFlatFeeLineConfig{config: bifflcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(bifflcq.modifiers) > 0 {
		_spec.Modifiers = bifflcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bifflcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bifflcq.querySpec()
	if len(bifflcq.modifiers) > 0 {
		_spec.Modifiers = bifflcq.modifiers
	}
	_spec.Node.Columns = bifflcq.ctx.Fields
	if len(bifflcq.ctx.Fields) > 0 {
		_spec.Unique = bifflcq.ctx.Unique != nil && *bifflcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bifflcq.driver, _spec)
}

func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(billinginvoiceflatfeelineconfig.Table, billinginvoiceflatfeelineconfig.Columns, sqlgraph.NewFieldSpec(billinginvoiceflatfeelineconfig.FieldID, field.TypeString))
	_spec.From = bifflcq.sql
	if unique := bifflcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bifflcq.path != nil {
		_spec.Unique = true
	}
	if fields := bifflcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, billinginvoiceflatfeelineconfig.FieldID)
		for i := range fields {
			if fields[i] != billinginvoiceflatfeelineconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bifflcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bifflcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bifflcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bifflcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bifflcq.driver.Dialect())
	t1 := builder.Table(billinginvoiceflatfeelineconfig.Table)
	columns := bifflcq.ctx.Fields
	if len(columns) == 0 {
		columns = billinginvoiceflatfeelineconfig.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bifflcq.sql != nil {
		selector = bifflcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bifflcq.ctx.Unique != nil && *bifflcq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range bifflcq.modifiers {
		m(selector)
	}
	for _, p := range bifflcq.predicates {
		p(selector)
	}
	for _, p := range bifflcq.order {
		p(selector)
	}
	if offset := bifflcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bifflcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) ForUpdate(opts ...sql.LockOption) *BillingInvoiceFlatFeeLineConfigQuery {
	if bifflcq.driver.Dialect() == dialect.Postgres {
		bifflcq.Unique(false)
	}
	bifflcq.modifiers = append(bifflcq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return bifflcq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (bifflcq *BillingInvoiceFlatFeeLineConfigQuery) ForShare(opts ...sql.LockOption) *BillingInvoiceFlatFeeLineConfigQuery {
	if bifflcq.driver.Dialect() == dialect.Postgres {
		bifflcq.Unique(false)
	}
	bifflcq.modifiers = append(bifflcq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return bifflcq
}

// BillingInvoiceFlatFeeLineConfigGroupBy is the group-by builder for BillingInvoiceFlatFeeLineConfig entities.
type BillingInvoiceFlatFeeLineConfigGroupBy struct {
	selector
	build *BillingInvoiceFlatFeeLineConfigQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bifflcgb *BillingInvoiceFlatFeeLineConfigGroupBy) Aggregate(fns ...AggregateFunc) *BillingInvoiceFlatFeeLineConfigGroupBy {
	bifflcgb.fns = append(bifflcgb.fns, fns...)
	return bifflcgb
}

// Scan applies the selector query and scans the result into the given value.
func (bifflcgb *BillingInvoiceFlatFeeLineConfigGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bifflcgb.build.ctx, ent.OpQueryGroupBy)
	if err := bifflcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillingInvoiceFlatFeeLineConfigQuery, *BillingInvoiceFlatFeeLineConfigGroupBy](ctx, bifflcgb.build, bifflcgb, bifflcgb.build.inters, v)
}

func (bifflcgb *BillingInvoiceFlatFeeLineConfigGroupBy) sqlScan(ctx context.Context, root *BillingInvoiceFlatFeeLineConfigQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bifflcgb.fns))
	for _, fn := range bifflcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bifflcgb.flds)+len(bifflcgb.fns))
		for _, f := range *bifflcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bifflcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bifflcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BillingInvoiceFlatFeeLineConfigSelect is the builder for selecting fields of BillingInvoiceFlatFeeLineConfig entities.
type BillingInvoiceFlatFeeLineConfigSelect struct {
	*BillingInvoiceFlatFeeLineConfigQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bifflcs *BillingInvoiceFlatFeeLineConfigSelect) Aggregate(fns ...AggregateFunc) *BillingInvoiceFlatFeeLineConfigSelect {
	bifflcs.fns = append(bifflcs.fns, fns...)
	return bifflcs
}

// Scan applies the selector query and scans the result into the given value.
func (bifflcs *BillingInvoiceFlatFeeLineConfigSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bifflcs.ctx, ent.OpQuerySelect)
	if err := bifflcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillingInvoiceFlatFeeLineConfigQuery, *BillingInvoiceFlatFeeLineConfigSelect](ctx, bifflcs.BillingInvoiceFlatFeeLineConfigQuery, bifflcs, bifflcs.inters, v)
}

func (bifflcs *BillingInvoiceFlatFeeLineConfigSelect) sqlScan(ctx context.Context, root *BillingInvoiceFlatFeeLineConfigQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bifflcs.fns))
	for _, fn := range bifflcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bifflcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bifflcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
