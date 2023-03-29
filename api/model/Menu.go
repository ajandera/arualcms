package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Id       uuid.UUID `gorm:"primary_key; unique"`
	ParentId string
	Name     string
	Url      string
	Children []*Menu `gorm:"foreignkey:ParentId"`
	SiteId   string
	Site     Site `gorm:"foreignKey:SiteId"`
}

func (menu *Menu) BeforeCreate(db *gorm.DB) error {
	menu.Id = uuid.New()
	return nil
}
