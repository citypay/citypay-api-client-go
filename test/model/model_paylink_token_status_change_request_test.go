package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var exampleChangeAfter = time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
var exampleChangeMax int32 = 10
var exampleChangeMerchant int32 = 1
var exampleChangeNext = "nxt"
var exampleChangeOrder = "p.id"

func buildExamplePaylinkTokenStatusChangeRequest() openapiclient.PaylinkTokenStatusChangeRequest {
	r := openapiclient.NewPaylinkTokenStatusChangeRequest(exampleChangeAfter, exampleChangeMerchant)
	r.SetMaxResults(exampleChangeMax)
	r.SetNextToken(exampleChangeNext)
	r.SetOrderBy(exampleChangeOrder)
	return *r
}

func TestNewPaylinkTokenStatusChangeRequest(t *testing.T) {
	model := openapiclient.NewPaylinkTokenStatusChangeRequest(exampleChangeAfter, exampleChangeMerchant)
	require.NotNil(t, model)
	assert.Equal(t, exampleChangeAfter, model.GetAfter())
	assert.Equal(t, exampleChangeMerchant, model.GetMerchantid())
	assert.False(t, model.HasMaxResults())
}

func TestNewPaylinkTokenStatusChangeRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewPaylinkTokenStatusChangeRequestWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasNextToken())
}

func TestPaylinkTokenStatusChangeRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewPaylinkTokenStatusChangeRequest(exampleChangeAfter, exampleChangeMerchant)
	model.SetMaxResults(exampleChangeMax)
	assert.True(t, model.HasMaxResults())
	assert.Equal(t, exampleChangeMax, model.GetMaxResults())

	model.SetNextToken(exampleChangeNext)
	assert.True(t, model.HasNextToken())
	assert.Equal(t, exampleChangeNext, model.GetNextToken())

	model.SetOrderBy(exampleChangeOrder)
	assert.True(t, model.HasOrderBy())
	assert.Equal(t, exampleChangeOrder, model.GetOrderBy())
}

func TestPaylinkTokenStatusChangeRequestJSONRoundTrip(t *testing.T) {
	model := buildExamplePaylinkTokenStatusChangeRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.PaylinkTokenStatusChangeRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasMaxResults())
	assert.Equal(t, exampleChangeMax, unmarshalled.GetMaxResults())
}

func TestPaylinkTokenStatusChangeRequestToMap(t *testing.T) {
	model := buildExamplePaylinkTokenStatusChangeRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	assert.Equal(t, exampleChangeAfter, m["after"])
	if assert.Contains(t, m, "nextToken") {
		assert.Equal(t, &exampleChangeNext, m["nextToken"])
	}
}

func TestNullablePaylinkTokenStatusChangeRequestGetSet(t *testing.T) {
	base := buildExamplePaylinkTokenStatusChangeRequest()
	n := openapiclient.NullablePaylinkTokenStatusChangeRequest{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkTokenStatusChangeRequestUnset(t *testing.T) {
	base := buildExamplePaylinkTokenStatusChangeRequest()
	n := openapiclient.NewNullablePaylinkTokenStatusChangeRequest(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullablePaylinkTokenStatusChangeRequestJSONRoundTrip(t *testing.T) {
	base := buildExamplePaylinkTokenStatusChangeRequest()
	n := openapiclient.NewNullablePaylinkTokenStatusChangeRequest(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullablePaylinkTokenStatusChangeRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleChangeMerchant, newN.Get().GetMerchantid())
	}
}
