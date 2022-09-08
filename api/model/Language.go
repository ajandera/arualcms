package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Language struct {
	gorm.Model
	Id      string `gorm:"primary_key; unique"`
	Name    string
	Key     string
	Default bool
	SiteId  string
}

func (language *Language) BeforeCreate(db *gorm.DB) error {
	language.Id = uuid.New().String()
	return nil
}
