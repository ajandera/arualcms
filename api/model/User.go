package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id          string `gorm:"primary_key; unique"`
	Name        string
	Username    string
	Password    string
	ForgotToken string
	ValidToken  time.Time
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	user.Id = uuid.New().String()
	return nil
}
