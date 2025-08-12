package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleDomain = []string{"example.com"}
var exampleLive = true
var exampleDomainMerchantID int32 = 50

func buildExampleDomainKeyRequest() *openapiclient.DomainKeyRequest {
	d := openapiclient.NewDomainKeyRequest(exampleDomain, exampleDomainMerchantID)
	d.SetLive(exampleLive)
	return d
}

func TestNewDomainKeyRequest(t *testing.T) {
	model := openapiclient.NewDomainKeyRequest(exampleDomain, exampleDomainMerchantID)
	require.NotNil(t, model)
	assert.Equal(t, exampleDomain, model.GetDomain())
	assert.Equal(t, exampleDomainMerchantID, model.GetMerchantid())
	assert.False(t, model.HasLive())
}

func TestNewDomainKeyRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewDomainKeyRequestWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, 0, len(model.GetDomain()))
	assert.Equal(t, int32(0), model.GetMerchantid())
}

func TestDomainKeyRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewDomainKeyRequest(exampleDomain, exampleDomainMerchantID)
	model.SetDomain([]string{"new.com"})
	assert.Equal(t, []string{"new.com"}, model.GetDomain())

	model.SetLive(exampleLive)
	assert.True(t, model.HasLive())
	assert.Equal(t, exampleLive, model.GetLive())
	if val, ok := model.GetLiveOk(); assert.True(t, ok) {
		assert.Equal(t, exampleLive, *val)
	}

	model.SetMerchantid(99)
	assert.Equal(t, int32(99), model.GetMerchantid())
}

func TestDomainKeyRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleDomainKeyRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.DomainKeyRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleDomainMerchantID, unmarshalled.GetMerchantid())
	assert.True(t, unmarshalled.HasLive())
}

func TestDomainKeyRequestToMap(t *testing.T) {
	model := buildExampleDomainKeyRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "domain") {
		assert.Equal(t, exampleDomain, m["domain"])
	}
	if assert.Contains(t, m, "live") {
		assert.Equal(t, &exampleLive, m["live"])
	}
}

func TestNullableDomainKeyRequestGetSet(t *testing.T) {
	base := buildExampleDomainKeyRequest()
	n := openapiclient.NullableDomainKeyRequest{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableDomainKeyRequestUnset(t *testing.T) {
	base := buildExampleDomainKeyRequest()
	n := openapiclient.NewNullableDomainKeyRequest(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableDomainKeyRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleDomainKeyRequest()
	n := openapiclient.NewNullableDomainKeyRequest(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableDomainKeyRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleDomainMerchantID, newN.Get().GetMerchantid())
	}
}
