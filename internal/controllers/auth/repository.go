package controllerAuth

import (
	"pelatihan-be/internal/model"

	"gorm.io/gorm"
)

type Repository interface {
	Create(e *model.UserLoginEntity) (*model.UserLoginEntityModel, error)
	FindByUsername(username *string) (*model.UserLoginEntityModel, error)
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
	err := r.db.Debug().Where("name = ? AND is_active = ? ", &username, true).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil

}
