package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleBillAddressee = "Customer"
var exampleBillDescriptor = "Invoice"
var exampleBillDue = "2024-07-01"
var exampleBillMemo = "memo"
var exampleBillNextTo = "notify@example.com"
var exampleBillSmsTo = "+44111111111"

func buildExampleAttachment() openapiclient.PaylinkAttachmentRequest {
	a := openapiclient.NewPaylinkAttachmentRequest("file.pdf", "application/pdf")
	a.SetData("ZGF0YQ==")
	return *a
}

func buildExampleEmailPath() openapiclient.PaylinkEmailNotificationPath {
	p := openapiclient.NewPaylinkEmailNotificationPath([]string{exampleBillNextTo})
	return *p
}

func buildExampleSmsPath() openapiclient.PaylinkSMSNotificationPath {
	s := openapiclient.NewPaylinkSMSNotificationPath(exampleBillSmsTo)
	return *s
}

func buildExampleRequestModel() openapiclient.PaylinkTokenRequestModel {
	r := openapiclient.NewPaylinkTokenRequestModel(100, "id1", 1)
	return *r
}

func buildExamplePaylinkBillPaymentTokenRequest() openapiclient.PaylinkBillPaymentTokenRequest {
	b := openapiclient.NewPaylinkBillPaymentTokenRequest(buildExampleRequestModel())
	b.SetAddressee(exampleBillAddressee)
	b.SetAttachments([]openapiclient.PaylinkAttachmentRequest{buildExampleAttachment()})
	b.SetDescriptor(exampleBillDescriptor)
	b.SetDue(exampleBillDue)
	email := buildExampleEmailPath()
	b.SetEmailNotificationPath(email)
	b.SetMemo(exampleBillMemo)
	sms := buildExampleSmsPath()
	b.SetSmsNotificationPath(sms)
	return *b
}

func TestNewPaylinkBillPaymentTokenRequest(t *testing.T) {
	model := openapiclient.NewPaylinkBillPaymentTokenRequest(buildExampleRequestModel())
	require.NotNil(t, model)
	assert.Equal(t, buildExampleRequestModel(), model.GetRequest())
	assert.False(t, model.HasAddressee())
}

func TestNewPaylinkBillPaymentTokenRequestWithDefaults(t *testing.T) {
	model := openapiclient.NewPaylinkBillPaymentTokenRequestWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasAttachments())
}

func TestPaylinkBillPaymentTokenRequestSetGetCycle(t *testing.T) {
	model := openapiclient.NewPaylinkBillPaymentTokenRequest(buildExampleRequestModel())
	model.SetAddressee(exampleBillAddressee)
	assert.True(t, model.HasAddressee())
	assert.Equal(t, exampleBillAddressee, model.GetAddressee())

	model.SetAttachments([]openapiclient.PaylinkAttachmentRequest{buildExampleAttachment()})
	assert.True(t, model.HasAttachments())
	assert.Len(t, model.GetAttachments(), 1)

	model.SetDescriptor(exampleBillDescriptor)
	assert.True(t, model.HasDescriptor())
	assert.Equal(t, exampleBillDescriptor, model.GetDescriptor())

	model.SetDue(exampleBillDue)
	assert.True(t, model.HasDue())
	assert.Equal(t, exampleBillDue, model.GetDue())

	email := buildExampleEmailPath()
	model.SetEmailNotificationPath(email)
	assert.True(t, model.HasEmailNotificationPath())
	assert.Equal(t, email, model.GetEmailNotificationPath())

	model.SetMemo(exampleBillMemo)
	assert.True(t, model.HasMemo())
	assert.Equal(t, exampleBillMemo, model.GetMemo())

	sms := buildExampleSmsPath()
	model.SetSmsNotificationPath(sms)
	assert.True(t, model.HasSmsNotificationPath())
	assert.Equal(t, sms, model.GetSmsNotificationPath())
}

func TestPaylinkBillPaymentTokenRequestJSONRoundTrip(t *testing.T) {
	model := buildExamplePaylinkBillPaymentTokenRequest()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.PaylinkBillPaymentTokenRequest
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasDescriptor())
	assert.Equal(t, exampleBillDescriptor, unmarshalled.GetDescriptor())
}

func TestPaylinkBillPaymentTokenRequestToMap(t *testing.T) {
	model := buildExamplePaylinkBillPaymentTokenRequest()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "addressee") {
		assert.Equal(t, &exampleBillAddressee, m["addressee"])
	}
	if assert.Contains(t, m, "attachments") {
		assert.Len(t, m["attachments"], 1)
	}
	if assert.Contains(t, m, "descriptor") {
		assert.Equal(t, &exampleBillDescriptor, m["descriptor"])
	}
	if assert.Contains(t, m, "due") {
		assert.Equal(t, &exampleBillDue, m["due"])
	}
	if assert.Contains(t, m, "memo") {
		assert.Equal(t, &exampleBillMemo, m["memo"])
	}
}

func TestNullablePaylinkBillPaymentTokenRequestGetSet(t *testing.T) {
	base := buildExamplePaylinkBillPaymentTokenRequest()
	n := openapiclient.NullablePaylinkBillPaymentTokenRequest{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkBillPaymentTokenRequestUnset(t *testing.T) {
	base := buildExamplePaylinkBillPaymentTokenRequest()
	n := openapiclient.NewNullablePaylinkBillPaymentTokenRequest(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullablePaylinkBillPaymentTokenRequestJSONRoundTrip(t *testing.T) {
	base := buildExamplePaylinkBillPaymentTokenRequest()
	n := openapiclient.NewNullablePaylinkBillPaymentTokenRequest(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullablePaylinkBillPaymentTokenRequest
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.True(t, newN.Get().HasAddressee())
	}
}
