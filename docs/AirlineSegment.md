# AirlineSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArrivalLocationCode** | **string** | A standard airline routing code (airport code or location identifier) applicable to the arrival portion of this segment.  | 
**CarrierCode** | **string** | This field contains the two character airline designator code (air carrier code or airline code) that corresponds to the airline carrier applicable for up to four flight segments of this trip itinerary.  | 
**ClassServiceCode** | **string** | This field contains a code that corresponds to the fare class (A, B, C, D, Premium, Discounted, etc.) within the overall class of service (e.g., First Class, Business, Economy) applicable to this travel segment, as specified in the IATA Standard Code allocation table.  | 
**DepartureDate** | **string** | The Departure Date for the travel segment in ISO Date Format (yyyy-MM-dd). | 
**DepartureLocationCode** | Pointer to **string** | A standard airline routing code (airport code or location identifier) applicable to the departure portion of this segment.  | [optional] 
**FlightNumber** | **string** | This field contains the carrier-assigned Flight Number for this travel segment. | 
**SegmentFare** | Pointer to **int32** | This field contains the total Fare for this travel segment. | [optional] 
**StopOverIndicator** | Pointer to **string** | O &#x3D; Stopover allowed, X &#x3D; Stopover not allowed. | [optional] 

## Methods

### NewAirlineSegment

`func NewAirlineSegment(arrivalLocationCode string, carrierCode string, classServiceCode string, departureDate string, flightNumber string, ) *AirlineSegment`

NewAirlineSegment instantiates a new AirlineSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirlineSegmentWithDefaults

`func NewAirlineSegmentWithDefaults() *AirlineSegment`

NewAirlineSegmentWithDefaults instantiates a new AirlineSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArrivalLocationCode

`func (o *AirlineSegment) GetArrivalLocationCode() string`

GetArrivalLocationCode returns the ArrivalLocationCode field if non-nil, zero value otherwise.

### GetArrivalLocationCodeOk

`func (o *AirlineSegment) GetArrivalLocationCodeOk() (*string, bool)`

GetArrivalLocationCodeOk returns a tuple with the ArrivalLocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalLocationCode

`func (o *AirlineSegment) SetArrivalLocationCode(v string)`

SetArrivalLocationCode sets ArrivalLocationCode field to given value.


### GetCarrierCode

`func (o *AirlineSegment) GetCarrierCode() string`

GetCarrierCode returns the CarrierCode field if non-nil, zero value otherwise.

### GetCarrierCodeOk

`func (o *AirlineSegment) GetCarrierCodeOk() (*string, bool)`

GetCarrierCodeOk returns a tuple with the CarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierCode

`func (o *AirlineSegment) SetCarrierCode(v string)`

SetCarrierCode sets CarrierCode field to given value.


### GetClassServiceCode

`func (o *AirlineSegment) GetClassServiceCode() string`

GetClassServiceCode returns the ClassServiceCode field if non-nil, zero value otherwise.

### GetClassServiceCodeOk

`func (o *AirlineSegment) GetClassServiceCodeOk() (*string, bool)`

GetClassServiceCodeOk returns a tuple with the ClassServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassServiceCode

`func (o *AirlineSegment) SetClassServiceCode(v string)`

SetClassServiceCode sets ClassServiceCode field to given value.


### GetDepartureDate

`func (o *AirlineSegment) GetDepartureDate() string`

GetDepartureDate returns the DepartureDate field if non-nil, zero value otherwise.

### GetDepartureDateOk

`func (o *AirlineSegment) GetDepartureDateOk() (*string, bool)`

GetDepartureDateOk returns a tuple with the DepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureDate

`func (o *AirlineSegment) SetDepartureDate(v string)`

SetDepartureDate sets DepartureDate field to given value.


### GetDepartureLocationCode

`func (o *AirlineSegment) GetDepartureLocationCode() string`

GetDepartureLocationCode returns the DepartureLocationCode field if non-nil, zero value otherwise.

### GetDepartureLocationCodeOk

`func (o *AirlineSegment) GetDepartureLocationCodeOk() (*string, bool)`

GetDepartureLocationCodeOk returns a tuple with the DepartureLocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureLocationCode

`func (o *AirlineSegment) SetDepartureLocationCode(v string)`

SetDepartureLocationCode sets DepartureLocationCode field to given value.

### HasDepartureLocationCode

`func (o *AirlineSegment) HasDepartureLocationCode() bool`

HasDepartureLocationCode returns a boolean if a field has been set.

### GetFlightNumber

`func (o *AirlineSegment) GetFlightNumber() string`

GetFlightNumber returns the FlightNumber field if non-nil, zero value otherwise.

### GetFlightNumberOk

`func (o *AirlineSegment) GetFlightNumberOk() (*string, bool)`

GetFlightNumberOk returns a tuple with the FlightNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlightNumber

`func (o *AirlineSegment) SetFlightNumber(v string)`

SetFlightNumber sets FlightNumber field to given value.


### GetSegmentFare

`func (o *AirlineSegment) GetSegmentFare() int32`

GetSegmentFare returns the SegmentFare field if non-nil, zero value otherwise.

### GetSegmentFareOk

`func (o *AirlineSegment) GetSegmentFareOk() (*int32, bool)`

GetSegmentFareOk returns a tuple with the SegmentFare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentFare

`func (o *AirlineSegment) SetSegmentFare(v int32)`

SetSegmentFare sets SegmentFare field to given value.

### HasSegmentFare

`func (o *AirlineSegment) HasSegmentFare() bool`

HasSegmentFare returns a boolean if a field has been set.

### GetStopOverIndicator

`func (o *AirlineSegment) GetStopOverIndicator() string`

GetStopOverIndicator returns the StopOverIndicator field if non-nil, zero value otherwise.

### GetStopOverIndicatorOk

`func (o *AirlineSegment) GetStopOverIndicatorOk() (*string, bool)`

GetStopOverIndicatorOk returns a tuple with the StopOverIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopOverIndicator

`func (o *AirlineSegment) SetStopOverIndicator(v string)`

SetStopOverIndicator sets StopOverIndicator field to given value.

### HasStopOverIndicator

`func (o *AirlineSegment) HasStopOverIndicator() bool`

HasStopOverIndicator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


