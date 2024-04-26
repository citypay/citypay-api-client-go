# PaylinkUI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressMandatory** | Pointer to **bool** | whether the address is forced as mandatory. | [optional] 
**FormAutoComplete** | Pointer to **string** | specify the form autocomplete setting, default to on. If set to off the UI will set autocomplete&#x3D;\&quot;off\&quot; on the form level and prevent elements from adding it. | [optional] 
**Ordering** | Pointer to **int32** | the logical ordering of the ui groups. | [optional] 
**PostcodeMandatory** | Pointer to **bool** | whether the postcode is forced as mandatory. | [optional] 

## Methods

### NewPaylinkUI

`func NewPaylinkUI() *PaylinkUI`

NewPaylinkUI instantiates a new PaylinkUI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkUIWithDefaults

`func NewPaylinkUIWithDefaults() *PaylinkUI`

NewPaylinkUIWithDefaults instantiates a new PaylinkUI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressMandatory

`func (o *PaylinkUI) GetAddressMandatory() bool`

GetAddressMandatory returns the AddressMandatory field if non-nil, zero value otherwise.

### GetAddressMandatoryOk

`func (o *PaylinkUI) GetAddressMandatoryOk() (*bool, bool)`

GetAddressMandatoryOk returns a tuple with the AddressMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressMandatory

`func (o *PaylinkUI) SetAddressMandatory(v bool)`

SetAddressMandatory sets AddressMandatory field to given value.

### HasAddressMandatory

`func (o *PaylinkUI) HasAddressMandatory() bool`

HasAddressMandatory returns a boolean if a field has been set.

### GetFormAutoComplete

`func (o *PaylinkUI) GetFormAutoComplete() string`

GetFormAutoComplete returns the FormAutoComplete field if non-nil, zero value otherwise.

### GetFormAutoCompleteOk

`func (o *PaylinkUI) GetFormAutoCompleteOk() (*string, bool)`

GetFormAutoCompleteOk returns a tuple with the FormAutoComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormAutoComplete

`func (o *PaylinkUI) SetFormAutoComplete(v string)`

SetFormAutoComplete sets FormAutoComplete field to given value.

### HasFormAutoComplete

`func (o *PaylinkUI) HasFormAutoComplete() bool`

HasFormAutoComplete returns a boolean if a field has been set.

### GetOrdering

`func (o *PaylinkUI) GetOrdering() int32`

GetOrdering returns the Ordering field if non-nil, zero value otherwise.

### GetOrderingOk

`func (o *PaylinkUI) GetOrderingOk() (*int32, bool)`

GetOrderingOk returns a tuple with the Ordering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdering

`func (o *PaylinkUI) SetOrdering(v int32)`

SetOrdering sets Ordering field to given value.

### HasOrdering

`func (o *PaylinkUI) HasOrdering() bool`

HasOrdering returns a boolean if a field has been set.

### GetPostcodeMandatory

`func (o *PaylinkUI) GetPostcodeMandatory() bool`

GetPostcodeMandatory returns the PostcodeMandatory field if non-nil, zero value otherwise.

### GetPostcodeMandatoryOk

`func (o *PaylinkUI) GetPostcodeMandatoryOk() (*bool, bool)`

GetPostcodeMandatoryOk returns a tuple with the PostcodeMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcodeMandatory

`func (o *PaylinkUI) SetPostcodeMandatory(v bool)`

SetPostcodeMandatory sets PostcodeMandatory field to given value.

### HasPostcodeMandatory

`func (o *PaylinkUI) HasPostcodeMandatory() bool`

HasPostcodeMandatory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


