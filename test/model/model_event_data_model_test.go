package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleEndDate = "2024-12-31"
var exampleEventId = "EVT1"
var exampleOrganiser = "ORG1"
var exampleStartDate = "2024-01-01"
var examplePayType = "deposit"

func buildExampleEventDataModel() *openapiclient.EventDataModel {
	e := openapiclient.NewEventDataModel()
	e.SetEventEndDate(exampleEndDate)
	e.SetEventId(exampleEventId)
	e.SetEventOrganiserId(exampleOrganiser)
	e.SetEventStartDate(exampleStartDate)
	e.SetPaymentType(examplePayType)
	return e
}

func TestNewEventDataModel(t *testing.T) {
	model := openapiclient.NewEventDataModel()
	require.NotNil(t, model)
	assert.False(t, model.HasEventId())
}

func TestNewEventDataModelWithDefaults(t *testing.T) {
	model := openapiclient.NewEventDataModelWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasEventId())
}

func TestEventDataModelSetGetCycle(t *testing.T) {
	model := openapiclient.NewEventDataModel()
	model.SetEventEndDate(exampleEndDate)
	assert.True(t, model.HasEventEndDate())
	assert.Equal(t, exampleEndDate, model.GetEventEndDate())

	model.SetEventId(exampleEventId)
	assert.True(t, model.HasEventId())
	assert.Equal(t, exampleEventId, model.GetEventId())

	model.SetEventOrganiserId(exampleOrganiser)
	assert.True(t, model.HasEventOrganiserId())
	assert.Equal(t, exampleOrganiser, model.GetEventOrganiserId())

	model.SetEventStartDate(exampleStartDate)
	assert.True(t, model.HasEventStartDate())
	assert.Equal(t, exampleStartDate, model.GetEventStartDate())

	model.SetPaymentType(examplePayType)
	assert.True(t, model.HasPaymentType())
	assert.Equal(t, examplePayType, model.GetPaymentType())
}

func TestEventDataModelJSONRoundTrip(t *testing.T) {
	model := buildExampleEventDataModel()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.EventDataModel
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleEventId, unmarshalled.GetEventId())
	assert.True(t, unmarshalled.HasEventEndDate())
}

func TestEventDataModelToMap(t *testing.T) {
	model := buildExampleEventDataModel()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "event_id") {
		assert.Equal(t, &exampleEventId, m["event_id"])
	}
	if assert.Contains(t, m, "payment_type") {
		assert.Equal(t, &examplePayType, m["payment_type"])
	}
}

func TestNullableEventDataModelGetSet(t *testing.T) {
	base := buildExampleEventDataModel()
	n := openapiclient.NullableEventDataModel{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableEventDataModelUnset(t *testing.T) {
	base := buildExampleEventDataModel()
	n := openapiclient.NewNullableEventDataModel(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableEventDataModelJSONRoundTrip(t *testing.T) {
	base := buildExampleEventDataModel()
	n := openapiclient.NewNullableEventDataModel(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableEventDataModel
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleEventId, newN.Get().GetEventId())
	}
}
