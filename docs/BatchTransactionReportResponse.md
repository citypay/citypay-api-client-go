# BatchTransactionReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The count of items returned in this page. | [optional] 
**Data** | [**[]AuthReference**](AuthReference.md) |  | 
**MaxResults** | Pointer to **int32** | The max results requested in this page. | [optional] 
**NextToken** | Pointer to **string** | A token that identifies the starting point of the page of results to be returned. An empty value indicates the start of the dataset. When supplied, it is validated and used to fetch the subsequent page of results. This token is typically obtained from the response of a previous pagination request. | [optional] 

## Methods

### NewBatchTransactionReportResponse

`func NewBatchTransactionReportResponse(data []AuthReference, ) *BatchTransactionReportResponse`

NewBatchTransactionReportResponse instantiates a new BatchTransactionReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchTransactionReportResponseWithDefaults

`func NewBatchTransactionReportResponseWithDefaults() *BatchTransactionReportResponse`

NewBatchTransactionReportResponseWithDefaults instantiates a new BatchTransactionReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *BatchTransactionReportResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BatchTransactionReportResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BatchTransactionReportResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *BatchTransactionReportResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetData

`func (o *BatchTransactionReportResponse) GetData() []AuthReference`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BatchTransactionReportResponse) GetDataOk() (*[]AuthReference, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BatchTransactionReportResponse) SetData(v []AuthReference)`

SetData sets Data field to given value.


### GetMaxResults

`func (o *BatchTransactionReportResponse) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *BatchTransactionReportResponse) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *BatchTransactionReportResponse) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *BatchTransactionReportResponse) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetNextToken

`func (o *BatchTransactionReportResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *BatchTransactionReportResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *BatchTransactionReportResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *BatchTransactionReportResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


