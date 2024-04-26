# Merchant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The currency of the merchant. | [optional] 
**Merchantid** | Pointer to **int32** | The merchant id which uniquely identifies the merchant account. | [optional] 
**Name** | Pointer to **string** | The name of the merchant. | [optional] 
**Status** | Pointer to **string** | The status of the account. | [optional] 
**StatusLabel** | Pointer to **string** | The status label of the account. | [optional] 

## Methods

### NewMerchant

`func NewMerchant() *Merchant`

NewMerchant instantiates a new Merchant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantWithDefaults

`func NewMerchantWithDefaults() *Merchant`

NewMerchantWithDefaults instantiates a new Merchant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *Merchant) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Merchant) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Merchant) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Merchant) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetMerchantid

`func (o *Merchant) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *Merchant) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *Merchant) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.

### HasMerchantid

`func (o *Merchant) HasMerchantid() bool`

HasMerchantid returns a boolean if a field has been set.

### GetName

`func (o *Merchant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Merchant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Merchant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Merchant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *Merchant) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Merchant) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Merchant) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Merchant) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusLabel

`func (o *Merchant) GetStatusLabel() string`

GetStatusLabel returns the StatusLabel field if non-nil, zero value otherwise.

### GetStatusLabelOk

`func (o *Merchant) GetStatusLabelOk() (*string, bool)`

GetStatusLabelOk returns a tuple with the StatusLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusLabel

`func (o *Merchant) SetStatusLabel(v string)`

SetStatusLabel sets StatusLabel field to given value.

### HasStatusLabel

`func (o *Merchant) HasStatusLabel() bool`

HasStatusLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


