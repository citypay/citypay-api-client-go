package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleAdjustmentConditionAnchor = "after_creation"
var exampleAdjustmentConditionDiscountCode = "PROMO"
var exampleAdjustmentConditionDuration int32 = 3

func buildExampleAdjustmentCondition() *openapiclient.AdjustmentCondition {
	m := openapiclient.NewAdjustmentCondition()
	m.SetAnchor(exampleAdjustmentConditionAnchor)
	m.SetDiscountCode(exampleAdjustmentConditionDiscountCode)
	m.SetDuration(exampleAdjustmentConditionDuration)
	return m
}

func TestNewAdjustmentCondition(t *testing.T) {
	model := openapiclient.NewAdjustmentCondition()
	require.NotNil(t, model)
	assert.False(t, model.HasAnchor())
}

func TestNewAdjustmentConditionWithDefaults(t *testing.T) {
	model := openapiclient.NewAdjustmentConditionWithDefaults()
	require.NotNil(t, model)
}

func TestAdjustmentConditionSetGetCycle(t *testing.T) {
	model := openapiclient.NewAdjustmentCondition()
	model.SetAnchor(exampleAdjustmentConditionAnchor)
	assert.True(t, model.HasAnchor())
	assert.Equal(t, exampleAdjustmentConditionAnchor, model.GetAnchor())
	anchorPtr, ok := model.GetAnchorOk()
	require.True(t, ok)
	if assert.NotNil(t, anchorPtr) {
		assert.Equal(t, exampleAdjustmentConditionAnchor, *anchorPtr)
	}
}

func TestAdjustmentConditionJSONRoundTrip(t *testing.T) {
	model := buildExampleAdjustmentCondition()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.AdjustmentCondition
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleAdjustmentConditionAnchor, unmarshalled.GetAnchor())
	assert.Equal(t, exampleAdjustmentConditionDiscountCode, unmarshalled.GetDiscountCode())
}

func TestAdjustmentConditionToMap(t *testing.T) {
	model := buildExampleAdjustmentCondition()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "anchor") {
		assert.Equal(t, &exampleAdjustmentConditionAnchor, m["anchor"])
	}
}

func TestNullableAdjustmentConditionGetSet(t *testing.T) {
	base := buildExampleAdjustmentCondition()
	n := openapiclient.NullableAdjustmentCondition{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableAdjustmentConditionUnset(t *testing.T) {
	base := buildExampleAdjustmentCondition()
	n := openapiclient.NewNullableAdjustmentCondition(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableAdjustmentConditionJSONRoundTrip(t *testing.T) {
	base := buildExampleAdjustmentCondition()
	n := openapiclient.NewNullableAdjustmentCondition(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableAdjustmentCondition
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleAdjustmentConditionAnchor, newN.Get().GetAnchor())
	}
}
