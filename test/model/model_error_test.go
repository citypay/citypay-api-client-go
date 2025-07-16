package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var exampleErrCode = "E001"
var exampleErrContext = "CTX"
var exampleErrId = "ID123"
var exampleErrMsg = "message"
var exampleErrDate = time.Date(2024, time.June, 1, 0, 0, 0, 0, time.UTC)

func buildExampleError() *openapiclient.Error {
	e := openapiclient.NewError()
	e.SetCode(exampleErrCode)
	e.SetContext(exampleErrContext)
	e.SetIdentifier(exampleErrId)
	e.SetMessage(exampleErrMsg)
	e.SetResponseDt(exampleErrDate)
	return e
}

func TestNewError(t *testing.T) {
	model := openapiclient.NewError()
	require.NotNil(t, model)
	assert.False(t, model.HasCode())
}

func TestNewErrorWithDefaults(t *testing.T) {
	model := openapiclient.NewErrorWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasCode())
}

func TestErrorSetGetCycle(t *testing.T) {
	model := openapiclient.NewError()
	model.SetCode(exampleErrCode)
	assert.True(t, model.HasCode())
	assert.Equal(t, exampleErrCode, model.GetCode())

	model.SetContext(exampleErrContext)
	assert.True(t, model.HasContext())
	assert.Equal(t, exampleErrContext, model.GetContext())

	model.SetIdentifier(exampleErrId)
	assert.True(t, model.HasIdentifier())
	assert.Equal(t, exampleErrId, model.GetIdentifier())

	model.SetMessage(exampleErrMsg)
	assert.True(t, model.HasMessage())
	assert.Equal(t, exampleErrMsg, model.GetMessage())

	model.SetResponseDt(exampleErrDate)
	assert.True(t, model.HasResponseDt())
	assert.Equal(t, exampleErrDate, model.GetResponseDt())
}

func TestErrorJSONRoundTrip(t *testing.T) {
	model := buildExampleError()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.Error
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasCode())
	assert.Equal(t, exampleErrCode, unmarshalled.GetCode())
}

func TestErrorToMap(t *testing.T) {
	model := buildExampleError()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "code") {
		assert.Equal(t, &exampleErrCode, m["code"])
	}
	if assert.Contains(t, m, "response_dt") {
		assert.Equal(t, &exampleErrDate, m["response_dt"])
	}
}

func TestNullableErrorGetSet(t *testing.T) {
	base := buildExampleError()
	n := openapiclient.NullableError{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableErrorUnset(t *testing.T) {
	base := buildExampleError()
	n := openapiclient.NewNullableError(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableErrorJSONRoundTrip(t *testing.T) {
	base := buildExampleError()
	n := openapiclient.NewNullableError(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableError
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleErrCode, newN.Get().GetCode())
	}
}
