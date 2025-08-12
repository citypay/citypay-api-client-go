# PaymentIntentResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adjustments** | Pointer to [**Adjustments**](Adjustments.md) |  | [optional] 
**Amount** | Pointer to **int32** | An amount of the intent. | [optional] 
**Created** | Pointer to **time.Time** | A date time of when the intent was created. | [optional] 
**Currency** | Pointer to **string** | The currency of the intent. | [optional] 
**Due** | Pointer to **string** | A due date of the intent. | [optional] 
**Expires** | Pointer to **string** | An expiration date of the intent. | [optional] 
**ExternalRef** | Pointer to **string** | An external reference of the intent. | [optional] 
**ExternalRefSource** | Pointer to **string** | An external reference source of the intent. | [optional] 
**Identifier** | **string** | An identifier of the intent. | 
**IntentStatus** | Pointer to **string** | A status of the intent such as &#x60;unknown&#x60;, &#x60;open&#x60;, &#x60;requires_payment_method&#x60;, &#x60;requires_confirmation&#x60;, &#x60;requires_confirmation&#x60;, &#x60;requires_action&#x60;, &#x60;processing&#x60;, &#x60;succeeded&#x60;, &#x60;cancelled&#x60;, &#x60;requires_capture&#x60;, &#x60;failed&#x60;, &#x60;expired&#x60;, &#x60;requires_refund&#x60;, &#x60;refunded&#x60;. | [optional] 
**Merchantid** | **int32** | The merchant id of the intent. | 
**PaymentType** | Pointer to **string** | A type of the intent such as &#x60;None&#x60;, &#x60;Single&#x60;, &#x60;Subscription&#x60;. | [optional] 
**PaymentIntentId** | **string** | The id of the intent. | 
**Transactions** | Pointer to [**AuthReference**](AuthReference.md) |  | [optional] 

## Methods

### NewPaymentIntentResponseModel

`func NewPaymentIntentResponseModel(identifier string, merchantid int32, paymentIntentId string, ) *PaymentIntentResponseModel`

NewPaymentIntentResponseModel instantiates a new PaymentIntentResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentIntentResponseModelWithDefaults

`func NewPaymentIntentResponseModelWithDefaults() *PaymentIntentResponseModel`

NewPaymentIntentResponseModelWithDefaults instantiates a new PaymentIntentResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjustments

`func (o *PaymentIntentResponseModel) GetAdjustments() Adjustments`

GetAdjustments returns the Adjustments field if non-nil, zero value otherwise.

### GetAdjustmentsOk

`func (o *PaymentIntentResponseModel) GetAdjustmentsOk() (*Adjustments, bool)`

GetAdjustmentsOk returns a tuple with the Adjustments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustments

`func (o *PaymentIntentResponseModel) SetAdjustments(v Adjustments)`

SetAdjustments sets Adjustments field to given value.

### HasAdjustments

`func (o *PaymentIntentResponseModel) HasAdjustments() bool`

HasAdjustments returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentIntentResponseModel) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentIntentResponseModel) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentIntentResponseModel) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentIntentResponseModel) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreated

`func (o *PaymentIntentResponseModel) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PaymentIntentResponseModel) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PaymentIntentResponseModel) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PaymentIntentResponseModel) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentIntentResponseModel) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentIntentResponseModel) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentIntentResponseModel) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentIntentResponseModel) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDue

`func (o *PaymentIntentResponseModel) GetDue() string`

GetDue returns the Due field if non-nil, zero value otherwise.

### GetDueOk

`func (o *PaymentIntentResponseModel) GetDueOk() (*string, bool)`

GetDueOk returns a tuple with the Due field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDue

`func (o *PaymentIntentResponseModel) SetDue(v string)`

SetDue sets Due field to given value.

### HasDue

`func (o *PaymentIntentResponseModel) HasDue() bool`

HasDue returns a boolean if a field has been set.

### GetExpires

`func (o *PaymentIntentResponseModel) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *PaymentIntentResponseModel) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *PaymentIntentResponseModel) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *PaymentIntentResponseModel) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetExternalRef

`func (o *PaymentIntentResponseModel) GetExternalRef() string`

GetExternalRef returns the ExternalRef field if non-nil, zero value otherwise.

### GetExternalRefOk

`func (o *PaymentIntentResponseModel) GetExternalRefOk() (*string, bool)`

GetExternalRefOk returns a tuple with the ExternalRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRef

`func (o *PaymentIntentResponseModel) SetExternalRef(v string)`

SetExternalRef sets ExternalRef field to given value.

### HasExternalRef

`func (o *PaymentIntentResponseModel) HasExternalRef() bool`

HasExternalRef returns a boolean if a field has been set.

### GetExternalRefSource

`func (o *PaymentIntentResponseModel) GetExternalRefSource() string`

GetExternalRefSource returns the ExternalRefSource field if non-nil, zero value otherwise.

### GetExternalRefSourceOk

`func (o *PaymentIntentResponseModel) GetExternalRefSourceOk() (*string, bool)`

GetExternalRefSourceOk returns a tuple with the ExternalRefSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefSource

`func (o *PaymentIntentResponseModel) SetExternalRefSource(v string)`

SetExternalRefSource sets ExternalRefSource field to given value.

### HasExternalRefSource

`func (o *PaymentIntentResponseModel) HasExternalRefSource() bool`

HasExternalRefSource returns a boolean if a field has been set.

### GetIdentifier

`func (o *PaymentIntentResponseModel) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PaymentIntentResponseModel) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PaymentIntentResponseModel) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetIntentStatus

`func (o *PaymentIntentResponseModel) GetIntentStatus() string`

GetIntentStatus returns the IntentStatus field if non-nil, zero value otherwise.

### GetIntentStatusOk

`func (o *PaymentIntentResponseModel) GetIntentStatusOk() (*string, bool)`

GetIntentStatusOk returns a tuple with the IntentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentStatus

`func (o *PaymentIntentResponseModel) SetIntentStatus(v string)`

SetIntentStatus sets IntentStatus field to given value.

### HasIntentStatus

`func (o *PaymentIntentResponseModel) HasIntentStatus() bool`

HasIntentStatus returns a boolean if a field has been set.

### GetMerchantid

`func (o *PaymentIntentResponseModel) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *PaymentIntentResponseModel) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *PaymentIntentResponseModel) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.


### GetPaymentType

`func (o *PaymentIntentResponseModel) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *PaymentIntentResponseModel) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *PaymentIntentResponseModel) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *PaymentIntentResponseModel) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetPaymentIntentId

`func (o *PaymentIntentResponseModel) GetPaymentIntentId() string`

GetPaymentIntentId returns the PaymentIntentId field if non-nil, zero value otherwise.

### GetPaymentIntentIdOk

`func (o *PaymentIntentResponseModel) GetPaymentIntentIdOk() (*string, bool)`

GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntentId

`func (o *PaymentIntentResponseModel) SetPaymentIntentId(v string)`

SetPaymentIntentId sets PaymentIntentId field to given value.


### GetTransactions

`func (o *PaymentIntentResponseModel) GetTransactions() AuthReference`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *PaymentIntentResponseModel) GetTransactionsOk() (*AuthReference, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *PaymentIntentResponseModel) SetTransactions(v AuthReference)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *PaymentIntentResponseModel) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


