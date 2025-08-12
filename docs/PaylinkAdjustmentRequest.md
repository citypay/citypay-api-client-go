# PaylinkAdjustmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | An amount to adjust to. | [optional] 
**Identifier** | Pointer to **string** | An identifier of the original request. | [optional] 
**Reason** | Pointer to **string** | A textual reason for the adjustment. | [optional] 

## Methods

### NewPaylinkAdjustmentRequest

`func NewPaylinkAdjustmentRequest() *PaylinkAdjustmentRequest`

NewPaylinkAdjustmentRequest instantiates a new PaylinkAdjustmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkAdjustmentRequestWithDefaults

`func NewPaylinkAdjustmentRequestWithDefaults() *PaylinkAdjustmentRequest`

NewPaylinkAdjustmentRequestWithDefaults instantiates a new PaylinkAdjustmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PaylinkAdjustmentRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaylinkAdjustmentRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaylinkAdjustmentRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaylinkAdjustmentRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetIdentifier

`func (o *PaylinkAdjustmentRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PaylinkAdjustmentRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PaylinkAdjustmentRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PaylinkAdjustmentRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetReason

`func (o *PaylinkAdjustmentRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PaylinkAdjustmentRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PaylinkAdjustmentRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PaylinkAdjustmentRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


