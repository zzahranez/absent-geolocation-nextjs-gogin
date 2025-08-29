package repository

import (
	"gin/model"

	"gorm.io/gorm"
)

type UsersRepository interface {
	QueryGetByEmail(email string) (*model.User, error)
}

type usersrepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &usersrepository{
		db: db,
	}
}

func (r *usersrepository) QueryGetByEmail(email string) (*model.User, error) {
	var users model.User
	if err := r.db.Where("email = ?", email).First(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}
