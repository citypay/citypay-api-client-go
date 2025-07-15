package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// helper to build example contact details
func buildExampleContact() openapiclient.ContactDetails {
	c := openapiclient.NewContactDetails()
	c.SetAddress1("123 Example Road")
	c.SetEmail("test@example.com")
	c.SetFirstname("Jane")
	c.SetLastname("Doe")
	c.SetCountry("GB")
	return *c
}

var exampleAccountID = "acc123"

func TestNewAccountCreate(t *testing.T) {
	model := openapiclient.NewAccountCreate(exampleAccountID)
	require.NotNil(t, model)
	assert.Equal(t, exampleAccountID, model.GetAccountId())
	assert.False(t, model.HasContact())
	contact, ok := model.GetContactOk()
	assert.False(t, ok)
	assert.Nil(t, contact)
}

func TestNewAccountCreateWithDefaults(t *testing.T) {
	model := openapiclient.NewAccountCreateWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, "", model.GetAccountId())
	assert.False(t, model.HasContact())
}

func TestAccountCreateSetGetCycle(t *testing.T) {
	model := openapiclient.NewAccountCreate(exampleAccountID)
	model.SetAccountId("updated")
	assert.Equal(t, "updated", model.GetAccountId())
	id, ok := model.GetAccountIdOk()
	require.True(t, ok)
	if assert.NotNil(t, id) {
		assert.Equal(t, "updated", *id)
	}

	contact := buildExampleContact()
	model.SetContact(contact)
	assert.True(t, model.HasContact())
	got := model.GetContact()
	assert.Equal(t, contact, got)
	gotPtr, ok := model.GetContactOk()
	require.True(t, ok)
	if assert.NotNil(t, gotPtr) {
		assert.Equal(t, contact, *gotPtr)
	}
}

func TestAccountCreateJSONRoundTrip(t *testing.T) {
	model := openapiclient.NewAccountCreate(exampleAccountID)
	contact := buildExampleContact()
	model.SetContact(contact)
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.AccountCreate
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleAccountID, unmarshalled.GetAccountId())
	assert.True(t, unmarshalled.HasContact())
	assert.Equal(t, contact, unmarshalled.GetContact())
}

func TestAccountCreateToMap(t *testing.T) {
	model := openapiclient.NewAccountCreate(exampleAccountID)
	contact := buildExampleContact()
	model.SetContact(contact)
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "account_id") {
		assert.Equal(t, exampleAccountID, m["account_id"])
	}
	if assert.Contains(t, m, "contact") {
		if assert.NotNil(t, m["contact"]) {
			assert.Equal(t, &contact, m["contact"])
		}
	}
}

func TestNullableAccountCreateGetSet(t *testing.T) {
	base := openapiclient.NewAccountCreate(exampleAccountID)
	base.SetContact(buildExampleContact())
	n := openapiclient.NullableAccountCreate{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableAccountCreateUnset(t *testing.T) {
	base := openapiclient.NewAccountCreate(exampleAccountID)
	n := openapiclient.NewNullableAccountCreate(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableAccountCreateJSONRoundTrip(t *testing.T) {
	base := openapiclient.NewAccountCreate(exampleAccountID)
	base.SetContact(buildExampleContact())
	n := openapiclient.NewNullableAccountCreate(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableAccountCreate
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, base.GetAccountId(), newN.Get().GetAccountId())
		assert.Equal(t, base.GetContact(), newN.Get().GetContact())
	}
}
