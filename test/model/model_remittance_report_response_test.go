package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var exampleRemitRespCount int32 = 2
var exampleRemitRespMax int32 = 10
var exampleRemitRespNext = "next"

func buildExampleMerchantBatchResponse() openapiclient.MerchantBatchResponse {
	m := openapiclient.NewMerchantBatchResponse()
	m.SetBatchId(1)
	return *m
}

func buildExampleRemittanceDataSlice() []openapiclient.RemittanceData {
	r := openapiclient.NewRemittanceData()
	r.SetDateCreated(time.Now())
	return []openapiclient.RemittanceData{*r}
}

func buildExampleRemittedClientData() openapiclient.RemittedClientData {
	d := openapiclient.NewRemittedClientData([]openapiclient.MerchantBatchResponse{buildExampleMerchantBatchResponse()}, buildExampleRemittanceDataSlice())
	d.SetClientid("c1")
	return *d
}

func buildExampleRemittanceReportResponse() openapiclient.RemittanceReportResponse {
	r := openapiclient.NewRemittanceReportResponse([]openapiclient.RemittedClientData{buildExampleRemittedClientData()})
	r.SetCount(exampleRemitRespCount)
	r.SetMaxResults(exampleRemitRespMax)
	r.SetNextToken(exampleRemitRespNext)
	return *r
}

func TestNewRemittanceReportResponse(t *testing.T) {
	model := openapiclient.NewRemittanceReportResponse([]openapiclient.RemittedClientData{})
	require.NotNil(t, model)
	assert.NotNil(t, model.GetData())
	assert.False(t, model.HasCount())
}

func TestNewRemittanceReportResponseWithDefaults(t *testing.T) {
	model := openapiclient.NewRemittanceReportResponseWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasMaxResults())
}

func TestRemittanceReportResponseSetGetCycle(t *testing.T) {
	model := openapiclient.NewRemittanceReportResponse([]openapiclient.RemittedClientData{})
	model.SetCount(exampleRemitRespCount)
	assert.True(t, model.HasCount())
	assert.Equal(t, exampleRemitRespCount, model.GetCount())

	model.SetData([]openapiclient.RemittedClientData{buildExampleRemittedClientData()})
	assert.Len(t, model.GetData(), 1)

	model.SetMaxResults(exampleRemitRespMax)
	assert.True(t, model.HasMaxResults())
	assert.Equal(t, exampleRemitRespMax, model.GetMaxResults())

	model.SetNextToken(exampleRemitRespNext)
	assert.True(t, model.HasNextToken())
	assert.Equal(t, exampleRemitRespNext, model.GetNextToken())
}

func TestRemittanceReportResponseJSONRoundTrip(t *testing.T) {
	model := buildExampleRemittanceReportResponse()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.RemittanceReportResponse
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasCount())
	assert.Equal(t, exampleRemitRespCount, unmarshalled.GetCount())
}

func TestRemittanceReportResponseToMap(t *testing.T) {
	model := buildExampleRemittanceReportResponse()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "nextToken") {
		assert.Equal(t, &exampleRemitRespNext, m["nextToken"])
	}
}

func TestNullableRemittanceReportResponseGetSet(t *testing.T) {
	base := buildExampleRemittanceReportResponse()
	n := openapiclient.NullableRemittanceReportResponse{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableRemittanceReportResponseUnset(t *testing.T) {
	base := buildExampleRemittanceReportResponse()
	n := openapiclient.NewNullableRemittanceReportResponse(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableRemittanceReportResponseJSONRoundTrip(t *testing.T) {
	base := buildExampleRemittanceReportResponse()
	n := openapiclient.NewNullableRemittanceReportResponse(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableRemittanceReportResponse
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleRemitRespMax, newN.Get().GetMaxResults())
	}
}
