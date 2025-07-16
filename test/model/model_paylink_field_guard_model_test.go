package citypay

import (
    "encoding/json"
    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
)

var exampleFGField = "text"
var exampleFGLabel = "Security"
var exampleFGMax int32 = 8
var exampleFGMin int32 = 2
var exampleFGName = "pin"
var exampleFGRegex = "\\d+"
var exampleFGValue = "12"

func buildExamplePaylinkFieldGuardModel() openapiclient.PaylinkFieldGuardModel {
    f := openapiclient.NewPaylinkFieldGuardModel()
    f.SetFieldType(exampleFGField)
    f.SetLabel(exampleFGLabel)
    f.SetMaxlen(exampleFGMax)
    f.SetMinlen(exampleFGMin)
    f.SetName(exampleFGName)
    f.SetRegex(exampleFGRegex)
    f.SetValue(exampleFGValue)
    return *f
}

func TestNewPaylinkFieldGuardModel(t *testing.T) {
    model := openapiclient.NewPaylinkFieldGuardModel()
    require.NotNil(t, model)
    assert.False(t, model.HasFieldType())
}

func TestNewPaylinkFieldGuardModelWithDefaults(t *testing.T) {
    model := openapiclient.NewPaylinkFieldGuardModelWithDefaults()
    require.NotNil(t, model)
    assert.False(t, model.HasName())
}

func TestPaylinkFieldGuardModelSetGetCycle(t *testing.T) {
    model := openapiclient.NewPaylinkFieldGuardModel()
    model.SetFieldType(exampleFGField)
    assert.True(t, model.HasFieldType())
    assert.Equal(t, exampleFGField, model.GetFieldType())

    model.SetLabel(exampleFGLabel)
    assert.True(t, model.HasLabel())
    assert.Equal(t, exampleFGLabel, model.GetLabel())

    model.SetMaxlen(exampleFGMax)
    assert.True(t, model.HasMaxlen())
    assert.Equal(t, exampleFGMax, model.GetMaxlen())

    model.SetMinlen(exampleFGMin)
    assert.True(t, model.HasMinlen())
    assert.Equal(t, exampleFGMin, model.GetMinlen())

    model.SetName(exampleFGName)
    assert.True(t, model.HasName())
    assert.Equal(t, exampleFGName, model.GetName())

    model.SetRegex(exampleFGRegex)
    assert.True(t, model.HasRegex())
    assert.Equal(t, exampleFGRegex, model.GetRegex())

    model.SetValue(exampleFGValue)
    assert.True(t, model.HasValue())
    assert.Equal(t, exampleFGValue, model.GetValue())
}

func TestPaylinkFieldGuardModelJSONRoundTrip(t *testing.T) {
    model := buildExamplePaylinkFieldGuardModel()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.PaylinkFieldGuardModel
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.True(t, unmarshalled.HasName())
    assert.Equal(t, exampleFGName, unmarshalled.GetName())
}

func TestPaylinkFieldGuardModelToMap(t *testing.T) {
    model := buildExamplePaylinkFieldGuardModel()
    m, err := model.ToMap()
    require.NoError(t, err)
    if assert.Contains(t, m, "regex") {
        assert.Equal(t, &exampleFGRegex, m["regex"])
    }
}

func TestNullablePaylinkFieldGuardModelGetSet(t *testing.T) {
    base := buildExamplePaylinkFieldGuardModel()
    n := openapiclient.NullablePaylinkFieldGuardModel{}
    n.Set(&base)
    require.True(t, n.IsSet())
    assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkFieldGuardModelUnset(t *testing.T) {
    base := buildExamplePaylinkFieldGuardModel()
    n := openapiclient.NewNullablePaylinkFieldGuardModel(&base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullablePaylinkFieldGuardModelJSONRoundTrip(t *testing.T) {
    base := buildExamplePaylinkFieldGuardModel()
    n := openapiclient.NewNullablePaylinkFieldGuardModel(&base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullablePaylinkFieldGuardModel
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, exampleFGMin, newN.Get().GetMinlen())
    }
}

