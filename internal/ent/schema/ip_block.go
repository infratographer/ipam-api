package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"go.infratographer.com/load-balancer-api/x/pubsubinfo"
	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/gidx"
)

// IPBlock holds the schema definition for the IPBlock entity.
type IPBlock struct {
	ent.Schema
}

// Mixin of the IP Block Type type
func (IPBlock) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the IPBlock.
func (IPBlock) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the IP Block.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(IPBlockPrefix) }),
		field.Text("prefix").
			NotEmpty().
			Comment("The prefix of the ip block.").
			Annotations(
				entgql.OrderField("PREFIX"),
			),
		field.String("block_type_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the block type for this ip block.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("BLOCK_TYPE"),
				pubsubinfo.AdditionalSubject(),
			),
		field.String("location_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the location for this ip block.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("LOCATION"),
				pubsubinfo.AdditionalSubject(),
			),
		field.String("parent_block_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the parent of this ip block.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("PARENT_BLOCK"),
				pubsubinfo.AdditionalSubject(),
			),
		field.Bool("allow_auto_subnet").
			Default(true).
			Comment("Allow carving this block into smaller subnets.").
			Annotations(
				entgql.OrderField("AUTOSUBNET"),
			),
		field.Bool("allow_auto_allocate").
			Default(true).
			Comment("Allow automatically assigning IPs directly from this block.").
			Annotations(
				entgql.OrderField("AUTOALLOCATE"),
			),
	}
}

// Edges of the IPBlock
func (IPBlock) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ip_block_type", IPBlockType.Type).
			Unique().
			Required().
			Immutable().
			Field("block_type_id").
			Annotations(),
		edge.From("ip_address", IPAddress.Type).
			Ref("ip_block").
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
				entgql.RelayConnection(),
			),
	}
}

// Indexes of the IPBlock.
func (IPBlock) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("block_type_id"),
	}
}

// Annotations of the IPBlock
func (IPBlock) Annotations() []schema.Annotation {
	return []schema.Annotation{
		pubsubinfo.Annotation{},
		entx.GraphKeyDirective("id"),
		schema.Comment("Represents an ip block node on the graph."),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Create a new ip block type node."),
			entgql.MutationUpdate().Description("Update an existing ip block type node."),
		),
	}
}
