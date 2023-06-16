package pubsubinfo

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

// Extension is an implementation of entc.Extension that adds all the templates
// that entx needs.
type Extension struct {
	entc.DefaultExtension

	templates []*gen.Template
}

// ExtensionOption allow for control over the behavior of the generator
type ExtensionOption func(*Extension) error

// NewExtension returns an entc Extension that allows the entx package to generate
// the schema changes and templates needed to function
func NewExtension(opts ...ExtensionOption) (*Extension, error) {
	e := &Extension{
		templates: MixinTemplates,
	}

	for _, opt := range opts {
		if err := opt(e); err != nil {
			return nil, err
		}
	}

	return e, nil
}

// Templates of the extension
func (e *Extension) Templates() []*gen.Template {
	return e.templates
}

var _ entc.Extension = (*Extension)(nil)
