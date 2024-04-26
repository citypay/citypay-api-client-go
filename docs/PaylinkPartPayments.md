# PaylinkPartPayments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **string** | Determines if part payments is enabled. Default is false. | [optional] 
**Floor** | Pointer to **string** | The floor amount specifies a value that the minimum rate cannot go under. If 0 the amount of min rate is applied. | [optional] 
**Max** | Pointer to **string** | a maximum percentage to charge i.e. 90%. | [optional] 
**MaxRate** | Pointer to **string** | a rate as fixed or percentage. | [optional] 
**Min** | Pointer to **string** | a minimum percentage to charge i.e. 10. | [optional] 
**MinRate** | Pointer to **string** | a rate as fixed or percentage. | [optional] 

## Methods

### NewPaylinkPartPayments

`func NewPaylinkPartPayments() *PaylinkPartPayments`

NewPaylinkPartPayments instantiates a new PaylinkPartPayments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkPartPaymentsWithDefaults

`func NewPaylinkPartPaymentsWithDefaults() *PaylinkPartPayments`

NewPaylinkPartPaymentsWithDefaults instantiates a new PaylinkPartPayments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *PaylinkPartPayments) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PaylinkPartPayments) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PaylinkPartPayments) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PaylinkPartPayments) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFloor

`func (o *PaylinkPartPayments) GetFloor() string`

GetFloor returns the Floor field if non-nil, zero value otherwise.

### GetFloorOk

`func (o *PaylinkPartPayments) GetFloorOk() (*string, bool)`

GetFloorOk returns a tuple with the Floor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloor

`func (o *PaylinkPartPayments) SetFloor(v string)`

SetFloor sets Floor field to given value.

### HasFloor

`func (o *PaylinkPartPayments) HasFloor() bool`

HasFloor returns a boolean if a field has been set.

### GetMax

`func (o *PaylinkPartPayments) GetMax() string`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *PaylinkPartPayments) GetMaxOk() (*string, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *PaylinkPartPayments) SetMax(v string)`

SetMax sets Max field to given value.

### HasMax

`func (o *PaylinkPartPayments) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMaxRate

`func (o *PaylinkPartPayments) GetMaxRate() string`

GetMaxRate returns the MaxRate field if non-nil, zero value otherwise.

### GetMaxRateOk

`func (o *PaylinkPartPayments) GetMaxRateOk() (*string, bool)`

GetMaxRateOk returns a tuple with the MaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRate

`func (o *PaylinkPartPayments) SetMaxRate(v string)`

SetMaxRate sets MaxRate field to given value.

### HasMaxRate

`func (o *PaylinkPartPayments) HasMaxRate() bool`

HasMaxRate returns a boolean if a field has been set.

### GetMin

`func (o *PaylinkPartPayments) GetMin() string`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *PaylinkPartPayments) GetMinOk() (*string, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *PaylinkPartPayments) SetMin(v string)`

SetMin sets Min field to given value.

### HasMin

`func (o *PaylinkPartPayments) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMinRate

`func (o *PaylinkPartPayments) GetMinRate() string`

GetMinRate returns the MinRate field if non-nil, zero value otherwise.

### GetMinRateOk

`func (o *PaylinkPartPayments) GetMinRateOk() (*string, bool)`

GetMinRateOk returns a tuple with the MinRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRate

`func (o *PaylinkPartPayments) SetMinRate(v string)`

SetMinRate sets MinRate field to given value.

### HasMinRate

`func (o *PaylinkPartPayments) HasMinRate() bool`

HasMinRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


