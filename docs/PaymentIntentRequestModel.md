# PaymentIntentRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adjustments** | Pointer to [**Adjustments**](Adjustments.md) |  | [optional] 
**Amount** | **int32** | The amount to authorise in the lowest unit of currency with a variable length to a maximum of 12 digits.  No decimal points are to be included and no divisional characters such as 1,024.  The amount should be the total amount required for the transaction.  For example with GBP £1,021.95 the amount value is 102195.  | 
**AvsPostcodePolicy** | Pointer to **string** | A policy value which determines whether an AVS postcode policy is enforced or bypassed.  Values are:   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be rejected if the AVS postcode numeric value does not match.   &#x60;2&#x60; to bypass. Transactions that are bypassed will be allowed through even if the postcode did not match.   &#x60;3&#x60; to ignore. Transactions that are ignored will bypass the result and not send postcode details for authorisation.  | [optional] 
**BillTo** | Pointer to [**ContactDetails**](ContactDetails.md) |  | [optional] 
**CscPolicy** | Pointer to **string** | A policy value which determines whether a CSC policy is enforced or bypassed.  Values are:   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be rejected if the CSC value does not match.   &#x60;2&#x60; to bypass. Transactions that are bypassed will be allowed through even if the CSC did not match.   &#x60;3&#x60; to ignore. Transactions that are ignored will bypass the result and not send the CSC details for authorisation.  | [optional] 
**Currency** | Pointer to **string** | The processing currency for the transaction. Will default to the merchant account currency. | [optional] 
**DuplicatePolicy** | Pointer to **string** | A policy value which determines whether a duplication policy is enforced or bypassed. A duplication check has a window of time set against your account within which it can action. If a previous transaction with matching values occurred within the window, any subsequent transaction will result in a T001 result.  Values are   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be checked for duplication within the duplication window.   &#x60;2&#x60; to bypass. Transactions that are bypassed will not be checked for duplication within the duplication window.   &#x60;3&#x60; to ignore. Transactions that are ignored will have the same affect as bypass.  | [optional] 
**ExternalRef** | Pointer to **string** | A unique identifier, such as an order ID or invoice number, provided by your accounting or billing system to link the payment intent with an external system reference. This ensures traceability across systems for audits and transaction validation. | [optional] 
**ExternalRefSource** | Pointer to **string** | Specifies the originating source or system of the external reference, helping to categorise and trace the context of the external identifier, whether it comes from an internal system, third-party vendor, or external financial platform. | [optional] 
**Identifier** | **string** | The identifier of the transaction to process. The value should be a valid reference and may be used to perform  post processing actions and to aid in reconciliation of transactions.  The value should be a valid printable string with ASCII character ranges from 0x32 to 0x127.  The identifier is recommended to be distinct for each transaction such as a [random unique identifier](https://en.wikipedia.org/wiki/Universally_unique_identifier) this will aid in ensuring each transaction is identifiable.  When transactions are processed they are also checked for duplicate requests. Changing the identifier on a subsequent request will ensure that a transaction is considered as different.  | 
**MatchAvsa** | Pointer to **string** | A policy value which determines whether an AVS address policy is enforced, bypassed or ignored.  Values are:   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be rejected if the AVS address numeric value does not match.   &#x60;2&#x60; to bypass. Transactions that are bypassed will be allowed through even if the address did not match.   &#x60;3&#x60; to ignore. Transactions that are ignored will bypass the result and not send address numeric details for authorisation.  | [optional] 
**Merchantid** | Pointer to **int32** | The merchant id of the intent, required if using the API key or not required if using a domain key. | [optional] 
**PreAuth** | Pointer to **string** | A policy value which determines whether a pre auth policy is enforced or bypassed.  Values are:   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy.  Enforces pre-authorisation when it does not pre-auth by default.   &#x60;2&#x60; to bypass. Bypasses pre-authorisation when it is enabled to pre auth by default.   &#x60;3&#x60; to ignore. The same as the default policy (0). Although it currently mirrors the default, this option is included for compatibility with other policies.  | [optional] 
**ShipTo** | Pointer to [**ContactDetails**](ContactDetails.md) |  | [optional] 
**Tag** | Pointer to **[]string** |  | [optional] 
**TransInfo** | Pointer to **string** | Further information that can be added to the transaction will display in reporting. Can be used for flexible values such as operator id. | [optional] 
**TransType** | Pointer to **string** | The type of transaction being submitted. Normally this value is not required and your account manager may request that you set this field. | [optional] 

## Methods

### NewPaymentIntentRequestModel

`func NewPaymentIntentRequestModel(amount int32, identifier string, ) *PaymentIntentRequestModel`

NewPaymentIntentRequestModel instantiates a new PaymentIntentRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentIntentRequestModelWithDefaults

`func NewPaymentIntentRequestModelWithDefaults() *PaymentIntentRequestModel`

NewPaymentIntentRequestModelWithDefaults instantiates a new PaymentIntentRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjustments

`func (o *PaymentIntentRequestModel) GetAdjustments() Adjustments`

GetAdjustments returns the Adjustments field if non-nil, zero value otherwise.

### GetAdjustmentsOk

`func (o *PaymentIntentRequestModel) GetAdjustmentsOk() (*Adjustments, bool)`

GetAdjustmentsOk returns a tuple with the Adjustments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustments

`func (o *PaymentIntentRequestModel) SetAdjustments(v Adjustments)`

SetAdjustments sets Adjustments field to given value.

### HasAdjustments

`func (o *PaymentIntentRequestModel) HasAdjustments() bool`

HasAdjustments returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentIntentRequestModel) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentIntentRequestModel) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentIntentRequestModel) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAvsPostcodePolicy

`func (o *PaymentIntentRequestModel) GetAvsPostcodePolicy() string`

GetAvsPostcodePolicy returns the AvsPostcodePolicy field if non-nil, zero value otherwise.

### GetAvsPostcodePolicyOk

`func (o *PaymentIntentRequestModel) GetAvsPostcodePolicyOk() (*string, bool)`

GetAvsPostcodePolicyOk returns a tuple with the AvsPostcodePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsPostcodePolicy

`func (o *PaymentIntentRequestModel) SetAvsPostcodePolicy(v string)`

SetAvsPostcodePolicy sets AvsPostcodePolicy field to given value.

### HasAvsPostcodePolicy

`func (o *PaymentIntentRequestModel) HasAvsPostcodePolicy() bool`

HasAvsPostcodePolicy returns a boolean if a field has been set.

### GetBillTo

`func (o *PaymentIntentRequestModel) GetBillTo() ContactDetails`

GetBillTo returns the BillTo field if non-nil, zero value otherwise.

### GetBillToOk

`func (o *PaymentIntentRequestModel) GetBillToOk() (*ContactDetails, bool)`

GetBillToOk returns a tuple with the BillTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTo

`func (o *PaymentIntentRequestModel) SetBillTo(v ContactDetails)`

SetBillTo sets BillTo field to given value.

### HasBillTo

`func (o *PaymentIntentRequestModel) HasBillTo() bool`

HasBillTo returns a boolean if a field has been set.

### GetCscPolicy

`func (o *PaymentIntentRequestModel) GetCscPolicy() string`

GetCscPolicy returns the CscPolicy field if non-nil, zero value otherwise.

### GetCscPolicyOk

`func (o *PaymentIntentRequestModel) GetCscPolicyOk() (*string, bool)`

GetCscPolicyOk returns a tuple with the CscPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCscPolicy

`func (o *PaymentIntentRequestModel) SetCscPolicy(v string)`

SetCscPolicy sets CscPolicy field to given value.

### HasCscPolicy

`func (o *PaymentIntentRequestModel) HasCscPolicy() bool`

HasCscPolicy returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentIntentRequestModel) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentIntentRequestModel) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentIntentRequestModel) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentIntentRequestModel) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDuplicatePolicy

`func (o *PaymentIntentRequestModel) GetDuplicatePolicy() string`

GetDuplicatePolicy returns the DuplicatePolicy field if non-nil, zero value otherwise.

### GetDuplicatePolicyOk

`func (o *PaymentIntentRequestModel) GetDuplicatePolicyOk() (*string, bool)`

GetDuplicatePolicyOk returns a tuple with the DuplicatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicatePolicy

`func (o *PaymentIntentRequestModel) SetDuplicatePolicy(v string)`

SetDuplicatePolicy sets DuplicatePolicy field to given value.

### HasDuplicatePolicy

`func (o *PaymentIntentRequestModel) HasDuplicatePolicy() bool`

HasDuplicatePolicy returns a boolean if a field has been set.

### GetExternalRef

`func (o *PaymentIntentRequestModel) GetExternalRef() string`

GetExternalRef returns the ExternalRef field if non-nil, zero value otherwise.

### GetExternalRefOk

`func (o *PaymentIntentRequestModel) GetExternalRefOk() (*string, bool)`

GetExternalRefOk returns a tuple with the ExternalRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRef

`func (o *PaymentIntentRequestModel) SetExternalRef(v string)`

SetExternalRef sets ExternalRef field to given value.

### HasExternalRef

`func (o *PaymentIntentRequestModel) HasExternalRef() bool`

HasExternalRef returns a boolean if a field has been set.

### GetExternalRefSource

`func (o *PaymentIntentRequestModel) GetExternalRefSource() string`

GetExternalRefSource returns the ExternalRefSource field if non-nil, zero value otherwise.

### GetExternalRefSourceOk

`func (o *PaymentIntentRequestModel) GetExternalRefSourceOk() (*string, bool)`

GetExternalRefSourceOk returns a tuple with the ExternalRefSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefSource

`func (o *PaymentIntentRequestModel) SetExternalRefSource(v string)`

SetExternalRefSource sets ExternalRefSource field to given value.

### HasExternalRefSource

`func (o *PaymentIntentRequestModel) HasExternalRefSource() bool`

HasExternalRefSource returns a boolean if a field has been set.

### GetIdentifier

`func (o *PaymentIntentRequestModel) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PaymentIntentRequestModel) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PaymentIntentRequestModel) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetMatchAvsa

`func (o *PaymentIntentRequestModel) GetMatchAvsa() string`

GetMatchAvsa returns the MatchAvsa field if non-nil, zero value otherwise.

### GetMatchAvsaOk

`func (o *PaymentIntentRequestModel) GetMatchAvsaOk() (*string, bool)`

GetMatchAvsaOk returns a tuple with the MatchAvsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAvsa

`func (o *PaymentIntentRequestModel) SetMatchAvsa(v string)`

SetMatchAvsa sets MatchAvsa field to given value.

### HasMatchAvsa

`func (o *PaymentIntentRequestModel) HasMatchAvsa() bool`

HasMatchAvsa returns a boolean if a field has been set.

### GetMerchantid

`func (o *PaymentIntentRequestModel) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *PaymentIntentRequestModel) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *PaymentIntentRequestModel) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.

### HasMerchantid

`func (o *PaymentIntentRequestModel) HasMerchantid() bool`

HasMerchantid returns a boolean if a field has been set.

### GetPreAuth

`func (o *PaymentIntentRequestModel) GetPreAuth() string`

GetPreAuth returns the PreAuth field if non-nil, zero value otherwise.

### GetPreAuthOk

`func (o *PaymentIntentRequestModel) GetPreAuthOk() (*string, bool)`

GetPreAuthOk returns a tuple with the PreAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuth

`func (o *PaymentIntentRequestModel) SetPreAuth(v string)`

SetPreAuth sets PreAuth field to given value.

### HasPreAuth

`func (o *PaymentIntentRequestModel) HasPreAuth() bool`

HasPreAuth returns a boolean if a field has been set.

### GetShipTo

`func (o *PaymentIntentRequestModel) GetShipTo() ContactDetails`

GetShipTo returns the ShipTo field if non-nil, zero value otherwise.

### GetShipToOk

`func (o *PaymentIntentRequestModel) GetShipToOk() (*ContactDetails, bool)`

GetShipToOk returns a tuple with the ShipTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipTo

`func (o *PaymentIntentRequestModel) SetShipTo(v ContactDetails)`

SetShipTo sets ShipTo field to given value.

### HasShipTo

`func (o *PaymentIntentRequestModel) HasShipTo() bool`

HasShipTo returns a boolean if a field has been set.

### GetTag

`func (o *PaymentIntentRequestModel) GetTag() []string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PaymentIntentRequestModel) GetTagOk() (*[]string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PaymentIntentRequestModel) SetTag(v []string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PaymentIntentRequestModel) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTransInfo

`func (o *PaymentIntentRequestModel) GetTransInfo() string`

GetTransInfo returns the TransInfo field if non-nil, zero value otherwise.

### GetTransInfoOk

`func (o *PaymentIntentRequestModel) GetTransInfoOk() (*string, bool)`

GetTransInfoOk returns a tuple with the TransInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransInfo

`func (o *PaymentIntentRequestModel) SetTransInfo(v string)`

SetTransInfo sets TransInfo field to given value.

### HasTransInfo

`func (o *PaymentIntentRequestModel) HasTransInfo() bool`

HasTransInfo returns a boolean if a field has been set.

### GetTransType

`func (o *PaymentIntentRequestModel) GetTransType() string`

GetTransType returns the TransType field if non-nil, zero value otherwise.

### GetTransTypeOk

`func (o *PaymentIntentRequestModel) GetTransTypeOk() (*string, bool)`

GetTransTypeOk returns a tuple with the TransType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransType

`func (o *PaymentIntentRequestModel) SetTransType(v string)`

SetTransType sets TransType field to given value.

### HasTransType

`func (o *PaymentIntentRequestModel) HasTransType() bool`

HasTransType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


