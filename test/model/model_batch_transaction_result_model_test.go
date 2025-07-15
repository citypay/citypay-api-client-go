package citypay

import (
	"encoding/json"
	"testing"
	"time"

	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var exampleResAccountId = "ac123"
var exampleResAmount int32 = 10
var exampleResAuthcode = "AUTH"
var exampleResDatetime = time.Date(2023, time.January, 1, 2, 3, 4, 0, time.UTC)
var exampleResIdentifier = "id001"
var exampleResMaskedpan = "411111******1111"
var exampleResMerchantid int32 = 44
var exampleResMessage = "OK"
var exampleResResult int32 = 1
var exampleResResultCode = "A1"
var exampleResScheme = "VISA"
var exampleResSchemeId = "VI"
var exampleResSchemeLogo = "logo"
var exampleResTransno int32 = 99

func buildExampleBatchTransactionResultModel() *openapiclient.BatchTransactionResultModel {
	m := openapiclient.NewBatchTransactionResultModel(exampleResAccountId, exampleResIdentifier, exampleResMerchantid, exampleResMessage, exampleResResult, exampleResResultCode)
	m.SetAmount(exampleResAmount)
	m.SetAuthcode(exampleResAuthcode)
	m.SetDatetime(exampleResDatetime)
	m.SetMaskedpan(exampleResMaskedpan)
	m.SetScheme(exampleResScheme)
	m.SetSchemeId(exampleResSchemeId)
	m.SetSchemeLogo(exampleResSchemeLogo)
	m.SetTransno(exampleResTransno)
	return m
}

func TestNewBatchTransactionResultModel(t *testing.T) {
	model := openapiclient.NewBatchTransactionResultModel(exampleResAccountId, exampleResIdentifier, exampleResMerchantid, exampleResMessage, exampleResResult, exampleResResultCode)
	require.NotNil(t, model)
	assert.Equal(t, exampleResAccountId, model.GetAccountId())
	assert.Equal(t, exampleResIdentifier, model.GetIdentifier())
	assert.Equal(t, exampleResMerchantid, model.GetMerchantid())
	assert.Equal(t, exampleResMessage, model.GetMessage())
	assert.Equal(t, exampleResResult, model.GetResult())
	assert.Equal(t, exampleResResultCode, model.GetResultCode())
	assert.False(t, model.HasAmount())
	assert.False(t, model.HasAuthcode())
	assert.False(t, model.HasDatetime())
	assert.False(t, model.HasMaskedpan())
	assert.False(t, model.HasScheme())
	assert.False(t, model.HasSchemeId())
	assert.False(t, model.HasSchemeLogo())
	assert.False(t, model.HasTransno())
}

func TestNewBatchTransactionResultModelWithDefaults(t *testing.T) {
	model := openapiclient.NewBatchTransactionResultModelWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, "", model.GetAccountId())
	assert.Equal(t, "", model.GetIdentifier())
	assert.Equal(t, int32(0), model.GetMerchantid())
	assert.Equal(t, "", model.GetMessage())
	assert.Equal(t, int32(0), model.GetResult())
	assert.Equal(t, "", model.GetResultCode())
}

func TestBatchTransactionResultModelSetGetCycle(t *testing.T) {
	model := openapiclient.NewBatchTransactionResultModel(exampleResAccountId, exampleResIdentifier, exampleResMerchantid, exampleResMessage, exampleResResult, exampleResResultCode)
	model.SetAmount(exampleResAmount)
	assert.True(t, model.HasAmount())
	assert.Equal(t, exampleResAmount, model.GetAmount())
	model.SetAuthcode(exampleResAuthcode)
	assert.True(t, model.HasAuthcode())
	assert.Equal(t, exampleResAuthcode, model.GetAuthcode())
	model.SetDatetime(exampleResDatetime)
	assert.True(t, model.HasDatetime())
	model.SetMaskedpan(exampleResMaskedpan)
	assert.True(t, model.HasMaskedpan())
	model.SetScheme(exampleResScheme)
	assert.True(t, model.HasScheme())
	model.SetSchemeId(exampleResSchemeId)
	assert.True(t, model.HasSchemeId())
	model.SetSchemeLogo(exampleResSchemeLogo)
	assert.True(t, model.HasSchemeLogo())
	model.SetTransno(exampleResTransno)
	assert.True(t, model.HasTransno())
}

func TestBatchTransactionResultModelJSONRoundTrip(t *testing.T) {
	model := buildExampleBatchTransactionResultModel()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.BatchTransactionResultModel
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleResAccountId, unmarshalled.GetAccountId())
	assert.True(t, unmarshalled.HasAmount())
	assert.Equal(t, exampleResTransno, unmarshalled.GetTransno())
}

func TestBatchTransactionResultModelToMap(t *testing.T) {
	model := buildExampleBatchTransactionResultModel()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "account_id") {
		assert.Equal(t, exampleResAccountId, m["account_id"])
	}
	if assert.Contains(t, m, "amount") {
		assert.Equal(t, &exampleResAmount, m["amount"])
	}
	if assert.Contains(t, m, "scheme_id") {
		assert.Equal(t, &exampleResSchemeId, m["scheme_id"])
	}
}

func TestNullableBatchTransactionResultModelGetSet(t *testing.T) {
	base := buildExampleBatchTransactionResultModel()
	n := openapiclient.NullableBatchTransactionResultModel{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableBatchTransactionResultModelUnset(t *testing.T) {
	base := buildExampleBatchTransactionResultModel()
	n := openapiclient.NewNullableBatchTransactionResultModel(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableBatchTransactionResultModelJSONRoundTrip(t *testing.T) {
	base := buildExampleBatchTransactionResultModel()
	n := openapiclient.NewNullableBatchTransactionResultModel(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableBatchTransactionResultModel
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleResAccountId, newN.Get().GetAccountId())
	}
}
