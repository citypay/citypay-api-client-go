package citypay

import (
    "encoding/json"
    "testing"

    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func buildExampleCheckBatchStatusResponse() *openapiclient.CheckBatchStatusResponse {
    c := openapiclient.NewCheckBatchStatusResponse()
    c.SetBatches([]openapiclient.Batch{*buildExampleBatch()})
    return c
}

func TestNewCheckBatchStatusResponse(t *testing.T) {
    model := openapiclient.NewCheckBatchStatusResponse()
    require.NotNil(t, model)
    assert.False(t, model.HasBatches())
}

func TestNewCheckBatchStatusResponseWithDefaults(t *testing.T) {
    model := openapiclient.NewCheckBatchStatusResponseWithDefaults()
    require.NotNil(t, model)
    assert.False(t, model.HasBatches())
}

func TestCheckBatchStatusResponseSetGetCycle(t *testing.T) {
    model := openapiclient.NewCheckBatchStatusResponse()
    model.SetBatches([]openapiclient.Batch{*buildExampleBatch()})
    assert.True(t, model.HasBatches())
    assert.Equal(t, 1, len(model.GetBatches()))
}

func TestCheckBatchStatusResponseJSONRoundTrip(t *testing.T) {
    model := buildExampleCheckBatchStatusResponse()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.CheckBatchStatusResponse
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.True(t, unmarshalled.HasBatches())
    assert.Equal(t, 1, len(unmarshalled.GetBatches()))
}

func TestCheckBatchStatusResponseToMap(t *testing.T) {
    model := buildExampleCheckBatchStatusResponse()
    m, err := model.ToMap()
    require.NoError(t, err)
    if assert.Contains(t, m, "batches") {
        assert.NotNil(t, m["batches"])
    }
}

func TestNullableCheckBatchStatusResponseGetSet(t *testing.T) {
    base := buildExampleCheckBatchStatusResponse()
    n := openapiclient.NullableCheckBatchStatusResponse{}
    n.Set(base)
    require.True(t, n.IsSet())
    assert.Equal(t, base, n.Get())
}

func TestNullableCheckBatchStatusResponseUnset(t *testing.T) {
    base := buildExampleCheckBatchStatusResponse()
    n := openapiclient.NewNullableCheckBatchStatusResponse(base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullableCheckBatchStatusResponseJSONRoundTrip(t *testing.T) {
    base := buildExampleCheckBatchStatusResponse()
    n := openapiclient.NewNullableCheckBatchStatusResponse(base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullableCheckBatchStatusResponse
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, 1, len(newN.Get().GetBatches()))
    }
}

