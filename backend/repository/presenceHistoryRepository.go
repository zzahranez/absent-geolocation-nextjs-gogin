package repository

import (
	"errors"
	"gin/model"
	"time"

	"gorm.io/gorm"
)

type PresenceHistoryRepository interface {
	QueryPresenceByEmail(email string, month, year int) ([]model.Presences, error)
	QueryTodayPresenceByEmail(email string) (*model.Presences, error)
}

type presencehistoryrepository struct {
	db *gorm.DB
}

func NewPresenceHistoryRepository(db *gorm.DB) PresenceHistoryRepository {
	return &presencehistoryrepository{
		db: db,
	}
}

func (r *presencehistoryrepository) QueryPresenceByEmail(email string, month, year int) ([]model.Presences, error) {
	var users model.User
	if err := r.db.Where("email = ? ", email).First(&users).Error; err != nil {
		return nil, err
	}

	var presence []model.Presences
	if err := r.db.
		Preload("User").
		Preload("PresenceStatus").
		Where("user_id = ? AND MONTH(presence_date) = ? AND YEAR(presence_date) = ?", users.ID, month, year).
		Order("presence_date DESC").
		Limit(7).
		Find(&presence).Error; err != nil {
		return nil, err
	}

	return presence, nil
}

func (r *presencehistoryrepository) QueryTodayPresenceByEmail(email string) (*model.Presences, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	var presence model.Presences
	if err := r.db.
		Preload("PresenceStatus").
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
