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

func TestQuery(t *testing.T) {
	client := graphTestClient()
	ctx := context.Background()

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	ownerID := gidx.MustNewID("testown")

	ipbt1 := (&IPBlockTypeBuilder{
		OwnerID: ownerID,
	}).MustNew(ctx)

	testCases := []struct {
		name     string
		queryID  gidx.PrefixedID
		expected *generated.IPBlockType
		errorMsg string
	}{
		{
			name:     "happy path ip block type",
			queryID:  ipbt1.ID,
			expected: ipbt1,
		},
		{
			name:     "invalid-id",
			queryID:  gidx.MustNewID("testing"),
			errorMsg: "ip_block_type not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			resp, err := client.GetIPBlockType(ctx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.IPBlockType)
		})
	}
}

func Test_HappyPath(t *testing.T) {
	client := graphTestClient()
	ctx := context.Background()
	ownerID := gidx.MustNewID("testown")

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	t.Run("Create + List + Update + Delete", func(t *testing.T) {
		ipbt, err := client.IPBlockTypeCreate(ctx, testclient.CreateIPBlockTypeInput{
			Name:    gofakeit.Dessert(),
			OwnerID: ownerID,
		})
		require.NoError(t, err)
		require.NotNil(t, ipbt)

		list, err := client.ListIPBlockTypes(ctx, ownerID, nil)
		require.NoError(t, err)
		require.NotNil(t, list)
		assert.Len(t, list.Entities[0].IPBlockType.Edges, 1)

		ipbtUpdate, err := client.IPBlockTypeUpdate(ctx, ipbt.CreateIPBlockType.IPBlockType.ID, testclient.UpdateIPBlockTypeInput{
			Name: newString(gofakeit.Dessert()),
		})

		require.NoError(t, err)
		require.NotNil(t, ipbtUpdate)

		assert.NotEqual(t, ipbt.CreateIPBlockType.IPBlockType.Name, ipbtUpdate.UpdateIPBlockType.IPBlockType.Name)

		deleteID, err := client.IPBlockTypeDelete(ctx, ipbt.CreateIPBlockType.IPBlockType.ID)
		require.NoError(t, err)
		require.NotNil(t, deleteID)
		assert.EqualValues(t, ipbt.CreateIPBlockType.IPBlockType.ID, deleteID.DeleteIPBlockType.DeletedID)
	})
}
