# RefundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The amount to refund in the lowest unit of currency with a variable length to a maximum of 12 digits.  The amount should be the total amount required to refund for the transaction up to the original processed amount.  No decimal points are to be included and no divisional characters such as 1,024.  For example with GBP £1,021.95 the amount value is 102195.  | 
**Identifier** | **string** | The identifier of the refund to process. The value should be a valid reference and may be used to perform  post processing actions and to aid in reconciliation of transactions.  The value should be a valid printable string with ASCII character ranges from 0x32 to 0x127.  The identifier is recommended to be distinct for each transaction such as a [random unique identifier](https://en.wikipedia.org/wiki/Universally_unique_identifier) this will aid in ensuring each transaction is identifiable.  When transactions are processed they are also checked for duplicate requests. Changing the identifier on a subsequent request will ensure that a transaction is considered as different.  | 
**Merchantid** | **int32** | Identifies the merchant account to perform the refund for. | 
**RefundRef** | **int32** | A reference to the original transaction number that is wanting to be refunded. The original  transaction must be on the same merchant id, previously authorised.  | 
**TransInfo** | Pointer to **string** | Further information that can be added to the transaction will display in reporting. Can be used for flexible values such as operator id. | [optional] 

## Methods

### NewRefundRequest

`func NewRefundRequest(amount int32, identifier string, merchantid int32, refundRef int32, ) *RefundRequest`

NewRefundRequest instantiates a new RefundRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundRequestWithDefaults

`func NewRefundRequestWithDefaults() *RefundRequest`

NewRefundRequestWithDefaults instantiates a new RefundRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *RefundRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RefundRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RefundRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetIdentifier

`func (o *RefundRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *RefundRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *RefundRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetMerchantid

`func (o *RefundRequest) GetMerchantid() int32`

GetMerchantid returns the Merchantid field if non-nil, zero value otherwise.

### GetMerchantidOk

`func (o *RefundRequest) GetMerchantidOk() (*int32, bool)`

GetMerchantidOk returns a tuple with the Merchantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantid

`func (o *RefundRequest) SetMerchantid(v int32)`

SetMerchantid sets Merchantid field to given value.


### GetRefundRef

`func (o *RefundRequest) GetRefundRef() int32`

GetRefundRef returns the RefundRef field if non-nil, zero value otherwise.

### GetRefundRefOk

`func (o *RefundRequest) GetRefundRefOk() (*int32, bool)`

GetRefundRefOk returns a tuple with the RefundRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundRef

`func (o *RefundRequest) SetRefundRef(v int32)`

SetRefundRef sets RefundRef field to given value.


### GetTransInfo

`func (o *RefundRequest) GetTransInfo() string`

GetTransInfo returns the TransInfo field if non-nil, zero value otherwise.

### GetTransInfoOk

`func (o *RefundRequest) GetTransInfoOk() (*string, bool)`

GetTransInfoOk returns a tuple with the TransInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransInfo

`func (o *RefundRequest) SetTransInfo(v string)`

SetTransInfo sets TransInfo field to given value.

### HasTransInfo

`func (o *RefundRequest) HasTransInfo() bool`

HasTransInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


