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

// ActiveDirectoryWriteRoleRequest struct for ActiveDirectoryWriteRoleRequest
type ActiveDirectoryWriteRoleRequest struct {
	// The username/logon name for the service account with which this role will be associated.
	ServiceAccountName string `json:"service_account_name"`
	// In seconds, the default password time-to-live.
	Ttl int32 `json:"ttl"`
}

// NewActiveDirectoryWriteRoleRequestWithDefaults instantiates a new ActiveDirectoryWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveDirectoryWriteRoleRequestWithDefaults() *ActiveDirectoryWriteRoleRequest {
	var this ActiveDirectoryWriteRoleRequest

	return &this
}

func (o ActiveDirectoryWriteRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["service_account_name"] = o.ServiceAccountName
	toSerialize["ttl"] = o.Ttl

	return json.Marshal(toSerialize)
}
