# MerchantBatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchClosed** | Pointer to **time.Time** | The date and time when the batch was closed. This is represented in ISO 8601 format (e.g., YYYY-MM-DDTHH:MM:SSZ) and indicates when the batch processing was completed. | [optional] 
**BatchNo** | Pointer to **string** | The incremental identifier of the batch. This number is used to track and reference the batch in subsequent operations or inquiries. | [optional] 
**BatchStatus** | Pointer to **string** | A descriptive string detailing the current status of the batch. This status provides a human-readable explanation of the batch&#39;s processing state. | [optional] 
**BatchStatusCode** | Pointer to **string** | A batch status code that represents the processing state of the batch. Batches will be one of  - &#39;O&#39; defining the batch is open for settlement and not yet settled  - &#39;X&#39; defines that the batch is external to our systems and managed elsewhere  - &#39;C&#39; defines that the batch is cancelled and not settled  - &#39;S&#39; defines that the batch has been settled and remitted.  | [optional] 
**Currency** | Pointer to **string** | The currency of the batch. | [optional] 
**Merchantid** | Pointer to **int32** | The Merchant ID (MID) associated with the batch. This identifier specifies which merchant account the batch was processed for, linking transactions to the merchant. | [optional] 
**NetSummary** | Pointer to [**NetSummaryResponse**](NetSummaryResponse.md) |  | [optional] 

## Methods

### NewMerchantBatchResponse

`func NewMerchantBatchResponse() *MerchantBatchResponse`

NewMerchantBatchResponse instantiates a new MerchantBatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantBatchResponseWithDefaults

`func NewMerchantBatchResponseWithDefaults() *MerchantBatchResponse`

NewMerchantBatchResponseWithDefaults instantiates a new MerchantBatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchClosed

`func (o *MerchantBatchResponse) GetBatchClosed() time.Time`

GetBatchClosed returns the BatchClosed field if non-nil, zero value otherwise.

### GetBatchClosedOk

`func (o *MerchantBatchResponse) GetBatchClosedOk() (*time.Time, bool)`

GetBatchClosedOk returns a tuple with the BatchClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchClosed

`func (o *MerchantBatchResponse) SetBatchClosed(v time.Time)`

SetBatchClosed sets BatchClosed field to given value.

### HasBatchClosed

`func (o *MerchantBatchResponse) HasBatchClosed() bool`

HasBatchClosed returns a boolean if a field has been set.

### GetBatchNo

`func (o *MerchantBatchResponse) GetBatchNo() string`

GetBatchNo returns the BatchNo field if non-nil, zero value otherwise.

### GetBatchNoOk

`func (o *MerchantBatchResponse) GetBatchNoOk() (*string, bool)`

GetBatchNoOk returns a tuple with the BatchNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchNo

`func (o *MerchantBatchResponse) SetBatchNo(v string)`

SetBatchNo sets BatchNo field to given value.

### HasBatchNo

`func (o *MerchantBatchResponse) HasBatchNo() bool`

HasBatchNo returns a boolean if a field has been set.

### GetBatchStatus

`func (o *MerchantBatchResponse) GetBatchStatus() string`

GetBatchStatus returns the BatchStatus field if non-nil, zero value otherwise.

### GetBatchStatusOk

`func (o *MerchantBatchResponse) GetBatchStatusOk() (*string, bool)`

GetBatchStatusOk returns a tuple with the BatchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchStatus

`func (o *MerchantBatchResponse) SetBatchStatus(v string)`

SetBatchStatus sets BatchStatus field to given value.

### HasBatchStatus

`func (o *MerchantBatchResponse) HasBatchStatus() bool`

HasBatchStatus returns a boolean if a field has been set.

### GetBatchStatusCode

`func (o *MerchantBatchResponse) GetBatchStatusCode() string`

GetBatchStatusCode returns the BatchStatusCode field if non-nil, zero value otherwise.

### GetBatchStatusCodeOk

`func (o *MerchantBatchResponse) GetBatchStatusCodeOk() (*string, bool)`

GetBatchStatusCodeOk returns a tuple with the BatchStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchStatusCode

`func (o *MerchantBatchResponse) SetBatchStatusCode(v string)`

SetBatchStatusCode sets BatchStatusCode field to given value.

### HasBatchStatusCode

`func (o *MerchantBatchResponse) HasBatchStatusCode() bool`

HasBatchStatusCode returns a boolean if a field has been set.

### GetCurrency

`func (o *MerchantBatchResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MerchantBatchResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MerchantBatchResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MerchantBatchResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetMerchantid

`func (o *MerchantBatchResponse) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *MerchantBatchResponse) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *MerchantBatchResponse) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.

### HasMerchantid

`func (o *MerchantBatchResponse) HasMerchantid() bool`

HasMerchantid returns a boolean if a field has been set.

### GetNetSummary

`func (o *MerchantBatchResponse) GetNetSummary() NetSummaryResponse`

GetNetSummary returns the NetSummary field if non-nil, zero value otherwise.

### GetNetSummaryOk

`func (o *MerchantBatchResponse) GetNetSummaryOk() (*NetSummaryResponse, bool)`

GetNetSummaryOk returns a tuple with the NetSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetSummary

`func (o *MerchantBatchResponse) SetNetSummary(v NetSummaryResponse)`

SetNetSummary sets NetSummary field to given value.

### HasNetSummary

`func (o *MerchantBatchResponse) HasNetSummary() bool`

HasNetSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


