package graphapi_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/x/gidx"

	"go.infratographer.com/ipam-api/internal/ent/generated"
	"go.infratographer.com/ipam-api/internal/testclient"
)

func TestQuery_IPAddress(t *testing.T) {
	client := graphTestClient()
	ctx := context.Background()

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

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

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	ipbt := (&IPBlockTypeBuilder{}).MustNew(ctx)
	ipb := (&IPBlockBuilder{IPBlockTypeID: ipbt.ID}).MustNew(ctx)
	node := gidx.MustNewID(nodePrefix)

	t.Run("Create", func(t *testing.T) {
		ipa, err := client.CreateIPAddress(ctx, testclient.CreateIPAddressInput{
			IP:          gofakeit.IPv4Address(),
			NodeID:      node,
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

		addr, err := client.GetIPAddressByNode(ctx, node)

		require.NoError(t, err)
		require.NotNil(t, addr)
		require.Equal(t, addr.IPAddressByNode.ID, ipa.CreateIPAddress.IPAddress.ID)

		ipaDelete, err := client.DeleteIPAddress(ctx, ipa.CreateIPAddress.IPAddress.ID)

		require.NoError(t, err)
		require.NotNil(t, ipaDelete)
	})
}

func Test_IPAddressable(t *testing.T) {
	client := graphTestClient()
	ctx := context.Background()

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	ipas := []gidx.PrefixedID{}

	ipbt := (&IPBlockTypeBuilder{}).MustNew(ctx)
	ipb := (&IPBlockBuilder{IPBlockTypeID: ipbt.ID}).MustNew(ctx)

	node := gidx.MustNewID(nodePrefix)

	// Create first IP address attached to node
	ipa, err := client.CreateIPAddress(ctx, testclient.CreateIPAddressInput{
		IP:          gofakeit.IPv4Address(),
		NodeID:      node,
		NodeOwnerID: gidx.MustNewID(ownerPrefix),
		Reserved:    newBool(true),
		IPBlockID:   ipb.ID,
	})

	require.NoError(t, err)
	require.NotNil(t, ipa)

	ipas = append(ipas, ipa.CreateIPAddress.IPAddress.ID)

	// Check that only 1 IP address is attached to node
	addrs, err := client.GetIPAddressesByNode(ctx, node)

	require.NoError(t, err)
	assert.Len(t, addrs.Entities[0].IPAddresses, 1)

	// Create second IP address attached to node
	ipa, err = client.CreateIPAddress(ctx, testclient.CreateIPAddressInput{
		IP:          gofakeit.IPv4Address(),
		NodeID:      node,
		NodeOwnerID: gidx.MustNewID(ownerPrefix),
		Reserved:    newBool(true),
		IPBlockID:   ipb.ID,
	})

	require.NoError(t, err)
	assert.NotNil(t, ipa)

	// Check that 2 IP addresses are attached to node
	addrs, err = client.GetIPAddressesByNode(ctx, node)

	require.NoError(t, err)
	assert.Len(t, addrs.Entities[0].IPAddresses, 2)

	ipas = append(ipas, ipa.CreateIPAddress.IPAddress.ID)

	// Delete IP addresses attached to node
	for _, ipa := range ipas {
		_, err := client.DeleteIPAddress(ctx, ipa)
		require.NoError(t, err)
	}

	// Check that no IP addresses are attached to node
	addrs, err = client.GetIPAddressesByNode(ctx, node)

	require.NoError(t, err)
	assert.Len(t, addrs.Entities[0].IPAddresses, 0)
}
