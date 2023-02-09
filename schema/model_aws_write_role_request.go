// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
//
// Code generated with OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schema

import (
	"encoding/json"
)

// AWSWriteRoleRequest struct for AWSWriteRoleRequest
type AWSWriteRoleRequest struct {
	// Use role_arns or policy_arns instead.
	// Deprecated
	Arn string `json:"arn"`

	// Type of credential to retrieve. Must be one of assumed_role, iam_user, or federation_token
	CredentialType string `json:"credential_type"`

	// Default TTL for assumed_role and federation_token credential types when no TTL is explicitly requested with the credentials
	DefaultStsTtl int32 `json:"default_sts_ttl"`

	// Names of IAM groups that generated IAM users will be added to. For a credential type of assumed_role or federation_token, the policies sent to the corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the policies from each group in iam_groups combined with the policy_document and policy_arns parameters.
	IamGroups []string `json:"iam_groups"`

	// IAM tags to be set for any users created by this role. These must be presented as Key-Value pairs. This can be represented as a map or a list of equal sign delimited key pairs.
	IamTags map[string]interface{} `json:"iam_tags"`

	// Max allowed TTL for assumed_role and federation_token credential types
	MaxStsTtl int32 `json:"max_sts_ttl"`

	// ARN of an IAM policy to attach as a permissions boundary on IAM user credentials; only valid when credential_type isiam_user
	PermissionsBoundaryArn string `json:"permissions_boundary_arn"`

	// Use policy_document instead.
	// Deprecated
	Policy string `json:"policy"`

	// ARNs of AWS policies. Behavior varies by credential_type. When credential_type is iam_user, then it will attach the specified policies to the generated IAM user. When credential_type is assumed_role or federation_token, the policies will be passed as the PolicyArns parameter, acting as a filter on permissions available.
	PolicyArns []string `json:"policy_arns"`

	// JSON-encoded IAM policy document. Behavior varies by credential_type. When credential_type is iam_user, then it will attach the contents of the policy_document to the IAM user generated. When credential_type is assumed_role or federation_token, this will be passed in as the Policy parameter to the AssumeRole or GetFederationToken API call, acting as a filter on permissions available.
	PolicyDocument string `json:"policy_document"`

	// ARNs of AWS roles allowed to be assumed. Only valid when credential_type is assumed_role
	RoleArns []string `json:"role_arns"`

	// Path for IAM User. Only valid when credential_type is iam_user
	UserPath string `json:"user_path"`
}

// NewAWSWriteRoleRequestWithDefaults instantiates a new AWSWriteRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSWriteRoleRequestWithDefaults() *AWSWriteRoleRequest {
	var this AWSWriteRoleRequest

	this.UserPath = "/"

	return &this
}

func (o AWSWriteRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(o)
}