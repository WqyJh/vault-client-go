/*
HashiCorp Vault API

HTTP API that gives you full access to Vault. All API routes are prefixed with `/v1/`.

API version: 1.13.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// RevokePrefixRequest struct for RevokePrefixRequest
type RevokePrefixRequest struct {
	// Whether or not to perform the revocation synchronously
	Sync bool `json:"sync"`
}

// NewRevokePrefixRequestWithDefaults instantiates a new RevokePrefixRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokePrefixRequestWithDefaults() *RevokePrefixRequest {
	var this RevokePrefixRequest

	this.Sync = true

	return &this
}

func (o RevokePrefixRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["sync"] = o.Sync

	return json.Marshal(toSerialize)
}
