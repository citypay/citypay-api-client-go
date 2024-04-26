# PaylinkTokenRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accountno** | Pointer to **string** | Specifies an alpha-numeric account number that the Paylink service uses when creating a Cardholder Account. The value should be no longer than 20 characters in length. | [optional] 
**Amount** | **int32** | Specifies the intended value of the transaction in the lowest denomination with no spacing characters or decimal point. This is the net total to be processed. An example of £74.95 would be presented as 7495. | 
**Cardholder** | Pointer to [**PaylinkCardHolder**](PaylinkCardHolder.md) |  | [optional] 
**Cart** | Pointer to [**PaylinkCart**](PaylinkCart.md) |  | [optional] 
**ClientVersion** | Pointer to **string** | The clientVersion field is used to specify the version of your application that has invoked the Paylink payment process. This feature is typically used for tracing issues relating to application deployments, or any Paylink integration module or plugin. | [optional] 
**Config** | Pointer to [**PaylinkConfig**](PaylinkConfig.md) |  | [optional] 
**Currency** | Pointer to **string** | A currency for the token. This value should be only used on multi-currency accounts and be an appropriate currency which the account is configured for. | [optional] 
**Email** | Pointer to **string** | The email field is used for the Merchant to be notified on completion of the transaction . The value may be supplied to override the default stored value. Emails sent to this address by the Paylink service should not be forwarded on to the cardholder as it may contain certain information that is used by the Paylink service to validate and authenticate Paylink Token Requests: for example, the Merchant ID and the licence key.  | [optional] 
**Identifier** | **string** | Identifies a particular transaction linked to a Merchant account. It enables accurate duplicate checking within a pre-configured time period, as well as transaction reporting and tracing. The identifier should be unique to prevent payment card processing attempts from being rejected due to duplication.  | 
**Merchantid** | **int32** | The merchant id you wish to process this transaction with. | 
**Recurring** | Pointer to **bool** | True if the intent of this cardholder initiated transaction is to establish a recurring payment model, processable as merchant initiated transactions. | [optional] 
**SubscriptionId** | Pointer to **string** | an id associated with a subscription to link the token request against. | [optional] 
**TxType** | Pointer to **string** | A value to override the transaction type if requested by your account manager. | [optional] 

## Methods

### NewPaylinkTokenRequestModel

`func NewPaylinkTokenRequestModel(amount int32, identifier string, merchantid int32, ) *PaylinkTokenRequestModel`

NewPaylinkTokenRequestModel instantiates a new PaylinkTokenRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkTokenRequestModelWithDefaults

`func NewPaylinkTokenRequestModelWithDefaults() *PaylinkTokenRequestModel`

NewPaylinkTokenRequestModelWithDefaults instantiates a new PaylinkTokenRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountno

`func (o *PaylinkTokenRequestModel) GetAccountno() string`

GetAccountno returns the Accountno field if non-nil, zero value otherwise.

### GetAccountnoOk

`func (o *PaylinkTokenRequestModel) GetAccountnoOk() (*string, bool)`

GetAccountnoOk returns a tuple with the Accountno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountno

`func (o *PaylinkTokenRequestModel) SetAccountno(v string)`

SetAccountno sets Accountno field to given value.

### HasAccountno

`func (o *PaylinkTokenRequestModel) HasAccountno() bool`

HasAccountno returns a boolean if a field has been set.

### GetAmount

`func (o *PaylinkTokenRequestModel) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaylinkTokenRequestModel) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaylinkTokenRequestModel) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCardholder

`func (o *PaylinkTokenRequestModel) GetCardholder() PaylinkCardHolder`

GetCardholder returns the Cardholder field if non-nil, zero value otherwise.

### GetCardholderOk

`func (o *PaylinkTokenRequestModel) GetCardholderOk() (*PaylinkCardHolder, bool)`

GetCardholderOk returns a tuple with the Cardholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholder

`func (o *PaylinkTokenRequestModel) SetCardholder(v PaylinkCardHolder)`

SetCardholder sets Cardholder field to given value.

### HasCardholder

`func (o *PaylinkTokenRequestModel) HasCardholder() bool`

HasCardholder returns a boolean if a field has been set.

### GetCart

`func (o *PaylinkTokenRequestModel) GetCart() PaylinkCart`

GetCart returns the Cart field if non-nil, zero value otherwise.

### GetCartOk

`func (o *PaylinkTokenRequestModel) GetCartOk() (*PaylinkCart, bool)`

GetCartOk returns a tuple with the Cart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCart

`func (o *PaylinkTokenRequestModel) SetCart(v PaylinkCart)`

SetCart sets Cart field to given value.

### HasCart

`func (o *PaylinkTokenRequestModel) HasCart() bool`

HasCart returns a boolean if a field has been set.

### GetClientVersion

`func (o *PaylinkTokenRequestModel) GetClientVersion() string`

GetClientVersion returns the ClientVersion field if non-nil, zero value otherwise.

### GetClientVersionOk

`func (o *PaylinkTokenRequestModel) GetClientVersionOk() (*string, bool)`

GetClientVersionOk returns a tuple with the ClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVersion

`func (o *PaylinkTokenRequestModel) SetClientVersion(v string)`

SetClientVersion sets ClientVersion field to given value.

### HasClientVersion

`func (o *PaylinkTokenRequestModel) HasClientVersion() bool`

HasClientVersion returns a boolean if a field has been set.

### GetConfig

`func (o *PaylinkTokenRequestModel) GetConfig() PaylinkConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PaylinkTokenRequestModel) GetConfigOk() (*PaylinkConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PaylinkTokenRequestModel) SetConfig(v PaylinkConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *PaylinkTokenRequestModel) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCurrency

`func (o *PaylinkTokenRequestModel) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaylinkTokenRequestModel) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaylinkTokenRequestModel) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaylinkTokenRequestModel) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEmail

`func (o *PaylinkTokenRequestModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PaylinkTokenRequestModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PaylinkTokenRequestModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PaylinkTokenRequestModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIdentifier

`func (o *PaylinkTokenRequestModel) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PaylinkTokenRequestModel) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PaylinkTokenRequestModel) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetMerchantid

`func (o *PaylinkTokenRequestModel) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *PaylinkTokenRequestModel) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *PaylinkTokenRequestModel) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.


### GetRecurring

`func (o *PaylinkTokenRequestModel) GetRecurring() bool`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *PaylinkTokenRequestModel) GetRecurringOk() (*bool, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *PaylinkTokenRequestModel) SetRecurring(v bool)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *PaylinkTokenRequestModel) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *PaylinkTokenRequestModel) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PaylinkTokenRequestModel) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PaylinkTokenRequestModel) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *PaylinkTokenRequestModel) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTxType

`func (o *PaylinkTokenRequestModel) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *PaylinkTokenRequestModel) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *PaylinkTokenRequestModel) SetTxType(v string)`

SetTxType sets TxType field to given value.

### HasTxType

`func (o *PaylinkTokenRequestModel) HasTxType() bool`

HasTxType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


