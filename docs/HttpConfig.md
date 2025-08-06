# HttpConfig

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

### NewHttpConfig

`func NewHttpConfig(url string, ) *HttpConfig`

NewHttpConfig instantiates a new HttpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpConfigWithDefaults

`func NewHttpConfigWithDefaults() *HttpConfig`

NewHttpConfigWithDefaults instantiates a new HttpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectTimeout

`func (o *HttpConfig) GetConnectTimeout() int32`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *HttpConfig) GetConnectTimeoutOk() (*int32, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *HttpConfig) SetConnectTimeout(v int32)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *HttpConfig) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetContentType

`func (o *HttpConfig) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *HttpConfig) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *HttpConfig) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *HttpConfig) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetHeaders

`func (o *HttpConfig) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HttpConfig) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HttpConfig) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HttpConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMethod

`func (o *HttpConfig) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HttpConfig) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HttpConfig) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *HttpConfig) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetReadTimeout

`func (o *HttpConfig) GetReadTimeout() int32`

GetReadTimeout returns the ReadTimeout field if non-nil, zero value otherwise.

### GetReadTimeoutOk

`func (o *HttpConfig) GetReadTimeoutOk() (*int32, bool)`

GetReadTimeoutOk returns a tuple with the ReadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadTimeout

`func (o *HttpConfig) SetReadTimeout(v int32)`

SetReadTimeout sets ReadTimeout field to given value.

### HasReadTimeout

`func (o *HttpConfig) HasReadTimeout() bool`

HasReadTimeout returns a boolean if a field has been set.

### GetUrl

`func (o *HttpConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HttpConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HttpConfig) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


