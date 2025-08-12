# WebHookChannelCreateRequestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectTimeout** | Pointer to **int32** | The connection timeout in milliseconds. | [optional] 
**ContentType** | Pointer to **string** | The content type of the http call. Defaults to &#x60;application/json&#x60;. | [optional] 
**Headers** | Pointer to **map[string]interface{}** | Http headers to add to the configuration. | [optional] 
**Method** | Pointer to **string** | The HTTP method to use. Defaults to &#x60;POST&#x60;. | [optional] 
**ReadTimeout** | Pointer to **int32** | The read timeout in milliseconds when waiting for a reply. | [optional] 
**Url** | **string** | The url of the endpoint to contact. The value should be https. | 

## Methods

### NewWebHookChannelCreateRequestConfig

`func NewWebHookChannelCreateRequestConfig(url string, ) *WebHookChannelCreateRequestConfig`

NewWebHookChannelCreateRequestConfig instantiates a new WebHookChannelCreateRequestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebHookChannelCreateRequestConfigWithDefaults

`func NewWebHookChannelCreateRequestConfigWithDefaults() *WebHookChannelCreateRequestConfig`

NewWebHookChannelCreateRequestConfigWithDefaults instantiates a new WebHookChannelCreateRequestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectTimeout

`func (o *WebHookChannelCreateRequestConfig) GetConnectTimeout() int32`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *WebHookChannelCreateRequestConfig) GetConnectTimeoutOk() (*int32, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *WebHookChannelCreateRequestConfig) SetConnectTimeout(v int32)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *WebHookChannelCreateRequestConfig) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetContentType

`func (o *WebHookChannelCreateRequestConfig) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *WebHookChannelCreateRequestConfig) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *WebHookChannelCreateRequestConfig) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *WebHookChannelCreateRequestConfig) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetHeaders

`func (o *WebHookChannelCreateRequestConfig) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebHookChannelCreateRequestConfig) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WebHookChannelCreateRequestConfig) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *WebHookChannelCreateRequestConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMethod

`func (o *WebHookChannelCreateRequestConfig) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *WebHookChannelCreateRequestConfig) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *WebHookChannelCreateRequestConfig) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *WebHookChannelCreateRequestConfig) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetReadTimeout

`func (o *WebHookChannelCreateRequestConfig) GetReadTimeout() int32`

GetReadTimeout returns the ReadTimeout field if non-nil, zero value otherwise.

### GetReadTimeoutOk

`func (o *WebHookChannelCreateRequestConfig) GetReadTimeoutOk() (*int32, bool)`

GetReadTimeoutOk returns a tuple with the ReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTimeout

`func (o *WebHookChannelCreateRequestConfig) SetReadTimeout(v int32)`

SetReadTimeout sets ReadTimeout field to given value.

### HasReadTimeout

`func (o *WebHookChannelCreateRequestConfig) HasReadTimeout() bool`

HasReadTimeout returns a boolean if a field has been set.

### GetUrl

`func (o *WebHookChannelCreateRequestConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebHookChannelCreateRequestConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebHookChannelCreateRequestConfig) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


