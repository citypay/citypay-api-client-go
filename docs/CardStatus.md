# CardStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardStatus** | Pointer to **string** | The status of the card to set, valid values are ACTIVE or INACTIVE. | [optional] 
**Default** | Pointer to **bool** | Defines if the card is set as the default. | [optional] 

## Methods

### NewCardStatus

`func NewCardStatus() *CardStatus`

NewCardStatus instantiates a new CardStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardStatusWithDefaults

`func NewCardStatusWithDefaults() *CardStatus`

NewCardStatusWithDefaults instantiates a new CardStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardStatus

`func (o *CardStatus) GetCardStatus() string`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *CardStatus) GetCardStatusOk() (*string, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *CardStatus) SetCardStatus(v string)`

SetCardStatus sets CardStatus field to given value.

### HasCardStatus

`func (o *CardStatus) HasCardStatus() bool`

HasCardStatus returns a boolean if a field has been set.

### GetDefault

`func (o *CardStatus) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CardStatus) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CardStatus) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *CardStatus) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


