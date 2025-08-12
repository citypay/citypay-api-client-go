# Bin

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
**Scheme** | Pointer to **string** | The scheme that issued the card. | [optional] 

## Methods

### NewBin

`func NewBin() *Bin`

NewBin instantiates a new Bin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinWithDefaults

`func NewBinWithDefaults() *Bin`

NewBinWithDefaults instantiates a new Bin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinCommercial

`func (o *Bin) GetBinCommercial() bool`

GetBinCommercial returns the BinCommercial field if non-nil, zero value otherwise.

### GetBinCommercialOk

`func (o *Bin) GetBinCommercialOk() (*bool, bool)`

GetBinCommercialOk returns a tuple with the BinCommercial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinCommercial

`func (o *Bin) SetBinCommercial(v bool)`

SetBinCommercial sets BinCommercial field to given value.

### HasBinCommercial

`func (o *Bin) HasBinCommercial() bool`

HasBinCommercial returns a boolean if a field has been set.

### GetBinCorporate

`func (o *Bin) GetBinCorporate() bool`

GetBinCorporate returns the BinCorporate field if non-nil, zero value otherwise.

### GetBinCorporateOk

`func (o *Bin) GetBinCorporateOk() (*bool, bool)`

GetBinCorporateOk returns a tuple with the BinCorporate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinCorporate

`func (o *Bin) SetBinCorporate(v bool)`

SetBinCorporate sets BinCorporate field to given value.

### HasBinCorporate

`func (o *Bin) HasBinCorporate() bool`

HasBinCorporate returns a boolean if a field has been set.

### GetBinCountryIssued

`func (o *Bin) GetBinCountryIssued() string`

GetBinCountryIssued returns the BinCountryIssued field if non-nil, zero value otherwise.

### GetBinCountryIssuedOk

`func (o *Bin) GetBinCountryIssuedOk() (*string, bool)`

GetBinCountryIssuedOk returns a tuple with the BinCountryIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinCountryIssued

`func (o *Bin) SetBinCountryIssued(v string)`

SetBinCountryIssued sets BinCountryIssued field to given value.

### HasBinCountryIssued

`func (o *Bin) HasBinCountryIssued() bool`

HasBinCountryIssued returns a boolean if a field has been set.

### GetBinCredit

`func (o *Bin) GetBinCredit() bool`

GetBinCredit returns the BinCredit field if non-nil, zero value otherwise.

### GetBinCreditOk

`func (o *Bin) GetBinCreditOk() (*bool, bool)`

GetBinCreditOk returns a tuple with the BinCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinCredit

`func (o *Bin) SetBinCredit(v bool)`

SetBinCredit sets BinCredit field to given value.

### HasBinCredit

`func (o *Bin) HasBinCredit() bool`

HasBinCredit returns a boolean if a field has been set.

### GetBinCurrency

`func (o *Bin) GetBinCurrency() string`

GetBinCurrency returns the BinCurrency field if non-nil, zero value otherwise.

### GetBinCurrencyOk

`func (o *Bin) GetBinCurrencyOk() (*string, bool)`

GetBinCurrencyOk returns a tuple with the BinCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinCurrency

`func (o *Bin) SetBinCurrency(v string)`

SetBinCurrency sets BinCurrency field to given value.

### HasBinCurrency

`func (o *Bin) HasBinCurrency() bool`

HasBinCurrency returns a boolean if a field has been set.

### GetBinDebit

`func (o *Bin) GetBinDebit() bool`

GetBinDebit returns the BinDebit field if non-nil, zero value otherwise.

### GetBinDebitOk

`func (o *Bin) GetBinDebitOk() (*bool, bool)`

GetBinDebitOk returns a tuple with the BinDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinDebit

`func (o *Bin) SetBinDebit(v bool)`

SetBinDebit sets BinDebit field to given value.

### HasBinDebit

`func (o *Bin) HasBinDebit() bool`

HasBinDebit returns a boolean if a field has been set.

### GetBinDescription

`func (o *Bin) GetBinDescription() string`

GetBinDescription returns the BinDescription field if non-nil, zero value otherwise.

### GetBinDescriptionOk

`func (o *Bin) GetBinDescriptionOk() (*string, bool)`

GetBinDescriptionOk returns a tuple with the BinDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinDescription

`func (o *Bin) SetBinDescription(v string)`

SetBinDescription sets BinDescription field to given value.

### HasBinDescription

`func (o *Bin) HasBinDescription() bool`

HasBinDescription returns a boolean if a field has been set.

### GetBinEu

`func (o *Bin) GetBinEu() bool`

GetBinEu returns the BinEu field if non-nil, zero value otherwise.

### GetBinEuOk

`func (o *Bin) GetBinEuOk() (*bool, bool)`

GetBinEuOk returns a tuple with the BinEu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinEu

`func (o *Bin) SetBinEu(v bool)`

SetBinEu sets BinEu field to given value.

### HasBinEu

`func (o *Bin) HasBinEu() bool`

HasBinEu returns a boolean if a field has been set.

### GetScheme

`func (o *Bin) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *Bin) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *Bin) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *Bin) HasScheme() bool`

HasScheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


