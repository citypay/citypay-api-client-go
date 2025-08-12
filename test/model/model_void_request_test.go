package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleVoidIdentifier = "vid"
var exampleVoidMid int32 = 9
var exampleVoidTrans int32 = 22

func buildExampleVoidRequest() openapiclient.VoidRequest {
	r := openapiclient.NewVoidRequest(exampleVoidMid)
	r.SetIdentifier(exampleVoidIdentifier)
	r.SetTransno(exampleVoidTrans)
	return *r
}

func TestNewVoidRequest(t *testing.T) {
	model := openapiclient.NewVoidRequest(exampleVoidMid)
	require.NotNil(t, model)
	assert.Equal(t, exampleVoidMid, model.GetMerchantid())
	assert.False(t, model.HasIdentifier())
}

func TestNewVoidRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewVoidRequestWithDefaults()
	require.NotNil(t, model)
}

func TestVoidRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewVoidRequest(exampleVoidMid)
	model.SetIdentifier(exampleVoidIdentifier)
	assert.True(t, model.HasIdentifier())
	assert.Equal(t, exampleVoidIdentifier, model.GetIdentifier())

	model.SetTransno(exampleVoidTrans)
	assert.True(t, model.HasTransno())
	assert.Equal(t, exampleVoidTrans, model.GetTransno())
}

func TestVoidRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleVoidRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.VoidRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasIdentifier())
	assert.Equal(t, exampleVoidIdentifier, unmarshalled.GetIdentifier())
}

func TestVoidRequestToMap(t *testing.T) {
	model := buildExampleVoidRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	assert.Equal(t, exampleVoidMid, m["merchantid"])
	if assert.Contains(t, m, "identifier") {
		assert.Equal(t, &exampleVoidIdentifier, m["identifier"])
	}
}

func TestNullableVoidRequestGetSet(t *testing.T) {
	base := buildExampleVoidRequest()
	n := openapiclient.NullableVoidRequest{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableVoidRequestUnset(t *testing.T) {
	base := buildExampleVoidRequest()
	n := openapiclient.NewNullableVoidRequest(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableVoidRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleVoidRequest()
	n := openapiclient.NewNullableVoidRequest(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableVoidRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleVoidTrans, newN.Get().GetTransno())
	}
}
