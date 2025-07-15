package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleAttData = "dGVzdA=="
var exampleAttFilename = "test.pdf"
var exampleAttMime = "application/pdf"
var exampleAttName = "Invoice"
var exampleAttRetention int32 = 10

func buildExamplePaylinkAttachmentRequest() openapiclient.PaylinkAttachmentRequest {
	a := openapiclient.NewPaylinkAttachmentRequest(exampleAttFilename, exampleAttMime)
	a.SetData(exampleAttData)
	a.SetName(exampleAttName)
	a.SetRetention(exampleAttRetention)
	return *a
}

func TestNewPaylinkAttachmentRequest(t *testing.T) {
	model := openapiclient.NewPaylinkAttachmentRequest(exampleAttFilename, exampleAttMime)
	require.NotNil(t, model)
	assert.Equal(t, exampleAttFilename, model.GetFilename())
	assert.Equal(t, exampleAttMime, model.GetMimeType())
	assert.False(t, model.HasData())
}

func TestNewPaylinkAttachmentRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewPaylinkAttachmentRequestWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasData())
	assert.False(t, model.HasName())
}

func TestPaylinkAttachmentRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewPaylinkAttachmentRequest(exampleAttFilename, exampleAttMime)
	model.SetData(exampleAttData)
	assert.True(t, model.HasData())
	assert.Equal(t, exampleAttData, model.GetData())

	model.SetName(exampleAttName)
	assert.True(t, model.HasName())
	assert.Equal(t, exampleAttName, model.GetName())

	model.SetRetention(exampleAttRetention)
	assert.True(t, model.HasRetention())
	assert.Equal(t, exampleAttRetention, model.GetRetention())
}

func TestPaylinkAttachmentRequestJSONRoundTrip(t *testing.T) {
	model := buildExamplePaylinkAttachmentRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.PaylinkAttachmentRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasName())
	assert.Equal(t, exampleAttName, unmarshalled.GetName())
}

func TestPaylinkAttachmentRequestToMap(t *testing.T) {
	model := buildExamplePaylinkAttachmentRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "data") {
		assert.Equal(t, &exampleAttData, m["data"])
	}
	assert.Equal(t, exampleAttFilename, m["filename"])
	assert.Equal(t, exampleAttMime, m["mime_type"])
	if assert.Contains(t, m, "name") {
		assert.Equal(t, &exampleAttName, m["name"])
	}
	if assert.Contains(t, m, "retention") {
		assert.Equal(t, &exampleAttRetention, m["retention"])
	}
}

func TestNullablePaylinkAttachmentRequestGetSet(t *testing.T) {
	base := buildExamplePaylinkAttachmentRequest()
	n := openapiclient.NullablePaylinkAttachmentRequest{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkAttachmentRequestUnset(t *testing.T) {
	base := buildExamplePaylinkAttachmentRequest()
	n := openapiclient.NewNullablePaylinkAttachmentRequest(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullablePaylinkAttachmentRequestJSONRoundTrip(t *testing.T) {
	base := buildExamplePaylinkAttachmentRequest()
	n := openapiclient.NewNullablePaylinkAttachmentRequest(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullablePaylinkAttachmentRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleAttMime, newN.Get().GetMimeType())
	}
}
