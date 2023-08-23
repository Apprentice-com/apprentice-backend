package configs

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

const (
	createUsersTableSQL = `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(100) UNIQUE NOT NULL,
		password VARCHAR(256) NOT NULL,
		role VARCHAR(50) NOT NULL,
		is_active BOOLEAN NOT NULL DEFAULT true,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	createLocationsTableSQL = `CREATE TABLE IF NOT EXISTS locations (
		id SERIAL PRIMARY KEY,
		country VARCHAR(255),
		city VARCHAR(255),
		state VARCHAR(255),
		zip_code VARCHAR(20),
		street_address TEXT
	);`

	createApplicantProfilesTableSQL = `CREATE TABLE IF NOT EXISTS applicant_profiles (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id),
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		profile_picture VARCHAR(255),
		bio TEXT,
		resume_pdf VARCHAR(255),
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	createEmployerProfilesTableSQL = `CREATE TABLE IF NOT EXISTS employer_profiles (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id),
		company_id INT REFERENCES companies(id),
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	createCompaniesTableSQL = `CREATE TABLE IF NOT EXISTS companies (
		id SERIAL PRIMARY KEY,
		location_id INT REFERENCES locations(id),
		name VARCHAR(255) NOT NULL,
		description TEXT,
		company_website_url VARCHAR(255),
		company_image VARCHAR(255),
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	createKeywordTableSQL = `CREATE TABLE IF NOT EXISTS keywords (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50)
	);`

	createJobPostsTableSQL = `CREATE TABLE IF NOT EXISTS job_posts (
		id SERIAL PRIMARY KEY,
		employer_id INT REFERENCES users(id),
		location_id INT REFERENCES locations(id),
		title VARCHAR(255) NOT NULL,
		description TEXT,
		level VARCHAR(50),
		experience_years INT,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	createJobPostKeywordTableSQL = `CREATE TABLE IF NOT EXISTS job_post_keywords (
		job_post_id INT REFERENCES job_posts(id),
		keyword_id INT REFERENCES keywords(id),
		PRIMARY KEY (job_post_id, keyword_id)
	);`

	createApplicationsTableSQL = `CREATE TABLE IF NOT EXISTS applications (
		id SERIAL PRIMARY KEY,
		applicant_id INT REFERENCES users(id),
		job_post_id INT REFERENCES job_posts(id),
		status VARCHAR(50),
		applied_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`
)

func NewPostgresDB() (*sqlx.DB, error) {
	cfg := PostgreConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: os.Getenv("POSTGRES_USERNAME"),
		DBName:   os.Getenv("POSTGRES_DBNAME"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("POSTGRES_PASSWORD")}
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	cfg.migrateDB(db)
	return db, nil
}

func (cfg *PostgreConfig) migrateDB(db *sqlx.DB)  {
	executeTable(db, createUsersTableSQL)
	executeTable(db, createLocationsTableSQL)
	executeTable(db, createCompaniesTableSQL)
	executeTable(db, createKeywordTableSQL)
	executeTable(db, createApplicantProfilesTableSQL)
	executeTable(db, createEmployerProfilesTableSQL)
	executeTable(db, createJobPostsTableSQL)
	executeTable(db, createJobPostKeywordTableSQL)
	executeTable(db, createApplicationsTableSQL)


}

func executeTable(db *sqlx.DB, table string) error {
	if _, err := db.Exec(table); err != nil {
		return fmt.Errorf("error creating database table: %s", err.Error())
	}
	return nil
}
