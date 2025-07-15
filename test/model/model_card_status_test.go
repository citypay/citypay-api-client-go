package citypay

import (
	"encoding/json"
	"testing"

	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var exampleCardStatusValue = "ACTIVE"
var exampleCardStatusDefault = true

func buildExampleCardStatus() *openapiclient.CardStatus {
	c := openapiclient.NewCardStatus()
	c.SetCardStatus(exampleCardStatusValue)
	c.SetDefault(exampleCardStatusDefault)
	return c
}

func TestNewCardStatus(t *testing.T) {
	model := openapiclient.NewCardStatus()
	require.NotNil(t, model)
	assert.False(t, model.HasCardStatus())
	assert.False(t, model.HasDefault())
}

func TestNewCardStatusWithDefaults(t *testing.T) {
	model := openapiclient.NewCardStatusWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasCardStatus())
}

func TestCardStatusSetGetCycle(t *testing.T) {
	model := openapiclient.NewCardStatus()
	model.SetCardStatus(exampleCardStatusValue)
	assert.True(t, model.HasCardStatus())
	assert.Equal(t, exampleCardStatusValue, model.GetCardStatus())

	model.SetDefault(exampleCardStatusDefault)
	assert.True(t, model.HasDefault())
	assert.Equal(t, exampleCardStatusDefault, model.GetDefault())
}

func TestCardStatusJSONRoundTrip(t *testing.T) {
	model := buildExampleCardStatus()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.CardStatus
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleCardStatusValue, unmarshalled.GetCardStatus())
	assert.True(t, unmarshalled.GetDefault())
}

func TestCardStatusToMap(t *testing.T) {
	model := buildExampleCardStatus()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "card_status") {
		assert.Equal(t, &exampleCardStatusValue, m["card_status"])
	}
	if assert.Contains(t, m, "default") {
		assert.Equal(t, &exampleCardStatusDefault, m["default"])
	}
}

func TestNullableCardStatusGetSet(t *testing.T) {
	base := buildExampleCardStatus()
	n := openapiclient.NullableCardStatus{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableCardStatusUnset(t *testing.T) {
	base := buildExampleCardStatus()
	n := openapiclient.NewNullableCardStatus(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableCardStatusJSONRoundTrip(t *testing.T) {
	base := buildExampleCardStatus()
	n := openapiclient.NewNullableCardStatus(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableCardStatus
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleCardStatusValue, newN.Get().GetCardStatus())
	}
}
