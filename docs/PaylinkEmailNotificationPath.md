# PaylinkEmailNotificationPath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bcc** | Pointer to **[]string** |  | [optional] 
**Cc** | Pointer to **[]string** |  | [optional] 
**ReplyTo** | Pointer to **[]string** |  | [optional] 
**Template** | Pointer to **string** | An optional template name to use a template other than the default. | [optional] 
**To** | **[]string** |  | 

## Methods

### NewPaylinkEmailNotificationPath

`func NewPaylinkEmailNotificationPath(to []string, ) *PaylinkEmailNotificationPath`

NewPaylinkEmailNotificationPath instantiates a new PaylinkEmailNotificationPath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkEmailNotificationPathWithDefaults

`func NewPaylinkEmailNotificationPathWithDefaults() *PaylinkEmailNotificationPath`

NewPaylinkEmailNotificationPathWithDefaults instantiates a new PaylinkEmailNotificationPath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBcc

`func (o *PaylinkEmailNotificationPath) GetBcc() []string`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *PaylinkEmailNotificationPath) GetBccOk() (*[]string, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *PaylinkEmailNotificationPath) SetBcc(v []string)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *PaylinkEmailNotificationPath) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetCc

`func (o *PaylinkEmailNotificationPath) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *PaylinkEmailNotificationPath) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *PaylinkEmailNotificationPath) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *PaylinkEmailNotificationPath) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetReplyTo

`func (o *PaylinkEmailNotificationPath) GetReplyTo() []string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *PaylinkEmailNotificationPath) GetReplyToOk() (*[]string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *PaylinkEmailNotificationPath) SetReplyTo(v []string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *PaylinkEmailNotificationPath) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetTemplate

`func (o *PaylinkEmailNotificationPath) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PaylinkEmailNotificationPath) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PaylinkEmailNotificationPath) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PaylinkEmailNotificationPath) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTo

`func (o *PaylinkEmailNotificationPath) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PaylinkEmailNotificationPath) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PaylinkEmailNotificationPath) SetTo(v []string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


