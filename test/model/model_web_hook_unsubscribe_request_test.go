package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleWebHookUnsubscribeRequestClient = "client1"
var exampleWebHookUnsubscribeRequestId = "whsub123"

func buildExampleWebHookUnsubscribeRequest() *openapiclient.WebHookUnsubscribeRequest {
	m := openapiclient.NewWebHookUnsubscribeRequest(exampleWebHookUnsubscribeRequestClient, exampleWebHookUnsubscribeRequestId)
	return m
}

func TestNewWebHookUnsubscribeRequest(t *testing.T) {
	model := openapiclient.NewWebHookUnsubscribeRequest(exampleWebHookUnsubscribeRequestClient, exampleWebHookUnsubscribeRequestId)
	require.NotNil(t, model)
	assert.Equal(t, exampleWebHookUnsubscribeRequestClient, model.GetClientid())
	assert.Equal(t, exampleWebHookUnsubscribeRequestId, model.GetWebHookId())
}

func TestNewWebHookUnsubscribeRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewWebHookUnsubscribeRequestWithDefaults()
	require.NotNil(t, model)
}

func TestWebHookUnsubscribeRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewWebHookUnsubscribeRequest(exampleWebHookUnsubscribeRequestClient, exampleWebHookUnsubscribeRequestId)
	model.SetClientid("updated")
	assert.Equal(t, "updated", model.GetClientid())
	cidPtr, ok := model.GetClientidOk()
	require.True(t, ok)
	if assert.NotNil(t, cidPtr) {
		assert.Equal(t, "updated", *cidPtr)
	}
}

func TestWebHookUnsubscribeRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleWebHookUnsubscribeRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.WebHookUnsubscribeRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleWebHookUnsubscribeRequestId, unmarshalled.GetWebHookId())
}

func TestWebHookUnsubscribeRequestToMap(t *testing.T) {
	model := buildExampleWebHookUnsubscribeRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "clientid") {
		assert.Equal(t, exampleWebHookUnsubscribeRequestClient, m["clientid"])
	}
}

func TestNullableWebHookUnsubscribeRequestGetSet(t *testing.T) {
	base := buildExampleWebHookUnsubscribeRequest()
	n := openapiclient.NullableWebHookUnsubscribeRequest{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableWebHookUnsubscribeRequestUnset(t *testing.T) {
	base := buildExampleWebHookUnsubscribeRequest()
	n := openapiclient.NewNullableWebHookUnsubscribeRequest(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableWebHookUnsubscribeRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleWebHookUnsubscribeRequest()
	n := openapiclient.NewNullableWebHookUnsubscribeRequest(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableWebHookUnsubscribeRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleWebHookUnsubscribeRequestId, newN.Get().GetWebHookId())
	}
}
