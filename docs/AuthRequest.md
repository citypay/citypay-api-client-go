# AuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AirlineData** | Pointer to [**AirlineAdvice**](AirlineAdvice.md) |  | [optional] 
**Amount** | **int32** | The amount to authorise in the lowest unit of currency with a variable length to a maximum of 12 digits.  No decimal points are to be included and no divisional characters such as 1,024.  The amount should be the total amount required for the transaction.  For example with GBP £1,021.95 the amount value is 102195.  | 
**AvsPostcodePolicy** | Pointer to **string** | A policy value which determines whether an AVS postcode policy is enforced or bypassed.  Values are:   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be rejected if the AVS postcode numeric value does not match.   &#x60;2&#x60; to bypass. Transactions that are bypassed will be allowed through even if the postcode did not match.   &#x60;3&#x60; to ignore. Transactions that are ignored will bypass the result and not send postcode details for authorisation.  | [optional] 
**BillTo** | Pointer to [**ContactDetails**](ContactDetails.md) |  | [optional] 
**Cardnumber** | **string** | The card number (PAN) with a variable length to a maximum of 21 digits in numerical form. Any non numeric characters will be stripped out of the card number, this includes whitespace or separators internal of the provided value.  The card number must be treated as sensitive data. We only provide an obfuscated value in logging and reporting.  The plaintext value is encrypted in our database using AES 256 GMC bit encryption for settlement or refund purposes.  When providing the card number to our gateway through the authorisation API you will be handling the card data on your application. This will require further PCI controls to be in place and this value must never be stored.  | 
**Csc** | Pointer to **string** | The Card Security Code (CSC) (also known as CV2/CVV2) is normally found on the back of the card (American Express has it on the front). The value helps to identify possession of the card as it is not available within the chip or magnetic swipe.  When forwarding the CSC, please ensure the value is a string as some values start with 0 and this will be stripped out by any integer parsing.  The CSC number aids fraud prevention in Mail Order and Internet payments.  Business rules are available on your account to identify whether to accept or decline transactions based on mismatched results of the CSC.  The Payment Card Industry (PCI) requires that at no stage of a transaction should the CSC be stored.  This applies to all entities handling card data.  It should also not be used in any hashing process.  CityPay do not store the value and have no method of retrieving the value once the transaction has been processed. For this reason, duplicate checking is unable to determine the CSC in its duplication check algorithm.  | [optional] 
**CscPolicy** | Pointer to **string** | A policy value which determines whether a CSC policy is enforced or bypassed.  Values are:   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be rejected if the CSC value does not match.   &#x60;2&#x60; to bypass. Transactions that are bypassed will be allowed through even if the CSC did not match.   &#x60;3&#x60; to ignore. Transactions that are ignored will bypass the result and not send the CSC details for authorisation.  | [optional] 
**Currency** | Pointer to **string** | The processing currency for the transaction. Will default to the merchant account currency. | [optional] 
**DuplicatePolicy** | Pointer to **string** | A policy value which determines whether a duplication policy is enforced or bypassed. A duplication check has a window of time set against your account within which it can action. If a previous transaction with matching values occurred within the window, any subsequent transaction will result in a T001 result.  Values are   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be checked for duplication within the duplication window.   &#x60;2&#x60; to bypass. Transactions that are bypassed will not be checked for duplication within the duplication window.   &#x60;3&#x60; to ignore. Transactions that are ignored will have the same affect as bypass.  | [optional] 
**EventManagement** | Pointer to [**EventDataModel**](EventDataModel.md) |  | [optional] 
**Expmonth** | **int32** | The month of expiry of the card. The month value should be a numerical value between 1 and 12.  | 
**Expyear** | **int32** | The year of expiry of the card.  | 
**ExternalMpi** | Pointer to [**ExternalMPI**](ExternalMPI.md) |  | [optional] 
**Identifier** | **string** | The identifier of the transaction to process. The value should be a valid reference and may be used to perform  post processing actions and to aid in reconciliation of transactions.  The value should be a valid printable string with ASCII character ranges from 0x32 to 0x127.  The identifier is recommended to be distinct for each transaction such as a [random unique identifier](https://en.wikipedia.org/wiki/Universally_unique_identifier) this will aid in ensuring each transaction is identifiable.  When transactions are processed they are also checked for duplicate requests. Changing the identifier on a subsequent request will ensure that a transaction is considered as different.  | 
**MatchAvsa** | Pointer to **string** | A policy value which determines whether an AVS address policy is enforced, bypassed or ignored.  Values are:   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be rejected if the AVS address numeric value does not match.   &#x60;2&#x60; to bypass. Transactions that are bypassed will be allowed through even if the address did not match.   &#x60;3&#x60; to ignore. Transactions that are ignored will bypass the result and not send address numeric details for authorisation.  | [optional] 
**Mcc6012** | Pointer to [**MCC6012**](MCC6012.md) |  | [optional] 
**Merchantid** | **int32** | Identifies the merchant account to perform processing for. | 
**NameOnCard** | Pointer to **string** | The card holder name as appears on the card such as MR N E BODY. Required for some acquirers.  | [optional] 
**ShipTo** | Pointer to [**ContactDetails**](ContactDetails.md) |  | [optional] 
**Tag** | Pointer to **[]string** |  | [optional] 
**Threedsecure** | Pointer to [**ThreeDSecure**](ThreeDSecure.md) |  | [optional] 
**TransInfo** | Pointer to **string** | Further information that can be added to the transaction will display in reporting. Can be used for flexible values such as operator id. | [optional] 
**TransType** | Pointer to **string** | The type of transaction being submitted. Normally this value is not required and your account manager may request that you set this field. | [optional] 

## Methods

### NewAuthRequest

`func NewAuthRequest(amount int32, cardnumber string, expmonth int32, expyear int32, identifier string, merchantid int32, ) *AuthRequest`

NewAuthRequest instantiates a new AuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthRequestWithDefaults

`func NewAuthRequestWithDefaults() *AuthRequest`

NewAuthRequestWithDefaults instantiates a new AuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAirlineData

`func (o *AuthRequest) GetAirlineData() AirlineAdvice`

GetAirlineData returns the AirlineData field if non-nil, zero value otherwise.

### GetAirlineDataOk

`func (o *AuthRequest) GetAirlineDataOk() (*AirlineAdvice, bool)`

GetAirlineDataOk returns a tuple with the AirlineData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineData

`func (o *AuthRequest) SetAirlineData(v AirlineAdvice)`

SetAirlineData sets AirlineData field to given value.

### HasAirlineData

`func (o *AuthRequest) HasAirlineData() bool`

HasAirlineData returns a boolean if a field has been set.

### GetAmount

`func (o *AuthRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AuthRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AuthRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAvsPostcodePolicy

`func (o *AuthRequest) GetAvsPostcodePolicy() string`

GetAvsPostcodePolicy returns the AvsPostcodePolicy field if non-nil, zero value otherwise.

### GetAvsPostcodePolicyOk

`func (o *AuthRequest) GetAvsPostcodePolicyOk() (*string, bool)`

GetAvsPostcodePolicyOk returns a tuple with the AvsPostcodePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsPostcodePolicy

`func (o *AuthRequest) SetAvsPostcodePolicy(v string)`

SetAvsPostcodePolicy sets AvsPostcodePolicy field to given value.

### HasAvsPostcodePolicy

`func (o *AuthRequest) HasAvsPostcodePolicy() bool`

HasAvsPostcodePolicy returns a boolean if a field has been set.

### GetBillTo

`func (o *AuthRequest) GetBillTo() ContactDetails`

GetBillTo returns the BillTo field if non-nil, zero value otherwise.

### GetBillToOk

`func (o *AuthRequest) GetBillToOk() (*ContactDetails, bool)`

GetBillToOk returns a tuple with the BillTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTo

`func (o *AuthRequest) SetBillTo(v ContactDetails)`

SetBillTo sets BillTo field to given value.

### HasBillTo

`func (o *AuthRequest) HasBillTo() bool`

HasBillTo returns a boolean if a field has been set.

### GetCardnumber

`func (o *AuthRequest) GetCardnumber() string`

GetCardnumber returns the Cardnumber field if non-nil, zero value otherwise.

### GetCardnumberOk

`func (o *AuthRequest) GetCardnumberOk() (*string, bool)`

GetCardnumberOk returns a tuple with the Cardnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardnumber

`func (o *AuthRequest) SetCardnumber(v string)`

SetCardnumber sets Cardnumber field to given value.


### GetCsc

`func (o *AuthRequest) GetCsc() string`

GetCsc returns the Csc field if non-nil, zero value otherwise.

### GetCscOk

`func (o *AuthRequest) GetCscOk() (*string, bool)`

GetCscOk returns a tuple with the Csc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsc

`func (o *AuthRequest) SetCsc(v string)`

SetCsc sets Csc field to given value.

### HasCsc

`func (o *AuthRequest) HasCsc() bool`

HasCsc returns a boolean if a field has been set.

### GetCscPolicy

`func (o *AuthRequest) GetCscPolicy() string`

GetCscPolicy returns the CscPolicy field if non-nil, zero value otherwise.

### GetCscPolicyOk

`func (o *AuthRequest) GetCscPolicyOk() (*string, bool)`

GetCscPolicyOk returns a tuple with the CscPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCscPolicy

`func (o *AuthRequest) SetCscPolicy(v string)`

SetCscPolicy sets CscPolicy field to given value.

### HasCscPolicy

`func (o *AuthRequest) HasCscPolicy() bool`

HasCscPolicy returns a boolean if a field has been set.

### GetCurrency

`func (o *AuthRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AuthRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AuthRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AuthRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDuplicatePolicy

`func (o *AuthRequest) GetDuplicatePolicy() string`

GetDuplicatePolicy returns the DuplicatePolicy field if non-nil, zero value otherwise.

### GetDuplicatePolicyOk

`func (o *AuthRequest) GetDuplicatePolicyOk() (*string, bool)`

GetDuplicatePolicyOk returns a tuple with the DuplicatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicatePolicy

`func (o *AuthRequest) SetDuplicatePolicy(v string)`

SetDuplicatePolicy sets DuplicatePolicy field to given value.

### HasDuplicatePolicy

`func (o *AuthRequest) HasDuplicatePolicy() bool`

HasDuplicatePolicy returns a boolean if a field has been set.

### GetEventManagement

`func (o *AuthRequest) GetEventManagement() EventDataModel`

GetEventManagement returns the EventManagement field if non-nil, zero value otherwise.

### GetEventManagementOk

`func (o *AuthRequest) GetEventManagementOk() (*EventDataModel, bool)`

GetEventManagementOk returns a tuple with the EventManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventManagement

`func (o *AuthRequest) SetEventManagement(v EventDataModel)`

SetEventManagement sets EventManagement field to given value.

### HasEventManagement

`func (o *AuthRequest) HasEventManagement() bool`

HasEventManagement returns a boolean if a field has been set.

### GetExpmonth

`func (o *AuthRequest) GetExpmonth() int32`

GetExpmonth returns the Expmonth field if non-nil, zero value otherwise.

### GetExpmonthOk

`func (o *AuthRequest) GetExpmonthOk() (*int32, bool)`

GetExpmonthOk returns a tuple with the Expmonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpmonth

`func (o *AuthRequest) SetExpmonth(v int32)`

SetExpmonth sets Expmonth field to given value.


### GetExpyear

`func (o *AuthRequest) GetExpyear() int32`

GetExpyear returns the Expyear field if non-nil, zero value otherwise.

### GetExpyearOk

`func (o *AuthRequest) GetExpyearOk() (*int32, bool)`

GetExpyearOk returns a tuple with the Expyear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpyear

`func (o *AuthRequest) SetExpyear(v int32)`

SetExpyear sets Expyear field to given value.


### GetExternalMpi

`func (o *AuthRequest) GetExternalMpi() ExternalMPI`

GetExternalMpi returns the ExternalMpi field if non-nil, zero value otherwise.

### GetExternalMpiOk

`func (o *AuthRequest) GetExternalMpiOk() (*ExternalMPI, bool)`

GetExternalMpiOk returns a tuple with the ExternalMpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalMpi

`func (o *AuthRequest) SetExternalMpi(v ExternalMPI)`

SetExternalMpi sets ExternalMpi field to given value.

### HasExternalMpi

`func (o *AuthRequest) HasExternalMpi() bool`

HasExternalMpi returns a boolean if a field has been set.

### GetIdentifier

`func (o *AuthRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *AuthRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *AuthRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetMatchAvsa

`func (o *AuthRequest) GetMatchAvsa() string`

GetMatchAvsa returns the MatchAvsa field if non-nil, zero value otherwise.

### GetMatchAvsaOk

`func (o *AuthRequest) GetMatchAvsaOk() (*string, bool)`

GetMatchAvsaOk returns a tuple with the MatchAvsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAvsa

`func (o *AuthRequest) SetMatchAvsa(v string)`

SetMatchAvsa sets MatchAvsa field to given value.

### HasMatchAvsa

`func (o *AuthRequest) HasMatchAvsa() bool`

HasMatchAvsa returns a boolean if a field has been set.

### GetMcc6012

`func (o *AuthRequest) GetMcc6012() MCC6012`

GetMcc6012 returns the Mcc6012 field if non-nil, zero value otherwise.

### GetMcc6012Ok

`func (o *AuthRequest) GetMcc6012Ok() (*MCC6012, bool)`

GetMcc6012Ok returns a tuple with the Mcc6012 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc6012

`func (o *AuthRequest) SetMcc6012(v MCC6012)`

SetMcc6012 sets Mcc6012 field to given value.

### HasMcc6012

`func (o *AuthRequest) HasMcc6012() bool`

HasMcc6012 returns a boolean if a field has been set.

### GetMerchantid

`func (o *AuthRequest) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *AuthRequest) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *AuthRequest) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.


### GetNameOnCard

`func (o *AuthRequest) GetNameOnCard() string`

GetNameOnCard returns the NameOnCard field if non-nil, zero value otherwise.

### GetNameOnCardOk

`func (o *AuthRequest) GetNameOnCardOk() (*string, bool)`

GetNameOnCardOk returns a tuple with the NameOnCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCard

`func (o *AuthRequest) SetNameOnCard(v string)`

SetNameOnCard sets NameOnCard field to given value.

### HasNameOnCard

`func (o *AuthRequest) HasNameOnCard() bool`

HasNameOnCard returns a boolean if a field has been set.

### GetShipTo

`func (o *AuthRequest) GetShipTo() ContactDetails`

GetShipTo returns the ShipTo field if non-nil, zero value otherwise.

### GetShipToOk

`func (o *AuthRequest) GetShipToOk() (*ContactDetails, bool)`

GetShipToOk returns a tuple with the ShipTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipTo

`func (o *AuthRequest) SetShipTo(v ContactDetails)`

SetShipTo sets ShipTo field to given value.

### HasShipTo

`func (o *AuthRequest) HasShipTo() bool`

HasShipTo returns a boolean if a field has been set.

### GetTag

`func (o *AuthRequest) GetTag() []string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AuthRequest) GetTagOk() (*[]string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AuthRequest) SetTag(v []string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AuthRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetThreedsecure

`func (o *AuthRequest) GetThreedsecure() ThreeDSecure`

GetThreedsecure returns the Threedsecure field if non-nil, zero value otherwise.

### GetThreedsecureOk

`func (o *AuthRequest) GetThreedsecureOk() (*ThreeDSecure, bool)`

GetThreedsecureOk returns a tuple with the Threedsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreedsecure

`func (o *AuthRequest) SetThreedsecure(v ThreeDSecure)`

SetThreedsecure sets Threedsecure field to given value.

### HasThreedsecure

`func (o *AuthRequest) HasThreedsecure() bool`

HasThreedsecure returns a boolean if a field has been set.

### GetTransInfo

`func (o *AuthRequest) GetTransInfo() string`

GetTransInfo returns the TransInfo field if non-nil, zero value otherwise.

### GetTransInfoOk

`func (o *AuthRequest) GetTransInfoOk() (*string, bool)`

GetTransInfoOk returns a tuple with the TransInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransInfo

`func (o *AuthRequest) SetTransInfo(v string)`

SetTransInfo sets TransInfo field to given value.

### HasTransInfo

`func (o *AuthRequest) HasTransInfo() bool`

HasTransInfo returns a boolean if a field has been set.

### GetTransType

`func (o *AuthRequest) GetTransType() string`

GetTransType returns the TransType field if non-nil, zero value otherwise.

### GetTransTypeOk

`func (o *AuthRequest) GetTransTypeOk() (*string, bool)`

GetTransTypeOk returns a tuple with the TransType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransType

`func (o *AuthRequest) SetTransType(v string)`

SetTransType sets TransType field to given value.

### HasTransType

`func (o *AuthRequest) HasTransType() bool`

HasTransType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


