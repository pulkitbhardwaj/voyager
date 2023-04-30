package schema

import "entgo.io/ent"

// Lead holds the schema definition for the Lead entity.
type Lead struct {
	ent.Schema
}

// Fields of the Lead.
func (Lead) Fields() []ent.Field {
	return nil
}

// Edges of the Lead.
func (Lead) Edges() []ent.Edge {
	return nil
}
