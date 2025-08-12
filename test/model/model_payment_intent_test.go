package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleIntentAmount int32 = 500
var exampleIntentIdentifier = "pid"
var exampleIntentAvsPolicy = "1"
var exampleIntentCsc = "123"
var exampleIntentCscPolicy = "2"
var exampleIntentCurrency = "GBP"
var exampleIntentDupPolicy = "1"
var exampleIntentMatchAvsa = "Y"
var exampleIntentTransInfo = "info"
var exampleIntentTransType = "SALE"
var exampleIntentTag = []string{"t1", "t2"}

func buildIntentContact() openapiclient.ContactDetails {
	c := openapiclient.NewContactDetails()
	c.SetAddress1("street")
	c.SetEmail("x@example.com")
	return *c
}

func buildExamplePaymentIntent() openapiclient.PaymentIntent {
	p := openapiclient.NewPaymentIntent(exampleIntentAmount, exampleIntentIdentifier)
	p.SetAvsPostcodePolicy(exampleIntentAvsPolicy)
	p.SetBillTo(buildIntentContact())
	p.SetCsc(exampleIntentCsc)
	p.SetCscPolicy(exampleIntentCscPolicy)
	p.SetCurrency(exampleIntentCurrency)
	p.SetDuplicatePolicy(exampleIntentDupPolicy)
	p.SetMatchAvsa(exampleIntentMatchAvsa)
	p.SetShipTo(buildIntentContact())
	p.SetTag(exampleIntentTag)
	p.SetTransInfo(exampleIntentTransInfo)
	p.SetTransType(exampleIntentTransType)
	return *p
}

func TestNewPaymentIntent(t *testing.T) {
	model := openapiclient.NewPaymentIntent(exampleIntentAmount, exampleIntentIdentifier)
	require.NotNil(t, model)
	assert.Equal(t, exampleIntentAmount, model.GetAmount())
	assert.Equal(t, exampleIntentIdentifier, model.GetIdentifier())
	assert.False(t, model.HasCurrency())
}

func TestNewPaymentIntentWithDefaults(t *testing.T) {
	model := openapiclient.NewPaymentIntentWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, int32(0), model.GetAmount())
	assert.False(t, model.HasTag())
}

func TestPaymentIntentSetGetCycle(t *testing.T) {
	model := openapiclient.NewPaymentIntent(exampleIntentAmount, exampleIntentIdentifier)
	model.SetAvsPostcodePolicy(exampleIntentAvsPolicy)
	assert.True(t, model.HasAvsPostcodePolicy())
	assert.Equal(t, exampleIntentAvsPolicy, model.GetAvsPostcodePolicy())

	ct := buildIntentContact()
	model.SetBillTo(ct)
	assert.True(t, model.HasBillTo())
	assert.Equal(t, ct, model.GetBillTo())

	model.SetCsc(exampleIntentCsc)
	assert.True(t, model.HasCsc())
	assert.Equal(t, exampleIntentCsc, model.GetCsc())

	model.SetCscPolicy(exampleIntentCscPolicy)
	assert.True(t, model.HasCscPolicy())
	assert.Equal(t, exampleIntentCscPolicy, model.GetCscPolicy())

	model.SetCurrency(exampleIntentCurrency)
	assert.True(t, model.HasCurrency())
	assert.Equal(t, exampleIntentCurrency, model.GetCurrency())

	model.SetDuplicatePolicy(exampleIntentDupPolicy)
	assert.True(t, model.HasDuplicatePolicy())
	assert.Equal(t, exampleIntentDupPolicy, model.GetDuplicatePolicy())

	model.SetMatchAvsa(exampleIntentMatchAvsa)
	assert.True(t, model.HasMatchAvsa())
	assert.Equal(t, exampleIntentMatchAvsa, model.GetMatchAvsa())

	ship := buildIntentContact()
	model.SetShipTo(ship)
	assert.True(t, model.HasShipTo())
	assert.Equal(t, ship, model.GetShipTo())

	model.SetTag(exampleIntentTag)
	assert.True(t, model.HasTag())
	assert.Equal(t, exampleIntentTag, model.GetTag())

	model.SetTransInfo(exampleIntentTransInfo)
	assert.True(t, model.HasTransInfo())
	assert.Equal(t, exampleIntentTransInfo, model.GetTransInfo())

	model.SetTransType(exampleIntentTransType)
	assert.True(t, model.HasTransType())
	assert.Equal(t, exampleIntentTransType, model.GetTransType())
}

func TestPaymentIntentJSONRoundTrip(t *testing.T) {
	model := buildExamplePaymentIntent()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.PaymentIntent
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasBillTo())
	assert.Equal(t, exampleIntentCurrency, unmarshalled.GetCurrency())
}

func TestPaymentIntentToMap(t *testing.T) {
	model := buildExamplePaymentIntent()
	m, err := model.ToMap()
	require.NoError(t, err)
	assert.Equal(t, exampleIntentAmount, m["amount"])
	if assert.Contains(t, m, "tag") {
		assert.Equal(t, exampleIntentTag, m["tag"])
	}
}

func TestNullablePaymentIntentGetSet(t *testing.T) {
	base := buildExamplePaymentIntent()
	n := openapiclient.NullablePaymentIntent{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullablePaymentIntentUnset(t *testing.T) {
	base := buildExamplePaymentIntent()
	n := openapiclient.NewNullablePaymentIntent(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullablePaymentIntentJSONRoundTrip(t *testing.T) {
	base := buildExamplePaymentIntent()
	n := openapiclient.NewNullablePaymentIntent(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullablePaymentIntent
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleIntentDupPolicy, newN.Get().GetDuplicatePolicy())
	}
}
