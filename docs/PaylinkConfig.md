# PaylinkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcsMode** | Pointer to **string** | Specifies the approach to be adopted by the Paylink form when displaying a 3-D Secure challenge window. The values may be  iframe: shows the 3-D Secure ACS in an iframe dialog, neatly embedding it in Paylink. This provides a more seamless flow for the cardholder who is able to validate and authenticate their card using a dialog provided by their card issuer.  inline: an inline mode transfers the full browser window to the authentication server, allowing the payment cardholder to see their payment card issuer&#39;s URL and the certificate status in the browser. If you request an iframe mode and the browser width is deemed as being small (&lt; 768px) then an inline mode will be enforced. This is to ensure that mobile users have an improved user experience.  | [optional] 
**CustomParams** | Pointer to [**[]PaylinkCustomParam**](PaylinkCustomParam.md) |  | [optional] 
**Descriptor** | Pointer to **string** | Directly specify the merchant descriptor used for the transaction to be displayed on the payment page. | [optional] 
**ExpireIn** | Pointer to **string** | Specifies a period of time in seconds after which the token cannot be used. A value of 0 defines that the token will never expire. The API will convert an expiry time based on a string value. For instance:   s - Time in seconds, for example 90s.   m - Time in minutes, for example 20m.   h - Time in hours, for example 4h.   w - Time in weeks, for example 4w.   M - Time in months, for example 6M.   y - Time in years, for example 1y.   Defaults to 30 minutes.  | [optional] 
**FieldGuard** | Pointer to [**[]PaylinkFieldGuardModel**](PaylinkFieldGuardModel.md) |  | [optional] 
**LockParams** | Pointer to **[]string** |  | [optional] 
**MerchLogo** | Pointer to **string** | A URL of a logo to include in the form. The URL should be delivered using HTTPS. | [optional] 
**MerchTerms** | Pointer to **string** | A URL of the merchant terms and conditions for payment. If a value is supplied, a checkbox will be required to be completed to confirm that the cardholder agrees to these conditions before payment. A modal dialogue is displayed with the content of the conditions displayed. | [optional] 
**Options** | Pointer to **[]string** |  | [optional] 
**PartPayments** | Pointer to [**PaylinkPartPayments**](PaylinkPartPayments.md) |  | [optional] 
**PassThroughData** | Pointer to **map[string]string** |  | [optional] 
**PassThroughHeaders** | Pointer to **map[string]string** |  | [optional] 
**Postback** | Pointer to **string** | Specifies a URL to use for a call back when the payment is completed. see Postback Handling }. | [optional] 
**PostbackPassword** | Pointer to **string** | A password to be added to the postback for HTTP Basic Authentication. | [optional] 
**PostbackPolicy** | Pointer to **string** | The policy setting for the postback see Postback Handling. | [optional] 
**PostbackUsername** | Pointer to **string** | A username to be added to the postback for HTTP Basic Authentication. | [optional] 
**RedirectDelay** | Pointer to **int32** | A value which can delay the redirection in seconds. A value of 0 will redirect immediately. | [optional] 
**RedirectFailure** | Pointer to **string** | A URL which the browser is redirected to on non-completion of a transaction. | [optional] 
**RedirectSuccess** | Pointer to **string** | A URL which the browser is redirected to on authorisation of a transaction. | [optional] 
**Renderer** | Pointer to **string** | The Paylink renderer engine to use. | [optional] 
**ReturnParams** | Pointer to **bool** | If a value of true is specified, any redirection will include the transaction result in parameters. It is recommended to use the postback integration rather than redirection parameters. | [optional] 
**Ui** | Pointer to [**PaylinkUI**](PaylinkUI.md) |  | [optional] 

## Methods

### NewPaylinkConfig

`func NewPaylinkConfig() *PaylinkConfig`

NewPaylinkConfig instantiates a new PaylinkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkConfigWithDefaults

`func NewPaylinkConfigWithDefaults() *PaylinkConfig`

NewPaylinkConfigWithDefaults instantiates a new PaylinkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcsMode

`func (o *PaylinkConfig) GetAcsMode() string`

GetAcsMode returns the AcsMode field if non-nil, zero value otherwise.

### GetAcsModeOk

`func (o *PaylinkConfig) GetAcsModeOk() (*string, bool)`

GetAcsModeOk returns a tuple with the AcsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsMode

`func (o *PaylinkConfig) SetAcsMode(v string)`

SetAcsMode sets AcsMode field to given value.

### HasAcsMode

`func (o *PaylinkConfig) HasAcsMode() bool`

HasAcsMode returns a boolean if a field has been set.

### GetCustomParams

`func (o *PaylinkConfig) GetCustomParams() []PaylinkCustomParam`

GetCustomParams returns the CustomParams field if non-nil, zero value otherwise.

### GetCustomParamsOk

`func (o *PaylinkConfig) GetCustomParamsOk() (*[]PaylinkCustomParam, bool)`

GetCustomParamsOk returns a tuple with the CustomParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParams

`func (o *PaylinkConfig) SetCustomParams(v []PaylinkCustomParam)`

SetCustomParams sets CustomParams field to given value.

### HasCustomParams

`func (o *PaylinkConfig) HasCustomParams() bool`

HasCustomParams returns a boolean if a field has been set.

### GetDescriptor

`func (o *PaylinkConfig) GetDescriptor() string`

GetDescriptor returns the Descriptor field if non-nil, zero value otherwise.

### GetDescriptorOk

`func (o *PaylinkConfig) GetDescriptorOk() (*string, bool)`

GetDescriptorOk returns a tuple with the Descriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptor

`func (o *PaylinkConfig) SetDescriptor(v string)`

SetDescriptor sets Descriptor field to given value.

### HasDescriptor

`func (o *PaylinkConfig) HasDescriptor() bool`

HasDescriptor returns a boolean if a field has been set.

### GetExpireIn

`func (o *PaylinkConfig) GetExpireIn() string`

GetExpireIn returns the ExpireIn field if non-nil, zero value otherwise.

### GetExpireInOk

`func (o *PaylinkConfig) GetExpireInOk() (*string, bool)`

GetExpireInOk returns a tuple with the ExpireIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireIn

`func (o *PaylinkConfig) SetExpireIn(v string)`

SetExpireIn sets ExpireIn field to given value.

### HasExpireIn

`func (o *PaylinkConfig) HasExpireIn() bool`

HasExpireIn returns a boolean if a field has been set.

### GetFieldGuard

`func (o *PaylinkConfig) GetFieldGuard() []PaylinkFieldGuardModel`

GetFieldGuard returns the FieldGuard field if non-nil, zero value otherwise.

### GetFieldGuardOk

`func (o *PaylinkConfig) GetFieldGuardOk() (*[]PaylinkFieldGuardModel, bool)`

GetFieldGuardOk returns a tuple with the FieldGuard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGuard

`func (o *PaylinkConfig) SetFieldGuard(v []PaylinkFieldGuardModel)`

SetFieldGuard sets FieldGuard field to given value.

### HasFieldGuard

`func (o *PaylinkConfig) HasFieldGuard() bool`

HasFieldGuard returns a boolean if a field has been set.

### GetLockParams

`func (o *PaylinkConfig) GetLockParams() []string`

GetLockParams returns the LockParams field if non-nil, zero value otherwise.

### GetLockParamsOk

`func (o *PaylinkConfig) GetLockParamsOk() (*[]string, bool)`

GetLockParamsOk returns a tuple with the LockParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockParams

`func (o *PaylinkConfig) SetLockParams(v []string)`

SetLockParams sets LockParams field to given value.

### HasLockParams

`func (o *PaylinkConfig) HasLockParams() bool`

HasLockParams returns a boolean if a field has been set.

### GetMerchLogo

`func (o *PaylinkConfig) GetMerchLogo() string`

GetMerchLogo returns the MerchLogo field if non-nil, zero value otherwise.

### GetMerchLogoOk

`func (o *PaylinkConfig) GetMerchLogoOk() (*string, bool)`

GetMerchLogoOk returns a tuple with the MerchLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchLogo

`func (o *PaylinkConfig) SetMerchLogo(v string)`

SetMerchLogo sets MerchLogo field to given value.

### HasMerchLogo

`func (o *PaylinkConfig) HasMerchLogo() bool`

HasMerchLogo returns a boolean if a field has been set.

### GetMerchTerms

`func (o *PaylinkConfig) GetMerchTerms() string`

GetMerchTerms returns the MerchTerms field if non-nil, zero value otherwise.

### GetMerchTermsOk

`func (o *PaylinkConfig) GetMerchTermsOk() (*string, bool)`

GetMerchTermsOk returns a tuple with the MerchTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchTerms

`func (o *PaylinkConfig) SetMerchTerms(v string)`

SetMerchTerms sets MerchTerms field to given value.

### HasMerchTerms

`func (o *PaylinkConfig) HasMerchTerms() bool`

HasMerchTerms returns a boolean if a field has been set.

### GetOptions

`func (o *PaylinkConfig) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PaylinkConfig) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PaylinkConfig) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PaylinkConfig) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPartPayments

`func (o *PaylinkConfig) GetPartPayments() PaylinkPartPayments`

GetPartPayments returns the PartPayments field if non-nil, zero value otherwise.

### GetPartPaymentsOk

`func (o *PaylinkConfig) GetPartPaymentsOk() (*PaylinkPartPayments, bool)`

GetPartPaymentsOk returns a tuple with the PartPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartPayments

`func (o *PaylinkConfig) SetPartPayments(v PaylinkPartPayments)`

SetPartPayments sets PartPayments field to given value.

### HasPartPayments

`func (o *PaylinkConfig) HasPartPayments() bool`

HasPartPayments returns a boolean if a field has been set.

### GetPassThroughData

`func (o *PaylinkConfig) GetPassThroughData() map[string]string`

GetPassThroughData returns the PassThroughData field if non-nil, zero value otherwise.

### GetPassThroughDataOk

`func (o *PaylinkConfig) GetPassThroughDataOk() (*map[string]string, bool)`

GetPassThroughDataOk returns a tuple with the PassThroughData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassThroughData

`func (o *PaylinkConfig) SetPassThroughData(v map[string]string)`

SetPassThroughData sets PassThroughData field to given value.

### HasPassThroughData

`func (o *PaylinkConfig) HasPassThroughData() bool`

HasPassThroughData returns a boolean if a field has been set.

### GetPassThroughHeaders

`func (o *PaylinkConfig) GetPassThroughHeaders() map[string]string`

GetPassThroughHeaders returns the PassThroughHeaders field if non-nil, zero value otherwise.

### GetPassThroughHeadersOk

`func (o *PaylinkConfig) GetPassThroughHeadersOk() (*map[string]string, bool)`

GetPassThroughHeadersOk returns a tuple with the PassThroughHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassThroughHeaders

`func (o *PaylinkConfig) SetPassThroughHeaders(v map[string]string)`

SetPassThroughHeaders sets PassThroughHeaders field to given value.

### HasPassThroughHeaders

`func (o *PaylinkConfig) HasPassThroughHeaders() bool`

HasPassThroughHeaders returns a boolean if a field has been set.

### GetPostback

`func (o *PaylinkConfig) GetPostback() string`

GetPostback returns the Postback field if non-nil, zero value otherwise.

### GetPostbackOk

`func (o *PaylinkConfig) GetPostbackOk() (*string, bool)`

GetPostbackOk returns a tuple with the Postback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostback

`func (o *PaylinkConfig) SetPostback(v string)`

SetPostback sets Postback field to given value.

### HasPostback

`func (o *PaylinkConfig) HasPostback() bool`

HasPostback returns a boolean if a field has been set.

### GetPostbackPassword

`func (o *PaylinkConfig) GetPostbackPassword() string`

GetPostbackPassword returns the PostbackPassword field if non-nil, zero value otherwise.

### GetPostbackPasswordOk

`func (o *PaylinkConfig) GetPostbackPasswordOk() (*string, bool)`

GetPostbackPasswordOk returns a tuple with the PostbackPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostbackPassword

`func (o *PaylinkConfig) SetPostbackPassword(v string)`

SetPostbackPassword sets PostbackPassword field to given value.

### HasPostbackPassword

`func (o *PaylinkConfig) HasPostbackPassword() bool`

HasPostbackPassword returns a boolean if a field has been set.

### GetPostbackPolicy

`func (o *PaylinkConfig) GetPostbackPolicy() string`

GetPostbackPolicy returns the PostbackPolicy field if non-nil, zero value otherwise.

### GetPostbackPolicyOk

`func (o *PaylinkConfig) GetPostbackPolicyOk() (*string, bool)`

GetPostbackPolicyOk returns a tuple with the PostbackPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostbackPolicy

`func (o *PaylinkConfig) SetPostbackPolicy(v string)`

SetPostbackPolicy sets PostbackPolicy field to given value.

### HasPostbackPolicy

`func (o *PaylinkConfig) HasPostbackPolicy() bool`

HasPostbackPolicy returns a boolean if a field has been set.

### GetPostbackUsername

`func (o *PaylinkConfig) GetPostbackUsername() string`

GetPostbackUsername returns the PostbackUsername field if non-nil, zero value otherwise.

### GetPostbackUsernameOk

`func (o *PaylinkConfig) GetPostbackUsernameOk() (*string, bool)`

GetPostbackUsernameOk returns a tuple with the PostbackUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostbackUsername

`func (o *PaylinkConfig) SetPostbackUsername(v string)`

SetPostbackUsername sets PostbackUsername field to given value.

### HasPostbackUsername

`func (o *PaylinkConfig) HasPostbackUsername() bool`

HasPostbackUsername returns a boolean if a field has been set.

### GetRedirectDelay

`func (o *PaylinkConfig) GetRedirectDelay() int32`

GetRedirectDelay returns the RedirectDelay field if non-nil, zero value otherwise.

### GetRedirectDelayOk

`func (o *PaylinkConfig) GetRedirectDelayOk() (*int32, bool)`

GetRedirectDelayOk returns a tuple with the RedirectDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectDelay

`func (o *PaylinkConfig) SetRedirectDelay(v int32)`

SetRedirectDelay sets RedirectDelay field to given value.

### HasRedirectDelay

`func (o *PaylinkConfig) HasRedirectDelay() bool`

HasRedirectDelay returns a boolean if a field has been set.

### GetRedirectFailure

`func (o *PaylinkConfig) GetRedirectFailure() string`

GetRedirectFailure returns the RedirectFailure field if non-nil, zero value otherwise.

### GetRedirectFailureOk

`func (o *PaylinkConfig) GetRedirectFailureOk() (*string, bool)`

GetRedirectFailureOk returns a tuple with the RedirectFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectFailure

`func (o *PaylinkConfig) SetRedirectFailure(v string)`

SetRedirectFailure sets RedirectFailure field to given value.

### HasRedirectFailure

`func (o *PaylinkConfig) HasRedirectFailure() bool`

HasRedirectFailure returns a boolean if a field has been set.

### GetRedirectSuccess

`func (o *PaylinkConfig) GetRedirectSuccess() string`

GetRedirectSuccess returns the RedirectSuccess field if non-nil, zero value otherwise.

### GetRedirectSuccessOk

`func (o *PaylinkConfig) GetRedirectSuccessOk() (*string, bool)`

GetRedirectSuccessOk returns a tuple with the RedirectSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectSuccess

`func (o *PaylinkConfig) SetRedirectSuccess(v string)`

SetRedirectSuccess sets RedirectSuccess field to given value.

### HasRedirectSuccess

`func (o *PaylinkConfig) HasRedirectSuccess() bool`

HasRedirectSuccess returns a boolean if a field has been set.

### GetRenderer

`func (o *PaylinkConfig) GetRenderer() string`

GetRenderer returns the Renderer field if non-nil, zero value otherwise.

### GetRendererOk

`func (o *PaylinkConfig) GetRendererOk() (*string, bool)`

GetRendererOk returns a tuple with the Renderer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderer

`func (o *PaylinkConfig) SetRenderer(v string)`

SetRenderer sets Renderer field to given value.

### HasRenderer

`func (o *PaylinkConfig) HasRenderer() bool`

HasRenderer returns a boolean if a field has been set.

### GetReturnParams

`func (o *PaylinkConfig) GetReturnParams() bool`

GetReturnParams returns the ReturnParams field if non-nil, zero value otherwise.

### GetReturnParamsOk

`func (o *PaylinkConfig) GetReturnParamsOk() (*bool, bool)`

GetReturnParamsOk returns a tuple with the ReturnParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnParams

`func (o *PaylinkConfig) SetReturnParams(v bool)`

SetReturnParams sets ReturnParams field to given value.

### HasReturnParams

`func (o *PaylinkConfig) HasReturnParams() bool`

HasReturnParams returns a boolean if a field has been set.

### GetUi

`func (o *PaylinkConfig) GetUi() PaylinkUI`

GetUi returns the Ui field if non-nil, zero value otherwise.

### GetUiOk

`func (o *PaylinkConfig) GetUiOk() (*PaylinkUI, bool)`

GetUiOk returns a tuple with the Ui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUi

`func (o *PaylinkConfig) SetUi(v PaylinkUI)`

SetUi sets Ui field to given value.

### HasUi

`func (o *PaylinkConfig) HasUi() bool`

HasUi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


