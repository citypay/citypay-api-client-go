# Decision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthResponse** | Pointer to [**AuthResponse**](AuthResponse.md) |  | [optional] 
**RequestChallenged** | Pointer to [**RequestChallenged**](RequestChallenged.md) |  | [optional] 

## Methods

### NewDecision

`func NewDecision() *Decision`

NewDecision instantiates a new Decision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecisionWithDefaults

`func NewDecisionWithDefaults() *Decision`

NewDecisionWithDefaults instantiates a new Decision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthResponse

`func (o *Decision) GetAuthResponse() AuthResponse`

GetAuthResponse returns the AuthResponse field if non-nil, zero value otherwise.

### GetAuthResponseOk

`func (o *Decision) GetAuthResponseOk() (*AuthResponse, bool)`

GetAuthResponseOk returns a tuple with the AuthResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthResponse

`func (o *Decision) SetAuthResponse(v AuthResponse)`

SetAuthResponse sets AuthResponse field to given value.

### HasAuthResponse

`func (o *Decision) HasAuthResponse() bool`

HasAuthResponse returns a boolean if a field has been set.

### GetRequestChallenged

`func (o *Decision) GetRequestChallenged() RequestChallenged`

GetRequestChallenged returns the RequestChallenged field if non-nil, zero value otherwise.

### GetRequestChallengedOk

`func (o *Decision) GetRequestChallengedOk() (*RequestChallenged, bool)`

GetRequestChallengedOk returns a tuple with the RequestChallenged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestChallenged

`func (o *Decision) SetRequestChallenged(v RequestChallenged)`

SetRequestChallenged sets RequestChallenged field to given value.

### HasRequestChallenged

`func (o *Decision) HasRequestChallenged() bool`

HasRequestChallenged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


