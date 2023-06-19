package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"go.infratographer.com/ipam-api/x/pubsubinfo"

	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/gidx"
)

// IPBlockType holds the schema definition for the IPBlockType entity.
type IPBlockType struct {
	ent.Schema
}

// Mixin of the IP Block Type type
func (IPBlockType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the IPBlockType.
func (IPBlockType) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the IP Block Type.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(IPBlockTypePrefix) }),
		field.Text("name").
			NotEmpty().
			Comment("The name of the ip block type.").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.String("owner_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the owner for this ip block type.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("OWNER"),
				pubsubinfo.AdditionalSubject(),
			),
	}
}

// Edges of the IPBlockType
func (IPBlockType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ip_block", IPBlock.Type).
			Ref("ip_block_type").
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
				entgql.RelayConnection(),
			),
	}
}

// Indexes of the IPBlockType.
func (IPBlockType) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_id"),
	}
}

// Annotations of the IPBlockType
func (IPBlockType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		pubsubinfo.Annotation{},
		entx.GraphKeyDirective("id"),
		schema.Comment("Represents an ip block type node on the graph."),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Create a new ip block type node."),
			entgql.MutationUpdate().Description("Update an existing ip block type node."),
		),
	}
}
