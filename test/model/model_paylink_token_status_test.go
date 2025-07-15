package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var exampleStatusAmount int32 = 100
var exampleStatusAuthCode = "AUTH"
var exampleStatusCard = "VISA"
var exampleStatusCreated = time.Date(2024, time.April, 1, 0, 0, 0, 0, time.UTC)
var exampleStatusDatetime = time.Date(2024, time.April, 2, 0, 0, 0, 0, time.UTC)
var exampleStatusIdentifier = "id"
var exampleStatusIsPaid = true
var exampleStatusMid int32 = 2
var exampleStatusToken = "tok"
var exampleStatusTransNo int32 = 5

func buildExamplePaylinkStateEventSlice() []openapiclient.PaylinkStateEvent {
	e := openapiclient.NewPaylinkStateEvent()
	e.SetState("CREATED")
	return []openapiclient.PaylinkStateEvent{*e}
}

func buildExamplePaylinkTokenStatus() openapiclient.PaylinkTokenStatus {
	s := openapiclient.NewPaylinkTokenStatus()
	s.SetAmountPaid(exampleStatusAmount)
	s.SetAuthCode(exampleStatusAuthCode)
	s.SetCard(exampleStatusCard)
	s.SetCreated(exampleStatusCreated)
	s.SetDatetime(exampleStatusDatetime)
	s.SetIdentifier(exampleStatusIdentifier)
	s.SetIsPaid(exampleStatusIsPaid)
	s.SetMid(exampleStatusMid)
	s.SetStateHistory(buildExamplePaylinkStateEventSlice())
	s.SetToken(exampleStatusToken)
	s.SetTransNo(exampleStatusTransNo)
	return *s
}

func TestNewPaylinkTokenStatus(t *testing.T) {
	model := openapiclient.NewPaylinkTokenStatus()
	require.NotNil(t, model)
	assert.False(t, model.HasToken())
}

func TestNewPaylinkTokenStatusWithDefaults(t *testing.T) {
	model := openapiclient.NewPaylinkTokenStatusWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasIsPaid())
}

func TestPaylinkTokenStatusSetGetCycle(t *testing.T) {
	model := openapiclient.NewPaylinkTokenStatus()
	model.SetAmountPaid(exampleStatusAmount)
	assert.True(t, model.HasAmountPaid())
	assert.Equal(t, exampleStatusAmount, model.GetAmountPaid())

	model.SetAuthCode(exampleStatusAuthCode)
	assert.True(t, model.HasAuthCode())
	assert.Equal(t, exampleStatusAuthCode, model.GetAuthCode())

	model.SetCard(exampleStatusCard)
	assert.True(t, model.HasCard())
	assert.Equal(t, exampleStatusCard, model.GetCard())

	model.SetCreated(exampleStatusCreated)
	assert.True(t, model.HasCreated())
	assert.Equal(t, exampleStatusCreated, model.GetCreated())

	model.SetDatetime(exampleStatusDatetime)
	assert.True(t, model.HasDatetime())
	assert.Equal(t, exampleStatusDatetime, model.GetDatetime())

	model.SetIdentifier(exampleStatusIdentifier)
	assert.True(t, model.HasIdentifier())
	assert.Equal(t, exampleStatusIdentifier, model.GetIdentifier())

	model.SetIsPaid(exampleStatusIsPaid)
	assert.True(t, model.HasIsPaid())
	assert.Equal(t, exampleStatusIsPaid, model.GetIsPaid())

	model.SetMid(exampleStatusMid)
	assert.True(t, model.HasMid())
	assert.Equal(t, exampleStatusMid, model.GetMid())

	history := buildExamplePaylinkStateEventSlice()
	model.SetStateHistory(history)
	assert.True(t, model.HasStateHistory())
	assert.Len(t, model.GetStateHistory(), 1)

	model.SetToken(exampleStatusToken)
	assert.True(t, model.HasToken())
	assert.Equal(t, exampleStatusToken, model.GetToken())

	model.SetTransNo(exampleStatusTransNo)
	assert.True(t, model.HasTransNo())
	assert.Equal(t, exampleStatusTransNo, model.GetTransNo())
}

func TestPaylinkTokenStatusJSONRoundTrip(t *testing.T) {
	model := buildExamplePaylinkTokenStatus()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.PaylinkTokenStatus
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasMid())
	assert.Equal(t, exampleStatusMid, unmarshalled.GetMid())
}

func TestPaylinkTokenStatusToMap(t *testing.T) {
	model := buildExamplePaylinkTokenStatus()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "token") {
		assert.Equal(t, &exampleStatusToken, m["token"])
	}
}

func TestNullablePaylinkTokenStatusGetSet(t *testing.T) {
	base := buildExamplePaylinkTokenStatus()
	n := openapiclient.NullablePaylinkTokenStatus{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkTokenStatusUnset(t *testing.T) {
	base := buildExamplePaylinkTokenStatus()
	n := openapiclient.NewNullablePaylinkTokenStatus(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullablePaylinkTokenStatusJSONRoundTrip(t *testing.T) {
	base := buildExamplePaylinkTokenStatus()
	n := openapiclient.NewNullablePaylinkTokenStatus(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullablePaylinkTokenStatus
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleStatusTransNo, newN.Get().GetTransNo())
	}
}
