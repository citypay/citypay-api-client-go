# AclCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | An ip address to check for an ACL against. The address should be a publicly routable IPv4 address. | 

## Methods

### NewAclCheckRequest

`func NewAclCheckRequest(ip string, ) *AclCheckRequest`

NewAclCheckRequest instantiates a new AclCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAclCheckRequestWithDefaults

`func NewAclCheckRequestWithDefaults() *AclCheckRequest`

NewAclCheckRequestWithDefaults instantiates a new AclCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *AclCheckRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AclCheckRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AclCheckRequest) SetIp(v string)`

SetIp sets Ip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


