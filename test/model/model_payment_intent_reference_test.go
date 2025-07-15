package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleIntentRefId = "intent123"

func buildExamplePaymentIntentReference() openapiclient.PaymentIntentReference {
	r := openapiclient.NewPaymentIntentReference(exampleIntentRefId)
	return *r
}

func TestNewPaymentIntentReference(t *testing.T) {
	model := openapiclient.NewPaymentIntentReference(exampleIntentRefId)
	require.NotNil(t, model)
	assert.Equal(t, exampleIntentRefId, model.GetPaymentIntentId())
}

func TestNewPaymentIntentReferenceWithDefaults(t *testing.T) {
	model := openapiclient.NewPaymentIntentReferenceWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, "", model.GetPaymentIntentId())
}

func TestPaymentIntentReferenceSetGetCycle(t *testing.T) {
	model := openapiclient.NewPaymentIntentReference(exampleIntentRefId)
	model.SetPaymentIntentId(exampleIntentRefId + "2")
	assert.Equal(t, exampleIntentRefId+"2", model.GetPaymentIntentId())
}

func TestPaymentIntentReferenceJSONRoundTrip(t *testing.T) {
	model := buildExamplePaymentIntentReference()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.PaymentIntentReference
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleIntentRefId, unmarshalled.GetPaymentIntentId())
}

func TestPaymentIntentReferenceToMap(t *testing.T) {
	model := buildExamplePaymentIntentReference()
	m, err := model.ToMap()
	require.NoError(t, err)
	assert.Equal(t, exampleIntentRefId, m["payment_intent_id"])
}

func TestNullablePaymentIntentReferenceGetSet(t *testing.T) {
	base := buildExamplePaymentIntentReference()
	n := openapiclient.NullablePaymentIntentReference{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullablePaymentIntentReferenceUnset(t *testing.T) {
	base := buildExamplePaymentIntentReference()
	n := openapiclient.NewNullablePaymentIntentReference(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullablePaymentIntentReferenceJSONRoundTrip(t *testing.T) {
	base := buildExamplePaymentIntentReference()
	n := openapiclient.NewNullablePaymentIntentReference(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullablePaymentIntentReference
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleIntentRefId, newN.Get().GetPaymentIntentId())
	}
}
