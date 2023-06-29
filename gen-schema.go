// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build ignore

package main

import (
	"log"
	"os"

	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/formatter"

	"go.infratographer.com/ipam-api/internal/graphapi"
)

// read in schema from internal package and save it to the schema file
func main() {
	execSchema := graphapi.NewExecutableSchema(graphapi.Config{})
	schema := execSchema.Schema()

	// remove codegen directives that we don't want in published schema
	for _, t := range schema.Types {
		dirs := ast.DirectiveList{}
		for _, td := range t.Directives {
			switch td.Name {
			case "goField", "goModel":
				continue
			default:
				dirs = append(dirs, td)
			}
		}
		t.Directives = dirs

		for _, f := range t.Fields {
			dirs := ast.DirectiveList{}
			for _, fd := range f.Directives {
				switch fd.Name {
				case "goField", "goModel":
					continue
				default:
					dirs = append(dirs, fd)
				}
			}
			f.Directives = dirs
		}
	}

	delete(schema.Directives, "goField")
	delete(schema.Directives, "goModel")

	// Some of our federation fields get marked as "BuiltIn" by gengql and the formatter doesn't print builtin types, this adds them for us.
	entityType := schema.Types["_Entity"]
	entityType.BuiltIn = false
	serviceType := schema.Types["_Service"]
	serviceType.BuiltIn = false
	anyType := schema.Types["_Any"]
	anyType.BuiltIn = false

	f, err := os.Create("schema.graphql")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmtr := formatter.NewFormatter(f)
	fmtr.FormatSchema(schema)

	f.Write(federationSchema)

	// Write testclient schema, include all federation params
	// find the internal federation src and mark it as not builtin. "interfaceObject" is a federation directive,
	// so we use that to look up the source
	intObj := schema.Directives["interfaceObject"]
	intObj.Position.Src.BuiltIn = false
	schema.Types["FieldSet"].BuiltIn = false

	clientSchema, err := os.Create("internal/testclient/schema/schema.graphql")
	if err != nil {
		log.Fatal(err)
	}
	defer clientSchema.Close()

	fmtr = formatter.NewFormatter(clientSchema)
	fmtr.FormatSchema(schema)
}

var federationSchema = []byte(`
extend schema
  @link(
    url: "https://specs.apollo.dev/federation/v2.3"
    import: [
			"@key",
			"@interfaceObject",
			"@shareable",
			"@inaccessible",
			"@override",
			"@provides",
			"@requires",
			"@tag"
      ]
  )
`)
