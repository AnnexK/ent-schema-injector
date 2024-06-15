package injector

import (
	"fmt"

	"entgo.io/contrib/schemast"
	"entgo.io/ent"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

// InjectSchema создает хук, внедряющий схему в граф.
func InjectSchema(schemaPath, typeName string) entc.Option {
	return func(c *gen.Config) error {
		ctx, err := schemast.Load(schemaPath)
		if err != nil {
			return fmt.Errorf("failed to load schema context: %w", err)
		}

		if ctx.HasType(typeName) {
			return fmt.Errorf("type %s already present in schema", typeName)
		}

		err = schemast.Mutate(ctx, mapTypeToMutation(TestSchema{}, typeName))
		if err != nil {
			return fmt.Errorf("failed to inject type: %w", err)
		}

		// Записать изменения схемы
		err = ctx.Print(schemaPath)
		if err != nil {
			return fmt.Errorf("failed to write injection: %w", err)
		}
		
		return nil
	}
}

func mapTypeToMutation(typ ent.Interface, typeName string) schemast.Mutator {
	return &schemast.UpsertSchema{
		Name: typeName,
		Fields: typ.Fields(),
		Edges: typ.Edges(),
		Annotations: typ.Annotations(),
		Indexes: typ.Indexes(),
	}
}
