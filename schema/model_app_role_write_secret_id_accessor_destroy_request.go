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

// AppRoleWriteSecretIDAccessorDestroyRequest struct for AppRoleWriteSecretIDAccessorDestroyRequest
type AppRoleWriteSecretIDAccessorDestroyRequest struct {
	// Accessor of the SecretID
	SecretIdAccessor string `json:"secret_id_accessor"`
}

// NewAppRoleWriteSecretIDAccessorDestroyRequestWithDefaults instantiates a new AppRoleWriteSecretIDAccessorDestroyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleWriteSecretIDAccessorDestroyRequestWithDefaults() *AppRoleWriteSecretIDAccessorDestroyRequest {
	var this AppRoleWriteSecretIDAccessorDestroyRequest

	return &this
}

func (o AppRoleWriteSecretIDAccessorDestroyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["secret_id_accessor"] = o.SecretIdAccessor

	return json.Marshal(toSerialize)
}
