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

// checks if the AuthReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthReference{}

// AuthReference struct for AuthReference
type AuthReference struct {
	// The amount of the transaction in decimal currency format.
	Amount *string `json:"amount,omitempty"`
	// The amount of the transaction in integer/request format.
	AmountValue *int32 `json:"amount_value,omitempty"`
	// A reference number provided by the acquiring services.
	Atrn *string `json:"atrn,omitempty"`
	// The authorisation code of the transaction returned by the acquirer or card issuer.
	Authcode *string `json:"authcode,omitempty"`
	// A batch number which the transaction has been end of day batched towards.
	Batchno *string `json:"batchno,omitempty"`
	// The currency of the transaction in ISO 4217 code format.
	Currency *string `json:"currency,omitempty"`
	// The date and time of the transaction.
	Datetime *time.Time `json:"datetime,omitempty"`
	// The identifier of the transaction used to process the transaction.
	Identifier *string `json:"identifier,omitempty"`
	// A masking of the card number which was used to process the tranasction.
	Maskedpan *string `json:"maskedpan,omitempty"`
	// The merchant id of the transaction result.
	Merchantid *int32 `json:"merchantid,omitempty"`
	// The result of the transaction.
	Result *string `json:"result,omitempty"`
	// The current status of the transaction through it's lifecycle.
	TransStatus *string `json:"trans_status,omitempty"`
	// The type of transaction that was processed.
	TransType *string `json:"trans_type,omitempty"`
	// The transaction number of the transaction.
	Transno *int32 `json:"transno,omitempty"`
}

// NewAuthReference instantiates a new AuthReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthReference() *AuthReference {
	this := AuthReference{}
	return &this
}

// NewAuthReferenceWithDefaults instantiates a new AuthReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthReferenceWithDefaults() *AuthReference {
	this := AuthReference{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AuthReference) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthReference) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AuthReference) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *AuthReference) SetAmount(v string) {
	o.Amount = &v
}

// GetAmountValue returns the AmountValue field value if set, zero value otherwise.
func (o *AuthReference) GetAmountValue() int32 {
	if o == nil || IsNil(o.AmountValue) {
		var ret int32
		return ret
	}
	return *o.AmountValue
}

// GetAmountValueOk returns a tuple with the AmountValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthReference) GetAmountValueOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountValue) {
		return nil, false
	}
	return o.AmountValue, true
}

// HasAmountValue returns a boolean if a field has been set.
func (o *AuthReference) HasAmountValue() bool {
	if o != nil && !IsNil(o.AmountValue) {
		return true
	}

	return false
}

// SetAmountValue gets a reference to the given int32 and assigns it to the AmountValue field.
func (o *AuthReference) SetAmountValue(v int32) {
	o.AmountValue = &v
}

// GetAtrn returns the Atrn field value if set, zero value otherwise.
func (o *AuthReference) GetAtrn() string {
	if o == nil || IsNil(o.Atrn) {
		var ret string
		return ret
	}
	return *o.Atrn
}

// GetAtrnOk returns a tuple with the Atrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthReference) GetAtrnOk() (*string, bool) {
	if o == nil || IsNil(o.Atrn) {
		return nil, false
	}
	return o.Atrn, true
}

// HasAtrn returns a boolean if a field has been set.
func (o *AuthReference) HasAtrn() bool {
	if o != nil && !IsNil(o.Atrn) {
		return true
	}

	return false
}

// SetAtrn gets a reference to the given string and assigns it to the Atrn field.
func (o *AuthReference) SetAtrn(v string) {
	o.Atrn = &v
}

// GetAuthcode returns the Authcode field value if set, zero value otherwise.
func (o *AuthReference) GetAuthcode() string {
	if o == nil || IsNil(o.Authcode) {
		var ret string
		return ret
	}
	return *o.Authcode
}

// GetAuthcodeOk returns a tuple with the Authcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthReference) GetAuthcodeOk() (*string, bool) {
	if o == nil || IsNil(o.Authcode) {
		return nil, false
	}
	return o.Authcode, true
}

// HasAuthcode returns a boolean if a field has been set.
func (o *AuthReference) HasAuthcode() bool {
	if o != nil && !IsNil(o.Authcode) {
		return true
	}

	return false
}

// SetAuthcode gets a reference to the given string and assigns it to the Authcode field.
func (o *AuthReference) SetAuthcode(v string) {
	o.Authcode = &v
}

// GetBatchno returns the Batchno field value if set, zero value otherwise.
func (o *AuthReference) GetBatchno() string {
	if o == nil || IsNil(o.Batchno) {
		var ret string
		return ret
	}
	return *o.Batchno
}

// GetBatchnoOk returns a tuple with the Batchno field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthReference) GetBatchnoOk() (*string, bool) {
	if o == nil || IsNil(o.Batchno) {
		return nil, false
	}
	return o.Batchno, true
}

// HasBatchno returns a boolean if a field has been set.
func (o *AuthReference) HasBatchno() bool {
	if o != nil && !IsNil(o.Batchno) {
		return true
	}

	return false
}

// SetBatchno gets a reference to the given string and assigns it to the Batchno field.
func (o *AuthReference) SetBatchno(v string) {
	o.Batchno = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AuthReference) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthReference) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AuthReference) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AuthReference) SetCurrency(v string) {
	o.Currency = &v
}

// GetDatetime returns the Datetime field value if set, zero value otherwise.
func (o *AuthReference) GetDatetime() time.Time {
	if o == nil || IsNil(o.Datetime) {
		var ret time.Time
		return ret
	}
	return *o.Datetime
}

// GetDatetimeOk returns a tuple with the Datetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthReference) GetDatetimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Datetime) {
		return nil, false
	}
	return o.Datetime, true
}

// HasDatetime returns a boolean if a field has been set.
func (o *AuthReference) HasDatetime() bool {
	if o != nil && !IsNil(o.Datetime) {
		return true
	}

	return false
}

// SetDatetime gets a reference to the given time.Time and assigns it to the Datetime field.
func (o *AuthReference) SetDatetime(v time.Time) {
	o.Datetime = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *AuthReference) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthReference) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *AuthReference) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *AuthReference) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetMaskedpan returns the Maskedpan field value if set, zero value otherwise.
func (o *AuthReference) GetMaskedpan() string {
	if o == nil || IsNil(o.Maskedpan) {
		var ret string
		return ret
	}
	return *o.Maskedpan
}

// GetMaskedpanOk returns a tuple with the Maskedpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthReference) GetMaskedpanOk() (*string, bool) {
	if o == nil || IsNil(o.Maskedpan) {
		return nil, false
	}
	return o.Maskedpan, true
}

// HasMaskedpan returns a boolean if a field has been set.
func (o *AuthReference) HasMaskedpan() bool {
	if o != nil && !IsNil(o.Maskedpan) {
		return true
	}

	return false
}

// SetMaskedpan gets a reference to the given string and assigns it to the Maskedpan field.
func (o *AuthReference) SetMaskedpan(v string) {
	o.Maskedpan = &v
}

// GetMerchantid returns the Merchantid field value if set, zero value otherwise.
func (o *AuthReference) GetMerchantid() int32 {
	if o == nil || IsNil(o.Merchantid) {
		var ret int32
		return ret
	}
	return *o.Merchantid
}

// GetMerchantidOk returns a tuple with the Merchantid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthReference) GetMerchantidOk() (*int32, bool) {
	if o == nil || IsNil(o.Merchantid) {
		return nil, false
	}
	return o.Merchantid, true
}

// HasMerchantid returns a boolean if a field has been set.
func (o *AuthReference) HasMerchantid() bool {
	if o != nil && !IsNil(o.Merchantid) {
		return true
	}

	return false
}

// SetMerchantid gets a reference to the given int32 and assigns it to the Merchantid field.
func (o *AuthReference) SetMerchantid(v int32) {
	o.Merchantid = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *AuthReference) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthReference) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *AuthReference) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *AuthReference) SetResult(v string) {
	o.Result = &v
}

// GetTransStatus returns the TransStatus field value if set, zero value otherwise.
func (o *AuthReference) GetTransStatus() string {
	if o == nil || IsNil(o.TransStatus) {
		var ret string
		return ret
	}
	return *o.TransStatus
}

// GetTransStatusOk returns a tuple with the TransStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthReference) GetTransStatusOk() (*string, bool) {
	if o == nil || IsNil(o.TransStatus) {
		return nil, false
	}
	return o.TransStatus, true
}

// HasTransStatus returns a boolean if a field has been set.
func (o *AuthReference) HasTransStatus() bool {
	if o != nil && !IsNil(o.TransStatus) {
		return true
	}

	return false
}

// SetTransStatus gets a reference to the given string and assigns it to the TransStatus field.
func (o *AuthReference) SetTransStatus(v string) {
	o.TransStatus = &v
}

// GetTransType returns the TransType field value if set, zero value otherwise.
func (o *AuthReference) GetTransType() string {
	if o == nil || IsNil(o.TransType) {
		var ret string
		return ret
	}
	return *o.TransType
}

// GetTransTypeOk returns a tuple with the TransType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthReference) GetTransTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TransType) {
		return nil, false
	}
	return o.TransType, true
}

// HasTransType returns a boolean if a field has been set.
func (o *AuthReference) HasTransType() bool {
	if o != nil && !IsNil(o.TransType) {
		return true
	}

	return false
}

// SetTransType gets a reference to the given string and assigns it to the TransType field.
func (o *AuthReference) SetTransType(v string) {
	o.TransType = &v
}

// GetTransno returns the Transno field value if set, zero value otherwise.
func (o *AuthReference) GetTransno() int32 {
	if o == nil || IsNil(o.Transno) {
		var ret int32
		return ret
	}
	return *o.Transno
}

// GetTransnoOk returns a tuple with the Transno field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthReference) GetTransnoOk() (*int32, bool) {
	if o == nil || IsNil(o.Transno) {
		return nil, false
	}
	return o.Transno, true
}

// HasTransno returns a boolean if a field has been set.
func (o *AuthReference) HasTransno() bool {
	if o != nil && !IsNil(o.Transno) {
		return true
	}

	return false
}

// SetTransno gets a reference to the given int32 and assigns it to the Transno field.
func (o *AuthReference) SetTransno(v int32) {
	o.Transno = &v
}

func (o AuthReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.AmountValue) {
		toSerialize["amount_value"] = o.AmountValue
	}
	if !IsNil(o.Atrn) {
		toSerialize["atrn"] = o.Atrn
	}
	if !IsNil(o.Authcode) {
		toSerialize["authcode"] = o.Authcode
	}
	if !IsNil(o.Batchno) {
		toSerialize["batchno"] = o.Batchno
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Datetime) {
		toSerialize["datetime"] = o.Datetime
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Maskedpan) {
		toSerialize["maskedpan"] = o.Maskedpan
	}
	if !IsNil(o.Merchantid) {
		toSerialize["merchantid"] = o.Merchantid
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.TransStatus) {
		toSerialize["trans_status"] = o.TransStatus
	}
	if !IsNil(o.TransType) {
		toSerialize["trans_type"] = o.TransType
	}
	if !IsNil(o.Transno) {
		toSerialize["transno"] = o.Transno
	}
	return toSerialize, nil
}

type NullableAuthReference struct {
	value *AuthReference
	isSet bool
}

func (v NullableAuthReference) Get() *AuthReference {
	return v.value
}

func (v *NullableAuthReference) Set(val *AuthReference) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthReference) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthReference(val *AuthReference) *NullableAuthReference {
	return &NullableAuthReference{value: val, isSet: true}
}

func (v NullableAuthReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
