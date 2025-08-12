# PaylinkStateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **time.Time** | the date and time of the event. | [optional] 
**Message** | Pointer to **string** | a message associated with the event. | [optional] 
**State** | Pointer to **string** | The name of the event that was actioned. | [optional] 

## Methods

### NewPaylinkStateEvent

`func NewPaylinkStateEvent() *PaylinkStateEvent`

NewPaylinkStateEvent instantiates a new PaylinkStateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkStateEventWithDefaults

`func NewPaylinkStateEventWithDefaults() *PaylinkStateEvent`

NewPaylinkStateEventWithDefaults instantiates a new PaylinkStateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *PaylinkStateEvent) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *PaylinkStateEvent) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *PaylinkStateEvent) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *PaylinkStateEvent) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetMessage

`func (o *PaylinkStateEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PaylinkStateEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PaylinkStateEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PaylinkStateEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetState

`func (o *PaylinkStateEvent) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PaylinkStateEvent) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PaylinkStateEvent) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PaylinkStateEvent) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


