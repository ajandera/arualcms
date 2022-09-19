package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Language struct {
	gorm.Model
	Id      uuid.UUID `gorm:"primary_key; unique"`
	Name    string
	Key     string
	Default bool
	SiteId  string
	Site    Site `gorm:"foreignKey:SiteId"`
}

func (language *Language) BeforeCreate(db *gorm.DB) error {
	language.Id = uuid.New()
	return nil
}
