package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Guild struct {
	ent.Schema
}

func (Guild) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty(),
		field.String("bot_prefix").NotEmpty().Default("!"),
	}
}

func (Guild) Edges() []ent.Edge {
	return nil
}
