// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AliasWriteByIDRequest struct for AliasWriteByIDRequest
type AliasWriteByIDRequest struct {
	// Entity ID to which this alias should be tied to
	CanonicalId string `json:"canonical_id"`

	// Entity ID to which this alias should be tied to. This field is deprecated in favor of 'canonical_id'.
	EntityId string `json:"entity_id"`

	// Mount accessor to which this alias belongs to
	MountAccessor string `json:"mount_accessor"`

	// Name of the alias
	Name string `json:"name"`
}

// NewAliasWriteByIDRequestWithDefaults instantiates a new AliasWriteByIDRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAliasWriteByIDRequestWithDefaults() *AliasWriteByIDRequest {
	var this AliasWriteByIDRequest

	return &this
}

func (o AliasWriteByIDRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}