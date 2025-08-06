# CardTokenisationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpCardToken** | **string** | The tokenised card value. The token is encrypted with integrity checks and scoped to a client id only allowing for the card value to be used.  The value may be used up and until the expiry date of the card. | 
**Last4digits** | Pointer to **string** | The last 4 digits of the card. | [optional] 
**Scheme** | Pointer to **string** | The card scheme of the card. | [optional] 
**SchemeLogo** | Pointer to **string** | The url of the logo card scheme of the card. | [optional] 

## Methods

### NewCardTokenisationResponse

`func NewCardTokenisationResponse(cpCardToken string, ) *CardTokenisationResponse`

NewCardTokenisationResponse instantiates a new CardTokenisationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardTokenisationResponseWithDefaults

`func NewCardTokenisationResponseWithDefaults() *CardTokenisationResponse`

NewCardTokenisationResponseWithDefaults instantiates a new CardTokenisationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpCardToken

`func (o *CardTokenisationResponse) GetCpCardToken() string`

GetCpCardToken returns the CpCardToken field if non-nil, zero value otherwise.

### GetCpCardTokenOk

`func (o *CardTokenisationResponse) GetCpCardTokenOk() (*string, bool)`

GetCpCardTokenOk returns a tuple with the CpCardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpCardToken

`func (o *CardTokenisationResponse) SetCpCardToken(v string)`

SetCpCardToken sets CpCardToken field to given value.


### GetLast4digits

`func (o *CardTokenisationResponse) GetLast4digits() string`

GetLast4digits returns the Last4digits field if non-nil, zero value otherwise.

### GetLast4digitsOk

`func (o *CardTokenisationResponse) GetLast4digitsOk() (*string, bool)`

GetLast4digitsOk returns a tuple with the Last4digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4digits

`func (o *CardTokenisationResponse) SetLast4digits(v string)`

SetLast4digits sets Last4digits field to given value.

### HasLast4digits

`func (o *CardTokenisationResponse) HasLast4digits() bool`

HasLast4digits returns a boolean if a field has been set.

### GetScheme

`func (o *CardTokenisationResponse) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *CardTokenisationResponse) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *CardTokenisationResponse) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *CardTokenisationResponse) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetSchemeLogo

`func (o *CardTokenisationResponse) GetSchemeLogo() string`

GetSchemeLogo returns the SchemeLogo field if non-nil, zero value otherwise.

### GetSchemeLogoOk

`func (o *CardTokenisationResponse) GetSchemeLogoOk() (*string, bool)`

GetSchemeLogoOk returns a tuple with the SchemeLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeLogo

`func (o *CardTokenisationResponse) SetSchemeLogo(v string)`

SetSchemeLogo sets SchemeLogo field to given value.

### HasSchemeLogo

`func (o *CardTokenisationResponse) HasSchemeLogo() bool`

HasSchemeLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


