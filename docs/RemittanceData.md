# RemittanceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **time.Time** | Represents the date and time when the remittance was processed. This timestamp follows the ISO 8601 format for datetime representation. | [optional] 
**NetAmount** | Pointer to **int32** | Represents the net amount after accounting for refunds. This is calculated as SalesAmount - RefundAmount and expressed in the smallest currency unit. | [optional] 
**RefundAmount** | Pointer to **int32** | The total amount refunded to customers. | [optional] 
**RefundCount** | Pointer to **int32** | The total number of refund transactions processed. This figure helps in understanding the frequency of refunds relative to sales. | [optional] 
**SalesAmount** | Pointer to **int32** | The total monetary amount of sales transactions. | [optional] 
**SalesCount** | Pointer to **int32** | Indicates the total number of sales transactions that occurred. This count provides insight into the volume of sales. | [optional] 

## Methods

### NewRemittanceData

`func NewRemittanceData() *RemittanceData`

NewRemittanceData instantiates a new RemittanceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemittanceDataWithDefaults

`func NewRemittanceDataWithDefaults() *RemittanceData`

NewRemittanceDataWithDefaults instantiates a new RemittanceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *RemittanceData) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *RemittanceData) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *RemittanceData) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *RemittanceData) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetNetAmount

`func (o *RemittanceData) GetNetAmount() int32`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *RemittanceData) GetNetAmountOk() (*int32, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *RemittanceData) SetNetAmount(v int32)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *RemittanceData) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### GetRefundAmount

`func (o *RemittanceData) GetRefundAmount() int32`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *RemittanceData) GetRefundAmountOk() (*int32, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *RemittanceData) SetRefundAmount(v int32)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *RemittanceData) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### GetRefundCount

`func (o *RemittanceData) GetRefundCount() int32`

GetRefundCount returns the RefundCount field if non-nil, zero value otherwise.

### GetRefundCountOk

`func (o *RemittanceData) GetRefundCountOk() (*int32, bool)`

GetRefundCountOk returns a tuple with the RefundCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundCount

`func (o *RemittanceData) SetRefundCount(v int32)`

SetRefundCount sets RefundCount field to given value.

### HasRefundCount

`func (o *RemittanceData) HasRefundCount() bool`

HasRefundCount returns a boolean if a field has been set.

### GetSalesAmount

`func (o *RemittanceData) GetSalesAmount() int32`

GetSalesAmount returns the SalesAmount field if non-nil, zero value otherwise.

### GetSalesAmountOk

`func (o *RemittanceData) GetSalesAmountOk() (*int32, bool)`

GetSalesAmountOk returns a tuple with the SalesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesAmount

`func (o *RemittanceData) SetSalesAmount(v int32)`

SetSalesAmount sets SalesAmount field to given value.

### HasSalesAmount

`func (o *RemittanceData) HasSalesAmount() bool`

HasSalesAmount returns a boolean if a field has been set.

### GetSalesCount

`func (o *RemittanceData) GetSalesCount() int32`

GetSalesCount returns the SalesCount field if non-nil, zero value otherwise.

### GetSalesCountOk

`func (o *RemittanceData) GetSalesCountOk() (*int32, bool)`

GetSalesCountOk returns a tuple with the SalesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesCount

`func (o *RemittanceData) SetSalesCount(v int32)`

SetSalesCount sets SalesCount field to given value.

### HasSalesCount

`func (o *RemittanceData) HasSalesCount() bool`

HasSalesCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


