package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Setting struct {
	gorm.Model
	Id     uuid.UUID `gorm:"primary_key; unique"`
	Key    string
	Value  string
	SiteId string
	Site   Site `gorm:"foreignKey:SiteId"`
}

func (setting *Setting) BeforeCreate(db *gorm.DB) error {
	setting.Id = uuid.New()
	return nil
}
