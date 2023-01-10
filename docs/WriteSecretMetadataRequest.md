# WriteSecretMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CasRequired** | Pointer to **bool** | If true the key will require the cas parameter to be set on all write requests. If false, the backend’s configuration will be used. | [optional] 
**CustomMetadata** | Pointer to **map[string]interface{}** | User-provided key-value pairs that are used to describe arbitrary and version-agnostic information about a secret. | [optional] 
**DeleteVersionAfter** | Pointer to **int32** | The length of time before a version is deleted. If not set, the backend&#39;s configured delete_version_after is used. Cannot be greater than the backend&#39;s delete_version_after. A zero duration clears the current setting. A negative duration will cause an error. | [optional] 
**MaxVersions** | Pointer to **int32** | The number of versions to keep. If not set, the backend’s configured max version is used. | [optional] 

## Methods

### NewWriteSecretMetadataRequest

`func NewWriteSecretMetadataRequest() *WriteSecretMetadataRequest`

NewWriteSecretMetadataRequest instantiates a new WriteSecretMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWriteSecretMetadataRequestWithDefaults

`func NewWriteSecretMetadataRequestWithDefaults() *WriteSecretMetadataRequest`

NewWriteSecretMetadataRequestWithDefaults instantiates a new WriteSecretMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCasRequired

`func (o *WriteSecretMetadataRequest) GetCasRequired() bool`

GetCasRequired returns the CasRequired field if non-nil, zero value otherwise.

### GetCasRequiredOk

`func (o *WriteSecretMetadataRequest) GetCasRequiredOk() (*bool, bool)`

GetCasRequiredOk returns a tuple with the CasRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasRequired

`func (o *WriteSecretMetadataRequest) SetCasRequired(v bool)`

SetCasRequired sets CasRequired field to given value.

### HasCasRequired

`func (o *WriteSecretMetadataRequest) HasCasRequired() bool`

HasCasRequired returns a boolean if a field has been set.

### GetCustomMetadata

`func (o *WriteSecretMetadataRequest) GetCustomMetadata() map[string]interface{}`

GetCustomMetadata returns the CustomMetadata field if non-nil, zero value otherwise.

### GetCustomMetadataOk

`func (o *WriteSecretMetadataRequest) GetCustomMetadataOk() (*map[string]interface{}, bool)`

GetCustomMetadataOk returns a tuple with the CustomMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetadata

`func (o *WriteSecretMetadataRequest) SetCustomMetadata(v map[string]interface{})`

SetCustomMetadata sets CustomMetadata field to given value.

### HasCustomMetadata

`func (o *WriteSecretMetadataRequest) HasCustomMetadata() bool`

HasCustomMetadata returns a boolean if a field has been set.

### GetDeleteVersionAfter

`func (o *WriteSecretMetadataRequest) GetDeleteVersionAfter() int32`

GetDeleteVersionAfter returns the DeleteVersionAfter field if non-nil, zero value otherwise.

### GetDeleteVersionAfterOk

`func (o *WriteSecretMetadataRequest) GetDeleteVersionAfterOk() (*int32, bool)`

GetDeleteVersionAfterOk returns a tuple with the DeleteVersionAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteVersionAfter

`func (o *WriteSecretMetadataRequest) SetDeleteVersionAfter(v int32)`

SetDeleteVersionAfter sets DeleteVersionAfter field to given value.

### HasDeleteVersionAfter

`func (o *WriteSecretMetadataRequest) HasDeleteVersionAfter() bool`

HasDeleteVersionAfter returns a boolean if a field has been set.

### GetMaxVersions

`func (o *WriteSecretMetadataRequest) GetMaxVersions() int32`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *WriteSecretMetadataRequest) GetMaxVersionsOk() (*int32, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *WriteSecretMetadataRequest) SetMaxVersions(v int32)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *WriteSecretMetadataRequest) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

