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

// checks if the PaylinkTokenStatusChangeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaylinkTokenStatusChangeResponse{}

// PaylinkTokenStatusChangeResponse struct for PaylinkTokenStatusChangeResponse
type PaylinkTokenStatusChangeResponse struct {
	// The count of items returned in this page.
	Count *int32 `json:"count,omitempty"`
	// The max results requested in this page.
	MaxResults *int32 `json:"maxResults,omitempty"`
	// A token that identifies the starting point of the page of results to be returned. An empty value indicates the start of the dataset. When supplied, it is validated and used to fetch the subsequent page of results. This token is typically obtained from the response of a previous pagination request.
	NextToken *string              `json:"nextToken,omitempty"`
	Tokens    []PaylinkTokenStatus `json:"tokens"`
}

type _PaylinkTokenStatusChangeResponse PaylinkTokenStatusChangeResponse

// NewPaylinkTokenStatusChangeResponse instantiates a new PaylinkTokenStatusChangeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaylinkTokenStatusChangeResponse(tokens []PaylinkTokenStatus) *PaylinkTokenStatusChangeResponse {
	this := PaylinkTokenStatusChangeResponse{}
	this.Tokens = tokens
	return &this
}

// NewPaylinkTokenStatusChangeResponseWithDefaults instantiates a new PaylinkTokenStatusChangeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaylinkTokenStatusChangeResponseWithDefaults() *PaylinkTokenStatusChangeResponse {
	this := PaylinkTokenStatusChangeResponse{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaylinkTokenStatusChangeResponse) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkTokenStatusChangeResponse) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaylinkTokenStatusChangeResponse) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaylinkTokenStatusChangeResponse) SetCount(v int32) {
	o.Count = &v
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *PaylinkTokenStatusChangeResponse) GetMaxResults() int32 {
	if o == nil || IsNil(o.MaxResults) {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkTokenStatusChangeResponse) GetMaxResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxResults) {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *PaylinkTokenStatusChangeResponse) HasMaxResults() bool {
	if o != nil && !IsNil(o.MaxResults) {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *PaylinkTokenStatusChangeResponse) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *PaylinkTokenStatusChangeResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkTokenStatusChangeResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *PaylinkTokenStatusChangeResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *PaylinkTokenStatusChangeResponse) SetNextToken(v string) {
	o.NextToken = &v
}

// GetTokens returns the Tokens field value
func (o *PaylinkTokenStatusChangeResponse) GetTokens() []PaylinkTokenStatus {
	if o == nil {
		var ret []PaylinkTokenStatus
		return ret
	}

	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value
// and a boolean to check if the value has been set.
func (o *PaylinkTokenStatusChangeResponse) GetTokensOk() ([]PaylinkTokenStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tokens, true
}

// SetTokens sets field value
func (o *PaylinkTokenStatusChangeResponse) SetTokens(v []PaylinkTokenStatus) {
	o.Tokens = v
}

func (o PaylinkTokenStatusChangeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaylinkTokenStatusChangeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.MaxResults) {
		toSerialize["maxResults"] = o.MaxResults
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	toSerialize["tokens"] = o.Tokens
	return toSerialize, nil
}

func (o *PaylinkTokenStatusChangeResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tokens",
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

	varPaylinkTokenStatusChangeResponse := _PaylinkTokenStatusChangeResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaylinkTokenStatusChangeResponse)

	if err != nil {
		return err
	}

	*o = PaylinkTokenStatusChangeResponse(varPaylinkTokenStatusChangeResponse)

	return err
}

type NullablePaylinkTokenStatusChangeResponse struct {
	value *PaylinkTokenStatusChangeResponse
	isSet bool
}

func (v NullablePaylinkTokenStatusChangeResponse) Get() *PaylinkTokenStatusChangeResponse {
	return v.value
}

func (v *NullablePaylinkTokenStatusChangeResponse) Set(val *PaylinkTokenStatusChangeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaylinkTokenStatusChangeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaylinkTokenStatusChangeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaylinkTokenStatusChangeResponse(val *PaylinkTokenStatusChangeResponse) *NullablePaylinkTokenStatusChangeResponse {
	return &NullablePaylinkTokenStatusChangeResponse{value: val, isSet: true}
}

func (v NullablePaylinkTokenStatusChangeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaylinkTokenStatusChangeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
