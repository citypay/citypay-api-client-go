# RemittanceReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateFrom** | Pointer to **string** | Start date (YYYY-MM-DD) for batch Retrieval range, inclusive. Maximum value is 3 years ago. | [optional] 
**DateUntil** | Pointer to **string** | End date (YYYY-MM-DD) for batch Retrieval range, inclusive. | [optional] 
**MaxResults** | Pointer to **int32** | The maximum number of results to return in a single response. This value is used to limit the size of data returned by the API, enhancing performance and manageability. Values should be between 5 and 250. | [optional] 
**MerchantId** | Pointer to **[]int32** |  | [optional] 
**NextToken** | Pointer to **string** | A token that identifies the starting point of the page of results to be returned. An empty value indicates the start of the dataset. When supplied, it is validated and used to fetch the subsequent page of results. This token is typically obtained from the response of a previous pagination request. | [optional] 
**OrderBy** | Pointer to **string** | Specifies the field by which results are ordered. Available fields are [trans_no,date_when,amount]. By default, fields are ordered by OrderByExpression(trans_no,ASC). To order in descending order, prefix with &#39;-&#39; or suffix with &#39; DESC&#39;. | [optional] 

## Methods

### NewRemittanceReportRequest

`func NewRemittanceReportRequest() *RemittanceReportRequest`

NewRemittanceReportRequest instantiates a new RemittanceReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemittanceReportRequestWithDefaults

`func NewRemittanceReportRequestWithDefaults() *RemittanceReportRequest`

NewRemittanceReportRequestWithDefaults instantiates a new RemittanceReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateFrom

`func (o *RemittanceReportRequest) GetDateFrom() string`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *RemittanceReportRequest) GetDateFromOk() (*string, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *RemittanceReportRequest) SetDateFrom(v string)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *RemittanceReportRequest) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### GetDateUntil

`func (o *RemittanceReportRequest) GetDateUntil() string`

GetDateUntil returns the DateUntil field if non-nil, zero value otherwise.

### GetDateUntilOk

`func (o *RemittanceReportRequest) GetDateUntilOk() (*string, bool)`

GetDateUntilOk returns a tuple with the DateUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUntil

`func (o *RemittanceReportRequest) SetDateUntil(v string)`

SetDateUntil sets DateUntil field to given value.

### HasDateUntil

`func (o *RemittanceReportRequest) HasDateUntil() bool`

HasDateUntil returns a boolean if a field has been set.

### GetMaxResults

`func (o *RemittanceReportRequest) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *RemittanceReportRequest) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *RemittanceReportRequest) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *RemittanceReportRequest) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetMerchantId

`func (o *RemittanceReportRequest) GetMerchantId() []int32`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *RemittanceReportRequest) GetMerchantIdOk() (*[]int32, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *RemittanceReportRequest) SetMerchantId(v []int32)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *RemittanceReportRequest) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetNextToken

`func (o *RemittanceReportRequest) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *RemittanceReportRequest) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *RemittanceReportRequest) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *RemittanceReportRequest) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetOrderBy

`func (o *RemittanceReportRequest) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *RemittanceReportRequest) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *RemittanceReportRequest) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *RemittanceReportRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


