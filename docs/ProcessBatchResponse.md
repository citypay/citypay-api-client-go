# ProcessBatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Information regarding the processing request. | [optional] 
**Valid** | **bool** | true if the request has been accepted for processing and is valid. | 

## Methods

### NewProcessBatchResponse

`func NewProcessBatchResponse(valid bool, ) *ProcessBatchResponse`

NewProcessBatchResponse instantiates a new ProcessBatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessBatchResponseWithDefaults

`func NewProcessBatchResponseWithDefaults() *ProcessBatchResponse`

NewProcessBatchResponseWithDefaults instantiates a new ProcessBatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ProcessBatchResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProcessBatchResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProcessBatchResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProcessBatchResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetValid

`func (o *ProcessBatchResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ProcessBatchResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ProcessBatchResponse) SetValid(v bool)`

SetValid sets Valid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


