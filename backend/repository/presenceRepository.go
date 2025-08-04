package repository

import (
	"gin/model"

	"gorm.io/gorm"
)

type PresenceRepository interface {
	QueryPresenceByEmail(email string) ([]model.Presences, error)
}

type presencerepository struct {
	db *gorm.DB
}

func NewPresenceRepository(db *gorm.DB) PresenceRepository {
	return &presencerepository{
		db: db,
	}
}

func (r *presencerepository) QueryPresenceByEmail(email string) ([]model.Presences, error) {
	var users model.User
	if err := r.db.Where("email = ? ", email).First(&users).Error; err != nil {
		return nil, err
	}

	var presence []model.Presences
	if err := r.db.
		Preload("User").
		Where("user_id = ?", users.ID).
		Order("presence_date DESC").
		Find(&presence).Error; err != nil {
		return nil, err
	}

	return presence, nil
}
