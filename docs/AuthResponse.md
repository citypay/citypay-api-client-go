# AuthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | The amount of the transaction processed. | [optional] 
**Atrn** | Pointer to **string** | A reference number provided by the acquirer for a transaction it can be used to cross reference transactions with an Acquirers reporting panel.  | [optional] 
**Atsd** | Pointer to **string** | Additional Transaction Security Data used for ecommerce transactions to decipher security capabilities and attempts against a transaction. | [optional] 
**Authcode** | Pointer to **string** | The authorisation code as returned by the card issuer or acquiring bank when a transaction has successfully   been authorised. Authorisation codes contain alphanumeric values. Whilst the code confirms authorisation it   should not be used to determine whether a transaction was successfully processed. For instance an auth code   may be returned when a transaction has been subsequently declined due to a CSC mismatch.  | [optional] 
**AuthenResult** | Pointer to **string** | The result of any authentication using 3d_secure authorisation against ecommerce transactions. Values are:  &lt;table&gt; &lt;tr&gt; &lt;th&gt;Value&lt;/th&gt; &lt;th&gt;Description&lt;/th&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;Y&lt;/td&gt; &lt;td&gt;Authentication Successful. The Cardholder&#39;s password was successfully validated.&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;N&lt;/td&gt; &lt;td&gt;Authentication Failed. Customer failed or cancelled authentication, transaction denied.&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;A&lt;/td&gt; &lt;td&gt;Attempts Processing Performed Authentication could not be completed but a proof of authentication attempt (CAVV) was generated.&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;U&lt;/td&gt; &lt;td&gt;Authentication Could Not Be Performed Authentication could not be completed, due to technical or other problem.&lt;/td&gt; &lt;/tr&gt; &lt;/table&gt;  | [optional] 
**Authorised** | Pointer to **bool** | A boolean definition that indicates that the transaction was authorised. It will return false if the transaction  was declined, rejected or cancelled due to CSC matching failures.  Attention should be referenced to the AuthResult and Response code for accurate determination of the result.  | [optional] 
**AvsResult** | Pointer to **string** | The AVS result codes determine the result of checking the AVS values within the Address Verification fraud system. If a transaction is declined due to the AVS code not matching, this value can help determine the reason for the decline.  &lt;table&gt; &lt;tr&gt; &lt;th&gt;Code&lt;/th&gt; &lt;th&gt;Description&lt;/th&gt; &lt;/tr&gt; &lt;tr&gt;&lt;td&gt;Y&lt;/td&gt;&lt;td&gt;Address and 5 digit post code match&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;M&lt;/td&gt;&lt;td&gt;Street address and Postal codes match for international transaction&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;U&lt;/td&gt;&lt;td&gt;No AVS data available from issuer auth system&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;A&lt;/td&gt;&lt;td&gt;Addres matches, post code does not&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;I&lt;/td&gt;&lt;td&gt;Address information verified for international transaction&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;Z&lt;/td&gt;&lt;td&gt;5 digit post code matches, Address does not&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;W&lt;/td&gt;&lt;td&gt;9 digit post code matches, Address does not&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;X&lt;/td&gt;&lt;td&gt;Postcode and address match&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;B&lt;/td&gt;&lt;td&gt;Postal code not verified due to incompatible formats&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;P&lt;/td&gt;&lt;td&gt;Postal codes match. Street address not verified due to to incompatible formats&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;E&lt;/td&gt;&lt;td&gt;AVS Error&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;C&lt;/td&gt;&lt;td&gt;Street address and Postal code not verified due to incompatible formats&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;D&lt;/td&gt;&lt;td&gt;Street address and postal codes match&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt; &lt;/td&gt;&lt;td&gt;No information&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;N&lt;/td&gt;&lt;td&gt;Neither postcode nor address match&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;R&lt;/td&gt;&lt;td&gt;Retry, System unavailble or Timed Out&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;S&lt;/td&gt;&lt;td&gt;AVS Service not supported by issuer or processor&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;G&lt;/td&gt;&lt;td&gt;Issuer does not participate in AVS&lt;/td&gt;&lt;/tr&gt; &lt;/table&gt;  | [optional] 
**BinCommercial** | Pointer to **bool** | Determines whether the bin range was found to be a commercial or business card. | [optional] 
**BinDebit** | Pointer to **bool** | Determines whether the bin range was found to be a debit card. If false the card was considered as a credit card. | [optional] 
**BinDescription** | Pointer to **string** | A description of the bin range found for the card. | [optional] 
**Cavv** | Pointer to **string** | The cardholder authentication verification value which can be returned for verification purposes of the authenticated  transaction for dispute realisation.  | [optional] 
**Context** | Pointer to **string** | The context which processed the transaction, can be used for support purposes to trace transactions. | [optional] 
**CscResult** | Pointer to **string** | The CSC result codes determine the result of checking the provided CSC value within the Card Security Code fraud system. If a transaction is declined due to the CSC code not matching, this value can help determine the reason for the decline.  &lt;table&gt; &lt;tr&gt; &lt;th&gt;Code&lt;/th&gt; &lt;th&gt;Description&lt;/th&gt; &lt;/tr&gt; &lt;tr&gt;&lt;td&gt; &lt;/td&gt;&lt;td&gt;No information&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;M&lt;/td&gt;&lt;td&gt;Card verification data matches&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;N&lt;/td&gt;&lt;td&gt;Card verification data was checked but did not match&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;P&lt;/td&gt;&lt;td&gt;Card verification was not processed&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;S&lt;/td&gt;&lt;td&gt;The card verification data should be on the card but the merchant indicates that it is not&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;U&lt;/td&gt;&lt;td&gt;The card issuer is not certified&lt;/td&gt;&lt;/tr&gt; &lt;/table&gt;  | [optional] 
**Currency** | Pointer to **string** | The currency the transaction was processed in. This is an &#x60;ISO4217&#x60; alpha currency value. | [optional] 
**Datetime** | Pointer to **time.Time** | The UTC date time of the transaction in ISO data time format.  | [optional] 
**Eci** | Pointer to **string** | An Electronic Commerce Indicator (ECI) used to identify the result of authentication using 3DSecure.  | [optional] 
**Identifier** | Pointer to **string** | The identifier provided within the request. | [optional] 
**Live** | Pointer to **bool** | Used to identify that a transaction was processed on a live authorisation platform. | [optional] 
**Maskedpan** | Pointer to **string** | A masked value of the card number used for processing displaying limited values that can be used on a receipt.  | [optional] 
**Merchantid** | **int32** | The merchant id that processed this transaction. | 
**Result** | **int32** | An integer result that indicates the outcome of the transaction. The Code value below maps to the result value  &lt;table&gt; &lt;tr&gt; &lt;th&gt;Code&lt;/th&gt; &lt;th&gt;Abbrev&lt;/th&gt; &lt;th&gt;Description&lt;/th&gt; &lt;/tr&gt; &lt;tr&gt;&lt;td&gt;0&lt;/td&gt;&lt;td&gt;Declined&lt;/td&gt;&lt;td&gt;Declined&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;1&lt;/td&gt;&lt;td&gt;Accepted&lt;/td&gt;&lt;td&gt;Accepted&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;2&lt;/td&gt;&lt;td&gt;Rejected&lt;/td&gt;&lt;td&gt;Rejected&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;3&lt;/td&gt;&lt;td&gt;Not Attempted&lt;/td&gt;&lt;td&gt;Not Attempted&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;4&lt;/td&gt;&lt;td&gt;Referred&lt;/td&gt;&lt;td&gt;Referred&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;5&lt;/td&gt;&lt;td&gt;PinRetry&lt;/td&gt;&lt;td&gt;Perform PIN Retry&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;6&lt;/td&gt;&lt;td&gt;ForSigVer&lt;/td&gt;&lt;td&gt;Force Signature Verification&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;7&lt;/td&gt;&lt;td&gt;Hold&lt;/td&gt;&lt;td&gt;Hold&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;8&lt;/td&gt;&lt;td&gt;SecErr&lt;/td&gt;&lt;td&gt;Security Error&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;9&lt;/td&gt;&lt;td&gt;CallAcq&lt;/td&gt;&lt;td&gt;Call Acquirer&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;10&lt;/td&gt;&lt;td&gt;DNH&lt;/td&gt;&lt;td&gt;Do Not Honour&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;11&lt;/td&gt;&lt;td&gt;RtnCrd&lt;/td&gt;&lt;td&gt;Retain Card&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;12&lt;/td&gt;&lt;td&gt;ExprdCrd&lt;/td&gt;&lt;td&gt;Expired Card&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;13&lt;/td&gt;&lt;td&gt;InvldCrd&lt;/td&gt;&lt;td&gt;Invalid Card No&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;14&lt;/td&gt;&lt;td&gt;PinExcd&lt;/td&gt;&lt;td&gt;Pin Tries Exceeded&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;15&lt;/td&gt;&lt;td&gt;PinInvld&lt;/td&gt;&lt;td&gt;Pin Invalid&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;16&lt;/td&gt;&lt;td&gt;AuthReq&lt;/td&gt;&lt;td&gt;Authentication Required&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;17&lt;/td&gt;&lt;td&gt;AuthenFail&lt;/td&gt;&lt;td&gt;Authentication Failed&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;18&lt;/td&gt;&lt;td&gt;Verified&lt;/td&gt;&lt;td&gt;Card Verified&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;19&lt;/td&gt;&lt;td&gt;Cancelled&lt;/td&gt;&lt;td&gt;Cancelled&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;20&lt;/td&gt;&lt;td&gt;Un&lt;/td&gt;&lt;td&gt;Unknown&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;21&lt;/td&gt;&lt;td&gt;Challenged&lt;/td&gt;&lt;td&gt;Challenged&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;22&lt;/td&gt;&lt;td&gt;Decoupled&lt;/td&gt;&lt;td&gt;Decoupled&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;23&lt;/td&gt;&lt;td&gt;Denied&lt;/td&gt;&lt;td&gt;Permission Denied&lt;/td&gt;&lt;/tr&gt; &lt;/table&gt;  | 
**ResultCode** | **string** | The result code as defined in the Response Codes Reference for example 000 is an accepted live transaction whilst 001 is an accepted test transaction. Result codes identify the source of success and failure.  Codes may start with an alpha character i.e. C001 indicating a type of error such as a card validation error.  | 
**ResultMessage** | **string** | The message regarding the result which provides further narrative to the result code.  | 
**Scheme** | Pointer to **string** | The name of the card scheme of the transaction that processed the transaction such as Visa or MasterCard.  | [optional] 
**SchemeId** | Pointer to **string** | The name of the card scheme of the transaction such as VI or MC.  | [optional] 
**SchemeLogo** | Pointer to **string** | A url containing a logo of the card scheme.  | [optional] 
**Sha256** | Pointer to **string** | A SHA256 digest value of the transaction used to validate the response data The digest is calculated by concatenating   * authcode   * amount   * response_code   * merchant_id   * trans_no   * identifier   * licence_key - which is not provided in the response.  | [optional] 
**TransStatus** | Pointer to **string** | Used to identify the status of a transaction. The status is used to track a transaction through its life cycle.  &lt;table&gt; &lt;tr&gt; &lt;th&gt;Id&lt;/th&gt; &lt;th&gt;Description&lt;/th&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;O&lt;/td&gt; &lt;td&gt;Transaction is open for settlement&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;A&lt;/td&gt; &lt;td&gt;Transaction is assigned for settlement and can no longer be voided&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;S&lt;/td&gt; &lt;td&gt;Transaction has been settled&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;D&lt;/td&gt; &lt;td&gt;Transaction has been declined&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;R&lt;/td&gt; &lt;td&gt;Transaction has been rejected&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;P&lt;/td&gt; &lt;td&gt;Transaction has been authorised only and awaiting a capture. Used in pre-auth situations&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;C&lt;/td&gt; &lt;td&gt;Transaction has been cancelled&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;E&lt;/td&gt; &lt;td&gt;Transaction has expired&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;I&lt;/td&gt; &lt;td&gt;Transaction has been initialised but no action was able to be carried out&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;H&lt;/td&gt; &lt;td&gt;Transaction is awaiting authorisation&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;.&lt;/td&gt; &lt;td&gt;Transaction is on hold&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;V&lt;/td&gt; &lt;td&gt;Transaction has been verified&lt;/td&gt; &lt;/tr&gt; &lt;/table&gt;  | [optional] 
**Transno** | Pointer to **int32** | The resulting transaction number, ordered incrementally from 1 for every merchant_id. The value will default to less than 1 for transactions that do not have a transaction number issued.  | [optional] 

## Methods

### NewAuthResponse

`func NewAuthResponse(merchantid int32, result int32, resultCode string, resultMessage string, ) *AuthResponse`

NewAuthResponse instantiates a new AuthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthResponseWithDefaults

`func NewAuthResponseWithDefaults() *AuthResponse`

NewAuthResponseWithDefaults instantiates a new AuthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AuthResponse) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AuthResponse) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AuthResponse) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AuthResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAtrn

`func (o *AuthResponse) GetAtrn() string`

GetAtrn returns the Atrn field if non-nil, zero value otherwise.

### GetAtrnOk

`func (o *AuthResponse) GetAtrnOk() (*string, bool)`

GetAtrnOk returns a tuple with the Atrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtrn

`func (o *AuthResponse) SetAtrn(v string)`

SetAtrn sets Atrn field to given value.

### HasAtrn

`func (o *AuthResponse) HasAtrn() bool`

HasAtrn returns a boolean if a field has been set.

### GetAtsd

`func (o *AuthResponse) GetAtsd() string`

GetAtsd returns the Atsd field if non-nil, zero value otherwise.

### GetAtsdOk

`func (o *AuthResponse) GetAtsdOk() (*string, bool)`

GetAtsdOk returns a tuple with the Atsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtsd

`func (o *AuthResponse) SetAtsd(v string)`

SetAtsd sets Atsd field to given value.

### HasAtsd

`func (o *AuthResponse) HasAtsd() bool`

HasAtsd returns a boolean if a field has been set.

### GetAuthcode

`func (o *AuthResponse) GetAuthcode() string`

GetAuthcode returns the Authcode field if non-nil, zero value otherwise.

### GetAuthcodeOk

`func (o *AuthResponse) GetAuthcodeOk() (*string, bool)`

GetAuthcodeOk returns a tuple with the Authcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthcode

`func (o *AuthResponse) SetAuthcode(v string)`

SetAuthcode sets Authcode field to given value.

### HasAuthcode

`func (o *AuthResponse) HasAuthcode() bool`

HasAuthcode returns a boolean if a field has been set.

### GetAuthenResult

`func (o *AuthResponse) GetAuthenResult() string`

GetAuthenResult returns the AuthenResult field if non-nil, zero value otherwise.

### GetAuthenResultOk

`func (o *AuthResponse) GetAuthenResultOk() (*string, bool)`

GetAuthenResultOk returns a tuple with the AuthenResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenResult

`func (o *AuthResponse) SetAuthenResult(v string)`

SetAuthenResult sets AuthenResult field to given value.

### HasAuthenResult

`func (o *AuthResponse) HasAuthenResult() bool`

HasAuthenResult returns a boolean if a field has been set.

### GetAuthorised

`func (o *AuthResponse) GetAuthorised() bool`

GetAuthorised returns the Authorised field if non-nil, zero value otherwise.

### GetAuthorisedOk

`func (o *AuthResponse) GetAuthorisedOk() (*bool, bool)`

GetAuthorisedOk returns a tuple with the Authorised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorised

`func (o *AuthResponse) SetAuthorised(v bool)`

SetAuthorised sets Authorised field to given value.

### HasAuthorised

`func (o *AuthResponse) HasAuthorised() bool`

HasAuthorised returns a boolean if a field has been set.

### GetAvsResult

`func (o *AuthResponse) GetAvsResult() string`

GetAvsResult returns the AvsResult field if non-nil, zero value otherwise.

### GetAvsResultOk

`func (o *AuthResponse) GetAvsResultOk() (*string, bool)`

GetAvsResultOk returns a tuple with the AvsResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsResult

`func (o *AuthResponse) SetAvsResult(v string)`

SetAvsResult sets AvsResult field to given value.

### HasAvsResult

`func (o *AuthResponse) HasAvsResult() bool`

HasAvsResult returns a boolean if a field has been set.

### GetBinCommercial

`func (o *AuthResponse) GetBinCommercial() bool`

GetBinCommercial returns the BinCommercial field if non-nil, zero value otherwise.

### GetBinCommercialOk

`func (o *AuthResponse) GetBinCommercialOk() (*bool, bool)`

GetBinCommercialOk returns a tuple with the BinCommercial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinCommercial

`func (o *AuthResponse) SetBinCommercial(v bool)`

SetBinCommercial sets BinCommercial field to given value.

### HasBinCommercial

`func (o *AuthResponse) HasBinCommercial() bool`

HasBinCommercial returns a boolean if a field has been set.

### GetBinDebit

`func (o *AuthResponse) GetBinDebit() bool`

GetBinDebit returns the BinDebit field if non-nil, zero value otherwise.

### GetBinDebitOk

`func (o *AuthResponse) GetBinDebitOk() (*bool, bool)`

GetBinDebitOk returns a tuple with the BinDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinDebit

`func (o *AuthResponse) SetBinDebit(v bool)`

SetBinDebit sets BinDebit field to given value.

### HasBinDebit

`func (o *AuthResponse) HasBinDebit() bool`

HasBinDebit returns a boolean if a field has been set.

### GetBinDescription

`func (o *AuthResponse) GetBinDescription() string`

GetBinDescription returns the BinDescription field if non-nil, zero value otherwise.

### GetBinDescriptionOk

`func (o *AuthResponse) GetBinDescriptionOk() (*string, bool)`

GetBinDescriptionOk returns a tuple with the BinDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinDescription

`func (o *AuthResponse) SetBinDescription(v string)`

SetBinDescription sets BinDescription field to given value.

### HasBinDescription

`func (o *AuthResponse) HasBinDescription() bool`

HasBinDescription returns a boolean if a field has been set.

### GetCavv

`func (o *AuthResponse) GetCavv() string`

GetCavv returns the Cavv field if non-nil, zero value otherwise.

### GetCavvOk

`func (o *AuthResponse) GetCavvOk() (*string, bool)`

GetCavvOk returns a tuple with the Cavv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavv

`func (o *AuthResponse) SetCavv(v string)`

SetCavv sets Cavv field to given value.

### HasCavv

`func (o *AuthResponse) HasCavv() bool`

HasCavv returns a boolean if a field has been set.

### GetContext

`func (o *AuthResponse) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AuthResponse) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AuthResponse) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *AuthResponse) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCscResult

`func (o *AuthResponse) GetCscResult() string`

GetCscResult returns the CscResult field if non-nil, zero value otherwise.

### GetCscResultOk

`func (o *AuthResponse) GetCscResultOk() (*string, bool)`

GetCscResultOk returns a tuple with the CscResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCscResult

`func (o *AuthResponse) SetCscResult(v string)`

SetCscResult sets CscResult field to given value.

### HasCscResult

`func (o *AuthResponse) HasCscResult() bool`

HasCscResult returns a boolean if a field has been set.

### GetCurrency

`func (o *AuthResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AuthResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AuthResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AuthResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDatetime

`func (o *AuthResponse) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *AuthResponse) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *AuthResponse) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *AuthResponse) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetEci

`func (o *AuthResponse) GetEci() string`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *AuthResponse) GetEciOk() (*string, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *AuthResponse) SetEci(v string)`

SetEci sets Eci field to given value.

### HasEci

`func (o *AuthResponse) HasEci() bool`

HasEci returns a boolean if a field has been set.

### GetIdentifier

`func (o *AuthResponse) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *AuthResponse) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *AuthResponse) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *AuthResponse) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetLive

`func (o *AuthResponse) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *AuthResponse) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *AuthResponse) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *AuthResponse) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetMaskedpan

`func (o *AuthResponse) GetMaskedpan() string`

GetMaskedpan returns the Maskedpan field if non-nil, zero value otherwise.

### GetMaskedpanOk

`func (o *AuthResponse) GetMaskedpanOk() (*string, bool)`

GetMaskedpanOk returns a tuple with the Maskedpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedpan

`func (o *AuthResponse) SetMaskedpan(v string)`

SetMaskedpan sets Maskedpan field to given value.

### HasMaskedpan

`func (o *AuthResponse) HasMaskedpan() bool`

HasMaskedpan returns a boolean if a field has been set.

### GetMerchantid

`func (o *AuthResponse) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *AuthResponse) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *AuthResponse) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.


### GetResult

`func (o *AuthResponse) GetResult() int32`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AuthResponse) GetResultOk() (*int32, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AuthResponse) SetResult(v int32)`

SetResult sets Result field to given value.


### GetResultCode

`func (o *AuthResponse) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *AuthResponse) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *AuthResponse) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.


### GetResultMessage

`func (o *AuthResponse) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *AuthResponse) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *AuthResponse) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.


### GetScheme

`func (o *AuthResponse) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *AuthResponse) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *AuthResponse) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *AuthResponse) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetSchemeId

`func (o *AuthResponse) GetSchemeId() string`

GetSchemeId returns the SchemeId field if non-nil, zero value otherwise.

### GetSchemeIdOk

`func (o *AuthResponse) GetSchemeIdOk() (*string, bool)`

GetSchemeIdOk returns a tuple with the SchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeId

`func (o *AuthResponse) SetSchemeId(v string)`

SetSchemeId sets SchemeId field to given value.

### HasSchemeId

`func (o *AuthResponse) HasSchemeId() bool`

HasSchemeId returns a boolean if a field has been set.

### GetSchemeLogo

`func (o *AuthResponse) GetSchemeLogo() string`

GetSchemeLogo returns the SchemeLogo field if non-nil, zero value otherwise.

### GetSchemeLogoOk

`func (o *AuthResponse) GetSchemeLogoOk() (*string, bool)`

GetSchemeLogoOk returns a tuple with the SchemeLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeLogo

`func (o *AuthResponse) SetSchemeLogo(v string)`

SetSchemeLogo sets SchemeLogo field to given value.

### HasSchemeLogo

`func (o *AuthResponse) HasSchemeLogo() bool`

HasSchemeLogo returns a boolean if a field has been set.

### GetSha256

`func (o *AuthResponse) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *AuthResponse) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *AuthResponse) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *AuthResponse) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetTransStatus

`func (o *AuthResponse) GetTransStatus() string`

GetTransStatus returns the TransStatus field if non-nil, zero value otherwise.

### GetTransStatusOk

`func (o *AuthResponse) GetTransStatusOk() (*string, bool)`

GetTransStatusOk returns a tuple with the TransStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransStatus

`func (o *AuthResponse) SetTransStatus(v string)`

SetTransStatus sets TransStatus field to given value.

### HasTransStatus

`func (o *AuthResponse) HasTransStatus() bool`

HasTransStatus returns a boolean if a field has been set.

### GetTransno

`func (o *AuthResponse) GetTransno() int32`

GetTransno returns the Transno field if non-nil, zero value otherwise.

### GetTransnoOk

`func (o *AuthResponse) GetTransnoOk() (*int32, bool)`

GetTransnoOk returns a tuple with the Transno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransno

`func (o *AuthResponse) SetTransno(v int32)`

SetTransno sets Transno field to given value.

### HasTransno

`func (o *AuthResponse) HasTransno() bool`

HasTransno returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


