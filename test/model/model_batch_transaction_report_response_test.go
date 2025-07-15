package citypay

import (
    "encoding/json"
    "testing"

    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

var exampleBTRRespCount int32 = 1
var exampleBTRRespMaxResults int32 = 10
var exampleBTRRespNextToken = "next"

var exampleAuthRefMerchant int32 = 10

func buildExampleAuthReferenceSimple() openapiclient.AuthReference {
    a := openapiclient.NewAuthReference()
    a.SetAuthcode("AC")
    a.SetMerchantid(exampleAuthRefMerchant)
    return *a
}

func buildExampleBatchTransactionReportResponse() *openapiclient.BatchTransactionReportResponse {
    data := []openapiclient.AuthReference{buildExampleAuthReferenceSimple()}
    r := openapiclient.NewBatchTransactionReportResponse(data)
    r.SetCount(exampleBTRRespCount)
    r.SetMaxResults(exampleBTRRespMaxResults)
    r.SetNextToken(exampleBTRRespNextToken)
    return r
}

func TestNewBatchTransactionReportResponse(t *testing.T) {
    data := []openapiclient.AuthReference{buildExampleAuthReferenceSimple()}
    model := openapiclient.NewBatchTransactionReportResponse(data)
    require.NotNil(t, model)
    assert.Equal(t, 1, len(model.GetData()))
    assert.False(t, model.HasCount())
}

func TestNewBatchTransactionReportResponseWithDefaults(t *testing.T) {
    model := openapiclient.NewBatchTransactionReportResponseWithDefaults()
    require.NotNil(t, model)
    assert.Equal(t, 0, len(model.GetData()))
}

func TestBatchTransactionReportResponseSetGetCycle(t *testing.T) {
    data := []openapiclient.AuthReference{buildExampleAuthReferenceSimple()}
    model := openapiclient.NewBatchTransactionReportResponse(data)
    model.SetCount(exampleBTRRespCount)
    assert.True(t, model.HasCount())
    assert.Equal(t, exampleBTRRespCount, model.GetCount())
    val, ok := model.GetCountOk()
    require.True(t, ok)
    if assert.NotNil(t, val) {
        assert.Equal(t, exampleBTRRespCount, *val)
    }
    model.SetMaxResults(exampleBTRRespMaxResults)
    assert.True(t, model.HasMaxResults())
    model.SetNextToken(exampleBTRRespNextToken)
    assert.True(t, model.HasNextToken())
}

func TestBatchTransactionReportResponseJSONRoundTrip(t *testing.T) {
    model := buildExampleBatchTransactionReportResponse()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.BatchTransactionReportResponse
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.Equal(t, 1, len(unmarshalled.GetData()))
    assert.True(t, unmarshalled.HasCount())
    assert.Equal(t, exampleBTRRespNextToken, unmarshalled.GetNextToken())
}

func TestBatchTransactionReportResponseToMap(t *testing.T) {
    model := buildExampleBatchTransactionReportResponse()
    m, err := model.ToMap()
    require.NoError(t, err)
    if assert.Contains(t, m, "data") {
        assert.Equal(t, 1, len(m["data"].([]openapiclient.AuthReference)))
    }
    if assert.Contains(t, m, "count") {
        assert.Equal(t, &exampleBTRRespCount, m["count"])
    }
    if assert.Contains(t, m, "nextToken") {
        assert.Equal(t, &exampleBTRRespNextToken, m["nextToken"])
    }
}

func TestNullableBatchTransactionReportResponseGetSet(t *testing.T) {
    base := buildExampleBatchTransactionReportResponse()
    n := openapiclient.NullableBatchTransactionReportResponse{}
    n.Set(base)
    require.True(t, n.IsSet())
    assert.Equal(t, base, n.Get())
}

func TestNullableBatchTransactionReportResponseUnset(t *testing.T) {
    base := buildExampleBatchTransactionReportResponse()
    n := openapiclient.NewNullableBatchTransactionReportResponse(base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullableBatchTransactionReportResponseJSONRoundTrip(t *testing.T) {
    base := buildExampleBatchTransactionReportResponse()
    n := openapiclient.NewNullableBatchTransactionReportResponse(base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullableBatchTransactionReportResponse
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, exampleBTRRespCount, newN.Get().GetCount())
    }
}

