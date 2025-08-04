package service

import (
	"errors"
	"gin/dto"
	"gin/repository"
)

type AuthService interface {
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authservice struct {
	repo repository.AuhtRepository
}

func NewAuthService(re repository.AuhtRepository) AuthService {
	return &authservice{repo: re}
}

func (s *authservice) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	users, err := s.repo.QueryLogin(req.Email)
	if err != nil {
		return nil, errors.New("email tidak ditemukan")
	}

	if users.Password != req.Password {
		return nil, errors.New("invalid email or password")
	}
	return &dto.LoginResponse{
		Name:  users.Name,
		Email: users.Email,
	}, nil

}
