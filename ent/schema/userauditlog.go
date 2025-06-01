package schema

import "entgo.io/ent"

// UserAuditLog holds the schema definition for the UserAuditLog entity.
type UserAuditLog struct {
	ent.Schema
}

// Fields of the UserAuditLog.
func (UserAuditLog) Fields() []ent.Field {
	return nil
}

// Edges of the UserAuditLog.
func (UserAuditLog) Edges() []ent.Edge {
	return nil
}
