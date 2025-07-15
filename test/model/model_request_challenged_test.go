package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleChallengedURL = "https://acs"
var exampleChallengedCreq = "creq"
var exampleChallengedMid int32 = 1
var exampleChallengedTransID = "tid"
var exampleChallengedTransNo int32 = 5

func buildExampleRequestChallenged() openapiclient.RequestChallenged {
	r := openapiclient.NewRequestChallenged()
	r.SetAcsUrl(exampleChallengedURL)
	r.SetCreq(exampleChallengedCreq)
	r.SetMerchantid(exampleChallengedMid)
	r.SetThreedserverTransId(exampleChallengedTransID)
	r.SetTransno(exampleChallengedTransNo)
	return *r
}

func TestNewRequestChallenged(t *testing.T) {
	model := openapiclient.NewRequestChallenged()
	require.NotNil(t, model)
	assert.False(t, model.HasAcsUrl())
}

func TestNewRequestChallengedWithDefaults(t *testing.T) {
	model := openapiclient.NewRequestChallengedWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasTransno())
}

func TestRequestChallengedSetGetCycle(t *testing.T) {
	model := openapiclient.NewRequestChallenged()
	model.SetAcsUrl(exampleChallengedURL)
	assert.True(t, model.HasAcsUrl())
	assert.Equal(t, exampleChallengedURL, model.GetAcsUrl())

	model.SetCreq(exampleChallengedCreq)
	assert.True(t, model.HasCreq())
	assert.Equal(t, exampleChallengedCreq, model.GetCreq())

	model.SetMerchantid(exampleChallengedMid)
	assert.True(t, model.HasMerchantid())
	assert.Equal(t, exampleChallengedMid, model.GetMerchantid())

	model.SetThreedserverTransId(exampleChallengedTransID)
	assert.True(t, model.HasThreedserverTransId())
	assert.Equal(t, exampleChallengedTransID, model.GetThreedserverTransId())

	model.SetTransno(exampleChallengedTransNo)
	assert.True(t, model.HasTransno())
	assert.Equal(t, exampleChallengedTransNo, model.GetTransno())
}

func TestRequestChallengedJSONRoundTrip(t *testing.T) {
	model := buildExampleRequestChallenged()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.RequestChallenged
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasMerchantid())
	assert.Equal(t, exampleChallengedMid, unmarshalled.GetMerchantid())
}

func TestRequestChallengedToMap(t *testing.T) {
	model := buildExampleRequestChallenged()
	m, err := model.ToMap()
	require.NoError(t, err)
	assert.Equal(t, &exampleChallengedURL, m["acs_url"])
}

func TestNullableRequestChallengedGetSet(t *testing.T) {
	base := buildExampleRequestChallenged()
	n := openapiclient.NullableRequestChallenged{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableRequestChallengedUnset(t *testing.T) {
	base := buildExampleRequestChallenged()
	n := openapiclient.NewNullableRequestChallenged(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableRequestChallengedJSONRoundTrip(t *testing.T) {
	base := buildExampleRequestChallenged()
	n := openapiclient.NewNullableRequestChallenged(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableRequestChallenged
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleChallengedTransID, newN.Get().GetThreedserverTransId())
	}
}
