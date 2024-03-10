package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	HRID              uint   `json:"hr_id" gorm:"index"`
	ReportAgainst     string `gorm:"type:varchar(500);not null;" json:"report_against" valid:"required~Report against is required"`
	ReportDescription string `gorm:"type:varchar(1000);not null;" json:"report_description" valid:"required~Report description is required"`
	ProofURL          string `json:"proof_url"`
}

type ReportResponse struct {
	ReportAgainst     string `json:"report_against"`
	ReportDescription string `json:"report_description"`
	ProofURL          string `json:"proof_url"`
	HRID              uint   `json:"hr_id"`
}
