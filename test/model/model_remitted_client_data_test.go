package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var exampleRemittedClientID = "cid"
var exampleRemittedDate = "2024-01-02"
var exampleRemittedNet int32 = 50

func buildExampleMerchantBatchResponse() openapiclient.MerchantBatchResponse {
	m := openapiclient.NewMerchantBatchResponse()
	m.SetBatchId(1)
	return *m
}

func buildExampleRemittanceDataSlice() []openapiclient.RemittanceData {
	r := openapiclient.NewRemittanceData()
	r.SetDateCreated(time.Now())
	return []openapiclient.RemittanceData{*r}
}

func buildExampleRemittedClientData() openapiclient.RemittedClientData {
	d := openapiclient.NewRemittedClientData([]openapiclient.MerchantBatchResponse{buildExampleMerchantBatchResponse()}, buildExampleRemittanceDataSlice())
	d.SetClientid(exampleRemittedClientID)
	d.SetDate(&exampleRemittedDate)
	d.SetNetAmount(exampleRemittedNet)
	return *d
}

func TestNewRemittedClientData(t *testing.T) {
	model := openapiclient.NewRemittedClientData([]openapiclient.MerchantBatchResponse{}, []openapiclient.RemittanceData{})
	require.NotNil(t, model)
	assert.False(t, model.HasClientid())
}

func TestNewRemittedClientDataWithDefaults(t *testing.T) {
	model := openapiclient.NewRemittedClientDataWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasRemittances())
}

func TestRemittedClientDataSetGetCycle(t *testing.T) {
	model := openapiclient.NewRemittedClientData([]openapiclient.MerchantBatchResponse{}, []openapiclient.RemittanceData{})
	model.SetClientid(exampleRemittedClientID)
	assert.True(t, model.HasClientid())
	assert.Equal(t, exampleRemittedClientID, model.GetClientid())

	model.SetDate(&exampleRemittedDate)
	assert.True(t, model.HasDate())
	assert.Equal(t, exampleRemittedDate, model.GetDate())

	now := time.Now()
	model.SetDateCreated(now)
	assert.True(t, model.HasDateCreated())
	assert.Equal(t, now, model.GetDateCreated())

	model.SetNetAmount(exampleRemittedNet)
	assert.True(t, model.HasNetAmount())
	assert.Equal(t, exampleRemittedNet, model.GetNetAmount())
}

func TestRemittedClientDataJSONRoundTrip(t *testing.T) {
	model := buildExampleRemittedClientData()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.RemittedClientData
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasClientid())
	assert.Equal(t, exampleRemittedClientID, unmarshalled.GetClientid())
}

func TestRemittedClientDataToMap(t *testing.T) {
	model := buildExampleRemittedClientData()
	m, err := model.ToMap()
	require.NoError(t, err)
	assert.Contains(t, m, "batches")
	if assert.Contains(t, m, "clientid") {
		assert.Equal(t, &exampleRemittedClientID, m["clientid"])
	}
}

func TestNullableRemittedClientDataGetSet(t *testing.T) {
	base := buildExampleRemittedClientData()
	n := openapiclient.NullableRemittedClientData{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableRemittedClientDataUnset(t *testing.T) {
	base := buildExampleRemittedClientData()
	n := openapiclient.NewNullableRemittedClientData(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableRemittedClientDataJSONRoundTrip(t *testing.T) {
	base := buildExampleRemittedClientData()
	n := openapiclient.NewNullableRemittedClientData(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableRemittedClientData
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleRemittedNet, newN.Get().GetNetAmount())
	}
}
