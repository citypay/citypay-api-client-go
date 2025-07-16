package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var examplePingID = "ping1"

func buildExamplePing() openapiclient.Ping {
	p := openapiclient.NewPing()
	p.SetIdentifier(examplePingID)
	return *p
}

func TestNewPing(t *testing.T) {
	model := openapiclient.NewPing()
	require.NotNil(t, model)
	assert.False(t, model.HasIdentifier())
}

func TestNewPingWithDefaults(t *testing.T) {
	model := openapiclient.NewPingWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasIdentifier())
}

func TestPingSetGetCycle(t *testing.T) {
	model := openapiclient.NewPing()
	model.SetIdentifier(examplePingID)
	assert.True(t, model.HasIdentifier())
	assert.Equal(t, examplePingID, model.GetIdentifier())
}

func TestPingJSONRoundTrip(t *testing.T) {
	model := buildExamplePing()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.Ping
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasIdentifier())
	assert.Equal(t, examplePingID, unmarshalled.GetIdentifier())
}

func TestPingToMap(t *testing.T) {
	model := buildExamplePing()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "identifier") {
		assert.Equal(t, &examplePingID, m["identifier"])
	}
}

func TestNullablePingGetSet(t *testing.T) {
	base := buildExamplePing()
	n := openapiclient.NullablePing{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullablePingUnset(t *testing.T) {
	base := buildExamplePing()
	n := openapiclient.NewNullablePing(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullablePingJSONRoundTrip(t *testing.T) {
	base := buildExamplePing()
	n := openapiclient.NewNullablePing(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullablePing
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, examplePingID, newN.Get().GetIdentifier())
	}
}
