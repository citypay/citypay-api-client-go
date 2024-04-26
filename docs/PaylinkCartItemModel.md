# PaylinkCartItemModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | The net amount of the item. The Paylink Payment Form does not multiply this figure by the value provided by the count value for the cart item, this is principally to avoid rounding errors to introduce discrepancies between the value of the goods charged for and the total amount represented by the collection of cart items. | [optional] 
**Brand** | Pointer to **string** | The brand of the item such as Nike. | [optional] 
**Category** | Pointer to **string** | The category of the item such as shoes. | [optional] 
**Count** | Pointer to **int32** | The count of how many of this item is being purchased, should the cart be editable, this value should be the default count required. The Paylink Payment Form assumes a count of 1 in the event that no value for the count field is provided for a cart item. | [optional] 
**Label** | Pointer to **string** | The label which describes the item. | [optional] 
**Max** | Pointer to **int32** | For an editable cart, the maximum number of items that can be purchased, defaults to 5. | [optional] 
**Sku** | Pointer to **string** | The stock control unit value. | [optional] 
**Variant** | Pointer to **string** | The variant field refers to the variant of the cart item to enable similar items to be distinguished according to certain criteria. For example, similar items may be distinguished in terms of size, weight and power. The Paylink Payment Form does not constrain the value of the variant field to a particular set of metrics. | [optional] 

## Methods

### NewPaylinkCartItemModel

`func NewPaylinkCartItemModel() *PaylinkCartItemModel`

NewPaylinkCartItemModel instantiates a new PaylinkCartItemModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkCartItemModelWithDefaults

`func NewPaylinkCartItemModelWithDefaults() *PaylinkCartItemModel`

NewPaylinkCartItemModelWithDefaults instantiates a new PaylinkCartItemModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PaylinkCartItemModel) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaylinkCartItemModel) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaylinkCartItemModel) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaylinkCartItemModel) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBrand

`func (o *PaylinkCartItemModel) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *PaylinkCartItemModel) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *PaylinkCartItemModel) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *PaylinkCartItemModel) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetCategory

`func (o *PaylinkCartItemModel) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PaylinkCartItemModel) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PaylinkCartItemModel) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PaylinkCartItemModel) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCount

`func (o *PaylinkCartItemModel) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaylinkCartItemModel) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaylinkCartItemModel) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaylinkCartItemModel) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLabel

`func (o *PaylinkCartItemModel) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PaylinkCartItemModel) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PaylinkCartItemModel) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PaylinkCartItemModel) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMax

`func (o *PaylinkCartItemModel) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *PaylinkCartItemModel) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *PaylinkCartItemModel) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *PaylinkCartItemModel) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetSku

`func (o *PaylinkCartItemModel) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *PaylinkCartItemModel) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *PaylinkCartItemModel) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *PaylinkCartItemModel) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetVariant

`func (o *PaylinkCartItemModel) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *PaylinkCartItemModel) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *PaylinkCartItemModel) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *PaylinkCartItemModel) HasVariant() bool`

HasVariant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


