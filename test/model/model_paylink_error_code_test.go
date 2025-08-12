package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var examplePaylinkErrCode = "E1"
var examplePaylinkErrMsg = "error"

func buildExamplePaylinkErrorCode() openapiclient.PaylinkErrorCode {
	e := openapiclient.NewPaylinkErrorCode(examplePaylinkErrCode, examplePaylinkErrMsg)
	return *e
}

func TestNewPaylinkErrorCode(t *testing.T) {
	model := openapiclient.NewPaylinkErrorCode(examplePaylinkErrCode, examplePaylinkErrMsg)
	require.NotNil(t, model)
	assert.Equal(t, examplePaylinkErrCode, model.GetCode())
	assert.Equal(t, examplePaylinkErrMsg, model.GetMsg())
}

func TestNewPaylinkErrorCodeWithDefaults(t *testing.T) {
	model := openapiclient.NewPaylinkErrorCodeWithDefaults()
	require.NotNil(t, model)
	assert.NotNil(t, model)
}

func TestPaylinkErrorCodeSetGetCycle(t *testing.T) {
	model := openapiclient.NewPaylinkErrorCode(examplePaylinkErrCode, examplePaylinkErrMsg)
	model.SetCode("E2")
	assert.Equal(t, "E2", model.GetCode())

	model.SetMsg("msg")
	assert.Equal(t, "msg", model.GetMsg())
}

func TestPaylinkErrorCodeJSONRoundTrip(t *testing.T) {
	model := buildExamplePaylinkErrorCode()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.PaylinkErrorCode
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, examplePaylinkErrMsg, unmarshalled.GetMsg())
}

func TestPaylinkErrorCodeToMap(t *testing.T) {
	model := buildExamplePaylinkErrorCode()
	m, err := model.ToMap()
	require.NoError(t, err)
	assert.Equal(t, examplePaylinkErrCode, m["code"])
	assert.Equal(t, examplePaylinkErrMsg, m["msg"])
}

func TestNullablePaylinkErrorCodeGetSet(t *testing.T) {
	base := buildExamplePaylinkErrorCode()
	n := openapiclient.NullablePaylinkErrorCode{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkErrorCodeUnset(t *testing.T) {
	base := buildExamplePaylinkErrorCode()
	n := openapiclient.NewNullablePaylinkErrorCode(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullablePaylinkErrorCodeJSONRoundTrip(t *testing.T) {
	base := buildExamplePaylinkErrorCode()
	n := openapiclient.NewNullablePaylinkErrorCode(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullablePaylinkErrorCode
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, examplePaylinkErrCode, newN.Get().GetCode())
	}
}
