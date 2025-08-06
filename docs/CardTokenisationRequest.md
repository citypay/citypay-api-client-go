# CardTokenisationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cardnumber** | Pointer to **string** | The card number (PAN) with a variable length to a maximum of 21 digits in numerical form. Any non numeric characters will be stripped out of the card number, this includes whitespace or separators internal of the provided value.  The card number must be treated as sensitive data. We only provide an obfuscated value in logging and reporting.  The plaintext value is encrypted in our database using AES 256 GMC bit encryption for settlement or refund purposes.  When providing the card number to our gateway through the authorisation API you will be handling the card data on your application. This will require further PCI controls to be in place and this value must never be stored.  | [optional] 
**Csc** | Pointer to **string** | The Card Security Code (CSC) (also known as CV2/CVV2) is normally found on the back of the card (American Express has it on the front). The value helps to identify possession of the card as it is not available within the chip or magnetic swipe.  When forwarding the CSC, please ensure the value is a string as some values start with 0 and this will be stripped out by any integer parsing.  The CSC number aids fraud prevention in Mail Order and Internet payments.  Business rules are available on your account to identify whether to accept or decline transactions based on mismatched results of the CSC.  The Payment Card Industry (PCI) requires that at no stage of a transaction should the CSC be stored.  This applies to all entities handling card data.  It should also not be used in any hashing process.  CityPay do not store the value and have no method of retrieving the value once the transaction has been processed. For this reason, duplicate checking is unable to determine the CSC in its duplication check algorithm.  | [optional] 
**Expmonth** | Pointer to **int32** | The month of expiry of the card. The month value should be a numerical value between 1 and 12.  | [optional] 
**Expyear** | Pointer to **int32** | The year of expiry of the card.  | [optional] 
**NameOnCard** | Pointer to **string** | The card holder name as appears on the card such as MR N E BODY. Required for some acquirers.  | [optional] 
**Uuid** | Pointer to **string** | A uuid for the session. The value tracks through 3ds session and therefore should be a valid v4 uuid. | [optional] 

## Methods

### NewCardTokenisationRequest

`func NewCardTokenisationRequest() *CardTokenisationRequest`

NewCardTokenisationRequest instantiates a new CardTokenisationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardTokenisationRequestWithDefaults

`func NewCardTokenisationRequestWithDefaults() *CardTokenisationRequest`

NewCardTokenisationRequestWithDefaults instantiates a new CardTokenisationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardnumber

`func (o *CardTokenisationRequest) GetCardnumber() string`

GetCardnumber returns the Cardnumber field if non-nil, zero value otherwise.

### GetCardnumberOk

`func (o *CardTokenisationRequest) GetCardnumberOk() (*string, bool)`

GetCardnumberOk returns a tuple with the Cardnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardnumber

`func (o *CardTokenisationRequest) SetCardnumber(v string)`

SetCardnumber sets Cardnumber field to given value.

### HasCardnumber

`func (o *CardTokenisationRequest) HasCardnumber() bool`

HasCardnumber returns a boolean if a field has been set.

### GetCsc

`func (o *CardTokenisationRequest) GetCsc() string`

GetCsc returns the Csc field if non-nil, zero value otherwise.

### GetCscOk

`func (o *CardTokenisationRequest) GetCscOk() (*string, bool)`

GetCscOk returns a tuple with the Csc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsc

`func (o *CardTokenisationRequest) SetCsc(v string)`

SetCsc sets Csc field to given value.

### HasCsc

`func (o *CardTokenisationRequest) HasCsc() bool`

HasCsc returns a boolean if a field has been set.

### GetExpmonth

`func (o *CardTokenisationRequest) GetExpmonth() int32`

GetExpmonth returns the Expmonth field if non-nil, zero value otherwise.

### GetExpmonthOk

`func (o *CardTokenisationRequest) GetExpmonthOk() (*int32, bool)`

GetExpmonthOk returns a tuple with the Expmonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpmonth

`func (o *CardTokenisationRequest) SetExpmonth(v int32)`

SetExpmonth sets Expmonth field to given value.

### HasExpmonth

`func (o *CardTokenisationRequest) HasExpmonth() bool`

HasExpmonth returns a boolean if a field has been set.

### GetExpyear

`func (o *CardTokenisationRequest) GetExpyear() int32`

GetExpyear returns the Expyear field if non-nil, zero value otherwise.

### GetExpyearOk

`func (o *CardTokenisationRequest) GetExpyearOk() (*int32, bool)`

GetExpyearOk returns a tuple with the Expyear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpyear

`func (o *CardTokenisationRequest) SetExpyear(v int32)`

SetExpyear sets Expyear field to given value.

### HasExpyear

`func (o *CardTokenisationRequest) HasExpyear() bool`

HasExpyear returns a boolean if a field has been set.

### GetNameOnCard

`func (o *CardTokenisationRequest) GetNameOnCard() string`

GetNameOnCard returns the NameOnCard field if non-nil, zero value otherwise.

### GetNameOnCardOk

`func (o *CardTokenisationRequest) GetNameOnCardOk() (*string, bool)`

GetNameOnCardOk returns a tuple with the NameOnCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCard

`func (o *CardTokenisationRequest) SetNameOnCard(v string)`

SetNameOnCard sets NameOnCard field to given value.

### HasNameOnCard

`func (o *CardTokenisationRequest) HasNameOnCard() bool`

HasNameOnCard returns a boolean if a field has been set.

### GetUuid

`func (o *CardTokenisationRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CardTokenisationRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CardTokenisationRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CardTokenisationRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


