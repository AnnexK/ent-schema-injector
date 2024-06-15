package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	injector "github.com/AnnexK/ent-schema-injector"
)

const schemaPath = "./ent/schema"

func main() {
	err := entc.Generate(schemaPath, &gen.Config{}, entc.Extensions(&injector.InjectionExtension{
		SchemaName: "Event",
		SchemaPath: schemaPath,
	}))

	if err != nil {
		log.Fatal("generation failed", err.Error())
	}	
}
