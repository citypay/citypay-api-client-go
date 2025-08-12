# RetrieveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | Pointer to **string** | The identifier of the transaction to retrieve. Optional if a transaction number is provided. | [optional] 
**Merchantid** | **int32** | The merchant account to retrieve data for. | 
**Transno** | Pointer to **int32** | The transaction number of a transaction to retrieve. Optional if an identifier is supplied. | [optional] 

## Methods

### NewRetrieveRequest

`func NewRetrieveRequest(merchantid int32, ) *RetrieveRequest`

NewRetrieveRequest instantiates a new RetrieveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetrieveRequestWithDefaults

`func NewRetrieveRequestWithDefaults() *RetrieveRequest`

NewRetrieveRequestWithDefaults instantiates a new RetrieveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *RetrieveRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *RetrieveRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *RetrieveRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *RetrieveRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetMerchantid

`func (o *RetrieveRequest) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *RetrieveRequest) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *RetrieveRequest) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.


### GetTransno

`func (o *RetrieveRequest) GetTransno() int32`

GetTransno returns the Transno field if non-nil, zero value otherwise.

### GetTransnoOk

`func (o *RetrieveRequest) GetTransnoOk() (*int32, bool)`

GetTransnoOk returns a tuple with the Transno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransno

`func (o *RetrieveRequest) SetTransno(v int32)`

SetTransno sets Transno field to given value.

### HasTransno

`func (o *RetrieveRequest) HasTransno() bool`

HasTransno returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


