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

// OIDCWriteProviderRequest struct for OIDCWriteProviderRequest
type OIDCWriteProviderRequest struct {
	// The client IDs that are permitted to use the provider
	AllowedClientIds []string `json:"allowed_client_ids"`
	// Specifies what will be used for the iss claim of ID tokens.
	Issuer string `json:"issuer"`
	// The scopes supported for requesting on the provider
	ScopesSupported []string `json:"scopes_supported"`
}

// NewOIDCWriteProviderRequestWithDefaults instantiates a new OIDCWriteProviderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOIDCWriteProviderRequestWithDefaults() *OIDCWriteProviderRequest {
	var this OIDCWriteProviderRequest

	return &this
}

func (o OIDCWriteProviderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["allowed_client_ids"] = o.AllowedClientIds
	toSerialize["issuer"] = o.Issuer
	toSerialize["scopes_supported"] = o.ScopesSupported

	return json.Marshal(toSerialize)
}
