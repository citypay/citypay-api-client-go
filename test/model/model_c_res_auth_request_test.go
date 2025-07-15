package citypay

import (
    "encoding/json"
    "testing"

    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

var exampleCRes = "cresdata"

func buildExampleCResAuthRequest() *openapiclient.CResAuthRequest {
    c := openapiclient.NewCResAuthRequest()
    c.SetCres(exampleCRes)
    return c
}

func TestNewCResAuthRequest(t *testing.T) {
    model := openapiclient.NewCResAuthRequest()
    require.NotNil(t, model)
    assert.False(t, model.HasCres())
}

func TestNewCResAuthRequestWithDefaults(t *testing.T) {
    model := openapiclient.NewCResAuthRequestWithDefaults()
    require.NotNil(t, model)
    assert.False(t, model.HasCres())
}

func TestCResAuthRequestSetGetCycle(t *testing.T) {
    model := openapiclient.NewCResAuthRequest()
    model.SetCres(exampleCRes)
    assert.True(t, model.HasCres())
    assert.Equal(t, exampleCRes, model.GetCres())
    val, ok := model.GetCresOk()
    require.True(t, ok)
    if assert.NotNil(t, val) {
        assert.Equal(t, exampleCRes, *val)
    }
}

func TestCResAuthRequestJSONRoundTrip(t *testing.T) {
    model := buildExampleCResAuthRequest()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.CResAuthRequest
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.Equal(t, exampleCRes, unmarshalled.GetCres())
}

func TestCResAuthRequestToMap(t *testing.T) {
    model := buildExampleCResAuthRequest()
    m, err := model.ToMap()
    require.NoError(t, err)
    if assert.Contains(t, m, "cres") {
        assert.Equal(t, &exampleCRes, m["cres"])
    }
}

func TestNullableCResAuthRequestGetSet(t *testing.T) {
    base := buildExampleCResAuthRequest()
    n := openapiclient.NullableCResAuthRequest{}
    n.Set(base)
    require.True(t, n.IsSet())
    assert.Equal(t, base, n.Get())
}

func TestNullableCResAuthRequestUnset(t *testing.T) {
    base := buildExampleCResAuthRequest()
    n := openapiclient.NewNullableCResAuthRequest(base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullableCResAuthRequestJSONRoundTrip(t *testing.T) {
    base := buildExampleCResAuthRequest()
    n := openapiclient.NewNullableCResAuthRequest(base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullableCResAuthRequest
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, exampleCRes, newN.Get().GetCres())
    }
}

