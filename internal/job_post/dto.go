package jobpost

// Create job post input
type CreateJobPostInput struct {
	EmployerID      int    `json:"employer_id"`      // ID of the employer who is creating the job post
	LocationID      int    `json:"location_id"`      // ID of the location where the job is located
	Title           string `json:"title"`            // Title of the job post
	Description     string `json:"description"`      // Description of the job
	Level           string `json:"level"`            // Level of the job (junior, middle, senior)
	ExperienceYears int    `json:"experience_years"` // Number of years of experience required
}
