package pubsubinfo

import (
	"embed"
	"strings"
	"text/template"

	"entgo.io/ent/entc/gen"
)

var (
	// PubSubInfoTemplate adds support for generating pubsub fields
	PubSubInfoTemplate = parseT("template/pubsub.tmpl")

	// TemplateFuncs contains the extra template functions used by entx.
	TemplateFuncs = template.FuncMap{
		"contains": strings.Contains,
	}

	// MixinTemplates includes all templates for extending ent to support entx mixins.
	MixinTemplates = []*gen.Template{
		PubSubInfoTemplate,
	}

	//go:embed template/*
	_templates embed.FS
)

func parseT(path string) *gen.Template {
	return gen.MustParse(gen.NewTemplate(path).
		Funcs(TemplateFuncs).
		ParseFS(_templates, path))
}
