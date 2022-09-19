package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Id      uuid.UUID `gorm:"primary_key; unique"`
	Name    string
	Src     string
	Gallery string
	SiteId  string
	Site    Site `gorm:"foreignKey:SiteId"`
}

func (file *File) BeforeCreate(db *gorm.DB) error {
	file.Id = uuid.New()
	return nil
}
