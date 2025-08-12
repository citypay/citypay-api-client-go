# AdjustmentCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Anchor** | Pointer to **string** | Specifies the reference point (\&quot;anchor\&quot;) and optional duration/range for when an adjustment  (such as a surcharge or discount) becomes effective in relation to the payment intent&#39;s lifecycle.   The period for which an adjustment is effective is determined by combining:    - An \&quot;anchor\&quot;: a reference event, timestamp, or explicit date.    - A &#x60;duration&#x60;: an integer offset in days, whose meaning depends on the anchor.    - Optionally an explicit &#x60;startDate&#x60; and/or &#x60;endDate&#x60;, which take precedence.   Anchor types:   - Explicit Dates:    If both &#x60;startDate&#x60; and &#x60;endDate&#x60; are provided, these take priority as the exact active range.   - after_creation:      Starts &#x60;duration&#x60; days after the intent&#39;s creation date.      If an &#x60;endDate&#x60; is supplied, it is used as the end of the range.      If not, the adjustment remains open-ended (no explicit end).   - after_due_date:      Starts &#x60;duration&#x60; days after the intent&#39;s due date (if available).      If an &#x60;endDate&#x60; is specified, it is used as the end of the range.      Otherwise, the adjustment remains open-ended.   - on_due_date:      Starts on the due date.      If an &#x60;endDate&#x60; is specified and is after the due date, it is used as the end.      Otherwise, the end is computed as the due date plus &#x60;duration&#x60; days.   - until_due_date:      If &#x60;duration &gt;&#x3D; 0&#x60;, starts &#x60;duration&#x60; days after creation, and ends at due date.      If &#x60;duration &lt; 0&#x60;, starts at creation and ends &#x60;abs(duration)&#x60; days before due date.   - before_expiry:      Starts at expiry date minus &#x60;duration&#x60; days (if expiry date is available).      Range extends up to expiry (default) or effectively open-ended.      If an &#x60;endDate&#x60; is specified and before expiry, the range is limited accordingly.   - yyyy-MM-dd (a valid ISO date string):      Effective on a specific date (midnight, or explicit time if provided), plus optional &#x60;duration&#x60; days.      In this case, the range is a single day, unless otherwise specified.    Example usages:    - Adjust 2 days after due: &#x60;anchor &#x3D; \&quot;after_due_date\&quot;, duration &#x3D; 2&#x60;    - Early-bird: &#x60;anchor &#x3D; \&quot;until_due_date\&quot;, duration &#x3D; -3&#x60; // from creation until 3 days before due    - From creation until due: &#x60;anchor &#x3D; \&quot;until_due_date\&quot;, duration &#x3D; 0&#x60;    - Discount for a single date: &#x60;anchor &#x3D; \&quot;2024-07-01\&quot;&#x60;   Notes:    - If &#x60;startDate&#x60; and &#x60;endDate&#x60; are both set, those take precedence over anchor/duration.    - If only &#x60;endDate&#x60; is specified with an anchor, the range starts at the anchor and ends at that date.    - If dates or anchor-related context are missing, no adjustment will be scheduled.  | [optional] 
**DiscountCode** | Pointer to **string** | The discount code condition ensures that an adjustment is only applied if the correct promotional code is provided during the transaction.  Example: A 10% discount might only apply if the customer enters the discount code SEP24.  | [optional] 
**Duration** | Pointer to **int32** | The duration specifies the time frame in days relative to the anchor point when the adjustment should be applied. | [optional] 
**EndDate** | Pointer to **string** | Define the exact date range within which an adjustment is valid, useful for promotional periods or specific events. | [optional] 
**StartDate** | Pointer to **string** | Define the exact date range within which an adjustment is valid, useful for promotional periods or specific events. | [optional] 

## Methods

### NewAdjustmentCondition

`func NewAdjustmentCondition() *AdjustmentCondition`

NewAdjustmentCondition instantiates a new AdjustmentCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdjustmentConditionWithDefaults

`func NewAdjustmentConditionWithDefaults() *AdjustmentCondition`

NewAdjustmentConditionWithDefaults instantiates a new AdjustmentCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnchor

`func (o *AdjustmentCondition) GetAnchor() string`

GetAnchor returns the Anchor field if non-nil, zero value otherwise.

### GetAnchorOk

`func (o *AdjustmentCondition) GetAnchorOk() (*string, bool)`

GetAnchorOk returns a tuple with the Anchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchor

`func (o *AdjustmentCondition) SetAnchor(v string)`

SetAnchor sets Anchor field to given value.

### HasAnchor

`func (o *AdjustmentCondition) HasAnchor() bool`

HasAnchor returns a boolean if a field has been set.

### GetDiscountCode

`func (o *AdjustmentCondition) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *AdjustmentCondition) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *AdjustmentCondition) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *AdjustmentCondition) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetDuration

`func (o *AdjustmentCondition) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AdjustmentCondition) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AdjustmentCondition) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AdjustmentCondition) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEndDate

`func (o *AdjustmentCondition) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AdjustmentCondition) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AdjustmentCondition) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AdjustmentCondition) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDate

`func (o *AdjustmentCondition) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AdjustmentCondition) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AdjustmentCondition) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AdjustmentCondition) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


