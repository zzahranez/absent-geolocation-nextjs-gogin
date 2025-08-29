package service

import (
	"gin/dto"
	"gin/model"
	"gin/repository"
	"time"
)

type PresenceService interface {
	HandleCreatePresence(email string, req dto.PresenceRequest) (*dto.CreatePresenceResponse, error)
}

type presenceservice struct {
	repo      repository.PresenceRepository
	usersRepo repository.UsersRepository
}

func NewPresenceService(serv repository.PresenceRepository, usersrepo repository.UsersRepository) PresenceService {
	return &presenceservice{
		repo:      serv,
		usersRepo: usersrepo,
	}
}

func (s *presenceservice) HandleCreatePresence(email string, req dto.PresenceRequest) (*dto.CreatePresenceResponse, error) {
	now := time.Now()
	isLate := false
	lateThreshold := time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, time.Local)

	// Ambil user
	user, err := s.usersRepo.QueryGetByEmail(email)
	if err != nil {
		return nil, err
	}

	latlong := ""
	if req.LatitudeLongitude != nil {
		latlong = *req.LatitudeLongitude
	}

	// Cek absen hari ini
	presenceToday, err := s.repo.QueryCheckPresenceToday(user.ID)
	if err != nil {
		return nil, err
	}

	// Helper pointer string
	ptrString := func(s string) *string { return &s }

	// Absen pagi
	if presenceToday == nil {
		if now.After(lateThreshold) {
			isLate = true
		}
		presence := model.Presences{
			UserID:              user.ID,
			PresenceDate:        now,
			TimeIn:              now.Format("15:04:05"),
			Is_late:             isLate,
			LatitudeLongitudeIn: latlong,
			StatusID:            1,
			CreatedAt:           now,
		}
		if err := s.repo.QueryCreatePresence(&presence); err != nil {
			return nil, err
		}
		presenceToday, _ = s.repo.QueryCheckPresenceToday(user.ID)
	}

	// Absen sore
	if presenceToday.TimeOut == nil {
		timeOutStr := now.Format("15:04:05")
		presenceToday.TimeOut = ptrString(timeOutStr)
		presenceToday.LatitudeLongitudeOut = ptrString(latlong)
		if err := s.repo.QueryCreateAbsentAfterNoon(presenceToday); err != nil {
			return nil, err
		}
		presenceToday, _ = s.repo.QueryCheckPresenceToday(user.ID)
	}

	// Ambil status
	status := "Hadir"
	if presenceToday.PresenceStatus != nil {
		status = presenceToday.PresenceStatus.StatusName
	}

	// Build response
	response := dto.CreatePresenceResponse{
		Name:                 user.Name,
		PresenceDate:         presenceToday.PresenceDate,
		TimeIn:               presenceToday.TimeIn,
		TimeOut:              presenceToday.TimeOut,
		IsLate:               presenceToday.Is_late,
		LatitudeLongitudeIn:  presenceToday.LatitudeLongitudeIn,
		LatitudeLongitudeOut: presenceToday.LatitudeLongitudeOut,
		Status:               status,
	}

	return &response, nil
}
