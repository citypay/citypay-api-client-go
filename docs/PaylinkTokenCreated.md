# PaylinkTokenCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to [**PaylinkAttachmentResult**](PaylinkAttachmentResult.md) |  | [optional] 
**Bps** | Pointer to **string** | true if BPS has been enabled on this token. | [optional] 
**DateCreated** | Pointer to **time.Time** | Date and time the token was generated. | [optional] 
**Errors** | Pointer to [**[]PaylinkErrorCode**](PaylinkErrorCode.md) |  | [optional] 
**Id** | **string** | A unique id of the request. | 
**Identifier** | Pointer to **string** | The identifier as presented in the TokenRequest. | [optional] 
**Mode** | Pointer to **string** | Determines whether the token is &#x60;live&#x60; or &#x60;test&#x60;. | [optional] 
**Qrcode** | Pointer to **string** | A URL of a qrcode which can be used to refer to the token URL. | [optional] 
**Result** | **int32** | The result field contains the result for the Paylink Token Request. 0 - indicates that an error was encountered while creating the token. 1 - which indicates that a Token was successfully created. | 
**ServerVersion** | Pointer to **string** | the version of the server performing the call. | [optional] 
**Source** | Pointer to **string** | The incoming IP address of the call. | [optional] 
**Token** | **string** | A token generated for the request used to refer to the transaction in consequential calls. | 
**Url** | Pointer to **string** | The Paylink token URL used to checkout by the card holder. | [optional] 
**Usc** | Pointer to **string** | A UrlShortCode (USC) used for short links. | [optional] 

## Methods

### NewPaylinkTokenCreated

`func NewPaylinkTokenCreated(id string, result int32, token string, ) *PaylinkTokenCreated`

NewPaylinkTokenCreated instantiates a new PaylinkTokenCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkTokenCreatedWithDefaults

`func NewPaylinkTokenCreatedWithDefaults() *PaylinkTokenCreated`

NewPaylinkTokenCreatedWithDefaults instantiates a new PaylinkTokenCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *PaylinkTokenCreated) GetAttachments() PaylinkAttachmentResult`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *PaylinkTokenCreated) GetAttachmentsOk() (*PaylinkAttachmentResult, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *PaylinkTokenCreated) SetAttachments(v PaylinkAttachmentResult)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *PaylinkTokenCreated) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetBps

`func (o *PaylinkTokenCreated) GetBps() string`

GetBps returns the Bps field if non-nil, zero value otherwise.

### GetBpsOk

`func (o *PaylinkTokenCreated) GetBpsOk() (*string, bool)`

GetBpsOk returns a tuple with the Bps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBps

`func (o *PaylinkTokenCreated) SetBps(v string)`

SetBps sets Bps field to given value.

### HasBps

`func (o *PaylinkTokenCreated) HasBps() bool`

HasBps returns a boolean if a field has been set.

### GetDateCreated

`func (o *PaylinkTokenCreated) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *PaylinkTokenCreated) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *PaylinkTokenCreated) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *PaylinkTokenCreated) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetErrors

`func (o *PaylinkTokenCreated) GetErrors() []PaylinkErrorCode`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *PaylinkTokenCreated) GetErrorsOk() (*[]PaylinkErrorCode, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *PaylinkTokenCreated) SetErrors(v []PaylinkErrorCode)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *PaylinkTokenCreated) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetId

`func (o *PaylinkTokenCreated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaylinkTokenCreated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaylinkTokenCreated) SetId(v string)`

SetId sets Id field to given value.


### GetIdentifier

`func (o *PaylinkTokenCreated) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PaylinkTokenCreated) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PaylinkTokenCreated) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PaylinkTokenCreated) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetMode

`func (o *PaylinkTokenCreated) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PaylinkTokenCreated) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PaylinkTokenCreated) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PaylinkTokenCreated) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetQrcode

`func (o *PaylinkTokenCreated) GetQrcode() string`

GetQrcode returns the Qrcode field if non-nil, zero value otherwise.

### GetQrcodeOk

`func (o *PaylinkTokenCreated) GetQrcodeOk() (*string, bool)`

GetQrcodeOk returns a tuple with the Qrcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrcode

`func (o *PaylinkTokenCreated) SetQrcode(v string)`

SetQrcode sets Qrcode field to given value.

### HasQrcode

`func (o *PaylinkTokenCreated) HasQrcode() bool`

HasQrcode returns a boolean if a field has been set.

### GetResult

`func (o *PaylinkTokenCreated) GetResult() int32`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PaylinkTokenCreated) GetResultOk() (*int32, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PaylinkTokenCreated) SetResult(v int32)`

SetResult sets Result field to given value.


### GetServerVersion

`func (o *PaylinkTokenCreated) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *PaylinkTokenCreated) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *PaylinkTokenCreated) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *PaylinkTokenCreated) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetSource

`func (o *PaylinkTokenCreated) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PaylinkTokenCreated) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PaylinkTokenCreated) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PaylinkTokenCreated) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetToken

`func (o *PaylinkTokenCreated) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaylinkTokenCreated) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaylinkTokenCreated) SetToken(v string)`

SetToken sets Token field to given value.


### GetUrl

`func (o *PaylinkTokenCreated) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PaylinkTokenCreated) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PaylinkTokenCreated) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PaylinkTokenCreated) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsc

`func (o *PaylinkTokenCreated) GetUsc() string`

GetUsc returns the Usc field if non-nil, zero value otherwise.

### GetUscOk

`func (o *PaylinkTokenCreated) GetUscOk() (*string, bool)`

GetUscOk returns a tuple with the Usc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsc

`func (o *PaylinkTokenCreated) SetUsc(v string)`

SetUsc sets Usc field to given value.

### HasUsc

`func (o *PaylinkTokenCreated) HasUsc() bool`

HasUsc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


