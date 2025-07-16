package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleMerchantCurrency = "GBP"
var exampleMerchantIDVal int32 = 99
var exampleMerchantName = "Test Shop"
var exampleMerchantStatus = "A"
var exampleMerchantLabel = "Active"

func buildExampleMerchant() openapiclient.Merchant {
	m := openapiclient.NewMerchant()
	m.SetCurrency(exampleMerchantCurrency)
	m.SetMerchantid(exampleMerchantIDVal)
	m.SetName(exampleMerchantName)
	m.SetStatus(exampleMerchantStatus)
	m.SetStatusLabel(exampleMerchantLabel)
	return *m
}

func TestNewMerchant(t *testing.T) {
	model := openapiclient.NewMerchant()
	require.NotNil(t, model)
	assert.False(t, model.HasCurrency())
	assert.False(t, model.HasMerchantid())
	assert.False(t, model.HasName())
	assert.False(t, model.HasStatus())
	assert.False(t, model.HasStatusLabel())
}

func TestNewMerchantWithDefaults(t *testing.T) {
	model := openapiclient.NewMerchantWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasCurrency())
	assert.False(t, model.HasMerchantid())
	assert.False(t, model.HasName())
	assert.False(t, model.HasStatus())
	assert.False(t, model.HasStatusLabel())
}

func TestMerchantSetGetCycle(t *testing.T) {
	model := openapiclient.NewMerchant()
	model.SetCurrency(exampleMerchantCurrency)
	assert.True(t, model.HasCurrency())
	assert.Equal(t, exampleMerchantCurrency, model.GetCurrency())
	if val, ok := model.GetCurrencyOk(); assert.True(t, ok) {
		assert.Equal(t, exampleMerchantCurrency, *val)
	}

	model.SetMerchantid(exampleMerchantIDVal)
	assert.True(t, model.HasMerchantid())
	assert.Equal(t, exampleMerchantIDVal, model.GetMerchantid())

	model.SetName(exampleMerchantName)
	assert.True(t, model.HasName())
	assert.Equal(t, exampleMerchantName, model.GetName())

	model.SetStatus(exampleMerchantStatus)
	assert.True(t, model.HasStatus())
	assert.Equal(t, exampleMerchantStatus, model.GetStatus())

	model.SetStatusLabel(exampleMerchantLabel)
	assert.True(t, model.HasStatusLabel())
	assert.Equal(t, exampleMerchantLabel, model.GetStatusLabel())
}

func TestMerchantJSONRoundTrip(t *testing.T) {
	model := buildExampleMerchant()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.Merchant
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasMerchantid())
	assert.Equal(t, exampleMerchantCurrency, unmarshalled.GetCurrency())
}

func TestMerchantToMap(t *testing.T) {
	model := buildExampleMerchant()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "currency") {
		assert.Equal(t, &exampleMerchantCurrency, m["currency"])
	}
	if assert.Contains(t, m, "merchantid") {
		assert.Equal(t, &exampleMerchantIDVal, m["merchantid"])
	}
	if assert.Contains(t, m, "name") {
		assert.Equal(t, &exampleMerchantName, m["name"])
	}
	if assert.Contains(t, m, "status") {
		assert.Equal(t, &exampleMerchantStatus, m["status"])
	}
	if assert.Contains(t, m, "status_label") {
		assert.Equal(t, &exampleMerchantLabel, m["status_label"])
	}
}

func TestNullableMerchantGetSet(t *testing.T) {
	base := buildExampleMerchant()
	n := openapiclient.NullableMerchant{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableMerchantUnset(t *testing.T) {
	base := buildExampleMerchant()
	n := openapiclient.NewNullableMerchant(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableMerchantJSONRoundTrip(t *testing.T) {
	base := buildExampleMerchant()
	n := openapiclient.NewNullableMerchant(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableMerchant
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleMerchantIDVal, newN.Get().GetMerchantid())
	}
}
