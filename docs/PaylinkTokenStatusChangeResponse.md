# PaylinkTokenStatusChangeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The count of items returned in this page. | [optional] 
**MaxResults** | Pointer to **int32** | The max results requested in this page. | [optional] 
**NextToken** | Pointer to **string** | A token that identifies the starting point of the page of results to be returned. An empty value indicates the start of the dataset. When supplied, it is validated and used to fetch the subsequent page of results. This token is typically obtained from the response of a previous pagination request. | [optional] 
**Tokens** | [**[]PaylinkTokenStatus**](PaylinkTokenStatus.md) |  | 

## Methods

### NewPaylinkTokenStatusChangeResponse

`func NewPaylinkTokenStatusChangeResponse(tokens []PaylinkTokenStatus, ) *PaylinkTokenStatusChangeResponse`

NewPaylinkTokenStatusChangeResponse instantiates a new PaylinkTokenStatusChangeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkTokenStatusChangeResponseWithDefaults

`func NewPaylinkTokenStatusChangeResponseWithDefaults() *PaylinkTokenStatusChangeResponse`

NewPaylinkTokenStatusChangeResponseWithDefaults instantiates a new PaylinkTokenStatusChangeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaylinkTokenStatusChangeResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaylinkTokenStatusChangeResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaylinkTokenStatusChangeResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaylinkTokenStatusChangeResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetMaxResults

`func (o *PaylinkTokenStatusChangeResponse) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *PaylinkTokenStatusChangeResponse) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *PaylinkTokenStatusChangeResponse) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *PaylinkTokenStatusChangeResponse) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetNextToken

`func (o *PaylinkTokenStatusChangeResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *PaylinkTokenStatusChangeResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *PaylinkTokenStatusChangeResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *PaylinkTokenStatusChangeResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetTokens

`func (o *PaylinkTokenStatusChangeResponse) GetTokens() []PaylinkTokenStatus`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *PaylinkTokenStatusChangeResponse) GetTokensOk() (*[]PaylinkTokenStatus, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *PaylinkTokenStatusChangeResponse) SetTokens(v []PaylinkTokenStatus)`

SetTokens sets Tokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


