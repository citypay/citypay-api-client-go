package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleAdjustmentsAdjustmentType = "surcharge"
var exampleAdjustmentsAmount int32 = 100
var exampleAdjustmentsAccumulate = "AccumulateBase"

func buildExampleAdjustments() *openapiclient.Adjustments {
	cond := buildExampleAdjustmentCondition()
	m := openapiclient.NewAdjustments(exampleAdjustmentsAdjustmentType)
	m.SetAmount(exampleAdjustmentsAmount)
	m.SetAccumulate(exampleAdjustmentsAccumulate)
	m.SetConditions(*cond)
	return m
}

func TestNewAdjustments(t *testing.T) {
	model := openapiclient.NewAdjustments(exampleAdjustmentsAdjustmentType)
	require.NotNil(t, model)
	assert.Equal(t, exampleAdjustmentsAdjustmentType, model.GetAdjustment())
	assert.False(t, model.HasAmount())
}

func TestNewAdjustmentsWithDefaults(t *testing.T) {
	model := openapiclient.NewAdjustmentsWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasAccumulate())
}

func TestAdjustmentsSetGetCycle(t *testing.T) {
	model := openapiclient.NewAdjustments(exampleAdjustmentsAdjustmentType)
	model.SetAmount(exampleAdjustmentsAmount)
	assert.True(t, model.HasAmount())
	assert.Equal(t, exampleAdjustmentsAmount, model.GetAmount())
	amtPtr, ok := model.GetAmountOk()
	require.True(t, ok)
	if assert.NotNil(t, amtPtr) {
		assert.Equal(t, exampleAdjustmentsAmount, *amtPtr)
	}
}

func TestAdjustmentsJSONRoundTrip(t *testing.T) {
	model := buildExampleAdjustments()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.Adjustments
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleAdjustmentsAdjustmentType, unmarshalled.GetAdjustment())
	assert.True(t, unmarshalled.HasConditions())
}

func TestAdjustmentsToMap(t *testing.T) {
	model := buildExampleAdjustments()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "adjustment") {
		assert.Equal(t, exampleAdjustmentsAdjustmentType, m["adjustment"])
	}
	if assert.Contains(t, m, "amount") {
		assert.Equal(t, &exampleAdjustmentsAmount, m["amount"])
	}
}

func TestNullableAdjustmentsGetSet(t *testing.T) {
	base := buildExampleAdjustments()
	n := openapiclient.NullableAdjustments{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableAdjustmentsUnset(t *testing.T) {
	base := buildExampleAdjustments()
	n := openapiclient.NewNullableAdjustments(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableAdjustmentsJSONRoundTrip(t *testing.T) {
	base := buildExampleAdjustments()
	n := openapiclient.NewNullableAdjustments(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableAdjustments
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleAdjustmentsAdjustmentType, newN.Get().GetAdjustment())
	}
}
