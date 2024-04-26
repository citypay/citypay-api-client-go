# ContactDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address1** | Pointer to **string** | The first line of the address for the card holder. | [optional] 
**Address2** | Pointer to **string** | The second line of the address for the card holder. | [optional] 
**Address3** | Pointer to **string** | The third line of the address for the card holder. | [optional] 
**Area** | Pointer to **string** | The area such as city, department, parish for the card holder. | [optional] 
**Company** | Pointer to **string** | The company name for the card holder if the contact is a corporate contact. | [optional] 
**Country** | Pointer to **string** | The country code in ISO 3166 format. The country value may be used for fraud analysis and for   acceptance of the transaction.  | [optional] 
**Email** | Pointer to **string** | An email address for the card holder which may be used for correspondence. | [optional] 
**Firstname** | Pointer to **string** | The first name  of the card holder. | [optional] 
**Lastname** | Pointer to **string** | The last name or surname of the card holder. | [optional] 
**MobileNo** | Pointer to **string** | A mobile number for the card holder the mobile number is often required by delivery companies to ensure they are able to be in contact when required. | [optional] 
**Postcode** | Pointer to **string** | The postcode or zip code of the address which may be used for fraud analysis. | [optional] 
**TelephoneNo** | Pointer to **string** | A telephone number for the card holder. | [optional] 
**Title** | Pointer to **string** | A title for the card holder such as Mr, Mrs, Ms, M. Mme. etc. | [optional] 

## Methods

### NewContactDetails

`func NewContactDetails() *ContactDetails`

NewContactDetails instantiates a new ContactDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactDetailsWithDefaults

`func NewContactDetailsWithDefaults() *ContactDetails`

NewContactDetailsWithDefaults instantiates a new ContactDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress1

`func (o *ContactDetails) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *ContactDetails) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *ContactDetails) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *ContactDetails) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *ContactDetails) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *ContactDetails) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *ContactDetails) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *ContactDetails) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetAddress3

`func (o *ContactDetails) GetAddress3() string`

GetAddress3 returns the Address3 field if non-nil, zero value otherwise.

### GetAddress3Ok

`func (o *ContactDetails) GetAddress3Ok() (*string, bool)`

GetAddress3Ok returns a tuple with the Address3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress3

`func (o *ContactDetails) SetAddress3(v string)`

SetAddress3 sets Address3 field to given value.

### HasAddress3

`func (o *ContactDetails) HasAddress3() bool`

HasAddress3 returns a boolean if a field has been set.

### GetArea

`func (o *ContactDetails) GetArea() string`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ContactDetails) GetAreaOk() (*string, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ContactDetails) SetArea(v string)`

SetArea sets Area field to given value.

### HasArea

`func (o *ContactDetails) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetCompany

`func (o *ContactDetails) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ContactDetails) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ContactDetails) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ContactDetails) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCountry

`func (o *ContactDetails) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ContactDetails) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ContactDetails) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ContactDetails) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *ContactDetails) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactDetails) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactDetails) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ContactDetails) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstname

`func (o *ContactDetails) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *ContactDetails) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *ContactDetails) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *ContactDetails) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *ContactDetails) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *ContactDetails) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *ContactDetails) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *ContactDetails) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetMobileNo

`func (o *ContactDetails) GetMobileNo() string`

GetMobileNo returns the MobileNo field if non-nil, zero value otherwise.

### GetMobileNoOk

`func (o *ContactDetails) GetMobileNoOk() (*string, bool)`

GetMobileNoOk returns a tuple with the MobileNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNo

`func (o *ContactDetails) SetMobileNo(v string)`

SetMobileNo sets MobileNo field to given value.

### HasMobileNo

`func (o *ContactDetails) HasMobileNo() bool`

HasMobileNo returns a boolean if a field has been set.

### GetPostcode

`func (o *ContactDetails) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *ContactDetails) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *ContactDetails) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.

### HasPostcode

`func (o *ContactDetails) HasPostcode() bool`

HasPostcode returns a boolean if a field has been set.

### GetTelephoneNo

`func (o *ContactDetails) GetTelephoneNo() string`

GetTelephoneNo returns the TelephoneNo field if non-nil, zero value otherwise.

### GetTelephoneNoOk

`func (o *ContactDetails) GetTelephoneNoOk() (*string, bool)`

GetTelephoneNoOk returns a tuple with the TelephoneNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephoneNo

`func (o *ContactDetails) SetTelephoneNo(v string)`

SetTelephoneNo sets TelephoneNo field to given value.

### HasTelephoneNo

`func (o *ContactDetails) HasTelephoneNo() bool`

HasTelephoneNo returns a boolean if a field has been set.

### GetTitle

`func (o *ContactDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ContactDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ContactDetails) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ContactDetails) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


