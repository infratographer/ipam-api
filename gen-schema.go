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

	"github.com/vektah/gqlparser/v2/formatter"

	"go.infratographer.com/ipam-api/internal/graphapi"
)

// read in schema from internal package and save it to the schema file
func main() {
	execSchema := graphapi.NewExecutableSchema(graphapi.Config{})
	schema := execSchema.Schema()

	// Some of our federation fields get marked as "BuiltIn" by gengql and the formatter doesn't print builtin types, this adds them for us.
	entities := schema.Types["_Entity"]
	entities.BuiltIn = false
	service := schema.Types["_Service"]
	service.BuiltIn = false
	// entities.Position.Src.BuiltIn = false

	f, err := os.Create("schema.graphql")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmtr := formatter.NewFormatter(f)

	fmtr.FormatSchema(schema)

	f.Write(federationSchema)
}

var federationSchema = []byte(`scalar _Any
scalar FieldSet
directive @requires(fields: FieldSet!) on FIELD_DEFINITION
directive @provides(fields: FieldSet!) on FIELD_DEFINITION
directive @extends on OBJECT | INTERFACE
directive @key(fields: FieldSet!, resolvable: Boolean = true) repeatable on OBJECT | INTERFACE
directive @link(import: [String!], url: String!) repeatable on SCHEMA
directive @external on FIELD_DEFINITION | OBJECT
directive @shareable on OBJECT | FIELD_DEFINITION
directive @tag(name: String!) repeatable on FIELD_DEFINITION | INTERFACE | OBJECT | UNION | ARGUMENT_DEFINITION | SCALAR | ENUM | ENUM_VALUE | INPUT_OBJECT | INPUT_FIELD_DEFINITION
directive @override(from: String!) on FIELD_DEFINITION
directive @inaccessible on SCALAR | OBJECT | FIELD_DEFINITION | ARGUMENT_DEFINITION | INTERFACE | UNION | ENUM | ENUM_VALUE | INPUT_OBJECT | INPUT_FIELD_DEFINITION
#directive @interfaceObject on OBJECT
extend schema
  @link(
	url: "https://specs.apollo.dev/federation/v2.3"
	import: [
	  "@key",
	  "@external",
	  "@shareable",
	  "@tag",
	  "@override",
	  "@inaccessible",
	  "@interfaceObject"
	  ]
  )
`)
