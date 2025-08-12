package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleNonce = "ABC"
var exampleRedirectFail = "https://fail"
var exampleRedirectSuccess = "https://ok"
var exampleToken = "tok123"

func buildExampleDirectTokenAuthRequest() *openapiclient.DirectTokenAuthRequest {
	d := openapiclient.NewDirectTokenAuthRequest()
	d.SetNonce(exampleNonce)
	d.SetRedirectFailure(exampleRedirectFail)
	d.SetRedirectSuccess(exampleRedirectSuccess)
	d.SetToken(exampleToken)
	return d
}

func TestNewDirectTokenAuthRequest(t *testing.T) {
	model := openapiclient.NewDirectTokenAuthRequest()
	require.NotNil(t, model)
	assert.False(t, model.HasToken())
}

func TestNewDirectTokenAuthRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewDirectTokenAuthRequestWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasToken())
}

func TestDirectTokenAuthRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewDirectTokenAuthRequest()
	model.SetNonce(exampleNonce)
	assert.True(t, model.HasNonce())
	assert.Equal(t, exampleNonce, model.GetNonce())

	model.SetRedirectFailure(exampleRedirectFail)
	assert.True(t, model.HasRedirectFailure())
	assert.Equal(t, exampleRedirectFail, model.GetRedirectFailure())

	model.SetRedirectSuccess(exampleRedirectSuccess)
	assert.True(t, model.HasRedirectSuccess())
	assert.Equal(t, exampleRedirectSuccess, model.GetRedirectSuccess())

	model.SetToken(exampleToken)
	assert.True(t, model.HasToken())
	assert.Equal(t, exampleToken, model.GetToken())
}

func TestDirectTokenAuthRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleDirectTokenAuthRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.DirectTokenAuthRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasToken())
	assert.Equal(t, exampleToken, unmarshalled.GetToken())
}

func TestDirectTokenAuthRequestToMap(t *testing.T) {
	model := buildExampleDirectTokenAuthRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "token") {
		assert.Equal(t, &exampleToken, m["token"])
	}
}

func TestNullableDirectTokenAuthRequestGetSet(t *testing.T) {
	base := buildExampleDirectTokenAuthRequest()
	n := openapiclient.NullableDirectTokenAuthRequest{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableDirectTokenAuthRequestUnset(t *testing.T) {
	base := buildExampleDirectTokenAuthRequest()
	n := openapiclient.NewNullableDirectTokenAuthRequest(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableDirectTokenAuthRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleDirectTokenAuthRequest()
	n := openapiclient.NewNullableDirectTokenAuthRequest(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableDirectTokenAuthRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleToken, newN.Get().GetToken())
	}
}
