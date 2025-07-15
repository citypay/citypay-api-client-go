package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleMBRRFrom = "2025-01-01"
var exampleMBRRUntil = "2025-01-31"
var exampleMBRRMax int32 = 50
var exampleMBRRMerchants = []int32{1, 2}
var exampleMBRRToken = "next"
var exampleMBRROrder = "merchant_id"

func buildExampleMerchantBatchReportRequest() openapiclient.MerchantBatchReportRequest {
	m := openapiclient.NewMerchantBatchReportRequest()
	m.SetDateFrom(exampleMBRRFrom)
	m.SetDateUntil(exampleMBRRUntil)
	m.SetMaxResults(exampleMBRRMax)
	m.SetMerchantId(exampleMBRRMerchants)
	m.SetNextToken(exampleMBRRToken)
	m.SetOrderBy(exampleMBRROrder)
	return *m
}

func TestNewMerchantBatchReportRequest(t *testing.T) {
	model := openapiclient.NewMerchantBatchReportRequest()
	require.NotNil(t, model)
	assert.False(t, model.HasDateFrom())
	assert.False(t, model.HasDateUntil())
	assert.False(t, model.HasMaxResults())
	assert.False(t, model.HasMerchantId())
	assert.False(t, model.HasNextToken())
	assert.False(t, model.HasOrderBy())
}

func TestNewMerchantBatchReportRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewMerchantBatchReportRequestWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasDateFrom())
	assert.False(t, model.HasDateUntil())
	assert.False(t, model.HasMaxResults())
	assert.False(t, model.HasMerchantId())
	assert.False(t, model.HasNextToken())
	assert.False(t, model.HasOrderBy())
}

func TestMerchantBatchReportRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewMerchantBatchReportRequest()
	model.SetDateFrom(exampleMBRRFrom)
	assert.True(t, model.HasDateFrom())
	assert.Equal(t, exampleMBRRFrom, model.GetDateFrom())
	if val, ok := model.GetDateFromOk(); assert.True(t, ok) {
		assert.Equal(t, exampleMBRRFrom, *val)
	}

	model.SetDateUntil(exampleMBRRUntil)
	assert.True(t, model.HasDateUntil())
	assert.Equal(t, exampleMBRRUntil, model.GetDateUntil())

	model.SetMaxResults(exampleMBRRMax)
	assert.True(t, model.HasMaxResults())
	assert.Equal(t, exampleMBRRMax, model.GetMaxResults())

	model.SetMerchantId(exampleMBRRMerchants)
	assert.True(t, model.HasMerchantId())
	assert.Equal(t, exampleMBRRMerchants, model.GetMerchantId())

	model.SetNextToken(exampleMBRRToken)
	assert.True(t, model.HasNextToken())
	assert.Equal(t, exampleMBRRToken, model.GetNextToken())

	model.SetOrderBy(exampleMBRROrder)
	assert.True(t, model.HasOrderBy())
	assert.Equal(t, exampleMBRROrder, model.GetOrderBy())
}

func TestMerchantBatchReportRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleMerchantBatchReportRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.MerchantBatchReportRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasDateFrom())
	assert.Equal(t, exampleMBRRMax, unmarshalled.GetMaxResults())
}

func TestMerchantBatchReportRequestToMap(t *testing.T) {
	model := buildExampleMerchantBatchReportRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "date_from") {
		assert.Equal(t, &exampleMBRRFrom, m["date_from"])
	}
	if assert.Contains(t, m, "date_until") {
		assert.Equal(t, &exampleMBRRUntil, m["date_until"])
	}
	if assert.Contains(t, m, "maxResults") {
		assert.Equal(t, &exampleMBRRMax, m["maxResults"])
	}
	if assert.Contains(t, m, "merchant_id") {
		assert.Equal(t, exampleMBRRMerchants, m["merchant_id"])
	}
	if assert.Contains(t, m, "nextToken") {
		assert.Equal(t, &exampleMBRRToken, m["nextToken"])
	}
	if assert.Contains(t, m, "orderBy") {
		assert.Equal(t, &exampleMBRROrder, m["orderBy"])
	}
}

func TestNullableMerchantBatchReportRequestGetSet(t *testing.T) {
	base := buildExampleMerchantBatchReportRequest()
	n := openapiclient.NullableMerchantBatchReportRequest{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableMerchantBatchReportRequestUnset(t *testing.T) {
	base := buildExampleMerchantBatchReportRequest()
	n := openapiclient.NewNullableMerchantBatchReportRequest(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableMerchantBatchReportRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleMerchantBatchReportRequest()
	n := openapiclient.NewNullableMerchantBatchReportRequest(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableMerchantBatchReportRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleMBRRUntil, newN.Get().GetDateUntil())
	}
}
