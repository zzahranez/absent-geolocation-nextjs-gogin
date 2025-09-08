package repository

import (
	"gin/model"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	QueryFindUserAccountByEmail(employeeID uint) (*model.User, error)
	QueryFindBiodataUserByEmail(employeeID uint) (*model.Profile, error)
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

func (r *profilerepository) QueryFindBiodataUserByEmail(employeeID uint) (*model.Profile, error) {
	var profile model.Profile
	if err := r.db.Preload("User").Where("user_id = ?", employeeID).First(&profile).Error; err != nil {
		return nil, err
	}

	return &profile, nil
}
