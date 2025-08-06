# RegisterIpModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exp** | Pointer to **int64** | When the ip address expires. At most an IP address can be registered for up to 720 hours. Will default to 12 hours if not supplied. | [optional] 
**Ip** | Pointer to **string** | The remote ip address to register. Will default to your current IP. | [optional] 

## Methods

### NewRegisterIpModel

`func NewRegisterIpModel() *RegisterIpModel`

NewRegisterIpModel instantiates a new RegisterIpModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterIpModelWithDefaults

`func NewRegisterIpModelWithDefaults() *RegisterIpModel`

NewRegisterIpModelWithDefaults instantiates a new RegisterIpModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExp

`func (o *RegisterIpModel) GetExp() int64`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *RegisterIpModel) GetExpOk() (*int64, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *RegisterIpModel) SetExp(v int64)`

SetExp sets Exp field to given value.

### HasExp

`func (o *RegisterIpModel) HasExp() bool`

HasExp returns a boolean if a field has been set.

### GetIp

`func (o *RegisterIpModel) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RegisterIpModel) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RegisterIpModel) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RegisterIpModel) HasIp() bool`

HasIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


