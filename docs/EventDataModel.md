# EventDataModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventEndDate** | Pointer to **string** | The date when the event ends in ISO format (yyyy-MM-dd). | [optional] 
**EventId** | Pointer to **string** | An id of the event. | [optional] 
**EventOrganiserId** | Pointer to **string** | An id of the event organiser. | [optional] 
**EventStartDate** | Pointer to **string** | The date when the event starts in ISO format (yyyy-MM-dd). | [optional] 
**PaymentType** | Pointer to **string** | The type of payment such as &#x60;deposit&#x60; or &#x60;balance&#x60;. | [optional] 

## Methods

### NewEventDataModel

`func NewEventDataModel() *EventDataModel`

NewEventDataModel instantiates a new EventDataModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDataModelWithDefaults

`func NewEventDataModelWithDefaults() *EventDataModel`

NewEventDataModelWithDefaults instantiates a new EventDataModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventEndDate

`func (o *EventDataModel) GetEventEndDate() string`

GetEventEndDate returns the EventEndDate field if non-nil, zero value otherwise.

### GetEventEndDateOk

`func (o *EventDataModel) GetEventEndDateOk() (*string, bool)`

GetEventEndDateOk returns a tuple with the EventEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventEndDate

`func (o *EventDataModel) SetEventEndDate(v string)`

SetEventEndDate sets EventEndDate field to given value.

### HasEventEndDate

`func (o *EventDataModel) HasEventEndDate() bool`

HasEventEndDate returns a boolean if a field has been set.

### GetEventId

`func (o *EventDataModel) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *EventDataModel) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *EventDataModel) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *EventDataModel) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventOrganiserId

`func (o *EventDataModel) GetEventOrganiserId() string`

GetEventOrganiserId returns the EventOrganiserId field if non-nil, zero value otherwise.

### GetEventOrganiserIdOk

`func (o *EventDataModel) GetEventOrganiserIdOk() (*string, bool)`

GetEventOrganiserIdOk returns a tuple with the EventOrganiserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventOrganiserId

`func (o *EventDataModel) SetEventOrganiserId(v string)`

SetEventOrganiserId sets EventOrganiserId field to given value.

### HasEventOrganiserId

`func (o *EventDataModel) HasEventOrganiserId() bool`

HasEventOrganiserId returns a boolean if a field has been set.

### GetEventStartDate

`func (o *EventDataModel) GetEventStartDate() string`

GetEventStartDate returns the EventStartDate field if non-nil, zero value otherwise.

### GetEventStartDateOk

`func (o *EventDataModel) GetEventStartDateOk() (*string, bool)`

GetEventStartDateOk returns a tuple with the EventStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventStartDate

`func (o *EventDataModel) SetEventStartDate(v string)`

SetEventStartDate sets EventStartDate field to given value.

### HasEventStartDate

`func (o *EventDataModel) HasEventStartDate() bool`

HasEventStartDate returns a boolean if a field has been set.

### GetPaymentType

`func (o *EventDataModel) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *EventDataModel) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *EventDataModel) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *EventDataModel) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


