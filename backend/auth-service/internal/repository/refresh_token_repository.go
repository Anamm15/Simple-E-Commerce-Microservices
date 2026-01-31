package repository

import (
	"context"

	"auth-service/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	GetByTokenHash(ctx context.Context, token string) (*model.RefreshToken, error)
	Create(ctx context.Context, refreshToken *model.RefreshToken) error
	Save(ctx context.Context, refreshToken *model.RefreshToken) error
	RevokeToken(ctx context.Context, token string) error
	RotateToken(ctx context.Context, oldTokenID uuid.UUID, newToken *model.RefreshToken) error
	RevokeFamily(ctx context.Context, familyID uuid.UUID) error
	Delete(ctx context.Context, userID uuid.UUID, token string) error
}

type refreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) RefreshTokenRepository {
	return &refreshTokenRepository{db: db}
}

func (r *refreshTokenRepository) GetByTokenHash(ctx context.Context, token string) (*model.RefreshToken, error) {
	refreshToken := &model.RefreshToken{}
	if err := r.db.WithContext(ctx).
		Where("token = ?", token).
		First(refreshToken).Error; err != nil {
		return nil, err
	}
	return refreshToken, nil
}

func (r *refreshTokenRepository) Create(ctx context.Context, RefreshToken *model.RefreshToken) error {
	if err := r.db.WithContext(ctx).Create(RefreshToken).Error; err != nil {
		return err
	}
	return nil
}

func (r *refreshTokenRepository) Save(ctx context.Context, token *model.RefreshToken) error {
	if err := r.db.WithContext(ctx).Save(token).Error; err != nil {
		return err
	}
	return nil
}

func (r *refreshTokenRepository) RevokeToken(ctx context.Context, token string) error {
	return r.db.
		WithContext(ctx).
		Model(&model.RefreshToken{}).
		Where("token = ?", token).
		Update("is_revoked", true).
		Error
}

func (r *refreshTokenRepository) RotateToken(ctx context.Context, oldTokenID uuid.UUID, newToken *model.RefreshToken) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.RefreshToken{}).Where("id = ?", oldTokenID).Update("is_used", true).Error; err != nil {
			return err
		}
		if err := tx.Create(newToken).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *refreshTokenRepository) RevokeFamily(ctx context.Context, familyID uuid.UUID) error {
	return r.db.Model(&model.RefreshToken{}).
		Where("family_id = ?", familyID).
		Update("is_revoked", true).Error
}

func (r *refreshTokenRepository) Delete(ctx context.Context, userID uuid.UUID, token string) error {
	if err := r.db.WithContext(ctx).
		Delete(&model.RefreshToken{}, "user_id = ? AND token = ?", userID, token).Error; err != nil {
		return err
	}
	return nil
}
