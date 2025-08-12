package citypay

import (
	"encoding/json"
	"testing"
	"time"

	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var exampleAuthRefAmount = "12.34"
var exampleAuthRefAmountValue int32 = 1234
var exampleAuthRefAtrn = "ATR123"
var exampleAuthRefAuthcode = "AUTHCODE"
var exampleAuthRefBatchno = "BN1"
var exampleAuthRefCurrency = "GBP"
var exampleAuthRefDatetime = time.Date(2023, time.January, 2, 15, 4, 5, 0, time.UTC)
var exampleAuthRefIdentifier = "ID42"
var exampleAuthRefMaskedpan = "411111******1111"
var exampleAuthRefMerchantid int32 = 100
var exampleAuthRefResult = "OK"
var exampleAuthRefTransStatus = "Approved"
var exampleAuthRefTransType = "SALE"
var exampleAuthRefTransno int32 = 99

func buildExampleAuthReference() *openapiclient.AuthReference {
	a := openapiclient.NewAuthReference()
	a.SetAmount(exampleAuthRefAmount)
	a.SetAmountValue(exampleAuthRefAmountValue)
	a.SetAtrn(exampleAuthRefAtrn)
	a.SetAuthcode(exampleAuthRefAuthcode)
	a.SetBatchno(exampleAuthRefBatchno)
	a.SetCurrency(exampleAuthRefCurrency)
	a.SetDatetime(exampleAuthRefDatetime)
	a.SetIdentifier(exampleAuthRefIdentifier)
	a.SetMaskedpan(exampleAuthRefMaskedpan)
	a.SetMerchantid(exampleAuthRefMerchantid)
	a.SetResult(exampleAuthRefResult)
	a.SetTransStatus(exampleAuthRefTransStatus)
	a.SetTransType(exampleAuthRefTransType)
	a.SetTransno(exampleAuthRefTransno)
	return a
}

func TestNewAuthReference(t *testing.T) {
	model := openapiclient.NewAuthReference()
	require.NotNil(t, model)
	assert.False(t, model.HasAmount())
	assert.False(t, model.HasAmountValue())
	assert.False(t, model.HasAtrn())
	assert.False(t, model.HasAuthcode())
	assert.False(t, model.HasBatchno())
	assert.False(t, model.HasCurrency())
	assert.False(t, model.HasDatetime())
	assert.False(t, model.HasIdentifier())
	assert.False(t, model.HasMaskedpan())
	assert.False(t, model.HasMerchantid())
	assert.False(t, model.HasResult())
	assert.False(t, model.HasTransStatus())
	assert.False(t, model.HasTransType())
	assert.False(t, model.HasTransno())
}

func TestNewAuthReferenceWithDefaults(t *testing.T) {
	model := openapiclient.NewAuthReferenceWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasAmount())
	assert.False(t, model.HasAmountValue())
}

func TestAuthReferenceSetGetCycle(t *testing.T) {
	model := openapiclient.NewAuthReference()
	model.SetAmount(exampleAuthRefAmount)
	assert.True(t, model.HasAmount())
	assert.Equal(t, exampleAuthRefAmount, model.GetAmount())

	model.SetAmountValue(exampleAuthRefAmountValue)
	assert.True(t, model.HasAmountValue())
	assert.Equal(t, exampleAuthRefAmountValue, model.GetAmountValue())

	model.SetAtrn(exampleAuthRefAtrn)
	assert.True(t, model.HasAtrn())
	assert.Equal(t, exampleAuthRefAtrn, model.GetAtrn())

	model.SetAuthcode(exampleAuthRefAuthcode)
	assert.True(t, model.HasAuthcode())
	assert.Equal(t, exampleAuthRefAuthcode, model.GetAuthcode())

	model.SetBatchno(exampleAuthRefBatchno)
	assert.True(t, model.HasBatchno())
	assert.Equal(t, exampleAuthRefBatchno, model.GetBatchno())

	model.SetCurrency(exampleAuthRefCurrency)
	assert.True(t, model.HasCurrency())
	assert.Equal(t, exampleAuthRefCurrency, model.GetCurrency())

	model.SetDatetime(exampleAuthRefDatetime)
	assert.True(t, model.HasDatetime())
	assert.Equal(t, exampleAuthRefDatetime, model.GetDatetime())

	model.SetIdentifier(exampleAuthRefIdentifier)
	assert.True(t, model.HasIdentifier())
	assert.Equal(t, exampleAuthRefIdentifier, model.GetIdentifier())

	model.SetMaskedpan(exampleAuthRefMaskedpan)
	assert.True(t, model.HasMaskedpan())
	assert.Equal(t, exampleAuthRefMaskedpan, model.GetMaskedpan())

	model.SetMerchantid(exampleAuthRefMerchantid)
	assert.True(t, model.HasMerchantid())
	assert.Equal(t, exampleAuthRefMerchantid, model.GetMerchantid())

	model.SetResult(exampleAuthRefResult)
	assert.True(t, model.HasResult())
	assert.Equal(t, exampleAuthRefResult, model.GetResult())

	model.SetTransStatus(exampleAuthRefTransStatus)
	assert.True(t, model.HasTransStatus())
	assert.Equal(t, exampleAuthRefTransStatus, model.GetTransStatus())

	model.SetTransType(exampleAuthRefTransType)
	assert.True(t, model.HasTransType())
	assert.Equal(t, exampleAuthRefTransType, model.GetTransType())

	model.SetTransno(exampleAuthRefTransno)
	assert.True(t, model.HasTransno())
	assert.Equal(t, exampleAuthRefTransno, model.GetTransno())
}

func TestAuthReferenceJSONRoundTrip(t *testing.T) {
	model := buildExampleAuthReference()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.AuthReference
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasAmount())
	assert.Equal(t, exampleAuthRefAmount, unmarshalled.GetAmount())
	assert.True(t, unmarshalled.HasTransno())
	assert.Equal(t, exampleAuthRefTransno, unmarshalled.GetTransno())
}

func TestAuthReferenceToMap(t *testing.T) {
	model := buildExampleAuthReference()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "amount") {
		assert.Equal(t, &exampleAuthRefAmount, m["amount"])
	}
	if assert.Contains(t, m, "transno") {
		assert.Equal(t, &exampleAuthRefTransno, m["transno"])
	}
}

func TestNullableAuthReferenceGetSet(t *testing.T) {
	base := buildExampleAuthReference()
	n := openapiclient.NullableAuthReference{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableAuthReferenceUnset(t *testing.T) {
	base := buildExampleAuthReference()
	n := openapiclient.NewNullableAuthReference(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableAuthReferenceJSONRoundTrip(t *testing.T) {
	base := buildExampleAuthReference()
	n := openapiclient.NewNullableAuthReference(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableAuthReference
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleAuthRefAtrn, newN.Get().GetAtrn())
	}
}
