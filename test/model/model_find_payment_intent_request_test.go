package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleFindPaymentIntentRequestExternalRef = "ext123"
var exampleFindPaymentIntentRequestMerchantId int32 = 123456

func buildExampleFindPaymentIntentRequest() *openapiclient.FindPaymentIntentRequest {
	m := openapiclient.NewFindPaymentIntentRequest()
	m.SetExternalRef(exampleFindPaymentIntentRequestExternalRef)
	m.SetMerchantid(exampleFindPaymentIntentRequestMerchantId)
	return m
}

func TestNewFindPaymentIntentRequest(t *testing.T) {
	model := openapiclient.NewFindPaymentIntentRequest()
	require.NotNil(t, model)
	assert.False(t, model.HasExternalRef())
}

func TestNewFindPaymentIntentRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewFindPaymentIntentRequestWithDefaults()
	require.NotNil(t, model)
}

func TestFindPaymentIntentRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewFindPaymentIntentRequest()
	model.SetExternalRef(exampleFindPaymentIntentRequestExternalRef)
	assert.True(t, model.HasExternalRef())
	assert.Equal(t, exampleFindPaymentIntentRequestExternalRef, model.GetExternalRef())
	refPtr, ok := model.GetExternalRefOk()
	require.True(t, ok)
	if assert.NotNil(t, refPtr) {
		assert.Equal(t, exampleFindPaymentIntentRequestExternalRef, *refPtr)
	}
}

func TestFindPaymentIntentRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleFindPaymentIntentRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.FindPaymentIntentRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleFindPaymentIntentRequestExternalRef, unmarshalled.GetExternalRef())
	assert.Equal(t, exampleFindPaymentIntentRequestMerchantId, unmarshalled.GetMerchantid())
}

func TestFindPaymentIntentRequestToMap(t *testing.T) {
	model := buildExampleFindPaymentIntentRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "external-ref") {
		assert.Equal(t, &exampleFindPaymentIntentRequestExternalRef, m["external-ref"])
	}
}

func TestNullableFindPaymentIntentRequestGetSet(t *testing.T) {
	base := buildExampleFindPaymentIntentRequest()
	n := openapiclient.NullableFindPaymentIntentRequest{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableFindPaymentIntentRequestUnset(t *testing.T) {
	base := buildExampleFindPaymentIntentRequest()
	n := openapiclient.NewNullableFindPaymentIntentRequest(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableFindPaymentIntentRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleFindPaymentIntentRequest()
	n := openapiclient.NewNullableFindPaymentIntentRequest(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableFindPaymentIntentRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleFindPaymentIntentRequestExternalRef, newN.Get().GetExternalRef())
	}
}
