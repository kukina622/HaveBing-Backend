package usecase

import (
	"HaveBing-Backend/internal/domain"
	"HaveBing-Backend/internal/util/jwt"
	passwordUtil "HaveBing-Backend/internal/util/password"
	"context"
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
		token, _ = jwt.Sign(key, user.ID)
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

func (u *UserUseCase) Update(ctx context.Context, user *domain.User) error {
	return nil
}
