package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleDPAmount int32 = 9999
var exampleDPAvsPolicy = "1"
var exampleDPCard = "4111111111111111"
var exampleDPCsc = "123"
var exampleDPCscPolicy = "2"
var exampleDPCurrency = "GBP"
var exampleDPDupPolicy = "1"
var exampleDPExpMonth int32 = 10
var exampleDPExpYear int32 = 2026
var exampleDPIdentifier = "id123"
var exampleDPMac = "MAC"
var exampleDPMatchAvsa = "0"
var exampleDPName = "John Doe"
var exampleDPNonce = "ABCDEF"
var exampleDPRedirectFail = "https://fail"
var exampleDPRedirectSuccess = "https://ok"
var exampleDPTag = []string{"one", "two"}
var exampleDPTransInfo = "info"
var exampleDPTransType = "SALE"

func buildDirectPostContact() openapiclient.ContactDetails {
	c := openapiclient.NewContactDetails()
	c.SetAddress1("addr")
	c.SetEmail("c@example.com")
	return *c
}

func buildExampleThreeDS() openapiclient.ThreeDSecure {
	t := openapiclient.NewThreeDSecure()
	t.SetAcceptHeaders("text/html")
	return *t
}

func buildExampleDirectPostRequest() *openapiclient.DirectPostRequest {
	d := openapiclient.NewDirectPostRequest(exampleDPAmount, exampleDPCard, exampleDPExpMonth, exampleDPExpYear, exampleDPIdentifier, exampleDPMac)
	d.SetAvsPostcodePolicy(exampleDPAvsPolicy)
	d.SetBillTo(buildDirectPostContact())
	d.SetCsc(exampleDPCsc)
	d.SetCscPolicy(exampleDPCscPolicy)
	d.SetCurrency(exampleDPCurrency)
	d.SetDuplicatePolicy(exampleDPDupPolicy)
	d.SetMatchAvsa(exampleDPMatchAvsa)
	d.SetNameOnCard(exampleDPName)
	d.SetNonce(exampleDPNonce)
	d.SetRedirectFailure(exampleDPRedirectFail)
	d.SetRedirectSuccess(exampleDPRedirectSuccess)
	d.SetShipTo(buildDirectPostContact())
	d.SetTag(exampleDPTag)
	d.SetThreedsecure(buildExampleThreeDS())
	d.SetTransInfo(exampleDPTransInfo)
	d.SetTransType(exampleDPTransType)
	return d
}

func TestNewDirectPostRequest(t *testing.T) {
	model := openapiclient.NewDirectPostRequest(exampleDPAmount, exampleDPCard, exampleDPExpMonth, exampleDPExpYear, exampleDPIdentifier, exampleDPMac)
	require.NotNil(t, model)
	assert.Equal(t, exampleDPAmount, model.GetAmount())
	assert.Equal(t, exampleDPCard, model.GetCardnumber())
	assert.False(t, model.HasCsc())
}

func TestNewDirectPostRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewDirectPostRequestWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, int32(0), model.GetAmount())
	assert.Equal(t, "", model.GetCardnumber())
}

func TestDirectPostRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewDirectPostRequest(exampleDPAmount, exampleDPCard, exampleDPExpMonth, exampleDPExpYear, exampleDPIdentifier, exampleDPMac)
	model.SetAvsPostcodePolicy(exampleDPAvsPolicy)
	assert.True(t, model.HasAvsPostcodePolicy())
	assert.Equal(t, exampleDPAvsPolicy, model.GetAvsPostcodePolicy())

	ct := buildDirectPostContact()
	model.SetBillTo(ct)
	assert.True(t, model.HasBillTo())
	assert.Equal(t, ct, model.GetBillTo())

	model.SetCsc(exampleDPCsc)
	assert.True(t, model.HasCsc())
	assert.Equal(t, exampleDPCsc, model.GetCsc())

	model.SetCscPolicy(exampleDPCscPolicy)
	assert.True(t, model.HasCscPolicy())
	assert.Equal(t, exampleDPCscPolicy, model.GetCscPolicy())

	model.SetCurrency(exampleDPCurrency)
	assert.True(t, model.HasCurrency())
	assert.Equal(t, exampleDPCurrency, model.GetCurrency())

	model.SetDuplicatePolicy(exampleDPDupPolicy)
	assert.True(t, model.HasDuplicatePolicy())
	assert.Equal(t, exampleDPDupPolicy, model.GetDuplicatePolicy())

	model.SetMatchAvsa(exampleDPMatchAvsa)
	assert.True(t, model.HasMatchAvsa())
	assert.Equal(t, exampleDPMatchAvsa, model.GetMatchAvsa())

	model.SetNameOnCard(exampleDPName)
	assert.True(t, model.HasNameOnCard())
	assert.Equal(t, exampleDPName, model.GetNameOnCard())

	model.SetNonce(exampleDPNonce)
	assert.True(t, model.HasNonce())
	assert.Equal(t, exampleDPNonce, model.GetNonce())

	model.SetRedirectFailure(exampleDPRedirectFail)
	assert.True(t, model.HasRedirectFailure())
	assert.Equal(t, exampleDPRedirectFail, model.GetRedirectFailure())

	model.SetRedirectSuccess(exampleDPRedirectSuccess)
	assert.True(t, model.HasRedirectSuccess())
	assert.Equal(t, exampleDPRedirectSuccess, model.GetRedirectSuccess())

	ship := buildDirectPostContact()
	model.SetShipTo(ship)
	assert.True(t, model.HasShipTo())
	assert.Equal(t, ship, model.GetShipTo())

	model.SetTag(exampleDPTag)
	assert.True(t, model.HasTag())
	assert.Equal(t, exampleDPTag, model.GetTag())

	td := buildExampleThreeDS()
	model.SetThreedsecure(td)
	assert.True(t, model.HasThreedsecure())
	assert.Equal(t, td, model.GetThreedsecure())

	model.SetTransInfo(exampleDPTransInfo)
	assert.True(t, model.HasTransInfo())
	assert.Equal(t, exampleDPTransInfo, model.GetTransInfo())

	model.SetTransType(exampleDPTransType)
	assert.True(t, model.HasTransType())
	assert.Equal(t, exampleDPTransType, model.GetTransType())
}

func TestDirectPostRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleDirectPostRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.DirectPostRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleDPAmount, unmarshalled.GetAmount())
	assert.True(t, unmarshalled.HasCsc())
	assert.True(t, unmarshalled.HasThreedsecure())
}

func TestDirectPostRequestToMap(t *testing.T) {
	model := buildExampleDirectPostRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "amount") {
		assert.Equal(t, exampleDPAmount, m["amount"])
	}
	if assert.Contains(t, m, "avs_postcode_policy") {
		assert.Equal(t, &exampleDPAvsPolicy, m["avs_postcode_policy"])
	}
	if assert.Contains(t, m, "tag") {
		assert.Equal(t, exampleDPTag, m["tag"])
	}
}

func TestNullableDirectPostRequestGetSet(t *testing.T) {
	base := buildExampleDirectPostRequest()
	n := openapiclient.NullableDirectPostRequest{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableDirectPostRequestUnset(t *testing.T) {
	base := buildExampleDirectPostRequest()
	n := openapiclient.NewNullableDirectPostRequest(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableDirectPostRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleDirectPostRequest()
	n := openapiclient.NewNullableDirectPostRequest(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableDirectPostRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleDPIdentifier, newN.Get().GetIdentifier())
	}
}
