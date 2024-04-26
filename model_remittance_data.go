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

// checks if the RemittanceData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemittanceData{}

// RemittanceData struct for RemittanceData
type RemittanceData struct {
	// Represents the date and time when the remittance was processed. This timestamp follows the ISO 8601 format for datetime representation.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// Represents the net amount after accounting for refunds. This is calculated as SalesAmount - RefundAmount and expressed in the smallest currency unit.
	NetAmount *int32 `json:"net_amount,omitempty"`
	// The total amount refunded to customers.
	RefundAmount *int32 `json:"refund_amount,omitempty"`
	// The total number of refund transactions processed. This figure helps in understanding the frequency of refunds relative to sales.
	RefundCount *int32 `json:"refund_count,omitempty"`
	// The total monetary amount of sales transactions.
	SalesAmount *int32 `json:"sales_amount,omitempty"`
	// Indicates the total number of sales transactions that occurred. This count provides insight into the volume of sales.
	SalesCount *int32 `json:"sales_count,omitempty"`
}

// NewRemittanceData instantiates a new RemittanceData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemittanceData() *RemittanceData {
	this := RemittanceData{}
	return &this
}

// NewRemittanceDataWithDefaults instantiates a new RemittanceData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemittanceDataWithDefaults() *RemittanceData {
	this := RemittanceData{}
	return &this
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *RemittanceData) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemittanceData) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *RemittanceData) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *RemittanceData) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetNetAmount returns the NetAmount field value if set, zero value otherwise.
func (o *RemittanceData) GetNetAmount() int32 {
	if o == nil || IsNil(o.NetAmount) {
		var ret int32
		return ret
	}
	return *o.NetAmount
}

// GetNetAmountOk returns a tuple with the NetAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemittanceData) GetNetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.NetAmount) {
		return nil, false
	}
	return o.NetAmount, true
}

// HasNetAmount returns a boolean if a field has been set.
func (o *RemittanceData) HasNetAmount() bool {
	if o != nil && !IsNil(o.NetAmount) {
		return true
	}

	return false
}

// SetNetAmount gets a reference to the given int32 and assigns it to the NetAmount field.
func (o *RemittanceData) SetNetAmount(v int32) {
	o.NetAmount = &v
}

// GetRefundAmount returns the RefundAmount field value if set, zero value otherwise.
func (o *RemittanceData) GetRefundAmount() int32 {
	if o == nil || IsNil(o.RefundAmount) {
		var ret int32
		return ret
	}
	return *o.RefundAmount
}

// GetRefundAmountOk returns a tuple with the RefundAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemittanceData) GetRefundAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.RefundAmount) {
		return nil, false
	}
	return o.RefundAmount, true
}

// HasRefundAmount returns a boolean if a field has been set.
func (o *RemittanceData) HasRefundAmount() bool {
	if o != nil && !IsNil(o.RefundAmount) {
		return true
	}

	return false
}

// SetRefundAmount gets a reference to the given int32 and assigns it to the RefundAmount field.
func (o *RemittanceData) SetRefundAmount(v int32) {
	o.RefundAmount = &v
}

// GetRefundCount returns the RefundCount field value if set, zero value otherwise.
func (o *RemittanceData) GetRefundCount() int32 {
	if o == nil || IsNil(o.RefundCount) {
		var ret int32
		return ret
	}
	return *o.RefundCount
}

// GetRefundCountOk returns a tuple with the RefundCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemittanceData) GetRefundCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RefundCount) {
		return nil, false
	}
	return o.RefundCount, true
}

// HasRefundCount returns a boolean if a field has been set.
func (o *RemittanceData) HasRefundCount() bool {
	if o != nil && !IsNil(o.RefundCount) {
		return true
	}

	return false
}

// SetRefundCount gets a reference to the given int32 and assigns it to the RefundCount field.
func (o *RemittanceData) SetRefundCount(v int32) {
	o.RefundCount = &v
}

// GetSalesAmount returns the SalesAmount field value if set, zero value otherwise.
func (o *RemittanceData) GetSalesAmount() int32 {
	if o == nil || IsNil(o.SalesAmount) {
		var ret int32
		return ret
	}
	return *o.SalesAmount
}

// GetSalesAmountOk returns a tuple with the SalesAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemittanceData) GetSalesAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.SalesAmount) {
		return nil, false
	}
	return o.SalesAmount, true
}

// HasSalesAmount returns a boolean if a field has been set.
func (o *RemittanceData) HasSalesAmount() bool {
	if o != nil && !IsNil(o.SalesAmount) {
		return true
	}

	return false
}

// SetSalesAmount gets a reference to the given int32 and assigns it to the SalesAmount field.
func (o *RemittanceData) SetSalesAmount(v int32) {
	o.SalesAmount = &v
}

// GetSalesCount returns the SalesCount field value if set, zero value otherwise.
func (o *RemittanceData) GetSalesCount() int32 {
	if o == nil || IsNil(o.SalesCount) {
		var ret int32
		return ret
	}
	return *o.SalesCount
}

// GetSalesCountOk returns a tuple with the SalesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemittanceData) GetSalesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SalesCount) {
		return nil, false
	}
	return o.SalesCount, true
}

// HasSalesCount returns a boolean if a field has been set.
func (o *RemittanceData) HasSalesCount() bool {
	if o != nil && !IsNil(o.SalesCount) {
		return true
	}

	return false
}

// SetSalesCount gets a reference to the given int32 and assigns it to the SalesCount field.
func (o *RemittanceData) SetSalesCount(v int32) {
	o.SalesCount = &v
}

func (o RemittanceData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemittanceData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateCreated) {
		toSerialize["date_created"] = o.DateCreated
	}
	if !IsNil(o.NetAmount) {
		toSerialize["net_amount"] = o.NetAmount
	}
	if !IsNil(o.RefundAmount) {
		toSerialize["refund_amount"] = o.RefundAmount
	}
	if !IsNil(o.RefundCount) {
		toSerialize["refund_count"] = o.RefundCount
	}
	if !IsNil(o.SalesAmount) {
		toSerialize["sales_amount"] = o.SalesAmount
	}
	if !IsNil(o.SalesCount) {
		toSerialize["sales_count"] = o.SalesCount
	}
	return toSerialize, nil
}

type NullableRemittanceData struct {
	value *RemittanceData
	isSet bool
}

func (v NullableRemittanceData) Get() *RemittanceData {
	return v.value
}

func (v *NullableRemittanceData) Set(val *RemittanceData) {
	v.value = val
	v.isSet = true
}

func (v NullableRemittanceData) IsSet() bool {
	return v.isSet
}

func (v *NullableRemittanceData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemittanceData(val *RemittanceData) *NullableRemittanceData {
	return &NullableRemittanceData{value: val, isSet: true}
}

func (v NullableRemittanceData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemittanceData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}