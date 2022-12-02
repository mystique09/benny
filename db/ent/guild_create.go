// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"voidmanager/db/ent/guild"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GuildCreate is the builder for creating a Guild entity.
type GuildCreate struct {
	config
	mutation *GuildMutation
	hooks    []Hook
}

// SetBotPrefix sets the "bot_prefix" field.
func (gc *GuildCreate) SetBotPrefix(s string) *GuildCreate {
	gc.mutation.SetBotPrefix(s)
	return gc
}

// SetNillableBotPrefix sets the "bot_prefix" field if the given value is not nil.
func (gc *GuildCreate) SetNillableBotPrefix(s *string) *GuildCreate {
	if s != nil {
		gc.SetBotPrefix(*s)
	}
	return gc
}

// SetID sets the "id" field.
func (gc *GuildCreate) SetID(s string) *GuildCreate {
	gc.mutation.SetID(s)
	return gc
}

// Mutation returns the GuildMutation object of the builder.
func (gc *GuildCreate) Mutation() *GuildMutation {
	return gc.mutation
}

// Save creates the Guild in the database.
func (gc *GuildCreate) Save(ctx context.Context) (*Guild, error) {
	var (
		err  error
		node *Guild
	)
	gc.defaults()
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GuildMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			if node, err = gc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			if gc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Guild)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GuildMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GuildCreate) SaveX(ctx context.Context) *Guild {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GuildCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GuildCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GuildCreate) defaults() {
	if _, ok := gc.mutation.BotPrefix(); !ok {
		v := guild.DefaultBotPrefix
		gc.mutation.SetBotPrefix(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GuildCreate) check() error {
	if _, ok := gc.mutation.BotPrefix(); !ok {
		return &ValidationError{Name: "bot_prefix", err: errors.New(`ent: missing required field "Guild.bot_prefix"`)}
	}
	if v, ok := gc.mutation.BotPrefix(); ok {
		if err := guild.BotPrefixValidator(v); err != nil {
			return &ValidationError{Name: "bot_prefix", err: fmt.Errorf(`ent: validator failed for field "Guild.bot_prefix": %w`, err)}
		}
	}
	if v, ok := gc.mutation.ID(); ok {
		if err := guild.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Guild.id": %w`, err)}
		}
	}
	return nil
}

func (gc *GuildCreate) sqlSave(ctx context.Context) (*Guild, error) {
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Guild.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (gc *GuildCreate) createSpec() (*Guild, *sqlgraph.CreateSpec) {
	var (
		_node = &Guild{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: guild.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: guild.FieldID,
			},
		}
	)
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gc.mutation.BotPrefix(); ok {
		_spec.SetField(guild.FieldBotPrefix, field.TypeString, value)
		_node.BotPrefix = value
	}
	return _node, _spec
}

// GuildCreateBulk is the builder for creating many Guild entities in bulk.
type GuildCreateBulk struct {
	config
	builders []*GuildCreate
}

// Save creates the Guild entities in the database.
func (gcb *GuildCreateBulk) Save(ctx context.Context) ([]*Guild, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Guild, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GuildMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GuildCreateBulk) SaveX(ctx context.Context) []*Guild {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GuildCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GuildCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}