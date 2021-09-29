package controllerAuth

import (
	"fmt"
	util "pelatihan-be/helpers/utils"
	"pelatihan-be/internal/model"
	//"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(payload *LoginRequest) (*LoginResponse, error)
	Register(payload *RegisterRequest) (*RegisterResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) Login(payload *LoginRequest) (*LoginResponse, error) {

	var result *LoginResponse

	data, err := s.repository.FindByUsername(&payload.Email)
	if data == nil {
		//fmt.Println("service 1")
		return result, err
	}

	comparePassword := util.ComparePassword(data.Password, payload.Password)

	if comparePassword != nil {
		//fmt.Println("service 2")
		return result, err
	}

	//dataUserLoginEntityModel := model.UserLoginEntityModel{}
	//dataUserLoginEntityModel.UserLoginEntity = *data

	token, err := data.GenerateToken()
	if err != nil {
		fmt.Println("service 3")
		return result, err
	}

	result = &LoginResponse{
		UserLoginEntity: data.UserLoginEntity,
		Token:           token,
	}

	return result, nil
}

func (s *service) Register(payload *RegisterRequest) (*RegisterResponse, error) {

	var result *RegisterResponse
	var data *model.UserLoginEntityModel

	data, err := s.repository.Create(&payload.UserLoginEntity)
	if data == nil {
		return result, err
	}

	result = &RegisterResponse{
		UserLoginEntity: data.UserLoginEntity,
	}

	return result, nil
}
