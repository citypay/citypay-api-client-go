package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var examplePaMd = "md123"
var examplePaPares = "paresval"

func buildExamplePaResAuthRequest() *openapiclient.PaResAuthRequest {
	p := openapiclient.NewPaResAuthRequest(examplePaMd, examplePaPares)
	return p
}

func TestNewPaResAuthRequest(t *testing.T) {
	model := openapiclient.NewPaResAuthRequest(examplePaMd, examplePaPares)
	require.NotNil(t, model)
	assert.Equal(t, examplePaMd, model.GetMd())
	assert.Equal(t, examplePaPares, model.GetPares())
}

func TestNewPaResAuthRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewPaResAuthRequestWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, "", model.GetMd())
	assert.Equal(t, "", model.GetPares())
}

func TestPaResAuthRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewPaResAuthRequest(examplePaMd, examplePaPares)
	model.SetMd("new")
	assert.Equal(t, "new", model.GetMd())
	if val, ok := model.GetMdOk(); assert.True(t, ok) {
		assert.Equal(t, "new", *val)
	}

	model.SetPares("p")
	assert.Equal(t, "p", model.GetPares())
	if val, ok := model.GetParesOk(); assert.True(t, ok) {
		assert.Equal(t, "p", *val)
	}
}

func TestPaResAuthRequestJSONRoundTrip(t *testing.T) {
	model := buildExamplePaResAuthRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.PaResAuthRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, examplePaMd, unmarshalled.GetMd())
	assert.Equal(t, examplePaPares, unmarshalled.GetPares())
}

func TestPaResAuthRequestToMap(t *testing.T) {
	model := buildExamplePaResAuthRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "md") {
		assert.Equal(t, examplePaMd, m["md"])
	}
	if assert.Contains(t, m, "pares") {
		assert.Equal(t, examplePaPares, m["pares"])
	}
}

func TestNullablePaResAuthRequestGetSet(t *testing.T) {
	base := buildExamplePaResAuthRequest()
	n := openapiclient.NullablePaResAuthRequest{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullablePaResAuthRequestUnset(t *testing.T) {
	base := buildExamplePaResAuthRequest()
	n := openapiclient.NewNullablePaResAuthRequest(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullablePaResAuthRequestJSONRoundTrip(t *testing.T) {
	base := buildExamplePaResAuthRequest()
	n := openapiclient.NewNullablePaResAuthRequest(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullablePaResAuthRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, examplePaMd, newN.Get().GetMd())
	}
}
