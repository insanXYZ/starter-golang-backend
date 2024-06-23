package repository

import (
	"backend/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{DB: DB}
}
