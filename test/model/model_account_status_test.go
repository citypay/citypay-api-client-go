package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// example data used across tests
var exampleStatus = "ACTIVE"

func TestNewAccountStatus(t *testing.T) {
	model := openapiclient.NewAccountStatus()
	require.NotNil(t, model)
	assert.False(t, model.HasStatus())
	assert.Equal(t, "", model.GetStatus())
	val, ok := model.GetStatusOk()
	assert.False(t, ok)
	assert.Nil(t, val)
}

func TestNewAccountStatusWithDefaults(t *testing.T) {
	model := openapiclient.NewAccountStatusWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasStatus())
	assert.Equal(t, "", model.GetStatus())
}

func TestAccountStatusSetGetCycle(t *testing.T) {
	model := openapiclient.NewAccountStatus()
	model.SetStatus(exampleStatus)
	assert.True(t, model.HasStatus())
	assert.Equal(t, exampleStatus, model.GetStatus())
	val, ok := model.GetStatusOk()
	require.True(t, ok)
	if assert.NotNil(t, val) {
		assert.Equal(t, exampleStatus, *val)
	}
}

func TestAccountStatusJSONRoundTrip(t *testing.T) {
	model := openapiclient.NewAccountStatus()
	model.SetStatus(exampleStatus)
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.AccountStatus
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasStatus())
	assert.Equal(t, exampleStatus, unmarshalled.GetStatus())
}

func TestAccountStatusToMap(t *testing.T) {
	model := openapiclient.NewAccountStatus()
	model.SetStatus(exampleStatus)
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "status") {
		assert.Equal(t, &exampleStatus, m["status"])
	}
}

func TestNullableAccountStatusGetSet(t *testing.T) {
	base := openapiclient.NewAccountStatus()
	base.SetStatus(exampleStatus)
	n := openapiclient.NullableAccountStatus{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableAccountStatusUnset(t *testing.T) {
	base := openapiclient.NewAccountStatus()
	n := openapiclient.NewNullableAccountStatus(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableAccountStatusJSONRoundTrip(t *testing.T) {
	base := openapiclient.NewAccountStatus()
	base.SetStatus(exampleStatus)
	n := openapiclient.NewNullableAccountStatus(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableAccountStatus
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleStatus, newN.Get().GetStatus())
	}
}
