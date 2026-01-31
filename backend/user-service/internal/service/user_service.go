package service

import (
	"context"

	"user-service/internal/dto"
	"user-service/internal/helper/mapper"
	"user-service/internal/model"
	userpb "user-service/internal/pb/user"
	"user-service/internal/repository"
	"user-service/pkg/util"

	"github.com/google/uuid"
)

type UserService interface {
	GetAll(ctx context.Context) ([]userpb.UserProfile, error)
	GetByID(ctx context.Context, id string) (*userpb.UserProfile, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, req dto.UpdateUserProfileRequestDTO) (*userpb.UserProfile, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) GetAll(ctx context.Context) ([]userpb.UserProfile, error) {
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
		return &userpb.UserProfile{}, err
	}

	user, err := s.userRepo.GetByID(ctx, idParsed)
	if err != nil {
		return &userpb.UserProfile{}, err
	}

	profile := mapper.UserProfileResponseMapper(user)
	return profile, nil
}

func (s *userService) Create(ctx context.Context, user *model.User) error {
	return s.userRepo.Create(ctx, user)
}

func (s *userService) Update(ctx context.Context, req dto.UpdateUserProfileRequestDTO) (*userpb.UserProfile, error) {
	userID, err := util.StringToUUID(req.UserID)
	if err != nil {
		return &userpb.UserProfile{}, err
	}

	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return &userpb.UserProfile{}, err
	}

	if req.FullName != "" {
		user.FullName = req.FullName
	}

	if req.PhoneNumber != "" {
		user.PhoneNumber = req.PhoneNumber
	}

	err = s.userRepo.Update(ctx, user)
	if err != nil {
		return &userpb.UserProfile{}, err
	}

	profile := mapper.UserProfileResponseMapper(user)
	return profile, nil
}
