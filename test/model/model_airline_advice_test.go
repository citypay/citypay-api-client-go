package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleCarrierName = "Airline"
var exampleTicketIssueCity = "London"
var exampleTicketIssueDate = "2023-01-05"
var exampleTicketIssueName = "Agency"
var exampleTicketNo = "1234567890123"
var exampleTransactionType = "TKT"
var examplePassengerName = "John Doe"
var exampleOriginalTicketNo = "9876543210987"
var exampleNoAirSegments int32 = 2
var exampleNumberInParty int32 = 1
var exampleConjunctionIndicator = true
var exampleEticketIndicator = false

func buildExampleAdvice() *openapiclient.AirlineAdvice {
	seg1 := buildExampleSegment()
	seg2 := buildExampleSegment()
	a := openapiclient.NewAirlineAdvice(exampleCarrierName, seg1, exampleTicketIssueCity, exampleTicketIssueDate, exampleTicketIssueName, exampleTicketNo, exampleTransactionType)
	a.SetConjunctionTicketIndicator(exampleConjunctionIndicator)
	a.SetEticketIndicator(exampleEticketIndicator)
	a.SetNoAirSegments(exampleNoAirSegments)
	a.SetNumberInParty(exampleNumberInParty)
	a.SetOriginalTicketNo(exampleOriginalTicketNo)
	a.SetPassengerName(examplePassengerName)
	a.SetSegment2(seg2)
	return a
}

func TestNewAirlineAdvice(t *testing.T) {
	seg := buildExampleSegment()
	model := openapiclient.NewAirlineAdvice(exampleCarrierName, seg, exampleTicketIssueCity, exampleTicketIssueDate, exampleTicketIssueName, exampleTicketNo, exampleTransactionType)
	require.NotNil(t, model)
	assert.Equal(t, exampleCarrierName, model.GetCarrierName())
	assert.False(t, model.HasConjunctionTicketIndicator())
	assert.False(t, model.HasEticketIndicator())
	assert.False(t, model.HasNoAirSegments())
	assert.False(t, model.HasNumberInParty())
	assert.False(t, model.HasOriginalTicketNo())
	assert.False(t, model.HasPassengerName())
	assert.False(t, model.HasSegment2())
	assert.False(t, model.HasSegment3())
	assert.False(t, model.HasSegment4())
}

func TestNewAirlineAdviceWithDefaults(t *testing.T) {
	model := openapiclient.NewAirlineAdviceWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, "", model.GetCarrierName())
	assert.False(t, model.HasConjunctionTicketIndicator())
}

func TestAirlineAdviceSetGetCycle(t *testing.T) {
	seg := buildExampleSegment()
	model := openapiclient.NewAirlineAdvice(exampleCarrierName, seg, exampleTicketIssueCity, exampleTicketIssueDate, exampleTicketIssueName, exampleTicketNo, exampleTransactionType)
	model.SetCarrierName("NewAir")
	assert.Equal(t, "NewAir", model.GetCarrierName())

	model.SetConjunctionTicketIndicator(exampleConjunctionIndicator)
	assert.True(t, model.HasConjunctionTicketIndicator())
	assert.Equal(t, exampleConjunctionIndicator, model.GetConjunctionTicketIndicator())

	model.SetEticketIndicator(exampleEticketIndicator)
	assert.True(t, model.HasEticketIndicator())
	assert.Equal(t, exampleEticketIndicator, model.GetEticketIndicator())

	model.SetNoAirSegments(exampleNoAirSegments)
	assert.True(t, model.HasNoAirSegments())
	assert.Equal(t, exampleNoAirSegments, model.GetNoAirSegments())

	model.SetNumberInParty(exampleNumberInParty)
	assert.True(t, model.HasNumberInParty())
	assert.Equal(t, exampleNumberInParty, model.GetNumberInParty())

	model.SetOriginalTicketNo(exampleOriginalTicketNo)
	assert.True(t, model.HasOriginalTicketNo())
	assert.Equal(t, exampleOriginalTicketNo, model.GetOriginalTicketNo())

	model.SetPassengerName(examplePassengerName)
	assert.True(t, model.HasPassengerName())
	assert.Equal(t, examplePassengerName, model.GetPassengerName())

	seg2 := buildExampleSegment()
	model.SetSegment2(seg2)
	assert.True(t, model.HasSegment2())
	assert.Equal(t, seg2, model.GetSegment2())
}

func TestAirlineAdviceJSONRoundTrip(t *testing.T) {
	model := buildExampleAdvice()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.AirlineAdvice
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleCarrierName, unmarshalled.GetCarrierName())
	assert.True(t, unmarshalled.HasConjunctionTicketIndicator())
	assert.Equal(t, exampleConjunctionIndicator, unmarshalled.GetConjunctionTicketIndicator())
	assert.True(t, unmarshalled.HasSegment2())
}

func TestAirlineAdviceToMap(t *testing.T) {
	model := buildExampleAdvice()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "carrier_name") {
		assert.Equal(t, exampleCarrierName, m["carrier_name"])
	}
	if assert.Contains(t, m, "conjunction_ticket_indicator") {
		assert.Equal(t, &exampleConjunctionIndicator, m["conjunction_ticket_indicator"])
	}
	if assert.Contains(t, m, "segment2") {
		if assert.NotNil(t, m["segment2"]) {
			assert.Equal(t, model.GetSegment2(), *m["segment2"].(*openapiclient.AirlineSegment))
		}
	}
}

func TestNullableAirlineAdviceGetSet(t *testing.T) {
	base := buildExampleAdvice()
	n := openapiclient.NullableAirlineAdvice{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableAirlineAdviceUnset(t *testing.T) {
	base := buildExampleAdvice()
	n := openapiclient.NewNullableAirlineAdvice(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableAirlineAdviceJSONRoundTrip(t *testing.T) {
	base := buildExampleAdvice()
	n := openapiclient.NewNullableAirlineAdvice(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableAirlineAdvice
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleCarrierName, newN.Get().GetCarrierName())
	}
}
