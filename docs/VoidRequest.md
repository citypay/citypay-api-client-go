# VoidRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The identifier of the transaction to void. If an empty value is supplied then a &#x60;trans_no&#x60; value must be supplied. | [optional] 
**Merchantid** | **int32** | Identifies the merchant account to perform the void for. | 
**Transno** | Pointer to **int32** | The transaction number of the transaction to look up and void. If an empty value is supplied then an identifier value must be supplied. | [optional] 

## Methods

### NewVoidRequest

`func NewVoidRequest(merchantid int32, ) *VoidRequest`

NewVoidRequest instantiates a new VoidRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoidRequestWithDefaults

`func NewVoidRequestWithDefaults() *VoidRequest`

NewVoidRequestWithDefaults instantiates a new VoidRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *VoidRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *VoidRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *VoidRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *VoidRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetMerchantid

`func (o *VoidRequest) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *VoidRequest) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *VoidRequest) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.


### GetTransno

`func (o *VoidRequest) GetTransno() int32`

GetTransno returns the Transno field if non-nil, zero value otherwise.

### GetTransnoOk

`func (o *VoidRequest) GetTransnoOk() (*int32, bool)`

GetTransnoOk returns a tuple with the Transno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransno

`func (o *VoidRequest) SetTransno(v int32)`

SetTransno sets Transno field to given value.

### HasTransno

`func (o *VoidRequest) HasTransno() bool`

HasTransno returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


