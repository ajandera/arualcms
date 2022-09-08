package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Setting struct {
	gorm.Model
	Id     string `gorm:"primary_key; unique"`
	Key    string
	Value  string
	SiteId string
}

func (setting *Setting) BeforeCreate(db *gorm.DB) error {
	setting.Id = uuid.New().String()
	return nil
}
