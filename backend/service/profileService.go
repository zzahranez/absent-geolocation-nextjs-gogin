package service

import (
	"errors"
	"gin/dto"
	"gin/model"
	"gin/repository"
	"sync"
)

type ProfileService interface {
	HandleGetAccountUser(prefereedUsername string) (*dto.ProfilePageResponse, error)
}

type profileservice struct {
	repo     repository.ProfileRepository
	userrepo repository.UsersRepository
}

func NewProfileService(repo repository.ProfileRepository, userrepo repository.UsersRepository) ProfileService {
	return &profileservice{
		repo:     repo,
		userrepo: userrepo,
	}
}

func (s *profileservice) HandleGetAccountUser(prefereedUsername string) (*dto.ProfilePageResponse, error) {
	user, err := s.userrepo.QueryGetByEmail(prefereedUsername)
	if err != nil {
		return nil, err
	}

	userID := user.ID

	var (
		userAccount *model.User
		dataUser    *model.Profile
		err1, err2  error
	)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		userAccount, err1 = s.repo.QueryFindUserAccountByEmail(userID)
		if err1 != nil {
			err1 = errors.New("account data not found")
		}
	}()

	go func() {
		defer wg.Done()
		dataUser, err2 = s.repo.QueryFindBiodataUserByEmail(userID)
		if err2 != nil {
			err2 = errors.New("biodata not found")
		}
	}()
	wg.Wait()

	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}

	account := dto.UserAccountResponse{
		Name:  userAccount.Name,
		Email: userAccount.Email,
		Role:  userAccount.Role,
	}
	var avatar string
	if dataUser.Avatar != "" {
		avatar = dataUser.Avatar
	}

	response := dto.ProfilePageResponse{
		Bio:         dataUser.Bio,
		NoHp:        dataUser.Phone,
		Address:     dataUser.Address,
		Avatar:      &avatar,
		AccountData: account,
	}

	return &response, nil

}
