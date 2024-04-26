# PaylinkAttachmentResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the attachment. | 
**Result** | **string** | The result of an uploaded attachment such as &#x60;OK&#x60; or &#x60;UPLOAD&#x60;. | 
**Url** | Pointer to **string** | If the attachment is to be uploaded, a URL that can be used for Multipart upload of the attachment. | [optional] 

## Methods

### NewPaylinkAttachmentResult

`func NewPaylinkAttachmentResult(name string, result string, ) *PaylinkAttachmentResult`

NewPaylinkAttachmentResult instantiates a new PaylinkAttachmentResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkAttachmentResultWithDefaults

`func NewPaylinkAttachmentResultWithDefaults() *PaylinkAttachmentResult`

NewPaylinkAttachmentResultWithDefaults instantiates a new PaylinkAttachmentResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PaylinkAttachmentResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaylinkAttachmentResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaylinkAttachmentResult) SetName(v string)`

SetName sets Name field to given value.


### GetResult

`func (o *PaylinkAttachmentResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PaylinkAttachmentResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PaylinkAttachmentResult) SetResult(v string)`

SetResult sets Result field to given value.


### GetUrl

`func (o *PaylinkAttachmentResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PaylinkAttachmentResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PaylinkAttachmentResult) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PaylinkAttachmentResult) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


