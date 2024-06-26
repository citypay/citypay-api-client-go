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

// checks if the AuthReferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthReferences{}

// AuthReferences struct for AuthReferences
type AuthReferences struct {
	Auths []AuthReference `json:"auths,omitempty"`
}

// NewAuthReferences instantiates a new AuthReferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthReferences() *AuthReferences {
	this := AuthReferences{}
	return &this
}

// NewAuthReferencesWithDefaults instantiates a new AuthReferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthReferencesWithDefaults() *AuthReferences {
	this := AuthReferences{}
	return &this
}

// GetAuths returns the Auths field value if set, zero value otherwise.
func (o *AuthReferences) GetAuths() []AuthReference {
	if o == nil || IsNil(o.Auths) {
		var ret []AuthReference
		return ret
	}
	return o.Auths
}

// GetAuthsOk returns a tuple with the Auths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthReferences) GetAuthsOk() ([]AuthReference, bool) {
	if o == nil || IsNil(o.Auths) {
		return nil, false
	}
	return o.Auths, true
}

// HasAuths returns a boolean if a field has been set.
func (o *AuthReferences) HasAuths() bool {
	if o != nil && !IsNil(o.Auths) {
		return true
	}

	return false
}

// SetAuths gets a reference to the given []AuthReference and assigns it to the Auths field.
func (o *AuthReferences) SetAuths(v []AuthReference) {
	o.Auths = v
}

func (o AuthReferences) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthReferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Auths) {
		toSerialize["auths"] = o.Auths
	}
	return toSerialize, nil
}

type NullableAuthReferences struct {
	value *AuthReferences
	isSet bool
}

func (v NullableAuthReferences) Get() *AuthReferences {
	return v.value
}

func (v *NullableAuthReferences) Set(val *AuthReferences) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthReferences) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthReferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthReferences(val *AuthReferences) *NullableAuthReferences {
	return &NullableAuthReferences{value: val, isSet: true}
}

func (v NullableAuthReferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthReferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
