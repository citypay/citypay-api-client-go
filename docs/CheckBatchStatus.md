# CheckBatchStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | **[]int32** |  | 
**ClientAccountId** | Pointer to **string** | The batch account id to obtain the batch for. Defaults to your client id if not provided. | [optional] 

## Methods

### NewCheckBatchStatus

`func NewCheckBatchStatus(batchId []int32, ) *CheckBatchStatus`

NewCheckBatchStatus instantiates a new CheckBatchStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckBatchStatusWithDefaults

`func NewCheckBatchStatusWithDefaults() *CheckBatchStatus`

NewCheckBatchStatusWithDefaults instantiates a new CheckBatchStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *CheckBatchStatus) GetBatchId() []int32`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *CheckBatchStatus) GetBatchIdOk() (*[]int32, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *CheckBatchStatus) SetBatchId(v []int32)`

SetBatchId sets BatchId field to given value.


### GetClientAccountId

`func (o *CheckBatchStatus) GetClientAccountId() string`

GetClientAccountId returns the ClientAccountId field if non-nil, zero value otherwise.

### GetClientAccountIdOk

`func (o *CheckBatchStatus) GetClientAccountIdOk() (*string, bool)`

GetClientAccountIdOk returns a tuple with the ClientAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAccountId

`func (o *CheckBatchStatus) SetClientAccountId(v string)`

SetClientAccountId sets ClientAccountId field to given value.

### HasClientAccountId

`func (o *CheckBatchStatus) HasClientAccountId() bool`

HasClientAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


