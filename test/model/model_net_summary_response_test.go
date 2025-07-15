package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleCreditAmount = "10.00"
var exampleCreditCount int32 = 2
var exampleCreditValue int32 = 1000
var exampleDebitAmount = "20.00"
var exampleDebitCount int32 = 3
var exampleDebitValue int32 = 2000
var exampleNetAmount int32 = 1000
var exampleTotalCount int32 = 5

func buildExampleNetSummary() openapiclient.NetSummaryResponse {
	n := openapiclient.NewNetSummaryResponse()
	n.SetCreditItemsAmount(exampleCreditAmount)
	n.SetCreditItemsCount(exampleCreditCount)
	n.SetCreditItemsValue(exampleCreditValue)
	n.SetDebitItemsAmount(exampleDebitAmount)
	n.SetDebitItemsCount(exampleDebitCount)
	n.SetDebitItemsValue(exampleDebitValue)
	n.SetNetAmount(exampleNetAmount)
	n.SetTotalCount(exampleTotalCount)
	return *n
}

func TestNewNetSummaryResponse(t *testing.T) {
	model := openapiclient.NewNetSummaryResponse()
	require.NotNil(t, model)
	assert.False(t, model.HasCreditItemsAmount())
	assert.False(t, model.HasCreditItemsCount())
	assert.False(t, model.HasCreditItemsValue())
	assert.False(t, model.HasDebitItemsAmount())
	assert.False(t, model.HasDebitItemsCount())
	assert.False(t, model.HasDebitItemsValue())
	assert.False(t, model.HasNetAmount())
	assert.False(t, model.HasTotalCount())
}

func TestNewNetSummaryResponseWithDefaults(t *testing.T) {
	model := openapiclient.NewNetSummaryResponseWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasCreditItemsAmount())
	assert.False(t, model.HasCreditItemsCount())
	assert.False(t, model.HasCreditItemsValue())
	assert.False(t, model.HasDebitItemsAmount())
	assert.False(t, model.HasDebitItemsCount())
	assert.False(t, model.HasDebitItemsValue())
	assert.False(t, model.HasNetAmount())
	assert.False(t, model.HasTotalCount())
}

func TestNetSummaryResponseSetGetCycle(t *testing.T) {
	model := openapiclient.NewNetSummaryResponse()
	model.SetCreditItemsAmount(exampleCreditAmount)
	assert.True(t, model.HasCreditItemsAmount())
	assert.Equal(t, exampleCreditAmount, model.GetCreditItemsAmount())

	model.SetCreditItemsCount(exampleCreditCount)
	assert.True(t, model.HasCreditItemsCount())
	assert.Equal(t, exampleCreditCount, model.GetCreditItemsCount())

	model.SetCreditItemsValue(exampleCreditValue)
	assert.True(t, model.HasCreditItemsValue())
	assert.Equal(t, exampleCreditValue, model.GetCreditItemsValue())

	model.SetDebitItemsAmount(exampleDebitAmount)
	assert.True(t, model.HasDebitItemsAmount())
	assert.Equal(t, exampleDebitAmount, model.GetDebitItemsAmount())

	model.SetDebitItemsCount(exampleDebitCount)
	assert.True(t, model.HasDebitItemsCount())
	assert.Equal(t, exampleDebitCount, model.GetDebitItemsCount())

	model.SetDebitItemsValue(exampleDebitValue)
	assert.True(t, model.HasDebitItemsValue())
	assert.Equal(t, exampleDebitValue, model.GetDebitItemsValue())

	model.SetNetAmount(exampleNetAmount)
	assert.True(t, model.HasNetAmount())
	assert.Equal(t, exampleNetAmount, model.GetNetAmount())

	model.SetTotalCount(exampleTotalCount)
	assert.True(t, model.HasTotalCount())
	assert.Equal(t, exampleTotalCount, model.GetTotalCount())
}

func TestNetSummaryResponseJSONRoundTrip(t *testing.T) {
	model := buildExampleNetSummary()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.NetSummaryResponse
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasDebitItemsAmount())
	assert.Equal(t, exampleNetAmount, unmarshalled.GetNetAmount())
}

func TestNetSummaryResponseToMap(t *testing.T) {
	model := buildExampleNetSummary()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "credit_items_amount") {
		assert.Equal(t, &exampleCreditAmount, m["credit_items_amount"])
	}
	if assert.Contains(t, m, "credit_items_count") {
		assert.Equal(t, &exampleCreditCount, m["credit_items_count"])
	}
	if assert.Contains(t, m, "credit_items_value") {
		assert.Equal(t, &exampleCreditValue, m["credit_items_value"])
	}
	if assert.Contains(t, m, "debit_items_amount") {
		assert.Equal(t, &exampleDebitAmount, m["debit_items_amount"])
	}
	if assert.Contains(t, m, "debit_items_count") {
		assert.Equal(t, &exampleDebitCount, m["debit_items_count"])
	}
	if assert.Contains(t, m, "debit_items_value") {
		assert.Equal(t, &exampleDebitValue, m["debit_items_value"])
	}
	if assert.Contains(t, m, "net_amount") {
		assert.Equal(t, &exampleNetAmount, m["net_amount"])
	}
	if assert.Contains(t, m, "total_count") {
		assert.Equal(t, &exampleTotalCount, m["total_count"])
	}
}

func TestNullableNetSummaryResponseGetSet(t *testing.T) {
	base := buildExampleNetSummary()
	n := openapiclient.NullableNetSummaryResponse{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableNetSummaryResponseUnset(t *testing.T) {
	base := buildExampleNetSummary()
	n := openapiclient.NewNullableNetSummaryResponse(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableNetSummaryResponseJSONRoundTrip(t *testing.T) {
	base := buildExampleNetSummary()
	n := openapiclient.NewNullableNetSummaryResponse(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableNetSummaryResponse
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleTotalCount, newN.Get().GetTotalCount())
	}
}
