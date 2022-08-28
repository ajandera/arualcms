package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Site struct {
	Id          string `gorm:"primary_key; unique"`
	Name        string
	UserId    	string
}

func (site *Site) BeforeCreate(db *gorm.DB) error {
	site.Id = uuid.New().String()
	return nil
}
