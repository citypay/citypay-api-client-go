package citypay

import (
    "encoding/json"
    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
)

var exampleSmsTemplate = "tpl"
var exampleSmsTo = "+4411111"

func buildExamplePaylinkSMSNotificationPath() openapiclient.PaylinkSMSNotificationPath {
    p := openapiclient.NewPaylinkSMSNotificationPath(exampleSmsTo)
    p.SetTemplate(exampleSmsTemplate)
    return *p
}

func TestNewPaylinkSMSNotificationPath(t *testing.T) {
    model := openapiclient.NewPaylinkSMSNotificationPath(exampleSmsTo)
    require.NotNil(t, model)
    assert.Equal(t, exampleSmsTo, model.GetTo())
    assert.False(t, model.HasTemplate())
}

func TestNewPaylinkSMSNotificationPathWithDefaults(t *testing.T) {
    model := openapiclient.NewPaylinkSMSNotificationPathWithDefaults()
    require.NotNil(t, model)
    assert.False(t, model.HasTemplate())
}

func TestPaylinkSMSNotificationPathSetGetCycle(t *testing.T) {
    model := openapiclient.NewPaylinkSMSNotificationPath(exampleSmsTo)
    model.SetTemplate(exampleSmsTemplate)
    assert.True(t, model.HasTemplate())
    assert.Equal(t, exampleSmsTemplate, model.GetTemplate())
}

func TestPaylinkSMSNotificationPathJSONRoundTrip(t *testing.T) {
    model := buildExamplePaylinkSMSNotificationPath()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.PaylinkSMSNotificationPath
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.True(t, unmarshalled.HasTemplate())
    assert.Equal(t, exampleSmsTemplate, unmarshalled.GetTemplate())
}

func TestPaylinkSMSNotificationPathToMap(t *testing.T) {
    model := buildExamplePaylinkSMSNotificationPath()
    m, err := model.ToMap()
    require.NoError(t, err)
    assert.Equal(t, exampleSmsTo, m["to"])
    if assert.Contains(t, m, "template") {
        assert.Equal(t, &exampleSmsTemplate, m["template"])
    }
}

func TestNullablePaylinkSMSNotificationPathGetSet(t *testing.T) {
    base := buildExamplePaylinkSMSNotificationPath()
    n := openapiclient.NullablePaylinkSMSNotificationPath{}
    n.Set(&base)
    require.True(t, n.IsSet())
    assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkSMSNotificationPathUnset(t *testing.T) {
    base := buildExamplePaylinkSMSNotificationPath()
    n := openapiclient.NewNullablePaylinkSMSNotificationPath(&base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullablePaylinkSMSNotificationPathJSONRoundTrip(t *testing.T) {
    base := buildExamplePaylinkSMSNotificationPath()
    n := openapiclient.NewNullablePaylinkSMSNotificationPath(&base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullablePaylinkSMSNotificationPath
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, exampleSmsTo, newN.Get().GetTo())
    }
}

