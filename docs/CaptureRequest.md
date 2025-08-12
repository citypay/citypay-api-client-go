# CaptureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AirlineData** | Pointer to [**AirlineAdvice**](AirlineAdvice.md) |  | [optional] 
**Amount** | Pointer to **int32** | The completion amount provided in the lowest unit of currency for the specific currency of the merchant, with a variable length to a maximum of 12 digits. No decimal points to be included. For example with GBP 75.45 use the value 7545. Please check that you do not supply divisional characters such as 1,024 in the request which may be caused by some number formatters.  If no amount is supplied, the original processing amount is used.  | [optional] 
**EventManagement** | Pointer to [**EventDataModel**](EventDataModel.md) |  | [optional] 
**Identifier** | Pointer to **string** | The identifier of the transaction to capture. If an empty value is supplied then a &#x60;trans_no&#x60; value must be supplied. | [optional] 
**Merchantid** | **int32** | Identifies the merchant account to perform the capture for. | 
**Transno** | Pointer to **int32** | The transaction number of the transaction to look up and capture. If an empty value is supplied then an identifier value must be supplied. | [optional] 

## Methods

### NewCaptureRequest

`func NewCaptureRequest(merchantid int32, ) *CaptureRequest`

NewCaptureRequest instantiates a new CaptureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptureRequestWithDefaults

`func NewCaptureRequestWithDefaults() *CaptureRequest`

NewCaptureRequestWithDefaults instantiates a new CaptureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAirlineData

`func (o *CaptureRequest) GetAirlineData() AirlineAdvice`

GetAirlineData returns the AirlineData field if non-nil, zero value otherwise.

### GetAirlineDataOk

`func (o *CaptureRequest) GetAirlineDataOk() (*AirlineAdvice, bool)`

GetAirlineDataOk returns a tuple with the AirlineData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirlineData

`func (o *CaptureRequest) SetAirlineData(v AirlineAdvice)`

SetAirlineData sets AirlineData field to given value.

### HasAirlineData

`func (o *CaptureRequest) HasAirlineData() bool`

HasAirlineData returns a boolean if a field has been set.

### GetAmount

`func (o *CaptureRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CaptureRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CaptureRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CaptureRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetEventManagement

`func (o *CaptureRequest) GetEventManagement() EventDataModel`

GetEventManagement returns the EventManagement field if non-nil, zero value otherwise.

### GetEventManagementOk

`func (o *CaptureRequest) GetEventManagementOk() (*EventDataModel, bool)`

GetEventManagementOk returns a tuple with the EventManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventManagement

`func (o *CaptureRequest) SetEventManagement(v EventDataModel)`

SetEventManagement sets EventManagement field to given value.

### HasEventManagement

`func (o *CaptureRequest) HasEventManagement() bool`

HasEventManagement returns a boolean if a field has been set.

### GetIdentifier

`func (o *CaptureRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CaptureRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CaptureRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *CaptureRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetMerchantid

`func (o *CaptureRequest) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *CaptureRequest) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *CaptureRequest) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.


### GetTransno

`func (o *CaptureRequest) GetTransno() int32`

GetTransno returns the Transno field if non-nil, zero value otherwise.

### GetTransnoOk

`func (o *CaptureRequest) GetTransnoOk() (*int32, bool)`

GetTransnoOk returns a tuple with the Transno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransno

`func (o *CaptureRequest) SetTransno(v int32)`

SetTransno sets Transno field to given value.

### HasTransno

`func (o *CaptureRequest) HasTransno() bool`

HasTransno returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


