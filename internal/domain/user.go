package domain

import (
	"context"
	"time"
)

type User struct {
	ID        uint      `gorm:"type:int NOT NULL auto_increment;primary_key;" json:"id"`
	Email     string    `gorm:"type:varchar(255) NOT NULL;unique;" json:"email"`
	Password  string    `gorm:"type:varchar(255) NOT NULL;" json:"password"`
	Name      string    `gorm:"type:mediumtext NOT NULL;" json:"name"`
	Birthday  time.Time `gorm:"type:date;" json:"birthday"`
	Phone     string    `gorm:"type:varchar(20);" json:"phone"`
	Available bool      `gorm:"type:bool NOT NULL;default:true"`
}

type UserLogin struct {
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
}

type UserRepository interface {
	GetAll(ctx context.Context) ([]User, error)
	GetById(ctx context.Context, id int) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user *User) error
	Save(ctx context.Context, user *User) error
}

type UserUseCase interface {
	Login(ctx context.Context, email string, password string) bool
	Register(ctx context.Context, user *User) error
	GetAll(ctx context.Context) ([]User, error)
	DisableAccount(ctx context.Context, id int) error
	Update(ctx context.Context, user *User) error
}
