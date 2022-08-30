package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Id      string `gorm:"primary_key; unique"`
	Name    string
	Src     string
	Gallery string
	SiteId  string
}

func (file *File) BeforeCreate(db *gorm.DB) error {
	file.Id = uuid.New().String()
	return nil
}
