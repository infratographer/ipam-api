package ipamclient

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.infratographer.com/ipam-api/pkg/ipamclient/internal/mock"
)

func newGQLClientMock() *mock.GQLClient {
	mockCli := &mock.GQLClient{}
	mockCli.DoMutate = func(ctx context.Context, m interface{}, variables map[string]interface{}) error {
		block, ok := m.(*GetIPBlock)
		if ok {
			block.IPBlock.ID = "ipamibk-12345"
			block.IPBlock.Prefix = "192.168.10.0/28"
			block.IPBlock.AllowAutoAllocate = true
			block.IPBlock.AllowAutoSubnet = true
		}

		address, ok := m.(*GetIPAddress)
		if ok {
			address.IPAddress.ID = "ipamipa-12345"
			address.IPAddress.IP = "192.168.10.1"
			address.IPAddress.Reserved = true
			address.IPAddress.IPBlock.ID = "ipamibk-12345"
			address.IPAddress.IPBlock.Prefix = "192.168.10.0/28"
			address.IPAddress.IPBlock.AllowAutoAllocate = true
			address.IPAddress.IPBlock.AllowAutoSubnet = true
			address.IPAddress.IPBlock.IPBlockType.ID = "ipamibt-12345"
			address.IPAddress.IPBlock.IPBlockType.Name = "test-block"
			address.IPAddress.IPBlock.IPBlockType.Owner.ID = "testtnt-12345"
		}

		return nil
	}

	mockCli.DoQuery = func(ctx context.Context, q interface{}, variables map[string]interface{}) error {
		block, ok := q.(*GetIPBlock)
		if ok {
			block.IPBlock.ID = "ipamibk-12345"
			block.IPBlock.Prefix = "192.168.10.0/28"
			block.IPBlock.AllowAutoAllocate = true
			block.IPBlock.AllowAutoSubnet = true
		}

		address, ok := q.(*GetIPAddress)
		if ok {
			address.IPAddress.ID = "ipamipa-12345"
			address.IPAddress.IP = "192.168.10.1"
			address.IPAddress.Reserved = true
			address.IPAddress.IPBlock.ID = "ipamibk-12345"
			address.IPAddress.IPBlock.Prefix = "192.168.10.0/28"
			address.IPAddress.IPBlock.AllowAutoAllocate = true
			address.IPAddress.IPBlock.AllowAutoSubnet = true
			address.IPAddress.IPBlock.IPBlockType.ID = "ipamibt-12345"
			address.IPAddress.IPBlock.IPBlockType.Name = "test-block"
			address.IPAddress.IPBlock.IPBlockType.Owner.ID = "testtnt-12345"
		}

		return nil
	}

	return mockCli
}

func TestGetIPBlock(t *testing.T) {
	cli := Client{
		gqlCli: newGQLClientMock(),
	}

	block, err := cli.GetIPBlock(context.TODO(), "ipamibk-12345")
	require.NoError(t, err)
	require.NotNil(t, block)

	assert.Equal(t, block.IPBlock.ID, "ipamibk-12345")
	assert.Equal(t, block.IPBlock.Prefix, "192.168.10.0/28")

	badblock, err := cli.GetIPBlock(context.Background(), "badprefix-test")
	require.Error(t, err)
	require.Nil(t, badblock)
	assert.ErrorContains(t, err, "invalid id")
}

func TestGetIPAddress(t *testing.T) {
	cli := Client{
		gqlCli: newGQLClientMock(),
	}

	address, err := cli.GetIPAddress(context.Background(), "ipamipa-12345")
	require.NoError(t, err)
	require.NotNil(t, address)

	assert.Equal(t, address.IPAddress.ID, "ipamipa-12345")
	assert.Equal(t, address.IPAddress.IP, "192.168.10.1")

	baddress, err := cli.GetIPAddress(context.TODO(), "badprefix-test")
	require.Error(t, err)
	require.Nil(t, baddress)
}

func TestGetNextAvailableAddressFromBlock(t *testing.T) {
	cli := Client{
		gqlCli: newGQLClientMock(),
	}

	ip, err := cli.GetNextAvailableAddressFromBlock(context.Background(), "ipamibk-12345")
	require.NoError(t, err)
	require.NotNil(t, ip)

	assert.Equal(t, ip, "192.168.10.1")

	// TODO: test for exhausted block
	// TODO: test for invalid block prefix

	ip, err = cli.GetNextAvailableAddressFromBlock(context.Background(), "badprefix-test")
	require.Error(t, err)
	require.Equal(t, ip, "")
}

func TestDeleteIPAddress(t *testing.T) {
	cli := Client{
		gqlCli: newGQLClientMock(),
	}

	// TODO: build better mock to build more expressive tests

	ip, err := cli.DeleteIPAddress(context.Background(), "ipamipa-12345")
	require.NoError(t, err)
	require.NotNil(t, ip)

	ip, err = cli.DeleteIPAddress(context.Background(), "badprefix-test")
	require.Error(t, err)
	require.Nil(t, ip)
}

func TestCreateIPAddressFromBlock(t *testing.T) {
	cli := Client{
		gqlCli: newGQLClientMock(),
	}

	ip, err := cli.CreateIPAddressFromBlock(context.Background(), "ipamibk-12345", "loadbal-12345", "testtnt-123456", true)
	require.NoError(t, err)
	require.NotNil(t, ip)

	// TODO: need better mock to build more expressive tests

	ip, err = cli.CreateIPAddressFromBlock(context.Background(), "badprefix-test", "loadbal-12345", "testtnt-123456", true)
	require.Error(t, err)
	require.Nil(t, ip)
}
