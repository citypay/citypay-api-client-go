# RemittedClientData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batches** | [**[]MerchantBatchResponse**](MerchantBatchResponse.md) |  | 
**Clientid** | Pointer to **string** | The client id that the remittance data is for. | [optional] 
**Date** | Pointer to **string** | The date of the remittance. | [optional] 
**DateCreated** | Pointer to **time.Time** | The date time that the remittance was created. | [optional] 
**NetAmount** | Pointer to **int32** | Represents the net amount after accounting for refunds. This is calculated as SalesAmount - RefundAmount and expressed in the smallest currency unit. | [optional] 
**ProcessedAmount** | Pointer to **int32** | The total monetary amount processed consisting of sale and refund transactions. | [optional] 
**ProcessedCount** | Pointer to **int32** | Indicates the total number of sales and refund transactions that occurred. This count provides insight into the volume of processing. | [optional] 
**RefundAmount** | Pointer to **int32** | The total amount refunded to customers. | [optional] 
**RefundCount** | Pointer to **int32** | The total number of refund transactions processed. This figure helps in understanding the frequency of refunds relative to sales. | [optional] 
**Remittances** | [**[]RemittanceData**](RemittanceData.md) |  | 
**SalesAmount** | Pointer to **int32** | The total monetary amount of sales transactions. | [optional] 
**SalesCount** | Pointer to **int32** | Indicates the total number of sales transactions that occurred. This count provides insight into the volume of sales. | [optional] 
**SettlementImplementation** | Pointer to **string** | The name of the implementation. | [optional] 
**Uuid** | Pointer to **string** | The uuid of the settlement file processed on. | [optional] 

## Methods

### NewRemittedClientData

`func NewRemittedClientData(batches []MerchantBatchResponse, remittances []RemittanceData, ) *RemittedClientData`

NewRemittedClientData instantiates a new RemittedClientData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemittedClientDataWithDefaults

`func NewRemittedClientDataWithDefaults() *RemittedClientData`

NewRemittedClientDataWithDefaults instantiates a new RemittedClientData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatches

`func (o *RemittedClientData) GetBatches() []MerchantBatchResponse`

GetBatches returns the Batches field if non-nil, zero value otherwise.

### GetBatchesOk

`func (o *RemittedClientData) GetBatchesOk() (*[]MerchantBatchResponse, bool)`

GetBatchesOk returns a tuple with the Batches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatches

`func (o *RemittedClientData) SetBatches(v []MerchantBatchResponse)`

SetBatches sets Batches field to given value.


### GetClientid

`func (o *RemittedClientData) GetClientid() string`

GetClientid returns the Clientid field if non-nil, zero value otherwise.

### GetClientidOk

`func (o *RemittedClientData) GetClientidOk() (*string, bool)`

GetClientidOk returns a tuple with the Clientid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientid

`func (o *RemittedClientData) SetClientid(v string)`

SetClientid sets Clientid field to given value.

### HasClientid

`func (o *RemittedClientData) HasClientid() bool`

HasClientid returns a boolean if a field has been set.

### GetDate

`func (o *RemittedClientData) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *RemittedClientData) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *RemittedClientData) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *RemittedClientData) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDateCreated

`func (o *RemittedClientData) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *RemittedClientData) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *RemittedClientData) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *RemittedClientData) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetNetAmount

`func (o *RemittedClientData) GetNetAmount() int32`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *RemittedClientData) GetNetAmountOk() (*int32, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *RemittedClientData) SetNetAmount(v int32)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *RemittedClientData) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### GetProcessedAmount

`func (o *RemittedClientData) GetProcessedAmount() int32`

GetProcessedAmount returns the ProcessedAmount field if non-nil, zero value otherwise.

### GetProcessedAmountOk

`func (o *RemittedClientData) GetProcessedAmountOk() (*int32, bool)`

GetProcessedAmountOk returns a tuple with the ProcessedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAmount

`func (o *RemittedClientData) SetProcessedAmount(v int32)`

SetProcessedAmount sets ProcessedAmount field to given value.

### HasProcessedAmount

`func (o *RemittedClientData) HasProcessedAmount() bool`

HasProcessedAmount returns a boolean if a field has been set.

### GetProcessedCount

`func (o *RemittedClientData) GetProcessedCount() int32`

GetProcessedCount returns the ProcessedCount field if non-nil, zero value otherwise.

### GetProcessedCountOk

`func (o *RemittedClientData) GetProcessedCountOk() (*int32, bool)`

GetProcessedCountOk returns a tuple with the ProcessedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedCount

`func (o *RemittedClientData) SetProcessedCount(v int32)`

SetProcessedCount sets ProcessedCount field to given value.

### HasProcessedCount

`func (o *RemittedClientData) HasProcessedCount() bool`

HasProcessedCount returns a boolean if a field has been set.

### GetRefundAmount

`func (o *RemittedClientData) GetRefundAmount() int32`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *RemittedClientData) GetRefundAmountOk() (*int32, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *RemittedClientData) SetRefundAmount(v int32)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *RemittedClientData) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### GetRefundCount

`func (o *RemittedClientData) GetRefundCount() int32`

GetRefundCount returns the RefundCount field if non-nil, zero value otherwise.

### GetRefundCountOk

`func (o *RemittedClientData) GetRefundCountOk() (*int32, bool)`

GetRefundCountOk returns a tuple with the RefundCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundCount

`func (o *RemittedClientData) SetRefundCount(v int32)`

SetRefundCount sets RefundCount field to given value.

### HasRefundCount

`func (o *RemittedClientData) HasRefundCount() bool`

HasRefundCount returns a boolean if a field has been set.

### GetRemittances

`func (o *RemittedClientData) GetRemittances() []RemittanceData`

GetRemittances returns the Remittances field if non-nil, zero value otherwise.

### GetRemittancesOk

`func (o *RemittedClientData) GetRemittancesOk() (*[]RemittanceData, bool)`

GetRemittancesOk returns a tuple with the Remittances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemittances

`func (o *RemittedClientData) SetRemittances(v []RemittanceData)`

SetRemittances sets Remittances field to given value.


### GetSalesAmount

`func (o *RemittedClientData) GetSalesAmount() int32`

GetSalesAmount returns the SalesAmount field if non-nil, zero value otherwise.

### GetSalesAmountOk

`func (o *RemittedClientData) GetSalesAmountOk() (*int32, bool)`

GetSalesAmountOk returns a tuple with the SalesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesAmount

`func (o *RemittedClientData) SetSalesAmount(v int32)`

SetSalesAmount sets SalesAmount field to given value.

### HasSalesAmount

`func (o *RemittedClientData) HasSalesAmount() bool`

HasSalesAmount returns a boolean if a field has been set.

### GetSalesCount

`func (o *RemittedClientData) GetSalesCount() int32`

GetSalesCount returns the SalesCount field if non-nil, zero value otherwise.

### GetSalesCountOk

`func (o *RemittedClientData) GetSalesCountOk() (*int32, bool)`

GetSalesCountOk returns a tuple with the SalesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesCount

`func (o *RemittedClientData) SetSalesCount(v int32)`

SetSalesCount sets SalesCount field to given value.

### HasSalesCount

`func (o *RemittedClientData) HasSalesCount() bool`

HasSalesCount returns a boolean if a field has been set.

### GetSettlementImplementation

`func (o *RemittedClientData) GetSettlementImplementation() string`

GetSettlementImplementation returns the SettlementImplementation field if non-nil, zero value otherwise.

### GetSettlementImplementationOk

`func (o *RemittedClientData) GetSettlementImplementationOk() (*string, bool)`

GetSettlementImplementationOk returns a tuple with the SettlementImplementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementImplementation

`func (o *RemittedClientData) SetSettlementImplementation(v string)`

SetSettlementImplementation sets SettlementImplementation field to given value.

### HasSettlementImplementation

`func (o *RemittedClientData) HasSettlementImplementation() bool`

HasSettlementImplementation returns a boolean if a field has been set.

### GetUuid

`func (o *RemittedClientData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RemittedClientData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RemittedClientData) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RemittedClientData) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


