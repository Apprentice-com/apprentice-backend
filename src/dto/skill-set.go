package dto

type CreateSkillSet struct {
	Name string    `json:"institution_name" validate:"required"`
}