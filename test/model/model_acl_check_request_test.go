package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleAclRequestIP = "192.0.2.1"

func TestNewAclCheckRequest(t *testing.T) {
	model := openapiclient.NewAclCheckRequest(exampleAclRequestIP)
	require.NotNil(t, model)
	assert.Equal(t, exampleAclRequestIP, model.GetIp())
}

func TestNewAclCheckRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewAclCheckRequestWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, "", model.GetIp())
}

func TestAclCheckRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewAclCheckRequest(exampleAclRequestIP)
	model.SetIp("203.0.113.5")
	assert.Equal(t, "203.0.113.5", model.GetIp())
	val, ok := model.GetIpOk()
	require.True(t, ok)
	if assert.NotNil(t, val) {
		assert.Equal(t, "203.0.113.5", *val)
	}
}

func TestAclCheckRequestJSONRoundTrip(t *testing.T) {
	model := openapiclient.NewAclCheckRequest(exampleAclRequestIP)
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.AclCheckRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleAclRequestIP, unmarshalled.GetIp())
}

func TestAclCheckRequestToMap(t *testing.T) {
	model := openapiclient.NewAclCheckRequest(exampleAclRequestIP)
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "ip") {
		assert.Equal(t, exampleAclRequestIP, m["ip"])
	}
}

func TestNullableAclCheckRequestGetSet(t *testing.T) {
	base := openapiclient.NewAclCheckRequest(exampleAclRequestIP)
	n := openapiclient.NullableAclCheckRequest{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableAclCheckRequestUnset(t *testing.T) {
	base := openapiclient.NewAclCheckRequest(exampleAclRequestIP)
	n := openapiclient.NewNullableAclCheckRequest(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableAclCheckRequestJSONRoundTrip(t *testing.T) {
	base := openapiclient.NewAclCheckRequest(exampleAclRequestIP)
	n := openapiclient.NewNullableAclCheckRequest(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableAclCheckRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleAclRequestIP, newN.Get().GetIp())
	}
}
