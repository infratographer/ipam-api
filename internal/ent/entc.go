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

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/x/entx"
)

func main() {
	// Ensure the schema directory exists before running entc.
	_ = os.Mkdir("schema", 0755)

	xExt, err := entx.NewExtension(
		entx.WithFederation(),
		// Disable the default event hooks generation, until the ExtraAdditionalSubjects support is added to x/entx
		// entx.WithEventHooks(),
		entx.WithJSONScalar(),
	)
	if err != nil {
		log.Fatalf("creating entx extension: %v", err)
	}

	gqlExt, err := entgql.NewExtension(
		// Tell Ent to generate a GraphQL schema for
		// the Ent schema in a file named ent.graphql.
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("schema/ent.graphql"),
		entgql.WithConfigPath("gqlgen.yml"),
		entgql.WithWhereInputs(true),
		entgql.WithSchemaHook(xExt.GQLSchemaHooks()...),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		entc.Extensions(
			xExt,
			gqlExt,
		),
		entc.Dependency(
			entc.DependencyName("EventsPublisher"),
			entc.DependencyTypeInfo(&field.TypeInfo{
				Ident:   "events.Connection",
				PkgPath: "go.infratographer.com/x/events",
			}),
		),
		// entc.TemplateDir("./internal/ent/templates"),
		// entc.FeatureNames("intercept"),
	}

	if err := entc.Generate("./internal/ent/schema", &gen.Config{
		Target:   "./internal/ent/generated",
		Package:  "go.infratographer.com/ipam-api/internal/ent/generated",
		Header:   entx.CopyrightHeader,
		Features: []gen.Feature{gen.FeatureVersionedMigration},
	}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
