# PaylinkCustomParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryMode** | Pointer to **string** | The type of entry mode. A value of &#39;pre&#39; will pre-render the custom parameter before the payment screen. Any other value will result in the custom parameter being displayed on the payment screen. | [optional] 
**FieldType** | Pointer to **string** | the type of html5 field, defaults to &#39;text&#39;. Other options are &#39;dob&#39; for a date of birth series of select list entry. | [optional] 
**Group** | Pointer to **string** | a group the parameter is linked with, allows for grouping with a title. | [optional] 
**Label** | Pointer to **string** | a label to show alongside the input. | [optional] 
**Locked** | Pointer to **bool** | whether the parameter is locked from entry. | [optional] 
**Name** | **string** | the name of the custom parameter used to converse with the submitter. | 
**Order** | Pointer to **int32** | an index order for the parameter. | [optional] 
**Pattern** | Pointer to **string** | a regex pattern to validate the custom parameter with. | [optional] 
**Placeholder** | Pointer to **string** | a placehold value to display in the input. | [optional] 
**Required** | Pointer to **bool** | whether the field is required. | [optional] 
**Value** | Pointer to **string** | a default value for the field. | [optional] 

## Methods

### NewPaylinkCustomParam

`func NewPaylinkCustomParam(name string, ) *PaylinkCustomParam`

NewPaylinkCustomParam instantiates a new PaylinkCustomParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkCustomParamWithDefaults

`func NewPaylinkCustomParamWithDefaults() *PaylinkCustomParam`

NewPaylinkCustomParamWithDefaults instantiates a new PaylinkCustomParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryMode

`func (o *PaylinkCustomParam) GetEntryMode() string`

GetEntryMode returns the EntryMode field if non-nil, zero value otherwise.

### GetEntryModeOk

`func (o *PaylinkCustomParam) GetEntryModeOk() (*string, bool)`

GetEntryModeOk returns a tuple with the EntryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryMode

`func (o *PaylinkCustomParam) SetEntryMode(v string)`

SetEntryMode sets EntryMode field to given value.

### HasEntryMode

`func (o *PaylinkCustomParam) HasEntryMode() bool`

HasEntryMode returns a boolean if a field has been set.

### GetFieldType

`func (o *PaylinkCustomParam) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *PaylinkCustomParam) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *PaylinkCustomParam) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *PaylinkCustomParam) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetGroup

`func (o *PaylinkCustomParam) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PaylinkCustomParam) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PaylinkCustomParam) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PaylinkCustomParam) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetLabel

`func (o *PaylinkCustomParam) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PaylinkCustomParam) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PaylinkCustomParam) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PaylinkCustomParam) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLocked

`func (o *PaylinkCustomParam) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *PaylinkCustomParam) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *PaylinkCustomParam) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *PaylinkCustomParam) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetName

`func (o *PaylinkCustomParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaylinkCustomParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaylinkCustomParam) SetName(v string)`

SetName sets Name field to given value.


### GetOrder

`func (o *PaylinkCustomParam) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PaylinkCustomParam) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PaylinkCustomParam) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PaylinkCustomParam) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPattern

`func (o *PaylinkCustomParam) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *PaylinkCustomParam) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *PaylinkCustomParam) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *PaylinkCustomParam) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetPlaceholder

`func (o *PaylinkCustomParam) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *PaylinkCustomParam) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *PaylinkCustomParam) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *PaylinkCustomParam) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetRequired

`func (o *PaylinkCustomParam) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *PaylinkCustomParam) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *PaylinkCustomParam) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *PaylinkCustomParam) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetValue

`func (o *PaylinkCustomParam) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PaylinkCustomParam) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PaylinkCustomParam) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PaylinkCustomParam) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


