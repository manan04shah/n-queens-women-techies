package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	FirstName     string `gorm:"type:varchar(100);not null;" json:"first_name" valid:"required~Name is required,matches(^[a-zA-Z ]+$)~Name must be alphabetic"`
	LastName      string `gorm:"type:varchar(100);not null;" json:"last_name" valid:"required~Name is required,matches(^[a-zA-Z ]+$)~Name must be alphabetic"`
	Email         string `json:"email"`
	ContactNumber string `gorm:"type:varchar(20);not null" json:"contact_number" valid:"required~Contact number is required,numeric~Contact number must be numeric"`
	CompanyID     string `gorm:"type:varchar(20);not null" json:"company_code" valid:"required~Company code is required"`
	Password      string `json:"password"`
}

type HR struct {
	gorm.Model
	CompanyName   string     `gorm:"type:varchar(100);not null;" json:"company_name" valid:"required~Company name is required,matches(^[a-zA-Z ]+$)~Company name must be alphabetic"`
	Email         string     `json:"email"`
	ContactNumber string     `gorm:"type:varchar(20);not null" json:"contact_number" valid:"required~Contact number is required,numeric~Contact number must be numeric"`
	Password      string     `json:"password"`
	Reports       []Report   `json:"reports" gorm:"foreignKey:HRID"`
	Employees     []Employee `json:"employees" gorm:"foreignKey:CompanyID"`
}

type HRResponse struct {
	CompanyName   string `json:"company_name"`
	Email         string `json:"email"`
	ContactNumber string `json:"contact_number"`
}

type EmployeeResponse struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	ContactNumber string `json:"contact_number"`
	CompanyID     string `json:"company_code"`
}
