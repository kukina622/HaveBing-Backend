package request

type LoginUserRequestDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterUserRequestDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Birthday string `json:"birthday" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type ToggleUserAvailableRequestDTO struct {
	ID        uint `json:"userId"`
	Available bool `json:"available"`
}

type UpdateUserRequestDTO struct {
	Email       string `json:"email" binding:"required"`
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Birthday    string `json:"birthday" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
}
