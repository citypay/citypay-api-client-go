# DirectPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The amount to authorise in the lowest unit of currency with a variable length to a maximum of 12 digits.  No decimal points are to be included and no divisional characters such as 1,024.  The amount should be the total amount required for the transaction.  For example with GBP £1,021.95 the amount value is 102195.  | 
**AvsPostcodePolicy** | Pointer to **string** | A policy value which determines whether an AVS postcode policy is enforced or bypassed.  Values are:   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be rejected if the AVS postcode numeric value does not match.   &#x60;2&#x60; to bypass. Transactions that are bypassed will be allowed through even if the postcode did not match.   &#x60;3&#x60; to ignore. Transactions that are ignored will bypass the result and not send postcode details for authorisation.  | [optional] 
**BillTo** | Pointer to [**ContactDetails**](ContactDetails.md) |  | [optional] 
**Cardnumber** | **string** | The card number (PAN) with a variable length to a maximum of 21 digits in numerical form. Any non numeric characters will be stripped out of the card number, this includes whitespace or separators internal of the provided value.  The card number must be treated as sensitive data. We only provide an obfuscated value in logging and reporting.  The plaintext value is encrypted in our database using AES 256 GMC bit encryption for settlement or refund purposes.  When providing the card number to our gateway through the authorisation API you will be handling the card data on your application. This will require further PCI controls to be in place and this value must never be stored.  | 
**Csc** | Pointer to **string** | The Card Security Code (CSC) (also known as CV2/CVV2) is normally found on the back of the card (American Express has it on the front). The value helps to identify possession of the card as it is not available within the chip or magnetic swipe.  When forwarding the CSC, please ensure the value is a string as some values start with 0 and this will be stripped out by any integer parsing.  The CSC number aids fraud prevention in Mail Order and Internet payments.  Business rules are available on your account to identify whether to accept or decline transactions based on mismatched results of the CSC.  The Payment Card Industry (PCI) requires that at no stage of a transaction should the CSC be stored.  This applies to all entities handling card data.  It should also not be used in any hashing process.  CityPay do not store the value and have no method of retrieving the value once the transaction has been processed. For this reason, duplicate checking is unable to determine the CSC in its duplication check algorithm.  | [optional] 
**CscPolicy** | Pointer to **string** | A policy value which determines whether a CSC policy is enforced or bypassed.  Values are:   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be rejected if the CSC value does not match.   &#x60;2&#x60; to bypass. Transactions that are bypassed will be allowed through even if the CSC did not match.   &#x60;3&#x60; to ignore. Transactions that are ignored will bypass the result and not send the CSC details for authorisation.  | [optional] 
**Currency** | Pointer to **string** | The processing currency for the transaction. Will default to the merchant account currency. | [optional] 
**DuplicatePolicy** | Pointer to **string** | A policy value which determines whether a duplication policy is enforced or bypassed. A duplication check has a window of time set against your account within which it can action. If a previous transaction with matching values occurred within the window, any subsequent transaction will result in a T001 result.  Values are   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be checked for duplication within the duplication window.   &#x60;2&#x60; to bypass. Transactions that are bypassed will not be checked for duplication within the duplication window.   &#x60;3&#x60; to ignore. Transactions that are ignored will have the same affect as bypass.  | [optional] 
**Expmonth** | **int32** | The month of expiry of the card. The month value should be a numerical value between 1 and 12.  | 
**Expyear** | **int32** | The year of expiry of the card.  | 
**Identifier** | **string** | The identifier of the transaction to process. The value should be a valid reference and may be used to perform  post processing actions and to aid in reconciliation of transactions.  The value should be a valid printable string with ASCII character ranges from 0x32 to 0x127.  The identifier is recommended to be distinct for each transaction such as a [random unique identifier](https://en.wikipedia.org/wiki/Universally_unique_identifier) this will aid in ensuring each transaction is identifiable.  When transactions are processed they are also checked for duplicate requests. Changing the identifier on a subsequent request will ensure that a transaction is considered as different.  | 
**Mac** | **string** | A message authentication code ensures the data is authentic and that the intended amount has not been tampered with. The mac value is generated using a hash-based mac value. The following algorithm is used. - A key (k) is derived from your licence key - A value (v) is produced by concatenating the nonce, amount value and identifier, such as a purchase   with nonce &#x60;0123456789ABCDEF&#x60; an amount of £275.95 and an identifier of OD-12345678 would become   &#x60;0123456789ABCDEF27595OD-12345678&#x60; and extracting the UTF-8 byte values - The result from HMAC_SHA256(k, v) is hex-encoded (upper-case) - For instance, a licence key of &#x60;LK123456789&#x60;, a nonce of &#x60;0123456789ABCDEF&#x60;, an amount of &#x60;27595&#x60; and an identifier of &#x60;OD-12345678&#x60;  would generate a MAC of &#x60;163DBAB194D743866A9BCC7FC9C8A88FCD99C6BBBF08D619291212D1B91EE12E&#x60;.  | 
**MatchAvsa** | Pointer to **string** | A policy value which determines whether an AVS address policy is enforced, bypassed or ignored.  Values are:   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions that are enforced will be rejected if the AVS address numeric value does not match.   &#x60;2&#x60; to bypass. Transactions that are bypassed will be allowed through even if the address did not match.   &#x60;3&#x60; to ignore. Transactions that are ignored will bypass the result and not send address numeric details for authorisation.  | [optional] 
**NameOnCard** | Pointer to **string** | The card holder name as appears on the card such as MR N E BODY. Required for some acquirers.  | [optional] 
**Nonce** | Pointer to **string** | A random value Hex string (uppercase) which is provided to the API to perform a digest. The value will be used in any digest function.  | [optional] 
**RedirectFailure** | Pointer to **string** | The URL used to redirect back to your site when a transaction has been rejected or declined. Required if a url-encoded request.  | [optional] 
**RedirectSuccess** | Pointer to **string** | The URL used to redirect back to your site when a transaction has been tokenised or authorised. Required if a url-encoded request.  | [optional] 
**ShipTo** | Pointer to [**ContactDetails**](ContactDetails.md) |  | [optional] 
**Tag** | Pointer to **[]string** |  | [optional] 
**Threedsecure** | Pointer to [**ThreeDSecure**](ThreeDSecure.md) |  | [optional] 
**TransInfo** | Pointer to **string** | Further information that can be added to the transaction will display in reporting. Can be used for flexible values such as operator id. | [optional] 
**TransType** | Pointer to **string** | The type of transaction being submitted. Normally this value is not required and your account manager may request that you set this field. | [optional] 

## Methods

### NewDirectPostRequest

`func NewDirectPostRequest(amount int32, cardnumber string, expmonth int32, expyear int32, identifier string, mac string, ) *DirectPostRequest`

NewDirectPostRequest instantiates a new DirectPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectPostRequestWithDefaults

`func NewDirectPostRequestWithDefaults() *DirectPostRequest`

NewDirectPostRequestWithDefaults instantiates a new DirectPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DirectPostRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DirectPostRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DirectPostRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAvsPostcodePolicy

`func (o *DirectPostRequest) GetAvsPostcodePolicy() string`

GetAvsPostcodePolicy returns the AvsPostcodePolicy field if non-nil, zero value otherwise.

### GetAvsPostcodePolicyOk

`func (o *DirectPostRequest) GetAvsPostcodePolicyOk() (*string, bool)`

GetAvsPostcodePolicyOk returns a tuple with the AvsPostcodePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsPostcodePolicy

`func (o *DirectPostRequest) SetAvsPostcodePolicy(v string)`

SetAvsPostcodePolicy sets AvsPostcodePolicy field to given value.

### HasAvsPostcodePolicy

`func (o *DirectPostRequest) HasAvsPostcodePolicy() bool`

HasAvsPostcodePolicy returns a boolean if a field has been set.

### GetBillTo

`func (o *DirectPostRequest) GetBillTo() ContactDetails`

GetBillTo returns the BillTo field if non-nil, zero value otherwise.

### GetBillToOk

`func (o *DirectPostRequest) GetBillToOk() (*ContactDetails, bool)`

GetBillToOk returns a tuple with the BillTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTo

`func (o *DirectPostRequest) SetBillTo(v ContactDetails)`

SetBillTo sets BillTo field to given value.

### HasBillTo

`func (o *DirectPostRequest) HasBillTo() bool`

HasBillTo returns a boolean if a field has been set.

### GetCardnumber

`func (o *DirectPostRequest) GetCardnumber() string`

GetCardnumber returns the Cardnumber field if non-nil, zero value otherwise.

### GetCardnumberOk

`func (o *DirectPostRequest) GetCardnumberOk() (*string, bool)`

GetCardnumberOk returns a tuple with the Cardnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardnumber

`func (o *DirectPostRequest) SetCardnumber(v string)`

SetCardnumber sets Cardnumber field to given value.


### GetCsc

`func (o *DirectPostRequest) GetCsc() string`

GetCsc returns the Csc field if non-nil, zero value otherwise.

### GetCscOk

`func (o *DirectPostRequest) GetCscOk() (*string, bool)`

GetCscOk returns a tuple with the Csc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsc

`func (o *DirectPostRequest) SetCsc(v string)`

SetCsc sets Csc field to given value.

### HasCsc

`func (o *DirectPostRequest) HasCsc() bool`

HasCsc returns a boolean if a field has been set.

### GetCscPolicy

`func (o *DirectPostRequest) GetCscPolicy() string`

GetCscPolicy returns the CscPolicy field if non-nil, zero value otherwise.

### GetCscPolicyOk

`func (o *DirectPostRequest) GetCscPolicyOk() (*string, bool)`

GetCscPolicyOk returns a tuple with the CscPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCscPolicy

`func (o *DirectPostRequest) SetCscPolicy(v string)`

SetCscPolicy sets CscPolicy field to given value.

### HasCscPolicy

`func (o *DirectPostRequest) HasCscPolicy() bool`

HasCscPolicy returns a boolean if a field has been set.

### GetCurrency

`func (o *DirectPostRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DirectPostRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DirectPostRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DirectPostRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDuplicatePolicy

`func (o *DirectPostRequest) GetDuplicatePolicy() string`

GetDuplicatePolicy returns the DuplicatePolicy field if non-nil, zero value otherwise.

### GetDuplicatePolicyOk

`func (o *DirectPostRequest) GetDuplicatePolicyOk() (*string, bool)`

GetDuplicatePolicyOk returns a tuple with the DuplicatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicatePolicy

`func (o *DirectPostRequest) SetDuplicatePolicy(v string)`

SetDuplicatePolicy sets DuplicatePolicy field to given value.

### HasDuplicatePolicy

`func (o *DirectPostRequest) HasDuplicatePolicy() bool`

HasDuplicatePolicy returns a boolean if a field has been set.

### GetExpmonth

`func (o *DirectPostRequest) GetExpmonth() int32`

GetExpmonth returns the Expmonth field if non-nil, zero value otherwise.

### GetExpmonthOk

`func (o *DirectPostRequest) GetExpmonthOk() (*int32, bool)`

GetExpmonthOk returns a tuple with the Expmonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpmonth

`func (o *DirectPostRequest) SetExpmonth(v int32)`

SetExpmonth sets Expmonth field to given value.


### GetExpyear

`func (o *DirectPostRequest) GetExpyear() int32`

GetExpyear returns the Expyear field if non-nil, zero value otherwise.

### GetExpyearOk

`func (o *DirectPostRequest) GetExpyearOk() (*int32, bool)`

GetExpyearOk returns a tuple with the Expyear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpyear

`func (o *DirectPostRequest) SetExpyear(v int32)`

SetExpyear sets Expyear field to given value.


### GetIdentifier

`func (o *DirectPostRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *DirectPostRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *DirectPostRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetMac

`func (o *DirectPostRequest) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DirectPostRequest) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DirectPostRequest) SetMac(v string)`

SetMac sets Mac field to given value.


### GetMatchAvsa

`func (o *DirectPostRequest) GetMatchAvsa() string`

GetMatchAvsa returns the MatchAvsa field if non-nil, zero value otherwise.

### GetMatchAvsaOk

`func (o *DirectPostRequest) GetMatchAvsaOk() (*string, bool)`

GetMatchAvsaOk returns a tuple with the MatchAvsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchAvsa

`func (o *DirectPostRequest) SetMatchAvsa(v string)`

SetMatchAvsa sets MatchAvsa field to given value.

### HasMatchAvsa

`func (o *DirectPostRequest) HasMatchAvsa() bool`

HasMatchAvsa returns a boolean if a field has been set.

### GetNameOnCard

`func (o *DirectPostRequest) GetNameOnCard() string`

GetNameOnCard returns the NameOnCard field if non-nil, zero value otherwise.

### GetNameOnCardOk

`func (o *DirectPostRequest) GetNameOnCardOk() (*string, bool)`

GetNameOnCardOk returns a tuple with the NameOnCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCard

`func (o *DirectPostRequest) SetNameOnCard(v string)`

SetNameOnCard sets NameOnCard field to given value.

### HasNameOnCard

`func (o *DirectPostRequest) HasNameOnCard() bool`

HasNameOnCard returns a boolean if a field has been set.

### GetNonce

`func (o *DirectPostRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *DirectPostRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *DirectPostRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *DirectPostRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetRedirectFailure

`func (o *DirectPostRequest) GetRedirectFailure() string`

GetRedirectFailure returns the RedirectFailure field if non-nil, zero value otherwise.

### GetRedirectFailureOk

`func (o *DirectPostRequest) GetRedirectFailureOk() (*string, bool)`

GetRedirectFailureOk returns a tuple with the RedirectFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectFailure

`func (o *DirectPostRequest) SetRedirectFailure(v string)`

SetRedirectFailure sets RedirectFailure field to given value.

### HasRedirectFailure

`func (o *DirectPostRequest) HasRedirectFailure() bool`

HasRedirectFailure returns a boolean if a field has been set.

### GetRedirectSuccess

`func (o *DirectPostRequest) GetRedirectSuccess() string`

GetRedirectSuccess returns the RedirectSuccess field if non-nil, zero value otherwise.

### GetRedirectSuccessOk

`func (o *DirectPostRequest) GetRedirectSuccessOk() (*string, bool)`

GetRedirectSuccessOk returns a tuple with the RedirectSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectSuccess

`func (o *DirectPostRequest) SetRedirectSuccess(v string)`

SetRedirectSuccess sets RedirectSuccess field to given value.

### HasRedirectSuccess

`func (o *DirectPostRequest) HasRedirectSuccess() bool`

HasRedirectSuccess returns a boolean if a field has been set.

### GetShipTo

`func (o *DirectPostRequest) GetShipTo() ContactDetails`

GetShipTo returns the ShipTo field if non-nil, zero value otherwise.

### GetShipToOk

`func (o *DirectPostRequest) GetShipToOk() (*ContactDetails, bool)`

GetShipToOk returns a tuple with the ShipTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipTo

`func (o *DirectPostRequest) SetShipTo(v ContactDetails)`

SetShipTo sets ShipTo field to given value.

### HasShipTo

`func (o *DirectPostRequest) HasShipTo() bool`

HasShipTo returns a boolean if a field has been set.

### GetTag

`func (o *DirectPostRequest) GetTag() []string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *DirectPostRequest) GetTagOk() (*[]string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *DirectPostRequest) SetTag(v []string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *DirectPostRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetThreedsecure

`func (o *DirectPostRequest) GetThreedsecure() ThreeDSecure`

GetThreedsecure returns the Threedsecure field if non-nil, zero value otherwise.

### GetThreedsecureOk

`func (o *DirectPostRequest) GetThreedsecureOk() (*ThreeDSecure, bool)`

GetThreedsecureOk returns a tuple with the Threedsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreedsecure

`func (o *DirectPostRequest) SetThreedsecure(v ThreeDSecure)`

SetThreedsecure sets Threedsecure field to given value.

### HasThreedsecure

`func (o *DirectPostRequest) HasThreedsecure() bool`

HasThreedsecure returns a boolean if a field has been set.

### GetTransInfo

`func (o *DirectPostRequest) GetTransInfo() string`

GetTransInfo returns the TransInfo field if non-nil, zero value otherwise.

### GetTransInfoOk

`func (o *DirectPostRequest) GetTransInfoOk() (*string, bool)`

GetTransInfoOk returns a tuple with the TransInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransInfo

`func (o *DirectPostRequest) SetTransInfo(v string)`

SetTransInfo sets TransInfo field to given value.

### HasTransInfo

`func (o *DirectPostRequest) HasTransInfo() bool`

HasTransInfo returns a boolean if a field has been set.

### GetTransType

`func (o *DirectPostRequest) GetTransType() string`

GetTransType returns the TransType field if non-nil, zero value otherwise.

### GetTransTypeOk

`func (o *DirectPostRequest) GetTransTypeOk() (*string, bool)`

GetTransTypeOk returns a tuple with the TransType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransType

`func (o *DirectPostRequest) SetTransType(v string)`

SetTransType sets TransType field to given value.

### HasTransType

`func (o *DirectPostRequest) HasTransType() bool`

HasTransType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


