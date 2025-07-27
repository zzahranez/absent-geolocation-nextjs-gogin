package repository

import (
	"gin/model"

	"gorm.io/gorm"
)

type AuhtRepository interface {
	QueryLogin(email string) (*model.User, error)
}

type authrepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuhtRepository {
	return &authrepository{db: db}
}

func (r *authrepository) QueryLogin(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
