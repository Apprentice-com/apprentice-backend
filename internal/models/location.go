package models

type Location struct {
	ID            int    `db:"id"`
	Country       string `db:"country"`
	City          string `db:"city"`
	State         string `db:"state"`
	ZipCode       string `db:"zip_code"`
	StreetAddress string `db:"street_address"`
}
