# RequestChallenged

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcsUrl** | Pointer to **string** | The url of the Access Control Server (ACS) to forward the user to.  | [optional] 
**Creq** | Pointer to **string** | The challenge request data which is encoded for usage by the ACS. | [optional] 
**Merchantid** | Pointer to **int32** | The merchant id that processed this transaction. | [optional] 
**ThreedserverTransId** | Pointer to **string** | The 3DSv2 trans id reference for the challenge process. May be used to create the ThreeDSSessionData value to send to the ACS. | [optional] 
**Transno** | Pointer to **int32** | The transaction number for the challenge, ordered incrementally from 1 for every merchant_id.  | [optional] 

## Methods

### NewRequestChallenged

`func NewRequestChallenged() *RequestChallenged`

NewRequestChallenged instantiates a new RequestChallenged object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestChallengedWithDefaults

`func NewRequestChallengedWithDefaults() *RequestChallenged`

NewRequestChallengedWithDefaults instantiates a new RequestChallenged object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcsUrl

`func (o *RequestChallenged) GetAcsUrl() string`

GetAcsUrl returns the AcsUrl field if non-nil, zero value otherwise.

### GetAcsUrlOk

`func (o *RequestChallenged) GetAcsUrlOk() (*string, bool)`

GetAcsUrlOk returns a tuple with the AcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsUrl

`func (o *RequestChallenged) SetAcsUrl(v string)`

SetAcsUrl sets AcsUrl field to given value.

### HasAcsUrl

`func (o *RequestChallenged) HasAcsUrl() bool`

HasAcsUrl returns a boolean if a field has been set.

### GetCreq

`func (o *RequestChallenged) GetCreq() string`

GetCreq returns the Creq field if non-nil, zero value otherwise.

### GetCreqOk

`func (o *RequestChallenged) GetCreqOk() (*string, bool)`

GetCreqOk returns a tuple with the Creq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreq

`func (o *RequestChallenged) SetCreq(v string)`

SetCreq sets Creq field to given value.

### HasCreq

`func (o *RequestChallenged) HasCreq() bool`

HasCreq returns a boolean if a field has been set.

### GetMerchantid

`func (o *RequestChallenged) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *RequestChallenged) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *RequestChallenged) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.

### HasMerchantid

`func (o *RequestChallenged) HasMerchantid() bool`

HasMerchantid returns a boolean if a field has been set.

### GetThreedserverTransId

`func (o *RequestChallenged) GetThreedserverTransId() string`

GetThreedserverTransId returns the ThreedserverTransId field if non-nil, zero value otherwise.

### GetThreedserverTransIdOk

`func (o *RequestChallenged) GetThreedserverTransIdOk() (*string, bool)`

GetThreedserverTransIdOk returns a tuple with the ThreedserverTransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreedserverTransId

`func (o *RequestChallenged) SetThreedserverTransId(v string)`

SetThreedserverTransId sets ThreedserverTransId field to given value.

### HasThreedserverTransId

`func (o *RequestChallenged) HasThreedserverTransId() bool`

HasThreedserverTransId returns a boolean if a field has been set.

### GetTransno

`func (o *RequestChallenged) GetTransno() int32`

GetTransno returns the Transno field if non-nil, zero value otherwise.

### GetTransnoOk

`func (o *RequestChallenged) GetTransnoOk() (*int32, bool)`

GetTransnoOk returns a tuple with the Transno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransno

`func (o *RequestChallenged) SetTransno(v int32)`

SetTransno sets Transno field to given value.

### HasTransno

`func (o *RequestChallenged) HasTransno() bool`

HasTransno returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


