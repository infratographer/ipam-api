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

func TestQuery_IPAddress(t *testing.T) {
	client := graphTestClient()
	ctx := context.Background()

	ipa1 := (&IPAddressBuilder{}).MustNew(ctx)

	testCases := []struct {
		name     string
		queryID  gidx.PrefixedID
		expected *generated.IPAddress
		errorMsg string
	}{
		{
			name:     "happy path ip address",
			queryID:  ipa1.ID,
			expected: ipa1,
		},
		{
			name:     "invalid-id",
			queryID:  gidx.MustNewID("testing"),
			errorMsg: "ip_address not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			resp, err := client.GetIPAddress(ctx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.IPAddress)
		})
	}
}

func Test_IPAddress_Lifecycle(t *testing.T) {
	client := graphTestClient()
	ctx := context.Background()

	ipbt := (&IPBlockTypeBuilder{}).MustNew(ctx)
	ipb := (&IPBlockBuilder{IPBlockTypeID: ipbt.ID}).MustNew(ctx)

	t.Run("Create", func(t *testing.T) {
		ipa, err := client.CreateIPAddress(ctx, testclient.CreateIPAddressInput{
			IP:          gofakeit.IPv4Address(),
			NodeID:      gidx.MustNewID(nodePrefix),
			NodeOwnerID: gidx.MustNewID(ownerPrefix),
			Reserved:    newBool(true),
			IPBlockID:   ipb.ID,
		})

		require.NoError(t, err)
		require.NotNil(t, ipa)

		ipaUpdate, err := client.UpdateIPAddress(ctx, ipa.CreateIPAddress.IPAddress.ID, testclient.UpdateIPAddressInput{
			Reserved: newBool(false),
		})

		require.NoError(t, err)
		require.NotNil(t, ipaUpdate)
		require.Equal(t, false, ipaUpdate.UpdateIPAddress.IPAddress.Reserved)

		ipaDelete, err := client.DeleteIPAddress(ctx, ipa.CreateIPAddress.IPAddress.ID)

		require.NoError(t, err)
		require.NotNil(t, ipaDelete)
	})
}
