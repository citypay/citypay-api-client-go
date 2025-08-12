package citypay

import (
	"encoding/json"
	"testing"

	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var exampleBinLookup int32 = 123456

func buildExampleBinLookup() *openapiclient.BinLookup {
	b := openapiclient.NewBinLookup(exampleBinLookup)
	return b
}

func TestNewBinLookup(t *testing.T) {
	model := openapiclient.NewBinLookup(exampleBinLookup)
	require.NotNil(t, model)
	assert.Equal(t, exampleBinLookup, model.GetBin())
}

func TestNewBinLookupWithDefaults(t *testing.T) {
	model := openapiclient.NewBinLookupWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, int32(0), model.GetBin())
}

func TestBinLookupSetGetCycle(t *testing.T) {
	model := openapiclient.NewBinLookup(exampleBinLookup)
	model.SetBin(654321)
	assert.Equal(t, int32(654321), model.GetBin())
	val, ok := model.GetBinOk()
	require.True(t, ok)
	if assert.NotNil(t, val) {
		assert.Equal(t, int32(654321), *val)
	}
}

func TestBinLookupJSONRoundTrip(t *testing.T) {
	model := buildExampleBinLookup()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.BinLookup
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleBinLookup, unmarshalled.GetBin())
}

func TestBinLookupToMap(t *testing.T) {
	model := buildExampleBinLookup()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "bin") {
		assert.Equal(t, exampleBinLookup, m["bin"])
	}
}

func TestNullableBinLookupGetSet(t *testing.T) {
	base := buildExampleBinLookup()
	n := openapiclient.NullableBinLookup{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableBinLookupUnset(t *testing.T) {
	base := buildExampleBinLookup()
	n := openapiclient.NewNullableBinLookup(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableBinLookupJSONRoundTrip(t *testing.T) {
	base := buildExampleBinLookup()
	n := openapiclient.NewNullableBinLookup(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableBinLookup
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleBinLookup, newN.Get().GetBin())
	}
}
