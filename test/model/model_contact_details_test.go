package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleAddress1 = "1 Test Way"
var exampleAddress2 = "Suite 5"
var exampleAddress3 = "Floor 3"
var exampleArea = "Test City"
var exampleCompany = "Test Co"
var exampleCountry = "GB"
var exampleEmail = "contact@example.com"
var exampleFirstname = "Jane"
var exampleLastname = "Doe"
var exampleMobile = "07123456789"
var examplePostcode = "TE5 7ST"
var exampleTelephone = "01111111111"
var exampleTitle = "Ms"

func buildExampleContactDetails() *openapiclient.ContactDetails {
	c := openapiclient.NewContactDetails()
	c.SetAddress1(exampleAddress1)
	c.SetAddress2(exampleAddress2)
	c.SetAddress3(exampleAddress3)
	c.SetArea(exampleArea)
	c.SetCompany(exampleCompany)
	c.SetCountry(exampleCountry)
	c.SetEmail(exampleEmail)
	c.SetFirstname(exampleFirstname)
	c.SetLastname(exampleLastname)
	c.SetMobileNo(exampleMobile)
	c.SetPostcode(examplePostcode)
	c.SetTelephoneNo(exampleTelephone)
	c.SetTitle(exampleTitle)
	return c
}

func TestNewContactDetails(t *testing.T) {
	model := openapiclient.NewContactDetails()
	require.NotNil(t, model)
	assert.False(t, model.HasAddress1())
	assert.False(t, model.HasCountry())
}

func TestNewContactDetailsWithDefaults(t *testing.T) {
	model := openapiclient.NewContactDetailsWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasAddress1())
	assert.Equal(t, "", model.GetAddress1())
}

func TestContactDetailsSetGetCycle(t *testing.T) {
	model := openapiclient.NewContactDetails()
	model.SetAddress1(exampleAddress1)
	assert.True(t, model.HasAddress1())
	assert.Equal(t, exampleAddress1, model.GetAddress1())
	if val, ok := model.GetAddress1Ok(); assert.True(t, ok) {
		assert.Equal(t, exampleAddress1, *val)
	}

	model.SetAddress2(exampleAddress2)
	assert.True(t, model.HasAddress2())
	assert.Equal(t, exampleAddress2, model.GetAddress2())

	model.SetAddress3(exampleAddress3)
	assert.True(t, model.HasAddress3())
	assert.Equal(t, exampleAddress3, model.GetAddress3())

	model.SetArea(exampleArea)
	assert.True(t, model.HasArea())
	assert.Equal(t, exampleArea, model.GetArea())

	model.SetCompany(exampleCompany)
	assert.True(t, model.HasCompany())
	assert.Equal(t, exampleCompany, model.GetCompany())

	model.SetCountry(exampleCountry)
	assert.True(t, model.HasCountry())
	assert.Equal(t, exampleCountry, model.GetCountry())

	model.SetEmail(exampleEmail)
	assert.True(t, model.HasEmail())
	assert.Equal(t, exampleEmail, model.GetEmail())

	model.SetFirstname(exampleFirstname)
	assert.True(t, model.HasFirstname())
	assert.Equal(t, exampleFirstname, model.GetFirstname())

	model.SetLastname(exampleLastname)
	assert.True(t, model.HasLastname())
	assert.Equal(t, exampleLastname, model.GetLastname())

	model.SetMobileNo(exampleMobile)
	assert.True(t, model.HasMobileNo())
	assert.Equal(t, exampleMobile, model.GetMobileNo())

	model.SetPostcode(examplePostcode)
	assert.True(t, model.HasPostcode())
	assert.Equal(t, examplePostcode, model.GetPostcode())

	model.SetTelephoneNo(exampleTelephone)
	assert.True(t, model.HasTelephoneNo())
	assert.Equal(t, exampleTelephone, model.GetTelephoneNo())

	model.SetTitle(exampleTitle)
	assert.True(t, model.HasTitle())
	assert.Equal(t, exampleTitle, model.GetTitle())
}

func TestContactDetailsJSONRoundTrip(t *testing.T) {
	model := buildExampleContactDetails()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.ContactDetails
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleAddress1, unmarshalled.GetAddress1())
	assert.Equal(t, exampleCountry, unmarshalled.GetCountry())
	assert.Equal(t, exampleEmail, unmarshalled.GetEmail())
}

func TestContactDetailsToMap(t *testing.T) {
	model := buildExampleContactDetails()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "address1") {
		assert.Equal(t, &exampleAddress1, m["address1"])
	}
	if assert.Contains(t, m, "country") {
		assert.Equal(t, &exampleCountry, m["country"])
	}
}

func TestNullableContactDetailsGetSet(t *testing.T) {
	base := buildExampleContactDetails()
	n := openapiclient.NullableContactDetails{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableContactDetailsUnset(t *testing.T) {
	base := buildExampleContactDetails()
	n := openapiclient.NewNullableContactDetails(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableContactDetailsJSONRoundTrip(t *testing.T) {
	base := buildExampleContactDetails()
	n := openapiclient.NewNullableContactDetails(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableContactDetails
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleEmail, newN.Get().GetEmail())
	}
}
