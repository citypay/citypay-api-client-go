# ListMerchantsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | Pointer to **string** | The client name that was requested. | [optional] 
**Clientid** | Pointer to **string** | The client id requested. | [optional] 
**Merchants** | Pointer to [**[]Merchant**](Merchant.md) |  | [optional] 

## Methods

### NewListMerchantsResponse

`func NewListMerchantsResponse() *ListMerchantsResponse`

NewListMerchantsResponse instantiates a new ListMerchantsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMerchantsResponseWithDefaults

`func NewListMerchantsResponseWithDefaults() *ListMerchantsResponse`

NewListMerchantsResponseWithDefaults instantiates a new ListMerchantsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *ListMerchantsResponse) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *ListMerchantsResponse) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *ListMerchantsResponse) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *ListMerchantsResponse) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetClientid

`func (o *ListMerchantsResponse) GetClientid() string`

GetClientid returns the Clientid field if non-nil, zero value otherwise.

### GetClientidOk

`func (o *ListMerchantsResponse) GetClientidOk() (*string, bool)`

GetClientidOk returns a tuple with the Clientid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientid

`func (o *ListMerchantsResponse) SetClientid(v string)`

SetClientid sets Clientid field to given value.

### HasClientid

`func (o *ListMerchantsResponse) HasClientid() bool`

HasClientid returns a boolean if a field has been set.

### GetMerchants

`func (o *ListMerchantsResponse) GetMerchants() []Merchant`

GetMerchants returns the Merchants field if non-nil, zero value otherwise.

### GetMerchantsOk

`func (o *ListMerchantsResponse) GetMerchantsOk() (*[]Merchant, bool)`

GetMerchantsOk returns a tuple with the Merchants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchants

`func (o *ListMerchantsResponse) SetMerchants(v []Merchant)`

SetMerchants sets Merchants field to given value.

### HasMerchants

`func (o *ListMerchantsResponse) HasMerchants() bool`

HasMerchants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


