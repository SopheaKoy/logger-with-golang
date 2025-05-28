package schema

import "entgo.io/ent"

// Track holds the schema definition for the Track entity.
type Track struct {
	ent.Schema
}

// Fields of the Track.
func (Track) Fields() []ent.Field {
	return nil
}

// Edges of the Track.
func (Track) Edges() []ent.Edge {
	return nil
}
