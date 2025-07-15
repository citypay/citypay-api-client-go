package citypay

import (
    "encoding/json"
    "testing"
    "time"

    openapiclient "github.com/citypay/citypay-api-client-go"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

var exampleCHAAccountId = "acc123"
var exampleCHADateCreated = time.Date(2023, time.February, 1, 9, 0, 0, 0, time.UTC)
var exampleCHADefaultCardId = "cardA"
var exampleCHADefaultCardIndex int32 = 0
var exampleCHALastModified = time.Date(2023, time.March, 1, 10, 0, 0, 0, time.UTC)
var exampleCHAStatus = "ACTIVE"
var exampleCHAUniqueId = "unique1"

func buildExampleCardHolderAccount() *openapiclient.CardHolderAccount {
    c := openapiclient.NewCardHolderAccount(exampleCHAAccountId, buildExampleContact())
    c.SetCards([]openapiclient.Card{*buildExampleCard()})
    c.SetDateCreated(exampleCHADateCreated)
    c.SetDefaultCardId(exampleCHADefaultCardId)
    c.SetDefaultCardIndex(exampleCHADefaultCardIndex)
    c.SetLastModified(exampleCHALastModified)
    c.SetStatus(exampleCHAStatus)
    c.SetUniqueId(exampleCHAUniqueId)
    return c
}

func TestNewCardHolderAccount(t *testing.T) {
    contact := buildExampleContact()
    model := openapiclient.NewCardHolderAccount(exampleCHAAccountId, contact)
    require.NotNil(t, model)
    assert.Equal(t, exampleCHAAccountId, model.GetAccountId())
    assert.Equal(t, contact, model.GetContact())
    assert.False(t, model.HasCards())
}

func TestNewCardHolderAccountWithDefaults(t *testing.T) {
    model := openapiclient.NewCardHolderAccountWithDefaults()
    require.NotNil(t, model)
    assert.Equal(t, "", model.GetAccountId())
}

func TestCardHolderAccountSetGetCycle(t *testing.T) {
    contact := buildExampleContact()
    model := openapiclient.NewCardHolderAccount(exampleCHAAccountId, contact)
    model.SetCards([]openapiclient.Card{*buildExampleCard()})
    assert.True(t, model.HasCards())
    assert.Equal(t, 1, len(model.GetCards()))

    model.SetContact(contact)
    assert.Equal(t, contact, model.GetContact())

    model.SetDateCreated(exampleCHADateCreated)
    assert.True(t, model.HasDateCreated())
    assert.Equal(t, exampleCHADateCreated, model.GetDateCreated())

    model.SetDefaultCardId(exampleCHADefaultCardId)
    assert.True(t, model.HasDefaultCardId())
    assert.Equal(t, exampleCHADefaultCardId, model.GetDefaultCardId())

    model.SetDefaultCardIndex(exampleCHADefaultCardIndex)
    assert.True(t, model.HasDefaultCardIndex())
    assert.Equal(t, exampleCHADefaultCardIndex, model.GetDefaultCardIndex())

    model.SetLastModified(exampleCHALastModified)
    assert.True(t, model.HasLastModified())
    assert.Equal(t, exampleCHALastModified, model.GetLastModified())

    model.SetStatus(exampleCHAStatus)
    assert.True(t, model.HasStatus())
    assert.Equal(t, exampleCHAStatus, model.GetStatus())

    model.SetUniqueId(exampleCHAUniqueId)
    assert.True(t, model.HasUniqueId())
    assert.Equal(t, exampleCHAUniqueId, model.GetUniqueId())
}

func TestCardHolderAccountJSONRoundTrip(t *testing.T) {
    model := buildExampleCardHolderAccount()
    data, err := json.Marshal(model)
    require.NoError(t, err)

    var unmarshalled openapiclient.CardHolderAccount
    err = json.Unmarshal(data, &unmarshalled)
    require.NoError(t, err)
    assert.Equal(t, exampleCHAAccountId, unmarshalled.GetAccountId())
    assert.True(t, unmarshalled.HasStatus())
}

func TestCardHolderAccountToMap(t *testing.T) {
    model := buildExampleCardHolderAccount()
    m, err := model.ToMap()
    require.NoError(t, err)
    if assert.Contains(t, m, "account_id") {
        assert.Equal(t, exampleCHAAccountId, m["account_id"])
    }
    if assert.Contains(t, m, "default_card_id") {
        assert.Equal(t, &exampleCHADefaultCardId, m["default_card_id"])
    }
}

func TestNullableCardHolderAccountGetSet(t *testing.T) {
    base := buildExampleCardHolderAccount()
    n := openapiclient.NullableCardHolderAccount{}
    n.Set(base)
    require.True(t, n.IsSet())
    assert.Equal(t, base, n.Get())
}

func TestNullableCardHolderAccountUnset(t *testing.T) {
    base := buildExampleCardHolderAccount()
    n := openapiclient.NewNullableCardHolderAccount(base)
    require.True(t, n.IsSet())
    n.Unset()
    assert.False(t, n.IsSet())
    assert.Nil(t, n.Get())
}

func TestNullableCardHolderAccountJSONRoundTrip(t *testing.T) {
    base := buildExampleCardHolderAccount()
    n := openapiclient.NewNullableCardHolderAccount(base)
    data, err := json.Marshal(n)
    require.NoError(t, err)

    var newN openapiclient.NullableCardHolderAccount
    err = json.Unmarshal(data, &newN)
    require.NoError(t, err)
    assert.True(t, newN.IsSet())
    if assert.NotNil(t, newN.Get()) {
        assert.Equal(t, exampleCHAUniqueId, newN.Get().GetUniqueId())
    }
}

