package citypay

import (
	"encoding/json"
	"testing"

	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func buildExampleAuthReferences() *openapiclient.AuthReferences {
	ref := buildExampleAuthReference()
	a := openapiclient.NewAuthReferences()
	a.SetAuths([]openapiclient.AuthReference{*ref})
	return a
}

func TestNewAuthReferences(t *testing.T) {
	model := openapiclient.NewAuthReferences()
	require.NotNil(t, model)
	assert.False(t, model.HasAuths())
}

func TestNewAuthReferencesWithDefaults(t *testing.T) {
	model := openapiclient.NewAuthReferencesWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasAuths())
}

func TestAuthReferencesSetGetCycle(t *testing.T) {
	model := openapiclient.NewAuthReferences()
	ref := buildExampleAuthReference()
	model.SetAuths([]openapiclient.AuthReference{*ref})
	assert.True(t, model.HasAuths())
	assert.Equal(t, 1, len(model.GetAuths()))
}

func TestAuthReferencesJSONRoundTrip(t *testing.T) {
	model := buildExampleAuthReferences()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.AuthReferences
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasAuths())
	assert.Equal(t, 1, len(unmarshalled.GetAuths()))
}

func TestAuthReferencesToMap(t *testing.T) {
	model := buildExampleAuthReferences()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "auths") {
		if assert.Len(t, m["auths"], 1) {
			first := m["auths"].([]openapiclient.AuthReference)[0]
			assert.Equal(t, exampleAuthRefAtrn, first.GetAtrn())
		}
	}
}

func TestNullableAuthReferencesGetSet(t *testing.T) {
	base := buildExampleAuthReferences()
	n := openapiclient.NullableAuthReferences{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableAuthReferencesUnset(t *testing.T) {
	base := buildExampleAuthReferences()
	n := openapiclient.NewNullableAuthReferences(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableAuthReferencesJSONRoundTrip(t *testing.T) {
	base := buildExampleAuthReferences()
	n := openapiclient.NewNullableAuthReferences(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableAuthReferences
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, 1, len(newN.Get().GetAuths()))
	}
}
