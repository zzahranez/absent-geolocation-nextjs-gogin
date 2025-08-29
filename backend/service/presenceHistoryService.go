package service

import (
	"gin/dto"
	"gin/repository"
	"time"
)

type PresenceHistoryService interface {
	GetPresenceByEmail(email string) (*dto.PresenceResponse, error)
}

type presencehistoryservice struct {
	repo repository.PresenceHistoryRepository
}

func NewPresenceHistoryService(repo repository.PresenceHistoryRepository) PresenceHistoryService {
	return &presencehistoryservice{repo: repo}
}

func (s *presencehistoryservice) GetPresenceByEmail(email string) (*dto.PresenceResponse, error) {
	// time variable
	now := time.Now()
	currentmonth := int(now.Month())
	currentYear := now.Year()

	presence, err := s.repo.QueryPresenceByEmail(email, currentmonth, currentYear)
	if err != nil {
		return nil, err
	}

	// 1. Looop Monly Summary
	var summary dto.MontlyPresenceSummary
	for _, v := range presence {
		switch v.PresenceStatus.StatusName {
		case "Hadir":
			summary.TotalHadir++
		case "Izin":
			summary.TotalIzin++
		case "Sakit":
			summary.TotalSakit++
		}

		if v.Is_late == true {
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
		if todaypresence.TimeIn != "" {
			presenceToday.Morning = "Sudah Absen"
		}
		if todaypresence.TimeOut != nil && *todaypresence.TimeOut != "" {
			presenceToday.AfterNon = "Sudah Absen"
		}
	}

	// Loop tiga
	var presenceAll []dto.PresencesAll
	for _, v := range presence {

		var latLongOut string
		if v.LatitudeLongitudeOut != nil {
			latLongOut = *v.LatitudeLongitudeOut
		} else {
			latLongOut = "" // atau kasih default value
		}
		presenceAll = append(presenceAll, dto.PresencesAll{
			PresenceDate:           v.PresenceDate,
			TimeIn:                 &v.TimeIn,
			TimeOut:                v.TimeOut,
			IsLate:                 v.Is_late,
			LatitudeLongTitudeIn:   v.LatitudeLongitudeIn,
			LatitudeLongTitudeeOut: latLongOut,
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
