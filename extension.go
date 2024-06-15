package injector

import "entgo.io/ent/entc"

type InjectionExtension struct {
	entc.DefaultExtension
	SchemaPath string
	SchemaName string
}

func (ext *InjectionExtension) Options() []entc.Option {
	return []entc.Option{
		InjectSchema(ext.SchemaPath, ext.SchemaName),
	}
}
