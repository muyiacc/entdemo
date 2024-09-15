package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("hash").Unique(),
		field.String("title"),
		field.String("dt"),
		field.String("cat"),
		field.Int("size"),
		field.String("ext_id").Optional(),
		field.String("imdb").Optional(),
	}
}

// Indexes of the Item.
func (Item) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("cat").Unique(),
		index.Fields("ext_id").Unique(),
		index.Fields("imdb").Unique(),
	}
}
