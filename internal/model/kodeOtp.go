package model

import (
	util "pelatihan-be/helpers/utils"
	"pelatihan-be/internal/abstraction"
	"time"

	"gorm.io/gorm"
)

type KodeOtpEntity struct {
	Token         string `gorm:"size:100;not null" json:"token"`
	Email         string `gorm:"size:100;not null" json:"email" validate:"required"`
	Request_count int    `gorm:"default:0" json:"request_count"`
}

type KodeOtpEntityModel struct {
	abstraction.EntityAI

	KodeOtpEntity

	abstraction.Filter
}

func (KodeOtpEntityModel) TableName() string {
	return "kode_otp"
}
func (m *KodeOtpEntityModel) BeforeCreate(db *gorm.DB) error {
	m.Token = util.RandStringRunes(6)
	m.CreatedAt = time.Now().Local()
	return nil
}
