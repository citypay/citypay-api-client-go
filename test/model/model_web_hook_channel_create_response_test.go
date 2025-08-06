package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleWebHookChannelCreateResponseEndpoint = "endpoint1"
var exampleWebHookChannelCreateResponseId = "channel123"

func buildExampleWebHookChannelCreateResponse() *openapiclient.WebHookChannelCreateResponse {
	m := openapiclient.NewWebHookChannelCreateResponse()
	m.SetEndpointId(exampleWebHookChannelCreateResponseEndpoint)
	m.SetWebChannelId(exampleWebHookChannelCreateResponseId)
	return m
}

func TestNewWebHookChannelCreateResponse(t *testing.T) {
	model := openapiclient.NewWebHookChannelCreateResponse()
	require.NotNil(t, model)
	assert.False(t, model.HasEndpointId())
}

func TestNewWebHookChannelCreateResponseWithDefaults(t *testing.T) {
	model := openapiclient.NewWebHookChannelCreateResponseWithDefaults()
	require.NotNil(t, model)
}

func TestWebHookChannelCreateResponseSetGetCycle(t *testing.T) {
	model := openapiclient.NewWebHookChannelCreateResponse()
	model.SetEndpointId(exampleWebHookChannelCreateResponseEndpoint)
	assert.True(t, model.HasEndpointId())
	assert.Equal(t, exampleWebHookChannelCreateResponseEndpoint, model.GetEndpointId())
	epPtr, ok := model.GetEndpointIdOk()
	require.True(t, ok)
	if assert.NotNil(t, epPtr) {
		assert.Equal(t, exampleWebHookChannelCreateResponseEndpoint, *epPtr)
	}
}

func TestWebHookChannelCreateResponseJSONRoundTrip(t *testing.T) {
	model := buildExampleWebHookChannelCreateResponse()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.WebHookChannelCreateResponse
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleWebHookChannelCreateResponseId, unmarshalled.GetWebChannelId())
}

func TestWebHookChannelCreateResponseToMap(t *testing.T) {
	model := buildExampleWebHookChannelCreateResponse()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "endpoint_id") {
		assert.Equal(t, &exampleWebHookChannelCreateResponseEndpoint, m["endpoint_id"])
	}
}

func TestNullableWebHookChannelCreateResponseGetSet(t *testing.T) {
	base := buildExampleWebHookChannelCreateResponse()
	n := openapiclient.NullableWebHookChannelCreateResponse{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableWebHookChannelCreateResponseUnset(t *testing.T) {
	base := buildExampleWebHookChannelCreateResponse()
	n := openapiclient.NewNullableWebHookChannelCreateResponse(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableWebHookChannelCreateResponseJSONRoundTrip(t *testing.T) {
	base := buildExampleWebHookChannelCreateResponse()
	n := openapiclient.NewNullableWebHookChannelCreateResponse(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableWebHookChannelCreateResponse
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleWebHookChannelCreateResponseId, newN.Get().GetWebChannelId())
	}
}
