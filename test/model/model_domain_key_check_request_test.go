package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleDomainKey = "ABCDEF"

func buildExampleDomainKeyCheckRequest() *openapiclient.DomainKeyCheckRequest {
	return openapiclient.NewDomainKeyCheckRequest(exampleDomainKey)
}

func TestNewDomainKeyCheckRequest(t *testing.T) {
	model := openapiclient.NewDomainKeyCheckRequest(exampleDomainKey)
	require.NotNil(t, model)
	assert.Equal(t, exampleDomainKey, model.GetDomainKey())
}

func TestNewDomainKeyCheckRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewDomainKeyCheckRequestWithDefaults()
	require.NotNil(t, model)
	assert.Equal(t, "", model.GetDomainKey())
}

func TestDomainKeyCheckRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewDomainKeyCheckRequest(exampleDomainKey)
	model.SetDomainKey("XYZ")
	assert.Equal(t, "XYZ", model.GetDomainKey())
	if val, ok := model.GetDomainKeyOk(); assert.True(t, ok) {
		assert.Equal(t, "XYZ", *val)
	}
}

func TestDomainKeyCheckRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleDomainKeyCheckRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.DomainKeyCheckRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.Equal(t, exampleDomainKey, unmarshalled.GetDomainKey())
}

func TestDomainKeyCheckRequestToMap(t *testing.T) {
	model := buildExampleDomainKeyCheckRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "domain_key") {
		assert.Equal(t, exampleDomainKey, m["domain_key"])
	}
}

func TestNullableDomainKeyCheckRequestGetSet(t *testing.T) {
	base := buildExampleDomainKeyCheckRequest()
	n := openapiclient.NullableDomainKeyCheckRequest{}
	n.Set(base)
	require.True(t, n.IsSet())
	assert.Equal(t, base, n.Get())
}

func TestNullableDomainKeyCheckRequestUnset(t *testing.T) {
	base := buildExampleDomainKeyCheckRequest()
	n := openapiclient.NewNullableDomainKeyCheckRequest(base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableDomainKeyCheckRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleDomainKeyCheckRequest()
	n := openapiclient.NewNullableDomainKeyCheckRequest(base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableDomainKeyCheckRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleDomainKey, newN.Get().GetDomainKey())
	}
}
