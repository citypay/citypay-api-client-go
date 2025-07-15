package citypay

import (
    "encoding/json"
    "testing"

    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

var exampleBTAccountId = "acc1"
var exampleBTAmount int32 = 500
var exampleBTIdentifier = "id1"
var exampleBTMerchantId int32 = 42

func buildExampleBatchTransaction() *openapiclient.BatchTransaction {
    b := openapiclient.NewBatchTransaction(exampleBTAccountId, exampleBTAmount)
    b.SetIdentifier(exampleBTIdentifier)
    b.SetMerchantid(exampleBTMerchantId)
    return b
}

func TestNewBatchTransaction(t *testing.T) {
    model := openapiclient.NewBatchTransaction(exampleBTAccountId, exampleBTAmount)
    require.NotNil(t, model)
    assert.Equal(t, exampleBTAccountId, model.GetAccountId())
    assert.Equal(t, exampleBTAmount, model.GetAmount())
    assert.False(t, model.HasIdentifier())
    assert.False(t, model.HasMerchantid())
}

func TestNewBatchTransactionWithDefaults(t *testing.T) {
    model := openapiclient.NewBatchTransactionWithDefaults()
    require.NotNil(t, model)
    assert.Equal(t, "", model.GetAccountId())
    assert.Equal(t, int32(0), model.GetAmount())
}

func TestBatchTransactionSetGetCycle(t *testing.T) {
    model := openapiclient.NewBatchTransaction(exampleBTAccountId, exampleBTAmount)
    model.SetIdentifier(exampleBTIdentifier)
    assert.True(t, model.HasIdentifier())
    assert.Equal(t, exampleBTIdentifier, model.GetIdentifier())
    val, ok := model.GetIdentifierOk()
    require.True(t, ok)
    if assert.NotNil(t, val) {
        assert.Equal(t, exampleBTIdentifier, *val)
    }
    model.SetMerchantid(exampleBTMerchantId)
    assert.True(t, model.HasMerchantid())
    assert.Equal(t, exampleBTMerchantId, model.GetMerchantid())
    val2, ok2 := model.GetMerchantidOk()
    require.True(t, ok2)
    if assert.NotNil(t, val2) {
        assert.Equal(t, exampleBTMerchantId, *val2)
    }
}

func TestBatchTransactionJSONRoundTrip(t *testing.T) {
    model := buildExampleBatchTransaction()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.BatchTransaction
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.Equal(t, exampleBTAccountId, unmarshalled.GetAccountId())
    assert.True(t, unmarshalled.HasIdentifier())
    assert.Equal(t, exampleBTMerchantId, unmarshalled.GetMerchantid())
}

func TestBatchTransactionToMap(t *testing.T) {
    model := buildExampleBatchTransaction()
    m, err := model.ToMap()
    require.NoError(t, err)
    if assert.Contains(t, m, "account_id") {
        assert.Equal(t, exampleBTAccountId, m["account_id"])
    }
    if assert.Contains(t, m, "amount") {
        assert.Equal(t, exampleBTAmount, m["amount"])
    }
    if assert.Contains(t, m, "identifier") {
        assert.Equal(t, &exampleBTIdentifier, m["identifier"])
    }
}

func TestNullableBatchTransactionGetSet(t *testing.T) {
    base := buildExampleBatchTransaction()
    n := openapiclient.NullableBatchTransaction{}
    n.Set(base)
    require.True(t, n.IsSet())
    assert.Equal(t, base, n.Get())
}

func TestNullableBatchTransactionUnset(t *testing.T) {
    base := buildExampleBatchTransaction()
    n := openapiclient.NewNullableBatchTransaction(base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullableBatchTransactionJSONRoundTrip(t *testing.T) {
    base := buildExampleBatchTransaction()
    n := openapiclient.NewNullableBatchTransaction(base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullableBatchTransaction
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, exampleBTIdentifier, newN.Get().GetIdentifier())
    }
}

