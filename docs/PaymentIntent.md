# PaymentIntent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The amount to authorise in the lowest unit of currency with a variable length to a maximum of 12 digits.  No decimal points are to be included and no divisional characters such as 1,024.  The amount should be the total amount required for the transaction.  For example with GBP £1,021.95 the amount value is 102195.  | 
**AvsPostcodePolicy** | Pointer to **string** | A policy value which determines whether an AVS postcode policy is enforced or bypassed.  Values are:   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be rejected if the AVS postcode numeric value does not match.   &#x60;2&#x60; to bypass. Transactions that are bypassed will be allowed through even if the postcode did not match.   &#x60;3&#x60; to ignore. Transactions that are ignored will bypass the result and not send postcode details for authorisation.  | [optional] 
**BillTo** | Pointer to [**ContactDetails**](ContactDetails.md) |  | [optional] 
**Csc** | Pointer to **string** | The Card Security Code (CSC) (also known as CV2/CVV2) is normally found on the back of the card (American Express has it on the front). The value helps to identify possession of the card as it is not available within the chip or magnetic swipe.  When forwarding the CSC, please ensure the value is a string as some values start with 0 and this will be stripped out by any integer parsing.  The CSC number aids fraud prevention in Mail Order and Internet payments.  Business rules are available on your account to identify whether to accept or decline transactions based on mismatched results of the CSC.  The Payment Card Industry (PCI) requires that at no stage of a transaction should the CSC be stored.  This applies to all entities handling card data.  It should also not be used in any hashing process.  CityPay do not store the value and have no method of retrieving the value once the transaction has been processed. For this reason, duplicate checking is unable to determine the CSC in its duplication check algorithm.  | [optional] 
**CscPolicy** | Pointer to **string** | A policy value which determines whether a CSC policy is enforced or bypassed.  Values are:   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be rejected if the CSC value does not match.   &#x60;2&#x60; to bypass. Transactions that are bypassed will be allowed through even if the CSC did not match.   &#x60;3&#x60; to ignore. Transactions that are ignored will bypass the result and not send the CSC details for authorisation.  | [optional] 
**Currency** | Pointer to **string** | The processing currency for the transaction. Will default to the merchant account currency. | [optional] 
**DuplicatePolicy** | Pointer to **string** | A policy value which determines whether a duplication policy is enforced or bypassed. A duplication check has a window of time set against your account within which it can action. If a previous transaction with matching values occurred within the window, any subsequent transaction will result in a T001 result.  Values are   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be checked for duplication within the duplication window.   &#x60;2&#x60; to bypass. Transactions that are bypassed will not be checked for duplication within the duplication window.   &#x60;3&#x60; to ignore. Transactions that are ignored will have the same affect as bypass.  | [optional] 
**Identifier** | **string** | The identifier of the transaction to process. The value should be a valid reference and may be used to perform  post processing actions and to aid in reconciliation of transactions.  The value should be a valid printable string with ASCII character ranges from 0x32 to 0x127.  The identifier is recommended to be distinct for each transaction such as a [random unique identifier](https://en.wikipedia.org/wiki/Universally_unique_identifier) this will aid in ensuring each transaction is identifiable.  When transactions are processed they are also checked for duplicate requests. Changing the identifier on a subsequent request will ensure that a transaction is considered as different.  | 
**MatchAvsa** | Pointer to **string** | A policy value which determines whether an AVS address policy is enforced, bypassed or ignored.  Values are:   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be rejected if the AVS address numeric value does not match.   &#x60;2&#x60; to bypass. Transactions that are bypassed will be allowed through even if the address did not match.   &#x60;3&#x60; to ignore. Transactions that are ignored will bypass the result and not send address numeric details for authorisation.  | [optional] 
**ShipTo** | Pointer to [**ContactDetails**](ContactDetails.md) |  | [optional] 
**Tag** | Pointer to **[]string** |  | [optional] 
**TransInfo** | Pointer to **string** | Further information that can be added to the transaction will display in reporting. Can be used for flexible values such as operator id. | [optional] 
**TransType** | Pointer to **string** | The type of transaction being submitted. Normally this value is not required and your account manager may request that you set this field. | [optional] 

## Methods

### NewPaymentIntent

`func NewPaymentIntent(amount int32, identifier string, ) *PaymentIntent`

NewPaymentIntent instantiates a new PaymentIntent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentIntentWithDefaults

`func NewPaymentIntentWithDefaults() *PaymentIntent`

NewPaymentIntentWithDefaults instantiates a new PaymentIntent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PaymentIntent) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentIntent) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentIntent) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAvsPostcodePolicy

`func (o *PaymentIntent) GetAvsPostcodePolicy() string`

GetAvsPostcodePolicy returns the AvsPostcodePolicy field if non-nil, zero value otherwise.

### GetAvsPostcodePolicyOk

`func (o *PaymentIntent) GetAvsPostcodePolicyOk() (*string, bool)`

GetAvsPostcodePolicyOk returns a tuple with the AvsPostcodePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsPostcodePolicy

`func (o *PaymentIntent) SetAvsPostcodePolicy(v string)`

SetAvsPostcodePolicy sets AvsPostcodePolicy field to given value.

### HasAvsPostcodePolicy

`func (o *PaymentIntent) HasAvsPostcodePolicy() bool`

HasAvsPostcodePolicy returns a boolean if a field has been set.

### GetBillTo

`func (o *PaymentIntent) GetBillTo() ContactDetails`

GetBillTo returns the BillTo field if non-nil, zero value otherwise.

### GetBillToOk

`func (o *PaymentIntent) GetBillToOk() (*ContactDetails, bool)`

GetBillToOk returns a tuple with the BillTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTo

`func (o *PaymentIntent) SetBillTo(v ContactDetails)`

SetBillTo sets BillTo field to given value.

### HasBillTo

`func (o *PaymentIntent) HasBillTo() bool`

HasBillTo returns a boolean if a field has been set.

### GetCsc

`func (o *PaymentIntent) GetCsc() string`

GetCsc returns the Csc field if non-nil, zero value otherwise.

### GetCscOk

`func (o *PaymentIntent) GetCscOk() (*string, bool)`

GetCscOk returns a tuple with the Csc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsc

`func (o *PaymentIntent) SetCsc(v string)`

SetCsc sets Csc field to given value.

### HasCsc

`func (o *PaymentIntent) HasCsc() bool`

HasCsc returns a boolean if a field has been set.

### GetCscPolicy

`func (o *PaymentIntent) GetCscPolicy() string`

GetCscPolicy returns the CscPolicy field if non-nil, zero value otherwise.

### GetCscPolicyOk

`func (o *PaymentIntent) GetCscPolicyOk() (*string, bool)`

GetCscPolicyOk returns a tuple with the CscPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCscPolicy

`func (o *PaymentIntent) SetCscPolicy(v string)`

SetCscPolicy sets CscPolicy field to given value.

### HasCscPolicy

`func (o *PaymentIntent) HasCscPolicy() bool`

HasCscPolicy returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentIntent) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentIntent) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentIntent) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentIntent) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDuplicatePolicy

`func (o *PaymentIntent) GetDuplicatePolicy() string`

GetDuplicatePolicy returns the DuplicatePolicy field if non-nil, zero value otherwise.

### GetDuplicatePolicyOk

`func (o *PaymentIntent) GetDuplicatePolicyOk() (*string, bool)`

GetDuplicatePolicyOk returns a tuple with the DuplicatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicatePolicy

`func (o *PaymentIntent) SetDuplicatePolicy(v string)`

SetDuplicatePolicy sets DuplicatePolicy field to given value.

### HasDuplicatePolicy

`func (o *PaymentIntent) HasDuplicatePolicy() bool`

HasDuplicatePolicy returns a boolean if a field has been set.

### GetIdentifier

`func (o *PaymentIntent) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PaymentIntent) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PaymentIntent) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetMatchAvsa

`func (o *PaymentIntent) GetMatchAvsa() string`

GetMatchAvsa returns the MatchAvsa field if non-nil, zero value otherwise.

### GetMatchAvsaOk

`func (o *PaymentIntent) GetMatchAvsaOk() (*string, bool)`

GetMatchAvsaOk returns a tuple with the MatchAvsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAvsa

`func (o *PaymentIntent) SetMatchAvsa(v string)`

SetMatchAvsa sets MatchAvsa field to given value.

### HasMatchAvsa

`func (o *PaymentIntent) HasMatchAvsa() bool`

HasMatchAvsa returns a boolean if a field has been set.

### GetShipTo

`func (o *PaymentIntent) GetShipTo() ContactDetails`

GetShipTo returns the ShipTo field if non-nil, zero value otherwise.

### GetShipToOk

`func (o *PaymentIntent) GetShipToOk() (*ContactDetails, bool)`

GetShipToOk returns a tuple with the ShipTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipTo

`func (o *PaymentIntent) SetShipTo(v ContactDetails)`

SetShipTo sets ShipTo field to given value.

### HasShipTo

`func (o *PaymentIntent) HasShipTo() bool`

HasShipTo returns a boolean if a field has been set.

### GetTag

`func (o *PaymentIntent) GetTag() []string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PaymentIntent) GetTagOk() (*[]string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PaymentIntent) SetTag(v []string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PaymentIntent) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTransInfo

`func (o *PaymentIntent) GetTransInfo() string`

GetTransInfo returns the TransInfo field if non-nil, zero value otherwise.

### GetTransInfoOk

`func (o *PaymentIntent) GetTransInfoOk() (*string, bool)`

GetTransInfoOk returns a tuple with the TransInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransInfo

`func (o *PaymentIntent) SetTransInfo(v string)`

SetTransInfo sets TransInfo field to given value.

### HasTransInfo

`func (o *PaymentIntent) HasTransInfo() bool`

HasTransInfo returns a boolean if a field has been set.

### GetTransType

`func (o *PaymentIntent) GetTransType() string`

GetTransType returns the TransType field if non-nil, zero value otherwise.

### GetTransTypeOk

`func (o *PaymentIntent) GetTransTypeOk() (*string, bool)`

GetTransTypeOk returns a tuple with the TransType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransType

`func (o *PaymentIntent) SetTransType(v string)`

SetTransType sets TransType field to given value.

### HasTransType

`func (o *PaymentIntent) HasTransType() bool`

HasTransType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


