package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleCardTokenisationRequestNumber = "4111111111111111"
var exampleCardTokenisationRequestCsc = "123"
var exampleCardTokenisationRequestExpMonth int32 = 12
var exampleCardTokenisationRequestExpYear int32 = 2030

func buildExampleCardTokenisationRequest() *openapiclient.CardTokenisationRequest {
	m := openapiclient.NewCardTokenisationRequest()
	m.SetCardnumber(exampleCardTokenisationRequestNumber)
	m.SetCsc(exampleCardTokenisationRequestCsc)
	m.SetExpmonth(exampleCardTokenisationRequestExpMonth)
	m.SetExpyear(exampleCardTokenisationRequestExpYear)
	return m
}

func TestNewCardTokenisationRequest(t *testing.T) {
	model := openapiclient.NewCardTokenisationRequest()
	require.NotNil(t, model)
	assert.False(t, model.HasCardnumber())
}

func TestNewCardTokenisationRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewCardTokenisationRequestWithDefaults()
	require.NotNil(t, model)
}

func TestCardTokenisationRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewCardTokenisationRequest()
	model.SetCsc(exampleCardTokenisationRequestCsc)
	assert.True(t, model.HasCsc())
	assert.Equal(t, exampleCardTokenisationRequestCsc, model.GetCsc())
	cscPtr, ok := model.GetCscOk()
	require.True(t, ok)
	if assert.NotNil(t, cscPtr) {
		assert.Equal(t, exampleCardTokenisationRequestCsc, *cscPtr)
	}
}

func TestCardTokenisationRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleCardTokenisationRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.CardTokenisationRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleCardTokenisationRequestNumber, unmarshalled.GetCardnumber())
	assert.Equal(t, exampleCardTokenisationRequestExpYear, unmarshalled.GetExpyear())
}

func TestCardTokenisationRequestToMap(t *testing.T) {
	model := buildExampleCardTokenisationRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "cardnumber") {
		assert.Equal(t, &exampleCardTokenisationRequestNumber, m["cardnumber"])
	}
}

func TestNullableCardTokenisationRequestGetSet(t *testing.T) {
	base := buildExampleCardTokenisationRequest()
	n := openapiclient.NullableCardTokenisationRequest{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableCardTokenisationRequestUnset(t *testing.T) {
	base := buildExampleCardTokenisationRequest()
	n := openapiclient.NewNullableCardTokenisationRequest(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableCardTokenisationRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleCardTokenisationRequest()
	n := openapiclient.NewNullableCardTokenisationRequest(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableCardTokenisationRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleCardTokenisationRequestNumber, newN.Get().GetCardnumber())
	}
}
