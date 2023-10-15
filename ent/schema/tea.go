package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Tea holds the schema definition for the Tea entity.
type Tea struct {
	ent.Schema
}

// Fields of the Tea.
func (Tea) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(100)",
			}),
		field.String("color").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(100)",
			}),
	}
}

// Edges of the Tea.
func (Tea) Edges() []ent.Edge {
	return nil
}
