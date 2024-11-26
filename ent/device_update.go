// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sectran_admin/ent/account"
	"sectran_admin/ent/department"
	"sectran_admin/ent/device"
	"sectran_admin/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeviceUpdate is the builder for updating Device entities.
type DeviceUpdate struct {
	config
	hooks    []Hook
	mutation *DeviceMutation
}

// Where appends a list predicates to the DeviceUpdate builder.
func (du *DeviceUpdate) Where(ps ...predicate.Device) *DeviceUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DeviceUpdate) SetUpdatedAt(t time.Time) *DeviceUpdate {
	du.mutation.SetUpdatedAt(t)
	return du
}

// SetName sets the "name" field.
func (du *DeviceUpdate) SetName(s string) *DeviceUpdate {
	du.mutation.SetName(s)
	return du
}

// SetNillableName sets the "name" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableName(s *string) *DeviceUpdate {
	if s != nil {
		du.SetName(*s)
	}
	return du
}

// SetDepartmentID sets the "department_id" field.
func (du *DeviceUpdate) SetDepartmentID(u uint64) *DeviceUpdate {
	du.mutation.SetDepartmentID(u)
	return du
}

// SetNillableDepartmentID sets the "department_id" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableDepartmentID(u *uint64) *DeviceUpdate {
	if u != nil {
		du.SetDepartmentID(*u)
	}
	return du
}

// SetHost sets the "host" field.
func (du *DeviceUpdate) SetHost(s string) *DeviceUpdate {
	du.mutation.SetHost(s)
	return du
}

// SetNillableHost sets the "host" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableHost(s *string) *DeviceUpdate {
	if s != nil {
		du.SetHost(*s)
	}
	return du
}

// SetType sets the "type" field.
func (du *DeviceUpdate) SetType(s string) *DeviceUpdate {
	du.mutation.SetType(s)
	return du
}

// SetNillableType sets the "type" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableType(s *string) *DeviceUpdate {
	if s != nil {
		du.SetType(*s)
	}
	return du
}

// SetDescription sets the "description" field.
func (du *DeviceUpdate) SetDescription(s string) *DeviceUpdate {
	du.mutation.SetDescription(s)
	return du
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableDescription(s *string) *DeviceUpdate {
	if s != nil {
		du.SetDescription(*s)
	}
	return du
}

// SetDepartmentsID sets the "departments" edge to the Department entity by ID.
func (du *DeviceUpdate) SetDepartmentsID(id uint64) *DeviceUpdate {
	du.mutation.SetDepartmentsID(id)
	return du
}

// SetDepartments sets the "departments" edge to the Department entity.
func (du *DeviceUpdate) SetDepartments(d *Department) *DeviceUpdate {
	return du.SetDepartmentsID(d.ID)
}

// AddAccountIDs adds the "accounts" edge to the Account entity by IDs.
func (du *DeviceUpdate) AddAccountIDs(ids ...uint64) *DeviceUpdate {
	du.mutation.AddAccountIDs(ids...)
	return du
}

// AddAccounts adds the "accounts" edges to the Account entity.
func (du *DeviceUpdate) AddAccounts(a ...*Account) *DeviceUpdate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return du.AddAccountIDs(ids...)
}

// Mutation returns the DeviceMutation object of the builder.
func (du *DeviceUpdate) Mutation() *DeviceMutation {
	return du.mutation
}

// ClearDepartments clears the "departments" edge to the Department entity.
func (du *DeviceUpdate) ClearDepartments() *DeviceUpdate {
	du.mutation.ClearDepartments()
	return du
}

// ClearAccounts clears all "accounts" edges to the Account entity.
func (du *DeviceUpdate) ClearAccounts() *DeviceUpdate {
	du.mutation.ClearAccounts()
	return du
}

// RemoveAccountIDs removes the "accounts" edge to Account entities by IDs.
func (du *DeviceUpdate) RemoveAccountIDs(ids ...uint64) *DeviceUpdate {
	du.mutation.RemoveAccountIDs(ids...)
	return du
}

// RemoveAccounts removes "accounts" edges to Account entities.
func (du *DeviceUpdate) RemoveAccounts(a ...*Account) *DeviceUpdate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return du.RemoveAccountIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DeviceUpdate) Save(ctx context.Context) (int, error) {
	du.defaults()
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DeviceUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DeviceUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DeviceUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DeviceUpdate) defaults() {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		v := device.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DeviceUpdate) check() error {
	if v, ok := du.mutation.Name(); ok {
		if err := device.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Device.name": %w`, err)}
		}
	}
	if v, ok := du.mutation.DepartmentID(); ok {
		if err := device.DepartmentIDValidator(v); err != nil {
			return &ValidationError{Name: "department_id", err: fmt.Errorf(`ent: validator failed for field "Device.department_id": %w`, err)}
		}
	}
	if v, ok := du.mutation.Host(); ok {
		if err := device.HostValidator(v); err != nil {
			return &ValidationError{Name: "host", err: fmt.Errorf(`ent: validator failed for field "Device.host": %w`, err)}
		}
	}
	if v, ok := du.mutation.GetType(); ok {
		if err := device.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Device.type": %w`, err)}
		}
	}
	if v, ok := du.mutation.Description(); ok {
		if err := device.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Device.description": %w`, err)}
		}
	}
	if du.mutation.DepartmentsCleared() && len(du.mutation.DepartmentsIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Device.departments"`)
	}
	return nil
}

func (du *DeviceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := du.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(device.Table, device.Columns, sqlgraph.NewFieldSpec(device.FieldID, field.TypeUint64))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.SetField(device.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.SetField(device.FieldName, field.TypeString, value)
	}
	if value, ok := du.mutation.Host(); ok {
		_spec.SetField(device.FieldHost, field.TypeString, value)
	}
	if value, ok := du.mutation.GetType(); ok {
		_spec.SetField(device.FieldType, field.TypeString, value)
	}
	if value, ok := du.mutation.Description(); ok {
		_spec.SetField(device.FieldDescription, field.TypeString, value)
	}
	if du.mutation.DepartmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   device.DepartmentsTable,
			Columns: []string{device.DepartmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DepartmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   device.DepartmentsTable,
			Columns: []string{device.DepartmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   device.AccountsTable,
			Columns: []string{device.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedAccountsIDs(); len(nodes) > 0 && !du.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   device.AccountsTable,
			Columns: []string{device.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   device.AccountsTable,
			Columns: []string{device.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{device.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DeviceUpdateOne is the builder for updating a single Device entity.
type DeviceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeviceMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DeviceUpdateOne) SetUpdatedAt(t time.Time) *DeviceUpdateOne {
	duo.mutation.SetUpdatedAt(t)
	return duo
}

// SetName sets the "name" field.
func (duo *DeviceUpdateOne) SetName(s string) *DeviceUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableName(s *string) *DeviceUpdateOne {
	if s != nil {
		duo.SetName(*s)
	}
	return duo
}

// SetDepartmentID sets the "department_id" field.
func (duo *DeviceUpdateOne) SetDepartmentID(u uint64) *DeviceUpdateOne {
	duo.mutation.SetDepartmentID(u)
	return duo
}

// SetNillableDepartmentID sets the "department_id" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableDepartmentID(u *uint64) *DeviceUpdateOne {
	if u != nil {
		duo.SetDepartmentID(*u)
	}
	return duo
}

// SetHost sets the "host" field.
func (duo *DeviceUpdateOne) SetHost(s string) *DeviceUpdateOne {
	duo.mutation.SetHost(s)
	return duo
}

// SetNillableHost sets the "host" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableHost(s *string) *DeviceUpdateOne {
	if s != nil {
		duo.SetHost(*s)
	}
	return duo
}

// SetType sets the "type" field.
func (duo *DeviceUpdateOne) SetType(s string) *DeviceUpdateOne {
	duo.mutation.SetType(s)
	return duo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableType(s *string) *DeviceUpdateOne {
	if s != nil {
		duo.SetType(*s)
	}
	return duo
}

// SetDescription sets the "description" field.
func (duo *DeviceUpdateOne) SetDescription(s string) *DeviceUpdateOne {
	duo.mutation.SetDescription(s)
	return duo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableDescription(s *string) *DeviceUpdateOne {
	if s != nil {
		duo.SetDescription(*s)
	}
	return duo
}

// SetDepartmentsID sets the "departments" edge to the Department entity by ID.
func (duo *DeviceUpdateOne) SetDepartmentsID(id uint64) *DeviceUpdateOne {
	duo.mutation.SetDepartmentsID(id)
	return duo
}

// SetDepartments sets the "departments" edge to the Department entity.
func (duo *DeviceUpdateOne) SetDepartments(d *Department) *DeviceUpdateOne {
	return duo.SetDepartmentsID(d.ID)
}

// AddAccountIDs adds the "accounts" edge to the Account entity by IDs.
func (duo *DeviceUpdateOne) AddAccountIDs(ids ...uint64) *DeviceUpdateOne {
	duo.mutation.AddAccountIDs(ids...)
	return duo
}

// AddAccounts adds the "accounts" edges to the Account entity.
func (duo *DeviceUpdateOne) AddAccounts(a ...*Account) *DeviceUpdateOne {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return duo.AddAccountIDs(ids...)
}

// Mutation returns the DeviceMutation object of the builder.
func (duo *DeviceUpdateOne) Mutation() *DeviceMutation {
	return duo.mutation
}

// ClearDepartments clears the "departments" edge to the Department entity.
func (duo *DeviceUpdateOne) ClearDepartments() *DeviceUpdateOne {
	duo.mutation.ClearDepartments()
	return duo
}

// ClearAccounts clears all "accounts" edges to the Account entity.
func (duo *DeviceUpdateOne) ClearAccounts() *DeviceUpdateOne {
	duo.mutation.ClearAccounts()
	return duo
}

// RemoveAccountIDs removes the "accounts" edge to Account entities by IDs.
func (duo *DeviceUpdateOne) RemoveAccountIDs(ids ...uint64) *DeviceUpdateOne {
	duo.mutation.RemoveAccountIDs(ids...)
	return duo
}

// RemoveAccounts removes "accounts" edges to Account entities.
func (duo *DeviceUpdateOne) RemoveAccounts(a ...*Account) *DeviceUpdateOne {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return duo.RemoveAccountIDs(ids...)
}

// Where appends a list predicates to the DeviceUpdate builder.
func (duo *DeviceUpdateOne) Where(ps ...predicate.Device) *DeviceUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DeviceUpdateOne) Select(field string, fields ...string) *DeviceUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Device entity.
func (duo *DeviceUpdateOne) Save(ctx context.Context) (*Device, error) {
	duo.defaults()
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DeviceUpdateOne) SaveX(ctx context.Context) *Device {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DeviceUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DeviceUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DeviceUpdateOne) defaults() {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		v := device.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DeviceUpdateOne) check() error {
	if v, ok := duo.mutation.Name(); ok {
		if err := device.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Device.name": %w`, err)}
		}
	}
	if v, ok := duo.mutation.DepartmentID(); ok {
		if err := device.DepartmentIDValidator(v); err != nil {
			return &ValidationError{Name: "department_id", err: fmt.Errorf(`ent: validator failed for field "Device.department_id": %w`, err)}
		}
	}
	if v, ok := duo.mutation.Host(); ok {
		if err := device.HostValidator(v); err != nil {
			return &ValidationError{Name: "host", err: fmt.Errorf(`ent: validator failed for field "Device.host": %w`, err)}
		}
	}
	if v, ok := duo.mutation.GetType(); ok {
		if err := device.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Device.type": %w`, err)}
		}
	}
	if v, ok := duo.mutation.Description(); ok {
		if err := device.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Device.description": %w`, err)}
		}
	}
	if duo.mutation.DepartmentsCleared() && len(duo.mutation.DepartmentsIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Device.departments"`)
	}
	return nil
}

func (duo *DeviceUpdateOne) sqlSave(ctx context.Context) (_node *Device, err error) {
	if err := duo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(device.Table, device.Columns, sqlgraph.NewFieldSpec(device.FieldID, field.TypeUint64))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Device.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, device.FieldID)
		for _, f := range fields {
			if !device.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != device.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.SetField(device.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.SetField(device.FieldName, field.TypeString, value)
	}
	if value, ok := duo.mutation.Host(); ok {
		_spec.SetField(device.FieldHost, field.TypeString, value)
	}
	if value, ok := duo.mutation.GetType(); ok {
		_spec.SetField(device.FieldType, field.TypeString, value)
	}
	if value, ok := duo.mutation.Description(); ok {
		_spec.SetField(device.FieldDescription, field.TypeString, value)
	}
	if duo.mutation.DepartmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   device.DepartmentsTable,
			Columns: []string{device.DepartmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DepartmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   device.DepartmentsTable,
			Columns: []string{device.DepartmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   device.AccountsTable,
			Columns: []string{device.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedAccountsIDs(); len(nodes) > 0 && !duo.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   device.AccountsTable,
			Columns: []string{device.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   device.AccountsTable,
			Columns: []string{device.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Device{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{device.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
