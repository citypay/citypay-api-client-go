# PaylinkTokenStatusChangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | **time.Time** | identifies the date and time to lookup changes after. | 
**MaxResults** | Pointer to **int32** | The maximum number of results to return in a single response. This value is used to limit the size of data returned by the API, enhancing performance and manageability. Values should be between 5 and 250. | [optional] 
**Merchantid** | **int32** | the merchant id to review tokens for. | 
**NextToken** | Pointer to **string** | A token that identifies the starting point of the page of results to be returned. An empty value indicates the start of the dataset. When supplied, it is validated and used to fetch the subsequent page of results. This token is typically obtained from the response of a previous pagination request. | [optional] 
**OrderBy** | Pointer to **string** | Specifies the field by which results are ordered. Available fields are [p.id]. By default, fields are ordered by OrderByExpression(p.id,ASC). To order in descending order, prefix with &#39;-&#39; or suffix with &#39; DESC&#39;. | [optional] 

## Methods

### NewPaylinkTokenStatusChangeRequest

`func NewPaylinkTokenStatusChangeRequest(after time.Time, merchantid int32, ) *PaylinkTokenStatusChangeRequest`

NewPaylinkTokenStatusChangeRequest instantiates a new PaylinkTokenStatusChangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkTokenStatusChangeRequestWithDefaults

`func NewPaylinkTokenStatusChangeRequestWithDefaults() *PaylinkTokenStatusChangeRequest`

NewPaylinkTokenStatusChangeRequestWithDefaults instantiates a new PaylinkTokenStatusChangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *PaylinkTokenStatusChangeRequest) GetAfter() time.Time`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *PaylinkTokenStatusChangeRequest) GetAfterOk() (*time.Time, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *PaylinkTokenStatusChangeRequest) SetAfter(v time.Time)`

SetAfter sets After field to given value.


### GetMaxResults

`func (o *PaylinkTokenStatusChangeRequest) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *PaylinkTokenStatusChangeRequest) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *PaylinkTokenStatusChangeRequest) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *PaylinkTokenStatusChangeRequest) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetMerchantid

`func (o *PaylinkTokenStatusChangeRequest) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *PaylinkTokenStatusChangeRequest) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *PaylinkTokenStatusChangeRequest) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.


### GetNextToken

`func (o *PaylinkTokenStatusChangeRequest) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *PaylinkTokenStatusChangeRequest) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *PaylinkTokenStatusChangeRequest) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *PaylinkTokenStatusChangeRequest) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetOrderBy

`func (o *PaylinkTokenStatusChangeRequest) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *PaylinkTokenStatusChangeRequest) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *PaylinkTokenStatusChangeRequest) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *PaylinkTokenStatusChangeRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


