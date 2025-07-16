package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleHolderAccept = "text/html"
var exampleHolderCompany = "Company"
var exampleHolderEmail = "a@example.com"
var exampleHolderFirst = "John"
var exampleHolderLast = "Smith"
var exampleHolderMobile = "+4411111111"
var exampleHolderRemote = "1.1.1.1"
var exampleHolderTitle = "Mr"
var exampleHolderUA = "UA"

func buildPaylinkCardHolderAddress() openapiclient.PaylinkAddress {
	a := openapiclient.NewPaylinkAddress()
	a.SetCountry("GB")
	return *a
}

func buildExamplePaylinkCardHolder() openapiclient.PaylinkCardHolder {
	h := openapiclient.NewPaylinkCardHolder()
	h.SetAcceptHeaders(exampleHolderAccept)
	addr := buildPaylinkCardHolderAddress()
	h.SetAddress(addr)
	h.SetCompany(exampleHolderCompany)
	h.SetEmail(exampleHolderEmail)
	h.SetFirstname(exampleHolderFirst)
	h.SetLastname(exampleHolderLast)
	h.SetMobileNo(exampleHolderMobile)
	h.SetRemoteAddr(exampleHolderRemote)
	h.SetTitle(exampleHolderTitle)
	h.SetUserAgent(exampleHolderUA)
	return *h
}

func TestNewPaylinkCardHolder(t *testing.T) {
	model := openapiclient.NewPaylinkCardHolder()
	require.NotNil(t, model)
	assert.False(t, model.HasEmail())
}

func TestNewPaylinkCardHolderWithDefaults(t *testing.T) {
	model := openapiclient.NewPaylinkCardHolderWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasFirstname())
}

func TestPaylinkCardHolderSetGetCycle(t *testing.T) {
	model := openapiclient.NewPaylinkCardHolder()
	model.SetAcceptHeaders(exampleHolderAccept)
	assert.True(t, model.HasAcceptHeaders())
	assert.Equal(t, exampleHolderAccept, model.GetAcceptHeaders())

	addr := buildPaylinkCardHolderAddress()
	model.SetAddress(addr)
	assert.True(t, model.HasAddress())
	assert.Equal(t, addr, model.GetAddress())

	model.SetCompany(exampleHolderCompany)
	assert.True(t, model.HasCompany())
	assert.Equal(t, exampleHolderCompany, model.GetCompany())

	model.SetEmail(exampleHolderEmail)
	assert.True(t, model.HasEmail())
	assert.Equal(t, exampleHolderEmail, model.GetEmail())

	model.SetFirstname(exampleHolderFirst)
	assert.True(t, model.HasFirstname())
	assert.Equal(t, exampleHolderFirst, model.GetFirstname())

	model.SetLastname(exampleHolderLast)
	assert.True(t, model.HasLastname())
	assert.Equal(t, exampleHolderLast, model.GetLastname())

	model.SetMobileNo(exampleHolderMobile)
	assert.True(t, model.HasMobileNo())
	assert.Equal(t, exampleHolderMobile, model.GetMobileNo())

	model.SetRemoteAddr(exampleHolderRemote)
	assert.True(t, model.HasRemoteAddr())
	assert.Equal(t, exampleHolderRemote, model.GetRemoteAddr())

	model.SetTitle(exampleHolderTitle)
	assert.True(t, model.HasTitle())
	assert.Equal(t, exampleHolderTitle, model.GetTitle())

	model.SetUserAgent(exampleHolderUA)
	assert.True(t, model.HasUserAgent())
	assert.Equal(t, exampleHolderUA, model.GetUserAgent())
}

func TestPaylinkCardHolderJSONRoundTrip(t *testing.T) {
	model := buildExamplePaylinkCardHolder()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.PaylinkCardHolder
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasRemoteAddr())
	assert.Equal(t, exampleHolderRemote, unmarshalled.GetRemoteAddr())
}

func TestPaylinkCardHolderToMap(t *testing.T) {
	model := buildExamplePaylinkCardHolder()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "accept_headers") {
		assert.Equal(t, &exampleHolderAccept, m["accept_headers"])
	}
	if assert.Contains(t, m, "company") {
		assert.Equal(t, &exampleHolderCompany, m["company"])
	}
}

func TestNullablePaylinkCardHolderGetSet(t *testing.T) {
	base := buildExamplePaylinkCardHolder()
	n := openapiclient.NullablePaylinkCardHolder{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkCardHolderUnset(t *testing.T) {
	base := buildExamplePaylinkCardHolder()
	n := openapiclient.NewNullablePaylinkCardHolder(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullablePaylinkCardHolderJSONRoundTrip(t *testing.T) {
	base := buildExamplePaylinkCardHolder()
	n := openapiclient.NewNullablePaylinkCardHolder(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullablePaylinkCardHolder
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleHolderFirst, newN.Get().GetFirstname())
	}
}
