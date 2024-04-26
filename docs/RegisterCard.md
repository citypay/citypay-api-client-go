# RegisterCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cardnumber** | **string** | The primary number of the card. | 
**Default** | Pointer to **bool** | Determines whether the card should be the new default card. | [optional] 
**Expmonth** | **int32** | The expiry month of the card. | 
**Expyear** | **int32** | The expiry year of the card. | 
**NameOnCard** | Pointer to **string** | The card holder name as it appears on the card. The value is required if the account is to be used for 3dsv2 processing, otherwise it is optional. | [optional] 

## Methods

### NewRegisterCard

`func NewRegisterCard(cardnumber string, expmonth int32, expyear int32, ) *RegisterCard`

NewRegisterCard instantiates a new RegisterCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterCardWithDefaults

`func NewRegisterCardWithDefaults() *RegisterCard`

NewRegisterCardWithDefaults instantiates a new RegisterCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardnumber

`func (o *RegisterCard) GetCardnumber() string`

GetCardnumber returns the Cardnumber field if non-nil, zero value otherwise.

### GetCardnumberOk

`func (o *RegisterCard) GetCardnumberOk() (*string, bool)`

GetCardnumberOk returns a tuple with the Cardnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardnumber

`func (o *RegisterCard) SetCardnumber(v string)`

SetCardnumber sets Cardnumber field to given value.


### GetDefault

`func (o *RegisterCard) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *RegisterCard) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *RegisterCard) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *RegisterCard) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetExpmonth

`func (o *RegisterCard) GetExpmonth() int32`

GetExpmonth returns the Expmonth field if non-nil, zero value otherwise.

### GetExpmonthOk

`func (o *RegisterCard) GetExpmonthOk() (*int32, bool)`

GetExpmonthOk returns a tuple with the Expmonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpmonth

`func (o *RegisterCard) SetExpmonth(v int32)`

SetExpmonth sets Expmonth field to given value.


### GetExpyear

`func (o *RegisterCard) GetExpyear() int32`

GetExpyear returns the Expyear field if non-nil, zero value otherwise.

### GetExpyearOk

`func (o *RegisterCard) GetExpyearOk() (*int32, bool)`

GetExpyearOk returns a tuple with the Expyear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpyear

`func (o *RegisterCard) SetExpyear(v int32)`

SetExpyear sets Expyear field to given value.


### GetNameOnCard

`func (o *RegisterCard) GetNameOnCard() string`

GetNameOnCard returns the NameOnCard field if non-nil, zero value otherwise.

### GetNameOnCardOk

`func (o *RegisterCard) GetNameOnCardOk() (*string, bool)`

GetNameOnCardOk returns a tuple with the NameOnCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCard

`func (o *RegisterCard) SetNameOnCard(v string)`

SetNameOnCard sets NameOnCard field to given value.

### HasNameOnCard

`func (o *RegisterCard) HasNameOnCard() bool`

HasNameOnCard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


