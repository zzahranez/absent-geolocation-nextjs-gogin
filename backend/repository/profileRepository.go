package repository

import (
	"gin/model"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	QueryFindUserAccountByEmail(employeeID uint) (*model.User, error)
}

type profilerepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profilerepository{
		db: db,
	}
}

func (r *profilerepository) QueryFindUserAccountByEmail(employeeID uint) (*model.User, error) {
	var user model.User
	if err := r.db.Select("name, email, role").Where("id = ?", employeeID).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
