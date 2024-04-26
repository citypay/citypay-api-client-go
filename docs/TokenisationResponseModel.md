# TokenisationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenResult** | Pointer to **string** | The result of any authentication using 3d_secure authorisation against ecommerce transactions. Values are:  &lt;table&gt; &lt;tr&gt; &lt;th&gt;Value&lt;/th&gt; &lt;th&gt;Description&lt;/th&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;Y&lt;/td&gt; &lt;td&gt;Authentication Successful. The Cardholder&#39;s password was successfully validated.&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;N&lt;/td&gt; &lt;td&gt;Authentication Failed. Customer failed or cancelled authentication, transaction denied.&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;A&lt;/td&gt; &lt;td&gt;Attempts Processing Performed Authentication could not be completed but a proof of authentication attempt (CAVV) was generated.&lt;/td&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt;U&lt;/td&gt; &lt;td&gt;Authentication Could Not Be Performed Authentication could not be completed, due to technical or other problem.&lt;/td&gt; &lt;/tr&gt; &lt;/table&gt;  | [optional] 
**BinCommercial** | Pointer to **bool** | Determines whether the bin range was found to be a commercial or business card. | [optional] 
**BinDebit** | Pointer to **bool** | Determines whether the bin range was found to be a debit card. If false the card was considered as a credit card. | [optional] 
**BinDescription** | Pointer to **string** | A description of the bin range found for the card. | [optional] 
**Eci** | Pointer to **string** | An Electronic Commerce Indicator (ECI) used to identify the result of authentication using 3DSecure.  | [optional] 
**Identifier** | Pointer to **string** | The identifier provided within the request. | [optional] 
**Maskedpan** | Pointer to **string** | A masked value of the card number used for processing displaying limited values that can be used on a receipt.  | [optional] 
**Scheme** | Pointer to **string** | The name of the card scheme of the transaction that processed the transaction such as Visa or MasterCard.  | [optional] 
**SigId** | Pointer to **string** | A Base58 encoded SHA-256 digest generated from the token value Base58 decoded and appended with the nonce value UTF-8 decoded. | [optional] 
**Token** | Pointer to **string** | The token used for presentment to authorisation later in the processing flow. | [optional] 

## Methods

### NewTokenisationResponseModel

`func NewTokenisationResponseModel() *TokenisationResponseModel`

NewTokenisationResponseModel instantiates a new TokenisationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenisationResponseModelWithDefaults

`func NewTokenisationResponseModelWithDefaults() *TokenisationResponseModel`

NewTokenisationResponseModelWithDefaults instantiates a new TokenisationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenResult

`func (o *TokenisationResponseModel) GetAuthenResult() string`

GetAuthenResult returns the AuthenResult field if non-nil, zero value otherwise.

### GetAuthenResultOk

`func (o *TokenisationResponseModel) GetAuthenResultOk() (*string, bool)`

GetAuthenResultOk returns a tuple with the AuthenResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenResult

`func (o *TokenisationResponseModel) SetAuthenResult(v string)`

SetAuthenResult sets AuthenResult field to given value.

### HasAuthenResult

`func (o *TokenisationResponseModel) HasAuthenResult() bool`

HasAuthenResult returns a boolean if a field has been set.

### GetBinCommercial

`func (o *TokenisationResponseModel) GetBinCommercial() bool`

GetBinCommercial returns the BinCommercial field if non-nil, zero value otherwise.

### GetBinCommercialOk

`func (o *TokenisationResponseModel) GetBinCommercialOk() (*bool, bool)`

GetBinCommercialOk returns a tuple with the BinCommercial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinCommercial

`func (o *TokenisationResponseModel) SetBinCommercial(v bool)`

SetBinCommercial sets BinCommercial field to given value.

### HasBinCommercial

`func (o *TokenisationResponseModel) HasBinCommercial() bool`

HasBinCommercial returns a boolean if a field has been set.

### GetBinDebit

`func (o *TokenisationResponseModel) GetBinDebit() bool`

GetBinDebit returns the BinDebit field if non-nil, zero value otherwise.

### GetBinDebitOk

`func (o *TokenisationResponseModel) GetBinDebitOk() (*bool, bool)`

GetBinDebitOk returns a tuple with the BinDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinDebit

`func (o *TokenisationResponseModel) SetBinDebit(v bool)`

SetBinDebit sets BinDebit field to given value.

### HasBinDebit

`func (o *TokenisationResponseModel) HasBinDebit() bool`

HasBinDebit returns a boolean if a field has been set.

### GetBinDescription

`func (o *TokenisationResponseModel) GetBinDescription() string`

GetBinDescription returns the BinDescription field if non-nil, zero value otherwise.

### GetBinDescriptionOk

`func (o *TokenisationResponseModel) GetBinDescriptionOk() (*string, bool)`

GetBinDescriptionOk returns a tuple with the BinDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinDescription

`func (o *TokenisationResponseModel) SetBinDescription(v string)`

SetBinDescription sets BinDescription field to given value.

### HasBinDescription

`func (o *TokenisationResponseModel) HasBinDescription() bool`

HasBinDescription returns a boolean if a field has been set.

### GetEci

`func (o *TokenisationResponseModel) GetEci() string`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *TokenisationResponseModel) GetEciOk() (*string, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *TokenisationResponseModel) SetEci(v string)`

SetEci sets Eci field to given value.

### HasEci

`func (o *TokenisationResponseModel) HasEci() bool`

HasEci returns a boolean if a field has been set.

### GetIdentifier

`func (o *TokenisationResponseModel) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *TokenisationResponseModel) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *TokenisationResponseModel) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *TokenisationResponseModel) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetMaskedpan

`func (o *TokenisationResponseModel) GetMaskedpan() string`

GetMaskedpan returns the Maskedpan field if non-nil, zero value otherwise.

### GetMaskedpanOk

`func (o *TokenisationResponseModel) GetMaskedpanOk() (*string, bool)`

GetMaskedpanOk returns a tuple with the Maskedpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedpan

`func (o *TokenisationResponseModel) SetMaskedpan(v string)`

SetMaskedpan sets Maskedpan field to given value.

### HasMaskedpan

`func (o *TokenisationResponseModel) HasMaskedpan() bool`

HasMaskedpan returns a boolean if a field has been set.

### GetScheme

`func (o *TokenisationResponseModel) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *TokenisationResponseModel) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *TokenisationResponseModel) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *TokenisationResponseModel) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetSigId

`func (o *TokenisationResponseModel) GetSigId() string`

GetSigId returns the SigId field if non-nil, zero value otherwise.

### GetSigIdOk

`func (o *TokenisationResponseModel) GetSigIdOk() (*string, bool)`

GetSigIdOk returns a tuple with the SigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigId

`func (o *TokenisationResponseModel) SetSigId(v string)`

SetSigId sets SigId field to given value.

### HasSigId

`func (o *TokenisationResponseModel) HasSigId() bool`

HasSigId returns a boolean if a field has been set.

### GetToken

`func (o *TokenisationResponseModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenisationResponseModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenisationResponseModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TokenisationResponseModel) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


