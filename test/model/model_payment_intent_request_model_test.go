package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var examplePaymentIntentRequestAmount int32 = 5000
var examplePaymentIntentRequestIdentifier = "pi-123"
var examplePaymentIntentRequestCurrency = "GBP"

func buildExamplePaymentIntentRequestModel() *openapiclient.PaymentIntentRequestModel {
	m := openapiclient.NewPaymentIntentRequestModel(examplePaymentIntentRequestAmount, examplePaymentIntentRequestIdentifier)
	m.SetCurrency(examplePaymentIntentRequestCurrency)
	m.SetAdjustments(*buildExampleAdjustments())
	m.SetBillTo(buildExampleContact())
	return m
}

func TestNewPaymentIntentRequestModel(t *testing.T) {
	model := openapiclient.NewPaymentIntentRequestModel(examplePaymentIntentRequestAmount, examplePaymentIntentRequestIdentifier)
	require.NotNil(t, model)
	assert.Equal(t, examplePaymentIntentRequestAmount, model.GetAmount())
	assert.Equal(t, examplePaymentIntentRequestIdentifier, model.GetIdentifier())
	assert.False(t, model.HasCurrency())
}

func TestNewPaymentIntentRequestModelWithDefaults(t *testing.T) {
	model := openapiclient.NewPaymentIntentRequestModelWithDefaults()
	require.NotNil(t, model)
}

func TestPaymentIntentRequestModelSetGetCycle(t *testing.T) {
	model := openapiclient.NewPaymentIntentRequestModel(examplePaymentIntentRequestAmount, examplePaymentIntentRequestIdentifier)
	model.SetCurrency(examplePaymentIntentRequestCurrency)
	assert.True(t, model.HasCurrency())
	assert.Equal(t, examplePaymentIntentRequestCurrency, model.GetCurrency())
	curPtr, ok := model.GetCurrencyOk()
	require.True(t, ok)
	if assert.NotNil(t, curPtr) {
		assert.Equal(t, examplePaymentIntentRequestCurrency, *curPtr)
	}
}

func TestPaymentIntentRequestModelJSONRoundTrip(t *testing.T) {
	model := buildExamplePaymentIntentRequestModel()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.PaymentIntentRequestModel
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, examplePaymentIntentRequestAmount, unmarshalled.GetAmount())
	assert.Equal(t, examplePaymentIntentRequestCurrency, unmarshalled.GetCurrency())
}

func TestPaymentIntentRequestModelToMap(t *testing.T) {
	model := buildExamplePaymentIntentRequestModel()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "amount") {
		assert.Equal(t, int32(examplePaymentIntentRequestAmount), m["amount"])
	}
}

func TestNullablePaymentIntentRequestModelGetSet(t *testing.T) {
	base := buildExamplePaymentIntentRequestModel()
	n := openapiclient.NullablePaymentIntentRequestModel{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullablePaymentIntentRequestModelUnset(t *testing.T) {
	base := buildExamplePaymentIntentRequestModel()
	n := openapiclient.NewNullablePaymentIntentRequestModel(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullablePaymentIntentRequestModelJSONRoundTrip(t *testing.T) {
	base := buildExamplePaymentIntentRequestModel()
	n := openapiclient.NewNullablePaymentIntentRequestModel(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullablePaymentIntentRequestModel
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, examplePaymentIntentRequestIdentifier, newN.Get().GetIdentifier())
	}
}
