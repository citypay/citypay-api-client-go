package citypay

import (
    "encoding/json"
    "testing"

    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

var exampleBatchDate = "2023-01-01"
var exampleBatchId int32 = 42
var exampleBatchStatus = "QUEUED"

func buildExampleBatch() *openapiclient.Batch {
    b := openapiclient.NewBatch(exampleBatchDate, exampleBatchStatus)
    b.SetBatchId(exampleBatchId)
    return b
}

func TestNewBatch(t *testing.T) {
    model := openapiclient.NewBatch(exampleBatchDate, exampleBatchStatus)
    require.NotNil(t, model)
    assert.Equal(t, exampleBatchDate, model.GetBatchDate())
    assert.Equal(t, exampleBatchStatus, model.GetBatchStatus())
    assert.False(t, model.HasBatchId())
}

func TestNewBatchWithDefaults(t *testing.T) {
    model := openapiclient.NewBatchWithDefaults()
    require.NotNil(t, model)
    assert.Equal(t, "", model.GetBatchDate())
    assert.Equal(t, "", model.GetBatchStatus())
}

func TestBatchSetGetCycle(t *testing.T) {
    model := openapiclient.NewBatch(exampleBatchDate, exampleBatchStatus)
    model.SetBatchDate("2023-02-02")
    assert.Equal(t, "2023-02-02", model.GetBatchDate())

    model.SetBatchId(exampleBatchId)
    assert.True(t, model.HasBatchId())
    assert.Equal(t, exampleBatchId, model.GetBatchId())

    model.SetBatchStatus("COMPLETE")
    assert.Equal(t, "COMPLETE", model.GetBatchStatus())
}

func TestBatchJSONRoundTrip(t *testing.T) {
    model := buildExampleBatch()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.Batch
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.Equal(t, exampleBatchDate, unmarshalled.GetBatchDate())
    assert.True(t, unmarshalled.HasBatchId())
    assert.Equal(t, exampleBatchId, unmarshalled.GetBatchId())
}

func TestBatchToMap(t *testing.T) {
    model := buildExampleBatch()
    m, err := model.ToMap()
    require.NoError(t, err)
    if assert.Contains(t, m, "batch_date") {
        assert.Equal(t, exampleBatchDate, m["batch_date"])
    }
    if assert.Contains(t, m, "batch_status") {
        assert.Equal(t, exampleBatchStatus, m["batch_status"])
    }
}

func TestNullableBatchGetSet(t *testing.T) {
    base := buildExampleBatch()
    n := openapiclient.NullableBatch{}
    n.Set(base)
    require.True(t, n.IsSet())
    assert.Equal(t, base, n.Get())
}

func TestNullableBatchUnset(t *testing.T) {
    base := buildExampleBatch()
    n := openapiclient.NewNullableBatch(base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullableBatchJSONRoundTrip(t *testing.T) {
    base := buildExampleBatch()
    n := openapiclient.NewNullableBatch(base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullableBatch
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, exampleBatchStatus, newN.Get().GetBatchStatus())
    }
}

