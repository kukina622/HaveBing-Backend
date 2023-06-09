package repository

import (
	"HaveBing-Backend/internal/domain"
	"context"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	var userList []domain.User
	err := u.db.Preload("Role").Find(&userList).Error
	return userList, err
}

func (u *UserRepository) GetById(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	err := u.db.Preload("Role").First(&user, id).Error
	return &user, err
}

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := u.db.Preload("Role").Where("email = ?", email).First(&user).Error
	return &user, err
}

func (u *UserRepository) Save(ctx context.Context, user *domain.User) error {
	return u.db.Save(user).Error
}
