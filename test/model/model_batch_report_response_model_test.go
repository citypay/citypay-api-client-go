package citypay

import (
	"encoding/json"
	"testing"
	"time"

	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var exampleBRRAmount int32 = 1000
var exampleBRRBatchDate = "2023-02-01T15:04:05Z"
var exampleBRRRespBatchId int32 = 55
var exampleBRRBatchStatus = "COMPLETE"
var exampleBRRRespClientAccountId = "clientB"

var exampleBRRResAccountId = "acc123"
var exampleBRRResAmount int32 = 10
var exampleBRRResAuthcode = "AUTH1"
var exampleBRRResDatetime = time.Date(2023, time.March, 4, 5, 6, 7, 0, time.UTC)
var exampleBRRResIdentifier = "id99"
var exampleBRRResMaskedpan = "411111******1111"
var exampleBRRResMerchantid int32 = 77
var exampleBRRResMessage = "Approved"
var exampleBRRResResult int32 = 1
var exampleBRRResResultCode = "OK"
var exampleBRRResScheme = "VISA"
var exampleBRRResSchemeId = "VI"
var exampleBRRResSchemeLogo = "logo"
var exampleBRRResTransno int32 = 101

func buildExampleBatchTransactionResult() openapiclient.BatchTransactionResultModel {
	r := openapiclient.NewBatchTransactionResultModel(exampleBRRResAccountId, exampleBRRResIdentifier, exampleBRRResMerchantid, exampleBRRResMessage, exampleBRRResResult, exampleBRRResResultCode)
	r.SetAmount(exampleBRRResAmount)
	r.SetAuthcode(exampleBRRResAuthcode)
	r.SetDatetime(exampleBRRResDatetime)
	r.SetMaskedpan(exampleBRRResMaskedpan)
	r.SetScheme(exampleBRRResScheme)
	r.SetSchemeId(exampleBRRResSchemeId)
	r.SetSchemeLogo(exampleBRRResSchemeLogo)
	r.SetTransno(exampleBRRResTransno)
	return *r
}

func buildExampleBatchReportResponseModel() *openapiclient.BatchReportResponseModel {
	tr := buildExampleBatchTransactionResult()
	m := openapiclient.NewBatchReportResponseModel(exampleBRRAmount, exampleBRRBatchDate, exampleBRRRespBatchId, exampleBRRBatchStatus, exampleBRRRespClientAccountId, []openapiclient.BatchTransactionResultModel{tr})
	return m
}

func TestNewBatchReportResponseModel(t *testing.T) {
	tr := buildExampleBatchTransactionResult()
	model := openapiclient.NewBatchReportResponseModel(exampleBRRAmount, exampleBRRBatchDate, exampleBRRRespBatchId, exampleBRRBatchStatus, exampleBRRRespClientAccountId, []openapiclient.BatchTransactionResultModel{tr})
	require.NotNil(t, model)
	assert.Equal(t, exampleBRRAmount, model.GetAmount())
	assert.Equal(t, exampleBRRBatchDate, model.GetBatchDate())
	assert.Equal(t, exampleBRRRespBatchId, model.GetBatchId())
	assert.Equal(t, exampleBRRBatchStatus, model.GetBatchStatus())
	assert.Equal(t, exampleBRRRespClientAccountId, model.GetClientAccountId())
	assert.Equal(t, 1, len(model.GetTransactions()))
}

func TestNewBatchReportResponseModelWithDefaults(t *testing.T) {
	model := openapiclient.NewBatchReportResponseModelWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, int32(0), model.GetAmount())
	assert.Equal(t, "", model.GetBatchDate())
	assert.Equal(t, int32(0), model.GetBatchId())
	assert.Equal(t, "", model.GetBatchStatus())
	assert.Equal(t, "", model.GetClientAccountId())
	assert.Equal(t, 0, len(model.GetTransactions()))
}

func TestBatchReportResponseModelSetGetCycle(t *testing.T) {
	tr := buildExampleBatchTransactionResult()
	model := openapiclient.NewBatchReportResponseModel(exampleBRRAmount, exampleBRRBatchDate, exampleBRRRespBatchId, exampleBRRBatchStatus, exampleBRRRespClientAccountId, []openapiclient.BatchTransactionResultModel{tr})
	model.SetAmount(2000)
	assert.Equal(t, int32(2000), model.GetAmount())
	model.SetBatchDate("2024-01-01")
	assert.Equal(t, "2024-01-01", model.GetBatchDate())
	model.SetBatchId(66)
	assert.Equal(t, int32(66), model.GetBatchId())
	model.SetBatchStatus("QUEUED")
	assert.Equal(t, "QUEUED", model.GetBatchStatus())
	model.SetClientAccountId("other")
	assert.Equal(t, "other", model.GetClientAccountId())
	model.SetTransactions([]openapiclient.BatchTransactionResultModel{tr, tr})
	assert.Equal(t, 2, len(model.GetTransactions()))
}

func TestBatchReportResponseModelJSONRoundTrip(t *testing.T) {
	model := buildExampleBatchReportResponseModel()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.BatchReportResponseModel
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleBRRAmount, unmarshalled.GetAmount())
	assert.Equal(t, exampleBRRBatchDate, unmarshalled.GetBatchDate())
	assert.Equal(t, exampleBRRBatchStatus, unmarshalled.GetBatchStatus())
	assert.Equal(t, 1, len(unmarshalled.GetTransactions()))
}

func TestBatchReportResponseModelToMap(t *testing.T) {
	model := buildExampleBatchReportResponseModel()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "amount") {
		assert.Equal(t, exampleBRRAmount, m["amount"])
	}
	if assert.Contains(t, m, "batch_status") {
		assert.Equal(t, exampleBRRBatchStatus, m["batch_status"])
	}
	if assert.Contains(t, m, "transactions") {
		assert.Equal(t, 1, len(m["transactions"].([]openapiclient.BatchTransactionResultModel)))
	}
}

func TestNullableBatchReportResponseModelGetSet(t *testing.T) {
	base := buildExampleBatchReportResponseModel()
	n := openapiclient.NullableBatchReportResponseModel{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableBatchReportResponseModelUnset(t *testing.T) {
	base := buildExampleBatchReportResponseModel()
	n := openapiclient.NewNullableBatchReportResponseModel(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableBatchReportResponseModelJSONRoundTrip(t *testing.T) {
	base := buildExampleBatchReportResponseModel()
	n := openapiclient.NewNullableBatchReportResponseModel(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableBatchReportResponseModel
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleBRRBatchStatus, newN.Get().GetBatchStatus())
	}
}
