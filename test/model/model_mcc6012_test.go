package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleAcc = "12345678"
var exampleDob = "1980-01-01"
var exampleLast = "Smith"
var examplePost = "AB12"

func buildExampleMCC6012() openapiclient.MCC6012 {
	m := openapiclient.NewMCC6012()
	m.SetRecipientAccount(exampleAcc)
	m.SetRecipientDob(exampleDob)
	m.SetRecipientLastname(exampleLast)
	m.SetRecipientPostcode(examplePost)
	return *m
}

func TestNewMCC6012(t *testing.T) {
	model := openapiclient.NewMCC6012()
	require.NotNil(t, model)
	assert.False(t, model.HasRecipientAccount())
	assert.False(t, model.HasRecipientDob())
	assert.False(t, model.HasRecipientLastname())
	assert.False(t, model.HasRecipientPostcode())
}

func TestNewMCC6012WithDefaults(t *testing.T) {
	model := openapiclient.NewMCC6012WithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasRecipientAccount())
	assert.False(t, model.HasRecipientDob())
	assert.False(t, model.HasRecipientLastname())
	assert.False(t, model.HasRecipientPostcode())
}

func TestMCC6012SetGetCycle(t *testing.T) {
	model := openapiclient.NewMCC6012()
	model.SetRecipientAccount(exampleAcc)
	assert.True(t, model.HasRecipientAccount())
	assert.Equal(t, exampleAcc, model.GetRecipientAccount())
	if val, ok := model.GetRecipientAccountOk(); assert.True(t, ok) {
		assert.Equal(t, exampleAcc, *val)
	}

	model.SetRecipientDob(exampleDob)
	assert.True(t, model.HasRecipientDob())
	assert.Equal(t, exampleDob, model.GetRecipientDob())

	model.SetRecipientLastname(exampleLast)
	assert.True(t, model.HasRecipientLastname())
	assert.Equal(t, exampleLast, model.GetRecipientLastname())

	model.SetRecipientPostcode(examplePost)
	assert.True(t, model.HasRecipientPostcode())
	assert.Equal(t, examplePost, model.GetRecipientPostcode())
}

func TestMCC6012JSONRoundTrip(t *testing.T) {
	model := buildExampleMCC6012()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.MCC6012
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasRecipientAccount())
	assert.Equal(t, examplePost, unmarshalled.GetRecipientPostcode())
}

func TestMCC6012ToMap(t *testing.T) {
	model := buildExampleMCC6012()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "recipient_account") {
		assert.Equal(t, &exampleAcc, m["recipient_account"])
	}
	if assert.Contains(t, m, "recipient_dob") {
		assert.Equal(t, &exampleDob, m["recipient_dob"])
	}
	if assert.Contains(t, m, "recipient_lastname") {
		assert.Equal(t, &exampleLast, m["recipient_lastname"])
	}
	if assert.Contains(t, m, "recipient_postcode") {
		assert.Equal(t, &examplePost, m["recipient_postcode"])
	}
}

func TestNullableMCC6012GetSet(t *testing.T) {
	base := buildExampleMCC6012()
	n := openapiclient.NullableMCC6012{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableMCC6012Unset(t *testing.T) {
	base := buildExampleMCC6012()
	n := openapiclient.NewNullableMCC6012(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableMCC6012JSONRoundTrip(t *testing.T) {
	base := buildExampleMCC6012()
	n := openapiclient.NewNullableMCC6012(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableMCC6012
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleAcc, newN.Get().GetRecipientAccount())
	}
}
