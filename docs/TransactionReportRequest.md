# TransactionReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to **[]string** |  | [optional] 
**From** | **time.Time** | The date and time of transactions from. | 
**IncludeAuthorised** | Pointer to **bool** | Include transactions fully authorised in the results. | [optional] 
**IncludeCancelled** | Pointer to **bool** | Include transactions that were cancelled in the results. | [optional] 
**IncludeDeclined** | Pointer to **bool** | Include transactions that were declined or not honoured in the results. | [optional] 
**IncludeRejected** | Pointer to **bool** | Include transactions that were rejected due to validation issues. | [optional] 
**IncludeUnfulfilled** | Pointer to **bool** | Includes transactions that were initiated but not completed—e.g. those pending authentication or challenge responses that were never fulfilled. | [optional] 
**MaxResults** | Pointer to **int32** | The maximum number of results to return in a single response. This value is used to limit the size of data returned by the API, enhancing performance and manageability. Values should be between 5 and 250. | [optional] 
**Merchantid** | **int32** | The merchant id of the transactions to review. | 
**Mode** | Pointer to **string** | Defines a preset profile for the level of detail in the returned fields. This simplifies response formatting for common use cases. Available values:  - &#x60;basic&#x60; (default): Returns a minimal, high-level view with key fields for reporting or dashboards.  - &#x60;extended&#x60;: Adds fields useful for customer support, settlement analysis, or more in-depth tracking, while still omitting sensitive personal or low-level fields.  - &#x60;full&#x60;: Returns all available transaction fields, including internal flags, personal data (where applicable), and detailed metadata. Use with care.  | [optional] 
**NextToken** | Pointer to **string** | A token that identifies the starting point of the page of results to be returned. An empty value indicates the start of the dataset. When supplied, it is validated and used to fetch the subsequent page of results. This token is typically obtained from the response of a previous pagination request. | [optional] 
**OrderBy** | Pointer to **string** | Specifies the field by which results are ordered. Available fields are [trans_no,date_when,amount]. By default, fields are ordered by OrderByExpression(trans_no,ASC). To order in descending order, prefix with &#39;-&#39; or suffix with &#39; DESC&#39;. | [optional] 
**PiiMasked** | Pointer to **bool** | Defines whether personal identifiable information is masked which it is by default. | [optional] 
**TypeRefund** | Pointer to **bool** | Include refunds in the results. | [optional] 
**TypeSale** | Pointer to **bool** | Include sales in the results. | [optional] 
**TypeVerify** | Pointer to **bool** | Include verifications in the results. | [optional] 
**Until** | **time.Time** | The date and time of transactions until. | 

## Methods

### NewTransactionReportRequest

`func NewTransactionReportRequest(from time.Time, merchantid int32, until time.Time, ) *TransactionReportRequest`

NewTransactionReportRequest instantiates a new TransactionReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionReportRequestWithDefaults

`func NewTransactionReportRequestWithDefaults() *TransactionReportRequest`

NewTransactionReportRequestWithDefaults instantiates a new TransactionReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *TransactionReportRequest) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *TransactionReportRequest) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *TransactionReportRequest) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *TransactionReportRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFrom

`func (o *TransactionReportRequest) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TransactionReportRequest) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TransactionReportRequest) SetFrom(v time.Time)`

SetFrom sets From field to given value.


### GetIncludeAuthorised

`func (o *TransactionReportRequest) GetIncludeAuthorised() bool`

GetIncludeAuthorised returns the IncludeAuthorised field if non-nil, zero value otherwise.

### GetIncludeAuthorisedOk

`func (o *TransactionReportRequest) GetIncludeAuthorisedOk() (*bool, bool)`

GetIncludeAuthorisedOk returns a tuple with the IncludeAuthorised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAuthorised

`func (o *TransactionReportRequest) SetIncludeAuthorised(v bool)`

SetIncludeAuthorised sets IncludeAuthorised field to given value.

### HasIncludeAuthorised

`func (o *TransactionReportRequest) HasIncludeAuthorised() bool`

HasIncludeAuthorised returns a boolean if a field has been set.

### GetIncludeCancelled

`func (o *TransactionReportRequest) GetIncludeCancelled() bool`

GetIncludeCancelled returns the IncludeCancelled field if non-nil, zero value otherwise.

### GetIncludeCancelledOk

`func (o *TransactionReportRequest) GetIncludeCancelledOk() (*bool, bool)`

GetIncludeCancelledOk returns a tuple with the IncludeCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCancelled

`func (o *TransactionReportRequest) SetIncludeCancelled(v bool)`

SetIncludeCancelled sets IncludeCancelled field to given value.

### HasIncludeCancelled

`func (o *TransactionReportRequest) HasIncludeCancelled() bool`

HasIncludeCancelled returns a boolean if a field has been set.

### GetIncludeDeclined

`func (o *TransactionReportRequest) GetIncludeDeclined() bool`

GetIncludeDeclined returns the IncludeDeclined field if non-nil, zero value otherwise.

### GetIncludeDeclinedOk

`func (o *TransactionReportRequest) GetIncludeDeclinedOk() (*bool, bool)`

GetIncludeDeclinedOk returns a tuple with the IncludeDeclined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDeclined

`func (o *TransactionReportRequest) SetIncludeDeclined(v bool)`

SetIncludeDeclined sets IncludeDeclined field to given value.

### HasIncludeDeclined

`func (o *TransactionReportRequest) HasIncludeDeclined() bool`

HasIncludeDeclined returns a boolean if a field has been set.

### GetIncludeRejected

`func (o *TransactionReportRequest) GetIncludeRejected() bool`

GetIncludeRejected returns the IncludeRejected field if non-nil, zero value otherwise.

### GetIncludeRejectedOk

`func (o *TransactionReportRequest) GetIncludeRejectedOk() (*bool, bool)`

GetIncludeRejectedOk returns a tuple with the IncludeRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRejected

`func (o *TransactionReportRequest) SetIncludeRejected(v bool)`

SetIncludeRejected sets IncludeRejected field to given value.

### HasIncludeRejected

`func (o *TransactionReportRequest) HasIncludeRejected() bool`

HasIncludeRejected returns a boolean if a field has been set.

### GetIncludeUnfulfilled

`func (o *TransactionReportRequest) GetIncludeUnfulfilled() bool`

GetIncludeUnfulfilled returns the IncludeUnfulfilled field if non-nil, zero value otherwise.

### GetIncludeUnfulfilledOk

`func (o *TransactionReportRequest) GetIncludeUnfulfilledOk() (*bool, bool)`

GetIncludeUnfulfilledOk returns a tuple with the IncludeUnfulfilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUnfulfilled

`func (o *TransactionReportRequest) SetIncludeUnfulfilled(v bool)`

SetIncludeUnfulfilled sets IncludeUnfulfilled field to given value.

### HasIncludeUnfulfilled

`func (o *TransactionReportRequest) HasIncludeUnfulfilled() bool`

HasIncludeUnfulfilled returns a boolean if a field has been set.

### GetMaxResults

`func (o *TransactionReportRequest) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *TransactionReportRequest) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *TransactionReportRequest) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *TransactionReportRequest) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetMerchantid

`func (o *TransactionReportRequest) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *TransactionReportRequest) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *TransactionReportRequest) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.


### GetMode

`func (o *TransactionReportRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TransactionReportRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TransactionReportRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *TransactionReportRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNextToken

`func (o *TransactionReportRequest) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *TransactionReportRequest) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *TransactionReportRequest) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *TransactionReportRequest) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetOrderBy

`func (o *TransactionReportRequest) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *TransactionReportRequest) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *TransactionReportRequest) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *TransactionReportRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetPiiMasked

`func (o *TransactionReportRequest) GetPiiMasked() bool`

GetPiiMasked returns the PiiMasked field if non-nil, zero value otherwise.

### GetPiiMaskedOk

`func (o *TransactionReportRequest) GetPiiMaskedOk() (*bool, bool)`

GetPiiMaskedOk returns a tuple with the PiiMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPiiMasked

`func (o *TransactionReportRequest) SetPiiMasked(v bool)`

SetPiiMasked sets PiiMasked field to given value.

### HasPiiMasked

`func (o *TransactionReportRequest) HasPiiMasked() bool`

HasPiiMasked returns a boolean if a field has been set.

### GetTypeRefund

`func (o *TransactionReportRequest) GetTypeRefund() bool`

GetTypeRefund returns the TypeRefund field if non-nil, zero value otherwise.

### GetTypeRefundOk

`func (o *TransactionReportRequest) GetTypeRefundOk() (*bool, bool)`

GetTypeRefundOk returns a tuple with the TypeRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeRefund

`func (o *TransactionReportRequest) SetTypeRefund(v bool)`

SetTypeRefund sets TypeRefund field to given value.

### HasTypeRefund

`func (o *TransactionReportRequest) HasTypeRefund() bool`

HasTypeRefund returns a boolean if a field has been set.

### GetTypeSale

`func (o *TransactionReportRequest) GetTypeSale() bool`

GetTypeSale returns the TypeSale field if non-nil, zero value otherwise.

### GetTypeSaleOk

`func (o *TransactionReportRequest) GetTypeSaleOk() (*bool, bool)`

GetTypeSaleOk returns a tuple with the TypeSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeSale

`func (o *TransactionReportRequest) SetTypeSale(v bool)`

SetTypeSale sets TypeSale field to given value.

### HasTypeSale

`func (o *TransactionReportRequest) HasTypeSale() bool`

HasTypeSale returns a boolean if a field has been set.

### GetTypeVerify

`func (o *TransactionReportRequest) GetTypeVerify() bool`

GetTypeVerify returns the TypeVerify field if non-nil, zero value otherwise.

### GetTypeVerifyOk

`func (o *TransactionReportRequest) GetTypeVerifyOk() (*bool, bool)`

GetTypeVerifyOk returns a tuple with the TypeVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeVerify

`func (o *TransactionReportRequest) SetTypeVerify(v bool)`

SetTypeVerify sets TypeVerify field to given value.

### HasTypeVerify

`func (o *TransactionReportRequest) HasTypeVerify() bool`

HasTypeVerify returns a boolean if a field has been set.

### GetUntil

`func (o *TransactionReportRequest) GetUntil() time.Time`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *TransactionReportRequest) GetUntilOk() (*time.Time, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *TransactionReportRequest) SetUntil(v time.Time)`

SetUntil sets Until field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


