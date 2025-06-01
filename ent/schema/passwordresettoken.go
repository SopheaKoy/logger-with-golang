package schema

import "entgo.io/ent"

// PasswordResetToken holds the schema definition for the PasswordResetToken entity.
type PasswordResetToken struct {
	ent.Schema
}

// Fields of the PasswordResetToken.
func (PasswordResetToken) Fields() []ent.Field {
	return nil
}

// Edges of the PasswordResetToken.
func (PasswordResetToken) Edges() []ent.Edge {
	return nil
}
