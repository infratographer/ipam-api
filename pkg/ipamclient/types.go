package ipamclient

// GetIPAddressResult is the result of a query for an IP Address
type GetIPAddressResult struct {
	ID       string
	IP       string
	Reserved bool
	IPBlock  IPBlock
}

// GetIPBlockResult is the result of a query for an IP Block
type GetIPBlockResult struct {
	ID                string
	Prefix            string
	AllowAutoSubnet   bool
	AllowAutoAllocate bool
	IPBlockType       IPBlockType
	IPAddress         IPAddresses
}

// IPBlock is part of a returned query including an associated block
type IPBlock struct {
	ID                string
	Prefix            string
	AllowAutoSubnet   bool
	AllowAutoAllocate bool
	IPBlockType       IPBlockType
}

// IPAddresses is part of a returned query including the edges of an associated IP Address
type IPAddresses struct {
	Edges []IPAddressEdge
}

// IPAddressEdge is part of a returned query including an edge(s) associated to an IP Address
type IPAddressEdge struct {
	Node IPAddressNode
}

// IPAddressNode is part of a returned query including an associated IP address
type IPAddressNode struct {
	ID       string
	IP       string
	Reserved bool
}

// IPBlockType is part of a returned query including an associated block type
type IPBlockType struct {
	ID    string
	Name  string
	Owner Owner
}

// Owner is part of a returned query including an associated owner
type Owner struct {
	ID string
}

// GetIPAddress is the query used for getting an IP Address
type GetIPAddress struct {
	IPAddress GetIPAddressResult `graphql:"ip_address(id: $id)"`
}

// GetIPBlock is the query used for getting an IP Block
type GetIPBlock struct {
	IPBlock GetIPBlockResult `graphql:"ip_block(id: $id)"`
}

// CreateIPAddress is the mutation for creating an IP Address
type CreateIPAddress struct {
	IPAddress CreateIPAddressResult `graphql:"createIPAddress(input: $input)"`
}

// CreateIPAddressResult is the result of the mutation to create an IP Address
type CreateIPAddressResult struct {
	// IPAddress StupidIPAddress `graphql:"ip_address"`
	IPAddress struct {
		ID       string
		IP       string
		Reserved bool
		IPBlock  IPBlock
	} `graphql:"ip_address"`
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
