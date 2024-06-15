package injector

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type TestSchema struct {
	ent.Schema
}

func (TestSchema) Fields() []ent.Field {
	return []ent.Field{
		field.String("Field1").Optional(),
		field.Int("Foo"),
		field.Bool("Bar"),
	}
}

func (TestSchema) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("Foo", "Bar").Unique(),
	}
}
