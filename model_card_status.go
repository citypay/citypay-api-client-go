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

// checks if the CardStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardStatus{}

// CardStatus struct for CardStatus
type CardStatus struct {
	// The status of the card to set, valid values are ACTIVE or INACTIVE.
	CardStatus *string `json:"card_status,omitempty"`
	// Defines if the card is set as the default.
	Default *bool `json:"default,omitempty"`
}

// NewCardStatus instantiates a new CardStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardStatus() *CardStatus {
	this := CardStatus{}
	return &this
}

// NewCardStatusWithDefaults instantiates a new CardStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardStatusWithDefaults() *CardStatus {
	this := CardStatus{}
	return &this
}

// GetCardStatus returns the CardStatus field value if set, zero value otherwise.
func (o *CardStatus) GetCardStatus() string {
	if o == nil || IsNil(o.CardStatus) {
		var ret string
		return ret
	}
	return *o.CardStatus
}

// GetCardStatusOk returns a tuple with the CardStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardStatus) GetCardStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CardStatus) {
		return nil, false
	}
	return o.CardStatus, true
}

// HasCardStatus returns a boolean if a field has been set.
func (o *CardStatus) HasCardStatus() bool {
	if o != nil && !IsNil(o.CardStatus) {
		return true
	}

	return false
}

// SetCardStatus gets a reference to the given string and assigns it to the CardStatus field.
func (o *CardStatus) SetCardStatus(v string) {
	o.CardStatus = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *CardStatus) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardStatus) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *CardStatus) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *CardStatus) SetDefault(v bool) {
	o.Default = &v
}

func (o CardStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CardStatus) {
		toSerialize["card_status"] = o.CardStatus
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	return toSerialize, nil
}

type NullableCardStatus struct {
	value *CardStatus
	isSet bool
}

func (v NullableCardStatus) Get() *CardStatus {
	return v.value
}

func (v *NullableCardStatus) Set(val *CardStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCardStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCardStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardStatus(val *CardStatus) *NullableCardStatus {
	return &NullableCardStatus{value: val, isSet: true}
}

func (v NullableCardStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
