package models

import "time"

type ApplicantProfile struct {
	ID          int    `db:"id"`
	UserID      int    `db:"user_id"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	ProfilePicture string `db:"profile_picture"`
	Bio         string `db:"bio"`
	ResumePDF   string `db:"resume_pdf"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}