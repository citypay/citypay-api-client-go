package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleProcessMessage = "ok"
var exampleProcessValid = true

func buildExampleProcessBatchResponse() openapiclient.ProcessBatchResponse {
	r := openapiclient.NewProcessBatchResponse(exampleProcessValid)
	r.SetMessage(exampleProcessMessage)
	return *r
}

func TestNewProcessBatchResponse(t *testing.T) {
	model := openapiclient.NewProcessBatchResponse(exampleProcessValid)
	require.NotNil(t, model)
	assert.Equal(t, exampleProcessValid, model.GetValid())
	assert.False(t, model.HasMessage())
}

func TestNewProcessBatchResponseWithDefaults(t *testing.T) {
	model := openapiclient.NewProcessBatchResponseWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.GetValid())
}

func TestProcessBatchResponseSetGetCycle(t *testing.T) {
	model := openapiclient.NewProcessBatchResponse(exampleProcessValid)
	model.SetMessage(exampleProcessMessage)
	assert.True(t, model.HasMessage())
	assert.Equal(t, exampleProcessMessage, model.GetMessage())
}

func TestProcessBatchResponseJSONRoundTrip(t *testing.T) {
	model := buildExampleProcessBatchResponse()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.ProcessBatchResponse
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasMessage())
	assert.Equal(t, exampleProcessMessage, unmarshalled.GetMessage())
}

func TestProcessBatchResponseToMap(t *testing.T) {
	model := buildExampleProcessBatchResponse()
	m, err := model.ToMap()
	require.NoError(t, err)
	assert.Equal(t, exampleProcessValid, m["valid"])
	if assert.Contains(t, m, "message") {
		assert.Equal(t, &exampleProcessMessage, m["message"])
	}
}

func TestNullableProcessBatchResponseGetSet(t *testing.T) {
	base := buildExampleProcessBatchResponse()
	n := openapiclient.NullableProcessBatchResponse{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableProcessBatchResponseUnset(t *testing.T) {
	base := buildExampleProcessBatchResponse()
	n := openapiclient.NewNullableProcessBatchResponse(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableProcessBatchResponseJSONRoundTrip(t *testing.T) {
	base := buildExampleProcessBatchResponse()
	n := openapiclient.NewNullableProcessBatchResponse(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableProcessBatchResponse
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleProcessValid, newN.Get().GetValid())
	}
}
