package citypay

import (
    "encoding/json"
    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
)

var exampleAdjAmount int32 = 123
var exampleAdjIdentifier = "adj-id"
var exampleAdjReason = "Adjustment reason"

func buildExamplePaylinkAdjustmentRequest() openapiclient.PaylinkAdjustmentRequest {
    p := openapiclient.NewPaylinkAdjustmentRequest()
    p.SetAmount(exampleAdjAmount)
    p.SetIdentifier(exampleAdjIdentifier)
    p.SetReason(exampleAdjReason)
    return *p
}

func TestNewPaylinkAdjustmentRequest(t *testing.T) {
    model := openapiclient.NewPaylinkAdjustmentRequest()
    require.NotNil(t, model)
    assert.False(t, model.HasAmount())
    assert.False(t, model.HasIdentifier())
    assert.False(t, model.HasReason())
}

func TestNewPaylinkAdjustmentRequestWithDefaults(t *testing.T) {
    model := openapiclient.NewPaylinkAdjustmentRequestWithDefaults()
    require.NotNil(t, model)
    assert.False(t, model.HasAmount())
    assert.False(t, model.HasIdentifier())
    assert.False(t, model.HasReason())
}

func TestPaylinkAdjustmentRequestSetGetCycle(t *testing.T) {
    model := openapiclient.NewPaylinkAdjustmentRequest()
    model.SetAmount(exampleAdjAmount)
    assert.True(t, model.HasAmount())
    assert.Equal(t, exampleAdjAmount, model.GetAmount())

    model.SetIdentifier(exampleAdjIdentifier)
    assert.True(t, model.HasIdentifier())
    assert.Equal(t, exampleAdjIdentifier, model.GetIdentifier())

    model.SetReason(exampleAdjReason)
    assert.True(t, model.HasReason())
    assert.Equal(t, exampleAdjReason, model.GetReason())
}

func TestPaylinkAdjustmentRequestJSONRoundTrip(t *testing.T) {
    model := buildExamplePaylinkAdjustmentRequest()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.PaylinkAdjustmentRequest
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.True(t, unmarshalled.HasIdentifier())
    assert.Equal(t, exampleAdjIdentifier, unmarshalled.GetIdentifier())
}

func TestPaylinkAdjustmentRequestToMap(t *testing.T) {
    model := buildExamplePaylinkAdjustmentRequest()
    m, err := model.ToMap()
    require.NoError(t, err)
    if assert.Contains(t, m, "amount") {
        assert.Equal(t, &exampleAdjAmount, m["amount"])
    }
    if assert.Contains(t, m, "identifier") {
        assert.Equal(t, &exampleAdjIdentifier, m["identifier"])
    }
    if assert.Contains(t, m, "reason") {
        assert.Equal(t, &exampleAdjReason, m["reason"])
    }
}

func TestNullablePaylinkAdjustmentRequestGetSet(t *testing.T) {
    base := buildExamplePaylinkAdjustmentRequest()
    n := openapiclient.NullablePaylinkAdjustmentRequest{}
    n.Set(&base)
    require.True(t, n.IsSet())
    assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkAdjustmentRequestUnset(t *testing.T) {
    base := buildExamplePaylinkAdjustmentRequest()
    n := openapiclient.NewNullablePaylinkAdjustmentRequest(&base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullablePaylinkAdjustmentRequestJSONRoundTrip(t *testing.T) {
    base := buildExamplePaylinkAdjustmentRequest()
    n := openapiclient.NewNullablePaylinkAdjustmentRequest(&base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullablePaylinkAdjustmentRequest
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, exampleAdjReason, newN.Get().GetReason())
    }
}

