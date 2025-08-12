# Adjustments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accumulate** | Pointer to **string** | How adjustments are accumulated and therefore applied.  **None (Default)**: Only the last applicable adjustment is applied. The system ignores previous adjustments, and only the effect of the final adjustment is considered. Use Case: Use this mode when you want the final transaction amount to reflect only the last adjustment in the sequence, without any cumulative effect from prior adjustments.  **AccumulateBase**: Applies each adjustment independently to the original base amount of the transaction, regardless of any previous adjustments. The effects of all adjustments are then combined to produce the final amount. Use Case: This mode is useful when each adjustment should be applied as if it were the only adjustment, but their effects are accumulated together.  **AccumulatePrevious**: Applies each adjustment sequentially based on the amount resulting from the previous adjustment. This creates a cumulative effect where each adjustment builds upon the last one. Use Case: This mode is ideal when you need the final amount to reflect the cumulative effect of all adjustments in the order they are applied.  **AccumulateBaseOver**: The AccumulateBaseOver mode compares the effect of applying an adjustment to the original base amount with the result of the previously accumulated adjustments. The system then applies whichever adjustment produces a greater final amount. Use Case: This mode is useful when you want to ensure that the most impactful adjustment is applied, whether it comes from the base or the accumulated amount.  | [optional] 
**Adjustment** | **string** | The type of adjustment, valid values are &#x60;surcharge&#x60; or &#x60;discount&#x60;. | 
**Amount** | Pointer to **int32** | For fixed-amount adjustments, an amount to be discounted or surcharged. | [optional] 
**Conditions** | Pointer to [**AdjustmentCondition**](AdjustmentCondition.md) |  | [optional] 
**Description** | Pointer to **string** | A brief description of the adjustment, explaining its purpose or the conditions under which it is applied. For example. - Late Payment Fee - £15 fee for expedited processing on the same day - 5% discount for payments made within 5 days - 15% discount for first-time customers - 10% discount for loyalty program members.  | [optional] 
**Percentage** | Pointer to **float64** | For percentage-based adjustments, the percentage amount to be discounted or surcharged. | [optional] 

## Methods

### NewAdjustments

`func NewAdjustments(adjustment string, ) *Adjustments`

NewAdjustments instantiates a new Adjustments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdjustmentsWithDefaults

`func NewAdjustmentsWithDefaults() *Adjustments`

NewAdjustmentsWithDefaults instantiates a new Adjustments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccumulate

`func (o *Adjustments) GetAccumulate() string`

GetAccumulate returns the Accumulate field if non-nil, zero value otherwise.

### GetAccumulateOk

`func (o *Adjustments) GetAccumulateOk() (*string, bool)`

GetAccumulateOk returns a tuple with the Accumulate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulate

`func (o *Adjustments) SetAccumulate(v string)`

SetAccumulate sets Accumulate field to given value.

### HasAccumulate

`func (o *Adjustments) HasAccumulate() bool`

HasAccumulate returns a boolean if a field has been set.

### GetAdjustment

`func (o *Adjustments) GetAdjustment() string`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *Adjustments) GetAdjustmentOk() (*string, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *Adjustments) SetAdjustment(v string)`

SetAdjustment sets Adjustment field to given value.


### GetAmount

`func (o *Adjustments) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Adjustments) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Adjustments) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Adjustments) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetConditions

`func (o *Adjustments) GetConditions() AdjustmentCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *Adjustments) GetConditionsOk() (*AdjustmentCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *Adjustments) SetConditions(v AdjustmentCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *Adjustments) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetDescription

`func (o *Adjustments) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Adjustments) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Adjustments) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Adjustments) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPercentage

`func (o *Adjustments) GetPercentage() float64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *Adjustments) GetPercentageOk() (*float64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *Adjustments) SetPercentage(v float64)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *Adjustments) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


