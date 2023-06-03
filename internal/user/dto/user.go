package dto

import "HaveBing-Backend/internal/domain"

type UserLoginDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Birthday string `json:"birthday" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type UserResponseDTO struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Birthday  string `json:"birthday"`
	Phone     string `json:"phone"`
	Available bool   `json:"available"`
}

func NewUserResponse(user *domain.User) UserResponseDTO {
	return UserResponseDTO{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Birthday:  user.Birthday.Format("2006-01-02"),
		Phone:     user.Phone,
		Available: user.Available,
	}
}