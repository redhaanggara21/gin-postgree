package controllerAuth

import (
	"errors"
	"pelatihan-be/internal/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(e *model.UserLoginEntity) (*model.UserLoginEntityModel, error)
	FindByUsername(username *string) (*model.UserLoginEntityModel, error)
	FindByParameter(parameter string, condisi string) (*model.UserLoginEntityModel, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(e *model.UserLoginEntity) (*model.UserLoginEntityModel, error) {

	var data model.UserLoginEntityModel
	data.UserLoginEntity = *e
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

func (r *repository) FindByUsername(username *string) (*model.UserLoginEntityModel, error) {

	var data model.UserLoginEntityModel
	err := r.db.Debug().Where("email = ?", &username).First(&data).Error
	if err != nil {
		return nil, errors.New("email belum terdaftar")
	}
	return &data, nil

}
func (r *repository) FindByParameter(parameter string, condisi string) (*model.UserLoginEntityModel, error) {
	var data model.UserLoginEntityModel
	err := r.db.Debug().Where(condisi, parameter).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil

}
