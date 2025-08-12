package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleUIAddressMandatory = true
var exampleUIAutocomplete = "off"
var exampleUIOrdering int32 = 3
var exampleUIPostcode = true

func buildExamplePaylinkUI() openapiclient.PaylinkUI {
	u := openapiclient.NewPaylinkUI()
	u.SetAddressMandatory(exampleUIAddressMandatory)
	u.SetFormAutoComplete(exampleUIAutocomplete)
	u.SetOrdering(exampleUIOrdering)
	u.SetPostcodeMandatory(exampleUIPostcode)
	return *u
}

func TestNewPaylinkUI(t *testing.T) {
	model := openapiclient.NewPaylinkUI()
	require.NotNil(t, model)
	assert.False(t, model.HasOrdering())
}

func TestNewPaylinkUIWithDefaults(t *testing.T) {
	model := openapiclient.NewPaylinkUIWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasPostcodeMandatory())
}

func TestPaylinkUISetGetCycle(t *testing.T) {
	model := openapiclient.NewPaylinkUI()
	model.SetAddressMandatory(exampleUIAddressMandatory)
	assert.True(t, model.HasAddressMandatory())
	assert.Equal(t, exampleUIAddressMandatory, model.GetAddressMandatory())

	model.SetFormAutoComplete(exampleUIAutocomplete)
	assert.True(t, model.HasFormAutoComplete())
	assert.Equal(t, exampleUIAutocomplete, model.GetFormAutoComplete())

	model.SetOrdering(exampleUIOrdering)
	assert.True(t, model.HasOrdering())
	assert.Equal(t, exampleUIOrdering, model.GetOrdering())

	model.SetPostcodeMandatory(exampleUIPostcode)
	assert.True(t, model.HasPostcodeMandatory())
	assert.Equal(t, exampleUIPostcode, model.GetPostcodeMandatory())
}

func TestPaylinkUIJSONRoundTrip(t *testing.T) {
	model := buildExamplePaylinkUI()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.PaylinkUI
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasOrdering())
	assert.Equal(t, exampleUIOrdering, unmarshalled.GetOrdering())
}

func TestPaylinkUIToMap(t *testing.T) {
	model := buildExamplePaylinkUI()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "address_mandatory") {
		assert.Equal(t, &exampleUIAddressMandatory, m["address_mandatory"])
	}
}

func TestNullablePaylinkUIGetSet(t *testing.T) {
	base := buildExamplePaylinkUI()
	n := openapiclient.NullablePaylinkUI{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkUIUnset(t *testing.T) {
	base := buildExamplePaylinkUI()
	n := openapiclient.NewNullablePaylinkUI(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullablePaylinkUIJSONRoundTrip(t *testing.T) {
	base := buildExamplePaylinkUI()
	n := openapiclient.NewNullablePaylinkUI(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullablePaylinkUI
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleUIAutocomplete, newN.Get().GetFormAutoComplete())
	}
}
