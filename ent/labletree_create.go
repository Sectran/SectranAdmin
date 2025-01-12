// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sectran_admin/ent/labletree"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LableTreeCreate is the builder for creating a LableTree entity.
type LableTreeCreate struct {
	config
	mutation *LableTreeMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ltc *LableTreeCreate) SetCreatedAt(t time.Time) *LableTreeCreate {
	ltc.mutation.SetCreatedAt(t)
	return ltc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ltc *LableTreeCreate) SetNillableCreatedAt(t *time.Time) *LableTreeCreate {
	if t != nil {
		ltc.SetCreatedAt(*t)
	}
	return ltc
}

// SetUpdatedAt sets the "updated_at" field.
func (ltc *LableTreeCreate) SetUpdatedAt(t time.Time) *LableTreeCreate {
	ltc.mutation.SetUpdatedAt(t)
	return ltc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ltc *LableTreeCreate) SetNillableUpdatedAt(t *time.Time) *LableTreeCreate {
	if t != nil {
		ltc.SetUpdatedAt(*t)
	}
	return ltc
}

// SetName sets the "name" field.
func (ltc *LableTreeCreate) SetName(s string) *LableTreeCreate {
	ltc.mutation.SetName(s)
	return ltc
}

// SetType sets the "type" field.
func (ltc *LableTreeCreate) SetType(u uint) *LableTreeCreate {
	ltc.mutation.SetType(u)
	return ltc
}

// SetIcon sets the "icon" field.
func (ltc *LableTreeCreate) SetIcon(s string) *LableTreeCreate {
	ltc.mutation.SetIcon(s)
	return ltc
}

// SetContent sets the "content" field.
func (ltc *LableTreeCreate) SetContent(s string) *LableTreeCreate {
	ltc.mutation.SetContent(s)
	return ltc
}

// SetOwnership sets the "ownership" field.
func (ltc *LableTreeCreate) SetOwnership(u uint8) *LableTreeCreate {
	ltc.mutation.SetOwnership(u)
	return ltc
}

// SetOwnerID sets the "owner_id" field.
func (ltc *LableTreeCreate) SetOwnerID(u uint64) *LableTreeCreate {
	ltc.mutation.SetOwnerID(u)
	return ltc
}

// SetParentID sets the "parent_id" field.
func (ltc *LableTreeCreate) SetParentID(u uint64) *LableTreeCreate {
	ltc.mutation.SetParentID(u)
	return ltc
}

// SetDescription sets the "description" field.
func (ltc *LableTreeCreate) SetDescription(s string) *LableTreeCreate {
	ltc.mutation.SetDescription(s)
	return ltc
}

// SetTargetType sets the "target_type" field.
func (ltc *LableTreeCreate) SetTargetType(u uint16) *LableTreeCreate {
	ltc.mutation.SetTargetType(u)
	return ltc
}

// SetParents sets the "parents" field.
func (ltc *LableTreeCreate) SetParents(s string) *LableTreeCreate {
	ltc.mutation.SetParents(s)
	return ltc
}

// SetInherit sets the "inherit" field.
func (ltc *LableTreeCreate) SetInherit(b bool) *LableTreeCreate {
	ltc.mutation.SetInherit(b)
	return ltc
}

// SetID sets the "id" field.
func (ltc *LableTreeCreate) SetID(u uint64) *LableTreeCreate {
	ltc.mutation.SetID(u)
	return ltc
}

// Mutation returns the LableTreeMutation object of the builder.
func (ltc *LableTreeCreate) Mutation() *LableTreeMutation {
	return ltc.mutation
}

// Save creates the LableTree in the database.
func (ltc *LableTreeCreate) Save(ctx context.Context) (*LableTree, error) {
	ltc.defaults()
	return withHooks(ctx, ltc.sqlSave, ltc.mutation, ltc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ltc *LableTreeCreate) SaveX(ctx context.Context) *LableTree {
	v, err := ltc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ltc *LableTreeCreate) Exec(ctx context.Context) error {
	_, err := ltc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ltc *LableTreeCreate) ExecX(ctx context.Context) {
	if err := ltc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ltc *LableTreeCreate) defaults() {
	if _, ok := ltc.mutation.CreatedAt(); !ok {
		v := labletree.DefaultCreatedAt()
		ltc.mutation.SetCreatedAt(v)
	}
	if _, ok := ltc.mutation.UpdatedAt(); !ok {
		v := labletree.DefaultUpdatedAt()
		ltc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ltc *LableTreeCreate) check() error {
	if _, ok := ltc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "LableTree.created_at"`)}
	}
	if _, ok := ltc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "LableTree.updated_at"`)}
	}
	if _, ok := ltc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "LableTree.name"`)}
	}
	if v, ok := ltc.mutation.Name(); ok {
		if err := labletree.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "LableTree.name": %w`, err)}
		}
	}
	if _, ok := ltc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "LableTree.type"`)}
	}
	if _, ok := ltc.mutation.Icon(); !ok {
		return &ValidationError{Name: "icon", err: errors.New(`ent: missing required field "LableTree.icon"`)}
	}
	if v, ok := ltc.mutation.Icon(); ok {
		if err := labletree.IconValidator(v); err != nil {
			return &ValidationError{Name: "icon", err: fmt.Errorf(`ent: validator failed for field "LableTree.icon": %w`, err)}
		}
	}
	if _, ok := ltc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "LableTree.content"`)}
	}
	if v, ok := ltc.mutation.Content(); ok {
		if err := labletree.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "LableTree.content": %w`, err)}
		}
	}
	if _, ok := ltc.mutation.Ownership(); !ok {
		return &ValidationError{Name: "ownership", err: errors.New(`ent: missing required field "LableTree.ownership"`)}
	}
	if _, ok := ltc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`ent: missing required field "LableTree.owner_id"`)}
	}
	if _, ok := ltc.mutation.ParentID(); !ok {
		return &ValidationError{Name: "parent_id", err: errors.New(`ent: missing required field "LableTree.parent_id"`)}
	}
	if _, ok := ltc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "LableTree.description"`)}
	}
	if v, ok := ltc.mutation.Description(); ok {
		if err := labletree.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "LableTree.description": %w`, err)}
		}
	}
	if _, ok := ltc.mutation.TargetType(); !ok {
		return &ValidationError{Name: "target_type", err: errors.New(`ent: missing required field "LableTree.target_type"`)}
	}
	if _, ok := ltc.mutation.Parents(); !ok {
		return &ValidationError{Name: "parents", err: errors.New(`ent: missing required field "LableTree.parents"`)}
	}
	if _, ok := ltc.mutation.Inherit(); !ok {
		return &ValidationError{Name: "inherit", err: errors.New(`ent: missing required field "LableTree.inherit"`)}
	}
	return nil
}

func (ltc *LableTreeCreate) sqlSave(ctx context.Context) (*LableTree, error) {
	if err := ltc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ltc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ltc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	ltc.mutation.id = &_node.ID
	ltc.mutation.done = true
	return _node, nil
}

func (ltc *LableTreeCreate) createSpec() (*LableTree, *sqlgraph.CreateSpec) {
	var (
		_node = &LableTree{config: ltc.config}
		_spec = sqlgraph.NewCreateSpec(labletree.Table, sqlgraph.NewFieldSpec(labletree.FieldID, field.TypeUint64))
	)
	if id, ok := ltc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ltc.mutation.CreatedAt(); ok {
		_spec.SetField(labletree.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ltc.mutation.UpdatedAt(); ok {
		_spec.SetField(labletree.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ltc.mutation.Name(); ok {
		_spec.SetField(labletree.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ltc.mutation.GetType(); ok {
		_spec.SetField(labletree.FieldType, field.TypeUint, value)
		_node.Type = value
	}
	if value, ok := ltc.mutation.Icon(); ok {
		_spec.SetField(labletree.FieldIcon, field.TypeString, value)
		_node.Icon = value
	}
	if value, ok := ltc.mutation.Content(); ok {
		_spec.SetField(labletree.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := ltc.mutation.Ownership(); ok {
		_spec.SetField(labletree.FieldOwnership, field.TypeUint8, value)
		_node.Ownership = value
	}
	if value, ok := ltc.mutation.OwnerID(); ok {
		_spec.SetField(labletree.FieldOwnerID, field.TypeUint64, value)
		_node.OwnerID = value
	}
	if value, ok := ltc.mutation.ParentID(); ok {
		_spec.SetField(labletree.FieldParentID, field.TypeUint64, value)
		_node.ParentID = value
	}
	if value, ok := ltc.mutation.Description(); ok {
		_spec.SetField(labletree.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ltc.mutation.TargetType(); ok {
		_spec.SetField(labletree.FieldTargetType, field.TypeUint16, value)
		_node.TargetType = value
	}
	if value, ok := ltc.mutation.Parents(); ok {
		_spec.SetField(labletree.FieldParents, field.TypeString, value)
		_node.Parents = value
	}
	if value, ok := ltc.mutation.Inherit(); ok {
		_spec.SetField(labletree.FieldInherit, field.TypeBool, value)
		_node.Inherit = value
	}
	return _node, _spec
}

// LableTreeCreateBulk is the builder for creating many LableTree entities in bulk.
type LableTreeCreateBulk struct {
	config
	err      error
	builders []*LableTreeCreate
}

// Save creates the LableTree entities in the database.
func (ltcb *LableTreeCreateBulk) Save(ctx context.Context) ([]*LableTree, error) {
	if ltcb.err != nil {
		return nil, ltcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ltcb.builders))
	nodes := make([]*LableTree, len(ltcb.builders))
	mutators := make([]Mutator, len(ltcb.builders))
	for i := range ltcb.builders {
		func(i int, root context.Context) {
			builder := ltcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LableTreeMutation)
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
					_, err = mutators[i+1].Mutate(root, ltcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ltcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ltcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ltcb *LableTreeCreateBulk) SaveX(ctx context.Context) []*LableTree {
	v, err := ltcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ltcb *LableTreeCreateBulk) Exec(ctx context.Context) error {
	_, err := ltcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ltcb *LableTreeCreateBulk) ExecX(ctx context.Context) {
	if err := ltcb.Exec(ctx); err != nil {
		panic(err)
	}
}
