# WebHookSubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to **[]string** |  | [optional] 
**Clientid** | **string** | The client id that the hook is being registered for.  | 
**Live** | Pointer to **bool** | Specifies if the key is to be used for production. Defaults to false.  | [optional] 
**MerchantId** | Pointer to **[]int32** |  | [optional] 
**Triggers** | Pointer to **[]string** |  | [optional] 

## Methods

### NewWebHookSubscriptionRequest

`func NewWebHookSubscriptionRequest(clientid string, ) *WebHookSubscriptionRequest`

NewWebHookSubscriptionRequest instantiates a new WebHookSubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebHookSubscriptionRequestWithDefaults

`func NewWebHookSubscriptionRequestWithDefaults() *WebHookSubscriptionRequest`

NewWebHookSubscriptionRequestWithDefaults instantiates a new WebHookSubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *WebHookSubscriptionRequest) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *WebHookSubscriptionRequest) GetChannelsOk() (*[]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *WebHookSubscriptionRequest) SetChannels(v []string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *WebHookSubscriptionRequest) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetClientid

`func (o *WebHookSubscriptionRequest) GetClientid() string`

GetClientid returns the Clientid field if non-nil, zero value otherwise.

### GetClientidOk

`func (o *WebHookSubscriptionRequest) GetClientidOk() (*string, bool)`

GetClientidOk returns a tuple with the Clientid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientid

`func (o *WebHookSubscriptionRequest) SetClientid(v string)`

SetClientid sets Clientid field to given value.


### GetLive

`func (o *WebHookSubscriptionRequest) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *WebHookSubscriptionRequest) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *WebHookSubscriptionRequest) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *WebHookSubscriptionRequest) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetMerchantId

`func (o *WebHookSubscriptionRequest) GetMerchantId() []int32`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *WebHookSubscriptionRequest) GetMerchantIdOk() (*[]int32, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *WebHookSubscriptionRequest) SetMerchantId(v []int32)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *WebHookSubscriptionRequest) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetTriggers

`func (o *WebHookSubscriptionRequest) GetTriggers() []string`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *WebHookSubscriptionRequest) GetTriggersOk() (*[]string, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *WebHookSubscriptionRequest) SetTriggers(v []string)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *WebHookSubscriptionRequest) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


