# NetSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditItemsAmount** | Pointer to **string** | The total value of refund (credit) transaction items. Represents the sum of funds returned to customers. | [optional] 
**CreditItemsCount** | Pointer to **int32** | The count of refund (credit) transaction items. Reflects the number of refund transactions processed. | [optional] 
**CreditItemsValue** | Pointer to **int32** | The total value of refund (credit) transaction items. Represents the sum of funds returned to customers. | [optional] 
**DebitItemsAmount** | Pointer to **string** | The total value of charge (debit) transaction items. Represents the sum of funds received from charges. | [optional] 
**DebitItemsCount** | Pointer to **int32** | The count of charge (debit) transaction items. Indicates the number of charge transactions processed. | [optional] 
**DebitItemsValue** | Pointer to **int32** | The total value of charge (debit) transaction items. Represents the sum of funds received from charges. | [optional] 
**NetAmount** | Pointer to **int32** | The absolute net value, reflecting the net gain or loss from transactions. | [optional] 
**TotalCount** | Pointer to **int32** | The total count of all transaction items. | [optional] 

## Methods

### NewNetSummaryResponse

`func NewNetSummaryResponse() *NetSummaryResponse`

NewNetSummaryResponse instantiates a new NetSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetSummaryResponseWithDefaults

`func NewNetSummaryResponseWithDefaults() *NetSummaryResponse`

NewNetSummaryResponseWithDefaults instantiates a new NetSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditItemsAmount

`func (o *NetSummaryResponse) GetCreditItemsAmount() string`

GetCreditItemsAmount returns the CreditItemsAmount field if non-nil, zero value otherwise.

### GetCreditItemsAmountOk

`func (o *NetSummaryResponse) GetCreditItemsAmountOk() (*string, bool)`

GetCreditItemsAmountOk returns a tuple with the CreditItemsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditItemsAmount

`func (o *NetSummaryResponse) SetCreditItemsAmount(v string)`

SetCreditItemsAmount sets CreditItemsAmount field to given value.

### HasCreditItemsAmount

`func (o *NetSummaryResponse) HasCreditItemsAmount() bool`

HasCreditItemsAmount returns a boolean if a field has been set.

### GetCreditItemsCount

`func (o *NetSummaryResponse) GetCreditItemsCount() int32`

GetCreditItemsCount returns the CreditItemsCount field if non-nil, zero value otherwise.

### GetCreditItemsCountOk

`func (o *NetSummaryResponse) GetCreditItemsCountOk() (*int32, bool)`

GetCreditItemsCountOk returns a tuple with the CreditItemsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditItemsCount

`func (o *NetSummaryResponse) SetCreditItemsCount(v int32)`

SetCreditItemsCount sets CreditItemsCount field to given value.

### HasCreditItemsCount

`func (o *NetSummaryResponse) HasCreditItemsCount() bool`

HasCreditItemsCount returns a boolean if a field has been set.

### GetCreditItemsValue

`func (o *NetSummaryResponse) GetCreditItemsValue() int32`

GetCreditItemsValue returns the CreditItemsValue field if non-nil, zero value otherwise.

### GetCreditItemsValueOk

`func (o *NetSummaryResponse) GetCreditItemsValueOk() (*int32, bool)`

GetCreditItemsValueOk returns a tuple with the CreditItemsValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditItemsValue

`func (o *NetSummaryResponse) SetCreditItemsValue(v int32)`

SetCreditItemsValue sets CreditItemsValue field to given value.

### HasCreditItemsValue

`func (o *NetSummaryResponse) HasCreditItemsValue() bool`

HasCreditItemsValue returns a boolean if a field has been set.

### GetDebitItemsAmount

`func (o *NetSummaryResponse) GetDebitItemsAmount() string`

GetDebitItemsAmount returns the DebitItemsAmount field if non-nil, zero value otherwise.

### GetDebitItemsAmountOk

`func (o *NetSummaryResponse) GetDebitItemsAmountOk() (*string, bool)`

GetDebitItemsAmountOk returns a tuple with the DebitItemsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitItemsAmount

`func (o *NetSummaryResponse) SetDebitItemsAmount(v string)`

SetDebitItemsAmount sets DebitItemsAmount field to given value.

### HasDebitItemsAmount

`func (o *NetSummaryResponse) HasDebitItemsAmount() bool`

HasDebitItemsAmount returns a boolean if a field has been set.

### GetDebitItemsCount

`func (o *NetSummaryResponse) GetDebitItemsCount() int32`

GetDebitItemsCount returns the DebitItemsCount field if non-nil, zero value otherwise.

### GetDebitItemsCountOk

`func (o *NetSummaryResponse) GetDebitItemsCountOk() (*int32, bool)`

GetDebitItemsCountOk returns a tuple with the DebitItemsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitItemsCount

`func (o *NetSummaryResponse) SetDebitItemsCount(v int32)`

SetDebitItemsCount sets DebitItemsCount field to given value.

### HasDebitItemsCount

`func (o *NetSummaryResponse) HasDebitItemsCount() bool`

HasDebitItemsCount returns a boolean if a field has been set.

### GetDebitItemsValue

`func (o *NetSummaryResponse) GetDebitItemsValue() int32`

GetDebitItemsValue returns the DebitItemsValue field if non-nil, zero value otherwise.

### GetDebitItemsValueOk

`func (o *NetSummaryResponse) GetDebitItemsValueOk() (*int32, bool)`

GetDebitItemsValueOk returns a tuple with the DebitItemsValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitItemsValue

`func (o *NetSummaryResponse) SetDebitItemsValue(v int32)`

SetDebitItemsValue sets DebitItemsValue field to given value.

### HasDebitItemsValue

`func (o *NetSummaryResponse) HasDebitItemsValue() bool`

HasDebitItemsValue returns a boolean if a field has been set.

### GetNetAmount

`func (o *NetSummaryResponse) GetNetAmount() int32`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *NetSummaryResponse) GetNetAmountOk() (*int32, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *NetSummaryResponse) SetNetAmount(v int32)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *NetSummaryResponse) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### GetTotalCount

`func (o *NetSummaryResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *NetSummaryResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *NetSummaryResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *NetSummaryResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


