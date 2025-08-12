package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleAclName = "Public"
var exampleAclCache = true
var exampleAclIP = "192.0.2.1"
var exampleAclProvider = "cloud"

func buildExampleAclCheckResponse() *openapiclient.AclCheckResponseModel {
	m := openapiclient.NewAclCheckResponseModel()
	m.SetAcl(exampleAclName)
	m.SetCache(exampleAclCache)
	m.SetIp(exampleAclIP)
	m.SetProvider(exampleAclProvider)
	return m
}

func TestNewAclCheckResponseModel(t *testing.T) {
	model := openapiclient.NewAclCheckResponseModel()
	require.NotNil(t, model)
	assert.False(t, model.HasAcl())
	assert.False(t, model.HasCache())
	assert.False(t, model.HasIp())
	assert.False(t, model.HasProvider())
}

func TestNewAclCheckResponseModelWithDefaults(t *testing.T) {
	model := openapiclient.NewAclCheckResponseModelWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasAcl())
	assert.False(t, model.HasCache())
	assert.False(t, model.HasIp())
	assert.False(t, model.HasProvider())
}

func TestAclCheckResponseModelSetGetCycle(t *testing.T) {
	model := openapiclient.NewAclCheckResponseModel()
	model.SetAcl(exampleAclName)
	assert.True(t, model.HasAcl())
	assert.Equal(t, exampleAclName, model.GetAcl())

	model.SetCache(exampleAclCache)
	assert.True(t, model.HasCache())
	assert.Equal(t, exampleAclCache, model.GetCache())

	model.SetIp(exampleAclIP)
	assert.True(t, model.HasIp())
	assert.Equal(t, exampleAclIP, model.GetIp())

	model.SetProvider(exampleAclProvider)
	assert.True(t, model.HasProvider())
	assert.Equal(t, exampleAclProvider, model.GetProvider())
}

func TestAclCheckResponseModelJSONRoundTrip(t *testing.T) {
	model := buildExampleAclCheckResponse()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.AclCheckResponseModel
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleAclName, unmarshalled.GetAcl())
	assert.Equal(t, exampleAclCache, unmarshalled.GetCache())
	assert.Equal(t, exampleAclIP, unmarshalled.GetIp())
	assert.Equal(t, exampleAclProvider, unmarshalled.GetProvider())
}

func TestAclCheckResponseModelToMap(t *testing.T) {
	model := buildExampleAclCheckResponse()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "acl") {
		assert.Equal(t, &exampleAclName, m["acl"])
	}
	if assert.Contains(t, m, "cache") {
		assert.Equal(t, &exampleAclCache, m["cache"])
	}
	if assert.Contains(t, m, "ip") {
		assert.Equal(t, &exampleAclIP, m["ip"])
	}
	if assert.Contains(t, m, "provider") {
		assert.Equal(t, &exampleAclProvider, m["provider"])
	}
}

func TestNullableAclCheckResponseModelGetSet(t *testing.T) {
	base := buildExampleAclCheckResponse()
	n := openapiclient.NullableAclCheckResponseModel{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableAclCheckResponseModelUnset(t *testing.T) {
	base := buildExampleAclCheckResponse()
	n := openapiclient.NewNullableAclCheckResponseModel(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableAclCheckResponseModelJSONRoundTrip(t *testing.T) {
	base := buildExampleAclCheckResponse()
	n := openapiclient.NewNullableAclCheckResponseModel(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableAclCheckResponseModel
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleAclName, newN.Get().GetAcl())
		assert.Equal(t, exampleAclCache, newN.Get().GetCache())
		assert.Equal(t, exampleAclIP, newN.Get().GetIp())
		assert.Equal(t, exampleAclProvider, newN.Get().GetProvider())
	}
}
