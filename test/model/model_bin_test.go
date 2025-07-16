package citypay

import (
    "encoding/json"
    "testing"

    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

var exampleBinCommercial = true
var exampleBinCorporate = true
var exampleBinCountry = "GB"
var exampleBinCredit = true
var exampleBinCurrency = "GBP"
var exampleBinDebit = true
var exampleBinDescription = "Business"
var exampleBinEu = true
var exampleBinScheme = "VISA"

func buildExampleBin() *openapiclient.Bin {
    b := openapiclient.NewBin()
    b.SetBinCommercial(exampleBinCommercial)
    b.SetBinCorporate(exampleBinCorporate)
    b.SetBinCountryIssued(exampleBinCountry)
    b.SetBinCredit(exampleBinCredit)
    b.SetBinCurrency(exampleBinCurrency)
    b.SetBinDebit(exampleBinDebit)
    b.SetBinDescription(exampleBinDescription)
    b.SetBinEu(exampleBinEu)
    b.SetScheme(exampleBinScheme)
    return b
}

func TestNewBin(t *testing.T) {
    model := openapiclient.NewBin()
    require.NotNil(t, model)
    assert.False(t, model.HasBinCommercial())
    assert.False(t, model.HasScheme())
}

func TestNewBinWithDefaults(t *testing.T) {
    model := openapiclient.NewBinWithDefaults()
    require.NotNil(t, model)
    assert.False(t, model.HasBinDebit())
}

func TestBinSetGetCycle(t *testing.T) {
    model := openapiclient.NewBin()
    model.SetBinCommercial(exampleBinCommercial)
    assert.True(t, model.HasBinCommercial())
    assert.Equal(t, exampleBinCommercial, model.GetBinCommercial())
    if val, ok := model.GetBinCommercialOk(); assert.True(t, ok) {
        assert.NotNil(t, val)
        assert.Equal(t, exampleBinCommercial, *val)
    }

    model.SetBinCorporate(exampleBinCorporate)
    assert.True(t, model.HasBinCorporate())
    assert.Equal(t, exampleBinCorporate, model.GetBinCorporate())

    model.SetBinCountryIssued(exampleBinCountry)
    assert.True(t, model.HasBinCountryIssued())
    assert.Equal(t, exampleBinCountry, model.GetBinCountryIssued())

    model.SetBinCredit(exampleBinCredit)
    assert.True(t, model.HasBinCredit())
    assert.Equal(t, exampleBinCredit, model.GetBinCredit())

    model.SetBinCurrency(exampleBinCurrency)
    assert.True(t, model.HasBinCurrency())
    assert.Equal(t, exampleBinCurrency, model.GetBinCurrency())

    model.SetBinDebit(exampleBinDebit)
    assert.True(t, model.HasBinDebit())
    assert.Equal(t, exampleBinDebit, model.GetBinDebit())

    model.SetBinDescription(exampleBinDescription)
    assert.True(t, model.HasBinDescription())
    assert.Equal(t, exampleBinDescription, model.GetBinDescription())

    model.SetBinEu(exampleBinEu)
    assert.True(t, model.HasBinEu())
    assert.Equal(t, exampleBinEu, model.GetBinEu())

    model.SetScheme(exampleBinScheme)
    assert.True(t, model.HasScheme())
    assert.Equal(t, exampleBinScheme, model.GetScheme())
}

func TestBinJSONRoundTrip(t *testing.T) {
    model := buildExampleBin()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.Bin
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.Equal(t, exampleBinCommercial, unmarshalled.GetBinCommercial())
    assert.Equal(t, exampleBinScheme, unmarshalled.GetScheme())
}

func TestBinToMap(t *testing.T) {
    model := buildExampleBin()
    m, err := model.ToMap()
    require.NoError(t, err)
    if assert.Contains(t, m, "bin_commercial") {
        assert.Equal(t, &exampleBinCommercial, m["bin_commercial"])
    }
    if assert.Contains(t, m, "scheme") {
        assert.Equal(t, &exampleBinScheme, m["scheme"])
    }
}

func TestNullableBinGetSet(t *testing.T) {
    base := buildExampleBin()
    n := openapiclient.NullableBin{}
    n.Set(base)
    require.True(t, n.IsSet())
    assert.Equal(t, base, n.Get())
}

func TestNullableBinUnset(t *testing.T) {
    base := buildExampleBin()
    n := openapiclient.NewNullableBin(base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullableBinJSONRoundTrip(t *testing.T) {
    base := buildExampleBin()
    n := openapiclient.NewNullableBin(base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullableBin
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, exampleBinScheme, newN.Get().GetScheme())
    }
}

