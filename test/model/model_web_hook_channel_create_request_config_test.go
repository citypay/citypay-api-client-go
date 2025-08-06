package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func buildExampleWebHookChannelCreateRequestConfig() *openapiclient.WebHookChannelCreateRequestConfig {
	cfg := buildExampleHttpConfig()
	wrapped := openapiclient.HttpConfigAsWebHookChannelCreateRequestConfig(cfg)
	return &wrapped
}

func TestWebHookChannelCreateRequestConfigJSONRoundTrip(t *testing.T) {
	cfg := buildExampleWebHookChannelCreateRequestConfig()
	data, err := json.Marshal(cfg)
	require.NoError(t, err)

	var unmarshalled openapiclient.WebHookChannelCreateRequestConfig
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.NotNil(t, unmarshalled.GetActualInstance())
}

func TestNullableWebHookChannelCreateRequestConfigGetSet(t *testing.T) {
	base := buildExampleWebHookChannelCreateRequestConfig()
	n := openapiclient.NullableWebHookChannelCreateRequestConfig{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableWebHookChannelCreateRequestConfigUnset(t *testing.T) {
	base := buildExampleWebHookChannelCreateRequestConfig()
	n := openapiclient.NewNullableWebHookChannelCreateRequestConfig(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableWebHookChannelCreateRequestConfigJSONRoundTrip(t *testing.T) {
	base := buildExampleWebHookChannelCreateRequestConfig()
	n := openapiclient.NewNullableWebHookChannelCreateRequestConfig(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableWebHookChannelCreateRequestConfig
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	assert.NotNil(t, newN.Get())
}
