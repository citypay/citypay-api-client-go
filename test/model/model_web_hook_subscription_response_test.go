package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleWebHookSubscriptionResponseId = "whsub123"

func buildExampleWebHookSubscriptionResponse() *openapiclient.WebHookSubscriptionResponse {
	m := openapiclient.NewWebHookSubscriptionResponse()
	m.SetWebHookId(exampleWebHookSubscriptionResponseId)
	return m
}

func TestNewWebHookSubscriptionResponse(t *testing.T) {
	model := openapiclient.NewWebHookSubscriptionResponse()
	require.NotNil(t, model)
	assert.False(t, model.HasWebHookId())
}

func TestNewWebHookSubscriptionResponseWithDefaults(t *testing.T) {
	model := openapiclient.NewWebHookSubscriptionResponseWithDefaults()
	require.NotNil(t, model)
}

func TestWebHookSubscriptionResponseSetGetCycle(t *testing.T) {
	model := openapiclient.NewWebHookSubscriptionResponse()
	model.SetWebHookId(exampleWebHookSubscriptionResponseId)
	assert.True(t, model.HasWebHookId())
	assert.Equal(t, exampleWebHookSubscriptionResponseId, model.GetWebHookId())
	idPtr, ok := model.GetWebHookIdOk()
	require.True(t, ok)
	if assert.NotNil(t, idPtr) {
		assert.Equal(t, exampleWebHookSubscriptionResponseId, *idPtr)
	}
}

func TestWebHookSubscriptionResponseJSONRoundTrip(t *testing.T) {
	model := buildExampleWebHookSubscriptionResponse()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.WebHookSubscriptionResponse
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleWebHookSubscriptionResponseId, unmarshalled.GetWebHookId())
}

func TestWebHookSubscriptionResponseToMap(t *testing.T) {
	model := buildExampleWebHookSubscriptionResponse()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "web_hook_id") {
		assert.Equal(t, &exampleWebHookSubscriptionResponseId, m["web_hook_id"])
	}
}

func TestNullableWebHookSubscriptionResponseGetSet(t *testing.T) {
	base := buildExampleWebHookSubscriptionResponse()
	n := openapiclient.NullableWebHookSubscriptionResponse{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableWebHookSubscriptionResponseUnset(t *testing.T) {
	base := buildExampleWebHookSubscriptionResponse()
	n := openapiclient.NewNullableWebHookSubscriptionResponse(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableWebHookSubscriptionResponseJSONRoundTrip(t *testing.T) {
	base := buildExampleWebHookSubscriptionResponse()
	n := openapiclient.NewNullableWebHookSubscriptionResponse(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableWebHookSubscriptionResponse
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleWebHookSubscriptionResponseId, newN.Get().GetWebHookId())
	}
}
