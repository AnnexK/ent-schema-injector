package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Event struct {
	ent.Schema
}

func (Event) Fields() []ent.Field {
	return []ent.Field{field.String("Field1").Optional(), field.Int("Foo"), field.Bool("Bar")}
}
func (Event) Edges() []ent.Edge {
	return nil
}
func (Event) Annotations() []schema.Annotation {
	return nil
}
func (Event) Indexes() []ent.Index {
	return []ent.Index{index.Fields("Foo", "Bar").Unique()}
}
