package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleConfigAcs = "iframe"
var exampleConfigDescriptor = "desc"
var exampleConfigExpire = "30m"
var exampleConfigLock = []string{"x"}
var exampleConfigLogo = "https://img"
var exampleConfigTerms = "https://terms"
var exampleConfigOptions = []string{"o1"}
var exampleConfigPostback = "https://post"
var exampleConfigPassword = "pass"
var exampleConfigPolicy = "ALWAYS"
var exampleConfigUser = "user"
var exampleConfigDelay int32 = 5
var exampleConfigFail = "https://fail"
var exampleConfigSuccess = "https://ok"
var exampleConfigRenderer = "default"
var exampleConfigReturn = true

func buildExampleCustomParam() openapiclient.PaylinkCustomParam {
	c := openapiclient.NewPaylinkCustomParam("p")
	c.SetLabel("label")
	return *c
}

func buildExampleFieldGuard() openapiclient.PaylinkFieldGuardModel {
	f := openapiclient.NewPaylinkFieldGuardModel()
	f.SetName("n")
	return *f
}

func buildExamplePartPayments() openapiclient.PaylinkPartPayments {
	p := openapiclient.NewPaylinkPartPayments()
	p.SetEnabled("true")
	return *p
}

func buildExampleUI() openapiclient.PaylinkUI {
	u := openapiclient.NewPaylinkUI()
	u.SetOrdering(1)
	return *u
}

func buildExamplePaylinkConfig() openapiclient.PaylinkConfig {
	c := openapiclient.NewPaylinkConfig()
	c.SetAcsMode(exampleConfigAcs)
	c.SetCustomParams([]openapiclient.PaylinkCustomParam{buildExampleCustomParam()})
	c.SetDescriptor(exampleConfigDescriptor)
	c.SetExpireIn(exampleConfigExpire)
	c.SetFieldGuard([]openapiclient.PaylinkFieldGuardModel{buildExampleFieldGuard()})
	c.SetLockParams(exampleConfigLock)
	c.SetMerchLogo(exampleConfigLogo)
	c.SetMerchTerms(exampleConfigTerms)
	c.SetOptions(exampleConfigOptions)
	c.SetPartPayments(buildExamplePartPayments())
	m := map[string]string{"k": "v"}
	c.SetPassThroughData(m)
	c.SetPassThroughHeaders(m)
	c.SetPostback(exampleConfigPostback)
	c.SetPostbackPassword(exampleConfigPassword)
	c.SetPostbackPolicy(exampleConfigPolicy)
	c.SetPostbackUsername(exampleConfigUser)
	c.SetRedirectDelay(exampleConfigDelay)
	c.SetRedirectFailure(exampleConfigFail)
	c.SetRedirectSuccess(exampleConfigSuccess)
	c.SetRenderer(exampleConfigRenderer)
	c.SetReturnParams(exampleConfigReturn)
	c.SetUi(buildExampleUI())
	return *c
}

func TestNewPaylinkConfig(t *testing.T) {
	model := openapiclient.NewPaylinkConfig()
	require.NotNil(t, model)
	assert.False(t, model.HasAcsMode())
}

func TestNewPaylinkConfigWithDefaults(t *testing.T) {
	model := openapiclient.NewPaylinkConfigWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasOptions())
}

func TestPaylinkConfigSetGetCycle(t *testing.T) {
	model := openapiclient.NewPaylinkConfig()
	model.SetAcsMode(exampleConfigAcs)
	assert.True(t, model.HasAcsMode())
	assert.Equal(t, exampleConfigAcs, model.GetAcsMode())

	model.SetCustomParams([]openapiclient.PaylinkCustomParam{buildExampleCustomParam()})
	assert.True(t, model.HasCustomParams())
	assert.Len(t, model.GetCustomParams(), 1)

	model.SetDescriptor(exampleConfigDescriptor)
	assert.True(t, model.HasDescriptor())
	assert.Equal(t, exampleConfigDescriptor, model.GetDescriptor())

	model.SetExpireIn(exampleConfigExpire)
	assert.True(t, model.HasExpireIn())
	assert.Equal(t, exampleConfigExpire, model.GetExpireIn())

	model.SetFieldGuard([]openapiclient.PaylinkFieldGuardModel{buildExampleFieldGuard()})
	assert.True(t, model.HasFieldGuard())
	assert.Len(t, model.GetFieldGuard(), 1)

	model.SetLockParams(exampleConfigLock)
	assert.True(t, model.HasLockParams())
	assert.Equal(t, exampleConfigLock, model.GetLockParams())

	model.SetMerchLogo(exampleConfigLogo)
	assert.True(t, model.HasMerchLogo())
	assert.Equal(t, exampleConfigLogo, model.GetMerchLogo())

	model.SetMerchTerms(exampleConfigTerms)
	assert.True(t, model.HasMerchTerms())
	assert.Equal(t, exampleConfigTerms, model.GetMerchTerms())

	model.SetOptions(exampleConfigOptions)
	assert.True(t, model.HasOptions())
	assert.Equal(t, exampleConfigOptions, model.GetOptions())

	part := buildExamplePartPayments()
	model.SetPartPayments(part)
	assert.True(t, model.HasPartPayments())
	assert.Equal(t, part, model.GetPartPayments())

	m := map[string]string{"k": "v"}
	model.SetPassThroughData(m)
	assert.True(t, model.HasPassThroughData())

	model.SetPassThroughHeaders(m)
	assert.True(t, model.HasPassThroughHeaders())

	model.SetPostback(exampleConfigPostback)
	assert.True(t, model.HasPostback())
	assert.Equal(t, exampleConfigPostback, model.GetPostback())

	model.SetPostbackPassword(exampleConfigPassword)
	assert.True(t, model.HasPostbackPassword())
	assert.Equal(t, exampleConfigPassword, model.GetPostbackPassword())

	model.SetPostbackPolicy(exampleConfigPolicy)
	assert.True(t, model.HasPostbackPolicy())
	assert.Equal(t, exampleConfigPolicy, model.GetPostbackPolicy())

	model.SetPostbackUsername(exampleConfigUser)
	assert.True(t, model.HasPostbackUsername())
	assert.Equal(t, exampleConfigUser, model.GetPostbackUsername())

	model.SetRedirectDelay(exampleConfigDelay)
	assert.True(t, model.HasRedirectDelay())
	assert.Equal(t, exampleConfigDelay, model.GetRedirectDelay())

	model.SetRedirectFailure(exampleConfigFail)
	assert.True(t, model.HasRedirectFailure())
	assert.Equal(t, exampleConfigFail, model.GetRedirectFailure())

	model.SetRedirectSuccess(exampleConfigSuccess)
	assert.True(t, model.HasRedirectSuccess())
	assert.Equal(t, exampleConfigSuccess, model.GetRedirectSuccess())

	model.SetRenderer(exampleConfigRenderer)
	assert.True(t, model.HasRenderer())
	assert.Equal(t, exampleConfigRenderer, model.GetRenderer())

	model.SetReturnParams(exampleConfigReturn)
	assert.True(t, model.HasReturnParams())
	assert.Equal(t, exampleConfigReturn, model.GetReturnParams())

	ui := buildExampleUI()
	model.SetUi(ui)
	assert.True(t, model.HasUi())
	assert.Equal(t, ui, model.GetUi())
}

func TestPaylinkConfigJSONRoundTrip(t *testing.T) {
	model := buildExamplePaylinkConfig()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.PaylinkConfig
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasAcsMode())
	assert.Equal(t, exampleConfigAcs, unmarshalled.GetAcsMode())
}

func TestPaylinkConfigToMap(t *testing.T) {
	model := buildExamplePaylinkConfig()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "descriptor") {
		assert.Equal(t, &exampleConfigDescriptor, m["descriptor"])
	}
	if assert.Contains(t, m, "renderer") {
		assert.Equal(t, &exampleConfigRenderer, m["renderer"])
	}
}

func TestNullablePaylinkConfigGetSet(t *testing.T) {
	base := buildExamplePaylinkConfig()
	n := openapiclient.NullablePaylinkConfig{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkConfigUnset(t *testing.T) {
	base := buildExamplePaylinkConfig()
	n := openapiclient.NewNullablePaylinkConfig(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullablePaylinkConfigJSONRoundTrip(t *testing.T) {
	base := buildExamplePaylinkConfig()
	n := openapiclient.NewNullablePaylinkConfig(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullablePaylinkConfig
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleConfigFail, newN.Get().GetRedirectFailure())
	}
}
