package mapper

import (
	"user-service/internal/model"
	userpb "user-service/internal/pb/user"
	"user-service/internal/util"
)

func UserProfileResponseMapper(user *model.User) *userpb.UserProfile {
	ID := util.UUIDToString(user.ID)

	var userAddresses []*userpb.AddressDetail
	if user.Addresses != nil && len(user.Addresses) > 0 {
		for _, address := range user.Addresses {
			userAddresses = append(userAddresses, AddressDetailResponseMapper(&address))
		}
	}

	return &userpb.UserProfile{
		Id:          ID,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		Addresses:   userAddresses,
	}
}

func UserProfieListResponseMapper(users []model.User) []userpb.UserProfile {
	profiles := make([]userpb.UserProfile, 0, len(users))
	for _, user := range users {
		profiles = append(profiles, *UserProfileResponseMapper(&user))
	}
	return profiles
}
