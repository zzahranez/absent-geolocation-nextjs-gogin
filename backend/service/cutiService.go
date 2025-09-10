package service

import (
	"gin/dto"
	"gin/model"
	"gin/repository"
	"time"
)

type CutiService interface {
	HandleCreateCuti(email string, req dto.CreateCutiRequest) (*dto.CreateCutiResponse, error)
}

type cutiservice struct {
	repo      repository.CutiRepository
	usersrepo repository.UsersRepository
}

func NewCutiService(repo repository.CutiRepository, usersrepo repository.UsersRepository) CutiService {
	return &cutiservice{
		repo:      repo,
		usersrepo: usersrepo,
	}
}

func (s *cutiservice) HandleCreateCuti(email string, req dto.CreateCutiRequest) (*dto.CreateCutiResponse, error) {
	userlogin, err := s.usersrepo.QueryGetByEmail(email)
	if err != nil {
		return nil, err
	}
	const layout = "2006-01-02"
	userID := userlogin.ID

	// Var Date
	submissionDate := time.Now()
	startDate, err := time.Parse(layout, req.StartDate)
	if err != nil {
		return nil, err
	}
	endDate, err := time.Parse(layout, req.EndDate)
	if err != nil {
		return nil, err
	}

	createCuti := model.Cuti{
		UserID:         userID,
		SubmissionDate: submissionDate,
		StartDate:      startDate,
		EndDate:        endDate,
		Information:    req.Informations,
		StatusCutiID:   1,
	}

	// save to database
	if err := s.repo.QueryCreateCuti(&createCuti); err != nil {
		return nil, err
	}

	savedCuti, err := s.repo.GetCutiStatusById(userID)
	if err != nil {
		return nil, err
	}

	result := dto.CreateCutiResponse{
		UserName:       userlogin.Name,
		SubmissionDate: createCuti.SubmissionDate.Format(layout),
		StartDate:      createCuti.StartDate.Format(layout),
		EndDate:        createCuti.EndDate.Format(layout),
		Information:    createCuti.Information,
		StatusCuti:     savedCuti.StatusCuti.StatusName,
	}

	return &result, nil

}
