package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleWebHookSubscriptionRequestClient = "client1"
var exampleWebHookSubscriptionRequestChannel = "chan1"

func buildExampleWebHookSubscriptionRequest() *openapiclient.WebHookSubscriptionRequest {
	m := openapiclient.NewWebHookSubscriptionRequest(exampleWebHookSubscriptionRequestClient)
	m.SetChannels([]string{exampleWebHookSubscriptionRequestChannel})
	m.SetTriggers([]string{"payment"})
	return m
}

func TestNewWebHookSubscriptionRequest(t *testing.T) {
	model := openapiclient.NewWebHookSubscriptionRequest(exampleWebHookSubscriptionRequestClient)
	require.NotNil(t, model)
	assert.Equal(t, exampleWebHookSubscriptionRequestClient, model.GetClientid())
	assert.False(t, model.HasChannels())
}

func TestNewWebHookSubscriptionRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewWebHookSubscriptionRequestWithDefaults()
	require.NotNil(t, model)
}

func TestWebHookSubscriptionRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewWebHookSubscriptionRequest(exampleWebHookSubscriptionRequestClient)
	model.SetChannels([]string{exampleWebHookSubscriptionRequestChannel})
	assert.True(t, model.HasChannels())
	assert.Equal(t, []string{exampleWebHookSubscriptionRequestChannel}, model.GetChannels())
}

func TestWebHookSubscriptionRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleWebHookSubscriptionRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.WebHookSubscriptionRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleWebHookSubscriptionRequestClient, unmarshalled.GetClientid())
	assert.True(t, unmarshalled.HasTriggers())
}

func TestWebHookSubscriptionRequestToMap(t *testing.T) {
	model := buildExampleWebHookSubscriptionRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "clientid") {
		assert.Equal(t, exampleWebHookSubscriptionRequestClient, m["clientid"])
	}
}

func TestNullableWebHookSubscriptionRequestGetSet(t *testing.T) {
	base := buildExampleWebHookSubscriptionRequest()
	n := openapiclient.NullableWebHookSubscriptionRequest{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableWebHookSubscriptionRequestUnset(t *testing.T) {
	base := buildExampleWebHookSubscriptionRequest()
	n := openapiclient.NewNullableWebHookSubscriptionRequest(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableWebHookSubscriptionRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleWebHookSubscriptionRequest()
	n := openapiclient.NewNullableWebHookSubscriptionRequest(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableWebHookSubscriptionRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleWebHookSubscriptionRequestClient, newN.Get().GetClientid())
	}
}
