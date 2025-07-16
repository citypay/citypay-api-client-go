package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleListClientName = "ExampleClient"
var exampleListClientID = "cli123"

func buildListMerchantsMerchant() openapiclient.Merchant {
	m := openapiclient.NewMerchant()
	m.SetName("Shop")
	m.SetMerchantid(99)
	return *m
}

func buildExampleListResponse() openapiclient.ListMerchantsResponse {
	l := openapiclient.NewListMerchantsResponse()
	l.SetClientName(exampleListClientName)
	l.SetClientid(exampleListClientID)
	l.SetMerchants([]openapiclient.Merchant{buildListMerchantsMerchant()})
	return *l
}

func TestNewListMerchantsResponse(t *testing.T) {
	model := openapiclient.NewListMerchantsResponse()
	require.NotNil(t, model)
	assert.False(t, model.HasClientName())
	assert.False(t, model.HasClientid())
	assert.False(t, model.HasMerchants())
}

func TestNewListMerchantsResponseWithDefaults(t *testing.T) {
	model := openapiclient.NewListMerchantsResponseWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasClientName())
	assert.False(t, model.HasClientid())
	assert.False(t, model.HasMerchants())
}

func TestListMerchantsResponseSetGetCycle(t *testing.T) {
	model := openapiclient.NewListMerchantsResponse()
	model.SetClientName(exampleListClientName)
	assert.True(t, model.HasClientName())
	assert.Equal(t, exampleListClientName, model.GetClientName())
	if val, ok := model.GetClientNameOk(); assert.True(t, ok) {
		assert.Equal(t, exampleListClientName, *val)
	}

	model.SetClientid(exampleListClientID)
	assert.True(t, model.HasClientid())
	assert.Equal(t, exampleListClientID, model.GetClientid())

	mer := buildListMerchantsMerchant()
	model.SetMerchants([]openapiclient.Merchant{mer})
	assert.True(t, model.HasMerchants())
	assert.Equal(t, 1, len(model.GetMerchants()))
}

func TestListMerchantsResponseJSONRoundTrip(t *testing.T) {
	model := buildExampleListResponse()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.ListMerchantsResponse
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasClientName())
	assert.Equal(t, exampleListClientID, unmarshalled.GetClientid())
	assert.True(t, unmarshalled.HasMerchants())
	assert.Equal(t, 1, len(unmarshalled.GetMerchants()))
}

func TestListMerchantsResponseToMap(t *testing.T) {
	model := buildExampleListResponse()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "client_name") {
		assert.Equal(t, &exampleListClientName, m["client_name"])
	}
	if assert.Contains(t, m, "clientid") {
		assert.Equal(t, &exampleListClientID, m["clientid"])
	}
	if assert.Contains(t, m, "merchants") {
		assert.Equal(t, model.Merchants, m["merchants"])
	}
}

func TestNullableListMerchantsResponseGetSet(t *testing.T) {
	base := buildExampleListResponse()
	n := openapiclient.NullableListMerchantsResponse{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableListMerchantsResponseUnset(t *testing.T) {
	base := buildExampleListResponse()
	n := openapiclient.NewNullableListMerchantsResponse(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableListMerchantsResponseJSONRoundTrip(t *testing.T) {
	base := buildExampleListResponse()
	n := openapiclient.NewNullableListMerchantsResponse(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableListMerchantsResponse
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.True(t, newN.Get().HasMerchants())
	}
}
