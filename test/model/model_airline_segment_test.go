package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleSegArrival = "LHR"
var exampleSegCarrier = "BA"
var exampleSegClass = "E"
var exampleSegDate = "2023-01-01"
var exampleSegDeparture = "JFK"
var exampleSegFlight = "BA123"
var exampleSegFare int32 = 100
var exampleSegStop = "O"

func buildExampleSegment() openapiclient.AirlineSegment {
	s := openapiclient.NewAirlineSegment(exampleSegArrival, exampleSegCarrier, exampleSegClass, exampleSegDate, exampleSegFlight)
	s.SetDepartureLocationCode(exampleSegDeparture)
	s.SetSegmentFare(exampleSegFare)
	s.SetStopOverIndicator(exampleSegStop)
	return *s
}

func TestNewAirlineSegment(t *testing.T) {
	model := openapiclient.NewAirlineSegment(exampleSegArrival, exampleSegCarrier, exampleSegClass, exampleSegDate, exampleSegFlight)
	require.NotNil(t, model)
	assert.Equal(t, exampleSegArrival, model.GetArrivalLocationCode())
	assert.Equal(t, exampleSegCarrier, model.GetCarrierCode())
	assert.False(t, model.HasDepartureLocationCode())
	assert.False(t, model.HasSegmentFare())
	assert.False(t, model.HasStopOverIndicator())
}

func TestNewAirlineSegmentWithDefaults(t *testing.T) {
	model := openapiclient.NewAirlineSegmentWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, "", model.GetArrivalLocationCode())
	assert.Equal(t, "", model.GetCarrierCode())
	assert.False(t, model.HasDepartureLocationCode())
	assert.False(t, model.HasSegmentFare())
	assert.False(t, model.HasStopOverIndicator())
}

func TestAirlineSegmentSetGetCycle(t *testing.T) {
	model := openapiclient.NewAirlineSegment(exampleSegArrival, exampleSegCarrier, exampleSegClass, exampleSegDate, exampleSegFlight)
	model.SetArrivalLocationCode("AAA")
	assert.Equal(t, "AAA", model.GetArrivalLocationCode())
	val, ok := model.GetArrivalLocationCodeOk()
	require.True(t, ok)
	if assert.NotNil(t, val) {
		assert.Equal(t, "AAA", *val)
	}

	model.SetCarrierCode("CC")
	assert.Equal(t, "CC", model.GetCarrierCode())

	model.SetClassServiceCode("B")
	assert.Equal(t, "B", model.GetClassServiceCode())

	model.SetDepartureDate("2023-02-01")
	assert.Equal(t, "2023-02-01", model.GetDepartureDate())

	model.SetDepartureLocationCode(exampleSegDeparture)
	assert.True(t, model.HasDepartureLocationCode())
	assert.Equal(t, exampleSegDeparture, model.GetDepartureLocationCode())

	model.SetFlightNumber("FL001")
	assert.Equal(t, "FL001", model.GetFlightNumber())

	model.SetSegmentFare(exampleSegFare)
	assert.True(t, model.HasSegmentFare())
	assert.Equal(t, exampleSegFare, model.GetSegmentFare())

	model.SetStopOverIndicator(exampleSegStop)
	assert.True(t, model.HasStopOverIndicator())
	assert.Equal(t, exampleSegStop, model.GetStopOverIndicator())
}

func TestAirlineSegmentJSONRoundTrip(t *testing.T) {
	model := buildExampleSegment()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.AirlineSegment
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleSegArrival, unmarshalled.GetArrivalLocationCode())
	assert.Equal(t, exampleSegCarrier, unmarshalled.GetCarrierCode())
	assert.Equal(t, exampleSegClass, unmarshalled.GetClassServiceCode())
	assert.Equal(t, exampleSegDate, unmarshalled.GetDepartureDate())
	assert.True(t, unmarshalled.HasDepartureLocationCode())
	assert.Equal(t, exampleSegDeparture, unmarshalled.GetDepartureLocationCode())
	assert.Equal(t, exampleSegFlight, unmarshalled.GetFlightNumber())
	assert.True(t, unmarshalled.HasSegmentFare())
	assert.Equal(t, exampleSegFare, unmarshalled.GetSegmentFare())
	assert.True(t, unmarshalled.HasStopOverIndicator())
	assert.Equal(t, exampleSegStop, unmarshalled.GetStopOverIndicator())
}

func TestAirlineSegmentToMap(t *testing.T) {
	model := buildExampleSegment()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "arrival_location_code") {
		assert.Equal(t, exampleSegArrival, m["arrival_location_code"])
	}
	if assert.Contains(t, m, "carrier_code") {
		assert.Equal(t, exampleSegCarrier, m["carrier_code"])
	}
	if assert.Contains(t, m, "class_service_code") {
		assert.Equal(t, exampleSegClass, m["class_service_code"])
	}
	if assert.Contains(t, m, "departure_date") {
		assert.Equal(t, exampleSegDate, m["departure_date"])
	}
	if assert.Contains(t, m, "departure_location_code") {
		assert.Equal(t, &exampleSegDeparture, m["departure_location_code"])
	}
	if assert.Contains(t, m, "flight_number") {
		assert.Equal(t, exampleSegFlight, m["flight_number"])
	}
	if assert.Contains(t, m, "segment_fare") {
		assert.Equal(t, &exampleSegFare, m["segment_fare"])
	}
	if assert.Contains(t, m, "stop_over_indicator") {
		assert.Equal(t, &exampleSegStop, m["stop_over_indicator"])
	}
}

func TestNullableAirlineSegmentGetSet(t *testing.T) {
	base := buildExampleSegment()
	n := openapiclient.NullableAirlineSegment{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableAirlineSegmentUnset(t *testing.T) {
	base := buildExampleSegment()
	n := openapiclient.NewNullableAirlineSegment(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableAirlineSegmentJSONRoundTrip(t *testing.T) {
	base := buildExampleSegment()
	n := openapiclient.NewNullableAirlineSegment(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableAirlineSegment
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleSegArrival, newN.Get().GetArrivalLocationCode())
		assert.Equal(t, exampleSegCarrier, newN.Get().GetCarrierCode())
	}
}
