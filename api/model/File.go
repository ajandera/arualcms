package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type File struct {
	Id   string `gorm:"primary_key; unique"`
	Name string
	Src  string
}

func (file *File) BeforeCreate(db *gorm.DB) error {
	file.Id = uuid.New().String()
	return nil
}
