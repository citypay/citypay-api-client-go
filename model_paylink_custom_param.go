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

// checks if the PaylinkCustomParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaylinkCustomParam{}

// PaylinkCustomParam struct for PaylinkCustomParam
type PaylinkCustomParam struct {
	// The type of entry mode. A value of 'pre' will pre-render the custom parameter before the payment screen. Any other value will result in the custom parameter being displayed on the payment screen.
	EntryMode *string `json:"entry_mode,omitempty"`
	// the type of html5 field, defaults to 'text'. Other options are 'dob' for a date of birth series of select list entry.
	FieldType *string `json:"field_type,omitempty"`
	// a group the parameter is linked with, allows for grouping with a title.
	Group *string `json:"group,omitempty"`
	// a label to show alongside the input.
	Label *string `json:"label,omitempty"`
	// whether the parameter is locked from entry.
	Locked *bool `json:"locked,omitempty"`
	// the name of the custom parameter used to converse with the submitter.
	Name string `json:"name"`
	// an index order for the parameter.
	Order *int32 `json:"order,omitempty"`
	// a regex pattern to validate the custom parameter with.
	Pattern *string `json:"pattern,omitempty"`
	// a placehold value to display in the input.
	Placeholder *string `json:"placeholder,omitempty"`
	// whether the field is required.
	Required *bool `json:"required,omitempty"`
	// a default value for the field.
	Value *string `json:"value,omitempty"`
}

type _PaylinkCustomParam PaylinkCustomParam

// NewPaylinkCustomParam instantiates a new PaylinkCustomParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaylinkCustomParam(name string) *PaylinkCustomParam {
	this := PaylinkCustomParam{}
	this.Name = name
	return &this
}

// NewPaylinkCustomParamWithDefaults instantiates a new PaylinkCustomParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaylinkCustomParamWithDefaults() *PaylinkCustomParam {
	this := PaylinkCustomParam{}
	return &this
}

// GetEntryMode returns the EntryMode field value if set, zero value otherwise.
func (o *PaylinkCustomParam) GetEntryMode() string {
	if o == nil || IsNil(o.EntryMode) {
		var ret string
		return ret
	}
	return *o.EntryMode
}

// GetEntryModeOk returns a tuple with the EntryMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCustomParam) GetEntryModeOk() (*string, bool) {
	if o == nil || IsNil(o.EntryMode) {
		return nil, false
	}
	return o.EntryMode, true
}

// HasEntryMode returns a boolean if a field has been set.
func (o *PaylinkCustomParam) HasEntryMode() bool {
	if o != nil && !IsNil(o.EntryMode) {
		return true
	}

	return false
}

// SetEntryMode gets a reference to the given string and assigns it to the EntryMode field.
func (o *PaylinkCustomParam) SetEntryMode(v string) {
	o.EntryMode = &v
}

// GetFieldType returns the FieldType field value if set, zero value otherwise.
func (o *PaylinkCustomParam) GetFieldType() string {
	if o == nil || IsNil(o.FieldType) {
		var ret string
		return ret
	}
	return *o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCustomParam) GetFieldTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FieldType) {
		return nil, false
	}
	return o.FieldType, true
}

// HasFieldType returns a boolean if a field has been set.
func (o *PaylinkCustomParam) HasFieldType() bool {
	if o != nil && !IsNil(o.FieldType) {
		return true
	}

	return false
}

// SetFieldType gets a reference to the given string and assigns it to the FieldType field.
func (o *PaylinkCustomParam) SetFieldType(v string) {
	o.FieldType = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *PaylinkCustomParam) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCustomParam) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *PaylinkCustomParam) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *PaylinkCustomParam) SetGroup(v string) {
	o.Group = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PaylinkCustomParam) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCustomParam) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PaylinkCustomParam) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PaylinkCustomParam) SetLabel(v string) {
	o.Label = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *PaylinkCustomParam) GetLocked() bool {
	if o == nil || IsNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCustomParam) GetLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *PaylinkCustomParam) HasLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *PaylinkCustomParam) SetLocked(v bool) {
	o.Locked = &v
}

// GetName returns the Name field value
func (o *PaylinkCustomParam) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PaylinkCustomParam) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PaylinkCustomParam) SetName(v string) {
	o.Name = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *PaylinkCustomParam) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCustomParam) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *PaylinkCustomParam) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *PaylinkCustomParam) SetOrder(v int32) {
	o.Order = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *PaylinkCustomParam) GetPattern() string {
	if o == nil || IsNil(o.Pattern) {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCustomParam) GetPatternOk() (*string, bool) {
	if o == nil || IsNil(o.Pattern) {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *PaylinkCustomParam) HasPattern() bool {
	if o != nil && !IsNil(o.Pattern) {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *PaylinkCustomParam) SetPattern(v string) {
	o.Pattern = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *PaylinkCustomParam) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder) {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCustomParam) GetPlaceholderOk() (*string, bool) {
	if o == nil || IsNil(o.Placeholder) {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *PaylinkCustomParam) HasPlaceholder() bool {
	if o != nil && !IsNil(o.Placeholder) {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *PaylinkCustomParam) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *PaylinkCustomParam) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCustomParam) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *PaylinkCustomParam) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *PaylinkCustomParam) SetRequired(v bool) {
	o.Required = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PaylinkCustomParam) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaylinkCustomParam) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PaylinkCustomParam) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PaylinkCustomParam) SetValue(v string) {
	o.Value = &v
}

func (o PaylinkCustomParam) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaylinkCustomParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntryMode) {
		toSerialize["entry_mode"] = o.EntryMode
	}
	if !IsNil(o.FieldType) {
		toSerialize["field_type"] = o.FieldType
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Pattern) {
		toSerialize["pattern"] = o.Pattern
	}
	if !IsNil(o.Placeholder) {
		toSerialize["placeholder"] = o.Placeholder
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

func (o *PaylinkCustomParam) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varPaylinkCustomParam := _PaylinkCustomParam{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaylinkCustomParam)

	if err != nil {
		return err
	}

	*o = PaylinkCustomParam(varPaylinkCustomParam)

	return err
}

type NullablePaylinkCustomParam struct {
	value *PaylinkCustomParam
	isSet bool
}

func (v NullablePaylinkCustomParam) Get() *PaylinkCustomParam {
	return v.value
}

func (v *NullablePaylinkCustomParam) Set(val *PaylinkCustomParam) {
	v.value = val
	v.isSet = true
}

func (v NullablePaylinkCustomParam) IsSet() bool {
	return v.isSet
}

func (v *NullablePaylinkCustomParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaylinkCustomParam(val *PaylinkCustomParam) *NullablePaylinkCustomParam {
	return &NullablePaylinkCustomParam{value: val, isSet: true}
}

func (v NullablePaylinkCustomParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaylinkCustomParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
