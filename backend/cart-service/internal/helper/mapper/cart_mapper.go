package mapper

import (
	"cart-service/internal/model"
	cartpb "cart-service/internal/pb/cart"
	productpb "cart-service/internal/pb/product"
	"cart-service/pkg/util"
)

func MapCartResponse(cartItem model.CartItem, productItem *productpb.Product) *cartpb.CartItem {
	productID := util.UUIDToString(cartItem.ProductID)
	cartItemID := util.UUIDToString(cartItem.ID)
	return &cartpb.CartItem{
		Id:           cartItemID,
		ProductId:    productID,
		Quantity:     cartItem.Quantity,
		ProductName:  productItem.Name,
		ProductImage: productItem.Thumbnail,
		ProductPrice: productItem.Price,
		Subtotal:     int64(cartItem.Quantity) * productItem.Price,
	}
}

func MapCartListResponse(cartItems []model.CartItem, products map[string]*productpb.Product) []*cartpb.CartItem {
	var cartItemsResponse []*cartpb.CartItem

	for _, cartItem := range cartItems {
		cartItemsResponse = append(cartItemsResponse, MapCartResponse(cartItem, products[cartItem.ProductID.String()]))
	}
	return cartItemsResponse
}
