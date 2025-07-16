package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleChangeRespCount int32 = 2
var exampleChangeRespMax int32 = 10
var exampleChangeRespNext = "nxt"

func buildExamplePaylinkTokenStatusChangeResponse() openapiclient.PaylinkTokenStatusChangeResponse {
	tok := buildExamplePaylinkTokenStatus()
	r := openapiclient.NewPaylinkTokenStatusChangeResponse([]openapiclient.PaylinkTokenStatus{tok})
	r.SetCount(exampleChangeRespCount)
	r.SetMaxResults(exampleChangeRespMax)
	r.SetNextToken(exampleChangeRespNext)
	return *r
}

func TestNewPaylinkTokenStatusChangeResponse(t *testing.T) {
	model := openapiclient.NewPaylinkTokenStatusChangeResponse([]openapiclient.PaylinkTokenStatus{})
	require.NotNil(t, model)
	assert.NotNil(t, model.GetTokens())
	assert.False(t, model.HasCount())
}

func TestNewPaylinkTokenStatusChangeResponseWithDefaults(t *testing.T) {
	model := openapiclient.NewPaylinkTokenStatusChangeResponseWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasMaxResults())
}

func TestPaylinkTokenStatusChangeResponseSetGetCycle(t *testing.T) {
	model := openapiclient.NewPaylinkTokenStatusChangeResponse([]openapiclient.PaylinkTokenStatus{})
	model.SetCount(exampleChangeRespCount)
	assert.True(t, model.HasCount())
	assert.Equal(t, exampleChangeRespCount, model.GetCount())

	model.SetMaxResults(exampleChangeRespMax)
	assert.True(t, model.HasMaxResults())
	assert.Equal(t, exampleChangeRespMax, model.GetMaxResults())

	model.SetNextToken(exampleChangeRespNext)
	assert.True(t, model.HasNextToken())
	assert.Equal(t, exampleChangeRespNext, model.GetNextToken())

	tok := buildExamplePaylinkTokenStatus()
	model.SetTokens([]openapiclient.PaylinkTokenStatus{tok})
	assert.Len(t, model.GetTokens(), 1)
}

func TestPaylinkTokenStatusChangeResponseJSONRoundTrip(t *testing.T) {
	model := buildExamplePaylinkTokenStatusChangeResponse()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.PaylinkTokenStatusChangeResponse
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasCount())
	assert.Equal(t, exampleChangeRespCount, unmarshalled.GetCount())
}

func TestPaylinkTokenStatusChangeResponseToMap(t *testing.T) {
	model := buildExamplePaylinkTokenStatusChangeResponse()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "nextToken") {
		assert.Equal(t, &exampleChangeRespNext, m["nextToken"])
	}
	assert.Contains(t, m, "tokens")
}

func TestNullablePaylinkTokenStatusChangeResponseGetSet(t *testing.T) {
	base := buildExamplePaylinkTokenStatusChangeResponse()
	n := openapiclient.NullablePaylinkTokenStatusChangeResponse{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkTokenStatusChangeResponseUnset(t *testing.T) {
	base := buildExamplePaylinkTokenStatusChangeResponse()
	n := openapiclient.NewNullablePaylinkTokenStatusChangeResponse(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullablePaylinkTokenStatusChangeResponseJSONRoundTrip(t *testing.T) {
	base := buildExamplePaylinkTokenStatusChangeResponse()
	n := openapiclient.NewNullablePaylinkTokenStatusChangeResponse(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullablePaylinkTokenStatusChangeResponse
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleChangeRespMax, newN.Get().GetMaxResults())
	}
}
