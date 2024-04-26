# DomainKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to **time.Time** | The date the domain key was generated.  | [optional] 
**Domain** | **[]string** |  | 
**DomainKey** | Pointer to **string** | The domain key generated.  | [optional] 
**Live** | Pointer to **bool** | true if this key is a production key.  | [optional] 
**Merchantid** | **int32** | The merchant id the domain key is to be used for.  | 

## Methods

### NewDomainKeyResponse

`func NewDomainKeyResponse(domain []string, merchantid int32, ) *DomainKeyResponse`

NewDomainKeyResponse instantiates a new DomainKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainKeyResponseWithDefaults

`func NewDomainKeyResponseWithDefaults() *DomainKeyResponse`

NewDomainKeyResponseWithDefaults instantiates a new DomainKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateCreated

`func (o *DomainKeyResponse) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *DomainKeyResponse) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *DomainKeyResponse) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *DomainKeyResponse) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDomain

`func (o *DomainKeyResponse) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainKeyResponse) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainKeyResponse) SetDomain(v []string)`

SetDomain sets Domain field to given value.


### GetDomainKey

`func (o *DomainKeyResponse) GetDomainKey() string`

GetDomainKey returns the DomainKey field if non-nil, zero value otherwise.

### GetDomainKeyOk

`func (o *DomainKeyResponse) GetDomainKeyOk() (*string, bool)`

GetDomainKeyOk returns a tuple with the DomainKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainKey

`func (o *DomainKeyResponse) SetDomainKey(v string)`

SetDomainKey sets DomainKey field to given value.

### HasDomainKey

`func (o *DomainKeyResponse) HasDomainKey() bool`

HasDomainKey returns a boolean if a field has been set.

### GetLive

`func (o *DomainKeyResponse) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *DomainKeyResponse) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *DomainKeyResponse) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *DomainKeyResponse) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetMerchantid

`func (o *DomainKeyResponse) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *DomainKeyResponse) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *DomainKeyResponse) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


