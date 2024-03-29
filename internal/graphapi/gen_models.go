// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphapi

import (
	"go.infratographer.com/ipam-api/internal/ent/generated"
	"go.infratographer.com/x/gidx"
)

// Return response for createIPAddress mutation
type IPAddressCreatePayload struct {
	// Created ip address
	IPAddress *generated.IPAddress `json:"ipAddress"`
}

// Return response for deleteIPAddress mutation
type IPAddressDeletePayload struct {
	// Deleted ip address ID
	DeletedID gidx.PrefixedID `json:"deletedID"`
}

// Return response for updateIPAddress mutation
type IPAddressUpdatePayload struct {
	// Updated ip address
	IPAddress *generated.IPAddress `json:"ipAddress"`
}

// IPAddressable provides an interface for describing IP addresses attached to a node
type IPAddressable struct {
	ID gidx.PrefixedID `json:"id"`
	// IPAddressable describes IP addresses attached to a node
	IPAddresses []*generated.IPAddress `json:"IPAddresses"`
}

func (IPAddressable) IsEntity() {}

// Return response for createIPBlock mutation
type IPBlockCreatePayload struct {
	// Created ip block
	IPBlock *generated.IPBlock `json:"ipBlock"`
}

// Return response for deleteIPBlock mutation
type IPBlockDeletePayload struct {
	// Deleted ip block
	DeletedID gidx.PrefixedID `json:"deletedID"`
}

// Return response for createIPBlockType mutation
type IPBlockTypeCreatePayload struct {
	// Created ip block type
	IPBlockType *generated.IPBlockType `json:"ipBlockType"`
}

// Return response for deleteIPBlockType mutation
type IPBlockTypeDeletePayload struct {
	// Deleted ip block type
	DeletedID gidx.PrefixedID `json:"deletedID"`
}

// Return response for updateIPBlockType mutation
type IPBlockTypeUpdatePayload struct {
	// Updated ip block type
	IPBlockType *generated.IPBlockType `json:"ipBlockType"`
}

// Return response for updateIPBlock mutation
type IPBlockUpdatePayload struct {
	// Updated ip block
	IPBlock *generated.IPBlock `json:"ipBlock"`
}

type ResourceOwner struct {
	ID          gidx.PrefixedID                  `json:"id"`
	IPBlockType *generated.IPBlockTypeConnection `json:"ipBlockType"`
}

func (ResourceOwner) IsEntity() {}
