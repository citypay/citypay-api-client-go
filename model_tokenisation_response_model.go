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
)

// checks if the TokenisationResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenisationResponseModel{}

// TokenisationResponseModel struct for TokenisationResponseModel
type TokenisationResponseModel struct {
	// The result of any authentication using 3d_secure authorisation against ecommerce transactions. Values are:  <table> <tr> <th>Value</th> <th>Description</th> </tr> <tr> <td>Y</td> <td>Authentication Successful. The Cardholder's password was successfully validated.</td> </tr> <tr> <td>N</td> <td>Authentication Failed. Customer failed or cancelled authentication, transaction denied.</td> </tr> <tr> <td>A</td> <td>Attempts Processing Performed Authentication could not be completed but a proof of authentication attempt (CAVV) was generated.</td> </tr> <tr> <td>U</td> <td>Authentication Could Not Be Performed Authentication could not be completed, due to technical or other problem.</td> </tr> </table>
	AuthenResult *string `json:"authen_result,omitempty"`
	// Determines whether the bin range was found to be a commercial or business card.
	BinCommercial *bool `json:"bin_commercial,omitempty"`
	// Determines whether the bin range was found to be a debit card. If false the card was considered as a credit card.
	BinDebit *bool `json:"bin_debit,omitempty"`
	// A description of the bin range found for the card.
	BinDescription *string `json:"bin_description,omitempty"`
	// An Electronic Commerce Indicator (ECI) used to identify the result of authentication using 3DSecure.
	Eci *string `json:"eci,omitempty"`
	// The identifier provided within the request.
	Identifier *string `json:"identifier,omitempty"`
	// A masked value of the card number used for processing displaying limited values that can be used on a receipt.
	Maskedpan *string `json:"maskedpan,omitempty"`
	// The name of the card scheme of the transaction that processed the transaction such as Visa or MasterCard.
	Scheme *string `json:"scheme,omitempty"`
	// A Base58 encoded SHA-256 digest generated from the token value Base58 decoded and appended with the nonce value UTF-8 decoded.
	SigId *string `json:"sig_id,omitempty"`
	// The token used for presentment to authorisation later in the processing flow.
	Token *string `json:"token,omitempty"`
}

// NewTokenisationResponseModel instantiates a new TokenisationResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenisationResponseModel() *TokenisationResponseModel {
	this := TokenisationResponseModel{}
	return &this
}

// NewTokenisationResponseModelWithDefaults instantiates a new TokenisationResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenisationResponseModelWithDefaults() *TokenisationResponseModel {
	this := TokenisationResponseModel{}
	return &this
}

// GetAuthenResult returns the AuthenResult field value if set, zero value otherwise.
func (o *TokenisationResponseModel) GetAuthenResult() string {
	if o == nil || IsNil(o.AuthenResult) {
		var ret string
		return ret
	}
	return *o.AuthenResult
}

// GetAuthenResultOk returns a tuple with the AuthenResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenisationResponseModel) GetAuthenResultOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenResult) {
		return nil, false
	}
	return o.AuthenResult, true
}

// HasAuthenResult returns a boolean if a field has been set.
func (o *TokenisationResponseModel) HasAuthenResult() bool {
	if o != nil && !IsNil(o.AuthenResult) {
		return true
	}

	return false
}

// SetAuthenResult gets a reference to the given string and assigns it to the AuthenResult field.
func (o *TokenisationResponseModel) SetAuthenResult(v string) {
	o.AuthenResult = &v
}

// GetBinCommercial returns the BinCommercial field value if set, zero value otherwise.
func (o *TokenisationResponseModel) GetBinCommercial() bool {
	if o == nil || IsNil(o.BinCommercial) {
		var ret bool
		return ret
	}
	return *o.BinCommercial
}

// GetBinCommercialOk returns a tuple with the BinCommercial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenisationResponseModel) GetBinCommercialOk() (*bool, bool) {
	if o == nil || IsNil(o.BinCommercial) {
		return nil, false
	}
	return o.BinCommercial, true
}

// HasBinCommercial returns a boolean if a field has been set.
func (o *TokenisationResponseModel) HasBinCommercial() bool {
	if o != nil && !IsNil(o.BinCommercial) {
		return true
	}

	return false
}

// SetBinCommercial gets a reference to the given bool and assigns it to the BinCommercial field.
func (o *TokenisationResponseModel) SetBinCommercial(v bool) {
	o.BinCommercial = &v
}

// GetBinDebit returns the BinDebit field value if set, zero value otherwise.
func (o *TokenisationResponseModel) GetBinDebit() bool {
	if o == nil || IsNil(o.BinDebit) {
		var ret bool
		return ret
	}
	return *o.BinDebit
}

// GetBinDebitOk returns a tuple with the BinDebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenisationResponseModel) GetBinDebitOk() (*bool, bool) {
	if o == nil || IsNil(o.BinDebit) {
		return nil, false
	}
	return o.BinDebit, true
}

// HasBinDebit returns a boolean if a field has been set.
func (o *TokenisationResponseModel) HasBinDebit() bool {
	if o != nil && !IsNil(o.BinDebit) {
		return true
	}

	return false
}

// SetBinDebit gets a reference to the given bool and assigns it to the BinDebit field.
func (o *TokenisationResponseModel) SetBinDebit(v bool) {
	o.BinDebit = &v
}

// GetBinDescription returns the BinDescription field value if set, zero value otherwise.
func (o *TokenisationResponseModel) GetBinDescription() string {
	if o == nil || IsNil(o.BinDescription) {
		var ret string
		return ret
	}
	return *o.BinDescription
}

// GetBinDescriptionOk returns a tuple with the BinDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenisationResponseModel) GetBinDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.BinDescription) {
		return nil, false
	}
	return o.BinDescription, true
}

// HasBinDescription returns a boolean if a field has been set.
func (o *TokenisationResponseModel) HasBinDescription() bool {
	if o != nil && !IsNil(o.BinDescription) {
		return true
	}

	return false
}

// SetBinDescription gets a reference to the given string and assigns it to the BinDescription field.
func (o *TokenisationResponseModel) SetBinDescription(v string) {
	o.BinDescription = &v
}

// GetEci returns the Eci field value if set, zero value otherwise.
func (o *TokenisationResponseModel) GetEci() string {
	if o == nil || IsNil(o.Eci) {
		var ret string
		return ret
	}
	return *o.Eci
}

// GetEciOk returns a tuple with the Eci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenisationResponseModel) GetEciOk() (*string, bool) {
	if o == nil || IsNil(o.Eci) {
		return nil, false
	}
	return o.Eci, true
}

// HasEci returns a boolean if a field has been set.
func (o *TokenisationResponseModel) HasEci() bool {
	if o != nil && !IsNil(o.Eci) {
		return true
	}

	return false
}

// SetEci gets a reference to the given string and assigns it to the Eci field.
func (o *TokenisationResponseModel) SetEci(v string) {
	o.Eci = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *TokenisationResponseModel) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenisationResponseModel) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *TokenisationResponseModel) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *TokenisationResponseModel) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetMaskedpan returns the Maskedpan field value if set, zero value otherwise.
func (o *TokenisationResponseModel) GetMaskedpan() string {
	if o == nil || IsNil(o.Maskedpan) {
		var ret string
		return ret
	}
	return *o.Maskedpan
}

// GetMaskedpanOk returns a tuple with the Maskedpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenisationResponseModel) GetMaskedpanOk() (*string, bool) {
	if o == nil || IsNil(o.Maskedpan) {
		return nil, false
	}
	return o.Maskedpan, true
}

// HasMaskedpan returns a boolean if a field has been set.
func (o *TokenisationResponseModel) HasMaskedpan() bool {
	if o != nil && !IsNil(o.Maskedpan) {
		return true
	}

	return false
}

// SetMaskedpan gets a reference to the given string and assigns it to the Maskedpan field.
func (o *TokenisationResponseModel) SetMaskedpan(v string) {
	o.Maskedpan = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *TokenisationResponseModel) GetScheme() string {
	if o == nil || IsNil(o.Scheme) {
		var ret string
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenisationResponseModel) GetSchemeOk() (*string, bool) {
	if o == nil || IsNil(o.Scheme) {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *TokenisationResponseModel) HasScheme() bool {
	if o != nil && !IsNil(o.Scheme) {
		return true
	}

	return false
}

// SetScheme gets a reference to the given string and assigns it to the Scheme field.
func (o *TokenisationResponseModel) SetScheme(v string) {
	o.Scheme = &v
}

// GetSigId returns the SigId field value if set, zero value otherwise.
func (o *TokenisationResponseModel) GetSigId() string {
	if o == nil || IsNil(o.SigId) {
		var ret string
		return ret
	}
	return *o.SigId
}

// GetSigIdOk returns a tuple with the SigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenisationResponseModel) GetSigIdOk() (*string, bool) {
	if o == nil || IsNil(o.SigId) {
		return nil, false
	}
	return o.SigId, true
}

// HasSigId returns a boolean if a field has been set.
func (o *TokenisationResponseModel) HasSigId() bool {
	if o != nil && !IsNil(o.SigId) {
		return true
	}

	return false
}

// SetSigId gets a reference to the given string and assigns it to the SigId field.
func (o *TokenisationResponseModel) SetSigId(v string) {
	o.SigId = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *TokenisationResponseModel) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenisationResponseModel) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *TokenisationResponseModel) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *TokenisationResponseModel) SetToken(v string) {
	o.Token = &v
}

func (o TokenisationResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenisationResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthenResult) {
		toSerialize["authen_result"] = o.AuthenResult
	}
	if !IsNil(o.BinCommercial) {
		toSerialize["bin_commercial"] = o.BinCommercial
	}
	if !IsNil(o.BinDebit) {
		toSerialize["bin_debit"] = o.BinDebit
	}
	if !IsNil(o.BinDescription) {
		toSerialize["bin_description"] = o.BinDescription
	}
	if !IsNil(o.Eci) {
		toSerialize["eci"] = o.Eci
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Maskedpan) {
		toSerialize["maskedpan"] = o.Maskedpan
	}
	if !IsNil(o.Scheme) {
		toSerialize["scheme"] = o.Scheme
	}
	if !IsNil(o.SigId) {
		toSerialize["sig_id"] = o.SigId
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableTokenisationResponseModel struct {
	value *TokenisationResponseModel
	isSet bool
}

func (v NullableTokenisationResponseModel) Get() *TokenisationResponseModel {
	return v.value
}

func (v *NullableTokenisationResponseModel) Set(val *TokenisationResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenisationResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenisationResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenisationResponseModel(val *TokenisationResponseModel) *NullableTokenisationResponseModel {
	return &NullableTokenisationResponseModel{value: val, isSet: true}
}

func (v NullableTokenisationResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenisationResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}