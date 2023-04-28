// Code generated by ent, DO NOT EDIT.

package entity

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/pulkitbhardwaj/voyager/pkgs/goatcx/crm/entity/lead"
)

// LeadCreate is the builder for creating a Lead entity.
type LeadCreate struct {
	config
	mutation *LeadMutation
	hooks    []Hook
}

// Mutation returns the LeadMutation object of the builder.
func (lc *LeadCreate) Mutation() *LeadMutation {
	return lc.mutation
}

// Save creates the Lead in the database.
func (lc *LeadCreate) Save(ctx context.Context) (*Lead, error) {
	return withHooks[*Lead, LeadMutation](ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LeadCreate) SaveX(ctx context.Context) *Lead {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LeadCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LeadCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LeadCreate) check() error {
	return nil
}

func (lc *LeadCreate) sqlSave(ctx context.Context) (*Lead, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LeadCreate) createSpec() (*Lead, *sqlgraph.CreateSpec) {
	var (
		_node = &Lead{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(lead.Table, sqlgraph.NewFieldSpec(lead.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// LeadCreateBulk is the builder for creating many Lead entities in bulk.
type LeadCreateBulk struct {
	config
	builders []*LeadCreate
}

// Save creates the Lead entities in the database.
func (lcb *LeadCreateBulk) Save(ctx context.Context) ([]*Lead, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Lead, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LeadMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LeadCreateBulk) SaveX(ctx context.Context) []*Lead {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LeadCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LeadCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
