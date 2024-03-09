package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName     string `gorm:"type:varchar(100);not null;" json:"first_name" valid:"required~Name is required,matches(^[a-zA-Z ]+$)~Name must be alphabetic"`
	LastName      string `gorm:"type:varchar(100);not null;" json:"last_name" valid:"required~Name is required,matches(^[a-zA-Z ]+$)~Name must be alphabetic"`
	Email         string `json:"email"`
	ContactNumber string `gorm:"type:varchar(20);not null" json:"contact_number" valid:"required~Contact number is required,numeric~Contact number must be numeric"`
	Password      string `json:"password"`
}
