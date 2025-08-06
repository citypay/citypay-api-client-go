package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var exampleTransactionReportRequestFrom = time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
var exampleTransactionReportRequestUntil = time.Date(2024, time.January, 31, 0, 0, 0, 0, time.UTC)
var exampleTransactionReportRequestMerchantId int32 = 123456

func buildExampleTransactionReportRequest() *openapiclient.TransactionReportRequest {
	m := openapiclient.NewTransactionReportRequest(exampleTransactionReportRequestFrom, exampleTransactionReportRequestMerchantId, exampleTransactionReportRequestUntil)
	m.SetIncludeAuthorised(true)
	return m
}

func TestNewTransactionReportRequest(t *testing.T) {
	model := openapiclient.NewTransactionReportRequest(exampleTransactionReportRequestFrom, exampleTransactionReportRequestMerchantId, exampleTransactionReportRequestUntil)
	require.NotNil(t, model)
	assert.Equal(t, exampleTransactionReportRequestMerchantId, model.GetMerchantid())
	assert.False(t, model.HasIncludeAuthorised())
}

func TestNewTransactionReportRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewTransactionReportRequestWithDefaults()
	require.NotNil(t, model)
}

func TestTransactionReportRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewTransactionReportRequest(exampleTransactionReportRequestFrom, exampleTransactionReportRequestMerchantId, exampleTransactionReportRequestUntil)
	model.SetIncludeAuthorised(true)
	assert.True(t, model.HasIncludeAuthorised())
	assert.True(t, model.GetIncludeAuthorised())
}

func TestTransactionReportRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleTransactionReportRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.TransactionReportRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleTransactionReportRequestFrom, unmarshalled.GetFrom())
	assert.Equal(t, exampleTransactionReportRequestUntil, unmarshalled.GetUntil())
}

func TestTransactionReportRequestToMap(t *testing.T) {
	model := buildExampleTransactionReportRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "merchantid") {
		assert.Equal(t, exampleTransactionReportRequestMerchantId, m["merchantid"])
	}
}

func TestNullableTransactionReportRequestGetSet(t *testing.T) {
	base := buildExampleTransactionReportRequest()
	n := openapiclient.NullableTransactionReportRequest{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableTransactionReportRequestUnset(t *testing.T) {
	base := buildExampleTransactionReportRequest()
	n := openapiclient.NewNullableTransactionReportRequest(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableTransactionReportRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleTransactionReportRequest()
	n := openapiclient.NewNullableTransactionReportRequest(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableTransactionReportRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleTransactionReportRequestMerchantId, newN.Get().GetMerchantid())
	}
}
