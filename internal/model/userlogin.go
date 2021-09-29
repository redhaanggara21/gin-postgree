package model

import (
	"os"
	util "pelatihan-be/helpers/utils"
	"pelatihan-be/internal/abstraction"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type UserLoginEntity struct {
	Name     string `gorm:"size:100;not null" json:"name" validate:"required"`
	Email    string `gorm:"size:100;not null;unique_index" json:"email" validate:"required,email,unique"`
	Nik      string `gorm:"size:100;not null;unique_index" json:"nik" validate:"required,unique"`
	Password string `json:"password" validate:"required"`
	IsActive bool   `json:"is_active" validate:"required" gorm:"default:false"`
}

type UserLoginEntityModel struct {
	abstraction.EntityAI

	UserLoginEntity

	abstraction.Filter
}

func (UserLoginEntityModel) TableName() string {
	return "user"
}

func (m *UserLoginEntityModel) BeforeCreate(db *gorm.DB) error {
	m.Password = util.HashPassword(m.Password)
	m.CreatedAt = time.Now().Local()
	return nil
}

func (m *UserLoginEntityModel) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}

func (m *UserLoginEntityModel) GenerateToken() (string, error) {
	var (
		jwtKey = os.Getenv("JWT_KEY")
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       m.ID,
		"iss":      os.Getenv("KEY"),
		"username": m.Name,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	return tokenString, err
}
