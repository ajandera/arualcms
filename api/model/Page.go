package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Page struct {
	gorm.Model
	Id          uuid.UUID `gorm:"primary_key; unique"`
	Key         string
	Body        string
	Title       string
	MetaTitle   string
	Keywords    string
	Description string
	SiteId      string
	Site        Site `gorm:"foreignKey:SiteId"`
}

func (page *Page) BeforeCreate(db *gorm.DB) error {
	page.Id = uuid.New()
	return nil
}
