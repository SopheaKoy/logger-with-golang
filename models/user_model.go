package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type UserModel struct {
	ent.Schema
}

// Fields of the User.
func (UserModel) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable(),

		field.String("first_name").
			NotEmpty(),

		field.String("last_name").
			NotEmpty(),

		field.String("email").
			NotEmpty().
			Unique(),

		field.String("password").
			NotEmpty().
			Sensitive(),

		field.Time("created_at").
			Default(time.Now),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),

		field.Time("deleted_at").
			Optional().
			Nillable(),
	}
}

// Edges of the User.
func (UserModel) Edges() []ent.Edge {
	return nil
}
