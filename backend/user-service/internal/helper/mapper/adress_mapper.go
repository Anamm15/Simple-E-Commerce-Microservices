package mapper

import (
	"user-service/internal/model"
	userpb "user-service/internal/pb/user"
	"user-service/internal/util"
)

func AddressDetailResponseMapper(address *model.Address) *userpb.AddressDetail {
	ID := util.UUIDToString(address.ID)
	UserID := util.UUIDToString(address.UserID)

	return &userpb.AddressDetail{
		Id:        ID,
		UserId:    UserID,
		Street:    address.Street,
		City:      address.City,
		Province:  address.Province,
		ZipCode:   address.ZipCode,
		IsPrimary: address.IsPrimary,
	}
}

func AddressResponseMapper(address *model.Address) *userpb.AddressResponse {
	ID := util.UUIDToString(address.ID)

	return &userpb.AddressResponse{
		Id:        ID,
		IsPrimary: address.IsPrimary,
	}
}

func AddressDetailListResponseMapper(addresses []model.Address) []userpb.AddressDetail {
	var userAddresses []userpb.AddressDetail

	for _, address := range addresses {
		userAddresses = append(userAddresses, *AddressDetailResponseMapper(&address))
	}

	return userAddresses
}
