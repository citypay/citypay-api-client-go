package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleProcessDate = "2024-01-01"
var exampleProcessID int32 = 7
var exampleProcessClient = "cli"

func buildExampleBatchTransaction() openapiclient.BatchTransaction {
	b := openapiclient.NewBatchTransaction("acc", 10)
	b.SetIdentifier("id")
	b.SetMerchantid(1)
	return *b
}

func buildExampleProcessBatchRequest() openapiclient.ProcessBatchRequest {
	p := openapiclient.NewProcessBatchRequest(exampleProcessDate, exampleProcessID, []openapiclient.BatchTransaction{buildExampleBatchTransaction()})
	p.SetClientAccountId(exampleProcessClient)
	return *p
}

func TestNewProcessBatchRequest(t *testing.T) {
	model := openapiclient.NewProcessBatchRequest(exampleProcessDate, exampleProcessID, []openapiclient.BatchTransaction{})
	require.NotNil(t, model)
	assert.Equal(t, exampleProcessDate, model.GetBatchDate())
	assert.Equal(t, exampleProcessID, model.GetBatchId())
	assert.False(t, model.HasClientAccountId())
}

func TestNewProcessBatchRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewProcessBatchRequestWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasTransactions())
}

func TestProcessBatchRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewProcessBatchRequest(exampleProcessDate, exampleProcessID, []openapiclient.BatchTransaction{})
	model.SetClientAccountId(exampleProcessClient)
	assert.True(t, model.HasClientAccountId())
	assert.Equal(t, exampleProcessClient, model.GetClientAccountId())

	tx := buildExampleBatchTransaction()
	model.SetTransactions([]openapiclient.BatchTransaction{tx})
	assert.Len(t, model.GetTransactions(), 1)
}

func TestProcessBatchRequestJSONRoundTrip(t *testing.T) {
	model := buildExampleProcessBatchRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.ProcessBatchRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasClientAccountId())
	assert.Equal(t, exampleProcessClient, unmarshalled.GetClientAccountId())
}

func TestProcessBatchRequestToMap(t *testing.T) {
	model := buildExampleProcessBatchRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	assert.Equal(t, exampleProcessDate, m["batch_date"])
	if assert.Contains(t, m, "client_account_id") {
		assert.Equal(t, &exampleProcessClient, m["client_account_id"])
	}
}

func TestNullableProcessBatchRequestGetSet(t *testing.T) {
	base := buildExampleProcessBatchRequest()
	n := openapiclient.NullableProcessBatchRequest{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableProcessBatchRequestUnset(t *testing.T) {
	base := buildExampleProcessBatchRequest()
	n := openapiclient.NewNullableProcessBatchRequest(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableProcessBatchRequestJSONRoundTrip(t *testing.T) {
	base := buildExampleProcessBatchRequest()
	n := openapiclient.NewNullableProcessBatchRequest(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableProcessBatchRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleProcessID, newN.Get().GetBatchId())
	}
}
