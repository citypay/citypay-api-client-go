/*
CityPay Payment API

 Welcome to the CityPay API, a robust HTTP API payment solution designed for seamless server-to-server  transactional processing. Our API facilitates a wide array of payment operations, catering to diverse business needs.  Whether you're integrating Internet payments, handling Mail Order/Telephone Order (MOTO) transactions, managing  Subscriptions with Recurring and Continuous Authority payments, or navigating the complexities of 3-D Secure  authentication, our API is equipped to support your requirements. Additionally, we offer functionalities for  Authorisation, Refunding, Pre-Authorisation, Cancellation/Voids, and Completion processing, alongside the capability  for tokenised payments.  ## Compliance and Security Overview <aside class=\"notice\">   Ensuring the security of payment transactions and compliance with industry standards is paramount. Our API is    designed with stringent security measures and compliance protocols to safeguard sensitive information and meet    the rigorous requirements of Visa, MasterCard, and the PCI Security Standards Council. </aside>  ### Key Compliance and Security Measures  * **TLS Encryption**: All data transmissions must utilise TLS version 1.2 or higher, employing [strong cryptography](#enabled-tls-ciphers). Our infrastructure strictly enforces this requirement to maintain the integrity and confidentiality of data in transit. We conduct regular scans and assessments of our TLS endpoints to identify and mitigate vulnerabilities. * **Data Storage Prohibitions**: Storing sensitive cardholder data (CHD), such as the card security code (CSC) or primary account number (PAN), is strictly prohibited. Our API is designed to minimize your exposure to sensitive data, thereby reducing your compliance burden. * **Data Masking**: For consumer protection and compliance, full card numbers must not be displayed on receipts or any customer-facing materials. Our API automatically masks PANs, displaying only the last four digits to facilitate safe receipt generation. * **Network Scans**: If your application is web-based, regular scans of your hosting environment are mandatory to identify and rectify potential vulnerabilities. This proactive measure is crucial for maintaining a secure and compliant online presence. * **PCI Compliance**: Adherence to PCI DSS standards is not optional; it's a requirement for operating securely and legally in the payments ecosystem. For detailed information on compliance requirements and resources, please visit the PCI Security Standards Council website [https://www.pcisecuritystandards.org/](https://www.pcisecuritystandards.org/). * **Request Validation**: Our API includes mechanisms to verify the legitimacy of each request, ensuring it pertains to a valid account and originates from a trusted source. We leverage remote IP address verification alongside sophisticated application firewall technologies to thwart a wide array of common security threats.  ## Getting Started Before integrating with the CityPay API, ensure your application and development practices align with the outlined compliance and security measures. This preparatory step is crucial for a smooth integration process and the long-term success of your payment processing operations.  For further details on API endpoints, request/response formats, and code examples, proceed to the subsequent sections of our documentation. Our aim is to provide you with all the necessary tools and information to integrate our payment processing capabilities seamlessly into your application.  Thank you for choosing CityPay API. We look forward to supporting your payment processing needs with our secure, compliant, and versatile API solution.

API version: 6.6.40
Contact: support@citypay.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package citypay

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PaymentIntentReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentIntentReference{}

// PaymentIntentReference struct for PaymentIntentReference
type PaymentIntentReference struct {
	// The intent id used for future referencing of the intent.
	PaymentIntentId string `json:"payment_intent_id"`
}

type _PaymentIntentReference PaymentIntentReference

// NewPaymentIntentReference instantiates a new PaymentIntentReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentIntentReference(paymentIntentId string) *PaymentIntentReference {
	this := PaymentIntentReference{}
	this.PaymentIntentId = paymentIntentId
	return &this
}

// NewPaymentIntentReferenceWithDefaults instantiates a new PaymentIntentReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentIntentReferenceWithDefaults() *PaymentIntentReference {
	this := PaymentIntentReference{}
	return &this
}

// GetPaymentIntentId returns the PaymentIntentId field value
func (o *PaymentIntentReference) GetPaymentIntentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentIntentId
}

// GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field value
// and a boolean to check if the value has been set.
func (o *PaymentIntentReference) GetPaymentIntentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentIntentId, true
}

// SetPaymentIntentId sets field value
func (o *PaymentIntentReference) SetPaymentIntentId(v string) {
	o.PaymentIntentId = v
}

func (o PaymentIntentReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentIntentReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["payment_intent_id"] = o.PaymentIntentId
	return toSerialize, nil
}

func (o *PaymentIntentReference) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"payment_intent_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPaymentIntentReference := _PaymentIntentReference{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaymentIntentReference)

	if err != nil {
		return err
	}

	*o = PaymentIntentReference(varPaymentIntentReference)

	return err
}

type NullablePaymentIntentReference struct {
	value *PaymentIntentReference
	isSet bool
}

func (v NullablePaymentIntentReference) Get() *PaymentIntentReference {
	return v.value
}

func (v *NullablePaymentIntentReference) Set(val *PaymentIntentReference) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentIntentReference) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentIntentReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentIntentReference(val *PaymentIntentReference) *NullablePaymentIntentReference {
	return &NullablePaymentIntentReference{value: val, isSet: true}
}

func (v NullablePaymentIntentReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentIntentReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
