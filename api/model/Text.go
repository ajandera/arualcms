package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Text struct {
	Id    string `gorm:"primary_key; unique"`
	Key   string
	Value string
}

func (text *Text) BeforeCreate(db *gorm.DB) error {
	text.Id = uuid.New().String()
	return nil
}
