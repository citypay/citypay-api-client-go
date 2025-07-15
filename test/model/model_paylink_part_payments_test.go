package citypay

import (
    "encoding/json"
    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
)

var examplePartEnabled = "true"
var examplePartFloor = "5"
var examplePartMax = "90"
var examplePartMaxRate = "10%"
var examplePartMin = "10"
var examplePartMinRate = "5%"

func buildExamplePaylinkPartPayments() openapiclient.PaylinkPartPayments {
    p := openapiclient.NewPaylinkPartPayments()
    p.SetEnabled(examplePartEnabled)
    p.SetFloor(examplePartFloor)
    p.SetMax(examplePartMax)
    p.SetMaxRate(examplePartMaxRate)
    p.SetMin(examplePartMin)
    p.SetMinRate(examplePartMinRate)
    return *p
}

func TestNewPaylinkPartPayments(t *testing.T) {
    model := openapiclient.NewPaylinkPartPayments()
    require.NotNil(t, model)
    assert.False(t, model.HasEnabled())
}

func TestNewPaylinkPartPaymentsWithDefaults(t *testing.T) {
    model := openapiclient.NewPaylinkPartPaymentsWithDefaults()
    require.NotNil(t, model)
    assert.False(t, model.HasFloor())
}

func TestPaylinkPartPaymentsSetGetCycle(t *testing.T) {
    model := openapiclient.NewPaylinkPartPayments()
    model.SetEnabled(examplePartEnabled)
    assert.True(t, model.HasEnabled())
    assert.Equal(t, examplePartEnabled, model.GetEnabled())

    model.SetFloor(examplePartFloor)
    assert.True(t, model.HasFloor())
    assert.Equal(t, examplePartFloor, model.GetFloor())

    model.SetMax(examplePartMax)
    assert.True(t, model.HasMax())
    assert.Equal(t, examplePartMax, model.GetMax())

    model.SetMaxRate(examplePartMaxRate)
    assert.True(t, model.HasMaxRate())
    assert.Equal(t, examplePartMaxRate, model.GetMaxRate())

    model.SetMin(examplePartMin)
    assert.True(t, model.HasMin())
    assert.Equal(t, examplePartMin, model.GetMin())

    model.SetMinRate(examplePartMinRate)
    assert.True(t, model.HasMinRate())
    assert.Equal(t, examplePartMinRate, model.GetMinRate())
}

func TestPaylinkPartPaymentsJSONRoundTrip(t *testing.T) {
    model := buildExamplePaylinkPartPayments()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.PaylinkPartPayments
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.True(t, unmarshalled.HasFloor())
    assert.Equal(t, examplePartFloor, unmarshalled.GetFloor())
}

func TestPaylinkPartPaymentsToMap(t *testing.T) {
    model := buildExamplePaylinkPartPayments()
    m, err := model.ToMap()
    require.NoError(t, err)
    if assert.Contains(t, m, "enabled") {
        assert.Equal(t, &examplePartEnabled, m["enabled"])
    }
    if assert.Contains(t, m, "max_rate") {
        assert.Equal(t, &examplePartMaxRate, m["max_rate"])
    }
}

func TestNullablePaylinkPartPaymentsGetSet(t *testing.T) {
    base := buildExamplePaylinkPartPayments()
    n := openapiclient.NullablePaylinkPartPayments{}
    n.Set(&base)
    require.True(t, n.IsSet())
    assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkPartPaymentsUnset(t *testing.T) {
    base := buildExamplePaylinkPartPayments()
    n := openapiclient.NewNullablePaylinkPartPayments(&base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullablePaylinkPartPaymentsJSONRoundTrip(t *testing.T) {
    base := buildExamplePaylinkPartPayments()
    n := openapiclient.NewNullablePaylinkPartPayments(&base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullablePaylinkPartPayments
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, examplePartMax, newN.Get().GetMax())
    }
}

