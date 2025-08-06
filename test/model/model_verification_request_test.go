package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleVerificationRequestAmount int32 = 1000
var exampleVerificationRequestIdentifier = "verify-123"
var exampleVerificationRequestMerchantId int32 = 987654
var exampleVerificationRequestCurrency = "GBP"

func buildExampleVerificationRequest() *openapiclient.VerificationRequest {
	m := openapiclient.NewVerificationRequest(exampleVerificationRequestAmount, exampleVerificationRequestIdentifier, exampleVerificationRequestMerchantId)
	m.SetCurrency(exampleVerificationRequestCurrency)
	m.SetBillTo(buildExampleContact())
	m.SetCardnumber("4111111111111111")
	m.SetExpmonth(12)
	m.SetExpyear(2030)
	return m
}

func TestNewVerificationRequest(t *testing.T) {
	model := openapiclient.NewVerificationRequest(exampleVerificationRequestAmount, exampleVerificationRequestIdentifier, exampleVerificationRequestMerchantId)
	require.NotNil(t, model)
	assert.Equal(t, exampleVerificationRequestAmount, model.GetAmount())
	assert.False(t, model.HasCurrency())
}

func TestNewVerificationRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewVerificationRequestWithDefaults()
	require.NotNil(t, model)
}

func TestVerificationRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewVerificationRequest(exampleVerificationRequestAmount, exampleVerificationRequestIdentifier, exampleVerificationRequestMerchantId)
	model.SetCurrency(exampleVerificationRequestCurrency)
	assert.True(t, model.HasCurrency())
	assert.Equal(t, exampleVerificationRequestCurrency, model.GetCurrency())
	curPtr, ok := model.GetCurrencyOk()
	require.True(t, ok)
	if assert.NotNil(t, curPtr) {
		assert.Equal(t, exampleVerificationRequestCurrency, *curPtr)
	}
}

func TestVerificationRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleVerificationRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.VerificationRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleVerificationRequestAmount, unmarshalled.GetAmount())
	assert.Equal(t, exampleVerificationRequestMerchantId, unmarshalled.GetMerchantid())
}

func TestVerificationRequestToMap(t *testing.T) {
	model := buildExampleVerificationRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "amount") {
		assert.Equal(t, int32(exampleVerificationRequestAmount), m["amount"])
	}
}

func TestNullableVerificationRequestGetSet(t *testing.T) {
	base := buildExampleVerificationRequest()
	n := openapiclient.NullableVerificationRequest{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableVerificationRequestUnset(t *testing.T) {
	base := buildExampleVerificationRequest()
	n := openapiclient.NewNullableVerificationRequest(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableVerificationRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleVerificationRequest()
	n := openapiclient.NewNullableVerificationRequest(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableVerificationRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleVerificationRequestIdentifier, newN.Get().GetIdentifier())
	}
}
