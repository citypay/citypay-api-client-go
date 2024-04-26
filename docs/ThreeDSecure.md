# ThreeDSecure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptHeaders** | Pointer to **string** | Required for 3DSv1. Optional if the &#x60;cp_bx&#x60; value is provided otherwise required for 3Dv2 processing operating in browser authentication mode.  The &#x60;cp_bx&#x60; value will override any value supplied to this field.  The content of the HTTP accept header as sent to the merchant from the cardholder&#39;s user agent.  This value will be validated by the ACS when the card holder authenticates themselves to verify that no intermediary is performing this action. Required for 3DSv1.  | [optional] 
**BrowserColorDepth** | Pointer to **string** | BrowserColorDepth field used for 3DSv2 browser enablement. Recommendation is to use citypay.js and the &#x60;bx&#x60; function to gather this value. | [optional] 
**BrowserIP** | Pointer to **string** | BrowserIP field used for 3DSv2 browser enablement. Recommendation is to use citypay.js and the &#x60;bx&#x60; function to gather this value. | [optional] 
**BrowserJavaEnabled** | Pointer to **string** | BrowserJavaEnabled field used for 3DSv2 browser enablement. Recommendation is to use citypay.js and the &#x60;bx&#x60; function to gather this value. | [optional] 
**BrowserLanguage** | Pointer to **string** | BrowserLanguage field used for 3DSv2 browser enablement. Recommendation is to use citypay.js and the &#x60;bx&#x60; function to gather this value. | [optional] 
**BrowserScreenHeight** | Pointer to **string** | BrowserScreenHeight field used for 3DSv2 browser enablement. Recommendation is to use citypay.js and the &#x60;bx&#x60; function to gather this value. | [optional] 
**BrowserScreenWidth** | Pointer to **string** | BrowserScreenWidth field used for 3DSv2 browser enablement. Recommendation is to use citypay.js and the &#x60;bx&#x60; function to gather this value. | [optional] 
**BrowserTZ** | Pointer to **string** | BrowserTZ offset field used for 3DSv2 browser enablement. Recommendation is to use citypay.js and the &#x60;bx&#x60; function to gather this value. | [optional] 
**CpBx** | Pointer to **string** | Required for 3DSv2.  Browser extension value produced by the citypay.js &#x60;bx&#x60; function. See [https://sandbox.citypay.com/3dsv2/bx](https://sandbox.citypay.com/3dsv2/bx) for  details.  | [optional] 
**Downgrade1** | Pointer to **bool** | Where a merchant is configured for 3DSv2, setting this option will attempt to downgrade the transaction to  3DSv1.  | [optional] 
**MerchantTermurl** | Pointer to **string** | A controller URL for 3D-Secure processing that any response from an authentication request or challenge request should be sent to.  The controller should forward on the response from the URL back via this API for subsequent processing.  | [optional] 
**TdsPolicy** | Pointer to **string** | A policy value which determines whether ThreeDSecure is enforced or bypassed. Note that this will only work for e-commerce transactions and accounts that have 3DSecure enabled and fully registered with Visa, MasterCard or American Express. It is useful when transactions may be wanted to bypass processing rules.  Note that this may affect the liability shift of transactions and may occur a higher fee with the acquiring bank.  Values are   &#x60;0&#x60; for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   &#x60;1&#x60; for an enforced policy. Transactions will be enabled for 3DS processing   &#x60;2&#x60; to bypass. Transactions that are bypassed will switch off 3DS processing.  | [optional] 
**UserAgent** | Pointer to **string** | Required for 3DSv1.  Optional if the &#x60;cp_bx&#x60; value is provided otherwise required 3Dv2 processing operating in browser authentication mode.  The &#x60;cp_bx&#x60; value will override any value supplied to this field.  The content of the HTTP user-agent header as sent to the merchant from the cardholder&#39;s user agent.  This value will be validated by the ACS when the card holder authenticates themselves to verify that no intermediary is performing this action. Required for 3DSv1.  | [optional] 

## Methods

### NewThreeDSecure

`func NewThreeDSecure() *ThreeDSecure`

NewThreeDSecure instantiates a new ThreeDSecure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreeDSecureWithDefaults

`func NewThreeDSecureWithDefaults() *ThreeDSecure`

NewThreeDSecureWithDefaults instantiates a new ThreeDSecure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptHeaders

`func (o *ThreeDSecure) GetAcceptHeaders() string`

GetAcceptHeaders returns the AcceptHeaders field if non-nil, zero value otherwise.

### GetAcceptHeadersOk

`func (o *ThreeDSecure) GetAcceptHeadersOk() (*string, bool)`

GetAcceptHeadersOk returns a tuple with the AcceptHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptHeaders

`func (o *ThreeDSecure) SetAcceptHeaders(v string)`

SetAcceptHeaders sets AcceptHeaders field to given value.

### HasAcceptHeaders

`func (o *ThreeDSecure) HasAcceptHeaders() bool`

HasAcceptHeaders returns a boolean if a field has been set.

### GetBrowserColorDepth

`func (o *ThreeDSecure) GetBrowserColorDepth() string`

GetBrowserColorDepth returns the BrowserColorDepth field if non-nil, zero value otherwise.

### GetBrowserColorDepthOk

`func (o *ThreeDSecure) GetBrowserColorDepthOk() (*string, bool)`

GetBrowserColorDepthOk returns a tuple with the BrowserColorDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserColorDepth

`func (o *ThreeDSecure) SetBrowserColorDepth(v string)`

SetBrowserColorDepth sets BrowserColorDepth field to given value.

### HasBrowserColorDepth

`func (o *ThreeDSecure) HasBrowserColorDepth() bool`

HasBrowserColorDepth returns a boolean if a field has been set.

### GetBrowserIP

`func (o *ThreeDSecure) GetBrowserIP() string`

GetBrowserIP returns the BrowserIP field if non-nil, zero value otherwise.

### GetBrowserIPOk

`func (o *ThreeDSecure) GetBrowserIPOk() (*string, bool)`

GetBrowserIPOk returns a tuple with the BrowserIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserIP

`func (o *ThreeDSecure) SetBrowserIP(v string)`

SetBrowserIP sets BrowserIP field to given value.

### HasBrowserIP

`func (o *ThreeDSecure) HasBrowserIP() bool`

HasBrowserIP returns a boolean if a field has been set.

### GetBrowserJavaEnabled

`func (o *ThreeDSecure) GetBrowserJavaEnabled() string`

GetBrowserJavaEnabled returns the BrowserJavaEnabled field if non-nil, zero value otherwise.

### GetBrowserJavaEnabledOk

`func (o *ThreeDSecure) GetBrowserJavaEnabledOk() (*string, bool)`

GetBrowserJavaEnabledOk returns a tuple with the BrowserJavaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserJavaEnabled

`func (o *ThreeDSecure) SetBrowserJavaEnabled(v string)`

SetBrowserJavaEnabled sets BrowserJavaEnabled field to given value.

### HasBrowserJavaEnabled

`func (o *ThreeDSecure) HasBrowserJavaEnabled() bool`

HasBrowserJavaEnabled returns a boolean if a field has been set.

### GetBrowserLanguage

`func (o *ThreeDSecure) GetBrowserLanguage() string`

GetBrowserLanguage returns the BrowserLanguage field if non-nil, zero value otherwise.

### GetBrowserLanguageOk

`func (o *ThreeDSecure) GetBrowserLanguageOk() (*string, bool)`

GetBrowserLanguageOk returns a tuple with the BrowserLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserLanguage

`func (o *ThreeDSecure) SetBrowserLanguage(v string)`

SetBrowserLanguage sets BrowserLanguage field to given value.

### HasBrowserLanguage

`func (o *ThreeDSecure) HasBrowserLanguage() bool`

HasBrowserLanguage returns a boolean if a field has been set.

### GetBrowserScreenHeight

`func (o *ThreeDSecure) GetBrowserScreenHeight() string`

GetBrowserScreenHeight returns the BrowserScreenHeight field if non-nil, zero value otherwise.

### GetBrowserScreenHeightOk

`func (o *ThreeDSecure) GetBrowserScreenHeightOk() (*string, bool)`

GetBrowserScreenHeightOk returns a tuple with the BrowserScreenHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserScreenHeight

`func (o *ThreeDSecure) SetBrowserScreenHeight(v string)`

SetBrowserScreenHeight sets BrowserScreenHeight field to given value.

### HasBrowserScreenHeight

`func (o *ThreeDSecure) HasBrowserScreenHeight() bool`

HasBrowserScreenHeight returns a boolean if a field has been set.

### GetBrowserScreenWidth

`func (o *ThreeDSecure) GetBrowserScreenWidth() string`

GetBrowserScreenWidth returns the BrowserScreenWidth field if non-nil, zero value otherwise.

### GetBrowserScreenWidthOk

`func (o *ThreeDSecure) GetBrowserScreenWidthOk() (*string, bool)`

GetBrowserScreenWidthOk returns a tuple with the BrowserScreenWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserScreenWidth

`func (o *ThreeDSecure) SetBrowserScreenWidth(v string)`

SetBrowserScreenWidth sets BrowserScreenWidth field to given value.

### HasBrowserScreenWidth

`func (o *ThreeDSecure) HasBrowserScreenWidth() bool`

HasBrowserScreenWidth returns a boolean if a field has been set.

### GetBrowserTZ

`func (o *ThreeDSecure) GetBrowserTZ() string`

GetBrowserTZ returns the BrowserTZ field if non-nil, zero value otherwise.

### GetBrowserTZOk

`func (o *ThreeDSecure) GetBrowserTZOk() (*string, bool)`

GetBrowserTZOk returns a tuple with the BrowserTZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserTZ

`func (o *ThreeDSecure) SetBrowserTZ(v string)`

SetBrowserTZ sets BrowserTZ field to given value.

### HasBrowserTZ

`func (o *ThreeDSecure) HasBrowserTZ() bool`

HasBrowserTZ returns a boolean if a field has been set.

### GetCpBx

`func (o *ThreeDSecure) GetCpBx() string`

GetCpBx returns the CpBx field if non-nil, zero value otherwise.

### GetCpBxOk

`func (o *ThreeDSecure) GetCpBxOk() (*string, bool)`

GetCpBxOk returns a tuple with the CpBx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpBx

`func (o *ThreeDSecure) SetCpBx(v string)`

SetCpBx sets CpBx field to given value.

### HasCpBx

`func (o *ThreeDSecure) HasCpBx() bool`

HasCpBx returns a boolean if a field has been set.

### GetDowngrade1

`func (o *ThreeDSecure) GetDowngrade1() bool`

GetDowngrade1 returns the Downgrade1 field if non-nil, zero value otherwise.

### GetDowngrade1Ok

`func (o *ThreeDSecure) GetDowngrade1Ok() (*bool, bool)`

GetDowngrade1Ok returns a tuple with the Downgrade1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowngrade1

`func (o *ThreeDSecure) SetDowngrade1(v bool)`

SetDowngrade1 sets Downgrade1 field to given value.

### HasDowngrade1

`func (o *ThreeDSecure) HasDowngrade1() bool`

HasDowngrade1 returns a boolean if a field has been set.

### GetMerchantTermurl

`func (o *ThreeDSecure) GetMerchantTermurl() string`

GetMerchantTermurl returns the MerchantTermurl field if non-nil, zero value otherwise.

### GetMerchantTermurlOk

`func (o *ThreeDSecure) GetMerchantTermurlOk() (*string, bool)`

GetMerchantTermurlOk returns a tuple with the MerchantTermurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTermurl

`func (o *ThreeDSecure) SetMerchantTermurl(v string)`

SetMerchantTermurl sets MerchantTermurl field to given value.

### HasMerchantTermurl

`func (o *ThreeDSecure) HasMerchantTermurl() bool`

HasMerchantTermurl returns a boolean if a field has been set.

### GetTdsPolicy

`func (o *ThreeDSecure) GetTdsPolicy() string`

GetTdsPolicy returns the TdsPolicy field if non-nil, zero value otherwise.

### GetTdsPolicyOk

`func (o *ThreeDSecure) GetTdsPolicyOk() (*string, bool)`

GetTdsPolicyOk returns a tuple with the TdsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdsPolicy

`func (o *ThreeDSecure) SetTdsPolicy(v string)`

SetTdsPolicy sets TdsPolicy field to given value.

### HasTdsPolicy

`func (o *ThreeDSecure) HasTdsPolicy() bool`

HasTdsPolicy returns a boolean if a field has been set.

### GetUserAgent

`func (o *ThreeDSecure) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *ThreeDSecure) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *ThreeDSecure) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *ThreeDSecure) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


