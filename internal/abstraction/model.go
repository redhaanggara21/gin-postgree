package abstraction

import (
	"time"

	"gorm.io/gorm"
)

type Id struct {
	ID uint `gorm:"primaryKey"`
}

type IdAI struct {
	ID uint `gorm:"primaryKey;autoIncrement;"`
}

type Entity struct {
	Id
	Filter
}

type EntityAI struct {
	IdAI
	Filter
}

type Filter struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

