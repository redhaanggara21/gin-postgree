package controllerToken

import (
	"pelatihan-be/internal/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(e *model.KodeOtpEntity) (*model.KodeOtpEntityModel, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(e *model.KodeOtpEntity) (*model.KodeOtpEntityModel, error) {

	var data model.KodeOtpEntityModel
	data.KodeOtpEntity = *e
	err := r.db.Debug().Create(&data).Error
	if err != nil {
		return nil, err
	}
	err = r.db.Debug().Model(&data).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
