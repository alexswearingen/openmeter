// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/openmeter/ent/db/appstripe"
	"github.com/openmeterio/openmeter/openmeter/ent/db/appstripecustomer"
	"github.com/openmeterio/openmeter/openmeter/ent/db/customer"
)

// AppStripeCustomerCreate is the builder for creating a AppStripeCustomer entity.
type AppStripeCustomerCreate struct {
	config
	mutation *AppStripeCustomerMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetNamespace sets the "namespace" field.
func (ascc *AppStripeCustomerCreate) SetNamespace(s string) *AppStripeCustomerCreate {
	ascc.mutation.SetNamespace(s)
	return ascc
}

// SetCreatedAt sets the "created_at" field.
func (ascc *AppStripeCustomerCreate) SetCreatedAt(t time.Time) *AppStripeCustomerCreate {
	ascc.mutation.SetCreatedAt(t)
	return ascc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ascc *AppStripeCustomerCreate) SetNillableCreatedAt(t *time.Time) *AppStripeCustomerCreate {
	if t != nil {
		ascc.SetCreatedAt(*t)
	}
	return ascc
}

// SetUpdatedAt sets the "updated_at" field.
func (ascc *AppStripeCustomerCreate) SetUpdatedAt(t time.Time) *AppStripeCustomerCreate {
	ascc.mutation.SetUpdatedAt(t)
	return ascc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ascc *AppStripeCustomerCreate) SetNillableUpdatedAt(t *time.Time) *AppStripeCustomerCreate {
	if t != nil {
		ascc.SetUpdatedAt(*t)
	}
	return ascc
}

// SetDeletedAt sets the "deleted_at" field.
func (ascc *AppStripeCustomerCreate) SetDeletedAt(t time.Time) *AppStripeCustomerCreate {
	ascc.mutation.SetDeletedAt(t)
	return ascc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ascc *AppStripeCustomerCreate) SetNillableDeletedAt(t *time.Time) *AppStripeCustomerCreate {
	if t != nil {
		ascc.SetDeletedAt(*t)
	}
	return ascc
}

// SetAppID sets the "app_id" field.
func (ascc *AppStripeCustomerCreate) SetAppID(s string) *AppStripeCustomerCreate {
	ascc.mutation.SetAppID(s)
	return ascc
}

// SetCustomerID sets the "customer_id" field.
func (ascc *AppStripeCustomerCreate) SetCustomerID(s string) *AppStripeCustomerCreate {
	ascc.mutation.SetCustomerID(s)
	return ascc
}

// SetStripeCustomerID sets the "stripe_customer_id" field.
func (ascc *AppStripeCustomerCreate) SetStripeCustomerID(s string) *AppStripeCustomerCreate {
	ascc.mutation.SetStripeCustomerID(s)
	return ascc
}

// SetStripeDefaultPaymentMethodID sets the "stripe_default_payment_method_id" field.
func (ascc *AppStripeCustomerCreate) SetStripeDefaultPaymentMethodID(s string) *AppStripeCustomerCreate {
	ascc.mutation.SetStripeDefaultPaymentMethodID(s)
	return ascc
}

// SetNillableStripeDefaultPaymentMethodID sets the "stripe_default_payment_method_id" field if the given value is not nil.
func (ascc *AppStripeCustomerCreate) SetNillableStripeDefaultPaymentMethodID(s *string) *AppStripeCustomerCreate {
	if s != nil {
		ascc.SetStripeDefaultPaymentMethodID(*s)
	}
	return ascc
}

// SetStripeAppID sets the "stripe_app" edge to the AppStripe entity by ID.
func (ascc *AppStripeCustomerCreate) SetStripeAppID(id string) *AppStripeCustomerCreate {
	ascc.mutation.SetStripeAppID(id)
	return ascc
}

// SetStripeApp sets the "stripe_app" edge to the AppStripe entity.
func (ascc *AppStripeCustomerCreate) SetStripeApp(a *AppStripe) *AppStripeCustomerCreate {
	return ascc.SetStripeAppID(a.ID)
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (ascc *AppStripeCustomerCreate) SetCustomer(c *Customer) *AppStripeCustomerCreate {
	return ascc.SetCustomerID(c.ID)
}

// Mutation returns the AppStripeCustomerMutation object of the builder.
func (ascc *AppStripeCustomerCreate) Mutation() *AppStripeCustomerMutation {
	return ascc.mutation
}

// Save creates the AppStripeCustomer in the database.
func (ascc *AppStripeCustomerCreate) Save(ctx context.Context) (*AppStripeCustomer, error) {
	ascc.defaults()
	return withHooks(ctx, ascc.sqlSave, ascc.mutation, ascc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ascc *AppStripeCustomerCreate) SaveX(ctx context.Context) *AppStripeCustomer {
	v, err := ascc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ascc *AppStripeCustomerCreate) Exec(ctx context.Context) error {
	_, err := ascc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ascc *AppStripeCustomerCreate) ExecX(ctx context.Context) {
	if err := ascc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ascc *AppStripeCustomerCreate) defaults() {
	if _, ok := ascc.mutation.CreatedAt(); !ok {
		v := appstripecustomer.DefaultCreatedAt()
		ascc.mutation.SetCreatedAt(v)
	}
	if _, ok := ascc.mutation.UpdatedAt(); !ok {
		v := appstripecustomer.DefaultUpdatedAt()
		ascc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ascc *AppStripeCustomerCreate) check() error {
	if _, ok := ascc.mutation.Namespace(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`db: missing required field "AppStripeCustomer.namespace"`)}
	}
	if v, ok := ascc.mutation.Namespace(); ok {
		if err := appstripecustomer.NamespaceValidator(v); err != nil {
			return &ValidationError{Name: "namespace", err: fmt.Errorf(`db: validator failed for field "AppStripeCustomer.namespace": %w`, err)}
		}
	}
	if _, ok := ascc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "AppStripeCustomer.created_at"`)}
	}
	if _, ok := ascc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`db: missing required field "AppStripeCustomer.updated_at"`)}
	}
	if _, ok := ascc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`db: missing required field "AppStripeCustomer.app_id"`)}
	}
	if v, ok := ascc.mutation.AppID(); ok {
		if err := appstripecustomer.AppIDValidator(v); err != nil {
			return &ValidationError{Name: "app_id", err: fmt.Errorf(`db: validator failed for field "AppStripeCustomer.app_id": %w`, err)}
		}
	}
	if _, ok := ascc.mutation.CustomerID(); !ok {
		return &ValidationError{Name: "customer_id", err: errors.New(`db: missing required field "AppStripeCustomer.customer_id"`)}
	}
	if v, ok := ascc.mutation.CustomerID(); ok {
		if err := appstripecustomer.CustomerIDValidator(v); err != nil {
			return &ValidationError{Name: "customer_id", err: fmt.Errorf(`db: validator failed for field "AppStripeCustomer.customer_id": %w`, err)}
		}
	}
	if _, ok := ascc.mutation.StripeCustomerID(); !ok {
		return &ValidationError{Name: "stripe_customer_id", err: errors.New(`db: missing required field "AppStripeCustomer.stripe_customer_id"`)}
	}
	if v, ok := ascc.mutation.StripeCustomerID(); ok {
		if err := appstripecustomer.StripeCustomerIDValidator(v); err != nil {
			return &ValidationError{Name: "stripe_customer_id", err: fmt.Errorf(`db: validator failed for field "AppStripeCustomer.stripe_customer_id": %w`, err)}
		}
	}
	if len(ascc.mutation.StripeAppIDs()) == 0 {
		return &ValidationError{Name: "stripe_app", err: errors.New(`db: missing required edge "AppStripeCustomer.stripe_app"`)}
	}
	if len(ascc.mutation.CustomerIDs()) == 0 {
		return &ValidationError{Name: "customer", err: errors.New(`db: missing required edge "AppStripeCustomer.customer"`)}
	}
	return nil
}

func (ascc *AppStripeCustomerCreate) sqlSave(ctx context.Context) (*AppStripeCustomer, error) {
	if err := ascc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ascc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ascc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ascc.mutation.id = &_node.ID
	ascc.mutation.done = true
	return _node, nil
}

func (ascc *AppStripeCustomerCreate) createSpec() (*AppStripeCustomer, *sqlgraph.CreateSpec) {
	var (
		_node = &AppStripeCustomer{config: ascc.config}
		_spec = sqlgraph.NewCreateSpec(appstripecustomer.Table, sqlgraph.NewFieldSpec(appstripecustomer.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ascc.conflict
	if value, ok := ascc.mutation.Namespace(); ok {
		_spec.SetField(appstripecustomer.FieldNamespace, field.TypeString, value)
		_node.Namespace = value
	}
	if value, ok := ascc.mutation.CreatedAt(); ok {
		_spec.SetField(appstripecustomer.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ascc.mutation.UpdatedAt(); ok {
		_spec.SetField(appstripecustomer.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ascc.mutation.DeletedAt(); ok {
		_spec.SetField(appstripecustomer.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := ascc.mutation.StripeCustomerID(); ok {
		_spec.SetField(appstripecustomer.FieldStripeCustomerID, field.TypeString, value)
		_node.StripeCustomerID = value
	}
	if value, ok := ascc.mutation.StripeDefaultPaymentMethodID(); ok {
		_spec.SetField(appstripecustomer.FieldStripeDefaultPaymentMethodID, field.TypeString, value)
		_node.StripeDefaultPaymentMethodID = &value
	}
	if nodes := ascc.mutation.StripeAppIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appstripecustomer.StripeAppTable,
			Columns: []string{appstripecustomer.StripeAppColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appstripe.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AppID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ascc.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   appstripecustomer.CustomerTable,
			Columns: []string{appstripecustomer.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CustomerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppStripeCustomer.Create().
//		SetNamespace(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppStripeCustomerUpsert) {
//			SetNamespace(v+v).
//		}).
//		Exec(ctx)
func (ascc *AppStripeCustomerCreate) OnConflict(opts ...sql.ConflictOption) *AppStripeCustomerUpsertOne {
	ascc.conflict = opts
	return &AppStripeCustomerUpsertOne{
		create: ascc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppStripeCustomer.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ascc *AppStripeCustomerCreate) OnConflictColumns(columns ...string) *AppStripeCustomerUpsertOne {
	ascc.conflict = append(ascc.conflict, sql.ConflictColumns(columns...))
	return &AppStripeCustomerUpsertOne{
		create: ascc,
	}
}

type (
	// AppStripeCustomerUpsertOne is the builder for "upsert"-ing
	//  one AppStripeCustomer node.
	AppStripeCustomerUpsertOne struct {
		create *AppStripeCustomerCreate
	}

	// AppStripeCustomerUpsert is the "OnConflict" setter.
	AppStripeCustomerUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *AppStripeCustomerUpsert) SetUpdatedAt(v time.Time) *AppStripeCustomerUpsert {
	u.Set(appstripecustomer.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppStripeCustomerUpsert) UpdateUpdatedAt() *AppStripeCustomerUpsert {
	u.SetExcluded(appstripecustomer.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppStripeCustomerUpsert) SetDeletedAt(v time.Time) *AppStripeCustomerUpsert {
	u.Set(appstripecustomer.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppStripeCustomerUpsert) UpdateDeletedAt() *AppStripeCustomerUpsert {
	u.SetExcluded(appstripecustomer.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *AppStripeCustomerUpsert) ClearDeletedAt() *AppStripeCustomerUpsert {
	u.SetNull(appstripecustomer.FieldDeletedAt)
	return u
}

// SetStripeCustomerID sets the "stripe_customer_id" field.
func (u *AppStripeCustomerUpsert) SetStripeCustomerID(v string) *AppStripeCustomerUpsert {
	u.Set(appstripecustomer.FieldStripeCustomerID, v)
	return u
}

// UpdateStripeCustomerID sets the "stripe_customer_id" field to the value that was provided on create.
func (u *AppStripeCustomerUpsert) UpdateStripeCustomerID() *AppStripeCustomerUpsert {
	u.SetExcluded(appstripecustomer.FieldStripeCustomerID)
	return u
}

// SetStripeDefaultPaymentMethodID sets the "stripe_default_payment_method_id" field.
func (u *AppStripeCustomerUpsert) SetStripeDefaultPaymentMethodID(v string) *AppStripeCustomerUpsert {
	u.Set(appstripecustomer.FieldStripeDefaultPaymentMethodID, v)
	return u
}

// UpdateStripeDefaultPaymentMethodID sets the "stripe_default_payment_method_id" field to the value that was provided on create.
func (u *AppStripeCustomerUpsert) UpdateStripeDefaultPaymentMethodID() *AppStripeCustomerUpsert {
	u.SetExcluded(appstripecustomer.FieldStripeDefaultPaymentMethodID)
	return u
}

// ClearStripeDefaultPaymentMethodID clears the value of the "stripe_default_payment_method_id" field.
func (u *AppStripeCustomerUpsert) ClearStripeDefaultPaymentMethodID() *AppStripeCustomerUpsert {
	u.SetNull(appstripecustomer.FieldStripeDefaultPaymentMethodID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.AppStripeCustomer.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AppStripeCustomerUpsertOne) UpdateNewValues() *AppStripeCustomerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.Namespace(); exists {
			s.SetIgnore(appstripecustomer.FieldNamespace)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(appstripecustomer.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.AppID(); exists {
			s.SetIgnore(appstripecustomer.FieldAppID)
		}
		if _, exists := u.create.mutation.CustomerID(); exists {
			s.SetIgnore(appstripecustomer.FieldCustomerID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppStripeCustomer.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AppStripeCustomerUpsertOne) Ignore() *AppStripeCustomerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppStripeCustomerUpsertOne) DoNothing() *AppStripeCustomerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppStripeCustomerCreate.OnConflict
// documentation for more info.
func (u *AppStripeCustomerUpsertOne) Update(set func(*AppStripeCustomerUpsert)) *AppStripeCustomerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppStripeCustomerUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppStripeCustomerUpsertOne) SetUpdatedAt(v time.Time) *AppStripeCustomerUpsertOne {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppStripeCustomerUpsertOne) UpdateUpdatedAt() *AppStripeCustomerUpsertOne {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppStripeCustomerUpsertOne) SetDeletedAt(v time.Time) *AppStripeCustomerUpsertOne {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppStripeCustomerUpsertOne) UpdateDeletedAt() *AppStripeCustomerUpsertOne {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *AppStripeCustomerUpsertOne) ClearDeletedAt() *AppStripeCustomerUpsertOne {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.ClearDeletedAt()
	})
}

// SetStripeCustomerID sets the "stripe_customer_id" field.
func (u *AppStripeCustomerUpsertOne) SetStripeCustomerID(v string) *AppStripeCustomerUpsertOne {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.SetStripeCustomerID(v)
	})
}

// UpdateStripeCustomerID sets the "stripe_customer_id" field to the value that was provided on create.
func (u *AppStripeCustomerUpsertOne) UpdateStripeCustomerID() *AppStripeCustomerUpsertOne {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.UpdateStripeCustomerID()
	})
}

// SetStripeDefaultPaymentMethodID sets the "stripe_default_payment_method_id" field.
func (u *AppStripeCustomerUpsertOne) SetStripeDefaultPaymentMethodID(v string) *AppStripeCustomerUpsertOne {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.SetStripeDefaultPaymentMethodID(v)
	})
}

// UpdateStripeDefaultPaymentMethodID sets the "stripe_default_payment_method_id" field to the value that was provided on create.
func (u *AppStripeCustomerUpsertOne) UpdateStripeDefaultPaymentMethodID() *AppStripeCustomerUpsertOne {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.UpdateStripeDefaultPaymentMethodID()
	})
}

// ClearStripeDefaultPaymentMethodID clears the value of the "stripe_default_payment_method_id" field.
func (u *AppStripeCustomerUpsertOne) ClearStripeDefaultPaymentMethodID() *AppStripeCustomerUpsertOne {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.ClearStripeDefaultPaymentMethodID()
	})
}

// Exec executes the query.
func (u *AppStripeCustomerUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for AppStripeCustomerCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppStripeCustomerUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppStripeCustomerUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppStripeCustomerUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppStripeCustomerCreateBulk is the builder for creating many AppStripeCustomer entities in bulk.
type AppStripeCustomerCreateBulk struct {
	config
	err      error
	builders []*AppStripeCustomerCreate
	conflict []sql.ConflictOption
}

// Save creates the AppStripeCustomer entities in the database.
func (asccb *AppStripeCustomerCreateBulk) Save(ctx context.Context) ([]*AppStripeCustomer, error) {
	if asccb.err != nil {
		return nil, asccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(asccb.builders))
	nodes := make([]*AppStripeCustomer, len(asccb.builders))
	mutators := make([]Mutator, len(asccb.builders))
	for i := range asccb.builders {
		func(i int, root context.Context) {
			builder := asccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppStripeCustomerMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, asccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = asccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, asccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, asccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (asccb *AppStripeCustomerCreateBulk) SaveX(ctx context.Context) []*AppStripeCustomer {
	v, err := asccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (asccb *AppStripeCustomerCreateBulk) Exec(ctx context.Context) error {
	_, err := asccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asccb *AppStripeCustomerCreateBulk) ExecX(ctx context.Context) {
	if err := asccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppStripeCustomer.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppStripeCustomerUpsert) {
//			SetNamespace(v+v).
//		}).
//		Exec(ctx)
func (asccb *AppStripeCustomerCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppStripeCustomerUpsertBulk {
	asccb.conflict = opts
	return &AppStripeCustomerUpsertBulk{
		create: asccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppStripeCustomer.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (asccb *AppStripeCustomerCreateBulk) OnConflictColumns(columns ...string) *AppStripeCustomerUpsertBulk {
	asccb.conflict = append(asccb.conflict, sql.ConflictColumns(columns...))
	return &AppStripeCustomerUpsertBulk{
		create: asccb,
	}
}

// AppStripeCustomerUpsertBulk is the builder for "upsert"-ing
// a bulk of AppStripeCustomer nodes.
type AppStripeCustomerUpsertBulk struct {
	create *AppStripeCustomerCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppStripeCustomer.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AppStripeCustomerUpsertBulk) UpdateNewValues() *AppStripeCustomerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.Namespace(); exists {
				s.SetIgnore(appstripecustomer.FieldNamespace)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(appstripecustomer.FieldCreatedAt)
			}
			if _, exists := b.mutation.AppID(); exists {
				s.SetIgnore(appstripecustomer.FieldAppID)
			}
			if _, exists := b.mutation.CustomerID(); exists {
				s.SetIgnore(appstripecustomer.FieldCustomerID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppStripeCustomer.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AppStripeCustomerUpsertBulk) Ignore() *AppStripeCustomerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppStripeCustomerUpsertBulk) DoNothing() *AppStripeCustomerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppStripeCustomerCreateBulk.OnConflict
// documentation for more info.
func (u *AppStripeCustomerUpsertBulk) Update(set func(*AppStripeCustomerUpsert)) *AppStripeCustomerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppStripeCustomerUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppStripeCustomerUpsertBulk) SetUpdatedAt(v time.Time) *AppStripeCustomerUpsertBulk {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppStripeCustomerUpsertBulk) UpdateUpdatedAt() *AppStripeCustomerUpsertBulk {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppStripeCustomerUpsertBulk) SetDeletedAt(v time.Time) *AppStripeCustomerUpsertBulk {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppStripeCustomerUpsertBulk) UpdateDeletedAt() *AppStripeCustomerUpsertBulk {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *AppStripeCustomerUpsertBulk) ClearDeletedAt() *AppStripeCustomerUpsertBulk {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.ClearDeletedAt()
	})
}

// SetStripeCustomerID sets the "stripe_customer_id" field.
func (u *AppStripeCustomerUpsertBulk) SetStripeCustomerID(v string) *AppStripeCustomerUpsertBulk {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.SetStripeCustomerID(v)
	})
}

// UpdateStripeCustomerID sets the "stripe_customer_id" field to the value that was provided on create.
func (u *AppStripeCustomerUpsertBulk) UpdateStripeCustomerID() *AppStripeCustomerUpsertBulk {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.UpdateStripeCustomerID()
	})
}

// SetStripeDefaultPaymentMethodID sets the "stripe_default_payment_method_id" field.
func (u *AppStripeCustomerUpsertBulk) SetStripeDefaultPaymentMethodID(v string) *AppStripeCustomerUpsertBulk {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.SetStripeDefaultPaymentMethodID(v)
	})
}

// UpdateStripeDefaultPaymentMethodID sets the "stripe_default_payment_method_id" field to the value that was provided on create.
func (u *AppStripeCustomerUpsertBulk) UpdateStripeDefaultPaymentMethodID() *AppStripeCustomerUpsertBulk {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.UpdateStripeDefaultPaymentMethodID()
	})
}

// ClearStripeDefaultPaymentMethodID clears the value of the "stripe_default_payment_method_id" field.
func (u *AppStripeCustomerUpsertBulk) ClearStripeDefaultPaymentMethodID() *AppStripeCustomerUpsertBulk {
	return u.Update(func(s *AppStripeCustomerUpsert) {
		s.ClearStripeDefaultPaymentMethodID()
	})
}

// Exec executes the query.
func (u *AppStripeCustomerUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("db: OnConflict was set for builder %d. Set it on the AppStripeCustomerCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for AppStripeCustomerCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppStripeCustomerUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
