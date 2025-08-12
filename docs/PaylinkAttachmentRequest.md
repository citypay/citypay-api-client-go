# PaylinkAttachmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | base64 encoding of the file if less than 32kb in size. | [optional] 
**Filename** | **string** | The name of the attachment normally taken from the filename. You should not include the filename path as appropriate. | 
**MimeType** | **string** | The mime type of the attachment as defined in [RFC 9110](https://www.rfc-editor.org/rfc/rfc9110.html). Currently only &#x60;application/pdf&#x60; is supported. | 
**Name** | Pointer to **string** | A name for the file to identify it to the card holder when it is displayed in the payment form. For example Invoice, Statement. | [optional] 
**Retention** | Pointer to **int32** | The retention period in days of the attachment. Defaults to 180 days. | [optional] 

## Methods

### NewPaylinkAttachmentRequest

`func NewPaylinkAttachmentRequest(filename string, mimeType string, ) *PaylinkAttachmentRequest`

NewPaylinkAttachmentRequest instantiates a new PaylinkAttachmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkAttachmentRequestWithDefaults

`func NewPaylinkAttachmentRequestWithDefaults() *PaylinkAttachmentRequest`

NewPaylinkAttachmentRequestWithDefaults instantiates a new PaylinkAttachmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaylinkAttachmentRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaylinkAttachmentRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaylinkAttachmentRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *PaylinkAttachmentRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFilename

`func (o *PaylinkAttachmentRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *PaylinkAttachmentRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *PaylinkAttachmentRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetMimeType

`func (o *PaylinkAttachmentRequest) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *PaylinkAttachmentRequest) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *PaylinkAttachmentRequest) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetName

`func (o *PaylinkAttachmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaylinkAttachmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaylinkAttachmentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaylinkAttachmentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetention

`func (o *PaylinkAttachmentRequest) GetRetention() int32`

GetRetention returns the Retention field if non-nil, zero value otherwise.

### GetRetentionOk

`func (o *PaylinkAttachmentRequest) GetRetentionOk() (*int32, bool)`

GetRetentionOk returns a tuple with the Retention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetention

`func (o *PaylinkAttachmentRequest) SetRetention(v int32)`

SetRetention sets Retention field to given value.

### HasRetention

`func (o *PaylinkAttachmentRequest) HasRetention() bool`

HasRetention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


