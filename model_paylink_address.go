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

// checks if the PaylinkAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaylinkAddress{}

// PaylinkAddress struct for PaylinkAddress
type PaylinkAddress struct {
	// The first line of the address.
	Address1 *string `json:"address1,omitempty"`
	// The second line of the address.
	Address2 *string `json:"address2,omitempty"`
	// The third line of the address.
	Address3 *string `json:"address3,omitempty"`
	// The area such as city, department, town or parish.
	Area *string `json:"area,omitempty"`
	// The country code in ISO 3166 format. The country code should be an ISO-3166 2 or 3 digit country code.
	Country *string `json:"country,omitempty"`
	// A label for the address such as Head Office, Home Address.
	Label *string `json:"label,omitempty"`
	// The postcode or zip code of the address.
	Postcode *string `json:"postcode,omitempty"`
}

// NewPaylinkAddress instantiates a new PaylinkAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaylinkAddress() *PaylinkAddress {
	this := PaylinkAddress{}
	return &this
}

// NewPaylinkAddressWithDefaults instantiates a new PaylinkAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaylinkAddressWithDefaults() *PaylinkAddress {
	this := PaylinkAddress{}
	return &this
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *PaylinkAddress) GetAddress1() string {
	if o == nil || IsNil(o.Address1) {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkAddress) GetAddress1Ok() (*string, bool) {
	if o == nil || IsNil(o.Address1) {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *PaylinkAddress) HasAddress1() bool {
	if o != nil && !IsNil(o.Address1) {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *PaylinkAddress) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *PaylinkAddress) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkAddress) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *PaylinkAddress) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *PaylinkAddress) SetAddress2(v string) {
	o.Address2 = &v
}

// GetAddress3 returns the Address3 field value if set, zero value otherwise.
func (o *PaylinkAddress) GetAddress3() string {
	if o == nil || IsNil(o.Address3) {
		var ret string
		return ret
	}
	return *o.Address3
}

// GetAddress3Ok returns a tuple with the Address3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkAddress) GetAddress3Ok() (*string, bool) {
	if o == nil || IsNil(o.Address3) {
		return nil, false
	}
	return o.Address3, true
}

// HasAddress3 returns a boolean if a field has been set.
func (o *PaylinkAddress) HasAddress3() bool {
	if o != nil && !IsNil(o.Address3) {
		return true
	}

	return false
}

// SetAddress3 gets a reference to the given string and assigns it to the Address3 field.
func (o *PaylinkAddress) SetAddress3(v string) {
	o.Address3 = &v
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *PaylinkAddress) GetArea() string {
	if o == nil || IsNil(o.Area) {
		var ret string
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkAddress) GetAreaOk() (*string, bool) {
	if o == nil || IsNil(o.Area) {
		return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *PaylinkAddress) HasArea() bool {
	if o != nil && !IsNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given string and assigns it to the Area field.
func (o *PaylinkAddress) SetArea(v string) {
	o.Area = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *PaylinkAddress) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkAddress) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *PaylinkAddress) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *PaylinkAddress) SetCountry(v string) {
	o.Country = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PaylinkAddress) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkAddress) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PaylinkAddress) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PaylinkAddress) SetLabel(v string) {
	o.Label = &v
}

// GetPostcode returns the Postcode field value if set, zero value otherwise.
func (o *PaylinkAddress) GetPostcode() string {
	if o == nil || IsNil(o.Postcode) {
		var ret string
		return ret
	}
	return *o.Postcode
}

// GetPostcodeOk returns a tuple with the Postcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkAddress) GetPostcodeOk() (*string, bool) {
	if o == nil || IsNil(o.Postcode) {
		return nil, false
	}
	return o.Postcode, true
}

// HasPostcode returns a boolean if a field has been set.
func (o *PaylinkAddress) HasPostcode() bool {
	if o != nil && !IsNil(o.Postcode) {
		return true
	}

	return false
}

// SetPostcode gets a reference to the given string and assigns it to the Postcode field.
func (o *PaylinkAddress) SetPostcode(v string) {
	o.Postcode = &v
}

func (o PaylinkAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaylinkAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address1) {
		toSerialize["address1"] = o.Address1
	}
	if !IsNil(o.Address2) {
		toSerialize["address2"] = o.Address2
	}
	if !IsNil(o.Address3) {
		toSerialize["address3"] = o.Address3
	}
	if !IsNil(o.Area) {
		toSerialize["area"] = o.Area
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Postcode) {
		toSerialize["postcode"] = o.Postcode
	}
	return toSerialize, nil
}

type NullablePaylinkAddress struct {
	value *PaylinkAddress
	isSet bool
}

func (v NullablePaylinkAddress) Get() *PaylinkAddress {
	return v.value
}

func (v *NullablePaylinkAddress) Set(val *PaylinkAddress) {
	v.value = val
	v.isSet = true
}

func (v NullablePaylinkAddress) IsSet() bool {
	return v.isSet
}

func (v *NullablePaylinkAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaylinkAddress(val *PaylinkAddress) *NullablePaylinkAddress {
	return &NullablePaylinkAddress{value: val, isSet: true}
}

func (v NullablePaylinkAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaylinkAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}