package repositories

import (
	"context"
	"errors"
	errCommon "user-service/common/error"
	errConstants "user-service/constants/error"
	"user-service/domain/dto"
	"user-service/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	Register(context.Context, *dto.RegisterRequest) (*models.User, error)
	Update(context.Context, *dto.UpdateRequest, string) (*models.User, error)
	FindByUsername(context.Context, string) (*models.User, error)
	FindByEmail(context.Context, string) (*models.User, error)
	FindByUUID(context.Context, string) (*models.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) Register(ctx context.Context, req *dto.RegisterRequest) (*models.User, error) {
	user := models.User{
		UUID:        uuid.New(),
		Name:        req.Name,
		Username:    req.Username,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		RoleID:      req.RoleID,
	}

	if err := r.DB.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, errCommon.WrapError(errConstants.ErrSQLError)
	}

	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, req *dto.UpdateRequest, uuid string) (*models.User, error) {
	user := models.User{
		Name:        req.Name,
		Username:    req.Username,
		Password:    *req.Password,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
	}

	if err := r.DB.WithContext(ctx).Where("uuid = ?", uuid).Updates(&user).Error; err != nil {
		return nil, errCommon.WrapError(errConstants.ErrSQLError)
	}

	return &user, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User

	if err := r.DB.WithContext(ctx).Preload("Role").Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errConstants.ErrUserNotFound
		}

		return nil, errCommon.WrapError(errConstants.ErrSQLError)
	}

	return &user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	if err := r.DB.WithContext(ctx).Preload("Role").Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errConstants.ErrUserNotFound
		}

		return nil, errCommon.WrapError(errConstants.ErrSQLError)
	}

	return &user, nil
}

func (r *userRepository) FindByUUID(ctx context.Context, uuid string) (*models.User, error) {
	var user models.User

	if err := r.DB.WithContext(ctx).Preload("Role").Where("uuid = ?", uuid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errConstants.ErrUserNotFound
		}

		return nil, errCommon.WrapError(errConstants.ErrSQLError)
	}

	return &user, nil
}
