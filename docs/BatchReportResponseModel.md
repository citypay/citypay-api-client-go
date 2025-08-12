# BatchReportResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The total amount that the batch contains. | 
**BatchDate** | **string** | The date and time of the batch in ISO-8601 format. | 
**BatchId** | **int32** | The batch id specified in the batch processing request. | 
**BatchStatus** | **string** | The status of the batch. Possible values are:   - CANCELLED. The file has been cancelled by an administrator or server process.  - COMPLETE. The file has passed through the processing cycle and is determined as being complete further information should be obtained on the results of the processing - ERROR_IN_PROCESSING. Errors have occurred in the processing that has deemed that processing can not continue. - INITIALISED. The file has been initialised and no action has yet been performed - LOCKED. The file has been locked for processing - QUEUED. The file has been queued for processing yet no processing has yet been performed - UNKNOWN. The file is of an unknown status, that is the file can either not be determined by the information requested of the file has not yet been received.  | 
**ClientAccountId** | **string** | The batch account id that the batch was processed with. | 
**Transactions** | [**[]BatchTransactionResultModel**](BatchTransactionResultModel.md) |  | 

## Methods

### NewBatchReportResponseModel

`func NewBatchReportResponseModel(amount int32, batchDate string, batchId int32, batchStatus string, clientAccountId string, transactions []BatchTransactionResultModel, ) *BatchReportResponseModel`

NewBatchReportResponseModel instantiates a new BatchReportResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchReportResponseModelWithDefaults

`func NewBatchReportResponseModelWithDefaults() *BatchReportResponseModel`

NewBatchReportResponseModelWithDefaults instantiates a new BatchReportResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BatchReportResponseModel) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BatchReportResponseModel) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BatchReportResponseModel) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetBatchDate

`func (o *BatchReportResponseModel) GetBatchDate() string`

GetBatchDate returns the BatchDate field if non-nil, zero value otherwise.

### GetBatchDateOk

`func (o *BatchReportResponseModel) GetBatchDateOk() (*string, bool)`

GetBatchDateOk returns a tuple with the BatchDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchDate

`func (o *BatchReportResponseModel) SetBatchDate(v string)`

SetBatchDate sets BatchDate field to given value.


### GetBatchId

`func (o *BatchReportResponseModel) GetBatchId() int32`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *BatchReportResponseModel) GetBatchIdOk() (*int32, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *BatchReportResponseModel) SetBatchId(v int32)`

SetBatchId sets BatchId field to given value.


### GetBatchStatus

`func (o *BatchReportResponseModel) GetBatchStatus() string`

GetBatchStatus returns the BatchStatus field if non-nil, zero value otherwise.

### GetBatchStatusOk

`func (o *BatchReportResponseModel) GetBatchStatusOk() (*string, bool)`

GetBatchStatusOk returns a tuple with the BatchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchStatus

`func (o *BatchReportResponseModel) SetBatchStatus(v string)`

SetBatchStatus sets BatchStatus field to given value.


### GetClientAccountId

`func (o *BatchReportResponseModel) GetClientAccountId() string`

GetClientAccountId returns the ClientAccountId field if non-nil, zero value otherwise.

### GetClientAccountIdOk

`func (o *BatchReportResponseModel) GetClientAccountIdOk() (*string, bool)`

GetClientAccountIdOk returns a tuple with the ClientAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAccountId

`func (o *BatchReportResponseModel) SetClientAccountId(v string)`

SetClientAccountId sets ClientAccountId field to given value.


### GetTransactions

`func (o *BatchReportResponseModel) GetTransactions() []BatchTransactionResultModel`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *BatchReportResponseModel) GetTransactionsOk() (*[]BatchTransactionResultModel, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *BatchReportResponseModel) SetTransactions(v []BatchTransactionResultModel)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


