# PaylinkSMSNotificationPath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | Pointer to **string** | An optional template name to use a template other than the default. | [optional] 
**To** | **string** | The phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format to send the message to. | 

## Methods

### NewPaylinkSMSNotificationPath

`func NewPaylinkSMSNotificationPath(to string, ) *PaylinkSMSNotificationPath`

NewPaylinkSMSNotificationPath instantiates a new PaylinkSMSNotificationPath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkSMSNotificationPathWithDefaults

`func NewPaylinkSMSNotificationPathWithDefaults() *PaylinkSMSNotificationPath`

NewPaylinkSMSNotificationPathWithDefaults instantiates a new PaylinkSMSNotificationPath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *PaylinkSMSNotificationPath) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PaylinkSMSNotificationPath) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PaylinkSMSNotificationPath) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PaylinkSMSNotificationPath) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTo

`func (o *PaylinkSMSNotificationPath) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PaylinkSMSNotificationPath) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PaylinkSMSNotificationPath) SetTo(v string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


