package dto

type CreateJobPost struct {
	UserID           uint               `json:"user_id"`
	CompanyID        uint               `json:"company_id"`
	LocationID       uint               `json:"location_id"`
	Name             string             `json:"jobpost_name" validate:"required"`
	Description      string             `json:"jobpost_description" validate:"required"`
}