package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

// Fields of the Device.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.String("manufacturer").
			Immutable(),
		field.String("model").
			Immutable(),
		field.String("build_number").
			Immutable(),
		field.String("os").
			Immutable(),
		field.Uint32("screen_width").
			Immutable(),
		field.Uint32("screen_height").
			Immutable(),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("devices").
			Required(),
		edge.To("telemetries", Telemetry.Type),
	}
}

// Indexes of the Device.
func (Device) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("manufacturer", "model", "build_number").
			Unique(),
	}
}
