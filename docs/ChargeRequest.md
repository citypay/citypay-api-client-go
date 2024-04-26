# ChargeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The amount to authorise in the lowest unit of currency with a variable length to a maximum of 12 digits.  No decimal points are to be included and no divisional characters such as 1,024.  The amount should be the total amount required for the transaction.  For example with GBP £1,021.95 the amount value is 102195.  | 
**AvsPostcodePolicy** | Pointer to **string** | A policy value which determines whether an AVS postcode policy is enforced or bypassed.  Values are:   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be rejected if the AVS postcode numeric value does not match.   &#x60;2&#x60; to bypass. Transactions that are bypassed will be allowed through even if the postcode did not match.   &#x60;3&#x60; to ignore. Transactions that are ignored will bypass the result and not send postcode details for authorisation.  | [optional] 
**CardholderAgreement** | Pointer to **string** | Merchant-initiated transactions (MITs) are payments you trigger, where the cardholder has previously consented to you carrying out such payments. These may be scheduled (such as recurring payments and installments) or unscheduled (like account top-ups triggered by balance thresholds and no-show charges).  Scheduled --- These are regular payments using stored card details, like installments or a monthly subscription fee.  - &#x60;I&#x60; Instalment - A single purchase of goods or services billed to a cardholder in multiple transactions, over a period of time agreed by the cardholder and you.  - &#x60;R&#x60; Recurring - Transactions processed at fixed, regular intervals not to exceed one year between transactions, representing an agreement between a cardholder and you to purchase goods or services provided over a period of time.  Unscheduled --- These are payments using stored card details that do not occur on a regular schedule, like top-ups for a digital wallet triggered by the balance falling below a certain threshold.  - &#x60;A&#x60; Reauthorisation - a purchase made after the original purchase. A common scenario is delayed/split shipments.  - &#x60;C&#x60; Unscheduled Payment - A transaction using a stored credential for a fixed or variable amount that does not occur on a scheduled or regularly occurring transaction date. This includes account top-ups triggered by balance thresholds.  - &#x60;D&#x60; Delayed Charge - A delayed charge is typically used in hotel, cruise lines and vehicle rental environments to perform a supplemental account charge after original services are rendered.  - &#x60;L&#x60; Incremental - An incremental authorisation is typically found in hotel and car rental environments, where the cardholder has agreed to pay for any service incurred during the duration of the contract. An incremental authorisation is where you need to seek authorisation of further funds in addition to what you have originally requested. A common scenario is additional services charged to the contract, such as extending a stay in a hotel.  - &#x60;S&#x60; Resubmission - When the original purchase occurred, but you were not able to get authorisation at the time the goods or services were provided. It should be only used where the goods or services have already been provided, but the authorisation request is declined for insufficient funds.  - &#x60;X&#x60; No-show - A no-show is a transaction where you are enabled to charge for services which the cardholder entered into an agreement to purchase, but the cardholder did not meet the terms of the agreement.  | [optional] 
**Csc** | Pointer to **string** | The Card Security Code (CSC) (also known as CV2/CVV2) is normally found on the back of the card (American Express has it on the front). The value helps to identify possession of the card as it is not available within the chip or magnetic swipe.  When forwarding the CSC, please ensure the value is a string as some values start with 0 and this will be stripped out by any integer parsing.  The CSC number aids fraud prevention in Mail Order and Internet payments.  Business rules are available on your account to identify whether to accept or decline transactions based on mismatched results of the CSC.  The Payment Card Industry (PCI) requires that at no stage of a transaction should the CSC be stored.  This applies to all entities handling card data.  It should also not be used in any hashing process.  CityPay do not store the value and have no method of retrieving the value once the transaction has been processed. For this reason, duplicate checking is unable to determine the CSC in its duplication check algorithm.  | [optional] 
**CscPolicy** | Pointer to **string** | A policy value which determines whether a CSC policy is enforced or bypassed.  Values are:   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be rejected if the CSC value does not match.   &#x60;2&#x60; to bypass. Transactions that are bypassed will be allowed through even if the CSC did not match.   &#x60;3&#x60; to ignore. Transactions that are ignored will bypass the result and not send the CSC details for authorisation.  | [optional] 
**Currency** | Pointer to **string** | The processing currency for the transaction. Will default to the merchant account currency. | [optional] 
**DuplicatePolicy** | Pointer to **string** | A policy value which determines whether a duplication policy is enforced or bypassed. A duplication check has a window of time set against your account within which it can action. If a previous transaction with matching values occurred within the window, any subsequent transaction will result in a T001 result.  Values are   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be checked for duplication within the duplication window.   &#x60;2&#x60; to bypass. Transactions that are bypassed will not be checked for duplication within the duplication window.   &#x60;3&#x60; to ignore. Transactions that are ignored will have the same affect as bypass.  | [optional] 
**Identifier** | **string** | The identifier of the transaction to process. The value should be a valid reference and may be used to perform  post processing actions and to aid in reconciliation of transactions.  The value should be a valid printable string with ASCII character ranges from 0x32 to 0x127.  The identifier is recommended to be distinct for each transaction such as a [random unique identifier](https://en.wikipedia.org/wiki/Universally_unique_identifier) this will aid in ensuring each transaction is identifiable.  When transactions are processed they are also checked for duplicate requests. Changing the identifier on a subsequent request will ensure that a transaction is considered as different.  | 
**Initiation** | Pointer to **string** | Transactions charged using the API are defined as:  **Cardholder Initiated**: A _cardholder initiated transaction_ (CIT) is where the cardholder selects the card for use for a purchase using previously stored details. An example would be a customer buying an item from your website after being present with their saved card details at checkout.  **Merchant Intiated**: A _merchant initiated transaction_ (MIT) is an authorisation initiated where you as the  merchant submit a cardholders previously stored details without the cardholder&#39;s participation. An example would  be a subscription to a membership scheme to debit their card monthly.  MITs have different reasons such as reauthorisation, delayed, unscheduled, incremental, recurring, instalment, no-show or resubmission.  The following values apply   - &#x60;M&#x60; - specifies that the transaction is initiated by the merchant   - &#x60;C&#x60; - specifies that the transaction is initiated by the cardholder  Where transactions are merchant initiated, a valid cardholder agreement must be defined.  | [optional] 
**MatchAvsa** | Pointer to **string** | A policy value which determines whether an AVS address policy is enforced, bypassed or ignored.  Values are:   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be rejected if the AVS address numeric value does not match.   &#x60;2&#x60; to bypass. Transactions that are bypassed will be allowed through even if the address did not match.   &#x60;3&#x60; to ignore. Transactions that are ignored will bypass the result and not send address numeric details for authorisation.  | [optional] 
**Merchantid** | **int32** | Identifies the merchant account to perform processing for. | 
**Tag** | Pointer to **[]string** |  | [optional] 
**Threedsecure** | Pointer to [**ThreeDSecure**](ThreeDSecure.md) |  | [optional] 
**Token** | **string** | A tokenised form of a card that belongs to a card holder&#39;s account and that has been previously registered. The token is time based and will only be active for a short duration. The value is therefore designed not to be stored remotely for future use.   Tokens will start with ct and are resiliently tamper proof using HMacSHA-256. No sensitive card data is stored internally within the token.   Each card will contain a different token and the value may be different on any retrieval call.   The value can be presented for payment as a selection value to an end user in a web application.  | 
**TransInfo** | Pointer to **string** | Further information that can be added to the transaction will display in reporting. Can be used for flexible values such as operator id. | [optional] 
**TransType** | Pointer to **string** | The type of transaction being submitted. Normally this value is not required and your account manager may request that you set this field. | [optional] 

## Methods

### NewChargeRequest

`func NewChargeRequest(amount int32, identifier string, merchantid int32, token string, ) *ChargeRequest`

NewChargeRequest instantiates a new ChargeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeRequestWithDefaults

`func NewChargeRequestWithDefaults() *ChargeRequest`

NewChargeRequestWithDefaults instantiates a new ChargeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ChargeRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ChargeRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ChargeRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAvsPostcodePolicy

`func (o *ChargeRequest) GetAvsPostcodePolicy() string`

GetAvsPostcodePolicy returns the AvsPostcodePolicy field if non-nil, zero value otherwise.

### GetAvsPostcodePolicyOk

`func (o *ChargeRequest) GetAvsPostcodePolicyOk() (*string, bool)`

GetAvsPostcodePolicyOk returns a tuple with the AvsPostcodePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsPostcodePolicy

`func (o *ChargeRequest) SetAvsPostcodePolicy(v string)`

SetAvsPostcodePolicy sets AvsPostcodePolicy field to given value.

### HasAvsPostcodePolicy

`func (o *ChargeRequest) HasAvsPostcodePolicy() bool`

HasAvsPostcodePolicy returns a boolean if a field has been set.

### GetCardholderAgreement

`func (o *ChargeRequest) GetCardholderAgreement() string`

GetCardholderAgreement returns the CardholderAgreement field if non-nil, zero value otherwise.

### GetCardholderAgreementOk

`func (o *ChargeRequest) GetCardholderAgreementOk() (*string, bool)`

GetCardholderAgreementOk returns a tuple with the CardholderAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardholderAgreement

`func (o *ChargeRequest) SetCardholderAgreement(v string)`

SetCardholderAgreement sets CardholderAgreement field to given value.

### HasCardholderAgreement

`func (o *ChargeRequest) HasCardholderAgreement() bool`

HasCardholderAgreement returns a boolean if a field has been set.

### GetCsc

`func (o *ChargeRequest) GetCsc() string`

GetCsc returns the Csc field if non-nil, zero value otherwise.

### GetCscOk

`func (o *ChargeRequest) GetCscOk() (*string, bool)`

GetCscOk returns a tuple with the Csc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsc

`func (o *ChargeRequest) SetCsc(v string)`

SetCsc sets Csc field to given value.

### HasCsc

`func (o *ChargeRequest) HasCsc() bool`

HasCsc returns a boolean if a field has been set.

### GetCscPolicy

`func (o *ChargeRequest) GetCscPolicy() string`

GetCscPolicy returns the CscPolicy field if non-nil, zero value otherwise.

### GetCscPolicyOk

`func (o *ChargeRequest) GetCscPolicyOk() (*string, bool)`

GetCscPolicyOk returns a tuple with the CscPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCscPolicy

`func (o *ChargeRequest) SetCscPolicy(v string)`

SetCscPolicy sets CscPolicy field to given value.

### HasCscPolicy

`func (o *ChargeRequest) HasCscPolicy() bool`

HasCscPolicy returns a boolean if a field has been set.

### GetCurrency

`func (o *ChargeRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ChargeRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ChargeRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ChargeRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDuplicatePolicy

`func (o *ChargeRequest) GetDuplicatePolicy() string`

GetDuplicatePolicy returns the DuplicatePolicy field if non-nil, zero value otherwise.

### GetDuplicatePolicyOk

`func (o *ChargeRequest) GetDuplicatePolicyOk() (*string, bool)`

GetDuplicatePolicyOk returns a tuple with the DuplicatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicatePolicy

`func (o *ChargeRequest) SetDuplicatePolicy(v string)`

SetDuplicatePolicy sets DuplicatePolicy field to given value.

### HasDuplicatePolicy

`func (o *ChargeRequest) HasDuplicatePolicy() bool`

HasDuplicatePolicy returns a boolean if a field has been set.

### GetIdentifier

`func (o *ChargeRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ChargeRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ChargeRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetInitiation

`func (o *ChargeRequest) GetInitiation() string`

GetInitiation returns the Initiation field if non-nil, zero value otherwise.

### GetInitiationOk

`func (o *ChargeRequest) GetInitiationOk() (*string, bool)`

GetInitiationOk returns a tuple with the Initiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiation

`func (o *ChargeRequest) SetInitiation(v string)`

SetInitiation sets Initiation field to given value.

### HasInitiation

`func (o *ChargeRequest) HasInitiation() bool`

HasInitiation returns a boolean if a field has been set.

### GetMatchAvsa

`func (o *ChargeRequest) GetMatchAvsa() string`

GetMatchAvsa returns the MatchAvsa field if non-nil, zero value otherwise.

### GetMatchAvsaOk

`func (o *ChargeRequest) GetMatchAvsaOk() (*string, bool)`

GetMatchAvsaOk returns a tuple with the MatchAvsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAvsa

`func (o *ChargeRequest) SetMatchAvsa(v string)`

SetMatchAvsa sets MatchAvsa field to given value.

### HasMatchAvsa

`func (o *ChargeRequest) HasMatchAvsa() bool`

HasMatchAvsa returns a boolean if a field has been set.

### GetMerchantid

`func (o *ChargeRequest) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *ChargeRequest) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *ChargeRequest) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.


### GetTag

`func (o *ChargeRequest) GetTag() []string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ChargeRequest) GetTagOk() (*[]string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ChargeRequest) SetTag(v []string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ChargeRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetThreedsecure

`func (o *ChargeRequest) GetThreedsecure() ThreeDSecure`

GetThreedsecure returns the Threedsecure field if non-nil, zero value otherwise.

### GetThreedsecureOk

`func (o *ChargeRequest) GetThreedsecureOk() (*ThreeDSecure, bool)`

GetThreedsecureOk returns a tuple with the Threedsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreedsecure

`func (o *ChargeRequest) SetThreedsecure(v ThreeDSecure)`

SetThreedsecure sets Threedsecure field to given value.

### HasThreedsecure

`func (o *ChargeRequest) HasThreedsecure() bool`

HasThreedsecure returns a boolean if a field has been set.

### GetToken

`func (o *ChargeRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ChargeRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ChargeRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetTransInfo

`func (o *ChargeRequest) GetTransInfo() string`

GetTransInfo returns the TransInfo field if non-nil, zero value otherwise.

### GetTransInfoOk

`func (o *ChargeRequest) GetTransInfoOk() (*string, bool)`

GetTransInfoOk returns a tuple with the TransInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransInfo

`func (o *ChargeRequest) SetTransInfo(v string)`

SetTransInfo sets TransInfo field to given value.

### HasTransInfo

`func (o *ChargeRequest) HasTransInfo() bool`

HasTransInfo returns a boolean if a field has been set.

### GetTransType

`func (o *ChargeRequest) GetTransType() string`

GetTransType returns the TransType field if non-nil, zero value otherwise.

### GetTransTypeOk

`func (o *ChargeRequest) GetTransTypeOk() (*string, bool)`

GetTransTypeOk returns a tuple with the TransType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransType

`func (o *ChargeRequest) SetTransType(v string)`

SetTransType sets TransType field to given value.

### HasTransType

`func (o *ChargeRequest) HasTransType() bool`

HasTransType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


