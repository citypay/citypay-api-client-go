package citypay

import (
    "encoding/json"
    "testing"

    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

var exampleCBSBatchIds = []int32{1, 2}
var exampleCBSClientAccount = "client1"

func buildExampleCheckBatchStatus() *openapiclient.CheckBatchStatus {
    c := openapiclient.NewCheckBatchStatus(exampleCBSBatchIds)
    c.SetClientAccountId(exampleCBSClientAccount)
    return c
}

func TestNewCheckBatchStatus(t *testing.T) {
    model := openapiclient.NewCheckBatchStatus(exampleCBSBatchIds)
    require.NotNil(t, model)
    assert.Equal(t, exampleCBSBatchIds, model.GetBatchId())
    assert.False(t, model.HasClientAccountId())
}

func TestNewCheckBatchStatusWithDefaults(t *testing.T) {
    model := openapiclient.NewCheckBatchStatusWithDefaults()
    require.NotNil(t, model)
    assert.Len(t, model.GetBatchId(), 0)
}

func TestCheckBatchStatusSetGetCycle(t *testing.T) {
    model := openapiclient.NewCheckBatchStatus(exampleCBSBatchIds)
    model.SetClientAccountId(exampleCBSClientAccount)
    assert.True(t, model.HasClientAccountId())
    assert.Equal(t, exampleCBSClientAccount, model.GetClientAccountId())
    val, ok := model.GetClientAccountIdOk()
    require.True(t, ok)
    if assert.NotNil(t, val) {
        assert.Equal(t, exampleCBSClientAccount, *val)
    }
}

func TestCheckBatchStatusJSONRoundTrip(t *testing.T) {
    model := buildExampleCheckBatchStatus()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.CheckBatchStatus
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.Equal(t, exampleCBSClientAccount, unmarshalled.GetClientAccountId())
    assert.Equal(t, exampleCBSBatchIds, unmarshalled.GetBatchId())
}

func TestCheckBatchStatusToMap(t *testing.T) {
    model := buildExampleCheckBatchStatus()
    m, err := model.ToMap()
    require.NoError(t, err)
    if assert.Contains(t, m, "batch_id") {
        assert.Equal(t, exampleCBSBatchIds, m["batch_id"])
    }
    if assert.Contains(t, m, "client_account_id") {
        assert.Equal(t, &exampleCBSClientAccount, m["client_account_id"])
    }
}

func TestNullableCheckBatchStatusGetSet(t *testing.T) {
    base := buildExampleCheckBatchStatus()
    n := openapiclient.NullableCheckBatchStatus{}
    n.Set(base)
    require.True(t, n.IsSet())
    assert.Equal(t, base, n.Get())
}

func TestNullableCheckBatchStatusUnset(t *testing.T) {
    base := buildExampleCheckBatchStatus()
    n := openapiclient.NewNullableCheckBatchStatus(base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullableCheckBatchStatusJSONRoundTrip(t *testing.T) {
    base := buildExampleCheckBatchStatus()
    n := openapiclient.NewNullableCheckBatchStatus(base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullableCheckBatchStatus
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, exampleCBSClientAccount, newN.Get().GetClientAccountId())
    }
}

