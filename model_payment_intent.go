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

// checks if the PaymentIntent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentIntent{}

// PaymentIntent struct for PaymentIntent
type PaymentIntent struct {
	// The amount to authorise in the lowest unit of currency with a variable length to a maximum of 12 digits.  No decimal points are to be included and no divisional characters such as 1,024.  The amount should be the total amount required for the transaction.  For example with GBP £1,021.95 the amount value is 102195.
	Amount int32 `json:"amount"`
	// A policy value which determines whether an AVS postcode policy is enforced or bypassed.  Values are:   `0` for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   `1` for an enforced policy. Transactions that are enforced will be rejected if the AVS postcode numeric value does not match.   `2` to bypass. Transactions that are bypassed will be allowed through even if the postcode did not match.   `3` to ignore. Transactions that are ignored will bypass the result and not send postcode details for authorisation.
	AvsPostcodePolicy *string         `json:"avs_postcode_policy,omitempty"`
	BillTo            *ContactDetails `json:"bill_to,omitempty"`
	// The Card Security Code (CSC) (also known as CV2/CVV2) is normally found on the back of the card (American Express has it on the front). The value helps to identify possession of the card as it is not available within the chip or magnetic swipe.  When forwarding the CSC, please ensure the value is a string as some values start with 0 and this will be stripped out by any integer parsing.  The CSC number aids fraud prevention in Mail Order and Internet payments.  Business rules are available on your account to identify whether to accept or decline transactions based on mismatched results of the CSC.  The Payment Card Industry (PCI) requires that at no stage of a transaction should the CSC be stored.  This applies to all entities handling card data.  It should also not be used in any hashing process.  CityPay do not store the value and have no method of retrieving the value once the transaction has been processed. For this reason, duplicate checking is unable to determine the CSC in its duplication check algorithm.
	Csc *string `json:"csc,omitempty"`
	// A policy value which determines whether a CSC policy is enforced or bypassed.  Values are:   `0` for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   `1` for an enforced policy. Transactions that are enforced will be rejected if the CSC value does not match.   `2` to bypass. Transactions that are bypassed will be allowed through even if the CSC did not match.   `3` to ignore. Transactions that are ignored will bypass the result and not send the CSC details for authorisation.
	CscPolicy *string `json:"csc_policy,omitempty"`
	// The processing currency for the transaction. Will default to the merchant account currency.
	Currency *string `json:"currency,omitempty"`
	// A policy value which determines whether a duplication policy is enforced or bypassed. A duplication check has a window of time set against your account within which it can action. If a previous transaction with matching values occurred within the window, any subsequent transaction will result in a T001 result.  Values are   `0` for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   `1` for an enforced policy. Transactions that are enforced will be checked for duplication within the duplication window.   `2` to bypass. Transactions that are bypassed will not be checked for duplication within the duplication window.   `3` to ignore. Transactions that are ignored will have the same affect as bypass.
	DuplicatePolicy *string `json:"duplicate_policy,omitempty"`
	// The identifier of the transaction to process. The value should be a valid reference and may be used to perform  post processing actions and to aid in reconciliation of transactions.  The value should be a valid printable string with ASCII character ranges from 0x32 to 0x127.  The identifier is recommended to be distinct for each transaction such as a [random unique identifier](https://en.wikipedia.org/wiki/Universally_unique_identifier) this will aid in ensuring each transaction is identifiable.  When transactions are processed they are also checked for duplicate requests. Changing the identifier on a subsequent request will ensure that a transaction is considered as different.
	Identifier string `json:"identifier"`
	// A policy value which determines whether an AVS address policy is enforced, bypassed or ignored.  Values are:   `0` for the default policy (default value if not supplied). Your default values are determined by your account manager on setup of the account.   `1` for an enforced policy. Transactions that are enforced will be rejected if the AVS address numeric value does not match.   `2` to bypass. Transactions that are bypassed will be allowed through even if the address did not match.   `3` to ignore. Transactions that are ignored will bypass the result and not send address numeric details for authorisation.
	MatchAvsa *string         `json:"match_avsa,omitempty"`
	ShipTo    *ContactDetails `json:"ship_to,omitempty"`
	Tag       []string        `json:"tag,omitempty"`
	// Further information that can be added to the transaction will display in reporting. Can be used for flexible values such as operator id.
	TransInfo *string `json:"trans_info,omitempty"`
	// The type of transaction being submitted. Normally this value is not required and your account manager may request that you set this field.
	TransType *string `json:"trans_type,omitempty"`
}

type _PaymentIntent PaymentIntent

// NewPaymentIntent instantiates a new PaymentIntent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentIntent(amount int32, identifier string) *PaymentIntent {
	this := PaymentIntent{}
	this.Amount = amount
	this.Identifier = identifier
	return &this
}

// NewPaymentIntentWithDefaults instantiates a new PaymentIntent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentIntentWithDefaults() *PaymentIntent {
	this := PaymentIntent{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PaymentIntent) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentIntent) SetAmount(v int32) {
	o.Amount = v
}

// GetAvsPostcodePolicy returns the AvsPostcodePolicy field value if set, zero value otherwise.
func (o *PaymentIntent) GetAvsPostcodePolicy() string {
	if o == nil || IsNil(o.AvsPostcodePolicy) {
		var ret string
		return ret
	}
	return *o.AvsPostcodePolicy
}

// GetAvsPostcodePolicyOk returns a tuple with the AvsPostcodePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetAvsPostcodePolicyOk() (*string, bool) {
	if o == nil || IsNil(o.AvsPostcodePolicy) {
		return nil, false
	}
	return o.AvsPostcodePolicy, true
}

// HasAvsPostcodePolicy returns a boolean if a field has been set.
func (o *PaymentIntent) HasAvsPostcodePolicy() bool {
	if o != nil && !IsNil(o.AvsPostcodePolicy) {
		return true
	}

	return false
}

// SetAvsPostcodePolicy gets a reference to the given string and assigns it to the AvsPostcodePolicy field.
func (o *PaymentIntent) SetAvsPostcodePolicy(v string) {
	o.AvsPostcodePolicy = &v
}

// GetBillTo returns the BillTo field value if set, zero value otherwise.
func (o *PaymentIntent) GetBillTo() ContactDetails {
	if o == nil || IsNil(o.BillTo) {
		var ret ContactDetails
		return ret
	}
	return *o.BillTo
}

// GetBillToOk returns a tuple with the BillTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetBillToOk() (*ContactDetails, bool) {
	if o == nil || IsNil(o.BillTo) {
		return nil, false
	}
	return o.BillTo, true
}

// HasBillTo returns a boolean if a field has been set.
func (o *PaymentIntent) HasBillTo() bool {
	if o != nil && !IsNil(o.BillTo) {
		return true
	}

	return false
}

// SetBillTo gets a reference to the given ContactDetails and assigns it to the BillTo field.
func (o *PaymentIntent) SetBillTo(v ContactDetails) {
	o.BillTo = &v
}

// GetCsc returns the Csc field value if set, zero value otherwise.
func (o *PaymentIntent) GetCsc() string {
	if o == nil || IsNil(o.Csc) {
		var ret string
		return ret
	}
	return *o.Csc
}

// GetCscOk returns a tuple with the Csc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetCscOk() (*string, bool) {
	if o == nil || IsNil(o.Csc) {
		return nil, false
	}
	return o.Csc, true
}

// HasCsc returns a boolean if a field has been set.
func (o *PaymentIntent) HasCsc() bool {
	if o != nil && !IsNil(o.Csc) {
		return true
	}

	return false
}

// SetCsc gets a reference to the given string and assigns it to the Csc field.
func (o *PaymentIntent) SetCsc(v string) {
	o.Csc = &v
}

// GetCscPolicy returns the CscPolicy field value if set, zero value otherwise.
func (o *PaymentIntent) GetCscPolicy() string {
	if o == nil || IsNil(o.CscPolicy) {
		var ret string
		return ret
	}
	return *o.CscPolicy
}

// GetCscPolicyOk returns a tuple with the CscPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetCscPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.CscPolicy) {
		return nil, false
	}
	return o.CscPolicy, true
}

// HasCscPolicy returns a boolean if a field has been set.
func (o *PaymentIntent) HasCscPolicy() bool {
	if o != nil && !IsNil(o.CscPolicy) {
		return true
	}

	return false
}

// SetCscPolicy gets a reference to the given string and assigns it to the CscPolicy field.
func (o *PaymentIntent) SetCscPolicy(v string) {
	o.CscPolicy = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PaymentIntent) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PaymentIntent) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PaymentIntent) SetCurrency(v string) {
	o.Currency = &v
}

// GetDuplicatePolicy returns the DuplicatePolicy field value if set, zero value otherwise.
func (o *PaymentIntent) GetDuplicatePolicy() string {
	if o == nil || IsNil(o.DuplicatePolicy) {
		var ret string
		return ret
	}
	return *o.DuplicatePolicy
}

// GetDuplicatePolicyOk returns a tuple with the DuplicatePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetDuplicatePolicyOk() (*string, bool) {
	if o == nil || IsNil(o.DuplicatePolicy) {
		return nil, false
	}
	return o.DuplicatePolicy, true
}

// HasDuplicatePolicy returns a boolean if a field has been set.
func (o *PaymentIntent) HasDuplicatePolicy() bool {
	if o != nil && !IsNil(o.DuplicatePolicy) {
		return true
	}

	return false
}

// SetDuplicatePolicy gets a reference to the given string and assigns it to the DuplicatePolicy field.
func (o *PaymentIntent) SetDuplicatePolicy(v string) {
	o.DuplicatePolicy = &v
}

// GetIdentifier returns the Identifier field value
func (o *PaymentIntent) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *PaymentIntent) SetIdentifier(v string) {
	o.Identifier = v
}

// GetMatchAvsa returns the MatchAvsa field value if set, zero value otherwise.
func (o *PaymentIntent) GetMatchAvsa() string {
	if o == nil || IsNil(o.MatchAvsa) {
		var ret string
		return ret
	}
	return *o.MatchAvsa
}

// GetMatchAvsaOk returns a tuple with the MatchAvsa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetMatchAvsaOk() (*string, bool) {
	if o == nil || IsNil(o.MatchAvsa) {
		return nil, false
	}
	return o.MatchAvsa, true
}

// HasMatchAvsa returns a boolean if a field has been set.
func (o *PaymentIntent) HasMatchAvsa() bool {
	if o != nil && !IsNil(o.MatchAvsa) {
		return true
	}

	return false
}

// SetMatchAvsa gets a reference to the given string and assigns it to the MatchAvsa field.
func (o *PaymentIntent) SetMatchAvsa(v string) {
	o.MatchAvsa = &v
}

// GetShipTo returns the ShipTo field value if set, zero value otherwise.
func (o *PaymentIntent) GetShipTo() ContactDetails {
	if o == nil || IsNil(o.ShipTo) {
		var ret ContactDetails
		return ret
	}
	return *o.ShipTo
}

// GetShipToOk returns a tuple with the ShipTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetShipToOk() (*ContactDetails, bool) {
	if o == nil || IsNil(o.ShipTo) {
		return nil, false
	}
	return o.ShipTo, true
}

// HasShipTo returns a boolean if a field has been set.
func (o *PaymentIntent) HasShipTo() bool {
	if o != nil && !IsNil(o.ShipTo) {
		return true
	}

	return false
}

// SetShipTo gets a reference to the given ContactDetails and assigns it to the ShipTo field.
func (o *PaymentIntent) SetShipTo(v ContactDetails) {
	o.ShipTo = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *PaymentIntent) GetTag() []string {
	if o == nil || IsNil(o.Tag) {
		var ret []string
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetTagOk() ([]string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *PaymentIntent) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given []string and assigns it to the Tag field.
func (o *PaymentIntent) SetTag(v []string) {
	o.Tag = v
}

// GetTransInfo returns the TransInfo field value if set, zero value otherwise.
func (o *PaymentIntent) GetTransInfo() string {
	if o == nil || IsNil(o.TransInfo) {
		var ret string
		return ret
	}
	return *o.TransInfo
}

// GetTransInfoOk returns a tuple with the TransInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetTransInfoOk() (*string, bool) {
	if o == nil || IsNil(o.TransInfo) {
		return nil, false
	}
	return o.TransInfo, true
}

// HasTransInfo returns a boolean if a field has been set.
func (o *PaymentIntent) HasTransInfo() bool {
	if o != nil && !IsNil(o.TransInfo) {
		return true
	}

	return false
}

// SetTransInfo gets a reference to the given string and assigns it to the TransInfo field.
func (o *PaymentIntent) SetTransInfo(v string) {
	o.TransInfo = &v
}

// GetTransType returns the TransType field value if set, zero value otherwise.
func (o *PaymentIntent) GetTransType() string {
	if o == nil || IsNil(o.TransType) {
		var ret string
		return ret
	}
	return *o.TransType
}

// GetTransTypeOk returns a tuple with the TransType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentIntent) GetTransTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TransType) {
		return nil, false
	}
	return o.TransType, true
}

// HasTransType returns a boolean if a field has been set.
func (o *PaymentIntent) HasTransType() bool {
	if o != nil && !IsNil(o.TransType) {
		return true
	}

	return false
}

// SetTransType gets a reference to the given string and assigns it to the TransType field.
func (o *PaymentIntent) SetTransType(v string) {
	o.TransType = &v
}

func (o PaymentIntent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentIntent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.AvsPostcodePolicy) {
		toSerialize["avs_postcode_policy"] = o.AvsPostcodePolicy
	}
	if !IsNil(o.BillTo) {
		toSerialize["bill_to"] = o.BillTo
	}
	if !IsNil(o.Csc) {
		toSerialize["csc"] = o.Csc
	}
	if !IsNil(o.CscPolicy) {
		toSerialize["csc_policy"] = o.CscPolicy
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.DuplicatePolicy) {
		toSerialize["duplicate_policy"] = o.DuplicatePolicy
	}
	toSerialize["identifier"] = o.Identifier
	if !IsNil(o.MatchAvsa) {
		toSerialize["match_avsa"] = o.MatchAvsa
	}
	if !IsNil(o.ShipTo) {
		toSerialize["ship_to"] = o.ShipTo
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.TransInfo) {
		toSerialize["trans_info"] = o.TransInfo
	}
	if !IsNil(o.TransType) {
		toSerialize["trans_type"] = o.TransType
	}
	return toSerialize, nil
}

func (o *PaymentIntent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"identifier",
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

	varPaymentIntent := _PaymentIntent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaymentIntent)

	if err != nil {
		return err
	}

	*o = PaymentIntent(varPaymentIntent)

	return err
}

type NullablePaymentIntent struct {
	value *PaymentIntent
	isSet bool
}

func (v NullablePaymentIntent) Get() *PaymentIntent {
	return v.value
}

func (v *NullablePaymentIntent) Set(val *PaymentIntent) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentIntent) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentIntent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentIntent(val *PaymentIntent) *NullablePaymentIntent {
	return &NullablePaymentIntent{value: val, isSet: true}
}

func (v NullablePaymentIntent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentIntent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}