package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleRegisterIpModelIp = "1.2.3.4"
var exampleRegisterIpModelExp int64 = 3600

func buildExampleRegisterIpModel() *openapiclient.RegisterIpModel {
	m := openapiclient.NewRegisterIpModel()
	m.SetIp(exampleRegisterIpModelIp)
	m.SetExp(exampleRegisterIpModelExp)
	return m
}

func TestNewRegisterIpModel(t *testing.T) {
	model := openapiclient.NewRegisterIpModel()
	require.NotNil(t, model)
	assert.False(t, model.HasIp())
}

func TestNewRegisterIpModelWithDefaults(t *testing.T) {
	model := openapiclient.NewRegisterIpModelWithDefaults()
	require.NotNil(t, model)
}

func TestRegisterIpModelSetGetCycle(t *testing.T) {
	model := openapiclient.NewRegisterIpModel()
	model.SetIp(exampleRegisterIpModelIp)
	assert.True(t, model.HasIp())
	assert.Equal(t, exampleRegisterIpModelIp, model.GetIp())
	ipPtr, ok := model.GetIpOk()
	require.True(t, ok)
	if assert.NotNil(t, ipPtr) {
		assert.Equal(t, exampleRegisterIpModelIp, *ipPtr)
	}
}

func TestRegisterIpModelJSONRoundTrip(t *testing.T) {
	model := buildExampleRegisterIpModel()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.RegisterIpModel
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleRegisterIpModelIp, unmarshalled.GetIp())
	assert.Equal(t, exampleRegisterIpModelExp, unmarshalled.GetExp())
}

func TestRegisterIpModelToMap(t *testing.T) {
	model := buildExampleRegisterIpModel()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "ip") {
		assert.Equal(t, &exampleRegisterIpModelIp, m["ip"])
	}
}

func TestNullableRegisterIpModelGetSet(t *testing.T) {
	base := buildExampleRegisterIpModel()
	n := openapiclient.NullableRegisterIpModel{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableRegisterIpModelUnset(t *testing.T) {
	base := buildExampleRegisterIpModel()
	n := openapiclient.NewNullableRegisterIpModel(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableRegisterIpModelJSONRoundTrip(t *testing.T) {
	base := buildExampleRegisterIpModel()
	n := openapiclient.NewNullableRegisterIpModel(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableRegisterIpModel
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleRegisterIpModelIp, newN.Get().GetIp())
	}
}
