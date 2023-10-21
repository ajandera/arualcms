package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Id       uuid.UUID `gorm:"primary_key; unique"`
	ParentId uuid.UUID `gorm:"default:null"`
	Name     string
	Url      string
	SiteId   string
	Site     Site `gorm:"foreignKey:SiteId"`
	Root     bool
	Order    uint
}

func (menu *Menu) BeforeCreate(db *gorm.DB) error {
	menu.Id = uuid.New()
	return nil
}
