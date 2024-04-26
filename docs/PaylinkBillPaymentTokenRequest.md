# PaylinkBillPaymentTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addressee** | Pointer to **string** | Who the bill payment request intended for. This should be a readable name such as a person or company. | [optional] 
**Attachments** | Pointer to [**[]PaylinkAttachmentRequest**](PaylinkAttachmentRequest.md) |  | [optional] 
**Descriptor** | Pointer to **string** | A descriptor for the bill payment used to describe what the payment request is for for instance \&quot;Invoice\&quot;.  The descriptor can be used as descriptive text on emails or the payment page. For instance an invoice may have a button saying \&quot;View Invoice\&quot; or an email may say \&quot;to pay your Invoice online\&quot;.  | [optional] 
**Due** | Pointer to **string** | A date that the invoice is due. This can be displayed on the payment page. | [optional] 
**EmailNotificationPath** | Pointer to [**PaylinkEmailNotificationPath**](PaylinkEmailNotificationPath.md) |  | [optional] 
**Memo** | Pointer to **string** | A memo that can be added to the payment page and email to provide to the customer. | [optional] 
**Request** | [**PaylinkTokenRequestModel**](PaylinkTokenRequestModel.md) |  | 
**SmsNotificationPath** | Pointer to [**PaylinkSMSNotificationPath**](PaylinkSMSNotificationPath.md) |  | [optional] 

## Methods

### NewPaylinkBillPaymentTokenRequest

`func NewPaylinkBillPaymentTokenRequest(request PaylinkTokenRequestModel, ) *PaylinkBillPaymentTokenRequest`

NewPaylinkBillPaymentTokenRequest instantiates a new PaylinkBillPaymentTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkBillPaymentTokenRequestWithDefaults

`func NewPaylinkBillPaymentTokenRequestWithDefaults() *PaylinkBillPaymentTokenRequest`

NewPaylinkBillPaymentTokenRequestWithDefaults instantiates a new PaylinkBillPaymentTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressee

`func (o *PaylinkBillPaymentTokenRequest) GetAddressee() string`

GetAddressee returns the Addressee field if non-nil, zero value otherwise.

### GetAddresseeOk

`func (o *PaylinkBillPaymentTokenRequest) GetAddresseeOk() (*string, bool)`

GetAddresseeOk returns a tuple with the Addressee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressee

`func (o *PaylinkBillPaymentTokenRequest) SetAddressee(v string)`

SetAddressee sets Addressee field to given value.

### HasAddressee

`func (o *PaylinkBillPaymentTokenRequest) HasAddressee() bool`

HasAddressee returns a boolean if a field has been set.

### GetAttachments

`func (o *PaylinkBillPaymentTokenRequest) GetAttachments() []PaylinkAttachmentRequest`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *PaylinkBillPaymentTokenRequest) GetAttachmentsOk() (*[]PaylinkAttachmentRequest, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *PaylinkBillPaymentTokenRequest) SetAttachments(v []PaylinkAttachmentRequest)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *PaylinkBillPaymentTokenRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetDescriptor

`func (o *PaylinkBillPaymentTokenRequest) GetDescriptor() string`

GetDescriptor returns the Descriptor field if non-nil, zero value otherwise.

### GetDescriptorOk

`func (o *PaylinkBillPaymentTokenRequest) GetDescriptorOk() (*string, bool)`

GetDescriptorOk returns a tuple with the Descriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptor

`func (o *PaylinkBillPaymentTokenRequest) SetDescriptor(v string)`

SetDescriptor sets Descriptor field to given value.

### HasDescriptor

`func (o *PaylinkBillPaymentTokenRequest) HasDescriptor() bool`

HasDescriptor returns a boolean if a field has been set.

### GetDue

`func (o *PaylinkBillPaymentTokenRequest) GetDue() string`

GetDue returns the Due field if non-nil, zero value otherwise.

### GetDueOk

`func (o *PaylinkBillPaymentTokenRequest) GetDueOk() (*string, bool)`

GetDueOk returns a tuple with the Due field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDue

`func (o *PaylinkBillPaymentTokenRequest) SetDue(v string)`

SetDue sets Due field to given value.

### HasDue

`func (o *PaylinkBillPaymentTokenRequest) HasDue() bool`

HasDue returns a boolean if a field has been set.

### GetEmailNotificationPath

`func (o *PaylinkBillPaymentTokenRequest) GetEmailNotificationPath() PaylinkEmailNotificationPath`

GetEmailNotificationPath returns the EmailNotificationPath field if non-nil, zero value otherwise.

### GetEmailNotificationPathOk

`func (o *PaylinkBillPaymentTokenRequest) GetEmailNotificationPathOk() (*PaylinkEmailNotificationPath, bool)`

GetEmailNotificationPathOk returns a tuple with the EmailNotificationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailNotificationPath

`func (o *PaylinkBillPaymentTokenRequest) SetEmailNotificationPath(v PaylinkEmailNotificationPath)`

SetEmailNotificationPath sets EmailNotificationPath field to given value.

### HasEmailNotificationPath

`func (o *PaylinkBillPaymentTokenRequest) HasEmailNotificationPath() bool`

HasEmailNotificationPath returns a boolean if a field has been set.

### GetMemo

`func (o *PaylinkBillPaymentTokenRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *PaylinkBillPaymentTokenRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *PaylinkBillPaymentTokenRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *PaylinkBillPaymentTokenRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetRequest

`func (o *PaylinkBillPaymentTokenRequest) GetRequest() PaylinkTokenRequestModel`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *PaylinkBillPaymentTokenRequest) GetRequestOk() (*PaylinkTokenRequestModel, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *PaylinkBillPaymentTokenRequest) SetRequest(v PaylinkTokenRequestModel)`

SetRequest sets Request field to given value.


### GetSmsNotificationPath

`func (o *PaylinkBillPaymentTokenRequest) GetSmsNotificationPath() PaylinkSMSNotificationPath`

GetSmsNotificationPath returns the SmsNotificationPath field if non-nil, zero value otherwise.

### GetSmsNotificationPathOk

`func (o *PaylinkBillPaymentTokenRequest) GetSmsNotificationPathOk() (*PaylinkSMSNotificationPath, bool)`

GetSmsNotificationPathOk returns a tuple with the SmsNotificationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNotificationPath

`func (o *PaylinkBillPaymentTokenRequest) SetSmsNotificationPath(v PaylinkSMSNotificationPath)`

SetSmsNotificationPath sets SmsNotificationPath field to given value.

### HasSmsNotificationPath

`func (o *PaylinkBillPaymentTokenRequest) HasSmsNotificationPath() bool`

HasSmsNotificationPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


