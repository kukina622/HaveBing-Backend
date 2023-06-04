package usecase

import (
	"HaveBing-Backend/internal/domain"
	jwtUtil "HaveBing-Backend/internal/util/jwt"
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

func (u *UserUseCase) Login(ctx context.Context, email string, password string) (bool, *domain.User, string) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil || user == nil {
		return false, nil, ""
	}
	salt := os.Getenv("SALT")
	success := passwordUtil.VerifyPassword(password, user.Password, salt) && user.Available

	var token string
	if success {
		key := os.Getenv("SECRET_KEY")
		token, err = jwtUtil.Sign(key, user.ID)
		success = err == nil
	}

	return success, user, token
}

func (u *UserUseCase) Register(ctx context.Context, user *domain.User) error {
	salt := os.Getenv("SALT")
	hashedPassword, err := passwordUtil.HashPassword(user.Password, salt)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return u.repo.Save(ctx, user)
}

func (u *UserUseCase) GetAll(ctx context.Context) ([]domain.User, error) {
	return u.repo.GetAll(ctx)
}

func (u *UserUseCase) ToggleUserAvailable(ctx context.Context, user *domain.User) error {
	targetUser, err := u.repo.GetById(ctx, user.ID)
	if err != nil {
		return err
	}
	targetUser.Available = user.Available
	return u.repo.Save(ctx, targetUser)
}

func (u *UserUseCase) Update(ctx context.Context, user *domain.User, oldPassword, newPassword string) error {
	targetUser, err := u.repo.GetById(ctx, user.ID)
	if err != nil {
		return err
	}

	salt := os.Getenv("SALT")
	if !passwordUtil.VerifyPassword(oldPassword, targetUser.Password, salt) {
		return fmt.Errorf("old password is incorrect")
	}

	hashedPassword, err := passwordUtil.HashPassword(newPassword, salt)
	if err != nil {
		return err
	}

	targetUser.Email = user.Email
	targetUser.Password = hashedPassword
	targetUser.Name = user.Name
	targetUser.Birthday = user.Birthday
	targetUser.Phone = user.Phone
	return u.repo.Save(ctx, targetUser)
}
