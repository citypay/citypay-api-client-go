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

// checks if the AirlineAdvice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AirlineAdvice{}

// AirlineAdvice struct for AirlineAdvice
type AirlineAdvice struct {
	// The name of the airline carrier that generated the tickets for airline travel.
	CarrierName string `json:"carrier_name"`
	// true if a conjunction ticket (with additional coupons) was issued for an itinerary with more than four segments. Defaults to false.
	ConjunctionTicketIndicator *bool `json:"conjunction_ticket_indicator,omitempty"`
	// The Electronic Ticket Indicator, a code that indicates if an electronic ticket was issued.  Defaults to true.
	EticketIndicator *bool `json:"eticket_indicator,omitempty"`
	// A value that indicates the number of air travel segments included on this ticket.  Valid entries include the numerals “0” through “4”. Required only if the transaction type is TKT or EXC.
	NoAirSegments *int32 `json:"no_air_segments,omitempty"`
	// The number of people in the party.
	NumberInParty *int32 `json:"number_in_party,omitempty"`
	// Required if transaction type is EXC.
	OriginalTicketNo *string `json:"original_ticket_no,omitempty"`
	// The name of the passenger when the traveller is not the card member that purchased the ticket. Required only if the transaction type is TKT or EXC.
	PassengerName *string         `json:"passenger_name,omitempty"`
	Segment1      AirlineSegment  `json:"segment1"`
	Segment2      *AirlineSegment `json:"segment2,omitempty"`
	Segment3      *AirlineSegment `json:"segment3,omitempty"`
	Segment4      *AirlineSegment `json:"segment4,omitempty"`
	// The name of the city town or village where the transaction took place.
	TicketIssueCity string `json:"ticket_issue_city"`
	// The date the ticket was issued in ISO Date format (yyyy-MM-dd).
	TicketIssueDate string `json:"ticket_issue_date"`
	// The name of the agency generating the ticket.
	TicketIssueName string `json:"ticket_issue_name"`
	// This must be a valid ticket number, i.e. numeric (the first 3 digits must represent the valid IATA plate carrier code). The final check digit should be validated prior to submission. On credit charges, this field should contain the number of the original ticket, and not of a replacement.
	TicketNo string `json:"ticket_no"`
	// This field contains the Transaction Type code assigned to this transaction. Valid codes include:   - `TKT` = Ticket Purchase   - `REF` = Refund   - `EXC` = Exchange Ticket   - `MSC` = Miscellaneous (non-Ticket Purchase- and non-Exchange Ticket-related transactions only).
	TransactionType string `json:"transaction_type"`
}

type _AirlineAdvice AirlineAdvice

// NewAirlineAdvice instantiates a new AirlineAdvice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAirlineAdvice(carrierName string, segment1 AirlineSegment, ticketIssueCity string, ticketIssueDate string, ticketIssueName string, ticketNo string, transactionType string) *AirlineAdvice {
	this := AirlineAdvice{}
	this.CarrierName = carrierName
	this.Segment1 = segment1
	this.TicketIssueCity = ticketIssueCity
	this.TicketIssueDate = ticketIssueDate
	this.TicketIssueName = ticketIssueName
	this.TicketNo = ticketNo
	this.TransactionType = transactionType
	return &this
}

// NewAirlineAdviceWithDefaults instantiates a new AirlineAdvice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAirlineAdviceWithDefaults() *AirlineAdvice {
	this := AirlineAdvice{}
	return &this
}

// GetCarrierName returns the CarrierName field value
func (o *AirlineAdvice) GetCarrierName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierName
}

// GetCarrierNameOk returns a tuple with the CarrierName field value
// and a boolean to check if the value has been set.
func (o *AirlineAdvice) GetCarrierNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierName, true
}

// SetCarrierName sets field value
func (o *AirlineAdvice) SetCarrierName(v string) {
	o.CarrierName = v
}

// GetConjunctionTicketIndicator returns the ConjunctionTicketIndicator field value if set, zero value otherwise.
func (o *AirlineAdvice) GetConjunctionTicketIndicator() bool {
	if o == nil || IsNil(o.ConjunctionTicketIndicator) {
		var ret bool
		return ret
	}
	return *o.ConjunctionTicketIndicator
}

// GetConjunctionTicketIndicatorOk returns a tuple with the ConjunctionTicketIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AirlineAdvice) GetConjunctionTicketIndicatorOk() (*bool, bool) {
	if o == nil || IsNil(o.ConjunctionTicketIndicator) {
		return nil, false
	}
	return o.ConjunctionTicketIndicator, true
}

// HasConjunctionTicketIndicator returns a boolean if a field has been set.
func (o *AirlineAdvice) HasConjunctionTicketIndicator() bool {
	if o != nil && !IsNil(o.ConjunctionTicketIndicator) {
		return true
	}

	return false
}

// SetConjunctionTicketIndicator gets a reference to the given bool and assigns it to the ConjunctionTicketIndicator field.
func (o *AirlineAdvice) SetConjunctionTicketIndicator(v bool) {
	o.ConjunctionTicketIndicator = &v
}

// GetEticketIndicator returns the EticketIndicator field value if set, zero value otherwise.
func (o *AirlineAdvice) GetEticketIndicator() bool {
	if o == nil || IsNil(o.EticketIndicator) {
		var ret bool
		return ret
	}
	return *o.EticketIndicator
}

// GetEticketIndicatorOk returns a tuple with the EticketIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AirlineAdvice) GetEticketIndicatorOk() (*bool, bool) {
	if o == nil || IsNil(o.EticketIndicator) {
		return nil, false
	}
	return o.EticketIndicator, true
}

// HasEticketIndicator returns a boolean if a field has been set.
func (o *AirlineAdvice) HasEticketIndicator() bool {
	if o != nil && !IsNil(o.EticketIndicator) {
		return true
	}

	return false
}

// SetEticketIndicator gets a reference to the given bool and assigns it to the EticketIndicator field.
func (o *AirlineAdvice) SetEticketIndicator(v bool) {
	o.EticketIndicator = &v
}

// GetNoAirSegments returns the NoAirSegments field value if set, zero value otherwise.
func (o *AirlineAdvice) GetNoAirSegments() int32 {
	if o == nil || IsNil(o.NoAirSegments) {
		var ret int32
		return ret
	}
	return *o.NoAirSegments
}

// GetNoAirSegmentsOk returns a tuple with the NoAirSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AirlineAdvice) GetNoAirSegmentsOk() (*int32, bool) {
	if o == nil || IsNil(o.NoAirSegments) {
		return nil, false
	}
	return o.NoAirSegments, true
}

// HasNoAirSegments returns a boolean if a field has been set.
func (o *AirlineAdvice) HasNoAirSegments() bool {
	if o != nil && !IsNil(o.NoAirSegments) {
		return true
	}

	return false
}

// SetNoAirSegments gets a reference to the given int32 and assigns it to the NoAirSegments field.
func (o *AirlineAdvice) SetNoAirSegments(v int32) {
	o.NoAirSegments = &v
}

// GetNumberInParty returns the NumberInParty field value if set, zero value otherwise.
func (o *AirlineAdvice) GetNumberInParty() int32 {
	if o == nil || IsNil(o.NumberInParty) {
		var ret int32
		return ret
	}
	return *o.NumberInParty
}

// GetNumberInPartyOk returns a tuple with the NumberInParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AirlineAdvice) GetNumberInPartyOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberInParty) {
		return nil, false
	}
	return o.NumberInParty, true
}

// HasNumberInParty returns a boolean if a field has been set.
func (o *AirlineAdvice) HasNumberInParty() bool {
	if o != nil && !IsNil(o.NumberInParty) {
		return true
	}

	return false
}

// SetNumberInParty gets a reference to the given int32 and assigns it to the NumberInParty field.
func (o *AirlineAdvice) SetNumberInParty(v int32) {
	o.NumberInParty = &v
}

// GetOriginalTicketNo returns the OriginalTicketNo field value if set, zero value otherwise.
func (o *AirlineAdvice) GetOriginalTicketNo() string {
	if o == nil || IsNil(o.OriginalTicketNo) {
		var ret string
		return ret
	}
	return *o.OriginalTicketNo
}

// GetOriginalTicketNoOk returns a tuple with the OriginalTicketNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AirlineAdvice) GetOriginalTicketNoOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalTicketNo) {
		return nil, false
	}
	return o.OriginalTicketNo, true
}

// HasOriginalTicketNo returns a boolean if a field has been set.
func (o *AirlineAdvice) HasOriginalTicketNo() bool {
	if o != nil && !IsNil(o.OriginalTicketNo) {
		return true
	}

	return false
}

// SetOriginalTicketNo gets a reference to the given string and assigns it to the OriginalTicketNo field.
func (o *AirlineAdvice) SetOriginalTicketNo(v string) {
	o.OriginalTicketNo = &v
}

// GetPassengerName returns the PassengerName field value if set, zero value otherwise.
func (o *AirlineAdvice) GetPassengerName() string {
	if o == nil || IsNil(o.PassengerName) {
		var ret string
		return ret
	}
	return *o.PassengerName
}

// GetPassengerNameOk returns a tuple with the PassengerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AirlineAdvice) GetPassengerNameOk() (*string, bool) {
	if o == nil || IsNil(o.PassengerName) {
		return nil, false
	}
	return o.PassengerName, true
}

// HasPassengerName returns a boolean if a field has been set.
func (o *AirlineAdvice) HasPassengerName() bool {
	if o != nil && !IsNil(o.PassengerName) {
		return true
	}

	return false
}

// SetPassengerName gets a reference to the given string and assigns it to the PassengerName field.
func (o *AirlineAdvice) SetPassengerName(v string) {
	o.PassengerName = &v
}

// GetSegment1 returns the Segment1 field value
func (o *AirlineAdvice) GetSegment1() AirlineSegment {
	if o == nil {
		var ret AirlineSegment
		return ret
	}

	return o.Segment1
}

// GetSegment1Ok returns a tuple with the Segment1 field value
// and a boolean to check if the value has been set.
func (o *AirlineAdvice) GetSegment1Ok() (*AirlineSegment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Segment1, true
}

// SetSegment1 sets field value
func (o *AirlineAdvice) SetSegment1(v AirlineSegment) {
	o.Segment1 = v
}

// GetSegment2 returns the Segment2 field value if set, zero value otherwise.
func (o *AirlineAdvice) GetSegment2() AirlineSegment {
	if o == nil || IsNil(o.Segment2) {
		var ret AirlineSegment
		return ret
	}
	return *o.Segment2
}

// GetSegment2Ok returns a tuple with the Segment2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AirlineAdvice) GetSegment2Ok() (*AirlineSegment, bool) {
	if o == nil || IsNil(o.Segment2) {
		return nil, false
	}
	return o.Segment2, true
}

// HasSegment2 returns a boolean if a field has been set.
func (o *AirlineAdvice) HasSegment2() bool {
	if o != nil && !IsNil(o.Segment2) {
		return true
	}

	return false
}

// SetSegment2 gets a reference to the given AirlineSegment and assigns it to the Segment2 field.
func (o *AirlineAdvice) SetSegment2(v AirlineSegment) {
	o.Segment2 = &v
}

// GetSegment3 returns the Segment3 field value if set, zero value otherwise.
func (o *AirlineAdvice) GetSegment3() AirlineSegment {
	if o == nil || IsNil(o.Segment3) {
		var ret AirlineSegment
		return ret
	}
	return *o.Segment3
}

// GetSegment3Ok returns a tuple with the Segment3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AirlineAdvice) GetSegment3Ok() (*AirlineSegment, bool) {
	if o == nil || IsNil(o.Segment3) {
		return nil, false
	}
	return o.Segment3, true
}

// HasSegment3 returns a boolean if a field has been set.
func (o *AirlineAdvice) HasSegment3() bool {
	if o != nil && !IsNil(o.Segment3) {
		return true
	}

	return false
}

// SetSegment3 gets a reference to the given AirlineSegment and assigns it to the Segment3 field.
func (o *AirlineAdvice) SetSegment3(v AirlineSegment) {
	o.Segment3 = &v
}

// GetSegment4 returns the Segment4 field value if set, zero value otherwise.
func (o *AirlineAdvice) GetSegment4() AirlineSegment {
	if o == nil || IsNil(o.Segment4) {
		var ret AirlineSegment
		return ret
	}
	return *o.Segment4
}

// GetSegment4Ok returns a tuple with the Segment4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AirlineAdvice) GetSegment4Ok() (*AirlineSegment, bool) {
	if o == nil || IsNil(o.Segment4) {
		return nil, false
	}
	return o.Segment4, true
}

// HasSegment4 returns a boolean if a field has been set.
func (o *AirlineAdvice) HasSegment4() bool {
	if o != nil && !IsNil(o.Segment4) {
		return true
	}

	return false
}

// SetSegment4 gets a reference to the given AirlineSegment and assigns it to the Segment4 field.
func (o *AirlineAdvice) SetSegment4(v AirlineSegment) {
	o.Segment4 = &v
}

// GetTicketIssueCity returns the TicketIssueCity field value
func (o *AirlineAdvice) GetTicketIssueCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TicketIssueCity
}

// GetTicketIssueCityOk returns a tuple with the TicketIssueCity field value
// and a boolean to check if the value has been set.
func (o *AirlineAdvice) GetTicketIssueCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TicketIssueCity, true
}

// SetTicketIssueCity sets field value
func (o *AirlineAdvice) SetTicketIssueCity(v string) {
	o.TicketIssueCity = v
}

// GetTicketIssueDate returns the TicketIssueDate field value
func (o *AirlineAdvice) GetTicketIssueDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TicketIssueDate
}

// GetTicketIssueDateOk returns a tuple with the TicketIssueDate field value
// and a boolean to check if the value has been set.
func (o *AirlineAdvice) GetTicketIssueDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TicketIssueDate, true
}

// SetTicketIssueDate sets field value
func (o *AirlineAdvice) SetTicketIssueDate(v string) {
	o.TicketIssueDate = v
}

// GetTicketIssueName returns the TicketIssueName field value
func (o *AirlineAdvice) GetTicketIssueName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TicketIssueName
}

// GetTicketIssueNameOk returns a tuple with the TicketIssueName field value
// and a boolean to check if the value has been set.
func (o *AirlineAdvice) GetTicketIssueNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TicketIssueName, true
}

// SetTicketIssueName sets field value
func (o *AirlineAdvice) SetTicketIssueName(v string) {
	o.TicketIssueName = v
}

// GetTicketNo returns the TicketNo field value
func (o *AirlineAdvice) GetTicketNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TicketNo
}

// GetTicketNoOk returns a tuple with the TicketNo field value
// and a boolean to check if the value has been set.
func (o *AirlineAdvice) GetTicketNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TicketNo, true
}

// SetTicketNo sets field value
func (o *AirlineAdvice) SetTicketNo(v string) {
	o.TicketNo = v
}

// GetTransactionType returns the TransactionType field value
func (o *AirlineAdvice) GetTransactionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value
// and a boolean to check if the value has been set.
func (o *AirlineAdvice) GetTransactionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionType, true
}

// SetTransactionType sets field value
func (o *AirlineAdvice) SetTransactionType(v string) {
	o.TransactionType = v
}

func (o AirlineAdvice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AirlineAdvice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["carrier_name"] = o.CarrierName
	if !IsNil(o.ConjunctionTicketIndicator) {
		toSerialize["conjunction_ticket_indicator"] = o.ConjunctionTicketIndicator
	}
	if !IsNil(o.EticketIndicator) {
		toSerialize["eticket_indicator"] = o.EticketIndicator
	}
	if !IsNil(o.NoAirSegments) {
		toSerialize["no_air_segments"] = o.NoAirSegments
	}
	if !IsNil(o.NumberInParty) {
		toSerialize["number_in_party"] = o.NumberInParty
	}
	if !IsNil(o.OriginalTicketNo) {
		toSerialize["original_ticket_no"] = o.OriginalTicketNo
	}
	if !IsNil(o.PassengerName) {
		toSerialize["passenger_name"] = o.PassengerName
	}
	toSerialize["segment1"] = o.Segment1
	if !IsNil(o.Segment2) {
		toSerialize["segment2"] = o.Segment2
	}
	if !IsNil(o.Segment3) {
		toSerialize["segment3"] = o.Segment3
	}
	if !IsNil(o.Segment4) {
		toSerialize["segment4"] = o.Segment4
	}
	toSerialize["ticket_issue_city"] = o.TicketIssueCity
	toSerialize["ticket_issue_date"] = o.TicketIssueDate
	toSerialize["ticket_issue_name"] = o.TicketIssueName
	toSerialize["ticket_no"] = o.TicketNo
	toSerialize["transaction_type"] = o.TransactionType
	return toSerialize, nil
}

func (o *AirlineAdvice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"carrier_name",
		"segment1",
		"ticket_issue_city",
		"ticket_issue_date",
		"ticket_issue_name",
		"ticket_no",
		"transaction_type",
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

	varAirlineAdvice := _AirlineAdvice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAirlineAdvice)

	if err != nil {
		return err
	}

	*o = AirlineAdvice(varAirlineAdvice)

	return err
}

type NullableAirlineAdvice struct {
	value *AirlineAdvice
	isSet bool
}

func (v NullableAirlineAdvice) Get() *AirlineAdvice {
	return v.value
}

func (v *NullableAirlineAdvice) Set(val *AirlineAdvice) {
	v.value = val
	v.isSet = true
}

func (v NullableAirlineAdvice) IsSet() bool {
	return v.isSet
}

func (v *NullableAirlineAdvice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAirlineAdvice(val *AirlineAdvice) *NullableAirlineAdvice {
	return &NullableAirlineAdvice{value: val, isSet: true}
}

func (v NullableAirlineAdvice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAirlineAdvice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
