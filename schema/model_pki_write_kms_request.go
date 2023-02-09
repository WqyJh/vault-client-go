// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// PKIWriteKMSRequest struct for PKIWriteKMSRequest
type PKIWriteKMSRequest struct {
	// The number of bits to use. Allowed values are 0 (universal default); with rsa key_type: 2048 (default), 3072, or 4096; with ec key_type: 224, 256 (default), 384, or 521; ignored with ed25519.
	KeyBits int32 `json:"key_bits"`

	// Optional name to be used for this key
	KeyName string `json:"key_name"`

	// The type of key to use; defaults to RSA. \"rsa\" \"ec\" and \"ed25519\" are the only valid values.
	KeyType string `json:"key_type"`

	// The name of the managed key to use when the exported type is kms. When kms type is the key type, this field or managed_key_name is required. Ignored for other types.
	ManagedKeyId string `json:"managed_key_id"`

	// The name of the managed key to use when the exported type is kms. When kms type is the key type, this field or managed_key_id is required. Ignored for other types.
	ManagedKeyName string `json:"managed_key_name"`
}

// NewPKIWriteKMSRequestWithDefaults instantiates a new PKIWriteKMSRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPKIWriteKMSRequestWithDefaults() *PKIWriteKMSRequest {
	var this PKIWriteKMSRequest

	this.KeyBits = 0
	this.KeyType = "rsa"

	return &this
}

func (o PKIWriteKMSRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}