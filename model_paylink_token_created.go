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
	"time"
)

// checks if the PaylinkTokenCreated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaylinkTokenCreated{}

// PaylinkTokenCreated struct for PaylinkTokenCreated
type PaylinkTokenCreated struct {
	Attachments *PaylinkAttachmentResult `json:"attachments,omitempty"`
	// true if BPS has been enabled on this token.
	Bps *string `json:"bps,omitempty"`
	// Date and time the token was generated.
	DateCreated *time.Time         `json:"date_created,omitempty"`
	Errors      []PaylinkErrorCode `json:"errors,omitempty"`
	// A unique id of the request.
	Id string `json:"id"`
	// The identifier as presented in the TokenRequest.
	Identifier *string `json:"identifier,omitempty"`
	// Determines whether the token is `live` or `test`.
	Mode *string `json:"mode,omitempty"`
	// A URL of a qrcode which can be used to refer to the token URL.
	Qrcode *string `json:"qrcode,omitempty"`
	// The result field contains the result for the Paylink Token Request. 0 - indicates that an error was encountered while creating the token. 1 - which indicates that a Token was successfully created.
	Result int32 `json:"result"`
	// the version of the server performing the call.
	ServerVersion *string `json:"server_version,omitempty"`
	// The incoming IP address of the call.
	Source *string `json:"source,omitempty"`
	// A token generated for the request used to refer to the transaction in consequential calls.
	Token string `json:"token"`
	// The Paylink token URL used to checkout by the card holder.
	Url *string `json:"url,omitempty"`
	// A UrlShortCode (USC) used for short links.
	Usc *string `json:"usc,omitempty"`
}

type _PaylinkTokenCreated PaylinkTokenCreated

// NewPaylinkTokenCreated instantiates a new PaylinkTokenCreated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaylinkTokenCreated(id string, result int32, token string) *PaylinkTokenCreated {
	this := PaylinkTokenCreated{}
	this.Id = id
	this.Result = result
	this.Token = token
	return &this
}

// NewPaylinkTokenCreatedWithDefaults instantiates a new PaylinkTokenCreated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaylinkTokenCreatedWithDefaults() *PaylinkTokenCreated {
	this := PaylinkTokenCreated{}
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *PaylinkTokenCreated) GetAttachments() PaylinkAttachmentResult {
	if o == nil || IsNil(o.Attachments) {
		var ret PaylinkAttachmentResult
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkTokenCreated) GetAttachmentsOk() (*PaylinkAttachmentResult, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *PaylinkTokenCreated) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given PaylinkAttachmentResult and assigns it to the Attachments field.
func (o *PaylinkTokenCreated) SetAttachments(v PaylinkAttachmentResult) {
	o.Attachments = &v
}

// GetBps returns the Bps field value if set, zero value otherwise.
func (o *PaylinkTokenCreated) GetBps() string {
	if o == nil || IsNil(o.Bps) {
		var ret string
		return ret
	}
	return *o.Bps
}

// GetBpsOk returns a tuple with the Bps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkTokenCreated) GetBpsOk() (*string, bool) {
	if o == nil || IsNil(o.Bps) {
		return nil, false
	}
	return o.Bps, true
}

// HasBps returns a boolean if a field has been set.
func (o *PaylinkTokenCreated) HasBps() bool {
	if o != nil && !IsNil(o.Bps) {
		return true
	}

	return false
}

// SetBps gets a reference to the given string and assigns it to the Bps field.
func (o *PaylinkTokenCreated) SetBps(v string) {
	o.Bps = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *PaylinkTokenCreated) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkTokenCreated) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *PaylinkTokenCreated) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *PaylinkTokenCreated) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *PaylinkTokenCreated) GetErrors() []PaylinkErrorCode {
	if o == nil || IsNil(o.Errors) {
		var ret []PaylinkErrorCode
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkTokenCreated) GetErrorsOk() ([]PaylinkErrorCode, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *PaylinkTokenCreated) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []PaylinkErrorCode and assigns it to the Errors field.
func (o *PaylinkTokenCreated) SetErrors(v []PaylinkErrorCode) {
	o.Errors = v
}

// GetId returns the Id field value
func (o *PaylinkTokenCreated) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaylinkTokenCreated) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaylinkTokenCreated) SetId(v string) {
	o.Id = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *PaylinkTokenCreated) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkTokenCreated) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *PaylinkTokenCreated) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *PaylinkTokenCreated) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *PaylinkTokenCreated) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkTokenCreated) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *PaylinkTokenCreated) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *PaylinkTokenCreated) SetMode(v string) {
	o.Mode = &v
}

// GetQrcode returns the Qrcode field value if set, zero value otherwise.
func (o *PaylinkTokenCreated) GetQrcode() string {
	if o == nil || IsNil(o.Qrcode) {
		var ret string
		return ret
	}
	return *o.Qrcode
}

// GetQrcodeOk returns a tuple with the Qrcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkTokenCreated) GetQrcodeOk() (*string, bool) {
	if o == nil || IsNil(o.Qrcode) {
		return nil, false
	}
	return o.Qrcode, true
}

// HasQrcode returns a boolean if a field has been set.
func (o *PaylinkTokenCreated) HasQrcode() bool {
	if o != nil && !IsNil(o.Qrcode) {
		return true
	}

	return false
}

// SetQrcode gets a reference to the given string and assigns it to the Qrcode field.
func (o *PaylinkTokenCreated) SetQrcode(v string) {
	o.Qrcode = &v
}

// GetResult returns the Result field value
func (o *PaylinkTokenCreated) GetResult() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *PaylinkTokenCreated) GetResultOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *PaylinkTokenCreated) SetResult(v int32) {
	o.Result = v
}

// GetServerVersion returns the ServerVersion field value if set, zero value otherwise.
func (o *PaylinkTokenCreated) GetServerVersion() string {
	if o == nil || IsNil(o.ServerVersion) {
		var ret string
		return ret
	}
	return *o.ServerVersion
}

// GetServerVersionOk returns a tuple with the ServerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkTokenCreated) GetServerVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ServerVersion) {
		return nil, false
	}
	return o.ServerVersion, true
}

// HasServerVersion returns a boolean if a field has been set.
func (o *PaylinkTokenCreated) HasServerVersion() bool {
	if o != nil && !IsNil(o.ServerVersion) {
		return true
	}

	return false
}

// SetServerVersion gets a reference to the given string and assigns it to the ServerVersion field.
func (o *PaylinkTokenCreated) SetServerVersion(v string) {
	o.ServerVersion = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *PaylinkTokenCreated) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkTokenCreated) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *PaylinkTokenCreated) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *PaylinkTokenCreated) SetSource(v string) {
	o.Source = &v
}

// GetToken returns the Token field value
func (o *PaylinkTokenCreated) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *PaylinkTokenCreated) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *PaylinkTokenCreated) SetToken(v string) {
	o.Token = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PaylinkTokenCreated) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkTokenCreated) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PaylinkTokenCreated) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PaylinkTokenCreated) SetUrl(v string) {
	o.Url = &v
}

// GetUsc returns the Usc field value if set, zero value otherwise.
func (o *PaylinkTokenCreated) GetUsc() string {
	if o == nil || IsNil(o.Usc) {
		var ret string
		return ret
	}
	return *o.Usc
}

// GetUscOk returns a tuple with the Usc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkTokenCreated) GetUscOk() (*string, bool) {
	if o == nil || IsNil(o.Usc) {
		return nil, false
	}
	return o.Usc, true
}

// HasUsc returns a boolean if a field has been set.
func (o *PaylinkTokenCreated) HasUsc() bool {
	if o != nil && !IsNil(o.Usc) {
		return true
	}

	return false
}

// SetUsc gets a reference to the given string and assigns it to the Usc field.
func (o *PaylinkTokenCreated) SetUsc(v string) {
	o.Usc = &v
}

func (o PaylinkTokenCreated) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaylinkTokenCreated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Bps) {
		toSerialize["bps"] = o.Bps
	}
	if !IsNil(o.DateCreated) {
		toSerialize["date_created"] = o.DateCreated
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Qrcode) {
		toSerialize["qrcode"] = o.Qrcode
	}
	toSerialize["result"] = o.Result
	if !IsNil(o.ServerVersion) {
		toSerialize["server_version"] = o.ServerVersion
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	toSerialize["token"] = o.Token
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Usc) {
		toSerialize["usc"] = o.Usc
	}
	return toSerialize, nil
}

func (o *PaylinkTokenCreated) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"result",
		"token",
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

	varPaylinkTokenCreated := _PaylinkTokenCreated{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaylinkTokenCreated)

	if err != nil {
		return err
	}

	*o = PaylinkTokenCreated(varPaylinkTokenCreated)

	return err
}

type NullablePaylinkTokenCreated struct {
	value *PaylinkTokenCreated
	isSet bool
}

func (v NullablePaylinkTokenCreated) Get() *PaylinkTokenCreated {
	return v.value
}

func (v *NullablePaylinkTokenCreated) Set(val *PaylinkTokenCreated) {
	v.value = val
	v.isSet = true
}

func (v NullablePaylinkTokenCreated) IsSet() bool {
	return v.isSet
}

func (v *NullablePaylinkTokenCreated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaylinkTokenCreated(val *PaylinkTokenCreated) *NullablePaylinkTokenCreated {
	return &NullablePaylinkTokenCreated{value: val, isSet: true}
}

func (v NullablePaylinkTokenCreated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaylinkTokenCreated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
