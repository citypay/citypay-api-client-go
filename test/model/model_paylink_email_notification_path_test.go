package citypay

import (
    "encoding/json"
    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
)

var exampleEmailBCC = []string{"b@example.com"}
var exampleEmailCC = []string{"c@example.com"}
var exampleEmailReply = []string{"r@example.com"}
var exampleEmailTemplate = "tmpl"
var exampleEmailTo = []string{"to@example.com"}

func buildExamplePaylinkEmailNotificationPath() openapiclient.PaylinkEmailNotificationPath {
    p := openapiclient.NewPaylinkEmailNotificationPath(exampleEmailTo)
    p.SetBcc(exampleEmailBCC)
    p.SetCc(exampleEmailCC)
    p.SetReplyTo(exampleEmailReply)
    p.SetTemplate(exampleEmailTemplate)
    return *p
}

func TestNewPaylinkEmailNotificationPath(t *testing.T) {
    model := openapiclient.NewPaylinkEmailNotificationPath(exampleEmailTo)
    require.NotNil(t, model)
    assert.Equal(t, exampleEmailTo, model.GetTo())
    assert.False(t, model.HasBcc())
}

func TestNewPaylinkEmailNotificationPathWithDefaults(t *testing.T) {
    model := openapiclient.NewPaylinkEmailNotificationPathWithDefaults()
    require.NotNil(t, model)
    assert.False(t, model.HasTemplate())
}

func TestPaylinkEmailNotificationPathSetGetCycle(t *testing.T) {
    model := openapiclient.NewPaylinkEmailNotificationPath(exampleEmailTo)
    model.SetBcc(exampleEmailBCC)
    assert.True(t, model.HasBcc())
    assert.Equal(t, exampleEmailBCC, model.GetBcc())

    model.SetCc(exampleEmailCC)
    assert.True(t, model.HasCc())
    assert.Equal(t, exampleEmailCC, model.GetCc())

    model.SetReplyTo(exampleEmailReply)
    assert.True(t, model.HasReplyTo())
    assert.Equal(t, exampleEmailReply, model.GetReplyTo())

    model.SetTemplate(exampleEmailTemplate)
    assert.True(t, model.HasTemplate())
    assert.Equal(t, exampleEmailTemplate, model.GetTemplate())
}

func TestPaylinkEmailNotificationPathJSONRoundTrip(t *testing.T) {
    model := buildExamplePaylinkEmailNotificationPath()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.PaylinkEmailNotificationPath
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.True(t, unmarshalled.HasBcc())
    assert.Equal(t, exampleEmailBCC, unmarshalled.GetBcc())
}

func TestPaylinkEmailNotificationPathToMap(t *testing.T) {
    model := buildExamplePaylinkEmailNotificationPath()
    m, err := model.ToMap()
    require.NoError(t, err)
    if assert.Contains(t, m, "bcc") {
        assert.Equal(t, exampleEmailBCC, m["bcc"])
    }
    assert.Equal(t, exampleEmailTo, m["to"])
}

func TestNullablePaylinkEmailNotificationPathGetSet(t *testing.T) {
    base := buildExamplePaylinkEmailNotificationPath()
    n := openapiclient.NullablePaylinkEmailNotificationPath{}
    n.Set(&base)
    require.True(t, n.IsSet())
    assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkEmailNotificationPathUnset(t *testing.T) {
    base := buildExamplePaylinkEmailNotificationPath()
    n := openapiclient.NewNullablePaylinkEmailNotificationPath(&base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullablePaylinkEmailNotificationPathJSONRoundTrip(t *testing.T) {
    base := buildExamplePaylinkEmailNotificationPath()
    n := openapiclient.NewNullablePaylinkEmailNotificationPath(&base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullablePaylinkEmailNotificationPath
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, exampleEmailTemplate, newN.Get().GetTemplate())
    }
}

