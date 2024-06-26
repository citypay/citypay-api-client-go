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

// checks if the PaylinkCardHolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaylinkCardHolder{}

// PaylinkCardHolder struct for PaylinkCardHolder
type PaylinkCardHolder struct {
	// The accept headers string generated by the Customer Browser. This field may be used to lock the payment process to the customer's browser. If the customer were to attempt to use a different browser an error will be generated.
	AcceptHeaders *string         `json:"accept_headers,omitempty"`
	Address       *PaylinkAddress `json:"address,omitempty"`
	// The company name for the card holder.
	Company *string `json:"company,omitempty"`
	// The cardholder's email address. This field can be used to send a receipt to the payment cardholder. If this value is not supplied, no email will be sent.
	Email *string `json:"email,omitempty"`
	// The first name of the card holder.
	Firstname *string `json:"firstname,omitempty"`
	// The last name or surname of the card holder.
	Lastname *string `json:"lastname,omitempty"`
	// The mobile number of the cardholder. This can be used for data collection via the Paylink Payment Form or to send an SMS on completion of a transaction. This feature is a licensable option and is not configured by default.
	MobileNo *string `json:"mobile_no,omitempty"`
	// Specifies the remote IP address of the customer's browser. This field may be used to lock the payment form to the customer's IP address. Should the address change or a malicious third party attempted to hijack the transaction, an error will be generated.
	RemoteAddr *string `json:"remote_addr,omitempty"`
	// A title for the card holder such as Mr, Mrs, Ms, M. Mme. etc.
	Title *string `json:"title,omitempty"`
	// Specifies the user agent string of the Customer Browser. This field may be used to lock the payment form to the browser. Should a different user agent attempt to process the transaction or a malicious third party attempted to hijack the transaction, an error is generated.
	UserAgent *string `json:"user_agent,omitempty"`
}

// NewPaylinkCardHolder instantiates a new PaylinkCardHolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaylinkCardHolder() *PaylinkCardHolder {
	this := PaylinkCardHolder{}
	return &this
}

// NewPaylinkCardHolderWithDefaults instantiates a new PaylinkCardHolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaylinkCardHolderWithDefaults() *PaylinkCardHolder {
	this := PaylinkCardHolder{}
	return &this
}

// GetAcceptHeaders returns the AcceptHeaders field value if set, zero value otherwise.
func (o *PaylinkCardHolder) GetAcceptHeaders() string {
	if o == nil || IsNil(o.AcceptHeaders) {
		var ret string
		return ret
	}
	return *o.AcceptHeaders
}

// GetAcceptHeadersOk returns a tuple with the AcceptHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCardHolder) GetAcceptHeadersOk() (*string, bool) {
	if o == nil || IsNil(o.AcceptHeaders) {
		return nil, false
	}
	return o.AcceptHeaders, true
}

// HasAcceptHeaders returns a boolean if a field has been set.
func (o *PaylinkCardHolder) HasAcceptHeaders() bool {
	if o != nil && !IsNil(o.AcceptHeaders) {
		return true
	}

	return false
}

// SetAcceptHeaders gets a reference to the given string and assigns it to the AcceptHeaders field.
func (o *PaylinkCardHolder) SetAcceptHeaders(v string) {
	o.AcceptHeaders = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *PaylinkCardHolder) GetAddress() PaylinkAddress {
	if o == nil || IsNil(o.Address) {
		var ret PaylinkAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCardHolder) GetAddressOk() (*PaylinkAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *PaylinkCardHolder) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given PaylinkAddress and assigns it to the Address field.
func (o *PaylinkCardHolder) SetAddress(v PaylinkAddress) {
	o.Address = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *PaylinkCardHolder) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCardHolder) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *PaylinkCardHolder) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *PaylinkCardHolder) SetCompany(v string) {
	o.Company = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PaylinkCardHolder) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCardHolder) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PaylinkCardHolder) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PaylinkCardHolder) SetEmail(v string) {
	o.Email = &v
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *PaylinkCardHolder) GetFirstname() string {
	if o == nil || IsNil(o.Firstname) {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCardHolder) GetFirstnameOk() (*string, bool) {
	if o == nil || IsNil(o.Firstname) {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *PaylinkCardHolder) HasFirstname() bool {
	if o != nil && !IsNil(o.Firstname) {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *PaylinkCardHolder) SetFirstname(v string) {
	o.Firstname = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *PaylinkCardHolder) GetLastname() string {
	if o == nil || IsNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCardHolder) GetLastnameOk() (*string, bool) {
	if o == nil || IsNil(o.Lastname) {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *PaylinkCardHolder) HasLastname() bool {
	if o != nil && !IsNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *PaylinkCardHolder) SetLastname(v string) {
	o.Lastname = &v
}

// GetMobileNo returns the MobileNo field value if set, zero value otherwise.
func (o *PaylinkCardHolder) GetMobileNo() string {
	if o == nil || IsNil(o.MobileNo) {
		var ret string
		return ret
	}
	return *o.MobileNo
}

// GetMobileNoOk returns a tuple with the MobileNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCardHolder) GetMobileNoOk() (*string, bool) {
	if o == nil || IsNil(o.MobileNo) {
		return nil, false
	}
	return o.MobileNo, true
}

// HasMobileNo returns a boolean if a field has been set.
func (o *PaylinkCardHolder) HasMobileNo() bool {
	if o != nil && !IsNil(o.MobileNo) {
		return true
	}

	return false
}

// SetMobileNo gets a reference to the given string and assigns it to the MobileNo field.
func (o *PaylinkCardHolder) SetMobileNo(v string) {
	o.MobileNo = &v
}

// GetRemoteAddr returns the RemoteAddr field value if set, zero value otherwise.
func (o *PaylinkCardHolder) GetRemoteAddr() string {
	if o == nil || IsNil(o.RemoteAddr) {
		var ret string
		return ret
	}
	return *o.RemoteAddr
}

// GetRemoteAddrOk returns a tuple with the RemoteAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCardHolder) GetRemoteAddrOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteAddr) {
		return nil, false
	}
	return o.RemoteAddr, true
}

// HasRemoteAddr returns a boolean if a field has been set.
func (o *PaylinkCardHolder) HasRemoteAddr() bool {
	if o != nil && !IsNil(o.RemoteAddr) {
		return true
	}

	return false
}

// SetRemoteAddr gets a reference to the given string and assigns it to the RemoteAddr field.
func (o *PaylinkCardHolder) SetRemoteAddr(v string) {
	o.RemoteAddr = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PaylinkCardHolder) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCardHolder) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PaylinkCardHolder) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PaylinkCardHolder) SetTitle(v string) {
	o.Title = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *PaylinkCardHolder) GetUserAgent() string {
	if o == nil || IsNil(o.UserAgent) {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCardHolder) GetUserAgentOk() (*string, bool) {
	if o == nil || IsNil(o.UserAgent) {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *PaylinkCardHolder) HasUserAgent() bool {
	if o != nil && !IsNil(o.UserAgent) {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *PaylinkCardHolder) SetUserAgent(v string) {
	o.UserAgent = &v
}

func (o PaylinkCardHolder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaylinkCardHolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcceptHeaders) {
		toSerialize["accept_headers"] = o.AcceptHeaders
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Firstname) {
		toSerialize["firstname"] = o.Firstname
	}
	if !IsNil(o.Lastname) {
		toSerialize["lastname"] = o.Lastname
	}
	if !IsNil(o.MobileNo) {
		toSerialize["mobile_no"] = o.MobileNo
	}
	if !IsNil(o.RemoteAddr) {
		toSerialize["remote_addr"] = o.RemoteAddr
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.UserAgent) {
		toSerialize["user_agent"] = o.UserAgent
	}
	return toSerialize, nil
}

type NullablePaylinkCardHolder struct {
	value *PaylinkCardHolder
	isSet bool
}

func (v NullablePaylinkCardHolder) Get() *PaylinkCardHolder {
	return v.value
}

func (v *NullablePaylinkCardHolder) Set(val *PaylinkCardHolder) {
	v.value = val
	v.isSet = true
}

func (v NullablePaylinkCardHolder) IsSet() bool {
	return v.isSet
}

func (v *NullablePaylinkCardHolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaylinkCardHolder(val *PaylinkCardHolder) *NullablePaylinkCardHolder {
	return &NullablePaylinkCardHolder{value: val, isSet: true}
}

func (v NullablePaylinkCardHolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaylinkCardHolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
