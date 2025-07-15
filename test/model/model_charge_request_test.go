package citypay

import (
	"encoding/json"
	"testing"

	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var exampleChargeAmount int32 = 1000
var exampleChargeAvsPolicy = "1"
var exampleChargeAgreement = "agr"
var exampleChargeCsc = "123"
var exampleChargeCscPolicy = "2"
var exampleChargeCurrency = "GBP"
var exampleChargeDuplicate = "1"
var exampleChargeIdentifier = "ID123"
var exampleChargeInitiation = "M"
var exampleChargeMatchAvsa = "Y"
var exampleChargeMerchantID int32 = 55
var exampleChargeTag = []string{"t1", "t2"}
var exampleChargeToken = "tok"
var exampleChargeTransInfo = "info"
var exampleChargeTransType = "SALE"

func buildExampleChargeRequest() *openapiclient.ChargeRequest {
	c := openapiclient.NewChargeRequest(exampleChargeAmount, exampleChargeIdentifier, exampleChargeMerchantID, exampleChargeToken)
	c.SetAvsPostcodePolicy(exampleChargeAvsPolicy)
	c.SetCardholderAgreement(exampleChargeAgreement)
	c.SetCsc(exampleChargeCsc)
	c.SetCscPolicy(exampleChargeCscPolicy)
	c.SetCurrency(exampleChargeCurrency)
	c.SetDuplicatePolicy(exampleChargeDuplicate)
	c.SetInitiation(exampleChargeInitiation)
	c.SetMatchAvsa(exampleChargeMatchAvsa)
	c.SetTag(exampleChargeTag)
	c.SetThreedsecure(buildExampleThreeDSecure())
	c.SetTransInfo(exampleChargeTransInfo)
	c.SetTransType(exampleChargeTransType)
	return c
}

func TestNewChargeRequest(t *testing.T) {
	model := openapiclient.NewChargeRequest(exampleChargeAmount, exampleChargeIdentifier, exampleChargeMerchantID, exampleChargeToken)
	require.NotNil(t, model)
	assert.Equal(t, exampleChargeAmount, model.GetAmount())
	assert.Equal(t, exampleChargeIdentifier, model.GetIdentifier())
	assert.Equal(t, exampleChargeMerchantID, model.GetMerchantid())
	assert.Equal(t, exampleChargeToken, model.GetToken())
	assert.False(t, model.HasAvsPostcodePolicy())
}

func TestNewChargeRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewChargeRequestWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, int32(0), model.GetAmount())
}

func TestChargeRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewChargeRequest(exampleChargeAmount, exampleChargeIdentifier, exampleChargeMerchantID, exampleChargeToken)
	model.SetAvsPostcodePolicy(exampleChargeAvsPolicy)
	assert.True(t, model.HasAvsPostcodePolicy())
	assert.Equal(t, exampleChargeAvsPolicy, model.GetAvsPostcodePolicy())

	model.SetCardholderAgreement(exampleChargeAgreement)
	assert.True(t, model.HasCardholderAgreement())
	assert.Equal(t, exampleChargeAgreement, model.GetCardholderAgreement())

	model.SetCsc(exampleChargeCsc)
	assert.True(t, model.HasCsc())
	assert.Equal(t, exampleChargeCsc, model.GetCsc())

	model.SetCscPolicy(exampleChargeCscPolicy)
	assert.True(t, model.HasCscPolicy())
	assert.Equal(t, exampleChargeCscPolicy, model.GetCscPolicy())

	model.SetCurrency(exampleChargeCurrency)
	assert.True(t, model.HasCurrency())
	assert.Equal(t, exampleChargeCurrency, model.GetCurrency())

	model.SetDuplicatePolicy(exampleChargeDuplicate)
	assert.True(t, model.HasDuplicatePolicy())
	assert.Equal(t, exampleChargeDuplicate, model.GetDuplicatePolicy())

	model.SetInitiation(exampleChargeInitiation)
	assert.True(t, model.HasInitiation())
	assert.Equal(t, exampleChargeInitiation, model.GetInitiation())

	model.SetMatchAvsa(exampleChargeMatchAvsa)
	assert.True(t, model.HasMatchAvsa())
	assert.Equal(t, exampleChargeMatchAvsa, model.GetMatchAvsa())

	model.SetTag(exampleChargeTag)
	assert.True(t, model.HasTag())
	assert.Equal(t, exampleChargeTag, model.GetTag())

	model.SetThreedsecure(buildExampleThreeDSecure())
	assert.True(t, model.HasThreedsecure())

	model.SetTransInfo(exampleChargeTransInfo)
	assert.True(t, model.HasTransInfo())
	assert.Equal(t, exampleChargeTransInfo, model.GetTransInfo())

	model.SetTransType(exampleChargeTransType)
	assert.True(t, model.HasTransType())
	assert.Equal(t, exampleChargeTransType, model.GetTransType())
}

func TestChargeRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleChargeRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.ChargeRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleChargeToken, unmarshalled.GetToken())
	assert.True(t, unmarshalled.HasTransType())
}

func TestChargeRequestToMap(t *testing.T) {
	model := buildExampleChargeRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "amount") {
		assert.Equal(t, exampleChargeAmount, m["amount"])
	}
	if assert.Contains(t, m, "identifier") {
		assert.Equal(t, exampleChargeIdentifier, m["identifier"])
	}
}

func TestNullableChargeRequestGetSet(t *testing.T) {
	base := buildExampleChargeRequest()
	n := openapiclient.NullableChargeRequest{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableChargeRequestUnset(t *testing.T) {
	base := buildExampleChargeRequest()
	n := openapiclient.NewNullableChargeRequest(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableChargeRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleChargeRequest()
	n := openapiclient.NewNullableChargeRequest(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableChargeRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleChargeIdentifier, newN.Get().GetIdentifier())
	}
}
