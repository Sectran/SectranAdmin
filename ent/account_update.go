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

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AccountUpdate) SetUpdatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetUsername sets the "username" field.
func (au *AccountUpdate) SetUsername(s string) *AccountUpdate {
	au.mutation.SetUsername(s)
	return au
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (au *AccountUpdate) SetNillableUsername(s *string) *AccountUpdate {
	if s != nil {
		au.SetUsername(*s)
	}
	return au
}

// SetPort sets the "port" field.
func (au *AccountUpdate) SetPort(u uint32) *AccountUpdate {
	au.mutation.ResetPort()
	au.mutation.SetPort(u)
	return au
}

// SetNillablePort sets the "port" field if the given value is not nil.
func (au *AccountUpdate) SetNillablePort(u *uint32) *AccountUpdate {
	if u != nil {
		au.SetPort(*u)
	}
	return au
}

// AddPort adds u to the "port" field.
func (au *AccountUpdate) AddPort(u int32) *AccountUpdate {
	au.mutation.AddPort(u)
	return au
}

// SetProtocol sets the "protocol" field.
func (au *AccountUpdate) SetProtocol(u uint8) *AccountUpdate {
	au.mutation.ResetProtocol()
	au.mutation.SetProtocol(u)
	return au
}

// SetNillableProtocol sets the "protocol" field if the given value is not nil.
func (au *AccountUpdate) SetNillableProtocol(u *uint8) *AccountUpdate {
	if u != nil {
		au.SetProtocol(*u)
	}
	return au
}

// AddProtocol adds u to the "protocol" field.
func (au *AccountUpdate) AddProtocol(u int8) *AccountUpdate {
	au.mutation.AddProtocol(u)
	return au
}

// SetPassword sets the "password" field.
func (au *AccountUpdate) SetPassword(s string) *AccountUpdate {
	au.mutation.SetPassword(s)
	return au
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (au *AccountUpdate) SetNillablePassword(s *string) *AccountUpdate {
	if s != nil {
		au.SetPassword(*s)
	}
	return au
}

// ClearPassword clears the value of the "password" field.
func (au *AccountUpdate) ClearPassword() *AccountUpdate {
	au.mutation.ClearPassword()
	return au
}

// SetPrivateKey sets the "private_key" field.
func (au *AccountUpdate) SetPrivateKey(s string) *AccountUpdate {
	au.mutation.SetPrivateKey(s)
	return au
}

// SetNillablePrivateKey sets the "private_key" field if the given value is not nil.
func (au *AccountUpdate) SetNillablePrivateKey(s *string) *AccountUpdate {
	if s != nil {
		au.SetPrivateKey(*s)
	}
	return au
}

// ClearPrivateKey clears the value of the "private_key" field.
func (au *AccountUpdate) ClearPrivateKey() *AccountUpdate {
	au.mutation.ClearPrivateKey()
	return au
}

// SetPrivateKeyPassword sets the "private_key_password" field.
func (au *AccountUpdate) SetPrivateKeyPassword(s string) *AccountUpdate {
	au.mutation.SetPrivateKeyPassword(s)
	return au
}

// SetNillablePrivateKeyPassword sets the "private_key_password" field if the given value is not nil.
func (au *AccountUpdate) SetNillablePrivateKeyPassword(s *string) *AccountUpdate {
	if s != nil {
		au.SetPrivateKeyPassword(*s)
	}
	return au
}

// ClearPrivateKeyPassword clears the value of the "private_key_password" field.
func (au *AccountUpdate) ClearPrivateKeyPassword() *AccountUpdate {
	au.mutation.ClearPrivateKeyPassword()
	return au
}

// SetDeviceID sets the "device_id" field.
func (au *AccountUpdate) SetDeviceID(u uint64) *AccountUpdate {
	au.mutation.SetDeviceID(u)
	return au
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (au *AccountUpdate) SetNillableDeviceID(u *uint64) *AccountUpdate {
	if u != nil {
		au.SetDeviceID(*u)
	}
	return au
}

// SetDepartmentID sets the "department_id" field.
func (au *AccountUpdate) SetDepartmentID(u uint64) *AccountUpdate {
	au.mutation.SetDepartmentID(u)
	return au
}

// SetNillableDepartmentID sets the "department_id" field if the given value is not nil.
func (au *AccountUpdate) SetNillableDepartmentID(u *uint64) *AccountUpdate {
	if u != nil {
		au.SetDepartmentID(*u)
	}
	return au
}

// SetDevicesID sets the "devices" edge to the Device entity by ID.
func (au *AccountUpdate) SetDevicesID(id uint64) *AccountUpdate {
	au.mutation.SetDevicesID(id)
	return au
}

// SetDevices sets the "devices" edge to the Device entity.
func (au *AccountUpdate) SetDevices(d *Device) *AccountUpdate {
	return au.SetDevicesID(d.ID)
}

// SetDepartmentsID sets the "departments" edge to the Department entity by ID.
func (au *AccountUpdate) SetDepartmentsID(id uint64) *AccountUpdate {
	au.mutation.SetDepartmentsID(id)
	return au
}

// SetDepartments sets the "departments" edge to the Department entity.
func (au *AccountUpdate) SetDepartments(d *Department) *AccountUpdate {
	return au.SetDepartmentsID(d.ID)
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// ClearDevices clears the "devices" edge to the Device entity.
func (au *AccountUpdate) ClearDevices() *AccountUpdate {
	au.mutation.ClearDevices()
	return au
}

// ClearDepartments clears the "departments" edge to the Department entity.
func (au *AccountUpdate) ClearDepartments() *AccountUpdate {
	au.mutation.ClearDepartments()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AccountUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AccountUpdate) check() error {
	if v, ok := au.mutation.Username(); ok {
		if err := account.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Account.username": %w`, err)}
		}
	}
	if v, ok := au.mutation.Port(); ok {
		if err := account.PortValidator(v); err != nil {
			return &ValidationError{Name: "port", err: fmt.Errorf(`ent: validator failed for field "Account.port": %w`, err)}
		}
	}
	if v, ok := au.mutation.Protocol(); ok {
		if err := account.ProtocolValidator(v); err != nil {
			return &ValidationError{Name: "protocol", err: fmt.Errorf(`ent: validator failed for field "Account.protocol": %w`, err)}
		}
	}
	if v, ok := au.mutation.Password(); ok {
		if err := account.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Account.password": %w`, err)}
		}
	}
	if v, ok := au.mutation.PrivateKey(); ok {
		if err := account.PrivateKeyValidator(v); err != nil {
			return &ValidationError{Name: "private_key", err: fmt.Errorf(`ent: validator failed for field "Account.private_key": %w`, err)}
		}
	}
	if v, ok := au.mutation.PrivateKeyPassword(); ok {
		if err := account.PrivateKeyPasswordValidator(v); err != nil {
			return &ValidationError{Name: "private_key_password", err: fmt.Errorf(`ent: validator failed for field "Account.private_key_password": %w`, err)}
		}
	}
	if v, ok := au.mutation.DeviceID(); ok {
		if err := account.DeviceIDValidator(v); err != nil {
			return &ValidationError{Name: "device_id", err: fmt.Errorf(`ent: validator failed for field "Account.device_id": %w`, err)}
		}
	}
	if v, ok := au.mutation.DepartmentID(); ok {
		if err := account.DepartmentIDValidator(v); err != nil {
			return &ValidationError{Name: "department_id", err: fmt.Errorf(`ent: validator failed for field "Account.department_id": %w`, err)}
		}
	}
	if au.mutation.DevicesCleared() && len(au.mutation.DevicesIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Account.devices"`)
	}
	if au.mutation.DepartmentsCleared() && len(au.mutation.DepartmentsIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Account.departments"`)
	}
	return nil
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeUint64))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.Username(); ok {
		_spec.SetField(account.FieldUsername, field.TypeString, value)
	}
	if value, ok := au.mutation.Port(); ok {
		_spec.SetField(account.FieldPort, field.TypeUint32, value)
	}
	if value, ok := au.mutation.AddedPort(); ok {
		_spec.AddField(account.FieldPort, field.TypeUint32, value)
	}
	if value, ok := au.mutation.Protocol(); ok {
		_spec.SetField(account.FieldProtocol, field.TypeUint8, value)
	}
	if value, ok := au.mutation.AddedProtocol(); ok {
		_spec.AddField(account.FieldProtocol, field.TypeUint8, value)
	}
	if value, ok := au.mutation.Password(); ok {
		_spec.SetField(account.FieldPassword, field.TypeString, value)
	}
	if au.mutation.PasswordCleared() {
		_spec.ClearField(account.FieldPassword, field.TypeString)
	}
	if value, ok := au.mutation.PrivateKey(); ok {
		_spec.SetField(account.FieldPrivateKey, field.TypeString, value)
	}
	if au.mutation.PrivateKeyCleared() {
		_spec.ClearField(account.FieldPrivateKey, field.TypeString)
	}
	if value, ok := au.mutation.PrivateKeyPassword(); ok {
		_spec.SetField(account.FieldPrivateKeyPassword, field.TypeString, value)
	}
	if au.mutation.PrivateKeyPasswordCleared() {
		_spec.ClearField(account.FieldPrivateKeyPassword, field.TypeString)
	}
	if au.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   account.DevicesTable,
			Columns: []string{account.DevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.DevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   account.DevicesTable,
			Columns: []string{account.DevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.DepartmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   account.DepartmentsTable,
			Columns: []string{account.DepartmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.DepartmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   account.DepartmentsTable,
			Columns: []string{account.DepartmentsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AccountUpdateOne) SetUpdatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetUsername sets the "username" field.
func (auo *AccountUpdateOne) SetUsername(s string) *AccountUpdateOne {
	auo.mutation.SetUsername(s)
	return auo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableUsername(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetUsername(*s)
	}
	return auo
}

// SetPort sets the "port" field.
func (auo *AccountUpdateOne) SetPort(u uint32) *AccountUpdateOne {
	auo.mutation.ResetPort()
	auo.mutation.SetPort(u)
	return auo
}

// SetNillablePort sets the "port" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillablePort(u *uint32) *AccountUpdateOne {
	if u != nil {
		auo.SetPort(*u)
	}
	return auo
}

// AddPort adds u to the "port" field.
func (auo *AccountUpdateOne) AddPort(u int32) *AccountUpdateOne {
	auo.mutation.AddPort(u)
	return auo
}

// SetProtocol sets the "protocol" field.
func (auo *AccountUpdateOne) SetProtocol(u uint8) *AccountUpdateOne {
	auo.mutation.ResetProtocol()
	auo.mutation.SetProtocol(u)
	return auo
}

// SetNillableProtocol sets the "protocol" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableProtocol(u *uint8) *AccountUpdateOne {
	if u != nil {
		auo.SetProtocol(*u)
	}
	return auo
}

// AddProtocol adds u to the "protocol" field.
func (auo *AccountUpdateOne) AddProtocol(u int8) *AccountUpdateOne {
	auo.mutation.AddProtocol(u)
	return auo
}

// SetPassword sets the "password" field.
func (auo *AccountUpdateOne) SetPassword(s string) *AccountUpdateOne {
	auo.mutation.SetPassword(s)
	return auo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillablePassword(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetPassword(*s)
	}
	return auo
}

// ClearPassword clears the value of the "password" field.
func (auo *AccountUpdateOne) ClearPassword() *AccountUpdateOne {
	auo.mutation.ClearPassword()
	return auo
}

// SetPrivateKey sets the "private_key" field.
func (auo *AccountUpdateOne) SetPrivateKey(s string) *AccountUpdateOne {
	auo.mutation.SetPrivateKey(s)
	return auo
}

// SetNillablePrivateKey sets the "private_key" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillablePrivateKey(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetPrivateKey(*s)
	}
	return auo
}

// ClearPrivateKey clears the value of the "private_key" field.
func (auo *AccountUpdateOne) ClearPrivateKey() *AccountUpdateOne {
	auo.mutation.ClearPrivateKey()
	return auo
}

// SetPrivateKeyPassword sets the "private_key_password" field.
func (auo *AccountUpdateOne) SetPrivateKeyPassword(s string) *AccountUpdateOne {
	auo.mutation.SetPrivateKeyPassword(s)
	return auo
}

// SetNillablePrivateKeyPassword sets the "private_key_password" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillablePrivateKeyPassword(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetPrivateKeyPassword(*s)
	}
	return auo
}

// ClearPrivateKeyPassword clears the value of the "private_key_password" field.
func (auo *AccountUpdateOne) ClearPrivateKeyPassword() *AccountUpdateOne {
	auo.mutation.ClearPrivateKeyPassword()
	return auo
}

// SetDeviceID sets the "device_id" field.
func (auo *AccountUpdateOne) SetDeviceID(u uint64) *AccountUpdateOne {
	auo.mutation.SetDeviceID(u)
	return auo
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableDeviceID(u *uint64) *AccountUpdateOne {
	if u != nil {
		auo.SetDeviceID(*u)
	}
	return auo
}

// SetDepartmentID sets the "department_id" field.
func (auo *AccountUpdateOne) SetDepartmentID(u uint64) *AccountUpdateOne {
	auo.mutation.SetDepartmentID(u)
	return auo
}

// SetNillableDepartmentID sets the "department_id" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableDepartmentID(u *uint64) *AccountUpdateOne {
	if u != nil {
		auo.SetDepartmentID(*u)
	}
	return auo
}

// SetDevicesID sets the "devices" edge to the Device entity by ID.
func (auo *AccountUpdateOne) SetDevicesID(id uint64) *AccountUpdateOne {
	auo.mutation.SetDevicesID(id)
	return auo
}

// SetDevices sets the "devices" edge to the Device entity.
func (auo *AccountUpdateOne) SetDevices(d *Device) *AccountUpdateOne {
	return auo.SetDevicesID(d.ID)
}

// SetDepartmentsID sets the "departments" edge to the Department entity by ID.
func (auo *AccountUpdateOne) SetDepartmentsID(id uint64) *AccountUpdateOne {
	auo.mutation.SetDepartmentsID(id)
	return auo
}

// SetDepartments sets the "departments" edge to the Department entity.
func (auo *AccountUpdateOne) SetDepartments(d *Department) *AccountUpdateOne {
	return auo.SetDepartmentsID(d.ID)
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// ClearDevices clears the "devices" edge to the Device entity.
func (auo *AccountUpdateOne) ClearDevices() *AccountUpdateOne {
	auo.mutation.ClearDevices()
	return auo
}

// ClearDepartments clears the "departments" edge to the Department entity.
func (auo *AccountUpdateOne) ClearDepartments() *AccountUpdateOne {
	auo.mutation.ClearDepartments()
	return auo
}

// Where appends a list predicates to the AccountUpdate builder.
func (auo *AccountUpdateOne) Where(ps ...predicate.Account) *AccountUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AccountUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AccountUpdateOne) check() error {
	if v, ok := auo.mutation.Username(); ok {
		if err := account.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "Account.username": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Port(); ok {
		if err := account.PortValidator(v); err != nil {
			return &ValidationError{Name: "port", err: fmt.Errorf(`ent: validator failed for field "Account.port": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Protocol(); ok {
		if err := account.ProtocolValidator(v); err != nil {
			return &ValidationError{Name: "protocol", err: fmt.Errorf(`ent: validator failed for field "Account.protocol": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Password(); ok {
		if err := account.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Account.password": %w`, err)}
		}
	}
	if v, ok := auo.mutation.PrivateKey(); ok {
		if err := account.PrivateKeyValidator(v); err != nil {
			return &ValidationError{Name: "private_key", err: fmt.Errorf(`ent: validator failed for field "Account.private_key": %w`, err)}
		}
	}
	if v, ok := auo.mutation.PrivateKeyPassword(); ok {
		if err := account.PrivateKeyPasswordValidator(v); err != nil {
			return &ValidationError{Name: "private_key_password", err: fmt.Errorf(`ent: validator failed for field "Account.private_key_password": %w`, err)}
		}
	}
	if v, ok := auo.mutation.DeviceID(); ok {
		if err := account.DeviceIDValidator(v); err != nil {
			return &ValidationError{Name: "device_id", err: fmt.Errorf(`ent: validator failed for field "Account.device_id": %w`, err)}
		}
	}
	if v, ok := auo.mutation.DepartmentID(); ok {
		if err := account.DepartmentIDValidator(v); err != nil {
			return &ValidationError{Name: "department_id", err: fmt.Errorf(`ent: validator failed for field "Account.department_id": %w`, err)}
		}
	}
	if auo.mutation.DevicesCleared() && len(auo.mutation.DevicesIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Account.devices"`)
	}
	if auo.mutation.DepartmentsCleared() && len(auo.mutation.DepartmentsIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Account.departments"`)
	}
	return nil
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeUint64))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.Username(); ok {
		_spec.SetField(account.FieldUsername, field.TypeString, value)
	}
	if value, ok := auo.mutation.Port(); ok {
		_spec.SetField(account.FieldPort, field.TypeUint32, value)
	}
	if value, ok := auo.mutation.AddedPort(); ok {
		_spec.AddField(account.FieldPort, field.TypeUint32, value)
	}
	if value, ok := auo.mutation.Protocol(); ok {
		_spec.SetField(account.FieldProtocol, field.TypeUint8, value)
	}
	if value, ok := auo.mutation.AddedProtocol(); ok {
		_spec.AddField(account.FieldProtocol, field.TypeUint8, value)
	}
	if value, ok := auo.mutation.Password(); ok {
		_spec.SetField(account.FieldPassword, field.TypeString, value)
	}
	if auo.mutation.PasswordCleared() {
		_spec.ClearField(account.FieldPassword, field.TypeString)
	}
	if value, ok := auo.mutation.PrivateKey(); ok {
		_spec.SetField(account.FieldPrivateKey, field.TypeString, value)
	}
	if auo.mutation.PrivateKeyCleared() {
		_spec.ClearField(account.FieldPrivateKey, field.TypeString)
	}
	if value, ok := auo.mutation.PrivateKeyPassword(); ok {
		_spec.SetField(account.FieldPrivateKeyPassword, field.TypeString, value)
	}
	if auo.mutation.PrivateKeyPasswordCleared() {
		_spec.ClearField(account.FieldPrivateKeyPassword, field.TypeString)
	}
	if auo.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   account.DevicesTable,
			Columns: []string{account.DevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.DevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   account.DevicesTable,
			Columns: []string{account.DevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.DepartmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   account.DepartmentsTable,
			Columns: []string{account.DepartmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.DepartmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   account.DepartmentsTable,
			Columns: []string{account.DepartmentsColumn},
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
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
