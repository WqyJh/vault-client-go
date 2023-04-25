// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KvV2ReadConfigurationResponse struct for KvV2ReadConfigurationResponse
type KvV2ReadConfigurationResponse struct {
	// If true, the backend will require the cas parameter to be set for each write
	CasRequired bool `json:"cas_required"`

	// The length of time before a version is deleted.
	DeleteVersionAfter int32 `json:"delete_version_after"`

	// The number of versions to keep for each key.
	MaxVersions int32 `json:"max_versions"`
}

// NewKvV2ReadConfigurationResponseWithDefaults instantiates a new KvV2ReadConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvV2ReadConfigurationResponseWithDefaults() *KvV2ReadConfigurationResponse {
	var this KvV2ReadConfigurationResponse

	return &this
}

func (o KvV2ReadConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["cas_required"] = o.CasRequired
	toSerialize["delete_version_after"] = o.DeleteVersionAfter
	toSerialize["max_versions"] = o.MaxVersions

	return json.Marshal(toSerialize)
}
