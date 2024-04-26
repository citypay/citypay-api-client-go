# CardHolderAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The account id of the card holder account provided by the merchant which uniquely identifies the account.  | 
**Cards** | Pointer to [**[]Card**](Card.md) |  | [optional] 
**Contact** | [**ContactDetails**](ContactDetails.md) |  | 
**DateCreated** | Pointer to **time.Time** | The date and time the account was created. | [optional] 
**DefaultCardId** | Pointer to **string** | The id of the default card. | [optional] 
**DefaultCardIndex** | Pointer to **int32** | The index in the array of the default card. | [optional] 
**LastModified** | Pointer to **time.Time** | The date and time the account was last modified. | [optional] 
**Status** | Pointer to **string** | Defines the status of the account for processing valid values are   - ACTIVE for active accounts that are able to process   - DISABLED for accounts that are currently disabled for processing.  | [optional] 
**UniqueId** | Pointer to **string** | A unique id of the card holder account which uniquely identifies the stored account. This value is not searchable. | [optional] 

## Methods

### NewCardHolderAccount

`func NewCardHolderAccount(accountId string, contact ContactDetails, ) *CardHolderAccount`

NewCardHolderAccount instantiates a new CardHolderAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardHolderAccountWithDefaults

`func NewCardHolderAccountWithDefaults() *CardHolderAccount`

NewCardHolderAccountWithDefaults instantiates a new CardHolderAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CardHolderAccount) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CardHolderAccount) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CardHolderAccount) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCards

`func (o *CardHolderAccount) GetCards() []Card`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *CardHolderAccount) GetCardsOk() (*[]Card, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *CardHolderAccount) SetCards(v []Card)`

SetCards sets Cards field to given value.

### HasCards

`func (o *CardHolderAccount) HasCards() bool`

HasCards returns a boolean if a field has been set.

### GetContact

`func (o *CardHolderAccount) GetContact() ContactDetails`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *CardHolderAccount) GetContactOk() (*ContactDetails, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *CardHolderAccount) SetContact(v ContactDetails)`

SetContact sets Contact field to given value.


### GetDateCreated

`func (o *CardHolderAccount) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *CardHolderAccount) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *CardHolderAccount) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *CardHolderAccount) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDefaultCardId

`func (o *CardHolderAccount) GetDefaultCardId() string`

GetDefaultCardId returns the DefaultCardId field if non-nil, zero value otherwise.

### GetDefaultCardIdOk

`func (o *CardHolderAccount) GetDefaultCardIdOk() (*string, bool)`

GetDefaultCardIdOk returns a tuple with the DefaultCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCardId

`func (o *CardHolderAccount) SetDefaultCardId(v string)`

SetDefaultCardId sets DefaultCardId field to given value.

### HasDefaultCardId

`func (o *CardHolderAccount) HasDefaultCardId() bool`

HasDefaultCardId returns a boolean if a field has been set.

### GetDefaultCardIndex

`func (o *CardHolderAccount) GetDefaultCardIndex() int32`

GetDefaultCardIndex returns the DefaultCardIndex field if non-nil, zero value otherwise.

### GetDefaultCardIndexOk

`func (o *CardHolderAccount) GetDefaultCardIndexOk() (*int32, bool)`

GetDefaultCardIndexOk returns a tuple with the DefaultCardIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCardIndex

`func (o *CardHolderAccount) SetDefaultCardIndex(v int32)`

SetDefaultCardIndex sets DefaultCardIndex field to given value.

### HasDefaultCardIndex

`func (o *CardHolderAccount) HasDefaultCardIndex() bool`

HasDefaultCardIndex returns a boolean if a field has been set.

### GetLastModified

`func (o *CardHolderAccount) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *CardHolderAccount) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *CardHolderAccount) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *CardHolderAccount) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetStatus

`func (o *CardHolderAccount) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CardHolderAccount) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CardHolderAccount) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CardHolderAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUniqueId

`func (o *CardHolderAccount) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *CardHolderAccount) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *CardHolderAccount) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *CardHolderAccount) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


