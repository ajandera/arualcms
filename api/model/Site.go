package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Site struct {
	gorm.Model
	Id     uuid.UUID `gorm:"primary_key; unique"`
	Name   string
	UserId string
	User   User `gorm:"foreignKey:UserId"`
}

func (site *Site) BeforeCreate(db *gorm.DB) error {
	site.Id = uuid.New()
	return nil
}
