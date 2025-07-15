package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleRefundAmount int32 = 999
var exampleRefundIdentifier = "ref1"
var exampleRefundMid int32 = 5
var exampleRefundRef int32 = 10
var exampleRefundInfo = "details"

func buildExampleRefundRequest() openapiclient.RefundRequest {
	r := openapiclient.NewRefundRequest(exampleRefundAmount, exampleRefundIdentifier, exampleRefundMid, exampleRefundRef)
	r.SetTransInfo(exampleRefundInfo)
	return *r
}

func TestNewRefundRequest(t *testing.T) {
	model := openapiclient.NewRefundRequest(exampleRefundAmount, exampleRefundIdentifier, exampleRefundMid, exampleRefundRef)
	require.NotNil(t, model)
	assert.Equal(t, exampleRefundAmount, model.GetAmount())
	assert.Equal(t, exampleRefundIdentifier, model.GetIdentifier())
	assert.Equal(t, exampleRefundMid, model.GetMerchantid())
	assert.Equal(t, exampleRefundRef, model.GetRefundRef())
	assert.False(t, model.HasTransInfo())
}

func TestNewRefundRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewRefundRequestWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, int32(0), model.GetAmount())
}

func TestRefundRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewRefundRequest(exampleRefundAmount, exampleRefundIdentifier, exampleRefundMid, exampleRefundRef)
	model.SetTransInfo(exampleRefundInfo)
	assert.True(t, model.HasTransInfo())
	assert.Equal(t, exampleRefundInfo, model.GetTransInfo())
}

func TestRefundRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleRefundRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.RefundRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasTransInfo())
	assert.Equal(t, exampleRefundInfo, unmarshalled.GetTransInfo())
}

func TestRefundRequestToMap(t *testing.T) {
	model := buildExampleRefundRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	assert.Equal(t, exampleRefundAmount, m["amount"])
	if assert.Contains(t, m, "trans_info") {
		assert.Equal(t, &exampleRefundInfo, m["trans_info"])
	}
}

func TestNullableRefundRequestGetSet(t *testing.T) {
	base := buildExampleRefundRequest()
	n := openapiclient.NullableRefundRequest{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableRefundRequestUnset(t *testing.T) {
	base := buildExampleRefundRequest()
	n := openapiclient.NewNullableRefundRequest(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableRefundRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleRefundRequest()
	n := openapiclient.NewNullableRefundRequest(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableRefundRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleRefundMid, newN.Get().GetMerchantid())
	}
}
