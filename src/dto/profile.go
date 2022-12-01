package dto

import "time"

type GetSeekerProfile struct {
	UserID uint `validate:"required"`
}

type CreateEducationDetails struct {
	UserID          uint      `json:"user_id"`
	InstitutionName string    `json:"institution_name" validate:"required"`
	Major           string    `json:"major" validate:"required"`
	Degree          string    `json:"degree" validate:"required"`
	StartDate       time.Time `json:"start_date" validate:"required"`
	EndDate         time.Time `json:"end_date" validate:"required"`
}
