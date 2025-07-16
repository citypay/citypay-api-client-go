package citypay

import (
    "encoding/json"
    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
)

var exampleItemAmount int32 = 99
var exampleItemBrand = "Brand"
var exampleItemCategory = "Cat"
var exampleItemCount int32 = 2
var exampleItemLabel = "Label"
var exampleItemMax int32 = 5
var exampleItemSku = "SKU1"
var exampleItemVariant = "Var"

func buildExamplePaylinkCartItemModel() openapiclient.PaylinkCartItemModel {
    i := openapiclient.NewPaylinkCartItemModel()
    i.SetAmount(exampleItemAmount)
    i.SetBrand(exampleItemBrand)
    i.SetCategory(exampleItemCategory)
    i.SetCount(exampleItemCount)
    i.SetLabel(exampleItemLabel)
    i.SetMax(exampleItemMax)
    i.SetSku(exampleItemSku)
    i.SetVariant(exampleItemVariant)
    return *i
}

func TestNewPaylinkCartItemModel(t *testing.T) {
    model := openapiclient.NewPaylinkCartItemModel()
    require.NotNil(t, model)
    assert.False(t, model.HasAmount())
}

func TestNewPaylinkCartItemModelWithDefaults(t *testing.T) {
    model := openapiclient.NewPaylinkCartItemModelWithDefaults()
    require.NotNil(t, model)
    assert.False(t, model.HasLabel())
}

func TestPaylinkCartItemModelSetGetCycle(t *testing.T) {
    model := openapiclient.NewPaylinkCartItemModel()
    model.SetAmount(exampleItemAmount)
    assert.True(t, model.HasAmount())
    assert.Equal(t, exampleItemAmount, model.GetAmount())

    model.SetBrand(exampleItemBrand)
    assert.True(t, model.HasBrand())
    assert.Equal(t, exampleItemBrand, model.GetBrand())

    model.SetCategory(exampleItemCategory)
    assert.True(t, model.HasCategory())
    assert.Equal(t, exampleItemCategory, model.GetCategory())

    model.SetCount(exampleItemCount)
    assert.True(t, model.HasCount())
    assert.Equal(t, exampleItemCount, model.GetCount())

    model.SetLabel(exampleItemLabel)
    assert.True(t, model.HasLabel())
    assert.Equal(t, exampleItemLabel, model.GetLabel())

    model.SetMax(exampleItemMax)
    assert.True(t, model.HasMax())
    assert.Equal(t, exampleItemMax, model.GetMax())

    model.SetSku(exampleItemSku)
    assert.True(t, model.HasSku())
    assert.Equal(t, exampleItemSku, model.GetSku())

    model.SetVariant(exampleItemVariant)
    assert.True(t, model.HasVariant())
    assert.Equal(t, exampleItemVariant, model.GetVariant())
}

func TestPaylinkCartItemModelJSONRoundTrip(t *testing.T) {
    model := buildExamplePaylinkCartItemModel()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.PaylinkCartItemModel
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.True(t, unmarshalled.HasBrand())
    assert.Equal(t, exampleItemBrand, unmarshalled.GetBrand())
}

func TestPaylinkCartItemModelToMap(t *testing.T) {
    model := buildExamplePaylinkCartItemModel()
    m, err := model.ToMap()
    require.NoError(t, err)
    if assert.Contains(t, m, "sku") {
        assert.Equal(t, &exampleItemSku, m["sku"])
    }
    if assert.Contains(t, m, "variant") {
        assert.Equal(t, &exampleItemVariant, m["variant"])
    }
}

func TestNullablePaylinkCartItemModelGetSet(t *testing.T) {
    base := buildExamplePaylinkCartItemModel()
    n := openapiclient.NullablePaylinkCartItemModel{}
    n.Set(&base)
    require.True(t, n.IsSet())
    assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkCartItemModelUnset(t *testing.T) {
    base := buildExamplePaylinkCartItemModel()
    n := openapiclient.NewNullablePaylinkCartItemModel(&base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullablePaylinkCartItemModelJSONRoundTrip(t *testing.T) {
    base := buildExamplePaylinkCartItemModel()
    n := openapiclient.NewNullablePaylinkCartItemModel(&base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullablePaylinkCartItemModel
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, exampleItemAmount, newN.Get().GetAmount())
    }
}

