package controllerAuth

import (
	"errors"
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
	if err != nil {
		return result, errors.New("email belum terdaftar")
	}

	comparePassword := util.ComparePassword(data.Password, payload.Password)
	if comparePassword != nil {
		return result, errors.New("password yang anda gunakan salah")
	}

	// cek is activ
	if data.IsActive == false {
		return result, errors.New("akun anda blom aktif")
	}

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

	dataEmail, _ := s.repository.FindByParameter(payload.Email, "email = ?")
	if dataEmail != nil {
		return result, errors.New("email sudah terdaftar")
	}
	dataNik, _ := s.repository.FindByParameter(payload.Nik, "nik = ?")
	if dataNik != nil {
		return result, errors.New("nik sudah terdaftar")
	}
	dataNo, _ := s.repository.FindByParameter(payload.Nomor_hp, "nomor_hp = ?")
	if dataNo != nil {
		return result, errors.New("nomor handphone sudah terdaftar")
	}

	data, err := s.repository.Create(&payload.UserLoginEntity)
	if data == nil {
		return result, err
	}

	result = &RegisterResponse{
		UserLoginEntity: data.UserLoginEntity,
	}

	return result, nil
}
