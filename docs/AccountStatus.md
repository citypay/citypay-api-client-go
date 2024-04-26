# AccountStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The status of the account to set, valid values are ACTIVE or DISABLED. | [optional] 

## Methods

### NewAccountStatus

`func NewAccountStatus() *AccountStatus`

NewAccountStatus instantiates a new AccountStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountStatusWithDefaults

`func NewAccountStatusWithDefaults() *AccountStatus`

NewAccountStatusWithDefaults instantiates a new AccountStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AccountStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


