package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleTokAuthResult = "Y"
var exampleTokBinCommercial = true
var exampleTokBinDebit = true
var exampleTokDescription = "desc"
var exampleTokEci = "05"
var exampleTokIdentifier = "id"
var exampleTokMasked = "411111******1111"
var exampleTokScheme = "VISA"
var exampleTokSig = "sig"
var exampleTokToken = "tok"

func buildExampleTokenisationResponseModel() openapiclient.TokenisationResponseModel {
	t := openapiclient.NewTokenisationResponseModel()
	t.SetAuthenResult(exampleTokAuthResult)
	t.SetBinCommercial(exampleTokBinCommercial)
	t.SetBinDebit(exampleTokBinDebit)
	t.SetBinDescription(exampleTokDescription)
	t.SetEci(exampleTokEci)
	t.SetIdentifier(exampleTokIdentifier)
	t.SetMaskedpan(exampleTokMasked)
	t.SetScheme(exampleTokScheme)
	t.SetSigId(exampleTokSig)
	t.SetToken(exampleTokToken)
	return *t
}

func TestNewTokenisationResponseModel(t *testing.T) {
	model := openapiclient.NewTokenisationResponseModel()
	require.NotNil(t, model)
	assert.False(t, model.HasScheme())
}

func TestNewTokenisationResponseModelWithDefaults(t *testing.T) {
	model := openapiclient.NewTokenisationResponseModelWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasAuthenResult())
}

func TestTokenisationResponseModelSetGetCycle(t *testing.T) {
	model := openapiclient.NewTokenisationResponseModel()
	model.SetAuthenResult(exampleTokAuthResult)
	assert.True(t, model.HasAuthenResult())
	assert.Equal(t, exampleTokAuthResult, model.GetAuthenResult())

	model.SetBinCommercial(exampleTokBinCommercial)
	assert.True(t, model.HasBinCommercial())
	assert.Equal(t, exampleTokBinCommercial, model.GetBinCommercial())

	model.SetBinDebit(exampleTokBinDebit)
	assert.True(t, model.HasBinDebit())
	assert.Equal(t, exampleTokBinDebit, model.GetBinDebit())

	model.SetBinDescription(exampleTokDescription)
	assert.True(t, model.HasBinDescription())
	assert.Equal(t, exampleTokDescription, model.GetBinDescription())

	model.SetEci(exampleTokEci)
	assert.True(t, model.HasEci())
	assert.Equal(t, exampleTokEci, model.GetEci())

	model.SetIdentifier(exampleTokIdentifier)
	assert.True(t, model.HasIdentifier())
	assert.Equal(t, exampleTokIdentifier, model.GetIdentifier())

	model.SetMaskedpan(exampleTokMasked)
	assert.True(t, model.HasMaskedpan())
	assert.Equal(t, exampleTokMasked, model.GetMaskedpan())

	model.SetScheme(exampleTokScheme)
	assert.True(t, model.HasScheme())
	assert.Equal(t, exampleTokScheme, model.GetScheme())

	model.SetSigId(exampleTokSig)
	assert.True(t, model.HasSigId())
	assert.Equal(t, exampleTokSig, model.GetSigId())

	model.SetToken(exampleTokToken)
	assert.True(t, model.HasToken())
	assert.Equal(t, exampleTokToken, model.GetToken())
}

func TestTokenisationResponseModelJSONRoundTrip(t *testing.T) {
	model := buildExampleTokenisationResponseModel()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.TokenisationResponseModel
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasMaskedpan())
	assert.Equal(t, exampleTokMasked, unmarshalled.GetMaskedpan())
}

func TestTokenisationResponseModelToMap(t *testing.T) {
	model := buildExampleTokenisationResponseModel()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "token") {
		assert.Equal(t, &exampleTokToken, m["token"])
	}
}

func TestNullableTokenisationResponseModelGetSet(t *testing.T) {
	base := buildExampleTokenisationResponseModel()
	n := openapiclient.NullableTokenisationResponseModel{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableTokenisationResponseModelUnset(t *testing.T) {
	base := buildExampleTokenisationResponseModel()
	n := openapiclient.NewNullableTokenisationResponseModel(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableTokenisationResponseModelJSONRoundTrip(t *testing.T) {
	base := buildExampleTokenisationResponseModel()
	n := openapiclient.NewNullableTokenisationResponseModel(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableTokenisationResponseModel
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleTokEci, newN.Get().GetEci())
	}
}
