// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmeterio/openmeter/openmeter/ent/db/billingcustomerlock"
)

// BillingCustomerLockCreate is the builder for creating a BillingCustomerLock entity.
type BillingCustomerLockCreate struct {
	config
	mutation *BillingCustomerLockMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetNamespace sets the "namespace" field.
func (bclc *BillingCustomerLockCreate) SetNamespace(s string) *BillingCustomerLockCreate {
	bclc.mutation.SetNamespace(s)
	return bclc
}

// SetCustomerID sets the "customer_id" field.
func (bclc *BillingCustomerLockCreate) SetCustomerID(s string) *BillingCustomerLockCreate {
	bclc.mutation.SetCustomerID(s)
	return bclc
}

// SetID sets the "id" field.
func (bclc *BillingCustomerLockCreate) SetID(s string) *BillingCustomerLockCreate {
	bclc.mutation.SetID(s)
	return bclc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (bclc *BillingCustomerLockCreate) SetNillableID(s *string) *BillingCustomerLockCreate {
	if s != nil {
		bclc.SetID(*s)
	}
	return bclc
}

// Mutation returns the BillingCustomerLockMutation object of the builder.
func (bclc *BillingCustomerLockCreate) Mutation() *BillingCustomerLockMutation {
	return bclc.mutation
}

// Save creates the BillingCustomerLock in the database.
func (bclc *BillingCustomerLockCreate) Save(ctx context.Context) (*BillingCustomerLock, error) {
	bclc.defaults()
	return withHooks(ctx, bclc.sqlSave, bclc.mutation, bclc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bclc *BillingCustomerLockCreate) SaveX(ctx context.Context) *BillingCustomerLock {
	v, err := bclc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bclc *BillingCustomerLockCreate) Exec(ctx context.Context) error {
	_, err := bclc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bclc *BillingCustomerLockCreate) ExecX(ctx context.Context) {
	if err := bclc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bclc *BillingCustomerLockCreate) defaults() {
	if _, ok := bclc.mutation.ID(); !ok {
		v := billingcustomerlock.DefaultID()
		bclc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bclc *BillingCustomerLockCreate) check() error {
	if _, ok := bclc.mutation.Namespace(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`db: missing required field "BillingCustomerLock.namespace"`)}
	}
	if v, ok := bclc.mutation.Namespace(); ok {
		if err := billingcustomerlock.NamespaceValidator(v); err != nil {
			return &ValidationError{Name: "namespace", err: fmt.Errorf(`db: validator failed for field "BillingCustomerLock.namespace": %w`, err)}
		}
	}
	if _, ok := bclc.mutation.CustomerID(); !ok {
		return &ValidationError{Name: "customer_id", err: errors.New(`db: missing required field "BillingCustomerLock.customer_id"`)}
	}
	return nil
}

func (bclc *BillingCustomerLockCreate) sqlSave(ctx context.Context) (*BillingCustomerLock, error) {
	if err := bclc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bclc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bclc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected BillingCustomerLock.ID type: %T", _spec.ID.Value)
		}
	}
	bclc.mutation.id = &_node.ID
	bclc.mutation.done = true
	return _node, nil
}

func (bclc *BillingCustomerLockCreate) createSpec() (*BillingCustomerLock, *sqlgraph.CreateSpec) {
	var (
		_node = &BillingCustomerLock{config: bclc.config}
		_spec = sqlgraph.NewCreateSpec(billingcustomerlock.Table, sqlgraph.NewFieldSpec(billingcustomerlock.FieldID, field.TypeString))
	)
	_spec.OnConflict = bclc.conflict
	if id, ok := bclc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bclc.mutation.Namespace(); ok {
		_spec.SetField(billingcustomerlock.FieldNamespace, field.TypeString, value)
		_node.Namespace = value
	}
	if value, ok := bclc.mutation.CustomerID(); ok {
		_spec.SetField(billingcustomerlock.FieldCustomerID, field.TypeString, value)
		_node.CustomerID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.BillingCustomerLock.Create().
//		SetNamespace(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BillingCustomerLockUpsert) {
//			SetNamespace(v+v).
//		}).
//		Exec(ctx)
func (bclc *BillingCustomerLockCreate) OnConflict(opts ...sql.ConflictOption) *BillingCustomerLockUpsertOne {
	bclc.conflict = opts
	return &BillingCustomerLockUpsertOne{
		create: bclc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.BillingCustomerLock.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (bclc *BillingCustomerLockCreate) OnConflictColumns(columns ...string) *BillingCustomerLockUpsertOne {
	bclc.conflict = append(bclc.conflict, sql.ConflictColumns(columns...))
	return &BillingCustomerLockUpsertOne{
		create: bclc,
	}
}

type (
	// BillingCustomerLockUpsertOne is the builder for "upsert"-ing
	//  one BillingCustomerLock node.
	BillingCustomerLockUpsertOne struct {
		create *BillingCustomerLockCreate
	}

	// BillingCustomerLockUpsert is the "OnConflict" setter.
	BillingCustomerLockUpsert struct {
		*sql.UpdateSet
	}
)

// SetCustomerID sets the "customer_id" field.
func (u *BillingCustomerLockUpsert) SetCustomerID(v string) *BillingCustomerLockUpsert {
	u.Set(billingcustomerlock.FieldCustomerID, v)
	return u
}

// UpdateCustomerID sets the "customer_id" field to the value that was provided on create.
func (u *BillingCustomerLockUpsert) UpdateCustomerID() *BillingCustomerLockUpsert {
	u.SetExcluded(billingcustomerlock.FieldCustomerID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.BillingCustomerLock.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(billingcustomerlock.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BillingCustomerLockUpsertOne) UpdateNewValues() *BillingCustomerLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(billingcustomerlock.FieldID)
		}
		if _, exists := u.create.mutation.Namespace(); exists {
			s.SetIgnore(billingcustomerlock.FieldNamespace)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.BillingCustomerLock.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *BillingCustomerLockUpsertOne) Ignore() *BillingCustomerLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BillingCustomerLockUpsertOne) DoNothing() *BillingCustomerLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BillingCustomerLockCreate.OnConflict
// documentation for more info.
func (u *BillingCustomerLockUpsertOne) Update(set func(*BillingCustomerLockUpsert)) *BillingCustomerLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BillingCustomerLockUpsert{UpdateSet: update})
	}))
	return u
}

// SetCustomerID sets the "customer_id" field.
func (u *BillingCustomerLockUpsertOne) SetCustomerID(v string) *BillingCustomerLockUpsertOne {
	return u.Update(func(s *BillingCustomerLockUpsert) {
		s.SetCustomerID(v)
	})
}

// UpdateCustomerID sets the "customer_id" field to the value that was provided on create.
func (u *BillingCustomerLockUpsertOne) UpdateCustomerID() *BillingCustomerLockUpsertOne {
	return u.Update(func(s *BillingCustomerLockUpsert) {
		s.UpdateCustomerID()
	})
}

// Exec executes the query.
func (u *BillingCustomerLockUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for BillingCustomerLockCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BillingCustomerLockUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *BillingCustomerLockUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("db: BillingCustomerLockUpsertOne.ID is not supported by MySQL driver. Use BillingCustomerLockUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *BillingCustomerLockUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// BillingCustomerLockCreateBulk is the builder for creating many BillingCustomerLock entities in bulk.
type BillingCustomerLockCreateBulk struct {
	config
	err      error
	builders []*BillingCustomerLockCreate
	conflict []sql.ConflictOption
}

// Save creates the BillingCustomerLock entities in the database.
func (bclcb *BillingCustomerLockCreateBulk) Save(ctx context.Context) ([]*BillingCustomerLock, error) {
	if bclcb.err != nil {
		return nil, bclcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bclcb.builders))
	nodes := make([]*BillingCustomerLock, len(bclcb.builders))
	mutators := make([]Mutator, len(bclcb.builders))
	for i := range bclcb.builders {
		func(i int, root context.Context) {
			builder := bclcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BillingCustomerLockMutation)
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
					_, err = mutators[i+1].Mutate(root, bclcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = bclcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bclcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, bclcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bclcb *BillingCustomerLockCreateBulk) SaveX(ctx context.Context) []*BillingCustomerLock {
	v, err := bclcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bclcb *BillingCustomerLockCreateBulk) Exec(ctx context.Context) error {
	_, err := bclcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bclcb *BillingCustomerLockCreateBulk) ExecX(ctx context.Context) {
	if err := bclcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.BillingCustomerLock.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BillingCustomerLockUpsert) {
//			SetNamespace(v+v).
//		}).
//		Exec(ctx)
func (bclcb *BillingCustomerLockCreateBulk) OnConflict(opts ...sql.ConflictOption) *BillingCustomerLockUpsertBulk {
	bclcb.conflict = opts
	return &BillingCustomerLockUpsertBulk{
		create: bclcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.BillingCustomerLock.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (bclcb *BillingCustomerLockCreateBulk) OnConflictColumns(columns ...string) *BillingCustomerLockUpsertBulk {
	bclcb.conflict = append(bclcb.conflict, sql.ConflictColumns(columns...))
	return &BillingCustomerLockUpsertBulk{
		create: bclcb,
	}
}

// BillingCustomerLockUpsertBulk is the builder for "upsert"-ing
// a bulk of BillingCustomerLock nodes.
type BillingCustomerLockUpsertBulk struct {
	create *BillingCustomerLockCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.BillingCustomerLock.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(billingcustomerlock.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BillingCustomerLockUpsertBulk) UpdateNewValues() *BillingCustomerLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(billingcustomerlock.FieldID)
			}
			if _, exists := b.mutation.Namespace(); exists {
				s.SetIgnore(billingcustomerlock.FieldNamespace)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.BillingCustomerLock.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *BillingCustomerLockUpsertBulk) Ignore() *BillingCustomerLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BillingCustomerLockUpsertBulk) DoNothing() *BillingCustomerLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BillingCustomerLockCreateBulk.OnConflict
// documentation for more info.
func (u *BillingCustomerLockUpsertBulk) Update(set func(*BillingCustomerLockUpsert)) *BillingCustomerLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BillingCustomerLockUpsert{UpdateSet: update})
	}))
	return u
}

// SetCustomerID sets the "customer_id" field.
func (u *BillingCustomerLockUpsertBulk) SetCustomerID(v string) *BillingCustomerLockUpsertBulk {
	return u.Update(func(s *BillingCustomerLockUpsert) {
		s.SetCustomerID(v)
	})
}

// UpdateCustomerID sets the "customer_id" field to the value that was provided on create.
func (u *BillingCustomerLockUpsertBulk) UpdateCustomerID() *BillingCustomerLockUpsertBulk {
	return u.Update(func(s *BillingCustomerLockUpsert) {
		s.UpdateCustomerID()
	})
}

// Exec executes the query.
func (u *BillingCustomerLockUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("db: OnConflict was set for builder %d. Set it on the BillingCustomerLockCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for BillingCustomerLockCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BillingCustomerLockUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
