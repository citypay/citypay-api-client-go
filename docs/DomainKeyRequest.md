# DomainKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **[]string** |  | 
**Live** | Pointer to **bool** | Specifies if the key is to be used for production. Defaults to false.  | [optional] 
**Merchantid** | **int32** | The merchant id the domain key is to be used for.  | 

## Methods

### NewDomainKeyRequest

`func NewDomainKeyRequest(domain []string, merchantid int32, ) *DomainKeyRequest`

NewDomainKeyRequest instantiates a new DomainKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainKeyRequestWithDefaults

`func NewDomainKeyRequestWithDefaults() *DomainKeyRequest`

NewDomainKeyRequestWithDefaults instantiates a new DomainKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DomainKeyRequest) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainKeyRequest) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainKeyRequest) SetDomain(v []string)`

SetDomain sets Domain field to given value.


### GetLive

`func (o *DomainKeyRequest) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *DomainKeyRequest) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *DomainKeyRequest) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *DomainKeyRequest) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetMerchantid

`func (o *DomainKeyRequest) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *DomainKeyRequest) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *DomainKeyRequest) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


