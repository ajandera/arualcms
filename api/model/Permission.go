package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	Id     string `gorm:"primary_key; unique"`
	UserId string
	SiteId string
	Role   string
}

func (permission *Permission) BeforeCreate(db *gorm.DB) error {
	permission.Id = uuid.New().String()
	return nil
}
