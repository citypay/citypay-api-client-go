package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleHttpConfigURL = "https://example.com"
var exampleHttpConfigMethod = "POST"
var exampleHttpConfigConnectTimeout int32 = 500

func buildExampleHttpConfig() *openapiclient.HttpConfig {
	m := openapiclient.NewHttpConfig(exampleHttpConfigURL)
	m.SetMethod(exampleHttpConfigMethod)
	m.SetConnectTimeout(exampleHttpConfigConnectTimeout)
	return m
}

func TestNewHttpConfig(t *testing.T) {
	model := openapiclient.NewHttpConfig(exampleHttpConfigURL)
	require.NotNil(t, model)
	assert.Equal(t, exampleHttpConfigURL, model.GetUrl())
	assert.False(t, model.HasMethod())
}

func TestNewHttpConfigWithDefaults(t *testing.T) {
	model := openapiclient.NewHttpConfigWithDefaults()
	require.NotNil(t, model)
}

func TestHttpConfigSetGetCycle(t *testing.T) {
	model := openapiclient.NewHttpConfig(exampleHttpConfigURL)
	model.SetMethod(exampleHttpConfigMethod)
	assert.True(t, model.HasMethod())
	assert.Equal(t, exampleHttpConfigMethod, model.GetMethod())
	methPtr, ok := model.GetMethodOk()
	require.True(t, ok)
	if assert.NotNil(t, methPtr) {
		assert.Equal(t, exampleHttpConfigMethod, *methPtr)
	}
}

func TestHttpConfigJSONRoundTrip(t *testing.T) {
	model := buildExampleHttpConfig()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.HttpConfig
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleHttpConfigURL, unmarshalled.GetUrl())
	assert.Equal(t, exampleHttpConfigMethod, unmarshalled.GetMethod())
}

func TestHttpConfigToMap(t *testing.T) {
	model := buildExampleHttpConfig()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "url") {
		assert.Equal(t, exampleHttpConfigURL, m["url"])
	}
}

func TestNullableHttpConfigGetSet(t *testing.T) {
	base := buildExampleHttpConfig()
	n := openapiclient.NullableHttpConfig{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableHttpConfigUnset(t *testing.T) {
	base := buildExampleHttpConfig()
	n := openapiclient.NewNullableHttpConfig(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableHttpConfigJSONRoundTrip(t *testing.T) {
	base := buildExampleHttpConfig()
	n := openapiclient.NewNullableHttpConfig(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableHttpConfig
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleHttpConfigURL, newN.Get().GetUrl())
	}
}
