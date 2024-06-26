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

// checks if the BatchReportResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchReportResponseModel{}

// BatchReportResponseModel struct for BatchReportResponseModel
type BatchReportResponseModel struct {
	// The total amount that the batch contains.
	Amount int32 `json:"amount"`
	// The date and time of the batch in ISO-8601 format.
	BatchDate string `json:"batch_date"`
	// The batch id specified in the batch processing request.
	BatchId int32 `json:"batch_id"`
	// The status of the batch. Possible values are:   - CANCELLED. The file has been cancelled by an administrator or server process.  - COMPLETE. The file has passed through the processing cycle and is determined as being complete further information should be obtained on the results of the processing - ERROR_IN_PROCESSING. Errors have occurred in the processing that has deemed that processing can not continue. - INITIALISED. The file has been initialised and no action has yet been performed - LOCKED. The file has been locked for processing - QUEUED. The file has been queued for processing yet no processing has yet been performed - UNKNOWN. The file is of an unknown status, that is the file can either not be determined by the information requested of the file has not yet been received.
	BatchStatus string `json:"batch_status"`
	// The batch account id that the batch was processed with.
	ClientAccountId string                        `json:"client_account_id"`
	Transactions    []BatchTransactionResultModel `json:"transactions"`
}

type _BatchReportResponseModel BatchReportResponseModel

// NewBatchReportResponseModel instantiates a new BatchReportResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchReportResponseModel(amount int32, batchDate string, batchId int32, batchStatus string, clientAccountId string, transactions []BatchTransactionResultModel) *BatchReportResponseModel {
	this := BatchReportResponseModel{}
	this.Amount = amount
	this.BatchDate = batchDate
	this.BatchId = batchId
	this.BatchStatus = batchStatus
	this.ClientAccountId = clientAccountId
	this.Transactions = transactions
	return &this
}

// NewBatchReportResponseModelWithDefaults instantiates a new BatchReportResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchReportResponseModelWithDefaults() *BatchReportResponseModel {
	this := BatchReportResponseModel{}
	return &this
}

// GetAmount returns the Amount field value
func (o *BatchReportResponseModel) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *BatchReportResponseModel) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *BatchReportResponseModel) SetAmount(v int32) {
	o.Amount = v
}

// GetBatchDate returns the BatchDate field value
func (o *BatchReportResponseModel) GetBatchDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BatchDate
}

// GetBatchDateOk returns a tuple with the BatchDate field value
// and a boolean to check if the value has been set.
func (o *BatchReportResponseModel) GetBatchDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BatchDate, true
}

// SetBatchDate sets field value
func (o *BatchReportResponseModel) SetBatchDate(v string) {
	o.BatchDate = v
}

// GetBatchId returns the BatchId field value
func (o *BatchReportResponseModel) GetBatchId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value
// and a boolean to check if the value has been set.
func (o *BatchReportResponseModel) GetBatchIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BatchId, true
}

// SetBatchId sets field value
func (o *BatchReportResponseModel) SetBatchId(v int32) {
	o.BatchId = v
}

// GetBatchStatus returns the BatchStatus field value
func (o *BatchReportResponseModel) GetBatchStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BatchStatus
}

// GetBatchStatusOk returns a tuple with the BatchStatus field value
// and a boolean to check if the value has been set.
func (o *BatchReportResponseModel) GetBatchStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BatchStatus, true
}

// SetBatchStatus sets field value
func (o *BatchReportResponseModel) SetBatchStatus(v string) {
	o.BatchStatus = v
}

// GetClientAccountId returns the ClientAccountId field value
func (o *BatchReportResponseModel) GetClientAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientAccountId
}

// GetClientAccountIdOk returns a tuple with the ClientAccountId field value
// and a boolean to check if the value has been set.
func (o *BatchReportResponseModel) GetClientAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientAccountId, true
}

// SetClientAccountId sets field value
func (o *BatchReportResponseModel) SetClientAccountId(v string) {
	o.ClientAccountId = v
}

// GetTransactions returns the Transactions field value
func (o *BatchReportResponseModel) GetTransactions() []BatchTransactionResultModel {
	if o == nil {
		var ret []BatchTransactionResultModel
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *BatchReportResponseModel) GetTransactionsOk() ([]BatchTransactionResultModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transactions, true
}

// SetTransactions sets field value
func (o *BatchReportResponseModel) SetTransactions(v []BatchTransactionResultModel) {
	o.Transactions = v
}

func (o BatchReportResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchReportResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["batch_date"] = o.BatchDate
	toSerialize["batch_id"] = o.BatchId
	toSerialize["batch_status"] = o.BatchStatus
	toSerialize["client_account_id"] = o.ClientAccountId
	toSerialize["transactions"] = o.Transactions
	return toSerialize, nil
}

func (o *BatchReportResponseModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"batch_date",
		"batch_id",
		"batch_status",
		"client_account_id",
		"transactions",
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

	varBatchReportResponseModel := _BatchReportResponseModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchReportResponseModel)

	if err != nil {
		return err
	}

	*o = BatchReportResponseModel(varBatchReportResponseModel)

	return err
}

type NullableBatchReportResponseModel struct {
	value *BatchReportResponseModel
	isSet bool
}

func (v NullableBatchReportResponseModel) Get() *BatchReportResponseModel {
	return v.value
}

func (v *NullableBatchReportResponseModel) Set(val *BatchReportResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchReportResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchReportResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchReportResponseModel(val *BatchReportResponseModel) *NullableBatchReportResponseModel {
	return &NullableBatchReportResponseModel{value: val, isSet: true}
}

func (v NullableBatchReportResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchReportResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
