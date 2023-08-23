package models

import "time"

type Application struct {
	ID           int       `db:"id"`
	ApplicantID  int       `db:"applicant_id"`
	JobPostID    int       `db:"job_post_id"`
	Status       string    `db:"status"`
	AppliedAt    time.Time `db:"applied_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}