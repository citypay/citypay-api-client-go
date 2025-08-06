# WebHookUnsubscribeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clientid** | **string** | The client id that the hook is registered for.  | 
**WebHookId** | **string** | The webhook id that is to be removed.  | 

## Methods

### NewWebHookUnsubscribeRequest

`func NewWebHookUnsubscribeRequest(clientid string, webHookId string, ) *WebHookUnsubscribeRequest`

NewWebHookUnsubscribeRequest instantiates a new WebHookUnsubscribeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebHookUnsubscribeRequestWithDefaults

`func NewWebHookUnsubscribeRequestWithDefaults() *WebHookUnsubscribeRequest`

NewWebHookUnsubscribeRequestWithDefaults instantiates a new WebHookUnsubscribeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientid

`func (o *WebHookUnsubscribeRequest) GetClientid() string`

GetClientid returns the Clientid field if non-nil, zero value otherwise.

### GetClientidOk

`func (o *WebHookUnsubscribeRequest) GetClientidOk() (*string, bool)`

GetClientidOk returns a tuple with the Clientid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientid

`func (o *WebHookUnsubscribeRequest) SetClientid(v string)`

SetClientid sets Clientid field to given value.


### GetWebHookId

`func (o *WebHookUnsubscribeRequest) GetWebHookId() string`

GetWebHookId returns the WebHookId field if non-nil, zero value otherwise.

### GetWebHookIdOk

`func (o *WebHookUnsubscribeRequest) GetWebHookIdOk() (*string, bool)`

GetWebHookIdOk returns a tuple with the WebHookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHookId

`func (o *WebHookUnsubscribeRequest) SetWebHookId(v string)`

SetWebHookId sets WebHookId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


