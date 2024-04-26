# AccountCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | A card holder account id used for uniquely identifying the account. This value will be used for future referencing of the account oand to link your system to this API. This value is immutable and never changes.  | 
**Contact** | Pointer to [**ContactDetails**](ContactDetails.md) |  | [optional] 

## Methods

### NewAccountCreate

`func NewAccountCreate(accountId string, ) *AccountCreate`

NewAccountCreate instantiates a new AccountCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreateWithDefaults

`func NewAccountCreateWithDefaults() *AccountCreate`

NewAccountCreateWithDefaults instantiates a new AccountCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountCreate) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountCreate) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountCreate) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetContact

`func (o *AccountCreate) GetContact() ContactDetails`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *AccountCreate) GetContactOk() (*ContactDetails, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *AccountCreate) SetContact(v ContactDetails)`

SetContact sets Contact field to given value.

### HasContact

`func (o *AccountCreate) HasContact() bool`

HasContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


