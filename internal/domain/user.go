package domain

import (
	"context"
	"time"
)

type User struct {
	ID        uint      `gorm:"type:int NOT NULL auto_increment;primary_key;"`
	Email     string    `gorm:"type:varchar(255) NOT NULL;unique;"`
	Password  string    `gorm:"type:varchar(255) NOT NULL;"`
	Name      string    `gorm:"type:mediumtext NOT NULL;"`
	Birthday  time.Time `gorm:"type:date;"`
	Phone     string    `gorm:"type:varchar(20);"`
	Available bool      `gorm:"type:bool NOT NULL;default:true"`
}

type UserRepository interface {
	GetAll(ctx context.Context) ([]User, error)
	GetById(ctx context.Context, id int) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user *User) error
	Save(ctx context.Context, user *User) error
}

type UserUseCase interface {
	Login(ctx context.Context, email string, password string) (*User, bool)
	Register(ctx context.Context, user *User) error
	GetAll(ctx context.Context) ([]User, error)
	DisableAccount(ctx context.Context, id int) error
	Update(ctx context.Context, user *User) error
}
