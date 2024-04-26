# MerchantBatchReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batches** | [**[]MerchantBatchResponse**](MerchantBatchResponse.md) |  | 
**Count** | Pointer to **int32** | The count of items returned in this page. | [optional] 
**MaxResults** | Pointer to **int32** | The max results requested in this page. | [optional] 
**NextToken** | Pointer to **string** | A token that identifies the starting point of the page of results to be returned. An empty value indicates the start of the dataset. When supplied, it is validated and used to fetch the subsequent page of results. This token is typically obtained from the response of a previous pagination request. | [optional] 

## Methods

### NewMerchantBatchReportResponse

`func NewMerchantBatchReportResponse(batches []MerchantBatchResponse, ) *MerchantBatchReportResponse`

NewMerchantBatchReportResponse instantiates a new MerchantBatchReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantBatchReportResponseWithDefaults

`func NewMerchantBatchReportResponseWithDefaults() *MerchantBatchReportResponse`

NewMerchantBatchReportResponseWithDefaults instantiates a new MerchantBatchReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatches

`func (o *MerchantBatchReportResponse) GetBatches() []MerchantBatchResponse`

GetBatches returns the Batches field if non-nil, zero value otherwise.

### GetBatchesOk

`func (o *MerchantBatchReportResponse) GetBatchesOk() (*[]MerchantBatchResponse, bool)`

GetBatchesOk returns a tuple with the Batches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatches

`func (o *MerchantBatchReportResponse) SetBatches(v []MerchantBatchResponse)`

SetBatches sets Batches field to given value.


### GetCount

`func (o *MerchantBatchReportResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MerchantBatchReportResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MerchantBatchReportResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *MerchantBatchReportResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetMaxResults

`func (o *MerchantBatchReportResponse) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *MerchantBatchReportResponse) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *MerchantBatchReportResponse) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *MerchantBatchReportResponse) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetNextToken

`func (o *MerchantBatchReportResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *MerchantBatchReportResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *MerchantBatchReportResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *MerchantBatchReportResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


