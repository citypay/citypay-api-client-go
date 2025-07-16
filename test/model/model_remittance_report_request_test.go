package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleRepFrom = "2024-01-01"
var exampleRepUntil = "2024-01-31"
var exampleRepMax int32 = 20
var exampleRepMerchant = []int32{1, 2}
var exampleRepNext = "tok"
var exampleRepOrder = "-date"

func buildExampleRemittanceReportRequest() openapiclient.RemittanceReportRequest {
	r := openapiclient.NewRemittanceReportRequest()
	r.SetDateFrom(exampleRepFrom)
	r.SetDateUntil(exampleRepUntil)
	r.SetMaxResults(exampleRepMax)
	r.SetMerchantId(exampleRepMerchant)
	r.SetNextToken(exampleRepNext)
	r.SetOrderBy(exampleRepOrder)
	return *r
}

func TestNewRemittanceReportRequest(t *testing.T) {
	model := openapiclient.NewRemittanceReportRequest()
	require.NotNil(t, model)
	assert.False(t, model.HasDateFrom())
}

func TestNewRemittanceReportRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewRemittanceReportRequestWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasMerchantId())
}

func TestRemittanceReportRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewRemittanceReportRequest()
	model.SetDateFrom(exampleRepFrom)
	assert.True(t, model.HasDateFrom())
	assert.Equal(t, exampleRepFrom, model.GetDateFrom())

	model.SetDateUntil(exampleRepUntil)
	assert.True(t, model.HasDateUntil())
	assert.Equal(t, exampleRepUntil, model.GetDateUntil())

	model.SetMaxResults(exampleRepMax)
	assert.True(t, model.HasMaxResults())
	assert.Equal(t, exampleRepMax, model.GetMaxResults())

	model.SetMerchantId(exampleRepMerchant)
	assert.True(t, model.HasMerchantId())
	assert.Equal(t, exampleRepMerchant, model.GetMerchantId())

	model.SetNextToken(exampleRepNext)
	assert.True(t, model.HasNextToken())
	assert.Equal(t, exampleRepNext, model.GetNextToken())

	model.SetOrderBy(exampleRepOrder)
	assert.True(t, model.HasOrderBy())
	assert.Equal(t, exampleRepOrder, model.GetOrderBy())
}

func TestRemittanceReportRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleRemittanceReportRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.RemittanceReportRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasMerchantId())
	assert.Equal(t, exampleRepMerchant, unmarshalled.GetMerchantId())
}

func TestRemittanceReportRequestToMap(t *testing.T) {
	model := buildExampleRemittanceReportRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "maxResults") {
		assert.Equal(t, &exampleRepMax, m["maxResults"])
	}
}

func TestNullableRemittanceReportRequestGetSet(t *testing.T) {
	base := buildExampleRemittanceReportRequest()
	n := openapiclient.NullableRemittanceReportRequest{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableRemittanceReportRequestUnset(t *testing.T) {
	base := buildExampleRemittanceReportRequest()
	n := openapiclient.NewNullableRemittanceReportRequest(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableRemittanceReportRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleRemittanceReportRequest()
	n := openapiclient.NewNullableRemittanceReportRequest(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableRemittanceReportRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleRepOrder, newN.Get().GetOrderBy())
	}
}
