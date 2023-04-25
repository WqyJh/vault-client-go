// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// LeasesListResponse struct for LeasesListResponse
type LeasesListResponse struct {
	// Number of matching leases per mount
	Counts int32 `json:"counts"`

	// Number of matching leases
	LeaseCount int32 `json:"lease_count"`
}

// NewLeasesListResponseWithDefaults instantiates a new LeasesListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeasesListResponseWithDefaults() *LeasesListResponse {
	var this LeasesListResponse

	return &this
}

func (o LeasesListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["counts"] = o.Counts
	toSerialize["lease_count"] = o.LeaseCount

	return json.Marshal(toSerialize)
}
