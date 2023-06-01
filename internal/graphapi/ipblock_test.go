package graphapi_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.infratographer.com/x/gidx"

	"go.infratographer.com/ipam-api/internal/ent/generated"
	"go.infratographer.com/ipam-api/internal/testclient"
)

func TestQuery_IPBlock(t *testing.T) {
	client := graphTestClient()
	ctx := context.Background()

	ipb1 := (&IpBlockBuilder{}).MustNew(ctx)

	testCases := []struct {
		name     string
		queryID  gidx.PrefixedID
		expected *generated.IPBlock
		errorMsg string
	}{
		{
			name:     "happy path ip block",
			queryID:  ipb1.ID,
			expected: ipb1,
		},
		{
			name:     "invalid-id",
			queryID:  gidx.MustNewID("testing"),
			errorMsg: "ip_block not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			resp, err := client.GetIPBlock(ctx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.IPBlock)
		})
	}
}

func Test_IPBlock_Lifecycle(t *testing.T) {
	client := graphTestClient()
	ctx := context.Background()

	// Create a new IPBlockType to use for the IPBlock
	ipbt := (&IpBlockTypeBuilder{OwnerID: gidx.MustNewID(ownerPrefix)}).MustNew(ctx)

	t.Run("Create Update Delete", func(t *testing.T) {
		ipb, err := client.IPBlockCreate(ctx, testclient.CreateIPBlockInput{
			Prefix:            gofakeit.IPv4Address(),
			LocationID:        gidx.MustNewID(locationPrefix),
			IPBlockTypeID:     ipbt.ID,
			ParentBlockID:     gidx.MustNewID(ipBlockTypePrefix),
			AllowAutoSubnet:   newBool(false),
			AllowAutoAllocate: newBool(true),
		})

		require.NoError(t, err)
		require.NotNil(t, ipb)

		ibpUpdate, err := client.IPBlockUpdate(ctx, ipb.CreateIPBlock.IPBlock.ID, testclient.UpdateIPBlockInput{
			AllowAutoSubnet:   newBool(true),
			AllowAutoAllocate: newBool(false),
		})

		require.NoError(t, err)
		require.NotNil(t, ibpUpdate)
		require.Equal(t, true, ibpUpdate.UpdateIPBlock.IPBlock.AllowAutoSubnet)
		require.Equal(t, false, ibpUpdate.UpdateIPBlock.IPBlock.AllowAutoAllocate)

		ibpDelete, err := client.IPBlockDelete(ctx, ipb.CreateIPBlock.IPBlock.ID)

		require.NoError(t, err)
		require.NotNil(t, ibpDelete)
	})
}
