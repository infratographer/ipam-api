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
