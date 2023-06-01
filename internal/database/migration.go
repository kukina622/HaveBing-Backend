package database

import (
	"HaveBing-Backend/internal/domain"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	return db.AutoMigrate(&domain.ProductCategory{}, &domain.User{})
}
