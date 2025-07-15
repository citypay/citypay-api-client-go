package citypay

import (
	"encoding/json"
	openapiclient "github.com/citypay/citypay-api-client-go/citypay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var exampleCartCoupon = "NEW10"
var exampleCartMode int32 = 2
var exampleCartProdDesc = "Product"
var exampleCartProdInfo = "Info"
var exampleCartShipping int32 = 50
var exampleCartTax int32 = 20

func buildExampleCartItem() openapiclient.PaylinkCartItemModel {
	i := openapiclient.NewPaylinkCartItemModel()
	i.SetLabel("Item")
	return *i
}

func buildExamplePaylinkCart() openapiclient.PaylinkCart {
	c := openapiclient.NewPaylinkCart()
	c.SetContents([]openapiclient.PaylinkCartItemModel{buildExampleCartItem()})
	c.SetCoupon(exampleCartCoupon)
	c.SetMode(exampleCartMode)
	c.SetProductDescription(exampleCartProdDesc)
	c.SetProductInformation(exampleCartProdInfo)
	c.SetShipping(exampleCartShipping)
	c.SetTax(exampleCartTax)
	return *c
}

func TestNewPaylinkCart(t *testing.T) {
	model := openapiclient.NewPaylinkCart()
	require.NotNil(t, model)
	assert.False(t, model.HasCoupon())
}

func TestNewPaylinkCartWithDefaults(t *testing.T) {
	model := openapiclient.NewPaylinkCartWithDefaults()
	require.NotNil(t, model)
	assert.False(t, model.HasContents())
}

func TestPaylinkCartSetGetCycle(t *testing.T) {
	model := openapiclient.NewPaylinkCart()
	model.SetContents([]openapiclient.PaylinkCartItemModel{buildExampleCartItem()})
	assert.True(t, model.HasContents())
	assert.Len(t, model.GetContents(), 1)

	model.SetCoupon(exampleCartCoupon)
	assert.True(t, model.HasCoupon())
	assert.Equal(t, exampleCartCoupon, model.GetCoupon())

	model.SetMode(exampleCartMode)
	assert.True(t, model.HasMode())
	assert.Equal(t, exampleCartMode, model.GetMode())

	model.SetProductDescription(exampleCartProdDesc)
	assert.True(t, model.HasProductDescription())
	assert.Equal(t, exampleCartProdDesc, model.GetProductDescription())

	model.SetProductInformation(exampleCartProdInfo)
	assert.True(t, model.HasProductInformation())
	assert.Equal(t, exampleCartProdInfo, model.GetProductInformation())

	model.SetShipping(exampleCartShipping)
	assert.True(t, model.HasShipping())
	assert.Equal(t, exampleCartShipping, model.GetShipping())

	model.SetTax(exampleCartTax)
	assert.True(t, model.HasTax())
	assert.Equal(t, exampleCartTax, model.GetTax())
}

func TestPaylinkCartJSONRoundTrip(t *testing.T) {
	model := buildExamplePaylinkCart()
	data, err := json.Marshal(model)
	require.NoError(t, err)

	var unmarshalled openapiclient.PaylinkCart
	err = json.Unmarshal(data, &unmarshalled)
	require.NoError(t, err)
	assert.True(t, unmarshalled.HasMode())
	assert.Equal(t, exampleCartMode, unmarshalled.GetMode())
}

func TestPaylinkCartToMap(t *testing.T) {
	model := buildExamplePaylinkCart()
	m, err := model.ToMap()
	require.NoError(t, err)
	if assert.Contains(t, m, "coupon") {
		assert.Equal(t, &exampleCartCoupon, m["coupon"])
	}
	if assert.Contains(t, m, "tax") {
		assert.Equal(t, &exampleCartTax, m["tax"])
	}
}

func TestNullablePaylinkCartGetSet(t *testing.T) {
	base := buildExamplePaylinkCart()
	n := openapiclient.NullablePaylinkCart{}
	n.Set(&base)
	require.True(t, n.IsSet())
	assert.Equal(t, &base, n.Get())
}

func TestNullablePaylinkCartUnset(t *testing.T) {
	base := buildExamplePaylinkCart()
	n := openapiclient.NewNullablePaylinkCart(&base)
	require.True(t, n.IsSet())
	n.Unset()
	assert.False(t, n.IsSet())
	assert.Nil(t, n.Get())
}

func TestNullablePaylinkCartJSONRoundTrip(t *testing.T) {
	base := buildExamplePaylinkCart()
	n := openapiclient.NewNullablePaylinkCart(&base)
	data, err := json.Marshal(n)
	require.NoError(t, err)

	var newN openapiclient.NullablePaylinkCart
	err = json.Unmarshal(data, &newN)
	require.NoError(t, err)
	assert.True(t, newN.IsSet())
	if assert.NotNil(t, newN.Get()) {
		assert.Equal(t, exampleCartShipping, newN.Get().GetShipping())
	}
}
