package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var examplePaymentIntentResponseIdentifier = "pi-123"
var examplePaymentIntentResponseMerchantId int32 = 111111
var examplePaymentIntentResponseId = "intent123"
var examplePaymentIntentResponseAmount int32 = 5000
var examplePaymentIntentResponseCreated = time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)

func buildExamplePaymentIntentResponseModel() *openapiclient.PaymentIntentResponseModel {
	m := openapiclient.NewPaymentIntentResponseModel(examplePaymentIntentResponseIdentifier, examplePaymentIntentResponseMerchantId, examplePaymentIntentResponseId)
	m.SetAmount(examplePaymentIntentResponseAmount)
	m.SetCreated(examplePaymentIntentResponseCreated)
	m.SetAdjustments(*buildExampleAdjustments())
	return m
}

func TestNewPaymentIntentResponseModel(t *testing.T) {
	model := openapiclient.NewPaymentIntentResponseModel(examplePaymentIntentResponseIdentifier, examplePaymentIntentResponseMerchantId, examplePaymentIntentResponseId)
	require.NotNil(t, model)
	assert.Equal(t, examplePaymentIntentResponseIdentifier, model.GetIdentifier())
	assert.Equal(t, examplePaymentIntentResponseMerchantId, model.GetMerchantid())
	assert.False(t, model.HasAmount())
}

func TestNewPaymentIntentResponseModelWithDefaults(t *testing.T) {
	model := openapiclient.NewPaymentIntentResponseModelWithDefaults()
	require.NotNil(t, model)
}

func TestPaymentIntentResponseModelSetGetCycle(t *testing.T) {
	model := openapiclient.NewPaymentIntentResponseModel(examplePaymentIntentResponseIdentifier, examplePaymentIntentResponseMerchantId, examplePaymentIntentResponseId)
	model.SetAmount(examplePaymentIntentResponseAmount)
	assert.True(t, model.HasAmount())
	assert.Equal(t, examplePaymentIntentResponseAmount, model.GetAmount())
	amtPtr, ok := model.GetAmountOk()
	require.True(t, ok)
	if assert.NotNil(t, amtPtr) {
		assert.Equal(t, examplePaymentIntentResponseAmount, *amtPtr)
	}
}

func TestPaymentIntentResponseModelJSONRoundTrip(t *testing.T) {
	model := buildExamplePaymentIntentResponseModel()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.PaymentIntentResponseModel
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, examplePaymentIntentResponseIdentifier, unmarshalled.GetIdentifier())
	assert.Equal(t, examplePaymentIntentResponseId, unmarshalled.GetPaymentIntentId())
}

func TestPaymentIntentResponseModelToMap(t *testing.T) {
	model := buildExamplePaymentIntentResponseModel()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "identifier") {
		assert.Equal(t, examplePaymentIntentResponseIdentifier, m["identifier"])
	}
}

func TestNullablePaymentIntentResponseModelGetSet(t *testing.T) {
	base := buildExamplePaymentIntentResponseModel()
	n := openapiclient.NullablePaymentIntentResponseModel{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullablePaymentIntentResponseModelUnset(t *testing.T) {
	base := buildExamplePaymentIntentResponseModel()
	n := openapiclient.NewNullablePaymentIntentResponseModel(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullablePaymentIntentResponseModelJSONRoundTrip(t *testing.T) {
	base := buildExamplePaymentIntentResponseModel()
	n := openapiclient.NewNullablePaymentIntentResponseModel(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullablePaymentIntentResponseModel
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, examplePaymentIntentResponseIdentifier, newN.Get().GetIdentifier())
	}
}
