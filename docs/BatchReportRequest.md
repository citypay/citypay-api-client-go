# BatchReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | **int32** | The batch id specified in the batch processing request. | 
**ClientAccountId** | Pointer to **string** | The batch account id that the batch was processed for. Defaults to your client id if not provided. | [optional] 

## Methods

### NewBatchReportRequest

`func NewBatchReportRequest(batchId int32, ) *BatchReportRequest`

NewBatchReportRequest instantiates a new BatchReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchReportRequestWithDefaults

`func NewBatchReportRequestWithDefaults() *BatchReportRequest`

NewBatchReportRequestWithDefaults instantiates a new BatchReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *BatchReportRequest) GetBatchId() int32`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *BatchReportRequest) GetBatchIdOk() (*int32, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *BatchReportRequest) SetBatchId(v int32)`

SetBatchId sets BatchId field to given value.


### GetClientAccountId

`func (o *BatchReportRequest) GetClientAccountId() string`

GetClientAccountId returns the ClientAccountId field if non-nil, zero value otherwise.

### GetClientAccountIdOk

`func (o *BatchReportRequest) GetClientAccountIdOk() (*string, bool)`

GetClientAccountIdOk returns a tuple with the ClientAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAccountId

`func (o *BatchReportRequest) SetClientAccountId(v string)`

SetClientAccountId sets ClientAccountId field to given value.

### HasClientAccountId

`func (o *BatchReportRequest) HasClientAccountId() bool`

HasClientAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


