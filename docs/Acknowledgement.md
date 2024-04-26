# Acknowledgement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | A response code providing a result of the process. | [optional] 
**Context** | Pointer to **string** | A context id of the process used for referencing transactions through support. | [optional] 
**Identifier** | Pointer to **string** | An identifier if presented in the original request. | [optional] 
**Message** | Pointer to **string** | A response message providing a description of the result of the process. | [optional] 

## Methods

### NewAcknowledgement

`func NewAcknowledgement() *Acknowledgement`

NewAcknowledgement instantiates a new Acknowledgement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcknowledgementWithDefaults

`func NewAcknowledgementWithDefaults() *Acknowledgement`

NewAcknowledgementWithDefaults instantiates a new Acknowledgement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Acknowledgement) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Acknowledgement) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Acknowledgement) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Acknowledgement) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetContext

`func (o *Acknowledgement) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Acknowledgement) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Acknowledgement) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *Acknowledgement) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetIdentifier

`func (o *Acknowledgement) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Acknowledgement) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Acknowledgement) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Acknowledgement) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetMessage

`func (o *Acknowledgement) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Acknowledgement) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Acknowledgement) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Acknowledgement) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


