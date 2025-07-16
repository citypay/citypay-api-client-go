package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var exampleRespDomain = []string{"example.com"}
var exampleRespDate = time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
var exampleRespDomainKey = "KEY"
var exampleDomainRespLive = true
var exampleRespMerchant int32 = 55

func buildExampleDomainKeyResponse() *openapiclient.DomainKeyResponse {
	d := openapiclient.NewDomainKeyResponse(exampleRespDomain, exampleRespMerchant)
	d.SetDateCreated(exampleRespDate)
	d.SetDomainKey(exampleRespDomainKey)
	d.SetLive(exampleDomainRespLive)
	return d
}

func TestNewDomainKeyResponse(t *testing.T) {
	model := openapiclient.NewDomainKeyResponse(exampleRespDomain, exampleRespMerchant)
	require.NotNil(t, model)
	assert.Equal(t, exampleRespDomain, model.GetDomain())
	assert.Equal(t, exampleRespMerchant, model.GetMerchantid())
	assert.False(t, model.HasDomainKey())
}

func TestNewDomainKeyResponseWithDefaults(t *testing.T) {
	model := openapiclient.NewDomainKeyResponseWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, 0, len(model.GetDomain()))
	assert.Equal(t, int32(0), model.GetMerchantid())
}

func TestDomainKeyResponseSetGetCycle(t *testing.T) {
	model := openapiclient.NewDomainKeyResponse(exampleRespDomain, exampleRespMerchant)
	model.SetDateCreated(exampleRespDate)
	assert.True(t, model.HasDateCreated())
	assert.Equal(t, exampleRespDate, model.GetDateCreated())

	model.SetDomain(exampleRespDomain)
	assert.Equal(t, exampleRespDomain, model.GetDomain())

	model.SetDomainKey(exampleRespDomainKey)
	assert.True(t, model.HasDomainKey())
	assert.Equal(t, exampleRespDomainKey, model.GetDomainKey())

	model.SetLive(exampleDomainRespLive)
	assert.True(t, model.HasLive())
	assert.Equal(t, exampleDomainRespLive, model.GetLive())

	model.SetMerchantid(99)
	assert.Equal(t, int32(99), model.GetMerchantid())
}

func TestDomainKeyResponseJSONRoundTrip(t *testing.T) {
	model := buildExampleDomainKeyResponse()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.DomainKeyResponse
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasDomainKey())
	assert.Equal(t, exampleRespDomainKey, unmarshalled.GetDomainKey())
}

func TestDomainKeyResponseToMap(t *testing.T) {
	model := buildExampleDomainKeyResponse()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "domain_key") {
		assert.Equal(t, &exampleRespDomainKey, m["domain_key"])
	}
	if assert.Contains(t, m, "live") {
		assert.Equal(t, &exampleDomainRespLive, m["live"])
	}
}

func TestNullableDomainKeyResponseGetSet(t *testing.T) {
	base := buildExampleDomainKeyResponse()
	n := openapiclient.NullableDomainKeyResponse{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableDomainKeyResponseUnset(t *testing.T) {
	base := buildExampleDomainKeyResponse()
	n := openapiclient.NewNullableDomainKeyResponse(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableDomainKeyResponseJSONRoundTrip(t *testing.T) {
	base := buildExampleDomainKeyResponse()
	n := openapiclient.NewNullableDomainKeyResponse(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableDomainKeyResponse
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleRespDomainKey, newN.Get().GetDomainKey())
	}
}
