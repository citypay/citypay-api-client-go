# FindPaymentIntentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalRef** | Pointer to **string** | An external reference to lookup. | [optional] 
**ExternalRefSource** | Pointer to **string** | An external reference source to lookup. | [optional] 
**Merchantid** | Pointer to **int32** | The merchant id the payment intent is registered for. | [optional] 
**PaymentIntentId** | Pointer to **string** | The payment intent id, if known. | [optional] 

## Methods

### NewFindPaymentIntentRequest

`func NewFindPaymentIntentRequest() *FindPaymentIntentRequest`

NewFindPaymentIntentRequest instantiates a new FindPaymentIntentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindPaymentIntentRequestWithDefaults

`func NewFindPaymentIntentRequestWithDefaults() *FindPaymentIntentRequest`

NewFindPaymentIntentRequestWithDefaults instantiates a new FindPaymentIntentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalRef

`func (o *FindPaymentIntentRequest) GetExternalRef() string`

GetExternalRef returns the ExternalRef field if non-nil, zero value otherwise.

### GetExternalRefOk

`func (o *FindPaymentIntentRequest) GetExternalRefOk() (*string, bool)`

GetExternalRefOk returns a tuple with the ExternalRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRef

`func (o *FindPaymentIntentRequest) SetExternalRef(v string)`

SetExternalRef sets ExternalRef field to given value.

### HasExternalRef

`func (o *FindPaymentIntentRequest) HasExternalRef() bool`

HasExternalRef returns a boolean if a field has been set.

### GetExternalRefSource

`func (o *FindPaymentIntentRequest) GetExternalRefSource() string`

GetExternalRefSource returns the ExternalRefSource field if non-nil, zero value otherwise.

### GetExternalRefSourceOk

`func (o *FindPaymentIntentRequest) GetExternalRefSourceOk() (*string, bool)`

GetExternalRefSourceOk returns a tuple with the ExternalRefSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefSource

`func (o *FindPaymentIntentRequest) SetExternalRefSource(v string)`

SetExternalRefSource sets ExternalRefSource field to given value.

### HasExternalRefSource

`func (o *FindPaymentIntentRequest) HasExternalRefSource() bool`

HasExternalRefSource returns a boolean if a field has been set.

### GetMerchantid

`func (o *FindPaymentIntentRequest) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *FindPaymentIntentRequest) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *FindPaymentIntentRequest) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.

### HasMerchantid

`func (o *FindPaymentIntentRequest) HasMerchantid() bool`

HasMerchantid returns a boolean if a field has been set.

### GetPaymentIntentId

`func (o *FindPaymentIntentRequest) GetPaymentIntentId() string`

GetPaymentIntentId returns the PaymentIntentId field if non-nil, zero value otherwise.

### GetPaymentIntentIdOk

`func (o *FindPaymentIntentRequest) GetPaymentIntentIdOk() (*string, bool)`

GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntentId

`func (o *FindPaymentIntentRequest) SetPaymentIntentId(v string)`

SetPaymentIntentId sets PaymentIntentId field to given value.

### HasPaymentIntentId

`func (o *FindPaymentIntentRequest) HasPaymentIntentId() bool`

HasPaymentIntentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


