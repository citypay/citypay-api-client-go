package citypay

import (
    "encoding/json"
    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
)

var exampleResendEmail = true
var exampleResendSms = true

func buildExamplePaylinkResendNotificationRequest() openapiclient.PaylinkResendNotificationRequest {
    r := openapiclient.NewPaylinkResendNotificationRequest()
    r.SetEmail(exampleResendEmail)
    r.SetSms(exampleResendSms)
    return *r
}

func TestNewPaylinkResendNotificationRequest(t *testing.T) {
    model := openapiclient.NewPaylinkResendNotificationRequest()
    require.NotNil(t, model)
    assert.False(t, model.HasEmail())
}

func TestNewPaylinkResendNotificationRequestWithDefaults(t *testing.T) {
    model := openapiclient.NewPaylinkResendNotificationRequestWithDefaults()
    require.NotNil(t, model)
    assert.False(t, model.HasSms())
}

func TestPaylinkResendNotificationRequestSetGetCycle(t *testing.T) {
    model := openapiclient.NewPaylinkResendNotificationRequest()
    model.SetEmail(exampleResendEmail)
    assert.True(t, model.HasEmail())
    assert.Equal(t, exampleResendEmail, model.GetEmail())

    model.SetSms(exampleResendSms)
    assert.True(t, model.HasSms())
    assert.Equal(t, exampleResendSms, model.GetSms())
}

func TestPaylinkResendNotificationRequestJSONRoundTrip(t *testing.T) {
    model := buildExamplePaylinkResendNotificationRequest()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.PaylinkResendNotificationRequest
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.True(t, unmarshalled.HasEmail())
    assert.Equal(t, exampleResendEmail, unmarshalled.GetEmail())
}

func TestPaylinkResendNotificationRequestToMap(t *testing.T) {
    model := buildExamplePaylinkResendNotificationRequest()
    m, err := model.ToMap()
    require.NoError(t, err)
    if assert.Contains(t, m, "email") {
        assert.Equal(t, &exampleResendEmail, m["email"])
    }
    if assert.Contains(t, m, "sms") {
        assert.Equal(t, &exampleResendSms, m["sms"])
    }
}

func TestNullablePaylinkResendNotificationRequestGetSet(t *testing.T) {
    base := buildExamplePaylinkResendNotificationRequest()
    n := openapiclient.NullablePaylinkResendNotificationRequest{}
    n.Set(&base)
    require.True(t, n.IsSet())
    assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkResendNotificationRequestUnset(t *testing.T) {
    base := buildExamplePaylinkResendNotificationRequest()
    n := openapiclient.NewNullablePaylinkResendNotificationRequest(&base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullablePaylinkResendNotificationRequestJSONRoundTrip(t *testing.T) {
    base := buildExamplePaylinkResendNotificationRequest()
    n := openapiclient.NewNullablePaylinkResendNotificationRequest(&base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullablePaylinkResendNotificationRequest
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, exampleResendSms, newN.Get().GetSms())
    }
}

