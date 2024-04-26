# CResAuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cres** | Pointer to **string** | The challenge response data forwarded by the ACS in 3D-Secure V2 processing. Data should be forwarded to CityPay unchanged for subsequent authorisation and processing.  | [optional] 

## Methods

### NewCResAuthRequest

`func NewCResAuthRequest() *CResAuthRequest`

NewCResAuthRequest instantiates a new CResAuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCResAuthRequestWithDefaults

`func NewCResAuthRequestWithDefaults() *CResAuthRequest`

NewCResAuthRequestWithDefaults instantiates a new CResAuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCres

`func (o *CResAuthRequest) GetCres() string`

GetCres returns the Cres field if non-nil, zero value otherwise.

### GetCresOk

`func (o *CResAuthRequest) GetCresOk() (*string, bool)`

GetCresOk returns a tuple with the Cres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCres

`func (o *CResAuthRequest) SetCres(v string)`

SetCres sets Cres field to given value.

### HasCres

`func (o *CResAuthRequest) HasCres() bool`

HasCres returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


