# AclCheckResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | Pointer to **string** | The name or value of the acl which was found to match the ip address. | [optional] 
**Cache** | Pointer to **bool** | Whether the ACL was returned via a cached instance. | [optional] 
**Ip** | Pointer to **string** | The IP address used in the lookup. | [optional] 
**Provider** | Pointer to **string** | The source provider of the ACL such as cloud, subnet, country or IP based. | [optional] 

## Methods

### NewAclCheckResponseModel

`func NewAclCheckResponseModel() *AclCheckResponseModel`

NewAclCheckResponseModel instantiates a new AclCheckResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAclCheckResponseModelWithDefaults

`func NewAclCheckResponseModelWithDefaults() *AclCheckResponseModel`

NewAclCheckResponseModelWithDefaults instantiates a new AclCheckResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcl

`func (o *AclCheckResponseModel) GetAcl() string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *AclCheckResponseModel) GetAclOk() (*string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *AclCheckResponseModel) SetAcl(v string)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *AclCheckResponseModel) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetCache

`func (o *AclCheckResponseModel) GetCache() bool`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *AclCheckResponseModel) GetCacheOk() (*bool, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *AclCheckResponseModel) SetCache(v bool)`

SetCache sets Cache field to given value.

### HasCache

`func (o *AclCheckResponseModel) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetIp

`func (o *AclCheckResponseModel) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AclCheckResponseModel) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AclCheckResponseModel) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AclCheckResponseModel) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetProvider

`func (o *AclCheckResponseModel) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AclCheckResponseModel) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AclCheckResponseModel) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AclCheckResponseModel) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


