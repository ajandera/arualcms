package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Text struct {
	gorm.Model
	Id     uuid.UUID `gorm:"primary_key; unique"`
	Key    string
	Value  string
	SiteId string
	Site   Site `gorm:"foreignKey:SiteId"`
}

func (text *Text) BeforeCreate(db *gorm.DB) error {
	text.Id = uuid.New()
	return nil
}
