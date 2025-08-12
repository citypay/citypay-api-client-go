package citypay

import (
	"encoding/json"
	"testing"

	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var exampleBRRBatchId int32 = 99
var exampleBRRClientAccountId = "clientA"

func buildExampleBatchReportRequest() *openapiclient.BatchReportRequest {
	r := openapiclient.NewBatchReportRequest(exampleBRRBatchId)
	r.SetClientAccountId(exampleBRRClientAccountId)
	return r
}

func TestNewBatchReportRequest(t *testing.T) {
	model := openapiclient.NewBatchReportRequest(exampleBRRBatchId)
	require.NotNil(t, model)
	assert.Equal(t, exampleBRRBatchId, model.GetBatchId())
	assert.False(t, model.HasClientAccountId())
}

func TestNewBatchReportRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewBatchReportRequestWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, int32(0), model.GetBatchId())
}

func TestBatchReportRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewBatchReportRequest(exampleBRRBatchId)
	model.SetClientAccountId(exampleBRRClientAccountId)
	assert.True(t, model.HasClientAccountId())
	assert.Equal(t, exampleBRRClientAccountId, model.GetClientAccountId())
	val, ok := model.GetClientAccountIdOk()
	require.True(t, ok)
	if assert.NotNil(t, val) {
		assert.Equal(t, exampleBRRClientAccountId, *val)
	}
}

func TestBatchReportRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleBatchReportRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.BatchReportRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleBRRBatchId, unmarshalled.GetBatchId())
	assert.True(t, unmarshalled.HasClientAccountId())
	assert.Equal(t, exampleBRRClientAccountId, unmarshalled.GetClientAccountId())
}

func TestBatchReportRequestToMap(t *testing.T) {
	model := buildExampleBatchReportRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "batch_id") {
		assert.Equal(t, exampleBRRBatchId, m["batch_id"])
	}
	if assert.Contains(t, m, "client_account_id") {
		assert.Equal(t, &exampleBRRClientAccountId, m["client_account_id"])
	}
}

func TestNullableBatchReportRequestGetSet(t *testing.T) {
	base := buildExampleBatchReportRequest()
	n := openapiclient.NullableBatchReportRequest{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableBatchReportRequestUnset(t *testing.T) {
	base := buildExampleBatchReportRequest()
	n := openapiclient.NewNullableBatchReportRequest(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableBatchReportRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleBatchReportRequest()
	n := openapiclient.NewNullableBatchReportRequest(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableBatchReportRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleBRRClientAccountId, newN.Get().GetClientAccountId())
	}
}
