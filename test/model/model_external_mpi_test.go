package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleMPIAuthenResult = "Y"
var exampleMPICavv = "abc123"
var exampleMPIEci int32 = 5
var exampleMPIEnrolled = "Y"
var exampleMPIXid = "xyz987"

func buildExampleExternalMPI() openapiclient.ExternalMPI {
	m := openapiclient.NewExternalMPI()
	m.SetAuthenResult(exampleMPIAuthenResult)
	m.SetCavv(exampleMPICavv)
	m.SetEci(exampleMPIEci)
	m.SetEnrolled(exampleMPIEnrolled)
	m.SetXid(exampleMPIXid)
	return *m
}

func TestNewExternalMPI(t *testing.T) {
	model := openapiclient.NewExternalMPI()
	require.NotNil(t, model)
	assert.False(t, model.HasAuthenResult())
	assert.False(t, model.HasCavv())
	assert.False(t, model.HasEci())
	assert.False(t, model.HasEnrolled())
	assert.False(t, model.HasXid())
}

func TestNewExternalMPIWithDefaults(t *testing.T) {
	model := openapiclient.NewExternalMPIWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasAuthenResult())
	assert.False(t, model.HasCavv())
	assert.False(t, model.HasEci())
	assert.False(t, model.HasEnrolled())
	assert.False(t, model.HasXid())
}

func TestExternalMPISetGetCycle(t *testing.T) {
	model := openapiclient.NewExternalMPI()
	model.SetAuthenResult("N")
	assert.Equal(t, "N", model.GetAuthenResult())
	if val, ok := model.GetAuthenResultOk(); assert.True(t, ok) {
		assert.Equal(t, "N", *val)
	}

	model.SetCavv(exampleMPICavv)
	assert.True(t, model.HasCavv())
	assert.Equal(t, exampleMPICavv, model.GetCavv())

	model.SetEci(exampleMPIEci)
	assert.True(t, model.HasEci())
	assert.Equal(t, exampleMPIEci, model.GetEci())

	model.SetEnrolled(exampleMPIEnrolled)
	assert.True(t, model.HasEnrolled())
	assert.Equal(t, exampleMPIEnrolled, model.GetEnrolled())

	model.SetXid(exampleMPIXid)
	assert.True(t, model.HasXid())
	assert.Equal(t, exampleMPIXid, model.GetXid())
}

func TestExternalMPIJSONRoundTrip(t *testing.T) {
	model := buildExampleExternalMPI()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.ExternalMPI
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasAuthenResult())
	assert.Equal(t, exampleMPICavv, unmarshalled.GetCavv())
	assert.True(t, unmarshalled.HasEci())
	assert.Equal(t, exampleMPIEci, unmarshalled.GetEci())
	assert.True(t, unmarshalled.HasEnrolled())
	assert.Equal(t, exampleMPIXid, unmarshalled.GetXid())
}

func TestExternalMPIToMap(t *testing.T) {
	model := buildExampleExternalMPI()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "authen_result") {
		assert.Equal(t, &exampleMPIAuthenResult, m["authen_result"])
	}
	if assert.Contains(t, m, "cavv") {
		assert.Equal(t, &exampleMPICavv, m["cavv"])
	}
	if assert.Contains(t, m, "eci") {
		assert.Equal(t, &exampleMPIEci, m["eci"])
	}
	if assert.Contains(t, m, "enrolled") {
		assert.Equal(t, &exampleMPIEnrolled, m["enrolled"])
	}
	if assert.Contains(t, m, "xid") {
		assert.Equal(t, &exampleMPIXid, m["xid"])
	}
}

func TestNullableExternalMPIGetSet(t *testing.T) {
	base := buildExampleExternalMPI()
	n := openapiclient.NullableExternalMPI{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableExternalMPIUnset(t *testing.T) {
	base := buildExampleExternalMPI()
	n := openapiclient.NewNullableExternalMPI(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableExternalMPIJSONRoundTrip(t *testing.T) {
	base := buildExampleExternalMPI()
	n := openapiclient.NewNullableExternalMPI(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableExternalMPI
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleMPIEci, newN.Get().GetEci())
	}
}
