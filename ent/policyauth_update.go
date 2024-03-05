// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sectran_admin/ent/policyauth"
	"sectran_admin/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PolicyAuthUpdate is the builder for updating PolicyAuth entities.
type PolicyAuthUpdate struct {
	config
	hooks    []Hook
	mutation *PolicyAuthMutation
}

// Where appends a list predicates to the PolicyAuthUpdate builder.
func (pau *PolicyAuthUpdate) Where(ps ...predicate.PolicyAuth) *PolicyAuthUpdate {
	pau.mutation.Where(ps...)
	return pau
}

// SetUpdatedAt sets the "updated_at" field.
func (pau *PolicyAuthUpdate) SetUpdatedAt(t time.Time) *PolicyAuthUpdate {
	pau.mutation.SetUpdatedAt(t)
	return pau
}

// SetName sets the "name" field.
func (pau *PolicyAuthUpdate) SetName(s string) *PolicyAuthUpdate {
	pau.mutation.SetName(s)
	return pau
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pau *PolicyAuthUpdate) SetNillableName(s *string) *PolicyAuthUpdate {
	if s != nil {
		pau.SetName(*s)
	}
	return pau
}

// SetPower sets the "power" field.
func (pau *PolicyAuthUpdate) SetPower(i int32) *PolicyAuthUpdate {
	pau.mutation.ResetPower()
	pau.mutation.SetPower(i)
	return pau
}

// SetNillablePower sets the "power" field if the given value is not nil.
func (pau *PolicyAuthUpdate) SetNillablePower(i *int32) *PolicyAuthUpdate {
	if i != nil {
		pau.SetPower(*i)
	}
	return pau
}

// AddPower adds i to the "power" field.
func (pau *PolicyAuthUpdate) AddPower(i int32) *PolicyAuthUpdate {
	pau.mutation.AddPower(i)
	return pau
}

// SetDepartmentID sets the "department_id" field.
func (pau *PolicyAuthUpdate) SetDepartmentID(u uint64) *PolicyAuthUpdate {
	pau.mutation.ResetDepartmentID()
	pau.mutation.SetDepartmentID(u)
	return pau
}

// SetNillableDepartmentID sets the "department_id" field if the given value is not nil.
func (pau *PolicyAuthUpdate) SetNillableDepartmentID(u *uint64) *PolicyAuthUpdate {
	if u != nil {
		pau.SetDepartmentID(*u)
	}
	return pau
}

// AddDepartmentID adds u to the "department_id" field.
func (pau *PolicyAuthUpdate) AddDepartmentID(u int64) *PolicyAuthUpdate {
	pau.mutation.AddDepartmentID(u)
	return pau
}

// ClearDepartmentID clears the value of the "department_id" field.
func (pau *PolicyAuthUpdate) ClearDepartmentID() *PolicyAuthUpdate {
	pau.mutation.ClearDepartmentID()
	return pau
}

// SetUsers sets the "users" field.
func (pau *PolicyAuthUpdate) SetUsers(s string) *PolicyAuthUpdate {
	pau.mutation.SetUsers(s)
	return pau
}

// SetNillableUsers sets the "users" field if the given value is not nil.
func (pau *PolicyAuthUpdate) SetNillableUsers(s *string) *PolicyAuthUpdate {
	if s != nil {
		pau.SetUsers(*s)
	}
	return pau
}

// SetAccounts sets the "accounts" field.
func (pau *PolicyAuthUpdate) SetAccounts(s string) *PolicyAuthUpdate {
	pau.mutation.SetAccounts(s)
	return pau
}

// SetNillableAccounts sets the "accounts" field if the given value is not nil.
func (pau *PolicyAuthUpdate) SetNillableAccounts(s *string) *PolicyAuthUpdate {
	if s != nil {
		pau.SetAccounts(*s)
	}
	return pau
}

// Mutation returns the PolicyAuthMutation object of the builder.
func (pau *PolicyAuthUpdate) Mutation() *PolicyAuthMutation {
	return pau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pau *PolicyAuthUpdate) Save(ctx context.Context) (int, error) {
	pau.defaults()
	return withHooks(ctx, pau.sqlSave, pau.mutation, pau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pau *PolicyAuthUpdate) SaveX(ctx context.Context) int {
	affected, err := pau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pau *PolicyAuthUpdate) Exec(ctx context.Context) error {
	_, err := pau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pau *PolicyAuthUpdate) ExecX(ctx context.Context) {
	if err := pau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pau *PolicyAuthUpdate) defaults() {
	if _, ok := pau.mutation.UpdatedAt(); !ok {
		v := policyauth.UpdateDefaultUpdatedAt()
		pau.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pau *PolicyAuthUpdate) check() error {
	if v, ok := pau.mutation.Name(); ok {
		if err := policyauth.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "PolicyAuth.name": %w`, err)}
		}
	}
	if v, ok := pau.mutation.Power(); ok {
		if err := policyauth.PowerValidator(v); err != nil {
			return &ValidationError{Name: "power", err: fmt.Errorf(`ent: validator failed for field "PolicyAuth.power": %w`, err)}
		}
	}
	if v, ok := pau.mutation.DepartmentID(); ok {
		if err := policyauth.DepartmentIDValidator(v); err != nil {
			return &ValidationError{Name: "department_id", err: fmt.Errorf(`ent: validator failed for field "PolicyAuth.department_id": %w`, err)}
		}
	}
	if v, ok := pau.mutation.Users(); ok {
		if err := policyauth.UsersValidator(v); err != nil {
			return &ValidationError{Name: "users", err: fmt.Errorf(`ent: validator failed for field "PolicyAuth.users": %w`, err)}
		}
	}
	if v, ok := pau.mutation.Accounts(); ok {
		if err := policyauth.AccountsValidator(v); err != nil {
			return &ValidationError{Name: "accounts", err: fmt.Errorf(`ent: validator failed for field "PolicyAuth.accounts": %w`, err)}
		}
	}
	return nil
}

func (pau *PolicyAuthUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(policyauth.Table, policyauth.Columns, sqlgraph.NewFieldSpec(policyauth.FieldID, field.TypeUint64))
	if ps := pau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pau.mutation.UpdatedAt(); ok {
		_spec.SetField(policyauth.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pau.mutation.Name(); ok {
		_spec.SetField(policyauth.FieldName, field.TypeString, value)
	}
	if value, ok := pau.mutation.Power(); ok {
		_spec.SetField(policyauth.FieldPower, field.TypeInt32, value)
	}
	if value, ok := pau.mutation.AddedPower(); ok {
		_spec.AddField(policyauth.FieldPower, field.TypeInt32, value)
	}
	if value, ok := pau.mutation.DepartmentID(); ok {
		_spec.SetField(policyauth.FieldDepartmentID, field.TypeUint64, value)
	}
	if value, ok := pau.mutation.AddedDepartmentID(); ok {
		_spec.AddField(policyauth.FieldDepartmentID, field.TypeUint64, value)
	}
	if pau.mutation.DepartmentIDCleared() {
		_spec.ClearField(policyauth.FieldDepartmentID, field.TypeUint64)
	}
	if value, ok := pau.mutation.Users(); ok {
		_spec.SetField(policyauth.FieldUsers, field.TypeString, value)
	}
	if value, ok := pau.mutation.Accounts(); ok {
		_spec.SetField(policyauth.FieldAccounts, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{policyauth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pau.mutation.done = true
	return n, nil
}

// PolicyAuthUpdateOne is the builder for updating a single PolicyAuth entity.
type PolicyAuthUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PolicyAuthMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (pauo *PolicyAuthUpdateOne) SetUpdatedAt(t time.Time) *PolicyAuthUpdateOne {
	pauo.mutation.SetUpdatedAt(t)
	return pauo
}

// SetName sets the "name" field.
func (pauo *PolicyAuthUpdateOne) SetName(s string) *PolicyAuthUpdateOne {
	pauo.mutation.SetName(s)
	return pauo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pauo *PolicyAuthUpdateOne) SetNillableName(s *string) *PolicyAuthUpdateOne {
	if s != nil {
		pauo.SetName(*s)
	}
	return pauo
}

// SetPower sets the "power" field.
func (pauo *PolicyAuthUpdateOne) SetPower(i int32) *PolicyAuthUpdateOne {
	pauo.mutation.ResetPower()
	pauo.mutation.SetPower(i)
	return pauo
}

// SetNillablePower sets the "power" field if the given value is not nil.
func (pauo *PolicyAuthUpdateOne) SetNillablePower(i *int32) *PolicyAuthUpdateOne {
	if i != nil {
		pauo.SetPower(*i)
	}
	return pauo
}

// AddPower adds i to the "power" field.
func (pauo *PolicyAuthUpdateOne) AddPower(i int32) *PolicyAuthUpdateOne {
	pauo.mutation.AddPower(i)
	return pauo
}

// SetDepartmentID sets the "department_id" field.
func (pauo *PolicyAuthUpdateOne) SetDepartmentID(u uint64) *PolicyAuthUpdateOne {
	pauo.mutation.ResetDepartmentID()
	pauo.mutation.SetDepartmentID(u)
	return pauo
}

// SetNillableDepartmentID sets the "department_id" field if the given value is not nil.
func (pauo *PolicyAuthUpdateOne) SetNillableDepartmentID(u *uint64) *PolicyAuthUpdateOne {
	if u != nil {
		pauo.SetDepartmentID(*u)
	}
	return pauo
}

// AddDepartmentID adds u to the "department_id" field.
func (pauo *PolicyAuthUpdateOne) AddDepartmentID(u int64) *PolicyAuthUpdateOne {
	pauo.mutation.AddDepartmentID(u)
	return pauo
}

// ClearDepartmentID clears the value of the "department_id" field.
func (pauo *PolicyAuthUpdateOne) ClearDepartmentID() *PolicyAuthUpdateOne {
	pauo.mutation.ClearDepartmentID()
	return pauo
}

// SetUsers sets the "users" field.
func (pauo *PolicyAuthUpdateOne) SetUsers(s string) *PolicyAuthUpdateOne {
	pauo.mutation.SetUsers(s)
	return pauo
}

// SetNillableUsers sets the "users" field if the given value is not nil.
func (pauo *PolicyAuthUpdateOne) SetNillableUsers(s *string) *PolicyAuthUpdateOne {
	if s != nil {
		pauo.SetUsers(*s)
	}
	return pauo
}

// SetAccounts sets the "accounts" field.
func (pauo *PolicyAuthUpdateOne) SetAccounts(s string) *PolicyAuthUpdateOne {
	pauo.mutation.SetAccounts(s)
	return pauo
}

// SetNillableAccounts sets the "accounts" field if the given value is not nil.
func (pauo *PolicyAuthUpdateOne) SetNillableAccounts(s *string) *PolicyAuthUpdateOne {
	if s != nil {
		pauo.SetAccounts(*s)
	}
	return pauo
}

// Mutation returns the PolicyAuthMutation object of the builder.
func (pauo *PolicyAuthUpdateOne) Mutation() *PolicyAuthMutation {
	return pauo.mutation
}

// Where appends a list predicates to the PolicyAuthUpdate builder.
func (pauo *PolicyAuthUpdateOne) Where(ps ...predicate.PolicyAuth) *PolicyAuthUpdateOne {
	pauo.mutation.Where(ps...)
	return pauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pauo *PolicyAuthUpdateOne) Select(field string, fields ...string) *PolicyAuthUpdateOne {
	pauo.fields = append([]string{field}, fields...)
	return pauo
}

// Save executes the query and returns the updated PolicyAuth entity.
func (pauo *PolicyAuthUpdateOne) Save(ctx context.Context) (*PolicyAuth, error) {
	pauo.defaults()
	return withHooks(ctx, pauo.sqlSave, pauo.mutation, pauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pauo *PolicyAuthUpdateOne) SaveX(ctx context.Context) *PolicyAuth {
	node, err := pauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pauo *PolicyAuthUpdateOne) Exec(ctx context.Context) error {
	_, err := pauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pauo *PolicyAuthUpdateOne) ExecX(ctx context.Context) {
	if err := pauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pauo *PolicyAuthUpdateOne) defaults() {
	if _, ok := pauo.mutation.UpdatedAt(); !ok {
		v := policyauth.UpdateDefaultUpdatedAt()
		pauo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pauo *PolicyAuthUpdateOne) check() error {
	if v, ok := pauo.mutation.Name(); ok {
		if err := policyauth.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "PolicyAuth.name": %w`, err)}
		}
	}
	if v, ok := pauo.mutation.Power(); ok {
		if err := policyauth.PowerValidator(v); err != nil {
			return &ValidationError{Name: "power", err: fmt.Errorf(`ent: validator failed for field "PolicyAuth.power": %w`, err)}
		}
	}
	if v, ok := pauo.mutation.DepartmentID(); ok {
		if err := policyauth.DepartmentIDValidator(v); err != nil {
			return &ValidationError{Name: "department_id", err: fmt.Errorf(`ent: validator failed for field "PolicyAuth.department_id": %w`, err)}
		}
	}
	if v, ok := pauo.mutation.Users(); ok {
		if err := policyauth.UsersValidator(v); err != nil {
			return &ValidationError{Name: "users", err: fmt.Errorf(`ent: validator failed for field "PolicyAuth.users": %w`, err)}
		}
	}
	if v, ok := pauo.mutation.Accounts(); ok {
		if err := policyauth.AccountsValidator(v); err != nil {
			return &ValidationError{Name: "accounts", err: fmt.Errorf(`ent: validator failed for field "PolicyAuth.accounts": %w`, err)}
		}
	}
	return nil
}

func (pauo *PolicyAuthUpdateOne) sqlSave(ctx context.Context) (_node *PolicyAuth, err error) {
	if err := pauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(policyauth.Table, policyauth.Columns, sqlgraph.NewFieldSpec(policyauth.FieldID, field.TypeUint64))
	id, ok := pauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PolicyAuth.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, policyauth.FieldID)
		for _, f := range fields {
			if !policyauth.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != policyauth.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pauo.mutation.UpdatedAt(); ok {
		_spec.SetField(policyauth.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pauo.mutation.Name(); ok {
		_spec.SetField(policyauth.FieldName, field.TypeString, value)
	}
	if value, ok := pauo.mutation.Power(); ok {
		_spec.SetField(policyauth.FieldPower, field.TypeInt32, value)
	}
	if value, ok := pauo.mutation.AddedPower(); ok {
		_spec.AddField(policyauth.FieldPower, field.TypeInt32, value)
	}
	if value, ok := pauo.mutation.DepartmentID(); ok {
		_spec.SetField(policyauth.FieldDepartmentID, field.TypeUint64, value)
	}
	if value, ok := pauo.mutation.AddedDepartmentID(); ok {
		_spec.AddField(policyauth.FieldDepartmentID, field.TypeUint64, value)
	}
	if pauo.mutation.DepartmentIDCleared() {
		_spec.ClearField(policyauth.FieldDepartmentID, field.TypeUint64)
	}
	if value, ok := pauo.mutation.Users(); ok {
		_spec.SetField(policyauth.FieldUsers, field.TypeString, value)
	}
	if value, ok := pauo.mutation.Accounts(); ok {
		_spec.SetField(policyauth.FieldAccounts, field.TypeString, value)
	}
	_node = &PolicyAuth{config: pauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{policyauth.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pauo.mutation.done = true
	return _node, nil
}
