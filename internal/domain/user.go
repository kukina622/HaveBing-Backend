package domain

import (
	"context"
	"time"

	"gopkg.in/guregu/null.v4"
)

type User struct {
	ID        uint      `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;"`
	Email     string    `gorm:"type:varchar(255) NOT NULL;unique;"`
	Password  string    `gorm:"type:varchar(255) NOT NULL;"`
	Name      string    `gorm:"type:mediumtext NOT NULL;"`
	Birthday  time.Time `gorm:"type:date;"`
	Phone     string    `gorm:"type:varchar(20);"`
	Available bool      `gorm:"type:bool NOT NULL;default:true"`
	RoleID    null.Int  `gorm:"type:bigint(20);default:null"`
	Role      *Role
}

type UserRepository interface {
	GetAll(ctx context.Context) ([]User, error)
	GetById(ctx context.Context, id uint) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Save(ctx context.Context, user *User) error
}

type UserUseCase interface {
	Login(ctx context.Context, email string, password string) (bool, *User, string)
	Register(ctx context.Context, user *User) error
	GetAll(ctx context.Context) ([]User, error)
	GetCurrentUser(ctx context.Context) (*User, error)
	ToggleUserAvailable(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User, oldPassword, newPassword string) error
}
