package models

import "time"

type JobPost struct {
	ID             int       `db:"id"`
	EmployerID     int       `db:"employer_id"`
	Title          string    `db:"title"`
	Description    string    `db:"description"`
	Level          string    `db:"level"`
	ExperienceYears int      `db:"experience_years"`
	LocationID     int       `db:"location_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type Keyword struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type JobPostKeyword struct {
	JobPostID int `db:"job_post_id"`
	KeywordID int `db:"keyword_id"`
}