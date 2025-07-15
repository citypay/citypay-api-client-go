package citypay

import (
    "encoding/json"
    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
)

var exampleReqAccount = "acc"
var exampleReqAmount int32 = 100
var exampleReqIdentifier = "id"
var exampleReqMerchant int32 = 1
var exampleReqClientVersion = "1.0"
var exampleReqCurrency = "GBP"
var exampleReqEmail = "m@example.com"
var exampleReqRecurring = true
var exampleReqSub = "sub"
var exampleReqTxType = "SALE"

func buildExampleCardHolder() openapiclient.PaylinkCardHolder {
    c := openapiclient.NewPaylinkCardHolder()
    c.SetFirstname("f")
    return *c
}

func buildExampleCart() openapiclient.PaylinkCart {
    c := openapiclient.NewPaylinkCart()
    return *c
}

func buildExampleConfig() openapiclient.PaylinkConfig {
    c := openapiclient.NewPaylinkConfig()
    return *c
}

func buildExamplePaylinkTokenRequestModel() openapiclient.PaylinkTokenRequestModel {
    r := openapiclient.NewPaylinkTokenRequestModel(exampleReqAmount, exampleReqIdentifier, exampleReqMerchant)
    r.SetAccountno(exampleReqAccount)
    ch := buildExampleCardHolder()
    r.SetCardholder(ch)
    cart := buildExampleCart()
    r.SetCart(cart)
    r.SetClientVersion(exampleReqClientVersion)
    cfg := buildExampleConfig()
    r.SetConfig(cfg)
    r.SetCurrency(exampleReqCurrency)
    r.SetEmail(exampleReqEmail)
    r.SetRecurring(exampleReqRecurring)
    r.SetSubscriptionId(exampleReqSub)
    r.SetTxType(exampleReqTxType)
    return *r
}

func TestNewPaylinkTokenRequestModel(t *testing.T) {
    model := openapiclient.NewPaylinkTokenRequestModel(exampleReqAmount, exampleReqIdentifier, exampleReqMerchant)
    require.NotNil(t, model)
    assert.Equal(t, exampleReqAmount, model.GetAmount())
    assert.Equal(t, exampleReqIdentifier, model.GetIdentifier())
    assert.Equal(t, exampleReqMerchant, model.GetMerchantid())
    assert.False(t, model.HasEmail())
}

func TestNewPaylinkTokenRequestModelWithDefaults(t *testing.T) {
    model := openapiclient.NewPaylinkTokenRequestModelWithDefaults()
    require.NotNil(t, model)
    assert.False(t, model.HasAmount())
}

func TestPaylinkTokenRequestModelSetGetCycle(t *testing.T) {
    model := openapiclient.NewPaylinkTokenRequestModel(exampleReqAmount, exampleReqIdentifier, exampleReqMerchant)
    model.SetAccountno(exampleReqAccount)
    assert.True(t, model.HasAccountno())
    assert.Equal(t, exampleReqAccount, model.GetAccountno())

    ch := buildExampleCardHolder()
    model.SetCardholder(ch)
    assert.True(t, model.HasCardholder())

    cart := buildExampleCart()
    model.SetCart(cart)
    assert.True(t, model.HasCart())

    model.SetClientVersion(exampleReqClientVersion)
    assert.True(t, model.HasClientVersion())
    assert.Equal(t, exampleReqClientVersion, model.GetClientVersion())

    cfg := buildExampleConfig()
    model.SetConfig(cfg)
    assert.True(t, model.HasConfig())

    model.SetCurrency(exampleReqCurrency)
    assert.True(t, model.HasCurrency())
    assert.Equal(t, exampleReqCurrency, model.GetCurrency())

    model.SetEmail(exampleReqEmail)
    assert.True(t, model.HasEmail())
    assert.Equal(t, exampleReqEmail, model.GetEmail())

    model.SetRecurring(exampleReqRecurring)
    assert.True(t, model.HasRecurring())
    assert.Equal(t, exampleReqRecurring, model.GetRecurring())

    model.SetSubscriptionId(exampleReqSub)
    assert.True(t, model.HasSubscriptionId())
    assert.Equal(t, exampleReqSub, model.GetSubscriptionId())

    model.SetTxType(exampleReqTxType)
    assert.True(t, model.HasTxType())
    assert.Equal(t, exampleReqTxType, model.GetTxType())
}

func TestPaylinkTokenRequestModelJSONRoundTrip(t *testing.T) {
    model := buildExamplePaylinkTokenRequestModel()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.PaylinkTokenRequestModel
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.True(t, unmarshalled.HasAccountno())
    assert.Equal(t, exampleReqAccount, unmarshalled.GetAccountno())
}

func TestPaylinkTokenRequestModelToMap(t *testing.T) {
    model := buildExamplePaylinkTokenRequestModel()
    m, err := model.ToMap()
    require.NoError(t, err)
    assert.Equal(t, exampleReqAmount, m["amount"])
    if assert.Contains(t, m, "tx_type") {
        assert.Equal(t, &exampleReqTxType, m["tx_type"])
    }
}

func TestNullablePaylinkTokenRequestModelGetSet(t *testing.T) {
    base := buildExamplePaylinkTokenRequestModel()
    n := openapiclient.NullablePaylinkTokenRequestModel{}
    n.Set(&base)
    require.True(t, n.IsSet())
    assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkTokenRequestModelUnset(t *testing.T) {
    base := buildExamplePaylinkTokenRequestModel()
    n := openapiclient.NewNullablePaylinkTokenRequestModel(&base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullablePaylinkTokenRequestModelJSONRoundTrip(t *testing.T) {
    base := buildExamplePaylinkTokenRequestModel()
    n := openapiclient.NewNullablePaylinkTokenRequestModel(&base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullablePaylinkTokenRequestModel
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, exampleReqMerchant, newN.Get().GetMerchantid())
    }
}

