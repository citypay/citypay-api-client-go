package citypay

import (
	"encoding/json"
	"testing"

	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var exampleReqAmount int32 = 10000
var exampleReqCard = "4111111111111111"
var exampleReqExpMonth int32 = 12
var exampleReqExpYear int32 = 25
var exampleReqIdentifier = "REQ123"
var exampleReqMerchantID int32 = 999
var exampleReqAvsPostcodePolicy = "1"
var exampleReqCsc = "999"
var exampleReqCscPolicy = "2"
var exampleReqCurrency = "USD"
var exampleReqDuplicatePolicy = "1"
var exampleReqMatchAvsa = "2"
var exampleReqNameOnCard = "John Doe"
var exampleReqTransInfo = "info"
var exampleReqTransType = "SALE"
var exampleReqTags = []string{"one", "two"}

func buildExampleEventData() openapiclient.EventDataModel {
	e := openapiclient.NewEventDataModel()
	e.SetEventId("EVT")
	return *e
}

func buildAuthExternalMPI() openapiclient.ExternalMPI {
	e := openapiclient.NewExternalMPI()
	e.SetEnrolled("Y")
	return *e
}

func buildExampleMcc6012() openapiclient.MCC6012 {
	m := openapiclient.NewMCC6012()
	m.SetRecipientLastname("Smith")
	return *m
}

func buildAuthThreeDSecure() openapiclient.ThreeDSecure {
	t := openapiclient.NewThreeDSecure()
	t.SetTdsPolicy("1")
	return *t
}

func buildExampleAuthRequest() *openapiclient.AuthRequest {
	r := openapiclient.NewAuthRequest(exampleReqAmount, exampleReqCard, exampleReqExpMonth, exampleReqExpYear, exampleReqIdentifier, exampleReqMerchantID)
	r.SetAirlineData(*buildExampleAdvice())
	r.SetAvsPostcodePolicy(exampleReqAvsPostcodePolicy)
	r.SetBillTo(buildExampleContact())
	r.SetCsc(exampleReqCsc)
	r.SetCscPolicy(exampleReqCscPolicy)
	r.SetCurrency(exampleReqCurrency)
	r.SetDuplicatePolicy(exampleReqDuplicatePolicy)
	r.SetEventManagement(buildExampleEventData())
	r.SetExternalMpi(buildAuthExternalMPI())
	r.SetMatchAvsa(exampleReqMatchAvsa)
	r.SetMcc6012(buildExampleMcc6012())
	r.SetNameOnCard(exampleReqNameOnCard)
	r.SetShipTo(buildExampleContact())
	r.SetTag(exampleReqTags)
	r.SetThreedsecure(buildAuthThreeDSecure())
	r.SetTransInfo(exampleReqTransInfo)
	r.SetTransType(exampleReqTransType)
	return r
}

func TestNewAuthRequest(t *testing.T) {
	model := openapiclient.NewAuthRequest(exampleReqAmount, exampleReqCard, exampleReqExpMonth, exampleReqExpYear, exampleReqIdentifier, exampleReqMerchantID)
	require.NotNil(t, model)
	assert.Equal(t, exampleReqAmount, model.GetAmount())
	assert.Equal(t, exampleReqCard, model.GetCardnumber())
	assert.False(t, model.HasAirlineData())
	assert.False(t, model.HasAvsPostcodePolicy())
	assert.False(t, model.HasBillTo())
	assert.False(t, model.HasCsc())
}

func TestNewAuthRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewAuthRequestWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, int32(0), model.GetAmount())
	assert.Equal(t, "", model.GetCardnumber())
	assert.False(t, model.HasAirlineData())
}

func TestAuthRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewAuthRequest(exampleReqAmount, exampleReqCard, exampleReqExpMonth, exampleReqExpYear, exampleReqIdentifier, exampleReqMerchantID)
	model.SetAirlineData(*buildExampleAdvice())
	assert.True(t, model.HasAirlineData())

	model.SetAvsPostcodePolicy(exampleReqAvsPostcodePolicy)
	assert.True(t, model.HasAvsPostcodePolicy())
	assert.Equal(t, exampleReqAvsPostcodePolicy, model.GetAvsPostcodePolicy())

	model.SetBillTo(buildExampleContact())
	assert.True(t, model.HasBillTo())

	model.SetCsc(exampleReqCsc)
	assert.True(t, model.HasCsc())
	assert.Equal(t, exampleReqCsc, model.GetCsc())

	model.SetCscPolicy(exampleReqCscPolicy)
	assert.True(t, model.HasCscPolicy())
	assert.Equal(t, exampleReqCscPolicy, model.GetCscPolicy())

	model.SetCurrency(exampleReqCurrency)
	assert.True(t, model.HasCurrency())
	assert.Equal(t, exampleReqCurrency, model.GetCurrency())

	model.SetDuplicatePolicy(exampleReqDuplicatePolicy)
	assert.True(t, model.HasDuplicatePolicy())
	assert.Equal(t, exampleReqDuplicatePolicy, model.GetDuplicatePolicy())

	model.SetEventManagement(buildExampleEventData())
	assert.True(t, model.HasEventManagement())

	model.SetExternalMpi(buildAuthExternalMPI())
	assert.True(t, model.HasExternalMpi())

	model.SetMatchAvsa(exampleReqMatchAvsa)
	assert.True(t, model.HasMatchAvsa())
	assert.Equal(t, exampleReqMatchAvsa, model.GetMatchAvsa())

	model.SetMcc6012(buildExampleMcc6012())
	assert.True(t, model.HasMcc6012())

	model.SetNameOnCard(exampleReqNameOnCard)
	assert.True(t, model.HasNameOnCard())
	assert.Equal(t, exampleReqNameOnCard, model.GetNameOnCard())

	model.SetShipTo(buildExampleContact())
	assert.True(t, model.HasShipTo())

	model.SetTag(exampleReqTags)
	assert.True(t, model.HasTag())
	assert.Equal(t, exampleReqTags, model.GetTag())

	model.SetThreedsecure(buildAuthThreeDSecure())
	assert.True(t, model.HasThreedsecure())

	model.SetTransInfo(exampleReqTransInfo)
	assert.True(t, model.HasTransInfo())
	assert.Equal(t, exampleReqTransInfo, model.GetTransInfo())

	model.SetTransType(exampleReqTransType)
	assert.True(t, model.HasTransType())
	assert.Equal(t, exampleReqTransType, model.GetTransType())
}

func TestAuthRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleAuthRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.AuthRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleReqAmount, unmarshalled.GetAmount())
	assert.True(t, unmarshalled.HasTag())
	assert.Equal(t, exampleReqTransType, unmarshalled.GetTransType())
}

func TestAuthRequestToMap(t *testing.T) {
	model := buildExampleAuthRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "amount") {
		assert.Equal(t, exampleReqAmount, m["amount"])
	}
	if assert.Contains(t, m, "tag") {
		assert.Equal(t, exampleReqTags, m["tag"])
	}
}

func TestNullableAuthRequestGetSet(t *testing.T) {
	base := buildExampleAuthRequest()
	n := openapiclient.NullableAuthRequest{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableAuthRequestUnset(t *testing.T) {
	base := buildExampleAuthRequest()
	n := openapiclient.NewNullableAuthRequest(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableAuthRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleAuthRequest()
	n := openapiclient.NewNullableAuthRequest(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableAuthRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleReqIdentifier, newN.Get().GetIdentifier())
	}
}
