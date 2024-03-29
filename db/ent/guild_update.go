// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"voidmanager/db/ent/guild"
	"voidmanager/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GuildUpdate is the builder for updating Guild entities.
type GuildUpdate struct {
	config
	hooks    []Hook
	mutation *GuildMutation
}

// Where appends a list predicates to the GuildUpdate builder.
func (gu *GuildUpdate) Where(ps ...predicate.Guild) *GuildUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetBotPrefix sets the "bot_prefix" field.
func (gu *GuildUpdate) SetBotPrefix(s string) *GuildUpdate {
	gu.mutation.SetBotPrefix(s)
	return gu
}

// SetNillableBotPrefix sets the "bot_prefix" field if the given value is not nil.
func (gu *GuildUpdate) SetNillableBotPrefix(s *string) *GuildUpdate {
	if s != nil {
		gu.SetBotPrefix(*s)
	}
	return gu
}

// Mutation returns the GuildMutation object of the builder.
func (gu *GuildUpdate) Mutation() *GuildMutation {
	return gu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GuildUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gu.hooks) == 0 {
		if err = gu.check(); err != nil {
			return 0, err
		}
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GuildMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gu.check(); err != nil {
				return 0, err
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			if gu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GuildUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GuildUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GuildUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GuildUpdate) check() error {
	if v, ok := gu.mutation.BotPrefix(); ok {
		if err := guild.BotPrefixValidator(v); err != nil {
			return &ValidationError{Name: "bot_prefix", err: fmt.Errorf(`ent: validator failed for field "Guild.bot_prefix": %w`, err)}
		}
	}
	return nil
}

func (gu *GuildUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   guild.Table,
			Columns: guild.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: guild.FieldID,
			},
		},
	}
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.BotPrefix(); ok {
		_spec.SetField(guild.FieldBotPrefix, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{guild.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GuildUpdateOne is the builder for updating a single Guild entity.
type GuildUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GuildMutation
}

// SetBotPrefix sets the "bot_prefix" field.
func (guo *GuildUpdateOne) SetBotPrefix(s string) *GuildUpdateOne {
	guo.mutation.SetBotPrefix(s)
	return guo
}

// SetNillableBotPrefix sets the "bot_prefix" field if the given value is not nil.
func (guo *GuildUpdateOne) SetNillableBotPrefix(s *string) *GuildUpdateOne {
	if s != nil {
		guo.SetBotPrefix(*s)
	}
	return guo
}

// Mutation returns the GuildMutation object of the builder.
func (guo *GuildUpdateOne) Mutation() *GuildMutation {
	return guo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GuildUpdateOne) Select(field string, fields ...string) *GuildUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Guild entity.
func (guo *GuildUpdateOne) Save(ctx context.Context) (*Guild, error) {
	var (
		err  error
		node *Guild
	)
	if len(guo.hooks) == 0 {
		if err = guo.check(); err != nil {
			return nil, err
		}
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GuildMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = guo.check(); err != nil {
				return nil, err
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			if guo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = guo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, guo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (guo *GuildUpdateOne) SaveX(ctx context.Context) *Guild {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GuildUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GuildUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GuildUpdateOne) check() error {
	if v, ok := guo.mutation.BotPrefix(); ok {
		if err := guild.BotPrefixValidator(v); err != nil {
			return &ValidationError{Name: "bot_prefix", err: fmt.Errorf(`ent: validator failed for field "Guild.bot_prefix": %w`, err)}
		}
	}
	return nil
}

func (guo *GuildUpdateOne) sqlSave(ctx context.Context) (_node *Guild, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   guild.Table,
			Columns: guild.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: guild.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Guild.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, guild.FieldID)
		for _, f := range fields {
			if !guild.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != guild.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.BotPrefix(); ok {
		_spec.SetField(guild.FieldBotPrefix, field.TypeString, value)
	}
	_node = &Guild{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{guild.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
