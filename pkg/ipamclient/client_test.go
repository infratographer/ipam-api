package ipamclient

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hasura/go-graphql-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.infratographer.com/ipam-api/pkg/ipamclient/internal/mock"
)

func newGQLClientMock() *mock.GQLClient {
	mockCli := &mock.GQLClient{}
	mockCli.DoMutate = func(ctx context.Context, m interface{}, variables map[string]interface{}, options ...graphql.Option) error {
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

	mockCli.DoQuery = func(ctx context.Context, q interface{}, variables map[string]interface{}, options ...graphql.Option) error {
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
		gqlCli: mustNewGQLTestClient(`{
  "data": {
    "ipBlock": {
      "allowAutoAllocate": true,
      "allowAutoSubnet": true,
      "id": "ipamibk-12345",
      "prefix": "192.168.10.0/28"
    }
  }
}`),
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
		gqlCli: mustNewGQLTestClient(
			`{
  "data": {
    "ipAddress": {
      "id": "ipamipa-12345",
      "ip": "192.168.10.1",
	  "reserved": false
    }
  }
}`),
	}

	ipResult, err := cli.GetIPAddress(context.Background(), "ipamipa-12345")
	require.NoError(t, err)
	require.NotNil(t, ipResult)

	assert.Equal(t, ipResult.IPAddress.ID, "ipamipa-12345")
	assert.Equal(t, ipResult.IPAddress.IP, "192.168.10.1")
	assert.False(t, ipResult.IPAddress.Reserved)

	baddress, err := cli.GetIPAddress(context.TODO(), "badprefix-test")
	require.Error(t, err)
	require.Nil(t, baddress)
}

func TestGetNextAvailableAddressFromBlock(t *testing.T) {
	cli := Client{
		gqlCli: mustNewGQLTestClient(`{
  "data": {
    "ipBlock": {
      "allowAutoAllocate": true,
      "allowAutoSubnet": true,
      "id": "ipamibk-12345",
      "prefix": "192.168.10.0/28"
    }
  }
}`),
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
		gqlCli: mustNewGQLTestClient(`{
  "data": {
    "deleteIPAddress": {
      "deletedID": "ipamipa-12345"
    }
  }
}`),
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

func TestGetIPAddressesByNodeID(t *testing.T) {
	cli := Client{
		gqlCli: mustNewGQLTestClient(
			`{
  "data": {
    "node": {
	  "id": "loadbal-randovalue",
	  "IPAddresses": [
		{
	  	  "id": "ipamipa-8IPzP37YJ1iTxJdMrCods",
		  "ip": "192.168.1.142",
		  "reserved": false
		},
		{
		  "id": "ipamipa-rPBY83fPw6Ll5sueCMpDr",
		  "ip": "192.168.1.1",
		  "reserved": true
		}
	  ]
	}
  }
}`),
	}

	t.Run("invalid prefix", func(t *testing.T) {
		ips, err := cli.GetIPAddresses(context.Background(), "badprefix-test")
		require.Error(t, err)
		require.Nil(t, ips)
	})

	t.Run("retrieves nodeID ip addresses", func(t *testing.T) {
		ips, err := cli.GetIPAddresses(context.Background(), "loadbal-randovalue")
		require.NoError(t, err)
		require.Len(t, ips, 2)

		assert.Equal(t, "ipamipa-8IPzP37YJ1iTxJdMrCods", ips[0].ID)
		assert.Equal(t, "192.168.1.142", ips[0].IP)
		assert.False(t, ips[0].Reserved)

		assert.Equal(t, "ipamipa-rPBY83fPw6Ll5sueCMpDr", ips[1].ID)
		assert.Equal(t, "192.168.1.1", ips[1].IP)
		assert.True(t, ips[1].Reserved)
	})
}

func TestGetIPAddressByNode(t *testing.T) {
	cli := Client{
		gqlCli: mustNewGQLTestClient(
			`{
	  "data": {
		"ipAddressByNode": {
		  "id": "ipamipa-12345",
		  "ip": "192.168.10.1",
		  "reserved": false
		}
	  }
	}`),
	}

	ipResult, err := cli.GetIPAddressByNode(context.Background(), "nodeids-test")
	require.NoError(t, err)
	require.NotNil(t, ipResult)

	assert.Equal(t, ipResult.IPAddress.ID, "ipamipa-12345")
	assert.Equal(t, ipResult.IPAddress.IP, "192.168.10.1")
	assert.False(t, ipResult.IPAddress.Reserved)

	baddress, err := cli.GetIPAddressByNode(context.TODO(), "badprefix-test")
	require.Error(t, err)
	require.Nil(t, baddress)
}

func mustNewGQLTestClient(respJSON string) *graphql.Client {
	mux := http.NewServeMux()
	mux.HandleFunc("/query", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := io.WriteString(w, respJSON)
		if err != nil {
			panic(err)
		}
	})

	return graphql.NewClient("/query", &http.Client{Transport: localRoundTripper{handler: mux}})
}

type localRoundTripper struct {
	handler http.Handler
}

func (l localRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	l.handler.ServeHTTP(w, req)

	return w.Result(), nil
}
