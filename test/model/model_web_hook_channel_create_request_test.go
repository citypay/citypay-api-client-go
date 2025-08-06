package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleWebHookChannelCreateRequestName = "channel1"
var exampleWebHookChannelCreateRequestClient = "client1"
var exampleWebHookChannelCreateRequestEndpoint = "endpoint1"

func buildExampleWebHookChannelCreateRequest() *openapiclient.WebHookChannelCreateRequest {
	cfg := openapiclient.HttpConfigAsWebHookChannelCreateRequestConfig(buildExampleHttpConfig())
	m := openapiclient.NewWebHookChannelCreateRequest(exampleWebHookChannelCreateRequestName, exampleWebHookChannelCreateRequestClient, cfg, exampleWebHookChannelCreateRequestEndpoint)
	return m
}

func TestNewWebHookChannelCreateRequest(t *testing.T) {
	cfg := openapiclient.HttpConfigAsWebHookChannelCreateRequestConfig(buildExampleHttpConfig())
	model := openapiclient.NewWebHookChannelCreateRequest(exampleWebHookChannelCreateRequestName, exampleWebHookChannelCreateRequestClient, cfg, exampleWebHookChannelCreateRequestEndpoint)
	require.NotNil(t, model)
	assert.Equal(t, exampleWebHookChannelCreateRequestName, model.GetChannelName())
}

func TestNewWebHookChannelCreateRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewWebHookChannelCreateRequestWithDefaults()
	require.NotNil(t, model)
}

func TestWebHookChannelCreateRequestSetGetCycle(t *testing.T) {
	cfg := openapiclient.HttpConfigAsWebHookChannelCreateRequestConfig(buildExampleHttpConfig())
	model := openapiclient.NewWebHookChannelCreateRequest(exampleWebHookChannelCreateRequestName, exampleWebHookChannelCreateRequestClient, cfg, exampleWebHookChannelCreateRequestEndpoint)
	model.SetChannelName("updated")
	assert.Equal(t, "updated", model.GetChannelName())
	namePtr, ok := model.GetChannelNameOk()
	require.True(t, ok)
	if assert.NotNil(t, namePtr) {
		assert.Equal(t, "updated", *namePtr)
	}
}

func TestWebHookChannelCreateRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleWebHookChannelCreateRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.WebHookChannelCreateRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleWebHookChannelCreateRequestClient, unmarshalled.GetClientid())
	assert.Equal(t, exampleWebHookChannelCreateRequestEndpoint, unmarshalled.GetEndpointId())
}

func TestWebHookChannelCreateRequestToMap(t *testing.T) {
	model := buildExampleWebHookChannelCreateRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "channel_name") {
		assert.Equal(t, exampleWebHookChannelCreateRequestName, m["channel_name"])
	}
}

func TestNullableWebHookChannelCreateRequestGetSet(t *testing.T) {
	base := buildExampleWebHookChannelCreateRequest()
	n := openapiclient.NullableWebHookChannelCreateRequest{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableWebHookChannelCreateRequestUnset(t *testing.T) {
	base := buildExampleWebHookChannelCreateRequest()
	n := openapiclient.NewNullableWebHookChannelCreateRequest(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableWebHookChannelCreateRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleWebHookChannelCreateRequest()
	n := openapiclient.NewNullableWebHookChannelCreateRequest(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableWebHookChannelCreateRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleWebHookChannelCreateRequestName, newN.Get().GetChannelName())
	}
}
