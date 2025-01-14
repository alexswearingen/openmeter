// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/openmeter/ent/db/billinginvoice"
	"github.com/openmeterio/openmeter/openmeter/ent/db/billinginvoicediscount"
	"github.com/openmeterio/openmeter/openmeter/ent/db/billinginvoiceline"
	"github.com/openmeterio/openmeter/openmeter/ent/db/predicate"
)

// BillingInvoiceDiscountQuery is the builder for querying BillingInvoiceDiscount entities.
type BillingInvoiceDiscountQuery struct {
	config
	ctx         *QueryContext
	order       []billinginvoicediscount.OrderOption
	inters      []Interceptor
	predicates  []predicate.BillingInvoiceDiscount
	withInvoice *BillingInvoiceQuery
	withLines   *BillingInvoiceLineQuery
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BillingInvoiceDiscountQuery builder.
func (bidq *BillingInvoiceDiscountQuery) Where(ps ...predicate.BillingInvoiceDiscount) *BillingInvoiceDiscountQuery {
	bidq.predicates = append(bidq.predicates, ps...)
	return bidq
}

// Limit the number of records to be returned by this query.
func (bidq *BillingInvoiceDiscountQuery) Limit(limit int) *BillingInvoiceDiscountQuery {
	bidq.ctx.Limit = &limit
	return bidq
}

// Offset to start from.
func (bidq *BillingInvoiceDiscountQuery) Offset(offset int) *BillingInvoiceDiscountQuery {
	bidq.ctx.Offset = &offset
	return bidq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bidq *BillingInvoiceDiscountQuery) Unique(unique bool) *BillingInvoiceDiscountQuery {
	bidq.ctx.Unique = &unique
	return bidq
}

// Order specifies how the records should be ordered.
func (bidq *BillingInvoiceDiscountQuery) Order(o ...billinginvoicediscount.OrderOption) *BillingInvoiceDiscountQuery {
	bidq.order = append(bidq.order, o...)
	return bidq
}

// QueryInvoice chains the current query on the "invoice" edge.
func (bidq *BillingInvoiceDiscountQuery) QueryInvoice() *BillingInvoiceQuery {
	query := (&BillingInvoiceClient{config: bidq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bidq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bidq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(billinginvoicediscount.Table, billinginvoicediscount.FieldID, selector),
			sqlgraph.To(billinginvoice.Table, billinginvoice.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, billinginvoicediscount.InvoiceTable, billinginvoicediscount.InvoiceColumn),
		)
		fromU = sqlgraph.SetNeighbors(bidq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLines chains the current query on the "lines" edge.
func (bidq *BillingInvoiceDiscountQuery) QueryLines() *BillingInvoiceLineQuery {
	query := (&BillingInvoiceLineClient{config: bidq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bidq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bidq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(billinginvoicediscount.Table, billinginvoicediscount.FieldID, selector),
			sqlgraph.To(billinginvoiceline.Table, billinginvoiceline.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, billinginvoicediscount.LinesTable, billinginvoicediscount.LinesColumn),
		)
		fromU = sqlgraph.SetNeighbors(bidq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BillingInvoiceDiscount entity from the query.
// Returns a *NotFoundError when no BillingInvoiceDiscount was found.
func (bidq *BillingInvoiceDiscountQuery) First(ctx context.Context) (*BillingInvoiceDiscount, error) {
	nodes, err := bidq.Limit(1).All(setContextOp(ctx, bidq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{billinginvoicediscount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bidq *BillingInvoiceDiscountQuery) FirstX(ctx context.Context) *BillingInvoiceDiscount {
	node, err := bidq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BillingInvoiceDiscount ID from the query.
// Returns a *NotFoundError when no BillingInvoiceDiscount ID was found.
func (bidq *BillingInvoiceDiscountQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bidq.Limit(1).IDs(setContextOp(ctx, bidq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{billinginvoicediscount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bidq *BillingInvoiceDiscountQuery) FirstIDX(ctx context.Context) string {
	id, err := bidq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BillingInvoiceDiscount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BillingInvoiceDiscount entity is found.
// Returns a *NotFoundError when no BillingInvoiceDiscount entities are found.
func (bidq *BillingInvoiceDiscountQuery) Only(ctx context.Context) (*BillingInvoiceDiscount, error) {
	nodes, err := bidq.Limit(2).All(setContextOp(ctx, bidq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{billinginvoicediscount.Label}
	default:
		return nil, &NotSingularError{billinginvoicediscount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bidq *BillingInvoiceDiscountQuery) OnlyX(ctx context.Context) *BillingInvoiceDiscount {
	node, err := bidq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BillingInvoiceDiscount ID in the query.
// Returns a *NotSingularError when more than one BillingInvoiceDiscount ID is found.
// Returns a *NotFoundError when no entities are found.
func (bidq *BillingInvoiceDiscountQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = bidq.Limit(2).IDs(setContextOp(ctx, bidq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{billinginvoicediscount.Label}
	default:
		err = &NotSingularError{billinginvoicediscount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bidq *BillingInvoiceDiscountQuery) OnlyIDX(ctx context.Context) string {
	id, err := bidq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BillingInvoiceDiscounts.
func (bidq *BillingInvoiceDiscountQuery) All(ctx context.Context) ([]*BillingInvoiceDiscount, error) {
	ctx = setContextOp(ctx, bidq.ctx, ent.OpQueryAll)
	if err := bidq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BillingInvoiceDiscount, *BillingInvoiceDiscountQuery]()
	return withInterceptors[[]*BillingInvoiceDiscount](ctx, bidq, qr, bidq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bidq *BillingInvoiceDiscountQuery) AllX(ctx context.Context) []*BillingInvoiceDiscount {
	nodes, err := bidq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BillingInvoiceDiscount IDs.
func (bidq *BillingInvoiceDiscountQuery) IDs(ctx context.Context) (ids []string, err error) {
	if bidq.ctx.Unique == nil && bidq.path != nil {
		bidq.Unique(true)
	}
	ctx = setContextOp(ctx, bidq.ctx, ent.OpQueryIDs)
	if err = bidq.Select(billinginvoicediscount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bidq *BillingInvoiceDiscountQuery) IDsX(ctx context.Context) []string {
	ids, err := bidq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bidq *BillingInvoiceDiscountQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bidq.ctx, ent.OpQueryCount)
	if err := bidq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bidq, querierCount[*BillingInvoiceDiscountQuery](), bidq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bidq *BillingInvoiceDiscountQuery) CountX(ctx context.Context) int {
	count, err := bidq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bidq *BillingInvoiceDiscountQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bidq.ctx, ent.OpQueryExist)
	switch _, err := bidq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bidq *BillingInvoiceDiscountQuery) ExistX(ctx context.Context) bool {
	exist, err := bidq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BillingInvoiceDiscountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bidq *BillingInvoiceDiscountQuery) Clone() *BillingInvoiceDiscountQuery {
	if bidq == nil {
		return nil
	}
	return &BillingInvoiceDiscountQuery{
		config:      bidq.config,
		ctx:         bidq.ctx.Clone(),
		order:       append([]billinginvoicediscount.OrderOption{}, bidq.order...),
		inters:      append([]Interceptor{}, bidq.inters...),
		predicates:  append([]predicate.BillingInvoiceDiscount{}, bidq.predicates...),
		withInvoice: bidq.withInvoice.Clone(),
		withLines:   bidq.withLines.Clone(),
		// clone intermediate query.
		sql:  bidq.sql.Clone(),
		path: bidq.path,
	}
}

// WithInvoice tells the query-builder to eager-load the nodes that are connected to
// the "invoice" edge. The optional arguments are used to configure the query builder of the edge.
func (bidq *BillingInvoiceDiscountQuery) WithInvoice(opts ...func(*BillingInvoiceQuery)) *BillingInvoiceDiscountQuery {
	query := (&BillingInvoiceClient{config: bidq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bidq.withInvoice = query
	return bidq
}

// WithLines tells the query-builder to eager-load the nodes that are connected to
// the "lines" edge. The optional arguments are used to configure the query builder of the edge.
func (bidq *BillingInvoiceDiscountQuery) WithLines(opts ...func(*BillingInvoiceLineQuery)) *BillingInvoiceDiscountQuery {
	query := (&BillingInvoiceLineClient{config: bidq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bidq.withLines = query
	return bidq
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
//	client.BillingInvoiceDiscount.Query().
//		GroupBy(billinginvoicediscount.FieldNamespace).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (bidq *BillingInvoiceDiscountQuery) GroupBy(field string, fields ...string) *BillingInvoiceDiscountGroupBy {
	bidq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BillingInvoiceDiscountGroupBy{build: bidq}
	grbuild.flds = &bidq.ctx.Fields
	grbuild.label = billinginvoicediscount.Label
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
//	client.BillingInvoiceDiscount.Query().
//		Select(billinginvoicediscount.FieldNamespace).
//		Scan(ctx, &v)
func (bidq *BillingInvoiceDiscountQuery) Select(fields ...string) *BillingInvoiceDiscountSelect {
	bidq.ctx.Fields = append(bidq.ctx.Fields, fields...)
	sbuild := &BillingInvoiceDiscountSelect{BillingInvoiceDiscountQuery: bidq}
	sbuild.label = billinginvoicediscount.Label
	sbuild.flds, sbuild.scan = &bidq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BillingInvoiceDiscountSelect configured with the given aggregations.
func (bidq *BillingInvoiceDiscountQuery) Aggregate(fns ...AggregateFunc) *BillingInvoiceDiscountSelect {
	return bidq.Select().Aggregate(fns...)
}

func (bidq *BillingInvoiceDiscountQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bidq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bidq); err != nil {
				return err
			}
		}
	}
	for _, f := range bidq.ctx.Fields {
		if !billinginvoicediscount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if bidq.path != nil {
		prev, err := bidq.path(ctx)
		if err != nil {
			return err
		}
		bidq.sql = prev
	}
	return nil
}

func (bidq *BillingInvoiceDiscountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BillingInvoiceDiscount, error) {
	var (
		nodes       = []*BillingInvoiceDiscount{}
		_spec       = bidq.querySpec()
		loadedTypes = [2]bool{
			bidq.withInvoice != nil,
			bidq.withLines != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BillingInvoiceDiscount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BillingInvoiceDiscount{config: bidq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(bidq.modifiers) > 0 {
		_spec.Modifiers = bidq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bidq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bidq.withInvoice; query != nil {
		if err := bidq.loadInvoice(ctx, query, nodes, nil,
			func(n *BillingInvoiceDiscount, e *BillingInvoice) { n.Edges.Invoice = e }); err != nil {
			return nil, err
		}
	}
	if query := bidq.withLines; query != nil {
		if err := bidq.loadLines(ctx, query, nodes,
			func(n *BillingInvoiceDiscount) { n.Edges.Lines = []*BillingInvoiceLine{} },
			func(n *BillingInvoiceDiscount, e *BillingInvoiceLine) { n.Edges.Lines = append(n.Edges.Lines, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bidq *BillingInvoiceDiscountQuery) loadInvoice(ctx context.Context, query *BillingInvoiceQuery, nodes []*BillingInvoiceDiscount, init func(*BillingInvoiceDiscount), assign func(*BillingInvoiceDiscount, *BillingInvoice)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*BillingInvoiceDiscount)
	for i := range nodes {
		fk := nodes[i].InvoiceID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(billinginvoice.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "invoice_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (bidq *BillingInvoiceDiscountQuery) loadLines(ctx context.Context, query *BillingInvoiceLineQuery, nodes []*BillingInvoiceDiscount, init func(*BillingInvoiceDiscount), assign func(*BillingInvoiceDiscount, *BillingInvoiceLine)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*BillingInvoiceDiscount)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.BillingInvoiceLine(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(billinginvoicediscount.LinesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.line_ids
		if fk == nil {
			return fmt.Errorf(`foreign-key "line_ids" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "line_ids" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (bidq *BillingInvoiceDiscountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bidq.querySpec()
	if len(bidq.modifiers) > 0 {
		_spec.Modifiers = bidq.modifiers
	}
	_spec.Node.Columns = bidq.ctx.Fields
	if len(bidq.ctx.Fields) > 0 {
		_spec.Unique = bidq.ctx.Unique != nil && *bidq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bidq.driver, _spec)
}

func (bidq *BillingInvoiceDiscountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(billinginvoicediscount.Table, billinginvoicediscount.Columns, sqlgraph.NewFieldSpec(billinginvoicediscount.FieldID, field.TypeString))
	_spec.From = bidq.sql
	if unique := bidq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bidq.path != nil {
		_spec.Unique = true
	}
	if fields := bidq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, billinginvoicediscount.FieldID)
		for i := range fields {
			if fields[i] != billinginvoicediscount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if bidq.withInvoice != nil {
			_spec.Node.AddColumnOnce(billinginvoicediscount.FieldInvoiceID)
		}
	}
	if ps := bidq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bidq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bidq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bidq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bidq *BillingInvoiceDiscountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bidq.driver.Dialect())
	t1 := builder.Table(billinginvoicediscount.Table)
	columns := bidq.ctx.Fields
	if len(columns) == 0 {
		columns = billinginvoicediscount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bidq.sql != nil {
		selector = bidq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bidq.ctx.Unique != nil && *bidq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range bidq.modifiers {
		m(selector)
	}
	for _, p := range bidq.predicates {
		p(selector)
	}
	for _, p := range bidq.order {
		p(selector)
	}
	if offset := bidq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bidq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (bidq *BillingInvoiceDiscountQuery) ForUpdate(opts ...sql.LockOption) *BillingInvoiceDiscountQuery {
	if bidq.driver.Dialect() == dialect.Postgres {
		bidq.Unique(false)
	}
	bidq.modifiers = append(bidq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return bidq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (bidq *BillingInvoiceDiscountQuery) ForShare(opts ...sql.LockOption) *BillingInvoiceDiscountQuery {
	if bidq.driver.Dialect() == dialect.Postgres {
		bidq.Unique(false)
	}
	bidq.modifiers = append(bidq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return bidq
}

// BillingInvoiceDiscountGroupBy is the group-by builder for BillingInvoiceDiscount entities.
type BillingInvoiceDiscountGroupBy struct {
	selector
	build *BillingInvoiceDiscountQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bidgb *BillingInvoiceDiscountGroupBy) Aggregate(fns ...AggregateFunc) *BillingInvoiceDiscountGroupBy {
	bidgb.fns = append(bidgb.fns, fns...)
	return bidgb
}

// Scan applies the selector query and scans the result into the given value.
func (bidgb *BillingInvoiceDiscountGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bidgb.build.ctx, ent.OpQueryGroupBy)
	if err := bidgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillingInvoiceDiscountQuery, *BillingInvoiceDiscountGroupBy](ctx, bidgb.build, bidgb, bidgb.build.inters, v)
}

func (bidgb *BillingInvoiceDiscountGroupBy) sqlScan(ctx context.Context, root *BillingInvoiceDiscountQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bidgb.fns))
	for _, fn := range bidgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bidgb.flds)+len(bidgb.fns))
		for _, f := range *bidgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bidgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bidgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BillingInvoiceDiscountSelect is the builder for selecting fields of BillingInvoiceDiscount entities.
type BillingInvoiceDiscountSelect struct {
	*BillingInvoiceDiscountQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bids *BillingInvoiceDiscountSelect) Aggregate(fns ...AggregateFunc) *BillingInvoiceDiscountSelect {
	bids.fns = append(bids.fns, fns...)
	return bids
}

// Scan applies the selector query and scans the result into the given value.
func (bids *BillingInvoiceDiscountSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bids.ctx, ent.OpQuerySelect)
	if err := bids.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BillingInvoiceDiscountQuery, *BillingInvoiceDiscountSelect](ctx, bids.BillingInvoiceDiscountQuery, bids, bids.inters, v)
}

func (bids *BillingInvoiceDiscountSelect) sqlScan(ctx context.Context, root *BillingInvoiceDiscountQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bids.fns))
	for _, fn := range bids.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bids.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bids.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
