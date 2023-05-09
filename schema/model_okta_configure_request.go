// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// OktaConfigureRequest struct for OktaConfigureRequest
type OktaConfigureRequest struct {
	// Okta API key.
	ApiToken string `json:"api_token"`

	// The base domain to use for the Okta API. When not specified in the configuration, \"okta.com\" is used.
	BaseUrl string `json:"base_url"`

	// When set true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.
	BypassOktaMfa bool `json:"bypass_okta_mfa"`

	// Use \"token_max_ttl\" instead. If this and \"token_max_ttl\" are both specified, only \"token_max_ttl\" will be used.
	// Deprecated
	MaxTtl int32 `json:"max_ttl"`

	// Name of the organization to be used in the Okta API.
	OrgName string `json:"org_name"`

	// Use org_name instead.
	// Deprecated
	Organization string `json:"organization"`

	// Use base_url instead.
	// Deprecated
	Production bool `json:"production"`

	// Use api_token instead.
	// Deprecated
	Token string `json:"token"`

	// Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token.
	TokenBoundCidrs []string `json:"token_bound_cidrs"`

	// If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed.
	TokenExplicitMaxTtl int32 `json:"token_explicit_max_ttl"`

	// The maximum lifetime of the generated token
	TokenMaxTtl int32 `json:"token_max_ttl"`

	// If true, the 'default' policy will not automatically be added to generated tokens
	TokenNoDefaultPolicy bool `json:"token_no_default_policy"`

	// The maximum number of times a token may be used, a value of zero means unlimited
	TokenNumUses int32 `json:"token_num_uses"`

	// If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \"24h\").
	TokenPeriod int32 `json:"token_period"`

	// Comma-separated list of policies. This will apply to all tokens generated by this auth method, in addition to any configured for specific users/groups.
	TokenPolicies []string `json:"token_policies"`

	// The initial ttl of the token to generate
	TokenTtl int32 `json:"token_ttl"`

	// The type of token to generate, service or batch
	TokenType string `json:"token_type"`

	// Use \"token_ttl\" instead. If this and \"token_ttl\" are both specified, only \"token_ttl\" will be used.
	// Deprecated
	Ttl int32 `json:"ttl"`
}

// NewOktaConfigureRequestWithDefaults instantiates a new OktaConfigureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaConfigureRequestWithDefaults() *OktaConfigureRequest {
	var this OktaConfigureRequest

	this.TokenType = "default-service"

	return &this
}

func (o OktaConfigureRequest) MarshalJSON() ([]byte, error) {
	toSerialize := make(map[string]interface{})

	toSerialize["api_token"] = o.ApiToken
	toSerialize["base_url"] = o.BaseUrl
	toSerialize["bypass_okta_mfa"] = o.BypassOktaMfa
	toSerialize["max_ttl"] = o.MaxTtl
	toSerialize["org_name"] = o.OrgName
	toSerialize["organization"] = o.Organization
	toSerialize["production"] = o.Production
	toSerialize["token"] = o.Token
	toSerialize["token_bound_cidrs"] = o.TokenBoundCidrs
	toSerialize["token_explicit_max_ttl"] = o.TokenExplicitMaxTtl
	toSerialize["token_max_ttl"] = o.TokenMaxTtl
	toSerialize["token_no_default_policy"] = o.TokenNoDefaultPolicy
	toSerialize["token_num_uses"] = o.TokenNumUses
	toSerialize["token_period"] = o.TokenPeriod
	toSerialize["token_policies"] = o.TokenPolicies
	toSerialize["token_ttl"] = o.TokenTtl
	toSerialize["token_type"] = o.TokenType
	toSerialize["ttl"] = o.Ttl

	return json.Marshal(toSerialize)
}