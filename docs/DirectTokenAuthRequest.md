# DirectTokenAuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nonce** | Pointer to **string** | A random value string which is provided to the API to perform a digest. The value will be used by its UTF-8 byte representation of any digest function.  | [optional] 
**RedirectFailure** | Pointer to **string** | The URL used to redirect back to your site when a transaction has been rejected or declined. Required if a url-encoded request.  | [optional] 
**RedirectSuccess** | Pointer to **string** | The URL used to redirect back to your site when a transaction has been authorised. Required if a url-encoded request.  | [optional] 
**Token** | Pointer to **string** | The token required to process the transaction as presented by the direct post methodology.  | [optional] 

## Methods

### NewDirectTokenAuthRequest

`func NewDirectTokenAuthRequest() *DirectTokenAuthRequest`

NewDirectTokenAuthRequest instantiates a new DirectTokenAuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectTokenAuthRequestWithDefaults

`func NewDirectTokenAuthRequestWithDefaults() *DirectTokenAuthRequest`

NewDirectTokenAuthRequestWithDefaults instantiates a new DirectTokenAuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonce

`func (o *DirectTokenAuthRequest) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *DirectTokenAuthRequest) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *DirectTokenAuthRequest) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *DirectTokenAuthRequest) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetRedirectFailure

`func (o *DirectTokenAuthRequest) GetRedirectFailure() string`

GetRedirectFailure returns the RedirectFailure field if non-nil, zero value otherwise.

### GetRedirectFailureOk

`func (o *DirectTokenAuthRequest) GetRedirectFailureOk() (*string, bool)`

GetRedirectFailureOk returns a tuple with the RedirectFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectFailure

`func (o *DirectTokenAuthRequest) SetRedirectFailure(v string)`

SetRedirectFailure sets RedirectFailure field to given value.

### HasRedirectFailure

`func (o *DirectTokenAuthRequest) HasRedirectFailure() bool`

HasRedirectFailure returns a boolean if a field has been set.

### GetRedirectSuccess

`func (o *DirectTokenAuthRequest) GetRedirectSuccess() string`

GetRedirectSuccess returns the RedirectSuccess field if non-nil, zero value otherwise.

### GetRedirectSuccessOk

`func (o *DirectTokenAuthRequest) GetRedirectSuccessOk() (*string, bool)`

GetRedirectSuccessOk returns a tuple with the RedirectSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectSuccess

`func (o *DirectTokenAuthRequest) SetRedirectSuccess(v string)`

SetRedirectSuccess sets RedirectSuccess field to given value.

### HasRedirectSuccess

`func (o *DirectTokenAuthRequest) HasRedirectSuccess() bool`

HasRedirectSuccess returns a boolean if a field has been set.

### GetToken

`func (o *DirectTokenAuthRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DirectTokenAuthRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DirectTokenAuthRequest) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DirectTokenAuthRequest) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


