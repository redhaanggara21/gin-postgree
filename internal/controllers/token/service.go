package controllerToken

import (
	"pelatihan-be/internal/model"
)

type Service interface {
	Create(payload *CreateRequest) (*CreateRequest, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) Create(payload *CreateRequest) (*CreateRequest, error) {

	var result *CreateRequest
	var data *model.KodeOtpEntityModel

	data, err := s.repository.Create(&payload.KodeOtpEntity)
	if data == nil {
		return result, err
	}

	result = &CreateRequest{
		KodeOtpEntity: data.KodeOtpEntity,
	}

	return result, nil
}
