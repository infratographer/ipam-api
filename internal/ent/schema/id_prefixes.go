package schema

import (
	"entgo.io/contrib/entgql"
	"github.com/vektah/gqlparser/v2/ast"
)

const (
	// ApplicationPrefix is the prefix for all application IDs owned by ipam-api
	ApplicationPrefix = "ipam"
	// IPBlockTypePrefix is the prefix for all IP Block Types nodes
	IPBlockTypePrefix = ApplicationPrefix + "ibt"
	// IPBlockPrefix is the prefix for all IP Block nodes
	IPBlockPrefix = ApplicationPrefix + "ibk"
	// IPAddressPrefix is the prefix for all IP Block nodes
	IPAddressPrefix = ApplicationPrefix + "ipa"
)

func prefixIDDirective(prefix string) entgql.Annotation {
	var args []*ast.Argument
	if prefix != "" {
		args = append(args, &ast.Argument{
			Name: "prefix",
			Value: &ast.Value{
				Raw:  prefix,
				Kind: ast.StringValue,
			},
		})
	}

	return entgql.Directives(entgql.NewDirective("prefixedID", args...))
}
