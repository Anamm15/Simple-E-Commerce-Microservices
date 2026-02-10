package service

import (
	"context"

	"user-service/internal/dto"
	"user-service/internal/helper/mapper"
	userpb "user-service/internal/pb/user"
	"user-service/internal/repository"
	"user-service/pkg/util"

	"github.com/google/uuid"
)

type UserService interface {
	GetAll(ctx context.Context) ([]*userpb.UserProfile, error)
	GetByID(ctx context.Context, id string) (*userpb.UserProfile, error)
	Create(ctx context.Context, userID string, req dto.CreateProfileRequestDTO) (*userpb.CreateProfileResponse, error)
	Update(ctx context.Context, req dto.UpdateUserProfileRequestDTO) (*userpb.UserProfile, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetAll(ctx context.Context) ([]*userpb.UserProfile, error) {
	users, err := s.userRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	profiles := mapper.UserProfieListResponseMapper(users)
	return profiles, nil
}

func (s *userService) GetByID(ctx context.Context, id string) (*userpb.UserProfile, error) {
	idParsed, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetByID(ctx, idParsed)
	if err != nil {
		return nil, err
	}

	profile := mapper.UserProfileResponseMapper(user)
	return profile, nil
}

func (s *userService) Create(ctx context.Context, userIDReq string, req dto.CreateProfileRequestDTO) (*userpb.CreateProfileResponse, error) {
	userIDParsed, err := util.StringToUUID(userIDReq)
	if err != nil {
		return nil, err
	}

	user := req.ToModel()
	user.ID = userIDParsed

	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &userpb.CreateProfileResponse{
		Id:          user.ID.String(),
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}, nil
}

func (s *userService) Update(ctx context.Context, req dto.UpdateUserProfileRequestDTO) (*userpb.UserProfile, error) {
	userID, err := util.StringToUUID(req.UserID)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if req.FullName != "" {
		user.FullName = req.FullName
	}

	if req.PhoneNumber != "" {
		user.PhoneNumber = req.PhoneNumber
	}

	err = s.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	profile := mapper.UserProfileResponseMapper(user)
	return profile, nil
}
