package citypay

import (
    "encoding/json"
    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
)

var exampleParamEntry = "pre"
var exampleParamField = "text"
var exampleParamGroup = "g"
var exampleParamLabel = "lbl"
var exampleParamLocked = true
var exampleParamName = "p"
var exampleParamOrder int32 = 1
var exampleParamPattern = "[0-9]"
var exampleParamPlaceholder = "ph"
var exampleParamRequired = true
var exampleParamValue = "val"

func buildExamplePaylinkCustomParam() openapiclient.PaylinkCustomParam {
    p := openapiclient.NewPaylinkCustomParam(exampleParamName)
    p.SetEntryMode(exampleParamEntry)
    p.SetFieldType(exampleParamField)
    p.SetGroup(exampleParamGroup)
    p.SetLabel(exampleParamLabel)
    p.SetLocked(exampleParamLocked)
    p.SetOrder(exampleParamOrder)
    p.SetPattern(exampleParamPattern)
    p.SetPlaceholder(exampleParamPlaceholder)
    p.SetRequired(exampleParamRequired)
    p.SetValue(exampleParamValue)
    return *p
}

func TestNewPaylinkCustomParam(t *testing.T) {
    model := openapiclient.NewPaylinkCustomParam(exampleParamName)
    require.NotNil(t, model)
    assert.Equal(t, exampleParamName, model.GetName())
    assert.False(t, model.HasLocked())
}

func TestNewPaylinkCustomParamWithDefaults(t *testing.T) {
    model := openapiclient.NewPaylinkCustomParamWithDefaults()
    require.NotNil(t, model)
    assert.False(t, model.HasEntryMode())
}

func TestPaylinkCustomParamSetGetCycle(t *testing.T) {
    model := openapiclient.NewPaylinkCustomParam(exampleParamName)
    model.SetEntryMode(exampleParamEntry)
    assert.True(t, model.HasEntryMode())
    assert.Equal(t, exampleParamEntry, model.GetEntryMode())

    model.SetFieldType(exampleParamField)
    assert.True(t, model.HasFieldType())
    assert.Equal(t, exampleParamField, model.GetFieldType())

    model.SetGroup(exampleParamGroup)
    assert.True(t, model.HasGroup())
    assert.Equal(t, exampleParamGroup, model.GetGroup())

    model.SetLabel(exampleParamLabel)
    assert.True(t, model.HasLabel())
    assert.Equal(t, exampleParamLabel, model.GetLabel())

    model.SetLocked(exampleParamLocked)
    assert.True(t, model.HasLocked())
    assert.Equal(t, exampleParamLocked, model.GetLocked())

    model.SetOrder(exampleParamOrder)
    assert.True(t, model.HasOrder())
    assert.Equal(t, exampleParamOrder, model.GetOrder())

    model.SetPattern(exampleParamPattern)
    assert.True(t, model.HasPattern())
    assert.Equal(t, exampleParamPattern, model.GetPattern())

    model.SetPlaceholder(exampleParamPlaceholder)
    assert.True(t, model.HasPlaceholder())
    assert.Equal(t, exampleParamPlaceholder, model.GetPlaceholder())

    model.SetRequired(exampleParamRequired)
    assert.True(t, model.HasRequired())
    assert.Equal(t, exampleParamRequired, model.GetRequired())

    model.SetValue(exampleParamValue)
    assert.True(t, model.HasValue())
    assert.Equal(t, exampleParamValue, model.GetValue())
}

func TestPaylinkCustomParamJSONRoundTrip(t *testing.T) {
    model := buildExamplePaylinkCustomParam()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.PaylinkCustomParam
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.True(t, unmarshalled.HasOrder())
    assert.Equal(t, exampleParamOrder, unmarshalled.GetOrder())
}

func TestPaylinkCustomParamToMap(t *testing.T) {
    model := buildExamplePaylinkCustomParam()
    m, err := model.ToMap()
    require.NoError(t, err)
    assert.Equal(t, exampleParamName, m["name"])
    if assert.Contains(t, m, "value") {
        assert.Equal(t, &exampleParamValue, m["value"])
    }
}

func TestNullablePaylinkCustomParamGetSet(t *testing.T) {
    base := buildExamplePaylinkCustomParam()
    n := openapiclient.NullablePaylinkCustomParam{}
    n.Set(&base)
    require.True(t, n.IsSet())
    assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkCustomParamUnset(t *testing.T) {
    base := buildExamplePaylinkCustomParam()
    n := openapiclient.NewNullablePaylinkCustomParam(&base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullablePaylinkCustomParamJSONRoundTrip(t *testing.T) {
    base := buildExamplePaylinkCustomParam()
    n := openapiclient.NewNullablePaylinkCustomParam(&base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullablePaylinkCustomParam
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, exampleParamName, newN.Get().GetName())
    }
}

