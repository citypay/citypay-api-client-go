package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleTDSAccept = "text/html"
var exampleTDSColor = "32"
var exampleTDSIP = "127.0.0.1"
var exampleTDSJava = "false"
var exampleTDSLang = "en"
var exampleTDSScreenH = "100"
var exampleTDSScreenW = "200"
var exampleTDSTZ = "0"
var exampleTDSBx = "bx"
var exampleTDSDowngrade = true
var exampleTDSTerm = "https://term"
var exampleTDSPolicy = "1"
var exampleTDSUser = "UA"

func buildExampleThreeDSecure() openapiclient.ThreeDSecure {
	t := openapiclient.NewThreeDSecure()
	t.SetAcceptHeaders(exampleTDSAccept)
	t.SetBrowserColorDepth(exampleTDSColor)
	t.SetBrowserIP(exampleTDSIP)
	t.SetBrowserJavaEnabled(exampleTDSJava)
	t.SetBrowserLanguage(exampleTDSLang)
	t.SetBrowserScreenHeight(exampleTDSScreenH)
	t.SetBrowserScreenWidth(exampleTDSScreenW)
	t.SetBrowserTZ(exampleTDSTZ)
	t.SetCpBx(exampleTDSBx)
	t.SetDowngrade1(exampleTDSDowngrade)
	t.SetMerchantTermurl(exampleTDSTerm)
	t.SetTdsPolicy(exampleTDSPolicy)
	t.SetUserAgent(exampleTDSUser)
	return *t
}

func TestNewThreeDSecure(t *testing.T) {
	model := openapiclient.NewThreeDSecure()
	require.NotNil(t, model)
	assert.False(t, model.HasUserAgent())
}

func TestNewThreeDSecureWithDefaults(t *testing.T) {
	model := openapiclient.NewThreeDSecureWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasDowngrade1())
}

func TestThreeDSecureSetGetCycle(t *testing.T) {
	model := openapiclient.NewThreeDSecure()
	model.SetAcceptHeaders(exampleTDSAccept)
	assert.True(t, model.HasAcceptHeaders())
	assert.Equal(t, exampleTDSAccept, model.GetAcceptHeaders())

	model.SetBrowserColorDepth(exampleTDSColor)
	assert.True(t, model.HasBrowserColorDepth())
	assert.Equal(t, exampleTDSColor, model.GetBrowserColorDepth())

	model.SetBrowserIP(exampleTDSIP)
	assert.True(t, model.HasBrowserIP())
	assert.Equal(t, exampleTDSIP, model.GetBrowserIP())

	model.SetBrowserJavaEnabled(exampleTDSJava)
	assert.True(t, model.HasBrowserJavaEnabled())
	assert.Equal(t, exampleTDSJava, model.GetBrowserJavaEnabled())

	model.SetBrowserLanguage(exampleTDSLang)
	assert.True(t, model.HasBrowserLanguage())
	assert.Equal(t, exampleTDSLang, model.GetBrowserLanguage())

	model.SetBrowserScreenHeight(exampleTDSScreenH)
	assert.True(t, model.HasBrowserScreenHeight())
	assert.Equal(t, exampleTDSScreenH, model.GetBrowserScreenHeight())

	model.SetBrowserScreenWidth(exampleTDSScreenW)
	assert.True(t, model.HasBrowserScreenWidth())
	assert.Equal(t, exampleTDSScreenW, model.GetBrowserScreenWidth())

	model.SetBrowserTZ(exampleTDSTZ)
	assert.True(t, model.HasBrowserTZ())
	assert.Equal(t, exampleTDSTZ, model.GetBrowserTZ())

	model.SetCpBx(exampleTDSBx)
	assert.True(t, model.HasCpBx())
	assert.Equal(t, exampleTDSBx, model.GetCpBx())

	model.SetDowngrade1(exampleTDSDowngrade)
	assert.True(t, model.HasDowngrade1())
	assert.Equal(t, exampleTDSDowngrade, model.GetDowngrade1())

	model.SetMerchantTermurl(exampleTDSTerm)
	assert.True(t, model.HasMerchantTermurl())
	assert.Equal(t, exampleTDSTerm, model.GetMerchantTermurl())

	model.SetTdsPolicy(exampleTDSPolicy)
	assert.True(t, model.HasTdsPolicy())
	assert.Equal(t, exampleTDSPolicy, model.GetTdsPolicy())

	model.SetUserAgent(exampleTDSUser)
	assert.True(t, model.HasUserAgent())
	assert.Equal(t, exampleTDSUser, model.GetUserAgent())
}

func TestThreeDSecureJSONRoundTrip(t *testing.T) {
	model := buildExampleThreeDSecure()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.ThreeDSecure
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasCpBx())
	assert.Equal(t, exampleTDSBx, unmarshalled.GetCpBx())
}

func TestThreeDSecureToMap(t *testing.T) {
	model := buildExampleThreeDSecure()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "browserIP") {
		assert.Equal(t, &exampleTDSIP, m["browserIP"])
	}
}

func TestNullableThreeDSecureGetSet(t *testing.T) {
	base := buildExampleThreeDSecure()
	n := openapiclient.NullableThreeDSecure{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableThreeDSecureUnset(t *testing.T) {
	base := buildExampleThreeDSecure()
	n := openapiclient.NewNullableThreeDSecure(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableThreeDSecureJSONRoundTrip(t *testing.T) {
	base := buildExampleThreeDSecure()
	n := openapiclient.NewNullableThreeDSecure(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableThreeDSecure
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleTDSPolicy, newN.Get().GetTdsPolicy())
	}
}
