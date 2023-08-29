package employer

// Create employer profile input
type CreateEmployerProfileInput struct {
	UserID      int    `json:"user_id"`
	CompanyID   int    `json:"company_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

// TODO(): change to Applicant module! Update applicant status input
// TODO() Add json tags
type UpdateApplicantStatusInput struct {
	JobPostID   int
	ApplicantID int
	Status      string
}
