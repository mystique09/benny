// Code generated by ent, DO NOT EDIT.

package guild

import (
	"voidmanager/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BotPrefix applies equality check predicate on the "bot_prefix" field. It's identical to BotPrefixEQ.
func BotPrefix(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBotPrefix), v))
	})
}

// BotPrefixEQ applies the EQ predicate on the "bot_prefix" field.
func BotPrefixEQ(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBotPrefix), v))
	})
}

// BotPrefixNEQ applies the NEQ predicate on the "bot_prefix" field.
func BotPrefixNEQ(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBotPrefix), v))
	})
}

// BotPrefixIn applies the In predicate on the "bot_prefix" field.
func BotPrefixIn(vs ...string) predicate.Guild {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBotPrefix), v...))
	})
}

// BotPrefixNotIn applies the NotIn predicate on the "bot_prefix" field.
func BotPrefixNotIn(vs ...string) predicate.Guild {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBotPrefix), v...))
	})
}

// BotPrefixGT applies the GT predicate on the "bot_prefix" field.
func BotPrefixGT(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBotPrefix), v))
	})
}

// BotPrefixGTE applies the GTE predicate on the "bot_prefix" field.
func BotPrefixGTE(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBotPrefix), v))
	})
}

// BotPrefixLT applies the LT predicate on the "bot_prefix" field.
func BotPrefixLT(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBotPrefix), v))
	})
}

// BotPrefixLTE applies the LTE predicate on the "bot_prefix" field.
func BotPrefixLTE(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBotPrefix), v))
	})
}

// BotPrefixContains applies the Contains predicate on the "bot_prefix" field.
func BotPrefixContains(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBotPrefix), v))
	})
}

// BotPrefixHasPrefix applies the HasPrefix predicate on the "bot_prefix" field.
func BotPrefixHasPrefix(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBotPrefix), v))
	})
}

// BotPrefixHasSuffix applies the HasSuffix predicate on the "bot_prefix" field.
func BotPrefixHasSuffix(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBotPrefix), v))
	})
}

// BotPrefixEqualFold applies the EqualFold predicate on the "bot_prefix" field.
func BotPrefixEqualFold(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBotPrefix), v))
	})
}

// BotPrefixContainsFold applies the ContainsFold predicate on the "bot_prefix" field.
func BotPrefixContainsFold(v string) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBotPrefix), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Guild) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Guild) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Guild) predicate.Guild {
	return predicate.Guild(func(s *sql.Selector) {
		p(s.Not())
	})
}
