package controllerAuth

import "pelatihan-be/internal/model"

//login
type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type LoginResponse struct {
	model.UserLoginEntity
	Token string `json:"token"`
}

//register
type RegisterRequest struct {
	model.UserLoginEntity
}

type RegisterResponse struct {
	model.UserLoginEntity
}
