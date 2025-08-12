# PaylinkCustomParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMetaData** | Pointer to **bool** | Determines if the custom parameter is stored as meta data against an authorisation allowing for that authorisation to be searched and queried by the result of this value. Defaults to false. | [optional] 
**EntryMode** | Pointer to **string** | The type of entry mode. A value of &#x60;pre&#x60; will pre-render the custom parameter before the payment screen. Any other value will result in the custom parameter being displayed on the payment screen. | [optional] 
**FieldType** | Pointer to **string** | The type of html field, defaulting to &#x60;text&#x60;. Options are:   - &#x60;dob&#x60;      A date of birth field as a series of select list entries  - &#x60;text&#x60;     Allows the user to enter any text.  - &#x60;password&#x60; A field where the characters are masked to protect the input, typically used for passwords.  - &#x60;email&#x60;    Used for input fields that should contain an email address.  - &#x60;number&#x60;   For numeric input, can include controls for incrementing or decrementing the number.  - &#x60;tel&#x60;      For telephone numbers.  - &#x60;url&#x60;      A text field for entering a URL.  - &#x60;hidden&#x60;   Not visible to the user, but its value is sent when the form is submitted.  - &#x60;checkbox&#x60; A check box allowing single values to be selected/deselected.  - &#x60;radio&#x60;    Allows the user to select one of a limited number of choices.  - &#x60;select&#x60;   Renders as select items  Select Options:  Select options are constructed by providing a list of values in the value custom parameter field. Each value is delimited by a pipe character &#x60;|&#x60;. Value items can also be delimited with &#x60;:&#x60; as a value label pair.  For instance, a sports club requires identifying it&#39;s age group for membership entry:  &lt;CodeGroup title&#x3D;\&quot;Select Examples\&quot; label&#x3D;\&quot;ProcessBatchRequest\&quot;&gt;    &#x60;&#x60;&#x60;json {{ title: &#39;Basic Values&#39; }}      { \&quot;label\&quot; : \&quot;Age Group\&quot;,        \&quot;fieldType\&quot;: \&quot;select\&quot;,        \&quot;value\&quot; : \&quot;Under 18|18-30|30-50|50+\&quot; }...       &lt;select&gt;          &lt;option value&#x3D;\&quot;Under 18\&quot;&gt;Under 18&lt;/option&gt;          &lt;option value&#x3D;\&quot;18-30\&quot;&gt;18-30&lt;/option&gt;          &lt;option value&#x3D;\&quot;30-50\&quot;&gt;30-50&lt;/option&gt;          &lt;option value&#x3D;\&quot;50+\&quot;&gt;50+&lt;/option&gt;      &lt;/select&gt;    &#x60;&#x60;&#x60;    &#x60;&#x60;&#x60;json {{ title: &#39;Label and Values&#39; }}      { \&quot;label\&quot; : \&quot;Age Group\&quot;,        \&quot;fieldType\&quot;: \&quot;select\&quot;,        \&quot;value\&quot; : \&quot;0:Under 18|1:18-30|2:30-50|3:50+\&quot; }...       &lt;select&gt;          &lt;option value&#x3D;\&quot;0\&quot;&gt;Under 18&lt;/option&gt;          &lt;option value&#x3D;\&quot;1\&quot;&gt;18-30&lt;/option&gt;          &lt;option value&#x3D;\&quot;2\&quot;&gt;30-50&lt;/option&gt;          &lt;option value&#x3D;\&quot;3\&quot;&gt;50+&lt;/option&gt;      &lt;/select&gt;    &#x60;&#x60;&#x60; &lt;/CodeGroup&gt;  Fields may be requested as optional. If a select is required to be optional, provide a value such as &#x60;:Select an Option|options...&#x60; at the front of the list.  | [optional] 
**Group** | Pointer to **string** | A value which groups items for layout. The value should be a string title for rendering such as \&quot;Your Account Info\&quot;. If no value is provided, the parameter is added to a default parameter group. Group names are ordered alphabetically when rendered. | [optional] 
**Label** | Pointer to **string** | A label to show alongside the input. If this value is not supplied, the name value will be used. | [optional] 
**Locked** | Pointer to **bool** | States whether the field is locked, preventing entry or amendment by the person completing the form. | [optional] 
**Name** | **string** | Refers to the rendered HTML form element name. The value of this field is used in the postback and redirect dataset. | 
**Order** | Pointer to **int32** | A value which allows you to order the position of elements in a grouping. Values will order in ascending order. Negative values are possible. | [optional] 
**Pattern** | Pointer to **string** | A string value which specifies the validation logic of the form element, for example a value of QA[0-9]{3,4} will require a value such as QA221 or QA4433. | [optional] 
**Placeholder** | Pointer to **string** | A value to set as the placeholder attribute which will render in the browser. | [optional] 
**Required** | Pointer to **bool** | A boolean value that states whether the field is required or optional. When an element is required, validation will be performed on the end user&#39;s input form. | [optional] 
**Value** | Pointer to **string** | An initial value for the parameter as it appears on the Form. If your parameter is hidden, the value will be required. | [optional] 

## Methods

### NewPaylinkCustomParam

`func NewPaylinkCustomParam(name string, ) *PaylinkCustomParam`

NewPaylinkCustomParam instantiates a new PaylinkCustomParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkCustomParamWithDefaults

`func NewPaylinkCustomParamWithDefaults() *PaylinkCustomParam`

NewPaylinkCustomParamWithDefaults instantiates a new PaylinkCustomParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMetaData

`func (o *PaylinkCustomParam) GetAuthMetaData() bool`

GetAuthMetaData returns the AuthMetaData field if non-nil, zero value otherwise.

### GetAuthMetaDataOk

`func (o *PaylinkCustomParam) GetAuthMetaDataOk() (*bool, bool)`

GetAuthMetaDataOk returns a tuple with the AuthMetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMetaData

`func (o *PaylinkCustomParam) SetAuthMetaData(v bool)`

SetAuthMetaData sets AuthMetaData field to given value.

### HasAuthMetaData

`func (o *PaylinkCustomParam) HasAuthMetaData() bool`

HasAuthMetaData returns a boolean if a field has been set.

### GetEntryMode

`func (o *PaylinkCustomParam) GetEntryMode() string`

GetEntryMode returns the EntryMode field if non-nil, zero value otherwise.

### GetEntryModeOk

`func (o *PaylinkCustomParam) GetEntryModeOk() (*string, bool)`

GetEntryModeOk returns a tuple with the EntryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryMode

`func (o *PaylinkCustomParam) SetEntryMode(v string)`

SetEntryMode sets EntryMode field to given value.

### HasEntryMode

`func (o *PaylinkCustomParam) HasEntryMode() bool`

HasEntryMode returns a boolean if a field has been set.

### GetFieldType

`func (o *PaylinkCustomParam) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *PaylinkCustomParam) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *PaylinkCustomParam) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *PaylinkCustomParam) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetGroup

`func (o *PaylinkCustomParam) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PaylinkCustomParam) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PaylinkCustomParam) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PaylinkCustomParam) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetLabel

`func (o *PaylinkCustomParam) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PaylinkCustomParam) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PaylinkCustomParam) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PaylinkCustomParam) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLocked

`func (o *PaylinkCustomParam) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *PaylinkCustomParam) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *PaylinkCustomParam) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *PaylinkCustomParam) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetName

`func (o *PaylinkCustomParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaylinkCustomParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaylinkCustomParam) SetName(v string)`

SetName sets Name field to given value.


### GetOrder

`func (o *PaylinkCustomParam) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PaylinkCustomParam) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PaylinkCustomParam) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PaylinkCustomParam) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPattern

`func (o *PaylinkCustomParam) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *PaylinkCustomParam) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *PaylinkCustomParam) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *PaylinkCustomParam) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetPlaceholder

`func (o *PaylinkCustomParam) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *PaylinkCustomParam) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *PaylinkCustomParam) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *PaylinkCustomParam) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetRequired

`func (o *PaylinkCustomParam) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *PaylinkCustomParam) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *PaylinkCustomParam) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *PaylinkCustomParam) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetValue

`func (o *PaylinkCustomParam) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PaylinkCustomParam) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PaylinkCustomParam) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PaylinkCustomParam) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


