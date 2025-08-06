package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleCardTokenisationResponseToken = "tok_123"
var exampleCardTokenisationResponseLast4 = "1111"

func buildExampleCardTokenisationResponse() *openapiclient.CardTokenisationResponse {
	m := openapiclient.NewCardTokenisationResponse(exampleCardTokenisationResponseToken)
	m.SetLast4digits(exampleCardTokenisationResponseLast4)
	return m
}

func TestNewCardTokenisationResponse(t *testing.T) {
	model := openapiclient.NewCardTokenisationResponse(exampleCardTokenisationResponseToken)
	require.NotNil(t, model)
	assert.Equal(t, exampleCardTokenisationResponseToken, model.GetCpCardToken())
	assert.False(t, model.HasLast4digits())
}

func TestNewCardTokenisationResponseWithDefaults(t *testing.T) {
	model := openapiclient.NewCardTokenisationResponseWithDefaults()
	require.NotNil(t, model)
}

func TestCardTokenisationResponseSetGetCycle(t *testing.T) {
	model := openapiclient.NewCardTokenisationResponse(exampleCardTokenisationResponseToken)
	model.SetLast4digits(exampleCardTokenisationResponseLast4)
	assert.True(t, model.HasLast4digits())
	assert.Equal(t, exampleCardTokenisationResponseLast4, model.GetLast4digits())
	l4, ok := model.GetLast4digitsOk()
	require.True(t, ok)
	if assert.NotNil(t, l4) {
		assert.Equal(t, exampleCardTokenisationResponseLast4, *l4)
	}
}

func TestCardTokenisationResponseJSONRoundTrip(t *testing.T) {
	model := buildExampleCardTokenisationResponse()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.CardTokenisationResponse
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleCardTokenisationResponseToken, unmarshalled.GetCpCardToken())
	assert.Equal(t, exampleCardTokenisationResponseLast4, unmarshalled.GetLast4digits())
}

func TestCardTokenisationResponseToMap(t *testing.T) {
	model := buildExampleCardTokenisationResponse()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "cp_card_token") {
		assert.Equal(t, exampleCardTokenisationResponseToken, m["cp_card_token"])
	}
}

func TestNullableCardTokenisationResponseGetSet(t *testing.T) {
	base := buildExampleCardTokenisationResponse()
	n := openapiclient.NullableCardTokenisationResponse{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableCardTokenisationResponseUnset(t *testing.T) {
	base := buildExampleCardTokenisationResponse()
	n := openapiclient.NewNullableCardTokenisationResponse(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableCardTokenisationResponseJSONRoundTrip(t *testing.T) {
	base := buildExampleCardTokenisationResponse()
	n := openapiclient.NewNullableCardTokenisationResponse(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableCardTokenisationResponse
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleCardTokenisationResponseToken, newN.Get().GetCpCardToken())
	}
}
