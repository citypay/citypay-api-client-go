# PaylinkCart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | Pointer to [**[]PaylinkCartItemModel**](PaylinkCartItemModel.md) |  | [optional] 
**Coupon** | Pointer to **string** | A coupon redeemed with the transaction. | [optional] 
**Mode** | Pointer to **int32** | The mode field specifies the behaviour or functionality of the cart.  Valid values are:   0 - No cart - No cart is shown  1 - Read-only - The cart is shown with a breakdown of the item details provided by objects in the contents array.  2 - Selection cart - The cart is shown as a drop-down box of available cart items that the customer can a single item select from.  3 - Dynamic cart - a text box is rendered to enable the operator to input an amount.  4 - Multi cart - The cart is displayed with items rendered with selectable quantities.  | [optional] 
**ProductDescription** | Pointer to **string** | Specifies a description about the product or service that is the subject of the transaction. It will be rendered in the header of the page with no labels. | [optional] 
**ProductInformation** | Pointer to **string** | Specifies information about the product or service that is the subject of the transaction. It will be rendered in the header of the page. | [optional] 
**Shipping** | Pointer to **int32** | The shipping amount of the transaction in the lowest denomination of currency. | [optional] 
**Tax** | Pointer to **int32** | The tax amount of the transaction in the lowest denomination of currency. | [optional] 

## Methods

### NewPaylinkCart

`func NewPaylinkCart() *PaylinkCart`

NewPaylinkCart instantiates a new PaylinkCart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkCartWithDefaults

`func NewPaylinkCartWithDefaults() *PaylinkCart`

NewPaylinkCartWithDefaults instantiates a new PaylinkCart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *PaylinkCart) GetContents() []PaylinkCartItemModel`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *PaylinkCart) GetContentsOk() (*[]PaylinkCartItemModel, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *PaylinkCart) SetContents(v []PaylinkCartItemModel)`

SetContents sets Contents field to given value.

### HasContents

`func (o *PaylinkCart) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetCoupon

`func (o *PaylinkCart) GetCoupon() string`

GetCoupon returns the Coupon field if non-nil, zero value otherwise.

### GetCouponOk

`func (o *PaylinkCart) GetCouponOk() (*string, bool)`

GetCouponOk returns a tuple with the Coupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupon

`func (o *PaylinkCart) SetCoupon(v string)`

SetCoupon sets Coupon field to given value.

### HasCoupon

`func (o *PaylinkCart) HasCoupon() bool`

HasCoupon returns a boolean if a field has been set.

### GetMode

`func (o *PaylinkCart) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PaylinkCart) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PaylinkCart) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PaylinkCart) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetProductDescription

`func (o *PaylinkCart) GetProductDescription() string`

GetProductDescription returns the ProductDescription field if non-nil, zero value otherwise.

### GetProductDescriptionOk

`func (o *PaylinkCart) GetProductDescriptionOk() (*string, bool)`

GetProductDescriptionOk returns a tuple with the ProductDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescription

`func (o *PaylinkCart) SetProductDescription(v string)`

SetProductDescription sets ProductDescription field to given value.

### HasProductDescription

`func (o *PaylinkCart) HasProductDescription() bool`

HasProductDescription returns a boolean if a field has been set.

### GetProductInformation

`func (o *PaylinkCart) GetProductInformation() string`

GetProductInformation returns the ProductInformation field if non-nil, zero value otherwise.

### GetProductInformationOk

`func (o *PaylinkCart) GetProductInformationOk() (*string, bool)`

GetProductInformationOk returns a tuple with the ProductInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductInformation

`func (o *PaylinkCart) SetProductInformation(v string)`

SetProductInformation sets ProductInformation field to given value.

### HasProductInformation

`func (o *PaylinkCart) HasProductInformation() bool`

HasProductInformation returns a boolean if a field has been set.

### GetShipping

`func (o *PaylinkCart) GetShipping() int32`

GetShipping returns the Shipping field if non-nil, zero value otherwise.

### GetShippingOk

`func (o *PaylinkCart) GetShippingOk() (*int32, bool)`

GetShippingOk returns a tuple with the Shipping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipping

`func (o *PaylinkCart) SetShipping(v int32)`

SetShipping sets Shipping field to given value.

### HasShipping

`func (o *PaylinkCart) HasShipping() bool`

HasShipping returns a boolean if a field has been set.

### GetTax

`func (o *PaylinkCart) GetTax() int32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *PaylinkCart) GetTaxOk() (*int32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *PaylinkCart) SetTax(v int32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *PaylinkCart) HasTax() bool`

HasTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


