# PaylinkResendNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **bool** | Resends the bill payment email provided on the Paylink create token notification path. Emails can be sent up to 5 times per token. | [optional] 
**Sms** | Pointer to **bool** | Resends the bill payment SMS provided on the Paylink create token notification path. An SMS cannot be resent if it was previously sent less than 1 minute ago. There is a limit of 3 retries per token.  | [optional] 

## Methods

### NewPaylinkResendNotificationRequest

`func NewPaylinkResendNotificationRequest() *PaylinkResendNotificationRequest`

NewPaylinkResendNotificationRequest instantiates a new PaylinkResendNotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkResendNotificationRequestWithDefaults

`func NewPaylinkResendNotificationRequestWithDefaults() *PaylinkResendNotificationRequest`

NewPaylinkResendNotificationRequestWithDefaults instantiates a new PaylinkResendNotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *PaylinkResendNotificationRequest) GetEmail() bool`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PaylinkResendNotificationRequest) GetEmailOk() (*bool, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PaylinkResendNotificationRequest) SetEmail(v bool)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PaylinkResendNotificationRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSms

`func (o *PaylinkResendNotificationRequest) GetSms() bool`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *PaylinkResendNotificationRequest) GetSmsOk() (*bool, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *PaylinkResendNotificationRequest) SetSms(v bool)`

SetSms sets Sms field to given value.

### HasSms

`func (o *PaylinkResendNotificationRequest) HasSms() bool`

HasSms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


