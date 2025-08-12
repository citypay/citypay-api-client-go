package citypay

import (
	"encoding/json"
	"testing"

	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var exampleCaptureAmount int32 = 5000
var exampleCaptureIdentifier = "CAPID"
var exampleCaptureMerchantID int32 = 321
var exampleCaptureTransno int32 = 42

func buildExampleCaptureRequest() *openapiclient.CaptureRequest {
	c := openapiclient.NewCaptureRequest(exampleCaptureMerchantID)
	c.SetAirlineData(*buildExampleAdvice())
	c.SetAmount(exampleCaptureAmount)
	c.SetEventManagement(buildExampleEventData())
	c.SetIdentifier(exampleCaptureIdentifier)
	c.SetTransno(exampleCaptureTransno)
	return c
}

func TestNewCaptureRequest(t *testing.T) {
	model := openapiclient.NewCaptureRequest(exampleCaptureMerchantID)
	require.NotNil(t, model)
	assert.Equal(t, exampleCaptureMerchantID, model.GetMerchantid())
	assert.False(t, model.HasAmount())
}

func TestNewCaptureRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewCaptureRequestWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, int32(0), model.GetMerchantid())
}

func TestCaptureRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewCaptureRequest(exampleCaptureMerchantID)
	model.SetAirlineData(*buildExampleAdvice())
	assert.True(t, model.HasAirlineData())

	model.SetAmount(exampleCaptureAmount)
	assert.True(t, model.HasAmount())
	assert.Equal(t, exampleCaptureAmount, model.GetAmount())

	model.SetEventManagement(buildExampleEventData())
	assert.True(t, model.HasEventManagement())

	model.SetIdentifier(exampleCaptureIdentifier)
	assert.True(t, model.HasIdentifier())
	assert.Equal(t, exampleCaptureIdentifier, model.GetIdentifier())

	model.SetMerchantid(exampleCaptureMerchantID)
	assert.Equal(t, exampleCaptureMerchantID, model.GetMerchantid())

	model.SetTransno(exampleCaptureTransno)
	assert.True(t, model.HasTransno())
	assert.Equal(t, exampleCaptureTransno, model.GetTransno())
}

func TestCaptureRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleCaptureRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.CaptureRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleCaptureMerchantID, unmarshalled.GetMerchantid())
	assert.True(t, unmarshalled.HasIdentifier())
}

func TestCaptureRequestToMap(t *testing.T) {
	model := buildExampleCaptureRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "merchantid") {
		assert.Equal(t, exampleCaptureMerchantID, m["merchantid"])
	}
	if assert.Contains(t, m, "identifier") {
		assert.Equal(t, &exampleCaptureIdentifier, m["identifier"])
	}
}

func TestNullableCaptureRequestGetSet(t *testing.T) {
	base := buildExampleCaptureRequest()
	n := openapiclient.NullableCaptureRequest{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableCaptureRequestUnset(t *testing.T) {
	base := buildExampleCaptureRequest()
	n := openapiclient.NewNullableCaptureRequest(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableCaptureRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleCaptureRequest()
	n := openapiclient.NewNullableCaptureRequest(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableCaptureRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleCaptureIdentifier, newN.Get().GetIdentifier())
	}
}
