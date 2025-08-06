# DomainKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **[]string** |  | 
**Live** | Pointer to **bool** | Specifies if the key is to be used for production. Defaults to false.  | [optional] 
**Merchantid** | **int32** | The merchant id the domain key is to be used for.  | 
**Nonce** | Pointer to **string** | Specifies a random value for integrity. The value is used to generate the domain key to provide further integrity to the key.  | [optional] 

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


### GetNonce

`func (o *DomainKeyRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *DomainKeyRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *DomainKeyRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *DomainKeyRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


