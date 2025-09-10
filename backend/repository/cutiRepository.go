package repository

import (
	"gin/model"

	"gorm.io/gorm"
)

type CutiRepository interface {
	// Submission
	QueryCreateCuti(model *model.Cuti) error
	GetCutiStatusById(userID uint) (*model.Cuti, error)
}

type cutirepository struct {
	db *gorm.DB
}

func NewCutiRepository(db *gorm.DB) CutiRepository {
	return &cutirepository{db: db}
}

func (r *cutirepository) GetCutiStatusById(userID uint) (*model.Cuti, error) {
	var cuti model.Cuti
	if err := r.db.Preload("StatusCuti").Where("user_id = ?", userID).First(&cuti).Error; err != nil {
		return nil, err
	}
	return &cuti, nil
}

func (r *cutirepository) QueryCreateCuti(model *model.Cuti) error {
	return r.db.Create(model).Error
}
