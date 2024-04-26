# PaResAuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Md** | **string** | The Merchant Data (MD) which is a unique ID to reference the authentication session.  This value will be created by CityPay when required. When responding from the ACS, this value will be returned by the ACS.  | 
**Pares** | **string** | The Payer Authentication Response packet which is returned by the ACS containing the  response of the authentication session including verification values. The response  is a base64 encoded packet and should be forwarded to CityPay untouched.  | 

## Methods

### NewPaResAuthRequest

`func NewPaResAuthRequest(md string, pares string, ) *PaResAuthRequest`

NewPaResAuthRequest instantiates a new PaResAuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaResAuthRequestWithDefaults

`func NewPaResAuthRequestWithDefaults() *PaResAuthRequest`

NewPaResAuthRequestWithDefaults instantiates a new PaResAuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMd

`func (o *PaResAuthRequest) GetMd() string`

GetMd returns the Md field if non-nil, zero value otherwise.

### GetMdOk

`func (o *PaResAuthRequest) GetMdOk() (*string, bool)`

GetMdOk returns a tuple with the Md field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd

`func (o *PaResAuthRequest) SetMd(v string)`

SetMd sets Md field to given value.


### GetPares

`func (o *PaResAuthRequest) GetPares() string`

GetPares returns the Pares field if non-nil, zero value otherwise.

### GetParesOk

`func (o *PaResAuthRequest) GetParesOk() (*string, bool)`

GetParesOk returns a tuple with the Pares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPares

`func (o *PaResAuthRequest) SetPares(v string)`

SetPares sets Pares field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


