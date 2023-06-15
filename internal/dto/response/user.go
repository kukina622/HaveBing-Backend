package response

import "HaveBing-Backend/internal/domain"

type UserResponseDTO struct {
	ID        uint         `json:"userId"`
	Email     string       `json:"email"`
	Name      string       `json:"name"`
	Birthday  string       `json:"birthday"`
	Phone     string       `json:"phone"`
	Available bool         `json:"available"`
	Role      *domain.Role `json:"role"`
}

type UserWithTokenResponseDTO struct {
	UserResponseDTO
	Token string `json:"token"`
}

func NewUserResponse[in *domain.User | []domain.User](user in) any {
	switch rawUser := any(user).(type) {
	case *domain.User:
		return UserResponseDTO{
			ID:        rawUser.ID,
			Email:     rawUser.Email,
			Name:      rawUser.Name,
			Birthday:  rawUser.Birthday.Format("2006-01-02"),
			Phone:     rawUser.Phone,
			Available: rawUser.Available,
			Role:      rawUser.Role,
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
				Role:      _user.Role,
			})
		}
		return userResponseDTO
	}
	return nil
}
