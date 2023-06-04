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

type UserAvailableDTO struct {
	ID        uint `json:"userId"`
	Available bool `json:"available"`
}

type UserUpdateDTO struct {
	Email       string `json:"email" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Birthday    string `json:"birthday" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
}

type UserResponseDTO struct {
	ID        uint   `json:"userId"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Birthday  string `json:"birthday"`
	Phone     string `json:"phone"`
	Available bool   `json:"available"`
}

type UserLoginResponseDTO struct {
	UserResponseDTO
	Token string `json:"token"`
}

func NewUserResponse(user any) any {
	switch rawUser := user.(type) {
	case *domain.User:
		return UserResponseDTO{
			ID:        rawUser.ID,
			Email:     rawUser.Email,
			Name:      rawUser.Name,
			Birthday:  rawUser.Birthday.Format("2006-01-02"),
			Phone:     rawUser.Phone,
			Available: rawUser.Available,
		}
	case []domain.User:
		var userResponseDTO []UserResponseDTO
		for _, _user := range rawUser {
			userResponseDTO = append(userResponseDTO, UserResponseDTO{
				ID:        _user.ID,
				Email:     _user.Email,
				Name:      _user.Name,
				Birthday:  _user.Birthday.Format("2006-01-02"),
				Phone:     _user.Phone,
				Available: _user.Available,
			})
		}
		return userResponseDTO
	}
	return nil
}
