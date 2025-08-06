package citypay

import (
	"encoding/json"
	"testing"
	"time"

	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var exampleTokenId = "id"
var exampleTokenResult int32 = 1
var exampleTokenToken = "tok"
var exampleTokenBps = "Y"
var exampleTokenDate = time.Date(2024, time.July, 1, 0, 0, 0, 0, time.UTC)
var exampleTokenIdentifier = "ident"
var exampleTokenMode = "test"
var exampleTokenQrcode = "https://q"
var exampleTokenServer = "v1"
var exampleTokenSource = "1.1.1.1"
var exampleTokenURL = "https://u"
var exampleTokenUsc = "usc"

func buildExampleAttachmentRes() []openapiclient.PaylinkAttachmentResult {
	a := openapiclient.NewPaylinkAttachmentResult("name", "OK")
	return []openapiclient.PaylinkAttachmentResult{*a}
}

func buildExampleErrorCode() openapiclient.PaylinkErrorCode {
	e := openapiclient.NewPaylinkErrorCode("E", "m")
	return *e
}

func buildExamplePaylinkTokenCreated() openapiclient.PaylinkTokenCreated {
	t := openapiclient.NewPaylinkTokenCreated(exampleTokenId, exampleTokenResult, exampleTokenToken)
	t.SetAttachments(buildExampleAttachmentRes())
	t.SetBps(exampleTokenBps)
	t.SetDateCreated(exampleTokenDate)
	t.SetErrors([]openapiclient.PaylinkErrorCode{buildExampleErrorCode()})
	t.SetIdentifier(exampleTokenIdentifier)
	t.SetMode(exampleTokenMode)
	t.SetQrcode(exampleTokenQrcode)
	t.SetServerVersion(exampleTokenServer)
	t.SetSource(exampleTokenSource)
	t.SetUrl(exampleTokenURL)
	t.SetUsc(exampleTokenUsc)
	return *t
}

func TestNewPaylinkTokenCreated(t *testing.T) {
	model := openapiclient.NewPaylinkTokenCreated(exampleTokenId, exampleTokenResult, exampleTokenToken)
	require.NotNil(t, model)
	assert.Equal(t, exampleTokenId, model.GetId())
	assert.Equal(t, exampleTokenResult, model.GetResult())
	assert.Equal(t, exampleTokenToken, model.GetToken())
	assert.False(t, model.HasMode())
}

func TestNewPaylinkTokenCreatedWithDefaults(t *testing.T) {
	model := openapiclient.NewPaylinkTokenCreatedWithDefaults()
	require.NotNil(t, model)
}

func TestPaylinkTokenCreatedSetGetCycle(t *testing.T) {
	model := openapiclient.NewPaylinkTokenCreated(exampleTokenId, exampleTokenResult, exampleTokenToken)
	att := buildExampleAttachmentRes()
	model.SetAttachments(att)
	assert.True(t, model.HasAttachments())
	assert.Equal(t, att, model.GetAttachments())

	model.SetBps(exampleTokenBps)
	assert.True(t, model.HasBps())
	assert.Equal(t, exampleTokenBps, model.GetBps())

	model.SetDateCreated(exampleTokenDate)
	assert.True(t, model.HasDateCreated())
	assert.Equal(t, exampleTokenDate, model.GetDateCreated())

	errs := []openapiclient.PaylinkErrorCode{buildExampleErrorCode()}
	model.SetErrors(errs)
	assert.True(t, model.HasErrors())
	assert.Len(t, model.GetErrors(), 1)

	model.SetIdentifier(exampleTokenIdentifier)
	assert.True(t, model.HasIdentifier())
	assert.Equal(t, exampleTokenIdentifier, model.GetIdentifier())

	model.SetMode(exampleTokenMode)
	assert.True(t, model.HasMode())
	assert.Equal(t, exampleTokenMode, model.GetMode())

	model.SetQrcode(exampleTokenQrcode)
	assert.True(t, model.HasQrcode())
	assert.Equal(t, exampleTokenQrcode, model.GetQrcode())

	model.SetServerVersion(exampleTokenServer)
	assert.True(t, model.HasServerVersion())
	assert.Equal(t, exampleTokenServer, model.GetServerVersion())

	model.SetSource(exampleTokenSource)
	assert.True(t, model.HasSource())
	assert.Equal(t, exampleTokenSource, model.GetSource())

	model.SetUrl(exampleTokenURL)
	assert.True(t, model.HasUrl())
	assert.Equal(t, exampleTokenURL, model.GetUrl())

	model.SetUsc(exampleTokenUsc)
	assert.True(t, model.HasUsc())
	assert.Equal(t, exampleTokenUsc, model.GetUsc())
}

func TestPaylinkTokenCreatedJSONRoundTrip(t *testing.T) {
	model := buildExamplePaylinkTokenCreated()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.PaylinkTokenCreated
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleTokenResult, unmarshalled.GetResult())
	assert.True(t, unmarshalled.HasUrl())
}

func TestPaylinkTokenCreatedToMap(t *testing.T) {
	model := buildExamplePaylinkTokenCreated()
	m, err := model.ToMap()
	require.NoError(t, err)
	assert.Equal(t, exampleTokenId, m["id"])
	if assert.Contains(t, m, "usc") {
		assert.Equal(t, &exampleTokenUsc, m["usc"])
	}
}

func TestNullablePaylinkTokenCreatedGetSet(t *testing.T) {
	base := buildExamplePaylinkTokenCreated()
	n := openapiclient.NullablePaylinkTokenCreated{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkTokenCreatedUnset(t *testing.T) {
	base := buildExamplePaylinkTokenCreated()
	n := openapiclient.NewNullablePaylinkTokenCreated(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullablePaylinkTokenCreatedJSONRoundTrip(t *testing.T) {
	base := buildExamplePaylinkTokenCreated()
	n := openapiclient.NewNullablePaylinkTokenCreated(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullablePaylinkTokenCreated
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleTokenUsc, newN.Get().GetUsc())
	}
}
