package usecase

import (
	"HaveBing-Backend/internal/domain"
	passwordUtil "HaveBing-Backend/internal/util/password"
	"context"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type UserUseCase struct {
	repo domain.UserRepository
}

func New(repo domain.UserRepository) domain.UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

func (u *UserUseCase) Login(ctx context.Context, email string, password string) (*domain.User, bool) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil || user == nil {
		return nil, false
	}
	salt := os.Getenv("SALT")
	success := passwordUtil.VerifyPassword(password, user.Password, salt) && user.Available
	return user, success
}

func (u *UserUseCase) Register(ctx context.Context, user *domain.User) error {
	salt := os.Getenv("SALT")
	hashedPassword, err := passwordUtil.HashPassword(user.Password, salt)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	fmt.Println(hashedPassword)
	return u.repo.Save(ctx, user)
}

func (u *UserUseCase) GetAll(ctx context.Context) ([]domain.User, error) {
	return nil, nil
}

func (u *UserUseCase) DisableAccount(ctx context.Context, id int) error {
	return nil
}

func (u *UserUseCase) Update(ctx context.Context, user *domain.User) error {
	return nil
}
