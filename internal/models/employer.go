package models

import "time"

type EmployerProfile struct {
	ID          int       `db:"id"`
	UserID      int       `db:"user_id"`
	CompanyID   int       `db:"company_id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
