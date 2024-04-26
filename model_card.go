/*
CityPay Payment API

 Welcome to the CityPay API, a robust HTTP API payment solution designed for seamless server-to-server  transactional processing. Our API facilitates a wide array of payment operations, catering to diverse business needs.  Whether you're integrating Internet payments, handling Mail Order/Telephone Order (MOTO) transactions, managing  Subscriptions with Recurring and Continuous Authority payments, or navigating the complexities of 3-D Secure  authentication, our API is equipped to support your requirements. Additionally, we offer functionalities for  Authorisation, Refunding, Pre-Authorisation, Cancellation/Voids, and Completion processing, alongside the capability  for tokenised payments.  ## Compliance and Security Overview <aside class=\"notice\">   Ensuring the security of payment transactions and compliance with industry standards is paramount. Our API is    designed with stringent security measures and compliance protocols to safeguard sensitive information and meet    the rigorous requirements of Visa, MasterCard, and the PCI Security Standards Council. </aside>  ### Key Compliance and Security Measures  * **TLS Encryption**: All data transmissions must utilise TLS version 1.2 or higher, employing [strong cryptography](#enabled-tls-ciphers). Our infrastructure strictly enforces this requirement to maintain the integrity and confidentiality of data in transit. We conduct regular scans and assessments of our TLS endpoints to identify and mitigate vulnerabilities. * **Data Storage Prohibitions**: Storing sensitive cardholder data (CHD), such as the card security code (CSC) or primary account number (PAN), is strictly prohibited. Our API is designed to minimize your exposure to sensitive data, thereby reducing your compliance burden. * **Data Masking**: For consumer protection and compliance, full card numbers must not be displayed on receipts or any customer-facing materials. Our API automatically masks PANs, displaying only the last four digits to facilitate safe receipt generation. * **Network Scans**: If your application is web-based, regular scans of your hosting environment are mandatory to identify and rectify potential vulnerabilities. This proactive measure is crucial for maintaining a secure and compliant online presence. * **PCI Compliance**: Adherence to PCI DSS standards is not optional; it's a requirement for operating securely and legally in the payments ecosystem. For detailed information on compliance requirements and resources, please visit the PCI Security Standards Council website [https://www.pcisecuritystandards.org/](https://www.pcisecuritystandards.org/). * **Request Validation**: Our API includes mechanisms to verify the legitimacy of each request, ensuring it pertains to a valid account and originates from a trusted source. We leverage remote IP address verification alongside sophisticated application firewall technologies to thwart a wide array of common security threats.  ## Getting Started Before integrating with the CityPay API, ensure your application and development practices align with the outlined compliance and security measures. This preparatory step is crucial for a smooth integration process and the long-term success of your payment processing operations.  For further details on API endpoints, request/response formats, and code examples, proceed to the subsequent sections of our documentation. Our aim is to provide you with all the necessary tools and information to integrate our payment processing capabilities seamlessly into your application.  Thank you for choosing CityPay API. We look forward to supporting your payment processing needs with our secure, compliant, and versatile API solution.

API version: 6.6.40
Contact: support@citypay.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citypay

import (
	"encoding/json"
	"time"
)

// checks if the Card type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Card{}

// Card struct for Card
type Card struct {
	// Defines whether the card is a commercial card.
	BinCommercial *bool `json:"bin_commercial,omitempty"`
	// Defines whether the card is a corporate business card.
	BinCorporate *bool `json:"bin_corporate,omitempty"`
	// The determined country where the card was issued.
	BinCountryIssued *string `json:"bin_country_issued,omitempty"`
	// Defines whether the card is a credit card.
	BinCredit *bool `json:"bin_credit,omitempty"`
	// The default currency determined for the card.
	BinCurrency *string `json:"bin_currency,omitempty"`
	// Defines whether the card is a debit card.
	BinDebit *bool `json:"bin_debit,omitempty"`
	// A description of the bin on the card to identify what type of product the card is.
	BinDescription *string `json:"bin_description,omitempty"`
	// Defines whether the card is regulated within the EU.
	BinEu *bool `json:"bin_eu,omitempty"`
	// The id of the card that is returned. Should be used for referencing the card when perform any changes.
	CardId *string `json:"card_id,omitempty"`
	// The status of the card such, valid values are   - ACTIVE the card is active for processing   - INACTIVE the card is not active for processing   - EXPIRED for cards that have passed their expiry date.
	CardStatus *string `json:"card_status,omitempty"`
	// The date time of when the card was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// Determines if the card is the default card for the account and should be regarded as the first option to be used for processing.
	Default *bool `json:"default,omitempty"`
	// The expiry month of the card.
	Expmonth *int32 `json:"expmonth,omitempty"`
	// The expiry year of the card.
	Expyear *int32 `json:"expyear,omitempty"`
	// A label which identifies this card.
	Label *string `json:"label,omitempty"`
	// A label which also provides the expiry date of the card.
	Label2 *string `json:"label2,omitempty"`
	// The last 4 digits of the card to aid in identification.
	Last4digits *string `json:"last4digits,omitempty"`
	// The name on the card.
	NameOnCard *string `json:"name_on_card,omitempty"`
	// The scheme that issued the card.
	Scheme *string `json:"scheme,omitempty"`
	// A token that can be used to process against the card.
	Token *string `json:"token,omitempty"`
}

// NewCard instantiates a new Card object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCard() *Card {
	this := Card{}
	return &this
}

// NewCardWithDefaults instantiates a new Card object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardWithDefaults() *Card {
	this := Card{}
	return &this
}

// GetBinCommercial returns the BinCommercial field value if set, zero value otherwise.
func (o *Card) GetBinCommercial() bool {
	if o == nil || IsNil(o.BinCommercial) {
		var ret bool
		return ret
	}
	return *o.BinCommercial
}

// GetBinCommercialOk returns a tuple with the BinCommercial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetBinCommercialOk() (*bool, bool) {
	if o == nil || IsNil(o.BinCommercial) {
		return nil, false
	}
	return o.BinCommercial, true
}

// HasBinCommercial returns a boolean if a field has been set.
func (o *Card) HasBinCommercial() bool {
	if o != nil && !IsNil(o.BinCommercial) {
		return true
	}

	return false
}

// SetBinCommercial gets a reference to the given bool and assigns it to the BinCommercial field.
func (o *Card) SetBinCommercial(v bool) {
	o.BinCommercial = &v
}

// GetBinCorporate returns the BinCorporate field value if set, zero value otherwise.
func (o *Card) GetBinCorporate() bool {
	if o == nil || IsNil(o.BinCorporate) {
		var ret bool
		return ret
	}
	return *o.BinCorporate
}

// GetBinCorporateOk returns a tuple with the BinCorporate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetBinCorporateOk() (*bool, bool) {
	if o == nil || IsNil(o.BinCorporate) {
		return nil, false
	}
	return o.BinCorporate, true
}

// HasBinCorporate returns a boolean if a field has been set.
func (o *Card) HasBinCorporate() bool {
	if o != nil && !IsNil(o.BinCorporate) {
		return true
	}

	return false
}

// SetBinCorporate gets a reference to the given bool and assigns it to the BinCorporate field.
func (o *Card) SetBinCorporate(v bool) {
	o.BinCorporate = &v
}

// GetBinCountryIssued returns the BinCountryIssued field value if set, zero value otherwise.
func (o *Card) GetBinCountryIssued() string {
	if o == nil || IsNil(o.BinCountryIssued) {
		var ret string
		return ret
	}
	return *o.BinCountryIssued
}

// GetBinCountryIssuedOk returns a tuple with the BinCountryIssued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetBinCountryIssuedOk() (*string, bool) {
	if o == nil || IsNil(o.BinCountryIssued) {
		return nil, false
	}
	return o.BinCountryIssued, true
}

// HasBinCountryIssued returns a boolean if a field has been set.
func (o *Card) HasBinCountryIssued() bool {
	if o != nil && !IsNil(o.BinCountryIssued) {
		return true
	}

	return false
}

// SetBinCountryIssued gets a reference to the given string and assigns it to the BinCountryIssued field.
func (o *Card) SetBinCountryIssued(v string) {
	o.BinCountryIssued = &v
}

// GetBinCredit returns the BinCredit field value if set, zero value otherwise.
func (o *Card) GetBinCredit() bool {
	if o == nil || IsNil(o.BinCredit) {
		var ret bool
		return ret
	}
	return *o.BinCredit
}

// GetBinCreditOk returns a tuple with the BinCredit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetBinCreditOk() (*bool, bool) {
	if o == nil || IsNil(o.BinCredit) {
		return nil, false
	}
	return o.BinCredit, true
}

// HasBinCredit returns a boolean if a field has been set.
func (o *Card) HasBinCredit() bool {
	if o != nil && !IsNil(o.BinCredit) {
		return true
	}

	return false
}

// SetBinCredit gets a reference to the given bool and assigns it to the BinCredit field.
func (o *Card) SetBinCredit(v bool) {
	o.BinCredit = &v
}

// GetBinCurrency returns the BinCurrency field value if set, zero value otherwise.
func (o *Card) GetBinCurrency() string {
	if o == nil || IsNil(o.BinCurrency) {
		var ret string
		return ret
	}
	return *o.BinCurrency
}

// GetBinCurrencyOk returns a tuple with the BinCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetBinCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.BinCurrency) {
		return nil, false
	}
	return o.BinCurrency, true
}

// HasBinCurrency returns a boolean if a field has been set.
func (o *Card) HasBinCurrency() bool {
	if o != nil && !IsNil(o.BinCurrency) {
		return true
	}

	return false
}

// SetBinCurrency gets a reference to the given string and assigns it to the BinCurrency field.
func (o *Card) SetBinCurrency(v string) {
	o.BinCurrency = &v
}

// GetBinDebit returns the BinDebit field value if set, zero value otherwise.
func (o *Card) GetBinDebit() bool {
	if o == nil || IsNil(o.BinDebit) {
		var ret bool
		return ret
	}
	return *o.BinDebit
}

// GetBinDebitOk returns a tuple with the BinDebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetBinDebitOk() (*bool, bool) {
	if o == nil || IsNil(o.BinDebit) {
		return nil, false
	}
	return o.BinDebit, true
}

// HasBinDebit returns a boolean if a field has been set.
func (o *Card) HasBinDebit() bool {
	if o != nil && !IsNil(o.BinDebit) {
		return true
	}

	return false
}

// SetBinDebit gets a reference to the given bool and assigns it to the BinDebit field.
func (o *Card) SetBinDebit(v bool) {
	o.BinDebit = &v
}

// GetBinDescription returns the BinDescription field value if set, zero value otherwise.
func (o *Card) GetBinDescription() string {
	if o == nil || IsNil(o.BinDescription) {
		var ret string
		return ret
	}
	return *o.BinDescription
}

// GetBinDescriptionOk returns a tuple with the BinDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetBinDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.BinDescription) {
		return nil, false
	}
	return o.BinDescription, true
}

// HasBinDescription returns a boolean if a field has been set.
func (o *Card) HasBinDescription() bool {
	if o != nil && !IsNil(o.BinDescription) {
		return true
	}

	return false
}

// SetBinDescription gets a reference to the given string and assigns it to the BinDescription field.
func (o *Card) SetBinDescription(v string) {
	o.BinDescription = &v
}

// GetBinEu returns the BinEu field value if set, zero value otherwise.
func (o *Card) GetBinEu() bool {
	if o == nil || IsNil(o.BinEu) {
		var ret bool
		return ret
	}
	return *o.BinEu
}

// GetBinEuOk returns a tuple with the BinEu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetBinEuOk() (*bool, bool) {
	if o == nil || IsNil(o.BinEu) {
		return nil, false
	}
	return o.BinEu, true
}

// HasBinEu returns a boolean if a field has been set.
func (o *Card) HasBinEu() bool {
	if o != nil && !IsNil(o.BinEu) {
		return true
	}

	return false
}

// SetBinEu gets a reference to the given bool and assigns it to the BinEu field.
func (o *Card) SetBinEu(v bool) {
	o.BinEu = &v
}

// GetCardId returns the CardId field value if set, zero value otherwise.
func (o *Card) GetCardId() string {
	if o == nil || IsNil(o.CardId) {
		var ret string
		return ret
	}
	return *o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetCardIdOk() (*string, bool) {
	if o == nil || IsNil(o.CardId) {
		return nil, false
	}
	return o.CardId, true
}

// HasCardId returns a boolean if a field has been set.
func (o *Card) HasCardId() bool {
	if o != nil && !IsNil(o.CardId) {
		return true
	}

	return false
}

// SetCardId gets a reference to the given string and assigns it to the CardId field.
func (o *Card) SetCardId(v string) {
	o.CardId = &v
}

// GetCardStatus returns the CardStatus field value if set, zero value otherwise.
func (o *Card) GetCardStatus() string {
	if o == nil || IsNil(o.CardStatus) {
		var ret string
		return ret
	}
	return *o.CardStatus
}

// GetCardStatusOk returns a tuple with the CardStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetCardStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CardStatus) {
		return nil, false
	}
	return o.CardStatus, true
}

// HasCardStatus returns a boolean if a field has been set.
func (o *Card) HasCardStatus() bool {
	if o != nil && !IsNil(o.CardStatus) {
		return true
	}

	return false
}

// SetCardStatus gets a reference to the given string and assigns it to the CardStatus field.
func (o *Card) SetCardStatus(v string) {
	o.CardStatus = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *Card) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *Card) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *Card) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *Card) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *Card) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *Card) SetDefault(v bool) {
	o.Default = &v
}

// GetExpmonth returns the Expmonth field value if set, zero value otherwise.
func (o *Card) GetExpmonth() int32 {
	if o == nil || IsNil(o.Expmonth) {
		var ret int32
		return ret
	}
	return *o.Expmonth
}

// GetExpmonthOk returns a tuple with the Expmonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetExpmonthOk() (*int32, bool) {
	if o == nil || IsNil(o.Expmonth) {
		return nil, false
	}
	return o.Expmonth, true
}

// HasExpmonth returns a boolean if a field has been set.
func (o *Card) HasExpmonth() bool {
	if o != nil && !IsNil(o.Expmonth) {
		return true
	}

	return false
}

// SetExpmonth gets a reference to the given int32 and assigns it to the Expmonth field.
func (o *Card) SetExpmonth(v int32) {
	o.Expmonth = &v
}

// GetExpyear returns the Expyear field value if set, zero value otherwise.
func (o *Card) GetExpyear() int32 {
	if o == nil || IsNil(o.Expyear) {
		var ret int32
		return ret
	}
	return *o.Expyear
}

// GetExpyearOk returns a tuple with the Expyear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetExpyearOk() (*int32, bool) {
	if o == nil || IsNil(o.Expyear) {
		return nil, false
	}
	return o.Expyear, true
}

// HasExpyear returns a boolean if a field has been set.
func (o *Card) HasExpyear() bool {
	if o != nil && !IsNil(o.Expyear) {
		return true
	}

	return false
}

// SetExpyear gets a reference to the given int32 and assigns it to the Expyear field.
func (o *Card) SetExpyear(v int32) {
	o.Expyear = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Card) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Card) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Card) SetLabel(v string) {
	o.Label = &v
}

// GetLabel2 returns the Label2 field value if set, zero value otherwise.
func (o *Card) GetLabel2() string {
	if o == nil || IsNil(o.Label2) {
		var ret string
		return ret
	}
	return *o.Label2
}

// GetLabel2Ok returns a tuple with the Label2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetLabel2Ok() (*string, bool) {
	if o == nil || IsNil(o.Label2) {
		return nil, false
	}
	return o.Label2, true
}

// HasLabel2 returns a boolean if a field has been set.
func (o *Card) HasLabel2() bool {
	if o != nil && !IsNil(o.Label2) {
		return true
	}

	return false
}

// SetLabel2 gets a reference to the given string and assigns it to the Label2 field.
func (o *Card) SetLabel2(v string) {
	o.Label2 = &v
}

// GetLast4digits returns the Last4digits field value if set, zero value otherwise.
func (o *Card) GetLast4digits() string {
	if o == nil || IsNil(o.Last4digits) {
		var ret string
		return ret
	}
	return *o.Last4digits
}

// GetLast4digitsOk returns a tuple with the Last4digits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetLast4digitsOk() (*string, bool) {
	if o == nil || IsNil(o.Last4digits) {
		return nil, false
	}
	return o.Last4digits, true
}

// HasLast4digits returns a boolean if a field has been set.
func (o *Card) HasLast4digits() bool {
	if o != nil && !IsNil(o.Last4digits) {
		return true
	}

	return false
}

// SetLast4digits gets a reference to the given string and assigns it to the Last4digits field.
func (o *Card) SetLast4digits(v string) {
	o.Last4digits = &v
}

// GetNameOnCard returns the NameOnCard field value if set, zero value otherwise.
func (o *Card) GetNameOnCard() string {
	if o == nil || IsNil(o.NameOnCard) {
		var ret string
		return ret
	}
	return *o.NameOnCard
}

// GetNameOnCardOk returns a tuple with the NameOnCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetNameOnCardOk() (*string, bool) {
	if o == nil || IsNil(o.NameOnCard) {
		return nil, false
	}
	return o.NameOnCard, true
}

// HasNameOnCard returns a boolean if a field has been set.
func (o *Card) HasNameOnCard() bool {
	if o != nil && !IsNil(o.NameOnCard) {
		return true
	}

	return false
}

// SetNameOnCard gets a reference to the given string and assigns it to the NameOnCard field.
func (o *Card) SetNameOnCard(v string) {
	o.NameOnCard = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *Card) GetScheme() string {
	if o == nil || IsNil(o.Scheme) {
		var ret string
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetSchemeOk() (*string, bool) {
	if o == nil || IsNil(o.Scheme) {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *Card) HasScheme() bool {
	if o != nil && !IsNil(o.Scheme) {
		return true
	}

	return false
}

// SetScheme gets a reference to the given string and assigns it to the Scheme field.
func (o *Card) SetScheme(v string) {
	o.Scheme = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *Card) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *Card) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *Card) SetToken(v string) {
	o.Token = &v
}

func (o Card) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Card) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BinCommercial) {
		toSerialize["bin_commercial"] = o.BinCommercial
	}
	if !IsNil(o.BinCorporate) {
		toSerialize["bin_corporate"] = o.BinCorporate
	}
	if !IsNil(o.BinCountryIssued) {
		toSerialize["bin_country_issued"] = o.BinCountryIssued
	}
	if !IsNil(o.BinCredit) {
		toSerialize["bin_credit"] = o.BinCredit
	}
	if !IsNil(o.BinCurrency) {
		toSerialize["bin_currency"] = o.BinCurrency
	}
	if !IsNil(o.BinDebit) {
		toSerialize["bin_debit"] = o.BinDebit
	}
	if !IsNil(o.BinDescription) {
		toSerialize["bin_description"] = o.BinDescription
	}
	if !IsNil(o.BinEu) {
		toSerialize["bin_eu"] = o.BinEu
	}
	if !IsNil(o.CardId) {
		toSerialize["card_id"] = o.CardId
	}
	if !IsNil(o.CardStatus) {
		toSerialize["card_status"] = o.CardStatus
	}
	if !IsNil(o.DateCreated) {
		toSerialize["date_created"] = o.DateCreated
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Expmonth) {
		toSerialize["expmonth"] = o.Expmonth
	}
	if !IsNil(o.Expyear) {
		toSerialize["expyear"] = o.Expyear
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Label2) {
		toSerialize["label2"] = o.Label2
	}
	if !IsNil(o.Last4digits) {
		toSerialize["last4digits"] = o.Last4digits
	}
	if !IsNil(o.NameOnCard) {
		toSerialize["name_on_card"] = o.NameOnCard
	}
	if !IsNil(o.Scheme) {
		toSerialize["scheme"] = o.Scheme
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableCard struct {
	value *Card
	isSet bool
}

func (v NullableCard) Get() *Card {
	return v.value
}

func (v *NullableCard) Set(val *Card) {
	v.value = val
	v.isSet = true
}

func (v NullableCard) IsSet() bool {
	return v.isSet
}

func (v *NullableCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCard(val *Card) *NullableCard {
	return &NullableCard{value: val, isSet: true}
}

func (v NullableCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
