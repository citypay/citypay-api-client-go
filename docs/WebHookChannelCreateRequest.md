# WebHookChannelCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelName** | **string** | The name of the channel we are creating.  | 
**Clientid** | **string** | The client id that the hook is being registered for.  | 
**Config** | [**WebHookChannelCreateRequestConfig**](WebHookChannelCreateRequestConfig.md) |  | 
**EndpointId** | **string** | The id of the endpoint being used. The channel configuration is dependant upon the endpoint type.  | 

## Methods

### NewWebHookChannelCreateRequest

`func NewWebHookChannelCreateRequest(channelName string, clientid string, config WebHookChannelCreateRequestConfig, endpointId string, ) *WebHookChannelCreateRequest`

NewWebHookChannelCreateRequest instantiates a new WebHookChannelCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebHookChannelCreateRequestWithDefaults

`func NewWebHookChannelCreateRequestWithDefaults() *WebHookChannelCreateRequest`

NewWebHookChannelCreateRequestWithDefaults instantiates a new WebHookChannelCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelName

`func (o *WebHookChannelCreateRequest) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *WebHookChannelCreateRequest) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *WebHookChannelCreateRequest) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.


### GetClientid

`func (o *WebHookChannelCreateRequest) GetClientid() string`

GetClientid returns the Clientid field if non-nil, zero value otherwise.

### GetClientidOk

`func (o *WebHookChannelCreateRequest) GetClientidOk() (*string, bool)`

GetClientidOk returns a tuple with the Clientid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientid

`func (o *WebHookChannelCreateRequest) SetClientid(v string)`

SetClientid sets Clientid field to given value.


### GetConfig

`func (o *WebHookChannelCreateRequest) GetConfig() WebHookChannelCreateRequestConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *WebHookChannelCreateRequest) GetConfigOk() (*WebHookChannelCreateRequestConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *WebHookChannelCreateRequest) SetConfig(v WebHookChannelCreateRequestConfig)`

SetConfig sets Config field to given value.


### GetEndpointId

`func (o *WebHookChannelCreateRequest) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *WebHookChannelCreateRequest) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *WebHookChannelCreateRequest) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


