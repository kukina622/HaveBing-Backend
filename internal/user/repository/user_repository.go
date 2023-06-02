package repository

import (
	"HaveBing-Backend/internal/domain"
	"context"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	var userList []domain.User
	err := u.db.Find(&userList).Error
	return userList, err
}

func (u *UserRepository) GetById(ctx context.Context, id int) (*domain.User, error) {
	var user domain.User
	err := u.db.First(&user, id).Error
	return &user, err
}

func (u *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := u.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (u *UserRepository) Update(ctx context.Context, user *domain.User) error {
	return nil
}

func (u *UserRepository) Save(ctx context.Context, user *domain.User) error {
	return nil
}
