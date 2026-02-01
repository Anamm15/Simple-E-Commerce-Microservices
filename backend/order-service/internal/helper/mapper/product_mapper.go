package mapper

import (
	"order-service/internal/dto"
	orderpb "order-service/internal/pb/order"
)

func MapProductCheckoutProtoToDTO(products []*orderpb.ProductCheckout) []dto.ProductCheckout {
	var result []dto.ProductCheckout
	for _, product := range products {
		result = append(result, dto.ProductCheckout{
			ProductID: product.ProductId,
			Quantity:  product.Quantity,
		})
	}
	return result
}
