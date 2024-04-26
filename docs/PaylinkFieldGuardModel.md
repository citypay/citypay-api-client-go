# PaylinkFieldGuardModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldType** | Pointer to **string** | A type of HTML element that should be displayed such as text, password, url. Any HTML5 input type value may be supplied.  If a value of &#x60;date&#x60; is supplied the value format should be an ISO format YYYY-MM-DD format date i.e. 2024-03-01 If a value of &#x60;datetime-local&#x60; is supplied, the value format should be an ISO format YYYY-MM-DDTHH:mm i.e. 2024-06-01T19:30.  | [optional] 
**Label** | Pointer to **string** | A label for the field guard to display on the authentication page. | [optional] 
**Maxlen** | Pointer to **int32** | A maximum length of any value supplied in the field guard form. Used for validating entry. | [optional] 
**Minlen** | Pointer to **int32** | A minimum length of any value supplied in the field guard form. Used for validating entry. | [optional] 
**Name** | Pointer to **string** | A field name which is used to refer to a field which is guarded. | [optional] 
**Regex** | Pointer to **string** | A JavaScript regular expression value which can be used to validate the data provided in the field guard entry form. Used for validating entry. | [optional] 
**Value** | Pointer to **string** | A value directly associated with the field guard. Any value provided at this level will be considered as sensitive and not logged. | [optional] 

## Methods

### NewPaylinkFieldGuardModel

`func NewPaylinkFieldGuardModel() *PaylinkFieldGuardModel`

NewPaylinkFieldGuardModel instantiates a new PaylinkFieldGuardModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkFieldGuardModelWithDefaults

`func NewPaylinkFieldGuardModelWithDefaults() *PaylinkFieldGuardModel`

NewPaylinkFieldGuardModelWithDefaults instantiates a new PaylinkFieldGuardModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldType

`func (o *PaylinkFieldGuardModel) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *PaylinkFieldGuardModel) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *PaylinkFieldGuardModel) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *PaylinkFieldGuardModel) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetLabel

`func (o *PaylinkFieldGuardModel) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PaylinkFieldGuardModel) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PaylinkFieldGuardModel) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PaylinkFieldGuardModel) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMaxlen

`func (o *PaylinkFieldGuardModel) GetMaxlen() int32`

GetMaxlen returns the Maxlen field if non-nil, zero value otherwise.

### GetMaxlenOk

`func (o *PaylinkFieldGuardModel) GetMaxlenOk() (*int32, bool)`

GetMaxlenOk returns a tuple with the Maxlen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxlen

`func (o *PaylinkFieldGuardModel) SetMaxlen(v int32)`

SetMaxlen sets Maxlen field to given value.

### HasMaxlen

`func (o *PaylinkFieldGuardModel) HasMaxlen() bool`

HasMaxlen returns a boolean if a field has been set.

### GetMinlen

`func (o *PaylinkFieldGuardModel) GetMinlen() int32`

GetMinlen returns the Minlen field if non-nil, zero value otherwise.

### GetMinlenOk

`func (o *PaylinkFieldGuardModel) GetMinlenOk() (*int32, bool)`

GetMinlenOk returns a tuple with the Minlen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinlen

`func (o *PaylinkFieldGuardModel) SetMinlen(v int32)`

SetMinlen sets Minlen field to given value.

### HasMinlen

`func (o *PaylinkFieldGuardModel) HasMinlen() bool`

HasMinlen returns a boolean if a field has been set.

### GetName

`func (o *PaylinkFieldGuardModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaylinkFieldGuardModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaylinkFieldGuardModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaylinkFieldGuardModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegex

`func (o *PaylinkFieldGuardModel) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *PaylinkFieldGuardModel) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *PaylinkFieldGuardModel) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *PaylinkFieldGuardModel) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetValue

`func (o *PaylinkFieldGuardModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PaylinkFieldGuardModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PaylinkFieldGuardModel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PaylinkFieldGuardModel) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


