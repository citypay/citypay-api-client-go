# PaylinkTokenStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountPaid** | Pointer to **int32** | the amount that has been paid against the session. | [optional] 
**AuthCode** | Pointer to **string** | an authorisation code if the transaction was processed and isPaid is true. | [optional] 
**Card** | Pointer to **string** | a description of the card that was used for payment if paid. | [optional] 
**Created** | Pointer to **time.Time** | the date and time that the session was created. | [optional] 
**Datetime** | Pointer to **time.Time** | the date and time of the current status. | [optional] 
**Identifier** | Pointer to **string** | the merchant identifier, to help identifying the token. | [optional] 
**IsAttachment** | Pointer to **bool** | true if an attachment exists. | [optional] 
**IsCancelled** | Pointer to **bool** | true if the session was cancelled either by the user or by a system request. | [optional] 
**IsClosed** | Pointer to **bool** | true if the token has been closed. | [optional] 
**IsCustomerReceiptEmailSent** | Pointer to **bool** | true if a customer receipt has been sent. | [optional] 
**IsEmailSent** | Pointer to **bool** | true if an email was sent. | [optional] 
**IsExpired** | Pointer to **bool** | true if the session has expired. | [optional] 
**IsFormViewed** | Pointer to **bool** | true if the form was ever displayed to the addressee. | [optional] 
**IsMerchantNotificationEmailSent** | Pointer to **bool** | true if a merchant notification receipt was sent. | [optional] 
**IsOpenForPayment** | Pointer to **bool** | true if the session is still open for payment or false if it has been closed. | [optional] 
**IsPaid** | Pointer to **bool** | whether the session has been paid and therefore can be considered as complete. | [optional] 
**IsPaymentAttempted** | Pointer to **bool** | true if payment has been attempted. | [optional] 
**IsPostbackOk** | Pointer to **bool** | true if a post back was executed successfully. | [optional] 
**IsRequestChallenged** | Pointer to **bool** | true if the request has been challenged using 3-D Secure. | [optional] 
**IsSmsSent** | Pointer to **bool** | true if an SMS was sent. | [optional] 
**IsValidated** | Pointer to **bool** | whether the token generation was successfully validated. | [optional] 
**LastEventDateTime** | Pointer to **time.Time** | the date and time that the session last had an event actioned against it. | [optional] 
**LastPaymentResult** | Pointer to **string** | the result of the last payment if one exists. | [optional] 
**Mid** | Pointer to **int32** | identifies the merchant account. | [optional] 
**PaymentAttemptsCount** | Pointer to **int32** | the number of attempts made to pay. | [optional] 
**StateHistory** | Pointer to [**[]PaylinkStateEvent**](PaylinkStateEvent.md) |  | [optional] 
**Token** | Pointer to **string** | the token value which uniquely identifies the session. | [optional] 
**TransNo** | Pointer to **int32** | a transaction number if the transacstion was processed and isPaid is true. | [optional] 

## Methods

### NewPaylinkTokenStatus

`func NewPaylinkTokenStatus() *PaylinkTokenStatus`

NewPaylinkTokenStatus instantiates a new PaylinkTokenStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkTokenStatusWithDefaults

`func NewPaylinkTokenStatusWithDefaults() *PaylinkTokenStatus`

NewPaylinkTokenStatusWithDefaults instantiates a new PaylinkTokenStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountPaid

`func (o *PaylinkTokenStatus) GetAmountPaid() int32`

GetAmountPaid returns the AmountPaid field if non-nil, zero value otherwise.

### GetAmountPaidOk

`func (o *PaylinkTokenStatus) GetAmountPaidOk() (*int32, bool)`

GetAmountPaidOk returns a tuple with the AmountPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPaid

`func (o *PaylinkTokenStatus) SetAmountPaid(v int32)`

SetAmountPaid sets AmountPaid field to given value.

### HasAmountPaid

`func (o *PaylinkTokenStatus) HasAmountPaid() bool`

HasAmountPaid returns a boolean if a field has been set.

### GetAuthCode

`func (o *PaylinkTokenStatus) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *PaylinkTokenStatus) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *PaylinkTokenStatus) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *PaylinkTokenStatus) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetCard

`func (o *PaylinkTokenStatus) GetCard() string`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *PaylinkTokenStatus) GetCardOk() (*string, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *PaylinkTokenStatus) SetCard(v string)`

SetCard sets Card field to given value.

### HasCard

`func (o *PaylinkTokenStatus) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetCreated

`func (o *PaylinkTokenStatus) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PaylinkTokenStatus) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PaylinkTokenStatus) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PaylinkTokenStatus) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDatetime

`func (o *PaylinkTokenStatus) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *PaylinkTokenStatus) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *PaylinkTokenStatus) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *PaylinkTokenStatus) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetIdentifier

`func (o *PaylinkTokenStatus) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PaylinkTokenStatus) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PaylinkTokenStatus) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PaylinkTokenStatus) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetIsAttachment

`func (o *PaylinkTokenStatus) GetIsAttachment() bool`

GetIsAttachment returns the IsAttachment field if non-nil, zero value otherwise.

### GetIsAttachmentOk

`func (o *PaylinkTokenStatus) GetIsAttachmentOk() (*bool, bool)`

GetIsAttachmentOk returns a tuple with the IsAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAttachment

`func (o *PaylinkTokenStatus) SetIsAttachment(v bool)`

SetIsAttachment sets IsAttachment field to given value.

### HasIsAttachment

`func (o *PaylinkTokenStatus) HasIsAttachment() bool`

HasIsAttachment returns a boolean if a field has been set.

### GetIsCancelled

`func (o *PaylinkTokenStatus) GetIsCancelled() bool`

GetIsCancelled returns the IsCancelled field if non-nil, zero value otherwise.

### GetIsCancelledOk

`func (o *PaylinkTokenStatus) GetIsCancelledOk() (*bool, bool)`

GetIsCancelledOk returns a tuple with the IsCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCancelled

`func (o *PaylinkTokenStatus) SetIsCancelled(v bool)`

SetIsCancelled sets IsCancelled field to given value.

### HasIsCancelled

`func (o *PaylinkTokenStatus) HasIsCancelled() bool`

HasIsCancelled returns a boolean if a field has been set.

### GetIsClosed

`func (o *PaylinkTokenStatus) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *PaylinkTokenStatus) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *PaylinkTokenStatus) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *PaylinkTokenStatus) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### GetIsCustomerReceiptEmailSent

`func (o *PaylinkTokenStatus) GetIsCustomerReceiptEmailSent() bool`

GetIsCustomerReceiptEmailSent returns the IsCustomerReceiptEmailSent field if non-nil, zero value otherwise.

### GetIsCustomerReceiptEmailSentOk

`func (o *PaylinkTokenStatus) GetIsCustomerReceiptEmailSentOk() (*bool, bool)`

GetIsCustomerReceiptEmailSentOk returns a tuple with the IsCustomerReceiptEmailSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomerReceiptEmailSent

`func (o *PaylinkTokenStatus) SetIsCustomerReceiptEmailSent(v bool)`

SetIsCustomerReceiptEmailSent sets IsCustomerReceiptEmailSent field to given value.

### HasIsCustomerReceiptEmailSent

`func (o *PaylinkTokenStatus) HasIsCustomerReceiptEmailSent() bool`

HasIsCustomerReceiptEmailSent returns a boolean if a field has been set.

### GetIsEmailSent

`func (o *PaylinkTokenStatus) GetIsEmailSent() bool`

GetIsEmailSent returns the IsEmailSent field if non-nil, zero value otherwise.

### GetIsEmailSentOk

`func (o *PaylinkTokenStatus) GetIsEmailSentOk() (*bool, bool)`

GetIsEmailSentOk returns a tuple with the IsEmailSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEmailSent

`func (o *PaylinkTokenStatus) SetIsEmailSent(v bool)`

SetIsEmailSent sets IsEmailSent field to given value.

### HasIsEmailSent

`func (o *PaylinkTokenStatus) HasIsEmailSent() bool`

HasIsEmailSent returns a boolean if a field has been set.

### GetIsExpired

`func (o *PaylinkTokenStatus) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *PaylinkTokenStatus) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *PaylinkTokenStatus) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *PaylinkTokenStatus) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.

### GetIsFormViewed

`func (o *PaylinkTokenStatus) GetIsFormViewed() bool`

GetIsFormViewed returns the IsFormViewed field if non-nil, zero value otherwise.

### GetIsFormViewedOk

`func (o *PaylinkTokenStatus) GetIsFormViewedOk() (*bool, bool)`

GetIsFormViewedOk returns a tuple with the IsFormViewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFormViewed

`func (o *PaylinkTokenStatus) SetIsFormViewed(v bool)`

SetIsFormViewed sets IsFormViewed field to given value.

### HasIsFormViewed

`func (o *PaylinkTokenStatus) HasIsFormViewed() bool`

HasIsFormViewed returns a boolean if a field has been set.

### GetIsMerchantNotificationEmailSent

`func (o *PaylinkTokenStatus) GetIsMerchantNotificationEmailSent() bool`

GetIsMerchantNotificationEmailSent returns the IsMerchantNotificationEmailSent field if non-nil, zero value otherwise.

### GetIsMerchantNotificationEmailSentOk

`func (o *PaylinkTokenStatus) GetIsMerchantNotificationEmailSentOk() (*bool, bool)`

GetIsMerchantNotificationEmailSentOk returns a tuple with the IsMerchantNotificationEmailSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMerchantNotificationEmailSent

`func (o *PaylinkTokenStatus) SetIsMerchantNotificationEmailSent(v bool)`

SetIsMerchantNotificationEmailSent sets IsMerchantNotificationEmailSent field to given value.

### HasIsMerchantNotificationEmailSent

`func (o *PaylinkTokenStatus) HasIsMerchantNotificationEmailSent() bool`

HasIsMerchantNotificationEmailSent returns a boolean if a field has been set.

### GetIsOpenForPayment

`func (o *PaylinkTokenStatus) GetIsOpenForPayment() bool`

GetIsOpenForPayment returns the IsOpenForPayment field if non-nil, zero value otherwise.

### GetIsOpenForPaymentOk

`func (o *PaylinkTokenStatus) GetIsOpenForPaymentOk() (*bool, bool)`

GetIsOpenForPaymentOk returns a tuple with the IsOpenForPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpenForPayment

`func (o *PaylinkTokenStatus) SetIsOpenForPayment(v bool)`

SetIsOpenForPayment sets IsOpenForPayment field to given value.

### HasIsOpenForPayment

`func (o *PaylinkTokenStatus) HasIsOpenForPayment() bool`

HasIsOpenForPayment returns a boolean if a field has been set.

### GetIsPaid

`func (o *PaylinkTokenStatus) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *PaylinkTokenStatus) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *PaylinkTokenStatus) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.

### HasIsPaid

`func (o *PaylinkTokenStatus) HasIsPaid() bool`

HasIsPaid returns a boolean if a field has been set.

### GetIsPaymentAttempted

`func (o *PaylinkTokenStatus) GetIsPaymentAttempted() bool`

GetIsPaymentAttempted returns the IsPaymentAttempted field if non-nil, zero value otherwise.

### GetIsPaymentAttemptedOk

`func (o *PaylinkTokenStatus) GetIsPaymentAttemptedOk() (*bool, bool)`

GetIsPaymentAttemptedOk returns a tuple with the IsPaymentAttempted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaymentAttempted

`func (o *PaylinkTokenStatus) SetIsPaymentAttempted(v bool)`

SetIsPaymentAttempted sets IsPaymentAttempted field to given value.

### HasIsPaymentAttempted

`func (o *PaylinkTokenStatus) HasIsPaymentAttempted() bool`

HasIsPaymentAttempted returns a boolean if a field has been set.

### GetIsPostbackOk

`func (o *PaylinkTokenStatus) GetIsPostbackOk() bool`

GetIsPostbackOk returns the IsPostbackOk field if non-nil, zero value otherwise.

### GetIsPostbackOkOk

`func (o *PaylinkTokenStatus) GetIsPostbackOkOk() (*bool, bool)`

GetIsPostbackOkOk returns a tuple with the IsPostbackOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPostbackOk

`func (o *PaylinkTokenStatus) SetIsPostbackOk(v bool)`

SetIsPostbackOk sets IsPostbackOk field to given value.

### HasIsPostbackOk

`func (o *PaylinkTokenStatus) HasIsPostbackOk() bool`

HasIsPostbackOk returns a boolean if a field has been set.

### GetIsRequestChallenged

`func (o *PaylinkTokenStatus) GetIsRequestChallenged() bool`

GetIsRequestChallenged returns the IsRequestChallenged field if non-nil, zero value otherwise.

### GetIsRequestChallengedOk

`func (o *PaylinkTokenStatus) GetIsRequestChallengedOk() (*bool, bool)`

GetIsRequestChallengedOk returns a tuple with the IsRequestChallenged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequestChallenged

`func (o *PaylinkTokenStatus) SetIsRequestChallenged(v bool)`

SetIsRequestChallenged sets IsRequestChallenged field to given value.

### HasIsRequestChallenged

`func (o *PaylinkTokenStatus) HasIsRequestChallenged() bool`

HasIsRequestChallenged returns a boolean if a field has been set.

### GetIsSmsSent

`func (o *PaylinkTokenStatus) GetIsSmsSent() bool`

GetIsSmsSent returns the IsSmsSent field if non-nil, zero value otherwise.

### GetIsSmsSentOk

`func (o *PaylinkTokenStatus) GetIsSmsSentOk() (*bool, bool)`

GetIsSmsSentOk returns a tuple with the IsSmsSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSmsSent

`func (o *PaylinkTokenStatus) SetIsSmsSent(v bool)`

SetIsSmsSent sets IsSmsSent field to given value.

### HasIsSmsSent

`func (o *PaylinkTokenStatus) HasIsSmsSent() bool`

HasIsSmsSent returns a boolean if a field has been set.

### GetIsValidated

`func (o *PaylinkTokenStatus) GetIsValidated() bool`

GetIsValidated returns the IsValidated field if non-nil, zero value otherwise.

### GetIsValidatedOk

`func (o *PaylinkTokenStatus) GetIsValidatedOk() (*bool, bool)`

GetIsValidatedOk returns a tuple with the IsValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValidated

`func (o *PaylinkTokenStatus) SetIsValidated(v bool)`

SetIsValidated sets IsValidated field to given value.

### HasIsValidated

`func (o *PaylinkTokenStatus) HasIsValidated() bool`

HasIsValidated returns a boolean if a field has been set.

### GetLastEventDateTime

`func (o *PaylinkTokenStatus) GetLastEventDateTime() time.Time`

GetLastEventDateTime returns the LastEventDateTime field if non-nil, zero value otherwise.

### GetLastEventDateTimeOk

`func (o *PaylinkTokenStatus) GetLastEventDateTimeOk() (*time.Time, bool)`

GetLastEventDateTimeOk returns a tuple with the LastEventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventDateTime

`func (o *PaylinkTokenStatus) SetLastEventDateTime(v time.Time)`

SetLastEventDateTime sets LastEventDateTime field to given value.

### HasLastEventDateTime

`func (o *PaylinkTokenStatus) HasLastEventDateTime() bool`

HasLastEventDateTime returns a boolean if a field has been set.

### GetLastPaymentResult

`func (o *PaylinkTokenStatus) GetLastPaymentResult() string`

GetLastPaymentResult returns the LastPaymentResult field if non-nil, zero value otherwise.

### GetLastPaymentResultOk

`func (o *PaylinkTokenStatus) GetLastPaymentResultOk() (*string, bool)`

GetLastPaymentResultOk returns a tuple with the LastPaymentResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentResult

`func (o *PaylinkTokenStatus) SetLastPaymentResult(v string)`

SetLastPaymentResult sets LastPaymentResult field to given value.

### HasLastPaymentResult

`func (o *PaylinkTokenStatus) HasLastPaymentResult() bool`

HasLastPaymentResult returns a boolean if a field has been set.

### GetMid

`func (o *PaylinkTokenStatus) GetMid() int32`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *PaylinkTokenStatus) GetMidOk() (*int32, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *PaylinkTokenStatus) SetMid(v int32)`

SetMid sets Mid field to given value.

### HasMid

`func (o *PaylinkTokenStatus) HasMid() bool`

HasMid returns a boolean if a field has been set.

### GetPaymentAttemptsCount

`func (o *PaylinkTokenStatus) GetPaymentAttemptsCount() int32`

GetPaymentAttemptsCount returns the PaymentAttemptsCount field if non-nil, zero value otherwise.

### GetPaymentAttemptsCountOk

`func (o *PaylinkTokenStatus) GetPaymentAttemptsCountOk() (*int32, bool)`

GetPaymentAttemptsCountOk returns a tuple with the PaymentAttemptsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAttemptsCount

`func (o *PaylinkTokenStatus) SetPaymentAttemptsCount(v int32)`

SetPaymentAttemptsCount sets PaymentAttemptsCount field to given value.

### HasPaymentAttemptsCount

`func (o *PaylinkTokenStatus) HasPaymentAttemptsCount() bool`

HasPaymentAttemptsCount returns a boolean if a field has been set.

### GetStateHistory

`func (o *PaylinkTokenStatus) GetStateHistory() []PaylinkStateEvent`

GetStateHistory returns the StateHistory field if non-nil, zero value otherwise.

### GetStateHistoryOk

`func (o *PaylinkTokenStatus) GetStateHistoryOk() (*[]PaylinkStateEvent, bool)`

GetStateHistoryOk returns a tuple with the StateHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateHistory

`func (o *PaylinkTokenStatus) SetStateHistory(v []PaylinkStateEvent)`

SetStateHistory sets StateHistory field to given value.

### HasStateHistory

`func (o *PaylinkTokenStatus) HasStateHistory() bool`

HasStateHistory returns a boolean if a field has been set.

### GetToken

`func (o *PaylinkTokenStatus) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PaylinkTokenStatus) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PaylinkTokenStatus) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PaylinkTokenStatus) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTransNo

`func (o *PaylinkTokenStatus) GetTransNo() int32`

GetTransNo returns the TransNo field if non-nil, zero value otherwise.

### GetTransNoOk

`func (o *PaylinkTokenStatus) GetTransNoOk() (*int32, bool)`

GetTransNoOk returns a tuple with the TransNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransNo

`func (o *PaylinkTokenStatus) SetTransNo(v int32)`

SetTransNo sets TransNo field to given value.

### HasTransNo

`func (o *PaylinkTokenStatus) HasTransNo() bool`

HasTransNo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


