# AuthReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | The amount of the transaction in decimal currency format. | [optional] 
**AmountValue** | Pointer to **int32** | The amount of the transaction in integer/request format. | [optional] 
**Atrn** | Pointer to **string** | A reference number provided by the acquiring services. | [optional] 
**Authcode** | Pointer to **string** | The authorisation code of the transaction returned by the acquirer or card issuer. | [optional] 
**Batchno** | Pointer to **string** | A batch number which the transaction has been end of day batched towards. | [optional] 
**Currency** | Pointer to **string** | The currency of the transaction in ISO 4217 code format. | [optional] 
**Datetime** | Pointer to **time.Time** | The date and time of the transaction. | [optional] 
**Identifier** | Pointer to **string** | The identifier of the transaction used to process the transaction. | [optional] 
**Maskedpan** | Pointer to **string** | A masking of the card number which was used to process the tranasction. | [optional] 
**Merchantid** | Pointer to **int32** | The merchant id of the transaction result. | [optional] 
**Result** | Pointer to **string** | The result of the transaction. | [optional] 
**TransStatus** | Pointer to **string** | The current status of the transaction through it&#39;s lifecycle. | [optional] 
**TransType** | Pointer to **string** | The type of transaction that was processed. | [optional] 
**Transno** | Pointer to **int32** | The transaction number of the transaction. | [optional] 

## Methods

### NewAuthReference

`func NewAuthReference() *AuthReference`

NewAuthReference instantiates a new AuthReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthReferenceWithDefaults

`func NewAuthReferenceWithDefaults() *AuthReference`

NewAuthReferenceWithDefaults instantiates a new AuthReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AuthReference) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AuthReference) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AuthReference) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AuthReference) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountValue

`func (o *AuthReference) GetAmountValue() int32`

GetAmountValue returns the AmountValue field if non-nil, zero value otherwise.

### GetAmountValueOk

`func (o *AuthReference) GetAmountValueOk() (*int32, bool)`

GetAmountValueOk returns a tuple with the AmountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountValue

`func (o *AuthReference) SetAmountValue(v int32)`

SetAmountValue sets AmountValue field to given value.

### HasAmountValue

`func (o *AuthReference) HasAmountValue() bool`

HasAmountValue returns a boolean if a field has been set.

### GetAtrn

`func (o *AuthReference) GetAtrn() string`

GetAtrn returns the Atrn field if non-nil, zero value otherwise.

### GetAtrnOk

`func (o *AuthReference) GetAtrnOk() (*string, bool)`

GetAtrnOk returns a tuple with the Atrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtrn

`func (o *AuthReference) SetAtrn(v string)`

SetAtrn sets Atrn field to given value.

### HasAtrn

`func (o *AuthReference) HasAtrn() bool`

HasAtrn returns a boolean if a field has been set.

### GetAuthcode

`func (o *AuthReference) GetAuthcode() string`

GetAuthcode returns the Authcode field if non-nil, zero value otherwise.

### GetAuthcodeOk

`func (o *AuthReference) GetAuthcodeOk() (*string, bool)`

GetAuthcodeOk returns a tuple with the Authcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthcode

`func (o *AuthReference) SetAuthcode(v string)`

SetAuthcode sets Authcode field to given value.

### HasAuthcode

`func (o *AuthReference) HasAuthcode() bool`

HasAuthcode returns a boolean if a field has been set.

### GetBatchno

`func (o *AuthReference) GetBatchno() string`

GetBatchno returns the Batchno field if non-nil, zero value otherwise.

### GetBatchnoOk

`func (o *AuthReference) GetBatchnoOk() (*string, bool)`

GetBatchnoOk returns a tuple with the Batchno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchno

`func (o *AuthReference) SetBatchno(v string)`

SetBatchno sets Batchno field to given value.

### HasBatchno

`func (o *AuthReference) HasBatchno() bool`

HasBatchno returns a boolean if a field has been set.

### GetCurrency

`func (o *AuthReference) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AuthReference) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AuthReference) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AuthReference) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDatetime

`func (o *AuthReference) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *AuthReference) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *AuthReference) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *AuthReference) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetIdentifier

`func (o *AuthReference) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *AuthReference) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *AuthReference) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *AuthReference) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetMaskedpan

`func (o *AuthReference) GetMaskedpan() string`

GetMaskedpan returns the Maskedpan field if non-nil, zero value otherwise.

### GetMaskedpanOk

`func (o *AuthReference) GetMaskedpanOk() (*string, bool)`

GetMaskedpanOk returns a tuple with the Maskedpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedpan

`func (o *AuthReference) SetMaskedpan(v string)`

SetMaskedpan sets Maskedpan field to given value.

### HasMaskedpan

`func (o *AuthReference) HasMaskedpan() bool`

HasMaskedpan returns a boolean if a field has been set.

### GetMerchantid

`func (o *AuthReference) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *AuthReference) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *AuthReference) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.

### HasMerchantid

`func (o *AuthReference) HasMerchantid() bool`

HasMerchantid returns a boolean if a field has been set.

### GetResult

`func (o *AuthReference) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AuthReference) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AuthReference) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *AuthReference) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetTransStatus

`func (o *AuthReference) GetTransStatus() string`

GetTransStatus returns the TransStatus field if non-nil, zero value otherwise.

### GetTransStatusOk

`func (o *AuthReference) GetTransStatusOk() (*string, bool)`

GetTransStatusOk returns a tuple with the TransStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransStatus

`func (o *AuthReference) SetTransStatus(v string)`

SetTransStatus sets TransStatus field to given value.

### HasTransStatus

`func (o *AuthReference) HasTransStatus() bool`

HasTransStatus returns a boolean if a field has been set.

### GetTransType

`func (o *AuthReference) GetTransType() string`

GetTransType returns the TransType field if non-nil, zero value otherwise.

### GetTransTypeOk

`func (o *AuthReference) GetTransTypeOk() (*string, bool)`

GetTransTypeOk returns a tuple with the TransType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransType

`func (o *AuthReference) SetTransType(v string)`

SetTransType sets TransType field to given value.

### HasTransType

`func (o *AuthReference) HasTransType() bool`

HasTransType returns a boolean if a field has been set.

### GetTransno

`func (o *AuthReference) GetTransno() int32`

GetTransno returns the Transno field if non-nil, zero value otherwise.

### GetTransnoOk

`func (o *AuthReference) GetTransnoOk() (*int32, bool)`

GetTransnoOk returns a tuple with the Transno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransno

`func (o *AuthReference) SetTransno(v int32)`

SetTransno sets Transno field to given value.

### HasTransno

`func (o *AuthReference) HasTransno() bool`

HasTransno returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


