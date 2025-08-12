package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var exampleMBRRespClosed = time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC)
var exampleMBRRespNo = "001"
var exampleMBRRespStatus = "Settled"
var exampleMBRRespCode = "S"
var exampleMBRRespCurrency = "GBP"
var exampleMBRRespMID int32 = 44

func buildExampleNetSummaryResponse() openapiclient.NetSummaryResponse {
	n := openapiclient.NewNetSummaryResponse()
	val := int32(55)
	n.SetNetAmount(val)
	return *n
}

func buildExampleMerchantBatchResponse() openapiclient.MerchantBatchResponse {
	m := openapiclient.NewMerchantBatchResponse()
	m.SetBatchClosed(exampleMBRRespClosed)
	m.SetBatchNo(exampleMBRRespNo)
	m.SetBatchStatus(exampleMBRRespStatus)
	m.SetBatchStatusCode(exampleMBRRespCode)
	m.SetCurrency(exampleMBRRespCurrency)
	m.SetMerchantid(exampleMBRRespMID)
	summary := buildExampleNetSummaryResponse()
	m.SetNetSummary(summary)
	return *m
}

func TestNewMerchantBatchResponse(t *testing.T) {
	model := openapiclient.NewMerchantBatchResponse()
	require.NotNil(t, model)
	assert.False(t, model.HasBatchClosed())
	assert.False(t, model.HasBatchNo())
	assert.False(t, model.HasBatchStatus())
	assert.False(t, model.HasBatchStatusCode())
	assert.False(t, model.HasCurrency())
	assert.False(t, model.HasMerchantid())
	assert.False(t, model.HasNetSummary())
}

func TestNewMerchantBatchResponseWithDefaults(t *testing.T) {
	model := openapiclient.NewMerchantBatchResponseWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasBatchClosed())
	assert.False(t, model.HasBatchNo())
	assert.False(t, model.HasBatchStatus())
	assert.False(t, model.HasBatchStatusCode())
	assert.False(t, model.HasCurrency())
	assert.False(t, model.HasMerchantid())
	assert.False(t, model.HasNetSummary())
}

func TestMerchantBatchResponseSetGetCycle(t *testing.T) {
	model := openapiclient.NewMerchantBatchResponse()
	model.SetBatchClosed(exampleMBRRespClosed)
	assert.True(t, model.HasBatchClosed())
	assert.Equal(t, exampleMBRRespClosed, model.GetBatchClosed())

	model.SetBatchNo(exampleMBRRespNo)
	assert.True(t, model.HasBatchNo())
	assert.Equal(t, exampleMBRRespNo, model.GetBatchNo())

	model.SetBatchStatus(exampleMBRRespStatus)
	assert.True(t, model.HasBatchStatus())
	assert.Equal(t, exampleMBRRespStatus, model.GetBatchStatus())

	model.SetBatchStatusCode(exampleMBRRespCode)
	assert.True(t, model.HasBatchStatusCode())
	assert.Equal(t, exampleMBRRespCode, model.GetBatchStatusCode())

	model.SetCurrency(exampleMBRRespCurrency)
	assert.True(t, model.HasCurrency())
	assert.Equal(t, exampleMBRRespCurrency, model.GetCurrency())

	model.SetMerchantid(exampleMBRRespMID)
	assert.True(t, model.HasMerchantid())
	assert.Equal(t, exampleMBRRespMID, model.GetMerchantid())

	summary := buildExampleNetSummaryResponse()
	model.SetNetSummary(summary)
	assert.True(t, model.HasNetSummary())
	assert.Equal(t, summary, model.GetNetSummary())
}

func TestMerchantBatchResponseJSONRoundTrip(t *testing.T) {
	model := buildExampleMerchantBatchResponse()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.MerchantBatchResponse
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasMerchantid())
	assert.Equal(t, exampleMBRRespStatus, unmarshalled.GetBatchStatus())
}

func TestMerchantBatchResponseToMap(t *testing.T) {
	model := buildExampleMerchantBatchResponse()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "batch_closed") {
		assert.Equal(t, &exampleMBRRespClosed, m["batch_closed"])
	}
	if assert.Contains(t, m, "batch_no") {
		assert.Equal(t, &exampleMBRRespNo, m["batch_no"])
	}
	if assert.Contains(t, m, "batch_status") {
		assert.Equal(t, &exampleMBRRespStatus, m["batch_status"])
	}
	if assert.Contains(t, m, "batch_status_code") {
		assert.Equal(t, &exampleMBRRespCode, m["batch_status_code"])
	}
	if assert.Contains(t, m, "currency") {
		assert.Equal(t, &exampleMBRRespCurrency, m["currency"])
	}
	if assert.Contains(t, m, "merchantid") {
		assert.Equal(t, &exampleMBRRespMID, m["merchantid"])
	}
	if assert.Contains(t, m, "net_summary") {
		assert.NotNil(t, m["net_summary"])
	}
}

func TestNullableMerchantBatchResponseGetSet(t *testing.T) {
	base := buildExampleMerchantBatchResponse()
	n := openapiclient.NullableMerchantBatchResponse{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableMerchantBatchResponseUnset(t *testing.T) {
	base := buildExampleMerchantBatchResponse()
	n := openapiclient.NewNullableMerchantBatchResponse(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableMerchantBatchResponseJSONRoundTrip(t *testing.T) {
	base := buildExampleMerchantBatchResponse()
	n := openapiclient.NewNullableMerchantBatchResponse(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableMerchantBatchResponse
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleMBRRespNo, newN.Get().GetBatchNo())
	}
}
