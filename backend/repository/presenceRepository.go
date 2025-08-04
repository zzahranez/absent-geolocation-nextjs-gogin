package repository

import (
	"errors"
	"gin/model"
	"time"

	"gorm.io/gorm"
)

type PresenceRepository interface {
	QueryPresenceByEmail(email string) ([]model.Presences, error)

	QueryTodayPresenceByEmail(email string) (*model.Presences, error)
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
		Preload("PresentStatus").
		Where("user_id = ?", users.ID).
		Order("presence_date DESC").
		Limit(7).
		Find(&presence).Error; err != nil {
		return nil, err
	}

	return presence, nil
}

func (r *presencerepository) QueryTodayPresenceByEmail(email string) (*model.Presences, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	var presence model.Presences
	if err := r.db.
		Preload("PresentStatus").
		Where("user_id = ?", user.ID).
		Where("DATE(presence_date) = ?", time.Now().Format("2006-01-02")).
		First(&presence).Error; err != nil {
		// Return nil tanpa error kalau memang belum absen hari ini
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &presence, nil
}
