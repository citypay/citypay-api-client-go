package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleRegCardNumber = "4111111111111111"
var exampleRegDefault = true
var exampleRegMonth int32 = 1
var exampleRegYear int32 = 25
var exampleRegName = "Jane Doe"

func buildExampleRegisterCard() openapiclient.RegisterCard {
	r := openapiclient.NewRegisterCard(exampleRegCardNumber, exampleRegMonth, exampleRegYear)
	r.SetDefault(exampleRegDefault)
	r.SetNameOnCard(exampleRegName)
	return *r
}

func TestNewRegisterCard(t *testing.T) {
	model := openapiclient.NewRegisterCard(exampleRegCardNumber, exampleRegMonth, exampleRegYear)
	require.NotNil(t, model)
	assert.Equal(t, exampleRegCardNumber, model.GetCardnumber())
	assert.Equal(t, exampleRegMonth, model.GetExpmonth())
	assert.Equal(t, exampleRegYear, model.GetExpyear())
	assert.False(t, model.HasDefault())
}

func TestNewRegisterCardWithDefaults(t *testing.T) {
	model := openapiclient.NewRegisterCardWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasNameOnCard())
}

func TestRegisterCardSetGetCycle(t *testing.T) {
	model := openapiclient.NewRegisterCard(exampleRegCardNumber, exampleRegMonth, exampleRegYear)
	model.SetDefault(exampleRegDefault)
	assert.True(t, model.HasDefault())
	assert.Equal(t, exampleRegDefault, model.GetDefault())

	model.SetNameOnCard(exampleRegName)
	assert.True(t, model.HasNameOnCard())
	assert.Equal(t, exampleRegName, model.GetNameOnCard())
}

func TestRegisterCardJSONRoundTrip(t *testing.T) {
	model := buildExampleRegisterCard()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.RegisterCard
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasDefault())
	assert.Equal(t, exampleRegName, unmarshalled.GetNameOnCard())
}

func TestRegisterCardToMap(t *testing.T) {
	model := buildExampleRegisterCard()
	m, err := model.ToMap()
	require.NoError(t, err)
	assert.Equal(t, exampleRegCardNumber, m["cardnumber"])
	if assert.Contains(t, m, "name_on_card") {
		assert.Equal(t, &exampleRegName, m["name_on_card"])
	}
}

func TestNullableRegisterCardGetSet(t *testing.T) {
	base := buildExampleRegisterCard()
	n := openapiclient.NullableRegisterCard{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullableRegisterCardUnset(t *testing.T) {
	base := buildExampleRegisterCard()
	n := openapiclient.NewNullableRegisterCard(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullableRegisterCardJSONRoundTrip(t *testing.T) {
	base := buildExampleRegisterCard()
	n := openapiclient.NewNullableRegisterCard(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullableRegisterCard
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleRegMonth, newN.Get().GetExpmonth())
	}
}
