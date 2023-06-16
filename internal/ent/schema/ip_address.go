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

// IPAddress holds the schema definition for the IPAddress entity.
type IPAddress struct {
	ent.Schema
}

// Mixin of the IP Block Type type
func (IPAddress) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

// Fields of the IPAddress.
func (IPAddress) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the IP Address.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(IPAddressPrefix) }),
		field.Text("IP").
			NotEmpty().
			Comment("The ip address.").
			Annotations(
				entgql.OrderField("IP"),
			),
		field.String("block_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the ip block for this ip address.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("BLOCK"),
				pubsubinfo.AdditionalSubject(),
			),
		field.String("node_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("The ID for the node this is assigned to.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("NODE"),
				pubsubinfo.AdditionalSubject(),
			),
		field.String("node_owner_id").
			GoType(gidx.PrefixedID("")).
			Immutable().
			Comment("Owner ID of the node this is assigned to.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput, entgql.SkipMutationUpdateInput, entgql.SkipType),
				entgql.OrderField("OWNER"),
				pubsubinfo.AdditionalSubject(),
			),
		field.Bool("reserved").
			Default(true).
			Comment("Reserve the IP without it being assigned.").
			Annotations(
				entgql.OrderField("RESERVED"),
			),
	}
}

// Edges of the IPAddress.
func (IPAddress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ip_block", IPBlock.Type).
			Unique().
			Required().
			Immutable().
			Field("block_id").
			Annotations(),
	}
}

// Indexes of the IPAddress.
func (IPAddress) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("node_owner_id", "block_id", "node_id"),
	}
}

// Annotations of the IPAddress
func (IPAddress) Annotations() []schema.Annotation {
	return []schema.Annotation{
		pubsubinfo.Annotation{},
		entx.GraphKeyDirective("id"),
		schema.Comment("Represents an ip address node on the graph."),
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Create a new ip address type node."),
			entgql.MutationUpdate().Description("Update an existing ip address type node."),
		),
	}
}
