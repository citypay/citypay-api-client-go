package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleAckCode = "200"
var exampleAckContext = "ctx1"
var exampleAckIdentifier = "id1"
var exampleAckMessage = "OK"

func buildExampleAcknowledgement() *openapiclient.Acknowledgement {
	a := openapiclient.NewAcknowledgement()
	a.SetCode(exampleAckCode)
	a.SetContext(exampleAckContext)
	a.SetIdentifier(exampleAckIdentifier)
	a.SetMessage(exampleAckMessage)
	return a
}

func TestNewAcknowledgement(t *testing.T) {
	model := openapiclient.NewAcknowledgement()
	require.NotNil(t, model)
	assert.False(t, model.HasCode())
	assert.False(t, model.HasContext())
	assert.False(t, model.HasIdentifier())
	assert.False(t, model.HasMessage())
}

func TestNewAcknowledgementWithDefaults(t *testing.T) {
	model := openapiclient.NewAcknowledgementWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasCode())
	assert.False(t, model.HasContext())
	assert.False(t, model.HasIdentifier())
	assert.False(t, model.HasMessage())
}

func TestAcknowledgementSetGetCycle(t *testing.T) {
	model := openapiclient.NewAcknowledgement()
	model.SetCode(exampleAckCode)
	assert.True(t, model.HasCode())
	assert.Equal(t, exampleAckCode, model.GetCode())
	val, ok := model.GetCodeOk()
	require.True(t, ok)
	if assert.NotNil(t, val) {
		assert.Equal(t, exampleAckCode, *val)
	}

	model.SetContext(exampleAckContext)
	assert.True(t, model.HasContext())
	assert.Equal(t, exampleAckContext, model.GetContext())

	model.SetIdentifier(exampleAckIdentifier)
	assert.True(t, model.HasIdentifier())
	assert.Equal(t, exampleAckIdentifier, model.GetIdentifier())

	model.SetMessage(exampleAckMessage)
	assert.True(t, model.HasMessage())
	assert.Equal(t, exampleAckMessage, model.GetMessage())
}

func TestAcknowledgementJSONRoundTrip(t *testing.T) {
	model := buildExampleAcknowledgement()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.Acknowledgement
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleAckCode, unmarshalled.GetCode())
	assert.Equal(t, exampleAckContext, unmarshalled.GetContext())
	assert.Equal(t, exampleAckIdentifier, unmarshalled.GetIdentifier())
	assert.Equal(t, exampleAckMessage, unmarshalled.GetMessage())
}

func TestAcknowledgementToMap(t *testing.T) {
	model := buildExampleAcknowledgement()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "code") {
		assert.Equal(t, &exampleAckCode, m["code"])
	}
	if assert.Contains(t, m, "context") {
		assert.Equal(t, &exampleAckContext, m["context"])
	}
	if assert.Contains(t, m, "identifier") {
		assert.Equal(t, &exampleAckIdentifier, m["identifier"])
	}
	if assert.Contains(t, m, "message") {
		assert.Equal(t, &exampleAckMessage, m["message"])
	}
}

func TestNullableAcknowledgementGetSet(t *testing.T) {
	base := buildExampleAcknowledgement()
	n := openapiclient.NullableAcknowledgement{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableAcknowledgementUnset(t *testing.T) {
	base := buildExampleAcknowledgement()
	n := openapiclient.NewNullableAcknowledgement(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableAcknowledgementJSONRoundTrip(t *testing.T) {
	base := buildExampleAcknowledgement()
	n := openapiclient.NewNullableAcknowledgement(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableAcknowledgement
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleAckCode, newN.Get().GetCode())
		assert.Equal(t, exampleAckContext, newN.Get().GetContext())
		assert.Equal(t, exampleAckIdentifier, newN.Get().GetIdentifier())
		assert.Equal(t, exampleAckMessage, newN.Get().GetMessage())
	}
}
