package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleRetrieveIdentifier = "rid"
var exampleRetrieveMid int32 = 2
var exampleRetrieveTrans int32 = 10

func buildExampleRetrieveRequest() openapiclient.RetrieveRequest {
	r := openapiclient.NewRetrieveRequest(exampleRetrieveMid)
	r.SetIdentifier(exampleRetrieveIdentifier)
	r.SetTransno(exampleRetrieveTrans)
	return *r
}

func TestNewRetrieveRequest(t *testing.T) {
	model := openapiclient.NewRetrieveRequest(exampleRetrieveMid)
	require.NotNil(t, model)
	assert.Equal(t, exampleRetrieveMid, model.GetMerchantid())
	assert.False(t, model.HasIdentifier())
}

func TestNewRetrieveRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewRetrieveRequestWithDefaults()
	require.NotNil(t, model)
}

func TestRetrieveRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewRetrieveRequest(exampleRetrieveMid)
	model.SetIdentifier(exampleRetrieveIdentifier)
	assert.True(t, model.HasIdentifier())
	assert.Equal(t, exampleRetrieveIdentifier, model.GetIdentifier())

	model.SetTransno(exampleRetrieveTrans)
	assert.True(t, model.HasTransno())
	assert.Equal(t, exampleRetrieveTrans, model.GetTransno())
}

func TestRetrieveRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleRetrieveRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.RetrieveRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasTransno())
	assert.Equal(t, exampleRetrieveTrans, unmarshalled.GetTransno())
}

func TestRetrieveRequestToMap(t *testing.T) {
	model := buildExampleRetrieveRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	assert.Equal(t, exampleRetrieveMid, m["merchantid"])
	if assert.Contains(t, m, "identifier") {
		assert.Equal(t, &exampleRetrieveIdentifier, m["identifier"])
	}
}

func TestNullableRetrieveRequestGetSet(t *testing.T) {
	base := buildExampleRetrieveRequest()
	n := openapiclient.NullableRetrieveRequest{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableRetrieveRequestUnset(t *testing.T) {
	base := buildExampleRetrieveRequest()
	n := openapiclient.NewNullableRetrieveRequest(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableRetrieveRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleRetrieveRequest()
	n := openapiclient.NewNullableRetrieveRequest(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableRetrieveRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleRetrieveIdentifier, newN.Get().GetIdentifier())
	}
}
