package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var exampleExistsActive = true
var exampleExistsValue = true
var exampleExistsDate = time.Date(2024, time.May, 1, 0, 0, 0, 0, time.UTC)

func buildExampleExists() *openapiclient.Exists {
	e := openapiclient.NewExists(exampleExistsValue)
	e.SetActive(exampleExistsActive)
	e.SetLastModified(exampleExistsDate)
	return e
}

func TestNewExists(t *testing.T) {
	model := openapiclient.NewExists(exampleExistsValue)
	require.NotNil(t, model)
	assert.Equal(t, exampleExistsValue, model.GetExists())
	assert.False(t, model.HasActive())
}

func TestNewExistsWithDefaults(t *testing.T) {
	model := openapiclient.NewExistsWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasActive())
	assert.Equal(t, false, model.GetExists())
}

func TestExistsSetGetCycle(t *testing.T) {
	model := openapiclient.NewExists(false)
	model.SetActive(exampleExistsActive)
	assert.True(t, model.HasActive())
	assert.Equal(t, exampleExistsActive, model.GetActive())

	model.SetExists(exampleExistsValue)
	assert.Equal(t, exampleExistsValue, model.GetExists())

	model.SetLastModified(exampleExistsDate)
	assert.True(t, model.HasLastModified())
	assert.Equal(t, exampleExistsDate, model.GetLastModified())
}

func TestExistsJSONRoundTrip(t *testing.T) {
	model := buildExampleExists()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.Exists
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasActive())
	assert.Equal(t, exampleExistsDate, unmarshalled.GetLastModified())
}

func TestExistsToMap(t *testing.T) {
	model := buildExampleExists()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "active") {
		assert.Equal(t, &exampleExistsActive, m["active"])
	}
	if assert.Contains(t, m, "exists") {
		assert.Equal(t, exampleExistsValue, m["exists"])
	}
}

func TestNullableExistsGetSet(t *testing.T) {
	base := buildExampleExists()
	n := openapiclient.NullableExists{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableExistsUnset(t *testing.T) {
	base := buildExampleExists()
	n := openapiclient.NewNullableExists(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableExistsJSONRoundTrip(t *testing.T) {
	base := buildExampleExists()
	n := openapiclient.NewNullableExists(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableExists
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleExistsValue, newN.Get().GetExists())
	}
}
