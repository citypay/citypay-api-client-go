package citypay

import (
    "encoding/json"
    "testing"
    "time"

    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

var exampleRespAmount int32 = 1000
var exampleRespAtrn = "ATRN"
var exampleRespAtsd = "A1"
var exampleRespAuthcode = "ACODE"
var exampleRespAuthenResult = "Y"
var exampleRespAuthorised = true
var exampleRespAvsResult = "Y"
var exampleRespBinCommercial = true
var exampleRespBinDebit = true
var exampleRespBinDescription = "DEBIT"
var exampleRespCavv = "CAVV"
var exampleRespContext = "CTX"
var exampleRespCscResult = "M"
var exampleRespCurrency = "GBP"
var exampleRespDatetime = time.Date(2023, time.March, 5, 10, 0, 0, 0, time.UTC)
var exampleRespEci = "05"
var exampleRespIdentifier = "ID1"
var exampleRespLive = true
var exampleRespMaskedpan = "411111******1111"
var exampleRespMerchantid int32 = 321
var exampleRespResult int32 = 1
var exampleRespResultCode = "A"
var exampleRespResultMessage = "Approved"
var exampleRespScheme = "VISA"
var exampleRespSchemeId = "VIS"
var exampleRespSchemeLogo = "logo"
var exampleRespSha256 = "hash"
var exampleRespTransStatus = "APPROVED"
var exampleRespTransno int32 = 55

func buildExampleAuthResponse() *openapiclient.AuthResponse {
    r := openapiclient.NewAuthResponse(exampleRespMerchantid, exampleRespResult, exampleRespResultCode, exampleRespResultMessage)
    r.SetAmount(exampleRespAmount)
    r.SetAtrn(exampleRespAtrn)
    r.SetAtsd(exampleRespAtsd)
    r.SetAuthcode(exampleRespAuthcode)
    r.SetAuthenResult(exampleRespAuthenResult)
    r.SetAuthorised(exampleRespAuthorised)
    r.SetAvsResult(exampleRespAvsResult)
    r.SetBinCommercial(exampleRespBinCommercial)
    r.SetBinDebit(exampleRespBinDebit)
    r.SetBinDescription(exampleRespBinDescription)
    r.SetCavv(exampleRespCavv)
    r.SetContext(exampleRespContext)
    r.SetCscResult(exampleRespCscResult)
    r.SetCurrency(exampleRespCurrency)
    r.SetDatetime(exampleRespDatetime)
    r.SetEci(exampleRespEci)
    r.SetIdentifier(exampleRespIdentifier)
    r.SetLive(exampleRespLive)
    r.SetMaskedpan(exampleRespMaskedpan)
    r.SetScheme(exampleRespScheme)
    r.SetSchemeId(exampleRespSchemeId)
    r.SetSchemeLogo(exampleRespSchemeLogo)
    r.SetSha256(exampleRespSha256)
    r.SetTransStatus(exampleRespTransStatus)
    r.SetTransno(exampleRespTransno)
    return r
}

func TestNewAuthResponse(t *testing.T) {
    model := openapiclient.NewAuthResponse(exampleRespMerchantid, exampleRespResult, exampleRespResultCode, exampleRespResultMessage)
    require.NotNil(t, model)
    assert.Equal(t, exampleRespMerchantid, model.GetMerchantid())
    assert.Equal(t, exampleRespResult, model.GetResult())
    assert.False(t, model.HasAmount())
}

func TestNewAuthResponseWithDefaults(t *testing.T) {
    model := openapiclient.NewAuthResponseWithDefaults()
    require.NotNil(t, model)
    assert.Equal(t, int32(0), model.GetMerchantid())
    assert.False(t, model.HasAmount())
}

func TestAuthResponseSetGetCycle(t *testing.T) {
    model := openapiclient.NewAuthResponse(exampleRespMerchantid, exampleRespResult, exampleRespResultCode, exampleRespResultMessage)
    model.SetAmount(exampleRespAmount)
    assert.True(t, model.HasAmount())
    assert.Equal(t, exampleRespAmount, model.GetAmount())

    model.SetAtrn(exampleRespAtrn)
    assert.True(t, model.HasAtrn())
    assert.Equal(t, exampleRespAtrn, model.GetAtrn())

    model.SetAtsd(exampleRespAtsd)
    assert.True(t, model.HasAtsd())
    assert.Equal(t, exampleRespAtsd, model.GetAtsd())

    model.SetAuthcode(exampleRespAuthcode)
    assert.True(t, model.HasAuthcode())
    assert.Equal(t, exampleRespAuthcode, model.GetAuthcode())

    model.SetAuthenResult(exampleRespAuthenResult)
    assert.True(t, model.HasAuthenResult())
    assert.Equal(t, exampleRespAuthenResult, model.GetAuthenResult())

    model.SetAuthorised(exampleRespAuthorised)
    assert.True(t, model.HasAuthorised())
    assert.Equal(t, exampleRespAuthorised, model.GetAuthorised())

    model.SetAvsResult(exampleRespAvsResult)
    assert.True(t, model.HasAvsResult())
    assert.Equal(t, exampleRespAvsResult, model.GetAvsResult())

    model.SetBinCommercial(exampleRespBinCommercial)
    assert.True(t, model.HasBinCommercial())
    assert.Equal(t, exampleRespBinCommercial, model.GetBinCommercial())

    model.SetBinDebit(exampleRespBinDebit)
    assert.True(t, model.HasBinDebit())
    assert.Equal(t, exampleRespBinDebit, model.GetBinDebit())

    model.SetBinDescription(exampleRespBinDescription)
    assert.True(t, model.HasBinDescription())
    assert.Equal(t, exampleRespBinDescription, model.GetBinDescription())

    model.SetCavv(exampleRespCavv)
    assert.True(t, model.HasCavv())
    assert.Equal(t, exampleRespCavv, model.GetCavv())

    model.SetContext(exampleRespContext)
    assert.True(t, model.HasContext())
    assert.Equal(t, exampleRespContext, model.GetContext())

    model.SetCscResult(exampleRespCscResult)
    assert.True(t, model.HasCscResult())
    assert.Equal(t, exampleRespCscResult, model.GetCscResult())

    model.SetCurrency(exampleRespCurrency)
    assert.True(t, model.HasCurrency())
    assert.Equal(t, exampleRespCurrency, model.GetCurrency())

    model.SetDatetime(exampleRespDatetime)
    assert.True(t, model.HasDatetime())
    assert.Equal(t, exampleRespDatetime, model.GetDatetime())

    model.SetEci(exampleRespEci)
    assert.True(t, model.HasEci())
    assert.Equal(t, exampleRespEci, model.GetEci())

    model.SetIdentifier(exampleRespIdentifier)
    assert.True(t, model.HasIdentifier())
    assert.Equal(t, exampleRespIdentifier, model.GetIdentifier())

    model.SetLive(exampleRespLive)
    assert.True(t, model.HasLive())
    assert.Equal(t, exampleRespLive, model.GetLive())

    model.SetMaskedpan(exampleRespMaskedpan)
    assert.True(t, model.HasMaskedpan())
    assert.Equal(t, exampleRespMaskedpan, model.GetMaskedpan())

    model.SetScheme(exampleRespScheme)
    assert.True(t, model.HasScheme())
    assert.Equal(t, exampleRespScheme, model.GetScheme())

    model.SetSchemeId(exampleRespSchemeId)
    assert.True(t, model.HasSchemeId())
    assert.Equal(t, exampleRespSchemeId, model.GetSchemeId())

    model.SetSchemeLogo(exampleRespSchemeLogo)
    assert.True(t, model.HasSchemeLogo())
    assert.Equal(t, exampleRespSchemeLogo, model.GetSchemeLogo())

    model.SetSha256(exampleRespSha256)
    assert.True(t, model.HasSha256())
    assert.Equal(t, exampleRespSha256, model.GetSha256())

    model.SetTransStatus(exampleRespTransStatus)
    assert.True(t, model.HasTransStatus())
    assert.Equal(t, exampleRespTransStatus, model.GetTransStatus())

    model.SetTransno(exampleRespTransno)
    assert.True(t, model.HasTransno())
    assert.Equal(t, exampleRespTransno, model.GetTransno())
}

func TestAuthResponseJSONRoundTrip(t *testing.T) {
    model := buildExampleAuthResponse()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.AuthResponse
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.Equal(t, exampleRespResultCode, unmarshalled.GetResultCode())
    assert.True(t, unmarshalled.HasAmount())
}

func TestAuthResponseToMap(t *testing.T) {
    model := buildExampleAuthResponse()
    m, err := model.ToMap()
    require.NoError(t, err)
    if assert.Contains(t, m, "amount") {
        assert.Equal(t, &exampleRespAmount, m["amount"])
    }
    if assert.Contains(t, m, "merchantid") {
        assert.Equal(t, exampleRespMerchantid, m["merchantid"])
    }
}

func TestNullableAuthResponseGetSet(t *testing.T) {
    base := buildExampleAuthResponse()
    n := openapiclient.NullableAuthResponse{}
    n.Set(base)
    require.True(t, n.IsSet())
    assert.Equal(t, base, n.Get())
}

func TestNullableAuthResponseUnset(t *testing.T) {
    base := buildExampleAuthResponse()
    n := openapiclient.NewNullableAuthResponse(base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullableAuthResponseJSONRoundTrip(t *testing.T) {
    base := buildExampleAuthResponse()
    n := openapiclient.NewNullableAuthResponse(base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullableAuthResponse
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, exampleRespResultMessage, newN.Get().GetResultMessage())
    }
}

