# BatchTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The card holder account id to process against. | 
**Amount** | **int32** | The amount required to process in the lowest denomination. | 
**Identifier** | Pointer to **string** | An identifier used to reference the transaction set by your integration. The value should be used to refer to the transaction in future calls. | [optional] 
**Merchantid** | Pointer to **int32** | The CityPay merchant id used to process the transaction. | [optional] 

## Methods

### NewBatchTransaction

`func NewBatchTransaction(accountId string, amount int32, ) *BatchTransaction`

NewBatchTransaction instantiates a new BatchTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchTransactionWithDefaults

`func NewBatchTransactionWithDefaults() *BatchTransaction`

NewBatchTransactionWithDefaults instantiates a new BatchTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *BatchTransaction) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BatchTransaction) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BatchTransaction) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAmount

`func (o *BatchTransaction) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BatchTransaction) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BatchTransaction) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetIdentifier

`func (o *BatchTransaction) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *BatchTransaction) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *BatchTransaction) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *BatchTransaction) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetMerchantid

`func (o *BatchTransaction) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *BatchTransaction) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *BatchTransaction) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.

### HasMerchantid

`func (o *BatchTransaction) HasMerchantid() bool`

HasMerchantid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


