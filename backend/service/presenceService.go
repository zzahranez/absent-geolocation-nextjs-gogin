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
		Name:     presence[0].User.Name,
		Presence: presenceAll,
	}

	return result, nil
}
