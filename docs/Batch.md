# Batch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchDate** | **string** | The date that the file was created in ISO-8601 format. | 
**BatchId** | Pointer to **int32** | The batch id requested. | [optional] 
**BatchStatus** | **string** | The status of the batch. Possible values are:    - CANCELLED. The file has been cancelled by an administrator or server process.  - COMPLETE. The file has passed through the processing cycle and is determined as being complete further information should be obtained on the results of the processing - ERROR_IN_PROCESSING. Errors have occurred in the processing that has deemed that processing can not continue. - INITIALISED. The file has been initialised and no action has yet been performed - LOCKED. The file has been locked for processing - QUEUED. The file has been queued for processing yet no processing has yet been performed - UNKNOWN. The file is of an unknown status, that is the file can either not be determined by the information requested of the file has not yet been received.  | 

## Methods

### NewBatch

`func NewBatch(batchDate string, batchStatus string, ) *Batch`

NewBatch instantiates a new Batch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchWithDefaults

`func NewBatchWithDefaults() *Batch`

NewBatchWithDefaults instantiates a new Batch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchDate

`func (o *Batch) GetBatchDate() string`

GetBatchDate returns the BatchDate field if non-nil, zero value otherwise.

### GetBatchDateOk

`func (o *Batch) GetBatchDateOk() (*string, bool)`

GetBatchDateOk returns a tuple with the BatchDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchDate

`func (o *Batch) SetBatchDate(v string)`

SetBatchDate sets BatchDate field to given value.


### GetBatchId

`func (o *Batch) GetBatchId() int32`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *Batch) GetBatchIdOk() (*int32, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *Batch) SetBatchId(v int32)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *Batch) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetBatchStatus

`func (o *Batch) GetBatchStatus() string`

GetBatchStatus returns the BatchStatus field if non-nil, zero value otherwise.

### GetBatchStatusOk

`func (o *Batch) GetBatchStatusOk() (*string, bool)`

GetBatchStatusOk returns a tuple with the BatchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchStatus

`func (o *Batch) SetBatchStatus(v string)`

SetBatchStatus sets BatchStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


