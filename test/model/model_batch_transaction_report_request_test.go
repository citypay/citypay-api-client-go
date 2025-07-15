package citypay

import (
	"encoding/json"
	"testing"

	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var exampleBTRRMaxResults int32 = 20
var exampleBTRRNextToken = "token"
var exampleBTRROrderBy = "trans_no"

func buildExampleBatchTransactionReportRequest() *openapiclient.BatchTransactionReportRequest {
	r := openapiclient.NewBatchTransactionReportRequest()
	r.SetMaxResults(exampleBTRRMaxResults)
	r.SetNextToken(exampleBTRRNextToken)
	r.SetOrderBy(exampleBTRROrderBy)
	return r
}

func TestNewBatchTransactionReportRequest(t *testing.T) {
	model := openapiclient.NewBatchTransactionReportRequest()
	require.NotNil(t, model)
	assert.False(t, model.HasMaxResults())
	assert.False(t, model.HasNextToken())
	assert.False(t, model.HasOrderBy())
}

func TestNewBatchTransactionReportRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewBatchTransactionReportRequestWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasMaxResults())
	assert.False(t, model.HasNextToken())
	assert.False(t, model.HasOrderBy())
}

func TestBatchTransactionReportRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewBatchTransactionReportRequest()
	model.SetMaxResults(exampleBTRRMaxResults)
	assert.True(t, model.HasMaxResults())
	assert.Equal(t, exampleBTRRMaxResults, model.GetMaxResults())
	val, ok := model.GetMaxResultsOk()
	require.True(t, ok)
	if assert.NotNil(t, val) {
		assert.Equal(t, exampleBTRRMaxResults, *val)
	}
	model.SetNextToken(exampleBTRRNextToken)
	assert.True(t, model.HasNextToken())
	assert.Equal(t, exampleBTRRNextToken, model.GetNextToken())
	val2, ok2 := model.GetNextTokenOk()
	require.True(t, ok2)
	if assert.NotNil(t, val2) {
		assert.Equal(t, exampleBTRRNextToken, *val2)
	}
	model.SetOrderBy(exampleBTRROrderBy)
	assert.True(t, model.HasOrderBy())
	assert.Equal(t, exampleBTRROrderBy, model.GetOrderBy())
}

func TestBatchTransactionReportRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleBatchTransactionReportRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.BatchTransactionReportRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasMaxResults())
	assert.Equal(t, exampleBTRRNextToken, unmarshalled.GetNextToken())
}

func TestBatchTransactionReportRequestToMap(t *testing.T) {
	model := buildExampleBatchTransactionReportRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "maxResults") {
		assert.Equal(t, &exampleBTRRMaxResults, m["maxResults"])
	}
	if assert.Contains(t, m, "nextToken") {
		assert.Equal(t, &exampleBTRRNextToken, m["nextToken"])
	}
	if assert.Contains(t, m, "orderBy") {
		assert.Equal(t, &exampleBTRROrderBy, m["orderBy"])
	}
}

func TestNullableBatchTransactionReportRequestGetSet(t *testing.T) {
	base := buildExampleBatchTransactionReportRequest()
	n := openapiclient.NullableBatchTransactionReportRequest{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableBatchTransactionReportRequestUnset(t *testing.T) {
	base := buildExampleBatchTransactionReportRequest()
	n := openapiclient.NewNullableBatchTransactionReportRequest(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableBatchTransactionReportRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleBatchTransactionReportRequest()
	n := openapiclient.NewNullableBatchTransactionReportRequest(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableBatchTransactionReportRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleBTRRNextToken, newN.Get().GetNextToken())
	}
}
