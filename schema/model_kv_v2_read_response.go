// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// KvV2ReadResponse struct for KvV2ReadResponse
type KvV2ReadResponse struct {
	Data map[string]interface{} `json:"data"`

	Metadata map[string]interface{} `json:"metadata"`
}

// NewKvV2ReadResponseWithDefaults instantiates a new KvV2ReadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvV2ReadResponseWithDefaults() *KvV2ReadResponse {
	var this KvV2ReadResponse

	return &this
}

func (o KvV2ReadResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["data"] = o.Data
	toSerialize["metadata"] = o.Metadata

	return json.Marshal(toSerialize)
}
