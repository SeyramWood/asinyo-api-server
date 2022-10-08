package schema

import "entgo.io/ent"

// AgentRequest holds the schema definition for the AgentRequest entity.
type AgentRequest struct {
	ent.Schema
}

func (AgentRequest) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the AgentRequest.
func (AgentRequest) Fields() []ent.Field {
	return nil
}

// Edges of the AgentRequest.
func (AgentRequest) Edges() []ent.Edge {
	return nil
}
