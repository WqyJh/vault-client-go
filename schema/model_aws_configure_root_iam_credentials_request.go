// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AwsConfigureRootIamCredentialsRequest struct for AwsConfigureRootIamCredentialsRequest
type AwsConfigureRootIamCredentialsRequest struct {
	// Access key with permission to create new keys.
	AccessKey string `json:"access_key"`

	// Endpoint to custom IAM server URL
	IamEndpoint string `json:"iam_endpoint"`

	// Maximum number of retries for recoverable exceptions of AWS APIs
	MaxRetries int32 `json:"max_retries"`

	// Region for API calls.
	Region string `json:"region"`

	// Secret key with permission to create new keys.
	SecretKey string `json:"secret_key"`

	// Endpoint to custom STS server URL
	StsEndpoint string `json:"sts_endpoint"`

	// Template to generate custom IAM usernames
	UsernameTemplate string `json:"username_template"`
}

// NewAwsConfigureRootIamCredentialsRequestWithDefaults instantiates a new AwsConfigureRootIamCredentialsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsConfigureRootIamCredentialsRequestWithDefaults() *AwsConfigureRootIamCredentialsRequest {
	var this AwsConfigureRootIamCredentialsRequest

	this.MaxRetries = -1

	return &this
}

func (o AwsConfigureRootIamCredentialsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["access_key"] = o.AccessKey
	toSerialize["iam_endpoint"] = o.IamEndpoint
	toSerialize["max_retries"] = o.MaxRetries
	toSerialize["region"] = o.Region
	toSerialize["secret_key"] = o.SecretKey
	toSerialize["sts_endpoint"] = o.StsEndpoint
	toSerialize["username_template"] = o.UsernameTemplate

	return json.Marshal(toSerialize)
}