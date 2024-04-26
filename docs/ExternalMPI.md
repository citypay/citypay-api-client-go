# ExternalMPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenResult** | Pointer to **string** | The authentication result available from the MPI. | [optional] 
**Cavv** | Pointer to **string** | A value determining the cardholder verification value supplied by the card scheme. | [optional] 
**Eci** | Pointer to **int32** | The obtained e-commerce indicator from the MPI. | [optional] 
**Enrolled** | Pointer to **string** | A value determining whether the card holder was enrolled. | [optional] 
**Xid** | Pointer to **string** | The XID used for processing with the MPI. | [optional] 

## Methods

### NewExternalMPI

`func NewExternalMPI() *ExternalMPI`

NewExternalMPI instantiates a new ExternalMPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalMPIWithDefaults

`func NewExternalMPIWithDefaults() *ExternalMPI`

NewExternalMPIWithDefaults instantiates a new ExternalMPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenResult

`func (o *ExternalMPI) GetAuthenResult() string`

GetAuthenResult returns the AuthenResult field if non-nil, zero value otherwise.

### GetAuthenResultOk

`func (o *ExternalMPI) GetAuthenResultOk() (*string, bool)`

GetAuthenResultOk returns a tuple with the AuthenResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenResult

`func (o *ExternalMPI) SetAuthenResult(v string)`

SetAuthenResult sets AuthenResult field to given value.

### HasAuthenResult

`func (o *ExternalMPI) HasAuthenResult() bool`

HasAuthenResult returns a boolean if a field has been set.

### GetCavv

`func (o *ExternalMPI) GetCavv() string`

GetCavv returns the Cavv field if non-nil, zero value otherwise.

### GetCavvOk

`func (o *ExternalMPI) GetCavvOk() (*string, bool)`

GetCavvOk returns a tuple with the Cavv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCavv

`func (o *ExternalMPI) SetCavv(v string)`

SetCavv sets Cavv field to given value.

### HasCavv

`func (o *ExternalMPI) HasCavv() bool`

HasCavv returns a boolean if a field has been set.

### GetEci

`func (o *ExternalMPI) GetEci() int32`

GetEci returns the Eci field if non-nil, zero value otherwise.

### GetEciOk

`func (o *ExternalMPI) GetEciOk() (*int32, bool)`

GetEciOk returns a tuple with the Eci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEci

`func (o *ExternalMPI) SetEci(v int32)`

SetEci sets Eci field to given value.

### HasEci

`func (o *ExternalMPI) HasEci() bool`

HasEci returns a boolean if a field has been set.

### GetEnrolled

`func (o *ExternalMPI) GetEnrolled() string`

GetEnrolled returns the Enrolled field if non-nil, zero value otherwise.

### GetEnrolledOk

`func (o *ExternalMPI) GetEnrolledOk() (*string, bool)`

GetEnrolledOk returns a tuple with the Enrolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolled

`func (o *ExternalMPI) SetEnrolled(v string)`

SetEnrolled sets Enrolled field to given value.

### HasEnrolled

`func (o *ExternalMPI) HasEnrolled() bool`

HasEnrolled returns a boolean if a field has been set.

### GetXid

`func (o *ExternalMPI) GetXid() string`

GetXid returns the Xid field if non-nil, zero value otherwise.

### GetXidOk

`func (o *ExternalMPI) GetXidOk() (*string, bool)`

GetXidOk returns a tuple with the Xid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXid

`func (o *ExternalMPI) SetXid(v string)`

SetXid sets Xid field to given value.

### HasXid

`func (o *ExternalMPI) HasXid() bool`

HasXid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


