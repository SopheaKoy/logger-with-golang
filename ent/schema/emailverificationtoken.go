package schema

import "entgo.io/ent"

// EmailVerificationToken holds the schema definition for the EmailVerificationToken entity.
type EmailVerificationToken struct {
	ent.Schema
}

// Fields of the EmailVerificationToken.
func (EmailVerificationToken) Fields() []ent.Field {
	return nil
}

// Edges of the EmailVerificationToken.
func (EmailVerificationToken) Edges() []ent.Edge {
	return nil
}
