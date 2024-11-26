// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sectran_admin/ent/account"
	"sectran_admin/ent/department"
	"sectran_admin/ent/device"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeviceCreate is the builder for creating a Device entity.
type DeviceCreate struct {
	config
	mutation *DeviceMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (dc *DeviceCreate) SetCreatedAt(t time.Time) *DeviceCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableCreatedAt(t *time.Time) *DeviceCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DeviceCreate) SetUpdatedAt(t time.Time) *DeviceCreate {
	dc.mutation.SetUpdatedAt(t)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableUpdatedAt(t *time.Time) *DeviceCreate {
	if t != nil {
		dc.SetUpdatedAt(*t)
	}
	return dc
}

// SetName sets the "name" field.
func (dc *DeviceCreate) SetName(s string) *DeviceCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetDepartmentID sets the "department_id" field.
func (dc *DeviceCreate) SetDepartmentID(u uint64) *DeviceCreate {
	dc.mutation.SetDepartmentID(u)
	return dc
}

// SetHost sets the "host" field.
func (dc *DeviceCreate) SetHost(s string) *DeviceCreate {
	dc.mutation.SetHost(s)
	return dc
}

// SetType sets the "type" field.
func (dc *DeviceCreate) SetType(s string) *DeviceCreate {
	dc.mutation.SetType(s)
	return dc
}

// SetDescription sets the "description" field.
func (dc *DeviceCreate) SetDescription(s string) *DeviceCreate {
	dc.mutation.SetDescription(s)
	return dc
}

// SetID sets the "id" field.
func (dc *DeviceCreate) SetID(u uint64) *DeviceCreate {
	dc.mutation.SetID(u)
	return dc
}

// SetDepartmentsID sets the "departments" edge to the Department entity by ID.
func (dc *DeviceCreate) SetDepartmentsID(id uint64) *DeviceCreate {
	dc.mutation.SetDepartmentsID(id)
	return dc
}

// SetDepartments sets the "departments" edge to the Department entity.
func (dc *DeviceCreate) SetDepartments(d *Department) *DeviceCreate {
	return dc.SetDepartmentsID(d.ID)
}

// AddAccountIDs adds the "accounts" edge to the Account entity by IDs.
func (dc *DeviceCreate) AddAccountIDs(ids ...uint64) *DeviceCreate {
	dc.mutation.AddAccountIDs(ids...)
	return dc
}

// AddAccounts adds the "accounts" edges to the Account entity.
func (dc *DeviceCreate) AddAccounts(a ...*Account) *DeviceCreate {
	ids := make([]uint64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return dc.AddAccountIDs(ids...)
}

// Mutation returns the DeviceMutation object of the builder.
func (dc *DeviceCreate) Mutation() *DeviceMutation {
	return dc.mutation
}

// Save creates the Device in the database.
func (dc *DeviceCreate) Save(ctx context.Context) (*Device, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeviceCreate) SaveX(ctx context.Context) *Device {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DeviceCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DeviceCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DeviceCreate) defaults() {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := device.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		v := device.DefaultUpdatedAt()
		dc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeviceCreate) check() error {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Device.created_at"`)}
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Device.updated_at"`)}
	}
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Device.name"`)}
	}
	if v, ok := dc.mutation.Name(); ok {
		if err := device.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Device.name": %w`, err)}
		}
	}
	if _, ok := dc.mutation.DepartmentID(); !ok {
		return &ValidationError{Name: "department_id", err: errors.New(`ent: missing required field "Device.department_id"`)}
	}
	if v, ok := dc.mutation.DepartmentID(); ok {
		if err := device.DepartmentIDValidator(v); err != nil {
			return &ValidationError{Name: "department_id", err: fmt.Errorf(`ent: validator failed for field "Device.department_id": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Host(); !ok {
		return &ValidationError{Name: "host", err: errors.New(`ent: missing required field "Device.host"`)}
	}
	if v, ok := dc.mutation.Host(); ok {
		if err := device.HostValidator(v); err != nil {
			return &ValidationError{Name: "host", err: fmt.Errorf(`ent: validator failed for field "Device.host": %w`, err)}
		}
	}
	if _, ok := dc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Device.type"`)}
	}
	if v, ok := dc.mutation.GetType(); ok {
		if err := device.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Device.type": %w`, err)}
		}
	}
	if _, ok := dc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Device.description"`)}
	}
	if v, ok := dc.mutation.Description(); ok {
		if err := device.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Device.description": %w`, err)}
		}
	}
	if _, ok := dc.mutation.DepartmentsID(); !ok {
		return &ValidationError{Name: "departments", err: errors.New(`ent: missing required edge "Device.departments"`)}
	}
	return nil
}

func (dc *DeviceCreate) sqlSave(ctx context.Context) (*Device, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DeviceCreate) createSpec() (*Device, *sqlgraph.CreateSpec) {
	var (
		_node = &Device{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(device.Table, sqlgraph.NewFieldSpec(device.FieldID, field.TypeUint64))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(device.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.SetField(device.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(device.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.Host(); ok {
		_spec.SetField(device.FieldHost, field.TypeString, value)
		_node.Host = value
	}
	if value, ok := dc.mutation.GetType(); ok {
		_spec.SetField(device.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := dc.mutation.Description(); ok {
		_spec.SetField(device.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := dc.mutation.DepartmentsIDs(); len(nodes) > 0 {
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
		_node.DepartmentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.AccountsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DeviceCreateBulk is the builder for creating many Device entities in bulk.
type DeviceCreateBulk struct {
	config
	err      error
	builders []*DeviceCreate
}

// Save creates the Device entities in the database.
func (dcb *DeviceCreateBulk) Save(ctx context.Context) ([]*Device, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Device, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeviceMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeviceCreateBulk) SaveX(ctx context.Context) []*Device {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DeviceCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DeviceCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
