package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id          uuid.UUID `gorm:"primary_key; unique"`
	Name        string
	Username    string
	Password    string
	ForgotToken string
	ValidToken  time.Time
	ParentId    string
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	user.Id = uuid.New()
	return nil
}
