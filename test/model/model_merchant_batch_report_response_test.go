package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var exampleMBRBatchNo = "BN1"
var exampleMBRStatus = "Open"
var exampleMBRCode = "O"
var exampleMBRCurrency = "GBP"
var exampleMBRMerchantID int32 = 9
var exampleMBRCount int32 = 1
var exampleMBRMax int32 = 10
var exampleMBRToken = "tok"
var exampleMBRTime = time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)

func buildMBReportNetSummary() openapiclient.NetSummaryResponse {
	n := openapiclient.NewNetSummaryResponse()
	val := int32(100)
	n.SetNetAmount(val)
	return *n
}

func buildMBReportMerchantBatchResponse() openapiclient.MerchantBatchResponse {
	m := openapiclient.NewMerchantBatchResponse()
	m.SetBatchClosed(exampleMBRTime)
	m.SetBatchNo(exampleMBRBatchNo)
	m.SetBatchStatus(exampleMBRStatus)
	m.SetBatchStatusCode(exampleMBRCode)
	m.SetCurrency(exampleMBRCurrency)
	m.SetMerchantid(exampleMBRMerchantID)
	summary := buildMBReportNetSummary()
	m.SetNetSummary(summary)
	return *m
}

func buildExampleMerchantBatchReportResponse() openapiclient.MerchantBatchReportResponse {
	m := openapiclient.NewMerchantBatchReportResponse([]openapiclient.MerchantBatchResponse{buildMBReportMerchantBatchResponse()})
	m.SetCount(exampleMBRCount)
	m.SetMaxResults(exampleMBRMax)
	m.SetNextToken(exampleMBRToken)
	return *m
}

func TestNewMerchantBatchReportResponse(t *testing.T) {
	model := openapiclient.NewMerchantBatchReportResponse([]openapiclient.MerchantBatchResponse{buildMBReportMerchantBatchResponse()})
	require.NotNil(t, model)
	assert.Equal(t, 1, len(model.GetBatches()))
	assert.False(t, model.HasCount())
	assert.False(t, model.HasMaxResults())
	assert.False(t, model.HasNextToken())
}

func TestNewMerchantBatchReportResponseWithDefaults(t *testing.T) {
	model := openapiclient.NewMerchantBatchReportResponseWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, 0, len(model.GetBatches()))
}

func TestMerchantBatchReportResponseSetGetCycle(t *testing.T) {
	model := openapiclient.NewMerchantBatchReportResponse([]openapiclient.MerchantBatchResponse{buildMBReportMerchantBatchResponse()})
	model.SetCount(exampleMBRCount)
	assert.True(t, model.HasCount())
	assert.Equal(t, exampleMBRCount, model.GetCount())
	model.SetMaxResults(exampleMBRMax)
	assert.True(t, model.HasMaxResults())
	assert.Equal(t, exampleMBRMax, model.GetMaxResults())
	model.SetNextToken(exampleMBRToken)
	assert.True(t, model.HasNextToken())
	assert.Equal(t, exampleMBRToken, model.GetNextToken())
}

func TestMerchantBatchReportResponseJSONRoundTrip(t *testing.T) {
	model := buildExampleMerchantBatchReportResponse()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.MerchantBatchReportResponse
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleMBRMax, unmarshalled.GetMaxResults())
	assert.True(t, unmarshalled.HasNextToken())
}

func TestMerchantBatchReportResponseToMap(t *testing.T) {
	model := buildExampleMerchantBatchReportResponse()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "batches") {
		assert.Equal(t, model.Batches, m["batches"])
	}
	if assert.Contains(t, m, "count") {
		assert.Equal(t, &exampleMBRCount, m["count"])
	}
	if assert.Contains(t, m, "maxResults") {
		assert.Equal(t, &exampleMBRMax, m["maxResults"])
	}
	if assert.Contains(t, m, "nextToken") {
		assert.Equal(t, &exampleMBRToken, m["nextToken"])
	}
}

func TestNullableMerchantBatchReportResponseGetSet(t *testing.T) {
	base := buildExampleMerchantBatchReportResponse()
	n := openapiclient.NullableMerchantBatchReportResponse{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableMerchantBatchReportResponseUnset(t *testing.T) {
	base := buildExampleMerchantBatchReportResponse()
	n := openapiclient.NewNullableMerchantBatchReportResponse(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableMerchantBatchReportResponseJSONRoundTrip(t *testing.T) {
	base := buildExampleMerchantBatchReportResponse()
	n := openapiclient.NewNullableMerchantBatchReportResponse(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableMerchantBatchReportResponse
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleMBRToken, newN.Get().GetNextToken())
	}
}
