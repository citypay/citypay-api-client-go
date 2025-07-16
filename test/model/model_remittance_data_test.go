package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var exampleRemitDate = time.Date(2024, time.January, 2, 0, 0, 0, 0, time.UTC)
var exampleRemitNet int32 = 50
var exampleRemitRefund int32 = 5
var exampleRemitRefundCount int32 = 1
var exampleRemitSales int32 = 55
var exampleRemitSalesCount int32 = 2

func buildExampleRemittanceData() openapiclient.RemittanceData {
	r := openapiclient.NewRemittanceData()
	r.SetDateCreated(exampleRemitDate)
	r.SetNetAmount(exampleRemitNet)
	r.SetRefundAmount(exampleRemitRefund)
	r.SetRefundCount(exampleRemitRefundCount)
	r.SetSalesAmount(exampleRemitSales)
	r.SetSalesCount(exampleRemitSalesCount)
	return *r
}

func TestNewRemittanceData(t *testing.T) {
	model := openapiclient.NewRemittanceData()
	require.NotNil(t, model)
	assert.False(t, model.HasNetAmount())
}

func TestNewRemittanceDataWithDefaults(t *testing.T) {
	model := openapiclient.NewRemittanceDataWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasRefundCount())
}

func TestRemittanceDataSetGetCycle(t *testing.T) {
	model := openapiclient.NewRemittanceData()
	model.SetDateCreated(exampleRemitDate)
	assert.True(t, model.HasDateCreated())
	assert.Equal(t, exampleRemitDate, model.GetDateCreated())

	model.SetNetAmount(exampleRemitNet)
	assert.True(t, model.HasNetAmount())
	assert.Equal(t, exampleRemitNet, model.GetNetAmount())

	model.SetRefundAmount(exampleRemitRefund)
	assert.True(t, model.HasRefundAmount())
	assert.Equal(t, exampleRemitRefund, model.GetRefundAmount())

	model.SetRefundCount(exampleRemitRefundCount)
	assert.True(t, model.HasRefundCount())
	assert.Equal(t, exampleRemitRefundCount, model.GetRefundCount())

	model.SetSalesAmount(exampleRemitSales)
	assert.True(t, model.HasSalesAmount())
	assert.Equal(t, exampleRemitSales, model.GetSalesAmount())

	model.SetSalesCount(exampleRemitSalesCount)
	assert.True(t, model.HasSalesCount())
	assert.Equal(t, exampleRemitSalesCount, model.GetSalesCount())
}

func TestRemittanceDataJSONRoundTrip(t *testing.T) {
	model := buildExampleRemittanceData()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.RemittanceData
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasRefundCount())
	assert.Equal(t, exampleRemitRefundCount, unmarshalled.GetRefundCount())
}

func TestRemittanceDataToMap(t *testing.T) {
	model := buildExampleRemittanceData()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "net_amount") {
		assert.Equal(t, &exampleRemitNet, m["net_amount"])
	}
}

func TestNullableRemittanceDataGetSet(t *testing.T) {
	base := buildExampleRemittanceData()
	n := openapiclient.NullableRemittanceData{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableRemittanceDataUnset(t *testing.T) {
	base := buildExampleRemittanceData()
	n := openapiclient.NewNullableRemittanceData(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableRemittanceDataJSONRoundTrip(t *testing.T) {
	base := buildExampleRemittanceData()
	n := openapiclient.NewNullableRemittanceData(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableRemittanceData
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleRemitSales, newN.Get().GetSalesAmount())
	}
}
