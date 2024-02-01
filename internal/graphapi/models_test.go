package graphapi_test

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
	"go.infratographer.com/x/gidx"

	ent "go.infratographer.com/ipam-api/internal/ent/generated"
)

type IPBlockTypeBuilder struct {
	Name    string
	OwnerID gidx.PrefixedID
}

func (i *IPBlockTypeBuilder) MustNew(ctx context.Context) *ent.IPBlockType {
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

type IPBlockBuilder struct {
	Prefix            string
	LocationID        gidx.PrefixedID
	ParentBlockID     gidx.PrefixedID
	AllowAutoSubnet   *bool
	AllowAutoAllocate *bool
	IPBlockTypeID     gidx.PrefixedID
}

func (i *IPBlockBuilder) MustNew(ctx context.Context) *ent.IPBlock {
	ipbCreate := EntClient.IPBlock.Create()

	if i.Prefix == "" {
		i.Prefix = gofakeit.IPv4Address() + "/25"
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
		ipbt := (&IPBlockTypeBuilder{OwnerID: gidx.MustNewID(ownerPrefix)}).MustNew(ctx)
		i.IPBlockTypeID = ipbt.ID
	}

	ipbCreate.SetIPBlockTypeID(i.IPBlockTypeID)

	return ipbCreate.SaveX(ctx)
}

type IPAddressBuilder struct {
	IP          string
	Reserved    bool
	IPBlockID   gidx.PrefixedID
	NodeID      gidx.PrefixedID
	NodeOwnerID gidx.PrefixedID
}

func (i *IPAddressBuilder) MustNew(ctx context.Context) *ent.IPAddress {
	ipaCreate := EntClient.IPAddress.Create()

	if i.IP == "" {
		i.IP = gofakeit.IPv4Address()
	}

	ipaCreate.SetIP(i.IP)

	ipaCreate.SetReserved(i.Reserved)

	if i.IPBlockID == "" {
		ipbt := (&IPBlockTypeBuilder{OwnerID: gidx.MustNewID(ownerPrefix)}).MustNew(ctx)
		ipb := (&IPBlockBuilder{IPBlockTypeID: ipbt.ID}).MustNew(ctx)
		i.IPBlockID = ipb.ID
	}

	ipaCreate.SetIPBlockID(i.IPBlockID)

	if i.NodeID == "" {
		i.NodeID = gidx.MustNewID(nodePrefix)
	}

	ipaCreate.SetNodeID(i.NodeID)

	if i.NodeOwnerID == "" {
		i.NodeOwnerID = gidx.MustNewID(ownerPrefix)
	}

	ipaCreate.SetNodeOwnerID(i.NodeOwnerID)

	return ipaCreate.SaveX(ctx)
}
