# BatchTransactionResultModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The card holder account id used for the transaction. | 
**Amount** | Pointer to **int32** | The amount of the transaction processed. | [optional] 
**Authcode** | Pointer to **string** | The authorisation code of a successful transaction. | [optional] 
**Datetime** | Pointer to **time.Time** | The datetime that the transaction was processed. | [optional] 
**Identifier** | **string** | The identifier of the transaction. | 
**Maskedpan** | Pointer to **string** | A masked value of the card number used for processing displaying limited values that can be used on a receipt.  | [optional] 
**Merchantid** | **int32** | The merchant id of the transaction. | 
**Message** | **string** | A response message pertaining to the transaction. | 
**Result** | **int32** | An integer result that indicates the outcome of the transaction. The Code value below maps to the result value  &lt;table&gt; &lt;tr&gt; &lt;th&gt;Code&lt;/th&gt; &lt;th&gt;Abbrev&lt;/th&gt; &lt;th&gt;Description&lt;/th&gt; &lt;/tr&gt; &lt;tr&gt;&lt;td&gt;0&lt;/td&gt;&lt;td&gt;Declined&lt;/td&gt;&lt;td&gt;Declined&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;1&lt;/td&gt;&lt;td&gt;Accepted&lt;/td&gt;&lt;td&gt;Accepted&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;2&lt;/td&gt;&lt;td&gt;Rejected&lt;/td&gt;&lt;td&gt;Rejected&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;3&lt;/td&gt;&lt;td&gt;Not Attempted&lt;/td&gt;&lt;td&gt;Not Attempted&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;4&lt;/td&gt;&lt;td&gt;Referred&lt;/td&gt;&lt;td&gt;Referred&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;5&lt;/td&gt;&lt;td&gt;PinRetry&lt;/td&gt;&lt;td&gt;Perform PIN Retry&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;6&lt;/td&gt;&lt;td&gt;ForSigVer&lt;/td&gt;&lt;td&gt;Force Signature Verification&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;7&lt;/td&gt;&lt;td&gt;Hold&lt;/td&gt;&lt;td&gt;Hold&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;8&lt;/td&gt;&lt;td&gt;SecErr&lt;/td&gt;&lt;td&gt;Security Error&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;9&lt;/td&gt;&lt;td&gt;CallAcq&lt;/td&gt;&lt;td&gt;Call Acquirer&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;10&lt;/td&gt;&lt;td&gt;DNH&lt;/td&gt;&lt;td&gt;Do Not Honour&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;11&lt;/td&gt;&lt;td&gt;RtnCrd&lt;/td&gt;&lt;td&gt;Retain Card&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;12&lt;/td&gt;&lt;td&gt;ExprdCrd&lt;/td&gt;&lt;td&gt;Expired Card&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;13&lt;/td&gt;&lt;td&gt;InvldCrd&lt;/td&gt;&lt;td&gt;Invalid Card No&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;14&lt;/td&gt;&lt;td&gt;PinExcd&lt;/td&gt;&lt;td&gt;Pin Tries Exceeded&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;15&lt;/td&gt;&lt;td&gt;PinInvld&lt;/td&gt;&lt;td&gt;Pin Invalid&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;16&lt;/td&gt;&lt;td&gt;AuthReq&lt;/td&gt;&lt;td&gt;Authentication Required&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;17&lt;/td&gt;&lt;td&gt;AuthenFail&lt;/td&gt;&lt;td&gt;Authentication Failed&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;18&lt;/td&gt;&lt;td&gt;Verified&lt;/td&gt;&lt;td&gt;Card Verified&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;19&lt;/td&gt;&lt;td&gt;Cancelled&lt;/td&gt;&lt;td&gt;Cancelled&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;20&lt;/td&gt;&lt;td&gt;Un&lt;/td&gt;&lt;td&gt;Unknown&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;21&lt;/td&gt;&lt;td&gt;Challenged&lt;/td&gt;&lt;td&gt;Challenged&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;22&lt;/td&gt;&lt;td&gt;Decoupled&lt;/td&gt;&lt;td&gt;Decoupled&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;&lt;td&gt;23&lt;/td&gt;&lt;td&gt;Denied&lt;/td&gt;&lt;td&gt;Permission Denied&lt;/td&gt;&lt;/tr&gt; &lt;/table&gt;  | 
**ResultCode** | **string** | A result code of the transaction identifying the result of the transaction for success, rejection or decline. | 
**Scheme** | Pointer to **string** | The name of the card scheme of the transaction that processed the transaction such as Visa or MasterCard.  | [optional] 
**SchemeId** | Pointer to **string** | The name of the card scheme of the transaction such as VI or MC.  | [optional] 
**SchemeLogo** | Pointer to **string** | A url containing a logo of the card scheme.  | [optional] 
**Transno** | Pointer to **int32** | The resulting transaction number, ordered incrementally from 1 for every merchant_id. The value will default to less than 1 for transactions that do not have a transaction number issued.  | [optional] 

## Methods

### NewBatchTransactionResultModel

`func NewBatchTransactionResultModel(accountId string, identifier string, merchantid int32, message string, result int32, resultCode string, ) *BatchTransactionResultModel`

NewBatchTransactionResultModel instantiates a new BatchTransactionResultModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchTransactionResultModelWithDefaults

`func NewBatchTransactionResultModelWithDefaults() *BatchTransactionResultModel`

NewBatchTransactionResultModelWithDefaults instantiates a new BatchTransactionResultModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *BatchTransactionResultModel) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BatchTransactionResultModel) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BatchTransactionResultModel) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAmount

`func (o *BatchTransactionResultModel) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BatchTransactionResultModel) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BatchTransactionResultModel) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BatchTransactionResultModel) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAuthcode

`func (o *BatchTransactionResultModel) GetAuthcode() string`

GetAuthcode returns the Authcode field if non-nil, zero value otherwise.

### GetAuthcodeOk

`func (o *BatchTransactionResultModel) GetAuthcodeOk() (*string, bool)`

GetAuthcodeOk returns a tuple with the Authcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthcode

`func (o *BatchTransactionResultModel) SetAuthcode(v string)`

SetAuthcode sets Authcode field to given value.

### HasAuthcode

`func (o *BatchTransactionResultModel) HasAuthcode() bool`

HasAuthcode returns a boolean if a field has been set.

### GetDatetime

`func (o *BatchTransactionResultModel) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *BatchTransactionResultModel) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *BatchTransactionResultModel) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *BatchTransactionResultModel) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetIdentifier

`func (o *BatchTransactionResultModel) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *BatchTransactionResultModel) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *BatchTransactionResultModel) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetMaskedpan

`func (o *BatchTransactionResultModel) GetMaskedpan() string`

GetMaskedpan returns the Maskedpan field if non-nil, zero value otherwise.

### GetMaskedpanOk

`func (o *BatchTransactionResultModel) GetMaskedpanOk() (*string, bool)`

GetMaskedpanOk returns a tuple with the Maskedpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedpan

`func (o *BatchTransactionResultModel) SetMaskedpan(v string)`

SetMaskedpan sets Maskedpan field to given value.

### HasMaskedpan

`func (o *BatchTransactionResultModel) HasMaskedpan() bool`

HasMaskedpan returns a boolean if a field has been set.

### GetMerchantid

`func (o *BatchTransactionResultModel) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *BatchTransactionResultModel) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *BatchTransactionResultModel) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.


### GetMessage

`func (o *BatchTransactionResultModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BatchTransactionResultModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BatchTransactionResultModel) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResult

`func (o *BatchTransactionResultModel) GetResult() int32`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *BatchTransactionResultModel) GetResultOk() (*int32, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *BatchTransactionResultModel) SetResult(v int32)`

SetResult sets Result field to given value.


### GetResultCode

`func (o *BatchTransactionResultModel) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *BatchTransactionResultModel) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *BatchTransactionResultModel) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.


### GetScheme

`func (o *BatchTransactionResultModel) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *BatchTransactionResultModel) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *BatchTransactionResultModel) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *BatchTransactionResultModel) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetSchemeId

`func (o *BatchTransactionResultModel) GetSchemeId() string`

GetSchemeId returns the SchemeId field if non-nil, zero value otherwise.

### GetSchemeIdOk

`func (o *BatchTransactionResultModel) GetSchemeIdOk() (*string, bool)`

GetSchemeIdOk returns a tuple with the SchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeId

`func (o *BatchTransactionResultModel) SetSchemeId(v string)`

SetSchemeId sets SchemeId field to given value.

### HasSchemeId

`func (o *BatchTransactionResultModel) HasSchemeId() bool`

HasSchemeId returns a boolean if a field has been set.

### GetSchemeLogo

`func (o *BatchTransactionResultModel) GetSchemeLogo() string`

GetSchemeLogo returns the SchemeLogo field if non-nil, zero value otherwise.

### GetSchemeLogoOk

`func (o *BatchTransactionResultModel) GetSchemeLogoOk() (*string, bool)`

GetSchemeLogoOk returns a tuple with the SchemeLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeLogo

`func (o *BatchTransactionResultModel) SetSchemeLogo(v string)`

SetSchemeLogo sets SchemeLogo field to given value.

### HasSchemeLogo

`func (o *BatchTransactionResultModel) HasSchemeLogo() bool`

HasSchemeLogo returns a boolean if a field has been set.

### GetTransno

`func (o *BatchTransactionResultModel) GetTransno() int32`

GetTransno returns the Transno field if non-nil, zero value otherwise.

### GetTransnoOk

`func (o *BatchTransactionResultModel) GetTransnoOk() (*int32, bool)`

GetTransnoOk returns a tuple with the Transno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransno

`func (o *BatchTransactionResultModel) SetTransno(v int32)`

SetTransno sets Transno field to given value.

### HasTransno

`func (o *BatchTransactionResultModel) HasTransno() bool`

HasTransno returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


