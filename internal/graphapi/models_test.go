package graphapi_test

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
	"go.infratographer.com/x/gidx"

	ent "go.infratographer.com/ipam-api/internal/ent/generated"
)

type IpBlockTypeBuilder struct {
	Name    string
	OwnerID gidx.PrefixedID
}

func (i *IpBlockTypeBuilder) MustNew(ctx context.Context) *ent.IPBlockType {
	ipbtCreate := EntClient.IPBlockType.Create()

	if i.Name == "" {
		i.Name = gofakeit.Name()
	}

	ipbtCreate.SetName(i.Name)

	if i.OwnerID == "" {
		i.OwnerID = gidx.MustNewID(ownerPrefix)
	}

	ipbtCreate.SetOwnerID(i.OwnerID)

	return ipbtCreate.SaveX(ctx)
}

type IpBlockBuilder struct {
	Prefix            string
	LocationID        gidx.PrefixedID
	ParentBlockID     gidx.PrefixedID
	AllowAutoSubnet   *bool
	AllowAutoAllocate *bool
	IPBlockTypeID     gidx.PrefixedID
}

func (i *IpBlockBuilder) MustNew(ctx context.Context) *ent.IPBlock {
	ipbCreate := EntClient.IPBlock.Create()

	if i.Prefix == "" {
		i.Prefix = gofakeit.IPv4Address()
	}

	ipbCreate.SetPrefix(i.Prefix)

	if i.LocationID == "" {
		i.LocationID = gidx.MustNewID(locationPrefix)
	}

	ipbCreate.SetLocationID(i.LocationID)

	if i.ParentBlockID != "" {
		i.ParentBlockID = gidx.MustNewID(ipBlockTypePrefix)
	}

	ipbCreate.SetParentBlockID(i.ParentBlockID)

	if i.AllowAutoSubnet != nil {
		ipbCreate.SetAllowAutoSubnet(*i.AllowAutoSubnet)
	}

	if i.AllowAutoAllocate != nil {
		ipbCreate.SetAllowAutoAllocate(*i.AllowAutoAllocate)
	}

	if i.IPBlockTypeID == "" {
		ipbt := (&IpBlockTypeBuilder{OwnerID: gidx.MustNewID(ownerPrefix)}).MustNew(ctx)
		i.IPBlockTypeID = ipbt.ID
	}

	ipbCreate.SetIPBlockTypeID(i.IPBlockTypeID)

	return ipbCreate.SaveX(ctx)
}
