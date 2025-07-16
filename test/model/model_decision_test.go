package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleAcsURL = "https://acs.example.com"
var exampleCreq = "creqdata"
var exampleDecisionMerchantID int32 = 100
var exampleThreeDSID = "threedid"
var exampleTransno int32 = 5
var exampleAuthMerchant int32 = 321
var exampleAuthResult int32 = 1
var exampleAuthCode = "A"
var exampleAuthMsg = "Approved"

func buildExampleRequestChallenged() openapiclient.RequestChallenged {
	r := openapiclient.NewRequestChallenged()
	r.SetAcsUrl(exampleAcsURL)
	r.SetCreq(exampleCreq)
	r.SetMerchantid(exampleDecisionMerchantID)
	r.SetThreedserverTransId(exampleThreeDSID)
	r.SetTransno(exampleTransno)
	return *r
}

func buildExampleAuthResp() openapiclient.AuthResponse {
	a := openapiclient.NewAuthResponse(exampleAuthMerchant, exampleAuthResult, exampleAuthCode, exampleAuthMsg)
	a.SetAuthcode("CODE")
	return *a
}

func buildExampleDecision() *openapiclient.Decision {
	d := openapiclient.NewDecision()
	d.SetAuthResponse(buildExampleAuthResp())
	d.SetRequestChallenged(buildExampleRequestChallenged())
	return d
}

func TestNewDecision(t *testing.T) {
	model := openapiclient.NewDecision()
	require.NotNil(t, model)
	assert.False(t, model.HasAuthResponse())
	assert.False(t, model.HasRequestChallenged())
}

func TestNewDecisionWithDefaults(t *testing.T) {
	model := openapiclient.NewDecisionWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasAuthResponse())
}

func TestDecisionSetGetCycle(t *testing.T) {
	model := openapiclient.NewDecision()
	ar := buildExampleAuthResp()
	model.SetAuthResponse(ar)
	assert.True(t, model.HasAuthResponse())
	assert.Equal(t, ar, model.GetAuthResponse())
	if val, ok := model.GetAuthResponseOk(); assert.True(t, ok) {
		assert.Equal(t, ar, *val)
	}

	rc := buildExampleRequestChallenged()
	model.SetRequestChallenged(rc)
	assert.True(t, model.HasRequestChallenged())
	assert.Equal(t, rc, model.GetRequestChallenged())
	if val, ok := model.GetRequestChallengedOk(); assert.True(t, ok) {
		assert.Equal(t, rc, *val)
	}
}

func TestDecisionJSONRoundTrip(t *testing.T) {
	model := buildExampleDecision()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.Decision
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasAuthResponse())
	assert.True(t, unmarshalled.HasRequestChallenged())
}

func TestDecisionToMap(t *testing.T) {
	model := buildExampleDecision()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "AuthResponse") {
		assert.NotNil(t, m["AuthResponse"])
	}
	if assert.Contains(t, m, "RequestChallenged") {
		assert.NotNil(t, m["RequestChallenged"])
	}
}

func TestNullableDecisionGetSet(t *testing.T) {
	base := buildExampleDecision()
	n := openapiclient.NullableDecision{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableDecisionUnset(t *testing.T) {
	base := buildExampleDecision()
	n := openapiclient.NewNullableDecision(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableDecisionJSONRoundTrip(t *testing.T) {
	base := buildExampleDecision()
	n := openapiclient.NewNullableDecision(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableDecision
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.True(t, newN.Get().HasRequestChallenged())
	}
}
