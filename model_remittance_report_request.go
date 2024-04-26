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

// checks if the RemittanceReportRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemittanceReportRequest{}

// RemittanceReportRequest struct for RemittanceReportRequest
type RemittanceReportRequest struct {
	// Start date (YYYY-MM-DD) for batch Retrieval range, inclusive. Maximum value is 3 years ago.
	DateFrom *string `json:"date_from,omitempty"`
	// End date (YYYY-MM-DD) for batch Retrieval range, inclusive.
	DateUntil *string `json:"date_until,omitempty"`
	// The maximum number of results to return in a single response. This value is used to limit the size of data returned by the API, enhancing performance and manageability. Values should be between 5 and 250.
	MaxResults *int32  `json:"maxResults,omitempty"`
	MerchantId []int32 `json:"merchant_id,omitempty"`
	// A token that identifies the starting point of the page of results to be returned. An empty value indicates the start of the dataset. When supplied, it is validated and used to fetch the subsequent page of results. This token is typically obtained from the response of a previous pagination request.
	NextToken *string `json:"nextToken,omitempty"`
	// Specifies the field by which results are ordered. Available fields are [trans_no,date_when,amount]. By default, fields are ordered by OrderByExpression(trans_no,ASC). To order in descending order, prefix with '-' or suffix with ' DESC'.
	OrderBy *string `json:"orderBy,omitempty"`
}

// NewRemittanceReportRequest instantiates a new RemittanceReportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemittanceReportRequest() *RemittanceReportRequest {
	this := RemittanceReportRequest{}
	return &this
}

// NewRemittanceReportRequestWithDefaults instantiates a new RemittanceReportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemittanceReportRequestWithDefaults() *RemittanceReportRequest {
	this := RemittanceReportRequest{}
	return &this
}

// GetDateFrom returns the DateFrom field value if set, zero value otherwise.
func (o *RemittanceReportRequest) GetDateFrom() string {
	if o == nil || IsNil(o.DateFrom) {
		var ret string
		return ret
	}
	return *o.DateFrom
}

// GetDateFromOk returns a tuple with the DateFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemittanceReportRequest) GetDateFromOk() (*string, bool) {
	if o == nil || IsNil(o.DateFrom) {
		return nil, false
	}
	return o.DateFrom, true
}

// HasDateFrom returns a boolean if a field has been set.
func (o *RemittanceReportRequest) HasDateFrom() bool {
	if o != nil && !IsNil(o.DateFrom) {
		return true
	}

	return false
}

// SetDateFrom gets a reference to the given string and assigns it to the DateFrom field.
func (o *RemittanceReportRequest) SetDateFrom(v string) {
	o.DateFrom = &v
}

// GetDateUntil returns the DateUntil field value if set, zero value otherwise.
func (o *RemittanceReportRequest) GetDateUntil() string {
	if o == nil || IsNil(o.DateUntil) {
		var ret string
		return ret
	}
	return *o.DateUntil
}

// GetDateUntilOk returns a tuple with the DateUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemittanceReportRequest) GetDateUntilOk() (*string, bool) {
	if o == nil || IsNil(o.DateUntil) {
		return nil, false
	}
	return o.DateUntil, true
}

// HasDateUntil returns a boolean if a field has been set.
func (o *RemittanceReportRequest) HasDateUntil() bool {
	if o != nil && !IsNil(o.DateUntil) {
		return true
	}

	return false
}

// SetDateUntil gets a reference to the given string and assigns it to the DateUntil field.
func (o *RemittanceReportRequest) SetDateUntil(v string) {
	o.DateUntil = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *RemittanceReportRequest) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemittanceReportRequest) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *RemittanceReportRequest) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *RemittanceReportRequest) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *RemittanceReportRequest) GetMerchantId() []int32 {
	if o == nil || IsNil(o.MerchantId) {
		var ret []int32
		return ret
	}
	return o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemittanceReportRequest) GetMerchantIdOk() ([]int32, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *RemittanceReportRequest) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given []int32 and assigns it to the MerchantId field.
func (o *RemittanceReportRequest) SetMerchantId(v []int32) {
	o.MerchantId = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *RemittanceReportRequest) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemittanceReportRequest) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *RemittanceReportRequest) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *RemittanceReportRequest) SetNextToken(v string) {
	o.NextToken = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *RemittanceReportRequest) GetOrderBy() string {
	if o == nil || IsNil(o.OrderBy) {
		var ret string
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemittanceReportRequest) GetOrderByOk() (*string, bool) {
	if o == nil || IsNil(o.OrderBy) {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *RemittanceReportRequest) HasOrderBy() bool {
	if o != nil && !IsNil(o.OrderBy) {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given string and assigns it to the OrderBy field.
func (o *RemittanceReportRequest) SetOrderBy(v string) {
	o.OrderBy = &v
}

func (o RemittanceReportRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemittanceReportRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateFrom) {
		toSerialize["date_from"] = o.DateFrom
	}
	if !IsNil(o.DateUntil) {
		toSerialize["date_until"] = o.DateUntil
	}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchant_id"] = o.MerchantId
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.OrderBy) {
		toSerialize["orderBy"] = o.OrderBy
	}
	return toSerialize, nil
}

type NullableRemittanceReportRequest struct {
	value *RemittanceReportRequest
	isSet bool
}

func (v NullableRemittanceReportRequest) Get() *RemittanceReportRequest {
	return v.value
}

func (v *NullableRemittanceReportRequest) Set(val *RemittanceReportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRemittanceReportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRemittanceReportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemittanceReportRequest(val *RemittanceReportRequest) *NullableRemittanceReportRequest {
	return &NullableRemittanceReportRequest{value: val, isSet: true}
}

func (v NullableRemittanceReportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemittanceReportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
