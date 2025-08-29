package repository

import (
	"errors"
	"gin/model"
	"time"

	"gorm.io/gorm"
)

type PresenceRepository interface {
	QueryCreatePresence(presence *model.Presences) error
	QueryCreateAbsentAfterNoon(presence *model.Presences) error
	QueryCheckPresenceToday(employeeID uint) (*model.Presences, error)
}

type presencerepository struct {
	db *gorm.DB
}

func NewPresenceRepository(db *gorm.DB) PresenceRepository {
	return &presencerepository{db: db}
}

// Absen pagi
func (r *presencerepository) QueryCreatePresence(presence *model.Presences) error {
	return r.db.Create(presence).Error
}

// Absen sore
func (r *presencerepository) QueryCreateAbsentAfterNoon(presence *model.Presences) error {
	timeOut := ""
	latlong := ""
	if presence.TimeOut != nil {
		timeOut = *presence.TimeOut
	}
	if presence.LatitudeLongitudeOut != nil {
		latlong = *presence.LatitudeLongitudeOut
	}

	return r.db.Model(&model.Presences{}).
		Where("user_id = ? AND DATE(presence_date) = ?", presence.UserID, presence.PresenceDate.Format("2006-01-02")).
		Updates(map[string]interface{}{
			"time_out":                timeOut,
			"latitude_longtitude_out": latlong,
		}).Error
}

// Cek absen hari ini
func (r *presencerepository) QueryCheckPresenceToday(employeeID uint) (*model.Presences, error) {
	var presence model.Presences
	today := time.Now().Format("2006-01-02")
	err := r.db.Preload("PresenceStatus").
		Where("user_id = ? AND DATE(presence_date) = ?", employeeID, today).
		First(&presence).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &presence, nil
}
