package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type User struct {
    ent.Schema
}

func (User) Fields() []ent.Field {
    return []ent.Field{
        field.String("firstname").NotEmpty(),
        field.String("lastname").NotEmpty(),
    }
}

func (User) Edges() []ent.Edge {
    return nil
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:  "user",
			Schema: "users",
		},
	}
}