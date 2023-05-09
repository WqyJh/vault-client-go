// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// SshIssueCertificateRequest struct for SshIssueCertificateRequest
type SshIssueCertificateRequest struct {
	// Type of certificate to be created; either \"user\" or \"host\".
	CertType string `json:"cert_type"`

	// Critical options that the certificate should be signed for.
	CriticalOptions map[string]interface{} `json:"critical_options"`

	// Extensions that the certificate should be signed for.
	Extensions map[string]interface{} `json:"extensions"`

	// Specifies the number of bits to use for the generated keys.
	KeyBits int32 `json:"key_bits"`

	// Key id that the created certificate should have. If not specified, the display name of the token will be used.
	KeyId string `json:"key_id"`

	// Specifies the desired key type; must be `rsa`, `ed25519` or `ec`
	KeyType string `json:"key_type"`

	// The requested Time To Live for the SSH certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be later than the role max TTL.
	Ttl int32 `json:"ttl"`

	// Valid principals, either usernames or hostnames, that the certificate should be signed for.
	ValidPrincipals string `json:"valid_principals"`
}

// NewSshIssueCertificateRequestWithDefaults instantiates a new SshIssueCertificateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshIssueCertificateRequestWithDefaults() *SshIssueCertificateRequest {
	var this SshIssueCertificateRequest

	this.CertType = "user"
	this.KeyBits = 0
	this.KeyType = "rsa"

	return &this
}

func (o SshIssueCertificateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["cert_type"] = o.CertType
	toSerialize["critical_options"] = o.CriticalOptions
	toSerialize["extensions"] = o.Extensions
	toSerialize["key_bits"] = o.KeyBits
	toSerialize["key_id"] = o.KeyId
	toSerialize["key_type"] = o.KeyType
	toSerialize["ttl"] = o.Ttl
	toSerialize["valid_principals"] = o.ValidPrincipals

	return json.Marshal(toSerialize)
}