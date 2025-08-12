package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var exampleEventDate = time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC)
var exampleEventMessage = "msg"
var exampleEventState = "CREATED"

func buildExamplePaylinkStateEvent() openapiclient.PaylinkStateEvent {
	e := openapiclient.NewPaylinkStateEvent()
	e.SetDatetime(exampleEventDate)
	e.SetMessage(exampleEventMessage)
	e.SetState(exampleEventState)
	return *e
}

func TestNewPaylinkStateEvent(t *testing.T) {
	model := openapiclient.NewPaylinkStateEvent()
	require.NotNil(t, model)
	assert.False(t, model.HasState())
}

func TestNewPaylinkStateEventWithDefaults(t *testing.T) {
	model := openapiclient.NewPaylinkStateEventWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasMessage())
}

func TestPaylinkStateEventSetGetCycle(t *testing.T) {
	model := openapiclient.NewPaylinkStateEvent()
	model.SetDatetime(exampleEventDate)
	assert.True(t, model.HasDatetime())
	assert.Equal(t, exampleEventDate, model.GetDatetime())

	model.SetMessage(exampleEventMessage)
	assert.True(t, model.HasMessage())
	assert.Equal(t, exampleEventMessage, model.GetMessage())

	model.SetState(exampleEventState)
	assert.True(t, model.HasState())
	assert.Equal(t, exampleEventState, model.GetState())
}

func TestPaylinkStateEventJSONRoundTrip(t *testing.T) {
	model := buildExamplePaylinkStateEvent()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.PaylinkStateEvent
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasState())
	assert.Equal(t, exampleEventState, unmarshalled.GetState())
}

func TestPaylinkStateEventToMap(t *testing.T) {
	model := buildExamplePaylinkStateEvent()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "message") {
		assert.Equal(t, &exampleEventMessage, m["message"])
	}
}

func TestNullablePaylinkStateEventGetSet(t *testing.T) {
	base := buildExamplePaylinkStateEvent()
	n := openapiclient.NullablePaylinkStateEvent{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkStateEventUnset(t *testing.T) {
	base := buildExamplePaylinkStateEvent()
	n := openapiclient.NewNullablePaylinkStateEvent(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullablePaylinkStateEventJSONRoundTrip(t *testing.T) {
	base := buildExamplePaylinkStateEvent()
	n := openapiclient.NewNullablePaylinkStateEvent(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullablePaylinkStateEvent
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleEventState, newN.Get().GetState())
	}
}
