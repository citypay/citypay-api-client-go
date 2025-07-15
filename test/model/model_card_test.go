package citypay

import (
    "encoding/json"
    "testing"
    "time"

    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

var exampleCardCommercial = true
var exampleCardCorporate = true
var exampleCardCountry = "US"
var exampleCardCredit = true
var exampleCardCurrency = "USD"
var exampleCardDebit = true
var exampleCardDescription = "Gold"
var exampleCardEu = false
var exampleCardId = "card1"
var exampleCardStatus = "ACTIVE"
var exampleCardDate = time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
var exampleCardDefault = true
var exampleCardExpMonth int32 = 12
var exampleCardExpYear int32 = 30
var exampleCardLabel = "Main"
var exampleCardLabel2 = "Main 12/30"
var exampleCardLast4 = "1234"
var exampleCardName = "John Smith"
var exampleCardScheme = "VISA"
var exampleCardToken = "tok123"

func buildExampleCard() *openapiclient.Card {
    c := openapiclient.NewCard()
    c.SetBinCommercial(exampleCardCommercial)
    c.SetBinCorporate(exampleCardCorporate)
    c.SetBinCountryIssued(exampleCardCountry)
    c.SetBinCredit(exampleCardCredit)
    c.SetBinCurrency(exampleCardCurrency)
    c.SetBinDebit(exampleCardDebit)
    c.SetBinDescription(exampleCardDescription)
    c.SetBinEu(exampleCardEu)
    c.SetCardId(exampleCardId)
    c.SetCardStatus(exampleCardStatus)
    c.SetDateCreated(exampleCardDate)
    c.SetDefault(exampleCardDefault)
    c.SetExpmonth(exampleCardExpMonth)
    c.SetExpyear(exampleCardExpYear)
    c.SetLabel(exampleCardLabel)
    c.SetLabel2(exampleCardLabel2)
    c.SetLast4digits(exampleCardLast4)
    c.SetNameOnCard(exampleCardName)
    c.SetScheme(exampleCardScheme)
    c.SetToken(exampleCardToken)
    return c
}

func TestNewCard(t *testing.T) {
    model := openapiclient.NewCard()
    require.NotNil(t, model)
    assert.False(t, model.HasCardId())
}

func TestNewCardWithDefaults(t *testing.T) {
    model := openapiclient.NewCardWithDefaults()
    require.NotNil(t, model)
    assert.False(t, model.HasCardStatus())
}

func TestCardSetGetCycle(t *testing.T) {
    model := openapiclient.NewCard()
    model.SetBinCommercial(exampleCardCommercial)
    assert.True(t, model.HasBinCommercial())
    assert.Equal(t, exampleCardCommercial, model.GetBinCommercial())

    model.SetBinCorporate(exampleCardCorporate)
    assert.True(t, model.HasBinCorporate())
    assert.Equal(t, exampleCardCorporate, model.GetBinCorporate())

    model.SetBinCountryIssued(exampleCardCountry)
    assert.True(t, model.HasBinCountryIssued())
    assert.Equal(t, exampleCardCountry, model.GetBinCountryIssued())

    model.SetBinCredit(exampleCardCredit)
    assert.True(t, model.HasBinCredit())
    assert.Equal(t, exampleCardCredit, model.GetBinCredit())

    model.SetBinCurrency(exampleCardCurrency)
    assert.True(t, model.HasBinCurrency())
    assert.Equal(t, exampleCardCurrency, model.GetBinCurrency())

    model.SetBinDebit(exampleCardDebit)
    assert.True(t, model.HasBinDebit())
    assert.Equal(t, exampleCardDebit, model.GetBinDebit())

    model.SetBinDescription(exampleCardDescription)
    assert.True(t, model.HasBinDescription())
    assert.Equal(t, exampleCardDescription, model.GetBinDescription())

    model.SetBinEu(exampleCardEu)
    assert.True(t, model.HasBinEu())
    assert.Equal(t, exampleCardEu, model.GetBinEu())

    model.SetCardId(exampleCardId)
    assert.True(t, model.HasCardId())
    assert.Equal(t, exampleCardId, model.GetCardId())

    model.SetCardStatus(exampleCardStatus)
    assert.True(t, model.HasCardStatus())
    assert.Equal(t, exampleCardStatus, model.GetCardStatus())

    model.SetDateCreated(exampleCardDate)
    assert.True(t, model.HasDateCreated())
    assert.Equal(t, exampleCardDate, model.GetDateCreated())

    model.SetDefault(exampleCardDefault)
    assert.True(t, model.HasDefault())
    assert.Equal(t, exampleCardDefault, model.GetDefault())

    model.SetExpmonth(exampleCardExpMonth)
    assert.True(t, model.HasExpmonth())
    assert.Equal(t, exampleCardExpMonth, model.GetExpmonth())

    model.SetExpyear(exampleCardExpYear)
    assert.True(t, model.HasExpyear())
    assert.Equal(t, exampleCardExpYear, model.GetExpyear())

    model.SetLabel(exampleCardLabel)
    assert.True(t, model.HasLabel())
    assert.Equal(t, exampleCardLabel, model.GetLabel())

    model.SetLabel2(exampleCardLabel2)
    assert.True(t, model.HasLabel2())
    assert.Equal(t, exampleCardLabel2, model.GetLabel2())

    model.SetLast4digits(exampleCardLast4)
    assert.True(t, model.HasLast4digits())
    assert.Equal(t, exampleCardLast4, model.GetLast4digits())

    model.SetNameOnCard(exampleCardName)
    assert.True(t, model.HasNameOnCard())
    assert.Equal(t, exampleCardName, model.GetNameOnCard())

    model.SetScheme(exampleCardScheme)
    assert.True(t, model.HasScheme())
    assert.Equal(t, exampleCardScheme, model.GetScheme())

    model.SetToken(exampleCardToken)
    assert.True(t, model.HasToken())
    assert.Equal(t, exampleCardToken, model.GetToken())
}

func TestCardJSONRoundTrip(t *testing.T) {
    model := buildExampleCard()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.Card
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.Equal(t, exampleCardId, unmarshalled.GetCardId())
    assert.Equal(t, exampleCardScheme, unmarshalled.GetScheme())
}

func TestCardToMap(t *testing.T) {
    model := buildExampleCard()
    m, err := model.ToMap()
    require.NoError(t, err)
    if assert.Contains(t, m, "card_id") {
        assert.Equal(t, &exampleCardId, m["card_id"])
    }
    if assert.Contains(t, m, "scheme") {
        assert.Equal(t, &exampleCardScheme, m["scheme"])
    }
}

func TestNullableCardGetSet(t *testing.T) {
    base := buildExampleCard()
    n := openapiclient.NullableCard{}
    n.Set(base)
    require.True(t, n.IsSet())
    assert.Equal(t, base, n.Get())
}

func TestNullableCardUnset(t *testing.T) {
    base := buildExampleCard()
    n := openapiclient.NewNullableCard(base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullableCardJSONRoundTrip(t *testing.T) {
    base := buildExampleCard()
    n := openapiclient.NewNullableCard(base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullableCard
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, exampleCardId, newN.Get().GetCardId())
    }
}

