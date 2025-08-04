package service

import (
	"gin/dto"
	"gin/repository"
)

type PresenceService interface {
	GetPresenceByEmail(email string) (*dto.PresenceResponse, error)
}

type presenceservice struct {
	repo repository.PresenceRepository
}

func NewPresenceService(repo repository.PresenceRepository) PresenceService {
	return &presenceservice{repo: repo}
}

func (s *presenceservice) GetPresenceByEmail(email string) (*dto.PresenceResponse, error) {
	presence, err := s.repo.QueryPresenceByEmail(email)
	if err != nil {
		return nil, err
	}

	// 1. Looop Monly Summary
	var summary dto.MontlyPresenceSummary

	for _, v := range presence {
		switch v.PresentStatus.StatusName {
		case "hadir":
			summary.TotalHadir++
		case "izin":
			summary.TotalIzin++
		case "sakit":
			summary.TotalSakit++
		case "terlambat":
			summary.TotalTerlambat++
		}
	}

	// Loop 2 Montly Sumary
	todaypresence, err := s.repo.QueryTodayPresenceByEmail(email)
	if err != nil {
		return nil, err
	}
	presenceToday := dto.PresenceToDay{
		Morning:  "Belum Absen",
		AfterNon: "Belum Absen",
	}

	if todaypresence != nil {
		if todaypresence.TimeIn != nil && *todaypresence.TimeIn != "" {
			presenceToday.Morning = "Sudah Absen"
		}
		if todaypresence.TimeOut != nil && *todaypresence.TimeOut != "" {
			presenceToday.AfterNon = "Sudah Absen"
		}
	}

	// Loop tiga
	var presenceAll []dto.PresencesAll
	for _, v := range presence {
		presenceAll = append(presenceAll, dto.PresencesAll{
			PresenceDate:           v.PresenceDate,
			TimeIn:                 v.TimeIn,
			TimeOut:                v.TimeOut,
			LatitudeLongTitudeIn:   *v.LatitudeLongitudeIn,
			LatitudeLongTitudeeOut: *v.LatitudeLongitudeOut,
		})
	}

	result := &dto.PresenceResponse{
		Name:                  presence[0].User.Name,
		MontlyPresenceSummary: summary,
		PresenceToday:         presenceToday,
		Presence:              presenceAll,
	}

	return result, nil
}
