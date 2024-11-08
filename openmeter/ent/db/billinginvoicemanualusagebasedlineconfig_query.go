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
	"github.com/openmeterio/openmeter/openmeter/ent/db/billinginvoicemanualusagebasedlineconfig"
	"github.com/openmeterio/openmeter/openmeter/ent/db/predicate"
)

// BillingInvoiceManualUsageBasedLineConfigQuery is the builder for querying BillingInvoiceManualUsageBasedLineConfig entities.
type BillingInvoiceManualUsageBasedLineConfigQuery struct {
	config
	ctx        *QueryContext
	order      []billinginvoicemanualusagebasedlineconfig.OrderOption
	inters     []Interceptor
	predicates []predicate.BillingInvoiceManualUsageBasedLineConfig
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BillingInvoiceManualUsageBasedLineConfigQuery builder.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) Where(ps ...predicate.BillingInvoiceManualUsageBasedLineConfig) *BillingInvoiceManualUsageBasedLineConfigQuery {
	bimublcq.predicates = append(bimublcq.predicates, ps...)
	return bimublcq
}

// Limit the number of records to be returned by this query.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) Limit(limit int) *BillingInvoiceManualUsageBasedLineConfigQuery {
	bimublcq.ctx.Limit = &limit
	return bimublcq
}

// Offset to start from.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) Offset(offset int) *BillingInvoiceManualUsageBasedLineConfigQuery {
	bimublcq.ctx.Offset = &offset
	return bimublcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) Unique(unique bool) *BillingInvoiceManualUsageBasedLineConfigQuery {
	bimublcq.ctx.Unique = &unique
	return bimublcq
}

// Order specifies how the records should be ordered.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) Order(o ...billinginvoicemanualusagebasedlineconfig.OrderOption) *BillingInvoiceManualUsageBasedLineConfigQuery {
	bimublcq.order = append(bimublcq.order, o...)
	return bimublcq
}

// First returns the first BillingInvoiceManualUsageBasedLineConfig entity from the query.
// Returns a *NotFoundError when no BillingInvoiceManualUsageBasedLineConfig was found.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) First(ctx context.Context) (*BillingInvoiceManualUsageBasedLineConfig, error) {
	nodes, err := bimublcq.Limit(1).All(setContextOp(ctx, bimublcq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{billinginvoicemanualusagebasedlineconfig.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) FirstX(ctx context.Context) *BillingInvoiceManualUsageBasedLineConfig {
	node, err := bimublcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BillingInvoiceManualUsageBasedLineConfig ID from the query.
// Returns a *NotFoundError when no BillingInvoiceManualUsageBasedLineConfig ID was found.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bimublcq.Limit(1).IDs(setContextOp(ctx, bimublcq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{billinginvoicemanualusagebasedlineconfig.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) FirstIDX(ctx context.Context) string {
	id, err := bimublcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BillingInvoiceManualUsageBasedLineConfig entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BillingInvoiceManualUsageBasedLineConfig entity is found.
// Returns a *NotFoundError when no BillingInvoiceManualUsageBasedLineConfig entities are found.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) Only(ctx context.Context) (*BillingInvoiceManualUsageBasedLineConfig, error) {
	nodes, err := bimublcq.Limit(2).All(setContextOp(ctx, bimublcq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{billinginvoicemanualusagebasedlineconfig.Label}
	default:
		return nil, &NotSingularError{billinginvoicemanualusagebasedlineconfig.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) OnlyX(ctx context.Context) *BillingInvoiceManualUsageBasedLineConfig {
	node, err := bimublcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BillingInvoiceManualUsageBasedLineConfig ID in the query.
// Returns a *NotSingularError when more than one BillingInvoiceManualUsageBasedLineConfig ID is found.
// Returns a *NotFoundError when no entities are found.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bimublcq.Limit(2).IDs(setContextOp(ctx, bimublcq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{billinginvoicemanualusagebasedlineconfig.Label}
	default:
		err = &NotSingularError{billinginvoicemanualusagebasedlineconfig.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) OnlyIDX(ctx context.Context) string {
	id, err := bimublcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BillingInvoiceManualUsageBasedLineConfigs.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) All(ctx context.Context) ([]*BillingInvoiceManualUsageBasedLineConfig, error) {
	ctx = setContextOp(ctx, bimublcq.ctx, ent.OpQueryAll)
	if err := bimublcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BillingInvoiceManualUsageBasedLineConfig, *BillingInvoiceManualUsageBasedLineConfigQuery]()
	return withInterceptors[[]*BillingInvoiceManualUsageBasedLineConfig](ctx, bimublcq, qr, bimublcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) AllX(ctx context.Context) []*BillingInvoiceManualUsageBasedLineConfig {
	nodes, err := bimublcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BillingInvoiceManualUsageBasedLineConfig IDs.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) IDs(ctx context.Context) (ids []string, err error) {
	if bimublcq.ctx.Unique == nil && bimublcq.path != nil {
		bimublcq.Unique(true)
	}
	ctx = setContextOp(ctx, bimublcq.ctx, ent.OpQueryIDs)
	if err = bimublcq.Select(billinginvoicemanualusagebasedlineconfig.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) IDsX(ctx context.Context) []string {
	ids, err := bimublcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bimublcq.ctx, ent.OpQueryCount)
	if err := bimublcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bimublcq, querierCount[*BillingInvoiceManualUsageBasedLineConfigQuery](), bimublcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) CountX(ctx context.Context) int {
	count, err := bimublcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bimublcq.ctx, ent.OpQueryExist)
	switch _, err := bimublcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) ExistX(ctx context.Context) bool {
	exist, err := bimublcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BillingInvoiceManualUsageBasedLineConfigQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) Clone() *BillingInvoiceManualUsageBasedLineConfigQuery {
	if bimublcq == nil {
		return nil
	}
	return &BillingInvoiceManualUsageBasedLineConfigQuery{
		config:     bimublcq.config,
		ctx:        bimublcq.ctx.Clone(),
		order:      append([]billinginvoicemanualusagebasedlineconfig.OrderOption{}, bimublcq.order...),
		inters:     append([]Interceptor{}, bimublcq.inters...),
		predicates: append([]predicate.BillingInvoiceManualUsageBasedLineConfig{}, bimublcq.predicates...),
		// clone intermediate query.
		sql:  bimublcq.sql.Clone(),
		path: bimublcq.path,
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
//	client.BillingInvoiceManualUsageBasedLineConfig.Query().
//		GroupBy(billinginvoicemanualusagebasedlineconfig.FieldNamespace).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) GroupBy(field string, fields ...string) *BillingInvoiceManualUsageBasedLineConfigGroupBy {
	bimublcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BillingInvoiceManualUsageBasedLineConfigGroupBy{build: bimublcq}
	grbuild.flds = &bimublcq.ctx.Fields
	grbuild.label = billinginvoicemanualusagebasedlineconfig.Label
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
//	client.BillingInvoiceManualUsageBasedLineConfig.Query().
//		Select(billinginvoicemanualusagebasedlineconfig.FieldNamespace).
//		Scan(ctx, &v)
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) Select(fields ...string) *BillingInvoiceManualUsageBasedLineConfigSelect {
	bimublcq.ctx.Fields = append(bimublcq.ctx.Fields, fields...)
	sbuild := &BillingInvoiceManualUsageBasedLineConfigSelect{BillingInvoiceManualUsageBasedLineConfigQuery: bimublcq}
	sbuild.label = billinginvoicemanualusagebasedlineconfig.Label
	sbuild.flds, sbuild.scan = &bimublcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BillingInvoiceManualUsageBasedLineConfigSelect configured with the given aggregations.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) Aggregate(fns ...AggregateFunc) *BillingInvoiceManualUsageBasedLineConfigSelect {
	return bimublcq.Select().Aggregate(fns...)
}

func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bimublcq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bimublcq); err != nil {
				return err
			}
		}
	}
	for _, f := range bimublcq.ctx.Fields {
		if !billinginvoicemanualusagebasedlineconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if bimublcq.path != nil {
		prev, err := bimublcq.path(ctx)
		if err != nil {
			return err
		}
		bimublcq.sql = prev
	}
	return nil
}

func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BillingInvoiceManualUsageBasedLineConfig, error) {
	var (
		nodes = []*BillingInvoiceManualUsageBasedLineConfig{}
		_spec = bimublcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BillingInvoiceManualUsageBasedLineConfig).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BillingInvoiceManualUsageBasedLineConfig{config: bimublcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(bimublcq.modifiers) > 0 {
		_spec.Modifiers = bimublcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bimublcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bimublcq.querySpec()
	if len(bimublcq.modifiers) > 0 {
		_spec.Modifiers = bimublcq.modifiers
	}
	_spec.Node.Columns = bimublcq.ctx.Fields
	if len(bimublcq.ctx.Fields) > 0 {
		_spec.Unique = bimublcq.ctx.Unique != nil && *bimublcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bimublcq.driver, _spec)
}

func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(billinginvoicemanualusagebasedlineconfig.Table, billinginvoicemanualusagebasedlineconfig.Columns, sqlgraph.NewFieldSpec(billinginvoicemanualusagebasedlineconfig.FieldID, field.TypeString))
	_spec.From = bimublcq.sql
	if unique := bimublcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bimublcq.path != nil {
		_spec.Unique = true
	}
	if fields := bimublcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, billinginvoicemanualusagebasedlineconfig.FieldID)
		for i := range fields {
			if fields[i] != billinginvoicemanualusagebasedlineconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bimublcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bimublcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bimublcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bimublcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bimublcq.driver.Dialect())
	t1 := builder.Table(billinginvoicemanualusagebasedlineconfig.Table)
	columns := bimublcq.ctx.Fields
	if len(columns) == 0 {
		columns = billinginvoicemanualusagebasedlineconfig.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bimublcq.sql != nil {
		selector = bimublcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bimublcq.ctx.Unique != nil && *bimublcq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range bimublcq.modifiers {
		m(selector)
	}
	for _, p := range bimublcq.predicates {
		p(selector)
	}
	for _, p := range bimublcq.order {
		p(selector)
	}
	if offset := bimublcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bimublcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) ForUpdate(opts ...sql.LockOption) *BillingInvoiceManualUsageBasedLineConfigQuery {
	if bimublcq.driver.Dialect() == dialect.Postgres {
		bimublcq.Unique(false)
	}
	bimublcq.modifiers = append(bimublcq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return bimublcq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (bimublcq *BillingInvoiceManualUsageBasedLineConfigQuery) ForShare(opts ...sql.LockOption) *BillingInvoiceManualUsageBasedLineConfigQuery {
	if bimublcq.driver.Dialect() == dialect.Postgres {
		bimublcq.Unique(false)
	}
	bimublcq.modifiers = append(bimublcq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return bimublcq
}

// BillingInvoiceManualUsageBasedLineConfigGroupBy is the group-by builder for BillingInvoiceManualUsageBasedLineConfig entities.
type BillingInvoiceManualUsageBasedLineConfigGroupBy struct {
	selector
	build *BillingInvoiceManualUsageBasedLineConfigQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bimublcgb *BillingInvoiceManualUsageBasedLineConfigGroupBy) Aggregate(fns ...AggregateFunc) *BillingInvoiceManualUsageBasedLineConfigGroupBy {
	bimublcgb.fns = append(bimublcgb.fns, fns...)
	return bimublcgb
}

// Scan applies the selector query and scans the result into the given value.
func (bimublcgb *BillingInvoiceManualUsageBasedLineConfigGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bimublcgb.build.ctx, ent.OpQueryGroupBy)
	if err := bimublcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillingInvoiceManualUsageBasedLineConfigQuery, *BillingInvoiceManualUsageBasedLineConfigGroupBy](ctx, bimublcgb.build, bimublcgb, bimublcgb.build.inters, v)
}

func (bimublcgb *BillingInvoiceManualUsageBasedLineConfigGroupBy) sqlScan(ctx context.Context, root *BillingInvoiceManualUsageBasedLineConfigQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bimublcgb.fns))
	for _, fn := range bimublcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bimublcgb.flds)+len(bimublcgb.fns))
		for _, f := range *bimublcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bimublcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bimublcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BillingInvoiceManualUsageBasedLineConfigSelect is the builder for selecting fields of BillingInvoiceManualUsageBasedLineConfig entities.
type BillingInvoiceManualUsageBasedLineConfigSelect struct {
	*BillingInvoiceManualUsageBasedLineConfigQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bimublcs *BillingInvoiceManualUsageBasedLineConfigSelect) Aggregate(fns ...AggregateFunc) *BillingInvoiceManualUsageBasedLineConfigSelect {
	bimublcs.fns = append(bimublcs.fns, fns...)
	return bimublcs
}

// Scan applies the selector query and scans the result into the given value.
func (bimublcs *BillingInvoiceManualUsageBasedLineConfigSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bimublcs.ctx, ent.OpQuerySelect)
	if err := bimublcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillingInvoiceManualUsageBasedLineConfigQuery, *BillingInvoiceManualUsageBasedLineConfigSelect](ctx, bimublcs.BillingInvoiceManualUsageBasedLineConfigQuery, bimublcs, bimublcs.inters, v)
}

func (bimublcs *BillingInvoiceManualUsageBasedLineConfigSelect) sqlScan(ctx context.Context, root *BillingInvoiceManualUsageBasedLineConfigQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bimublcs.fns))
	for _, fn := range bimublcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bimublcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bimublcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
