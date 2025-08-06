package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleWebHookChannelDeleteRequestId = "channel123"

func buildExampleWebHookChannelDeleteRequest() *openapiclient.WebHookChannelDeleteRequest {
	m := openapiclient.NewWebHookChannelDeleteRequest(exampleWebHookChannelDeleteRequestId)
	return m
}

func TestNewWebHookChannelDeleteRequest(t *testing.T) {
	model := openapiclient.NewWebHookChannelDeleteRequest(exampleWebHookChannelDeleteRequestId)
	require.NotNil(t, model)
	assert.Equal(t, exampleWebHookChannelDeleteRequestId, model.GetWebChannelId())
}

func TestNewWebHookChannelDeleteRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewWebHookChannelDeleteRequestWithDefaults()
	require.NotNil(t, model)
}

func TestWebHookChannelDeleteRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewWebHookChannelDeleteRequest(exampleWebHookChannelDeleteRequestId)
	model.SetWebChannelId("updated")
	assert.Equal(t, "updated", model.GetWebChannelId())
	idPtr, ok := model.GetWebChannelIdOk()
	require.True(t, ok)
	if assert.NotNil(t, idPtr) {
		assert.Equal(t, "updated", *idPtr)
	}
}

func TestWebHookChannelDeleteRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleWebHookChannelDeleteRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.WebHookChannelDeleteRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleWebHookChannelDeleteRequestId, unmarshalled.GetWebChannelId())
}

func TestWebHookChannelDeleteRequestToMap(t *testing.T) {
	model := buildExampleWebHookChannelDeleteRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "web_channel_id") {
		assert.Equal(t, exampleWebHookChannelDeleteRequestId, m["web_channel_id"])
	}
}

func TestNullableWebHookChannelDeleteRequestGetSet(t *testing.T) {
	base := buildExampleWebHookChannelDeleteRequest()
	n := openapiclient.NullableWebHookChannelDeleteRequest{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableWebHookChannelDeleteRequestUnset(t *testing.T) {
	base := buildExampleWebHookChannelDeleteRequest()
	n := openapiclient.NewNullableWebHookChannelDeleteRequest(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableWebHookChannelDeleteRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleWebHookChannelDeleteRequest()
	n := openapiclient.NewNullableWebHookChannelDeleteRequest(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableWebHookChannelDeleteRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleWebHookChannelDeleteRequestId, newN.Get().GetWebChannelId())
	}
}
