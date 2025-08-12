# Card

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinCommercial** | Pointer to **bool** | Defines whether the card is a commercial card. | [optional] 
**BinCorporate** | Pointer to **bool** | Defines whether the card is a corporate business card. | [optional] 
**BinCountryIssued** | Pointer to **string** | The determined country where the card was issued. | [optional] 
**BinCredit** | Pointer to **bool** | Defines whether the card is a credit card. | [optional] 
**BinCurrency** | Pointer to **string** | The default currency determined for the card. | [optional] 
**BinDebit** | Pointer to **bool** | Defines whether the card is a debit card. | [optional] 
**BinDescription** | Pointer to **string** | A description of the bin on the card to identify what type of product the card is. | [optional] 
**BinEu** | Pointer to **bool** | Defines whether the card is regulated within the EU. | [optional] 
**CardId** | Pointer to **string** | The id of the card that is returned. Should be used for referencing the card when perform any changes. | [optional] 
**CardStatus** | Pointer to **string** | The status of the card such, valid values are   - ACTIVE the card is active for processing   - INACTIVE the card is not active for processing   - EXPIRED for cards that have passed their expiry date.  | [optional] 
**DateCreated** | Pointer to **time.Time** | The date time of when the card was created. | [optional] 
**Default** | Pointer to **bool** | Determines if the card is the default card for the account and should be regarded as the first option to be used for processing. | [optional] 
**Expmonth** | Pointer to **int32** | The expiry month of the card. | [optional] 
**Expyear** | Pointer to **int32** | The expiry year of the card. | [optional] 
**Label** | Pointer to **string** | A label which identifies this card. | [optional] 
**Label2** | Pointer to **string** | A label which also provides the expiry date of the card. | [optional] 
**Last4digits** | Pointer to **string** | The last 4 digits of the card to aid in identification. | [optional] 
**NameOnCard** | Pointer to **string** | The name on the card. | [optional] 
**Scheme** | Pointer to **string** | The scheme that issued the card. | [optional] 
**Token** | Pointer to **string** | A token that can be used to process against the card. | [optional] 

## Methods

### NewCard

`func NewCard() *Card`

NewCard instantiates a new Card object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardWithDefaults

`func NewCardWithDefaults() *Card`

NewCardWithDefaults instantiates a new Card object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinCommercial

`func (o *Card) GetBinCommercial() bool`

GetBinCommercial returns the BinCommercial field if non-nil, zero value otherwise.

### GetBinCommercialOk

`func (o *Card) GetBinCommercialOk() (*bool, bool)`

GetBinCommercialOk returns a tuple with the BinCommercial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinCommercial

`func (o *Card) SetBinCommercial(v bool)`

SetBinCommercial sets BinCommercial field to given value.

### HasBinCommercial

`func (o *Card) HasBinCommercial() bool`

HasBinCommercial returns a boolean if a field has been set.

### GetBinCorporate

`func (o *Card) GetBinCorporate() bool`

GetBinCorporate returns the BinCorporate field if non-nil, zero value otherwise.

### GetBinCorporateOk

`func (o *Card) GetBinCorporateOk() (*bool, bool)`

GetBinCorporateOk returns a tuple with the BinCorporate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinCorporate

`func (o *Card) SetBinCorporate(v bool)`

SetBinCorporate sets BinCorporate field to given value.

### HasBinCorporate

`func (o *Card) HasBinCorporate() bool`

HasBinCorporate returns a boolean if a field has been set.

### GetBinCountryIssued

`func (o *Card) GetBinCountryIssued() string`

GetBinCountryIssued returns the BinCountryIssued field if non-nil, zero value otherwise.

### GetBinCountryIssuedOk

`func (o *Card) GetBinCountryIssuedOk() (*string, bool)`

GetBinCountryIssuedOk returns a tuple with the BinCountryIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinCountryIssued

`func (o *Card) SetBinCountryIssued(v string)`

SetBinCountryIssued sets BinCountryIssued field to given value.

### HasBinCountryIssued

`func (o *Card) HasBinCountryIssued() bool`

HasBinCountryIssued returns a boolean if a field has been set.

### GetBinCredit

`func (o *Card) GetBinCredit() bool`

GetBinCredit returns the BinCredit field if non-nil, zero value otherwise.

### GetBinCreditOk

`func (o *Card) GetBinCreditOk() (*bool, bool)`

GetBinCreditOk returns a tuple with the BinCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinCredit

`func (o *Card) SetBinCredit(v bool)`

SetBinCredit sets BinCredit field to given value.

### HasBinCredit

`func (o *Card) HasBinCredit() bool`

HasBinCredit returns a boolean if a field has been set.

### GetBinCurrency

`func (o *Card) GetBinCurrency() string`

GetBinCurrency returns the BinCurrency field if non-nil, zero value otherwise.

### GetBinCurrencyOk

`func (o *Card) GetBinCurrencyOk() (*string, bool)`

GetBinCurrencyOk returns a tuple with the BinCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinCurrency

`func (o *Card) SetBinCurrency(v string)`

SetBinCurrency sets BinCurrency field to given value.

### HasBinCurrency

`func (o *Card) HasBinCurrency() bool`

HasBinCurrency returns a boolean if a field has been set.

### GetBinDebit

`func (o *Card) GetBinDebit() bool`

GetBinDebit returns the BinDebit field if non-nil, zero value otherwise.

### GetBinDebitOk

`func (o *Card) GetBinDebitOk() (*bool, bool)`

GetBinDebitOk returns a tuple with the BinDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinDebit

`func (o *Card) SetBinDebit(v bool)`

SetBinDebit sets BinDebit field to given value.

### HasBinDebit

`func (o *Card) HasBinDebit() bool`

HasBinDebit returns a boolean if a field has been set.

### GetBinDescription

`func (o *Card) GetBinDescription() string`

GetBinDescription returns the BinDescription field if non-nil, zero value otherwise.

### GetBinDescriptionOk

`func (o *Card) GetBinDescriptionOk() (*string, bool)`

GetBinDescriptionOk returns a tuple with the BinDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinDescription

`func (o *Card) SetBinDescription(v string)`

SetBinDescription sets BinDescription field to given value.

### HasBinDescription

`func (o *Card) HasBinDescription() bool`

HasBinDescription returns a boolean if a field has been set.

### GetBinEu

`func (o *Card) GetBinEu() bool`

GetBinEu returns the BinEu field if non-nil, zero value otherwise.

### GetBinEuOk

`func (o *Card) GetBinEuOk() (*bool, bool)`

GetBinEuOk returns a tuple with the BinEu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinEu

`func (o *Card) SetBinEu(v bool)`

SetBinEu sets BinEu field to given value.

### HasBinEu

`func (o *Card) HasBinEu() bool`

HasBinEu returns a boolean if a field has been set.

### GetCardId

`func (o *Card) GetCardId() string`

GetCardId returns the CardId field if non-nil, zero value otherwise.

### GetCardIdOk

`func (o *Card) GetCardIdOk() (*string, bool)`

GetCardIdOk returns a tuple with the CardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardId

`func (o *Card) SetCardId(v string)`

SetCardId sets CardId field to given value.

### HasCardId

`func (o *Card) HasCardId() bool`

HasCardId returns a boolean if a field has been set.

### GetCardStatus

`func (o *Card) GetCardStatus() string`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *Card) GetCardStatusOk() (*string, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *Card) SetCardStatus(v string)`

SetCardStatus sets CardStatus field to given value.

### HasCardStatus

`func (o *Card) HasCardStatus() bool`

HasCardStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *Card) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Card) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Card) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Card) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDefault

`func (o *Card) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Card) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Card) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Card) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetExpmonth

`func (o *Card) GetExpmonth() int32`

GetExpmonth returns the Expmonth field if non-nil, zero value otherwise.

### GetExpmonthOk

`func (o *Card) GetExpmonthOk() (*int32, bool)`

GetExpmonthOk returns a tuple with the Expmonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpmonth

`func (o *Card) SetExpmonth(v int32)`

SetExpmonth sets Expmonth field to given value.

### HasExpmonth

`func (o *Card) HasExpmonth() bool`

HasExpmonth returns a boolean if a field has been set.

### GetExpyear

`func (o *Card) GetExpyear() int32`

GetExpyear returns the Expyear field if non-nil, zero value otherwise.

### GetExpyearOk

`func (o *Card) GetExpyearOk() (*int32, bool)`

GetExpyearOk returns a tuple with the Expyear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpyear

`func (o *Card) SetExpyear(v int32)`

SetExpyear sets Expyear field to given value.

### HasExpyear

`func (o *Card) HasExpyear() bool`

HasExpyear returns a boolean if a field has been set.

### GetLabel

`func (o *Card) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Card) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Card) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Card) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLabel2

`func (o *Card) GetLabel2() string`

GetLabel2 returns the Label2 field if non-nil, zero value otherwise.

### GetLabel2Ok

`func (o *Card) GetLabel2Ok() (*string, bool)`

GetLabel2Ok returns a tuple with the Label2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel2

`func (o *Card) SetLabel2(v string)`

SetLabel2 sets Label2 field to given value.

### HasLabel2

`func (o *Card) HasLabel2() bool`

HasLabel2 returns a boolean if a field has been set.

### GetLast4digits

`func (o *Card) GetLast4digits() string`

GetLast4digits returns the Last4digits field if non-nil, zero value otherwise.

### GetLast4digitsOk

`func (o *Card) GetLast4digitsOk() (*string, bool)`

GetLast4digitsOk returns a tuple with the Last4digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4digits

`func (o *Card) SetLast4digits(v string)`

SetLast4digits sets Last4digits field to given value.

### HasLast4digits

`func (o *Card) HasLast4digits() bool`

HasLast4digits returns a boolean if a field has been set.

### GetNameOnCard

`func (o *Card) GetNameOnCard() string`

GetNameOnCard returns the NameOnCard field if non-nil, zero value otherwise.

### GetNameOnCardOk

`func (o *Card) GetNameOnCardOk() (*string, bool)`

GetNameOnCardOk returns a tuple with the NameOnCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameOnCard

`func (o *Card) SetNameOnCard(v string)`

SetNameOnCard sets NameOnCard field to given value.

### HasNameOnCard

`func (o *Card) HasNameOnCard() bool`

HasNameOnCard returns a boolean if a field has been set.

### GetScheme

`func (o *Card) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *Card) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *Card) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *Card) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetToken

`func (o *Card) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Card) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Card) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Card) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


