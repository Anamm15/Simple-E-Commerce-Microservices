package mapper

import (
	"api-gateway/internal/dto"
	"api-gateway/internal/pb/order"
)

func ProductCheckoutRequestToProto(req []dto.ProductCheckoutRequest) []*orderpb.ProductCheckout {
	var result []*orderpb.ProductCheckout
	for _, v := range req {
		result = append(result, &orderpb.ProductCheckout{
			ProductId: v.ProductID,
			Quantity:  v.Quantity,
		})
	}
	return result
}

func AddressSnapshotToProto(req dto.AddressOrderSnapshot) *orderpb.AddressSnapshot {
	return &orderpb.AddressSnapshot{
		Street:        req.Street,
		City:          req.City,
		Province:      req.Province,
		ZipCode:       req.ZipCode,
		ReceiverName:  req.ReceiverName,
		ReceiverPhone: req.ReceiverPhone,
	}
}
