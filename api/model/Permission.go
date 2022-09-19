package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	Id     uuid.UUID `gorm:"primary_key; unique"`
	UserId string
	SiteId string
	Role   string
	Site   Site `gorm:"foreignKey:SiteId"`
	User   User `gorm:"foreignKey:UserId"`
}

func (permission *Permission) BeforeCreate(db *gorm.DB) error {
	permission.Id = uuid.New()
	return nil
}
