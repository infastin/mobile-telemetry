package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Telemetry holds the schema definition for the Telemetry entity.
type Telemetry struct {
	ent.Schema
}

// Fields of the Telemetry.
func (Telemetry) Fields() []ent.Field {
	return []ent.Field{
		field.String("app_version"),
		field.String("os_version"),
		field.String("action_type"),
		field.JSON("action_data", map[string]any{}),
		field.Time("action_at"),
	}
}

// Edges of the Telemetry.
func (Telemetry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("telemetries").
			Unique().
			Required().
			Immutable(),
		edge.From("device", Device.Type).
			Ref("telemetries").
			Unique().
			Required().
			Immutable(),
	}
}

// Indexes of the Telemetry.
func (Telemetry) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("user", "device"),
		index.Edges("user", "device").
			Fields("app_version"),
		index.Edges("user", "device").
			Fields("os_version"),
	}
}
