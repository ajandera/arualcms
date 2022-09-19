package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Id          uuid.UUID `gorm:"primary_key; unique"`
	File        string
	Body        string
	Title       string
	Excerpt     string
	Published   string
	MetaTitle   string
	Keywords    string
	Description string
	SiteId      string
	Site        Site `gorm:"foreignKey:SiteId"`
}

func (post *Post) BeforeCreate(db *gorm.DB) error {
	post.Id = uuid.New()
	return nil
}
