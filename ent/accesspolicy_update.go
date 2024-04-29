// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sectran_admin/ent/accesspolicy"
	"sectran_admin/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccessPolicyUpdate is the builder for updating AccessPolicy entities.
type AccessPolicyUpdate struct {
	config
	hooks    []Hook
	mutation *AccessPolicyMutation
}

// Where appends a list predicates to the AccessPolicyUpdate builder.
func (apu *AccessPolicyUpdate) Where(ps ...predicate.AccessPolicy) *AccessPolicyUpdate {
	apu.mutation.Where(ps...)
	return apu
}

// SetUpdatedAt sets the "updated_at" field.
func (apu *AccessPolicyUpdate) SetUpdatedAt(t time.Time) *AccessPolicyUpdate {
	apu.mutation.SetUpdatedAt(t)
	return apu
}

// SetName sets the "name" field.
func (apu *AccessPolicyUpdate) SetName(s string) *AccessPolicyUpdate {
	apu.mutation.SetName(s)
	return apu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (apu *AccessPolicyUpdate) SetNillableName(s *string) *AccessPolicyUpdate {
	if s != nil {
		apu.SetName(*s)
	}
	return apu
}

// SetPower sets the "power" field.
func (apu *AccessPolicyUpdate) SetPower(i int32) *AccessPolicyUpdate {
	apu.mutation.ResetPower()
	apu.mutation.SetPower(i)
	return apu
}

// SetNillablePower sets the "power" field if the given value is not nil.
func (apu *AccessPolicyUpdate) SetNillablePower(i *int32) *AccessPolicyUpdate {
	if i != nil {
		apu.SetPower(*i)
	}
	return apu
}

// AddPower adds i to the "power" field.
func (apu *AccessPolicyUpdate) AddPower(i int32) *AccessPolicyUpdate {
	apu.mutation.AddPower(i)
	return apu
}

// ClearPower clears the value of the "power" field.
func (apu *AccessPolicyUpdate) ClearPower() *AccessPolicyUpdate {
	apu.mutation.ClearPower()
	return apu
}

// SetDepartmentID sets the "department_id" field.
func (apu *AccessPolicyUpdate) SetDepartmentID(u uint64) *AccessPolicyUpdate {
	apu.mutation.ResetDepartmentID()
	apu.mutation.SetDepartmentID(u)
	return apu
}

// SetNillableDepartmentID sets the "department_id" field if the given value is not nil.
func (apu *AccessPolicyUpdate) SetNillableDepartmentID(u *uint64) *AccessPolicyUpdate {
	if u != nil {
		apu.SetDepartmentID(*u)
	}
	return apu
}

// AddDepartmentID adds u to the "department_id" field.
func (apu *AccessPolicyUpdate) AddDepartmentID(u int64) *AccessPolicyUpdate {
	apu.mutation.AddDepartmentID(u)
	return apu
}

// SetUsers sets the "users" field.
func (apu *AccessPolicyUpdate) SetUsers(s string) *AccessPolicyUpdate {
	apu.mutation.SetUsers(s)
	return apu
}

// SetNillableUsers sets the "users" field if the given value is not nil.
func (apu *AccessPolicyUpdate) SetNillableUsers(s *string) *AccessPolicyUpdate {
	if s != nil {
		apu.SetUsers(*s)
	}
	return apu
}

// SetAccounts sets the "accounts" field.
func (apu *AccessPolicyUpdate) SetAccounts(s string) *AccessPolicyUpdate {
	apu.mutation.SetAccounts(s)
	return apu
}

// SetNillableAccounts sets the "accounts" field if the given value is not nil.
func (apu *AccessPolicyUpdate) SetNillableAccounts(s *string) *AccessPolicyUpdate {
	if s != nil {
		apu.SetAccounts(*s)
	}
	return apu
}

// SetEffecteTimeStart sets the "effecte_time_start" field.
func (apu *AccessPolicyUpdate) SetEffecteTimeStart(t time.Time) *AccessPolicyUpdate {
	apu.mutation.SetEffecteTimeStart(t)
	return apu
}

// ClearEffecteTimeStart clears the value of the "effecte_time_start" field.
func (apu *AccessPolicyUpdate) ClearEffecteTimeStart() *AccessPolicyUpdate {
	apu.mutation.ClearEffecteTimeStart()
	return apu
}

// SetEffecteTimeEnd sets the "effecte_time_end" field.
func (apu *AccessPolicyUpdate) SetEffecteTimeEnd(t time.Time) *AccessPolicyUpdate {
	apu.mutation.SetEffecteTimeEnd(t)
	return apu
}

// ClearEffecteTimeEnd clears the value of the "effecte_time_end" field.
func (apu *AccessPolicyUpdate) ClearEffecteTimeEnd() *AccessPolicyUpdate {
	apu.mutation.ClearEffecteTimeEnd()
	return apu
}

// Mutation returns the AccessPolicyMutation object of the builder.
func (apu *AccessPolicyUpdate) Mutation() *AccessPolicyMutation {
	return apu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (apu *AccessPolicyUpdate) Save(ctx context.Context) (int, error) {
	apu.defaults()
	return withHooks(ctx, apu.sqlSave, apu.mutation, apu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (apu *AccessPolicyUpdate) SaveX(ctx context.Context) int {
	affected, err := apu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (apu *AccessPolicyUpdate) Exec(ctx context.Context) error {
	_, err := apu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apu *AccessPolicyUpdate) ExecX(ctx context.Context) {
	if err := apu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (apu *AccessPolicyUpdate) defaults() {
	if _, ok := apu.mutation.UpdatedAt(); !ok {
		v := accesspolicy.UpdateDefaultUpdatedAt()
		apu.mutation.SetUpdatedAt(v)
	}
	if _, ok := apu.mutation.EffecteTimeStart(); !ok && !apu.mutation.EffecteTimeStartCleared() {
		v := accesspolicy.UpdateDefaultEffecteTimeStart()
		apu.mutation.SetEffecteTimeStart(v)
	}
	if _, ok := apu.mutation.EffecteTimeEnd(); !ok && !apu.mutation.EffecteTimeEndCleared() {
		v := accesspolicy.UpdateDefaultEffecteTimeEnd()
		apu.mutation.SetEffecteTimeEnd(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (apu *AccessPolicyUpdate) check() error {
	if v, ok := apu.mutation.Name(); ok {
		if err := accesspolicy.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "AccessPolicy.name": %w`, err)}
		}
	}
	if v, ok := apu.mutation.Power(); ok {
		if err := accesspolicy.PowerValidator(v); err != nil {
			return &ValidationError{Name: "power", err: fmt.Errorf(`ent: validator failed for field "AccessPolicy.power": %w`, err)}
		}
	}
	if v, ok := apu.mutation.DepartmentID(); ok {
		if err := accesspolicy.DepartmentIDValidator(v); err != nil {
			return &ValidationError{Name: "department_id", err: fmt.Errorf(`ent: validator failed for field "AccessPolicy.department_id": %w`, err)}
		}
	}
	if v, ok := apu.mutation.Users(); ok {
		if err := accesspolicy.UsersValidator(v); err != nil {
			return &ValidationError{Name: "users", err: fmt.Errorf(`ent: validator failed for field "AccessPolicy.users": %w`, err)}
		}
	}
	if v, ok := apu.mutation.Accounts(); ok {
		if err := accesspolicy.AccountsValidator(v); err != nil {
			return &ValidationError{Name: "accounts", err: fmt.Errorf(`ent: validator failed for field "AccessPolicy.accounts": %w`, err)}
		}
	}
	return nil
}

func (apu *AccessPolicyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := apu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(accesspolicy.Table, accesspolicy.Columns, sqlgraph.NewFieldSpec(accesspolicy.FieldID, field.TypeUint64))
	if ps := apu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := apu.mutation.UpdatedAt(); ok {
		_spec.SetField(accesspolicy.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := apu.mutation.Name(); ok {
		_spec.SetField(accesspolicy.FieldName, field.TypeString, value)
	}
	if value, ok := apu.mutation.Power(); ok {
		_spec.SetField(accesspolicy.FieldPower, field.TypeInt32, value)
	}
	if value, ok := apu.mutation.AddedPower(); ok {
		_spec.AddField(accesspolicy.FieldPower, field.TypeInt32, value)
	}
	if apu.mutation.PowerCleared() {
		_spec.ClearField(accesspolicy.FieldPower, field.TypeInt32)
	}
	if value, ok := apu.mutation.DepartmentID(); ok {
		_spec.SetField(accesspolicy.FieldDepartmentID, field.TypeUint64, value)
	}
	if value, ok := apu.mutation.AddedDepartmentID(); ok {
		_spec.AddField(accesspolicy.FieldDepartmentID, field.TypeUint64, value)
	}
	if value, ok := apu.mutation.Users(); ok {
		_spec.SetField(accesspolicy.FieldUsers, field.TypeString, value)
	}
	if value, ok := apu.mutation.Accounts(); ok {
		_spec.SetField(accesspolicy.FieldAccounts, field.TypeString, value)
	}
	if value, ok := apu.mutation.EffecteTimeStart(); ok {
		_spec.SetField(accesspolicy.FieldEffecteTimeStart, field.TypeTime, value)
	}
	if apu.mutation.EffecteTimeStartCleared() {
		_spec.ClearField(accesspolicy.FieldEffecteTimeStart, field.TypeTime)
	}
	if value, ok := apu.mutation.EffecteTimeEnd(); ok {
		_spec.SetField(accesspolicy.FieldEffecteTimeEnd, field.TypeTime, value)
	}
	if apu.mutation.EffecteTimeEndCleared() {
		_spec.ClearField(accesspolicy.FieldEffecteTimeEnd, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, apu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accesspolicy.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	apu.mutation.done = true
	return n, nil
}

// AccessPolicyUpdateOne is the builder for updating a single AccessPolicy entity.
type AccessPolicyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccessPolicyMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (apuo *AccessPolicyUpdateOne) SetUpdatedAt(t time.Time) *AccessPolicyUpdateOne {
	apuo.mutation.SetUpdatedAt(t)
	return apuo
}

// SetName sets the "name" field.
func (apuo *AccessPolicyUpdateOne) SetName(s string) *AccessPolicyUpdateOne {
	apuo.mutation.SetName(s)
	return apuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (apuo *AccessPolicyUpdateOne) SetNillableName(s *string) *AccessPolicyUpdateOne {
	if s != nil {
		apuo.SetName(*s)
	}
	return apuo
}

// SetPower sets the "power" field.
func (apuo *AccessPolicyUpdateOne) SetPower(i int32) *AccessPolicyUpdateOne {
	apuo.mutation.ResetPower()
	apuo.mutation.SetPower(i)
	return apuo
}

// SetNillablePower sets the "power" field if the given value is not nil.
func (apuo *AccessPolicyUpdateOne) SetNillablePower(i *int32) *AccessPolicyUpdateOne {
	if i != nil {
		apuo.SetPower(*i)
	}
	return apuo
}

// AddPower adds i to the "power" field.
func (apuo *AccessPolicyUpdateOne) AddPower(i int32) *AccessPolicyUpdateOne {
	apuo.mutation.AddPower(i)
	return apuo
}

// ClearPower clears the value of the "power" field.
func (apuo *AccessPolicyUpdateOne) ClearPower() *AccessPolicyUpdateOne {
	apuo.mutation.ClearPower()
	return apuo
}

// SetDepartmentID sets the "department_id" field.
func (apuo *AccessPolicyUpdateOne) SetDepartmentID(u uint64) *AccessPolicyUpdateOne {
	apuo.mutation.ResetDepartmentID()
	apuo.mutation.SetDepartmentID(u)
	return apuo
}

// SetNillableDepartmentID sets the "department_id" field if the given value is not nil.
func (apuo *AccessPolicyUpdateOne) SetNillableDepartmentID(u *uint64) *AccessPolicyUpdateOne {
	if u != nil {
		apuo.SetDepartmentID(*u)
	}
	return apuo
}

// AddDepartmentID adds u to the "department_id" field.
func (apuo *AccessPolicyUpdateOne) AddDepartmentID(u int64) *AccessPolicyUpdateOne {
	apuo.mutation.AddDepartmentID(u)
	return apuo
}

// SetUsers sets the "users" field.
func (apuo *AccessPolicyUpdateOne) SetUsers(s string) *AccessPolicyUpdateOne {
	apuo.mutation.SetUsers(s)
	return apuo
}

// SetNillableUsers sets the "users" field if the given value is not nil.
func (apuo *AccessPolicyUpdateOne) SetNillableUsers(s *string) *AccessPolicyUpdateOne {
	if s != nil {
		apuo.SetUsers(*s)
	}
	return apuo
}

// SetAccounts sets the "accounts" field.
func (apuo *AccessPolicyUpdateOne) SetAccounts(s string) *AccessPolicyUpdateOne {
	apuo.mutation.SetAccounts(s)
	return apuo
}

// SetNillableAccounts sets the "accounts" field if the given value is not nil.
func (apuo *AccessPolicyUpdateOne) SetNillableAccounts(s *string) *AccessPolicyUpdateOne {
	if s != nil {
		apuo.SetAccounts(*s)
	}
	return apuo
}

// SetEffecteTimeStart sets the "effecte_time_start" field.
func (apuo *AccessPolicyUpdateOne) SetEffecteTimeStart(t time.Time) *AccessPolicyUpdateOne {
	apuo.mutation.SetEffecteTimeStart(t)
	return apuo
}

// ClearEffecteTimeStart clears the value of the "effecte_time_start" field.
func (apuo *AccessPolicyUpdateOne) ClearEffecteTimeStart() *AccessPolicyUpdateOne {
	apuo.mutation.ClearEffecteTimeStart()
	return apuo
}

// SetEffecteTimeEnd sets the "effecte_time_end" field.
func (apuo *AccessPolicyUpdateOne) SetEffecteTimeEnd(t time.Time) *AccessPolicyUpdateOne {
	apuo.mutation.SetEffecteTimeEnd(t)
	return apuo
}

// ClearEffecteTimeEnd clears the value of the "effecte_time_end" field.
func (apuo *AccessPolicyUpdateOne) ClearEffecteTimeEnd() *AccessPolicyUpdateOne {
	apuo.mutation.ClearEffecteTimeEnd()
	return apuo
}

// Mutation returns the AccessPolicyMutation object of the builder.
func (apuo *AccessPolicyUpdateOne) Mutation() *AccessPolicyMutation {
	return apuo.mutation
}

// Where appends a list predicates to the AccessPolicyUpdate builder.
func (apuo *AccessPolicyUpdateOne) Where(ps ...predicate.AccessPolicy) *AccessPolicyUpdateOne {
	apuo.mutation.Where(ps...)
	return apuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (apuo *AccessPolicyUpdateOne) Select(field string, fields ...string) *AccessPolicyUpdateOne {
	apuo.fields = append([]string{field}, fields...)
	return apuo
}

// Save executes the query and returns the updated AccessPolicy entity.
func (apuo *AccessPolicyUpdateOne) Save(ctx context.Context) (*AccessPolicy, error) {
	apuo.defaults()
	return withHooks(ctx, apuo.sqlSave, apuo.mutation, apuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (apuo *AccessPolicyUpdateOne) SaveX(ctx context.Context) *AccessPolicy {
	node, err := apuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (apuo *AccessPolicyUpdateOne) Exec(ctx context.Context) error {
	_, err := apuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apuo *AccessPolicyUpdateOne) ExecX(ctx context.Context) {
	if err := apuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (apuo *AccessPolicyUpdateOne) defaults() {
	if _, ok := apuo.mutation.UpdatedAt(); !ok {
		v := accesspolicy.UpdateDefaultUpdatedAt()
		apuo.mutation.SetUpdatedAt(v)
	}
	if _, ok := apuo.mutation.EffecteTimeStart(); !ok && !apuo.mutation.EffecteTimeStartCleared() {
		v := accesspolicy.UpdateDefaultEffecteTimeStart()
		apuo.mutation.SetEffecteTimeStart(v)
	}
	if _, ok := apuo.mutation.EffecteTimeEnd(); !ok && !apuo.mutation.EffecteTimeEndCleared() {
		v := accesspolicy.UpdateDefaultEffecteTimeEnd()
		apuo.mutation.SetEffecteTimeEnd(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (apuo *AccessPolicyUpdateOne) check() error {
	if v, ok := apuo.mutation.Name(); ok {
		if err := accesspolicy.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "AccessPolicy.name": %w`, err)}
		}
	}
	if v, ok := apuo.mutation.Power(); ok {
		if err := accesspolicy.PowerValidator(v); err != nil {
			return &ValidationError{Name: "power", err: fmt.Errorf(`ent: validator failed for field "AccessPolicy.power": %w`, err)}
		}
	}
	if v, ok := apuo.mutation.DepartmentID(); ok {
		if err := accesspolicy.DepartmentIDValidator(v); err != nil {
			return &ValidationError{Name: "department_id", err: fmt.Errorf(`ent: validator failed for field "AccessPolicy.department_id": %w`, err)}
		}
	}
	if v, ok := apuo.mutation.Users(); ok {
		if err := accesspolicy.UsersValidator(v); err != nil {
			return &ValidationError{Name: "users", err: fmt.Errorf(`ent: validator failed for field "AccessPolicy.users": %w`, err)}
		}
	}
	if v, ok := apuo.mutation.Accounts(); ok {
		if err := accesspolicy.AccountsValidator(v); err != nil {
			return &ValidationError{Name: "accounts", err: fmt.Errorf(`ent: validator failed for field "AccessPolicy.accounts": %w`, err)}
		}
	}
	return nil
}

func (apuo *AccessPolicyUpdateOne) sqlSave(ctx context.Context) (_node *AccessPolicy, err error) {
	if err := apuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(accesspolicy.Table, accesspolicy.Columns, sqlgraph.NewFieldSpec(accesspolicy.FieldID, field.TypeUint64))
	id, ok := apuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AccessPolicy.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := apuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accesspolicy.FieldID)
		for _, f := range fields {
			if !accesspolicy.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != accesspolicy.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := apuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := apuo.mutation.UpdatedAt(); ok {
		_spec.SetField(accesspolicy.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := apuo.mutation.Name(); ok {
		_spec.SetField(accesspolicy.FieldName, field.TypeString, value)
	}
	if value, ok := apuo.mutation.Power(); ok {
		_spec.SetField(accesspolicy.FieldPower, field.TypeInt32, value)
	}
	if value, ok := apuo.mutation.AddedPower(); ok {
		_spec.AddField(accesspolicy.FieldPower, field.TypeInt32, value)
	}
	if apuo.mutation.PowerCleared() {
		_spec.ClearField(accesspolicy.FieldPower, field.TypeInt32)
	}
	if value, ok := apuo.mutation.DepartmentID(); ok {
		_spec.SetField(accesspolicy.FieldDepartmentID, field.TypeUint64, value)
	}
	if value, ok := apuo.mutation.AddedDepartmentID(); ok {
		_spec.AddField(accesspolicy.FieldDepartmentID, field.TypeUint64, value)
	}
	if value, ok := apuo.mutation.Users(); ok {
		_spec.SetField(accesspolicy.FieldUsers, field.TypeString, value)
	}
	if value, ok := apuo.mutation.Accounts(); ok {
		_spec.SetField(accesspolicy.FieldAccounts, field.TypeString, value)
	}
	if value, ok := apuo.mutation.EffecteTimeStart(); ok {
		_spec.SetField(accesspolicy.FieldEffecteTimeStart, field.TypeTime, value)
	}
	if apuo.mutation.EffecteTimeStartCleared() {
		_spec.ClearField(accesspolicy.FieldEffecteTimeStart, field.TypeTime)
	}
	if value, ok := apuo.mutation.EffecteTimeEnd(); ok {
		_spec.SetField(accesspolicy.FieldEffecteTimeEnd, field.TypeTime, value)
	}
	if apuo.mutation.EffecteTimeEndCleared() {
		_spec.ClearField(accesspolicy.FieldEffecteTimeEnd, field.TypeTime)
	}
	_node = &AccessPolicy{config: apuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, apuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accesspolicy.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	apuo.mutation.done = true
	return _node, nil
}