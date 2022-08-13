package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	Id          string `gorm:"primary_key; unique"`
	File        string
	Body        string
	Title       string
	Excerpt     string
	Published   string
	MetaTitle   string
	Keywords    string
	Description string
}

func (post *Post) BeforeCreate(db *gorm.DB) error {
	post.Id = uuid.New().String()
	return nil
}
