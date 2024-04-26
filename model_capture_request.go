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

// checks if the CaptureRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaptureRequest{}

// CaptureRequest struct for CaptureRequest
type CaptureRequest struct {
	AirlineData *AirlineAdvice `json:"airline_data,omitempty"`
	// The completion amount provided in the lowest unit of currency for the specific currency of the merchant, with a variable length to a maximum of 12 digits. No decimal points to be included. For example with GBP 75.45 use the value 7545. Please check that you do not supply divisional characters such as 1,024 in the request which may be caused by some number formatters.  If no amount is supplied, the original processing amount is used.
	Amount          *int32          `json:"amount,omitempty"`
	EventManagement *EventDataModel `json:"event_management,omitempty"`
	// The identifier of the transaction to capture. If an empty value is supplied then a `trans_no` value must be supplied.
	Identifier *string `json:"identifier,omitempty"`
	// Identifies the merchant account to perform the capture for.
	Merchantid int32 `json:"merchantid"`
	// The transaction number of the transaction to look up and capture. If an empty value is supplied then an identifier value must be supplied.
	Transno *int32 `json:"transno,omitempty"`
}

type _CaptureRequest CaptureRequest

// NewCaptureRequest instantiates a new CaptureRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaptureRequest(merchantid int32) *CaptureRequest {
	this := CaptureRequest{}
	this.Merchantid = merchantid
	return &this
}

// NewCaptureRequestWithDefaults instantiates a new CaptureRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaptureRequestWithDefaults() *CaptureRequest {
	this := CaptureRequest{}
	return &this
}

// GetAirlineData returns the AirlineData field value if set, zero value otherwise.
func (o *CaptureRequest) GetAirlineData() AirlineAdvice {
	if o == nil || IsNil(o.AirlineData) {
		var ret AirlineAdvice
		return ret
	}
	return *o.AirlineData
}

// GetAirlineDataOk returns a tuple with the AirlineData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureRequest) GetAirlineDataOk() (*AirlineAdvice, bool) {
	if o == nil || IsNil(o.AirlineData) {
		return nil, false
	}
	return o.AirlineData, true
}

// HasAirlineData returns a boolean if a field has been set.
func (o *CaptureRequest) HasAirlineData() bool {
	if o != nil && !IsNil(o.AirlineData) {
		return true
	}

	return false
}

// SetAirlineData gets a reference to the given AirlineAdvice and assigns it to the AirlineData field.
func (o *CaptureRequest) SetAirlineData(v AirlineAdvice) {
	o.AirlineData = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CaptureRequest) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureRequest) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CaptureRequest) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *CaptureRequest) SetAmount(v int32) {
	o.Amount = &v
}

// GetEventManagement returns the EventManagement field value if set, zero value otherwise.
func (o *CaptureRequest) GetEventManagement() EventDataModel {
	if o == nil || IsNil(o.EventManagement) {
		var ret EventDataModel
		return ret
	}
	return *o.EventManagement
}

// GetEventManagementOk returns a tuple with the EventManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureRequest) GetEventManagementOk() (*EventDataModel, bool) {
	if o == nil || IsNil(o.EventManagement) {
		return nil, false
	}
	return o.EventManagement, true
}

// HasEventManagement returns a boolean if a field has been set.
func (o *CaptureRequest) HasEventManagement() bool {
	if o != nil && !IsNil(o.EventManagement) {
		return true
	}

	return false
}

// SetEventManagement gets a reference to the given EventDataModel and assigns it to the EventManagement field.
func (o *CaptureRequest) SetEventManagement(v EventDataModel) {
	o.EventManagement = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *CaptureRequest) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureRequest) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *CaptureRequest) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *CaptureRequest) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetMerchantid returns the Merchantid field value
func (o *CaptureRequest) GetMerchantid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Merchantid
}

// GetMerchantidOk returns a tuple with the Merchantid field value
// and a boolean to check if the value has been set.
func (o *CaptureRequest) GetMerchantidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Merchantid, true
}

// SetMerchantid sets field value
func (o *CaptureRequest) SetMerchantid(v int32) {
	o.Merchantid = v
}

// GetTransno returns the Transno field value if set, zero value otherwise.
func (o *CaptureRequest) GetTransno() int32 {
	if o == nil || IsNil(o.Transno) {
		var ret int32
		return ret
	}
	return *o.Transno
}

// GetTransnoOk returns a tuple with the Transno field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaptureRequest) GetTransnoOk() (*int32, bool) {
	if o == nil || IsNil(o.Transno) {
		return nil, false
	}
	return o.Transno, true
}

// HasTransno returns a boolean if a field has been set.
func (o *CaptureRequest) HasTransno() bool {
	if o != nil && !IsNil(o.Transno) {
		return true
	}

	return false
}

// SetTransno gets a reference to the given int32 and assigns it to the Transno field.
func (o *CaptureRequest) SetTransno(v int32) {
	o.Transno = &v
}

func (o CaptureRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaptureRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AirlineData) {
		toSerialize["airline_data"] = o.AirlineData
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.EventManagement) {
		toSerialize["event_management"] = o.EventManagement
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	toSerialize["merchantid"] = o.Merchantid
	if !IsNil(o.Transno) {
		toSerialize["transno"] = o.Transno
	}
	return toSerialize, nil
}

func (o *CaptureRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"merchantid",
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

	varCaptureRequest := _CaptureRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCaptureRequest)

	if err != nil {
		return err
	}

	*o = CaptureRequest(varCaptureRequest)

	return err
}

type NullableCaptureRequest struct {
	value *CaptureRequest
	isSet bool
}

func (v NullableCaptureRequest) Get() *CaptureRequest {
	return v.value
}

func (v *NullableCaptureRequest) Set(val *CaptureRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCaptureRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCaptureRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaptureRequest(val *CaptureRequest) *NullableCaptureRequest {
	return &NullableCaptureRequest{value: val, isSet: true}
}

func (v NullableCaptureRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaptureRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}