package models

import "time"

type Company struct {
	ID                int       `db:"id"`
	LocationID        int       `db:"location_id"`
	Name              string    `db:"name"`
	Description       string    `db:"description"`
	CompanyWebsiteUrl string    `db:"company_website_url"`
	CompanyImage      string    `db:"company_imange"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}
