package ipamclient

// GetIPAddressResult is the result of a query for an IP Address
type GetIPAddressResult struct {
	IPAddressNode
	IPBlock IPBlock `graphql:"ipBlock" json:"ipBlock"`
}

// GetIPBlockResult is the result of a query for an IP Block
type GetIPBlockResult struct {
	IPBlock

	IPAddress IPAddresses `graphql:"ipAddress" json:"ipAddress"`
}

// IPBlock is part of a returned query including an associated block
type IPBlock struct {
	ID                string      `graphql:"id" json:"id"`
	Prefix            string      `graphql:"prefix" json:"prefix"`
	AllowAutoSubnet   bool        `graphql:"allowAutoSubnet" json:"allowAutoSubnet"`
	AllowAutoAllocate bool        `graphql:"allowAutoAllocate" json:"allowAutoAllocate"`
	IPBlockType       IPBlockType `graphql:"ipBlockType" json:"ipBlockType"`
}

// IPAddresses is part of a returned query including the edges of an associated IP Address
type IPAddresses struct {
	Edges []IPAddressEdge `graphql:"edges" json:"edges"`
}

// IPAddressEdge is part of a returned query including an edge(s) associated to an IP Address
type IPAddressEdge struct {
	Node IPAddressNode `graphql:"node" json:"node"`
}

// IPAddressNode is part of a returned query including an associated IP address
type IPAddressNode struct {
	ID       string `graphql:"id" json:"id"`
	IP       string `graphql:"ip" json:"ip"`
	Reserved bool   `graphql:"reserved" json:"reserved"`
}

// IPBlockType is part of a returned query including an associated block type
type IPBlockType struct {
	ID    string `graphql:"id" json:"id"`
	Name  string `graphql:"name" json:"name"`
	Owner Owner  `graphql:"owner" json:"owner"`
}

// Owner is part of a returned query including an associated owner
type Owner struct {
	ID string `graphql:"id" json:"id"`
}

// GetIPAddress is the query used for getting an IP Address
type GetIPAddress struct {
	IPAddress GetIPAddressResult `graphql:"ipAddress(id: $id)"`
}

// GetIPBlock is the query used for getting an IP Block
type GetIPBlock struct {
	IPBlock GetIPBlockResult `graphql:"ipBlock(id: $id)"`
}

// CreateIPAddress is the mutation for creating an IP Address
type CreateIPAddress struct {
	CreateIPAddressResult `graphql:"createIPAddress(input: $input)"`
}

// CreateIPAddressResult is the result of the mutation to create an IP Address
type CreateIPAddressResult struct {
	// IPAddress MyIPAddress `graphql:"ipAddress"`
	IPAddress struct {
		IPAddressNode
		IPBlock IPBlock
	} `graphql:"ipAddress"`
}

// CreateIPAddressInput is the set of input required for creating an IP address
type CreateIPAddressInput struct {
	IP          string `graphql:"ip" json:"ip"`
	IPBlockID   string `graphql:"ipBlockID" json:"ipBlockID"`
	NodeID      string `graphql:"nodeID" json:"nodeID"`
	NodeOwnerID string `graphql:"nodeOwnerID" json:"nodeOwnerID"`
	Reserved    bool   `graphql:"reserved" json:"reserved"`
}

// DeleteIPAddress is the mutation for deleting an IP Address
type DeleteIPAddress struct {
	DeleteIPAddress DeleteIPAddressResult `graphql:"deleteIPAddress(id: $id)"`
}

// DeleteIPAddressResult is the result of the mutation to delete an IP Address
type DeleteIPAddressResult struct {
	DeletedID string `graphql:"deletedID"`
}

// IPAddressableFragment fragment for getting IP Addresses by node
type IPAddressableFragment struct {
	IPAddresses []IPAddressNode `graphql:"IPAddresses"`
}

// GetIPAddressesByNode query for getting IP Addresses by node
type GetIPAddressesByNode struct {
	NodeIPAddress struct {
		ID string `graphql:"id" json:"id"`

		LoadBalancerFragment struct {
			IPAddressableFragment
		} `graphql:"... on LoadBalancer"`
	} `graphql:"node(id: $id)"`
}
