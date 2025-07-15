package citypay

import (
    "encoding/json"
    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
)

var exampleAttResName = "Invoice"
var exampleAttResResult = "OK"
var exampleAttResURL = "https://example.com/upload"

func buildExamplePaylinkAttachmentResult() openapiclient.PaylinkAttachmentResult {
    r := openapiclient.NewPaylinkAttachmentResult(exampleAttResName, exampleAttResResult)
    r.SetUrl(exampleAttResURL)
    return *r
}

func TestNewPaylinkAttachmentResult(t *testing.T) {
    model := openapiclient.NewPaylinkAttachmentResult(exampleAttResName, exampleAttResResult)
    require.NotNil(t, model)
    assert.Equal(t, exampleAttResName, model.GetName())
    assert.Equal(t, exampleAttResResult, model.GetResult())
    assert.False(t, model.HasUrl())
}

func TestNewPaylinkAttachmentResultWithDefaults(t *testing.T) {
    model := openapiclient.NewPaylinkAttachmentResultWithDefaults()
    require.NotNil(t, model)
    assert.False(t, model.HasUrl())
}

func TestPaylinkAttachmentResultSetGetCycle(t *testing.T) {
    model := openapiclient.NewPaylinkAttachmentResult(exampleAttResName, exampleAttResResult)
    model.SetUrl(exampleAttResURL)
    assert.True(t, model.HasUrl())
    assert.Equal(t, exampleAttResURL, model.GetUrl())
}

func TestPaylinkAttachmentResultJSONRoundTrip(t *testing.T) {
    model := buildExamplePaylinkAttachmentResult()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.PaylinkAttachmentResult
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.Equal(t, exampleAttResResult, unmarshalled.GetResult())
    assert.True(t, unmarshalled.HasUrl())
}

func TestPaylinkAttachmentResultToMap(t *testing.T) {
    model := buildExamplePaylinkAttachmentResult()
    m, err := model.ToMap()
    require.NoError(t, err)
    assert.Equal(t, exampleAttResName, m["name"])
    assert.Equal(t, exampleAttResResult, m["result"])
    if assert.Contains(t, m, "url") {
        assert.Equal(t, &exampleAttResURL, m["url"])
    }
}

func TestNullablePaylinkAttachmentResultGetSet(t *testing.T) {
    base := buildExamplePaylinkAttachmentResult()
    n := openapiclient.NullablePaylinkAttachmentResult{}
    n.Set(&base)
    require.True(t, n.IsSet())
    assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkAttachmentResultUnset(t *testing.T) {
    base := buildExamplePaylinkAttachmentResult()
    n := openapiclient.NewNullablePaylinkAttachmentResult(&base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullablePaylinkAttachmentResultJSONRoundTrip(t *testing.T) {
    base := buildExamplePaylinkAttachmentResult()
    n := openapiclient.NewNullablePaylinkAttachmentResult(&base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullablePaylinkAttachmentResult
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, exampleAttResName, newN.Get().GetName())
    }
}

