package schema

import "entgo.io/ent"

// UserSession holds the schema definition for the UserSession entity.
type UserSession struct {
	ent.Schema
}

// Fields of the UserSession.
func (UserSession) Fields() []ent.Field {
	return nil
}

// Edges of the UserSession.
func (UserSession) Edges() []ent.Edge {
	return nil
}
