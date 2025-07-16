package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleAddr1 = "1 Road"
var exampleAddr2 = "Suite 2"
var exampleAddr3 = "Floor 3"
var exampleArea = "City"
var exampleCountry = "GB"
var exampleLabel = "Home"
var examplePostcode = "AB12"

func buildExamplePaylinkAddress() openapiclient.PaylinkAddress {
	p := openapiclient.NewPaylinkAddress()
	p.SetAddress1(exampleAddr1)
	p.SetAddress2(exampleAddr2)
	p.SetAddress3(exampleAddr3)
	p.SetArea(exampleArea)
	p.SetCountry(exampleCountry)
	p.SetLabel(exampleLabel)
	p.SetPostcode(examplePostcode)
	return *p
}

func TestNewPaylinkAddress(t *testing.T) {
	model := openapiclient.NewPaylinkAddress()
	require.NotNil(t, model)
	assert.False(t, model.HasAddress1())
	assert.False(t, model.HasAddress2())
	assert.False(t, model.HasAddress3())
	assert.False(t, model.HasArea())
	assert.False(t, model.HasCountry())
	assert.False(t, model.HasLabel())
	assert.False(t, model.HasPostcode())
}

func TestNewPaylinkAddressWithDefaults(t *testing.T) {
	model := openapiclient.NewPaylinkAddressWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasAddress1())
	assert.False(t, model.HasAddress2())
	assert.False(t, model.HasAddress3())
	assert.False(t, model.HasArea())
	assert.False(t, model.HasCountry())
	assert.False(t, model.HasLabel())
	assert.False(t, model.HasPostcode())
}

func TestPaylinkAddressSetGetCycle(t *testing.T) {
	model := openapiclient.NewPaylinkAddress()
	model.SetAddress1(exampleAddr1)
	assert.True(t, model.HasAddress1())
	assert.Equal(t, exampleAddr1, model.GetAddress1())
	if val, ok := model.GetAddress1Ok(); assert.True(t, ok) {
		assert.Equal(t, exampleAddr1, *val)
	}

	model.SetAddress2(exampleAddr2)
	assert.True(t, model.HasAddress2())
	assert.Equal(t, exampleAddr2, model.GetAddress2())

	model.SetAddress3(exampleAddr3)
	assert.True(t, model.HasAddress3())
	assert.Equal(t, exampleAddr3, model.GetAddress3())

	model.SetArea(exampleArea)
	assert.True(t, model.HasArea())
	assert.Equal(t, exampleArea, model.GetArea())

	model.SetCountry(exampleCountry)
	assert.True(t, model.HasCountry())
	assert.Equal(t, exampleCountry, model.GetCountry())

	model.SetLabel(exampleLabel)
	assert.True(t, model.HasLabel())
	assert.Equal(t, exampleLabel, model.GetLabel())

	model.SetPostcode(examplePostcode)
	assert.True(t, model.HasPostcode())
	assert.Equal(t, examplePostcode, model.GetPostcode())
}

func TestPaylinkAddressJSONRoundTrip(t *testing.T) {
	model := buildExamplePaylinkAddress()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.PaylinkAddress
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasCountry())
	assert.Equal(t, exampleAddr2, unmarshalled.GetAddress2())
}

func TestPaylinkAddressToMap(t *testing.T) {
	model := buildExamplePaylinkAddress()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "address1") {
		assert.Equal(t, &exampleAddr1, m["address1"])
	}
	if assert.Contains(t, m, "address2") {
		assert.Equal(t, &exampleAddr2, m["address2"])
	}
	if assert.Contains(t, m, "address3") {
		assert.Equal(t, &exampleAddr3, m["address3"])
	}
	if assert.Contains(t, m, "area") {
		assert.Equal(t, &exampleArea, m["area"])
	}
	if assert.Contains(t, m, "country") {
		assert.Equal(t, &exampleCountry, m["country"])
	}
	if assert.Contains(t, m, "label") {
		assert.Equal(t, &exampleLabel, m["label"])
	}
	if assert.Contains(t, m, "postcode") {
		assert.Equal(t, &examplePostcode, m["postcode"])
	}
}

func TestNullablePaylinkAddressGetSet(t *testing.T) {
	base := buildExamplePaylinkAddress()
	n := openapiclient.NullablePaylinkAddress{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkAddressUnset(t *testing.T) {
	base := buildExamplePaylinkAddress()
	n := openapiclient.NewNullablePaylinkAddress(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullablePaylinkAddressJSONRoundTrip(t *testing.T) {
	base := buildExamplePaylinkAddress()
	n := openapiclient.NewNullablePaylinkAddress(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullablePaylinkAddress
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, examplePostcode, newN.Get().GetPostcode())
	}
}
