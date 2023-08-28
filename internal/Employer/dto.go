package employer

// Create employer profile input
type CreateEmployerProfileInput struct {

}

// Create job post input
type CreateJobPostInput struct {

}

// Update applicant status input
type UpdateApplicantStatusInput struct {
	JobPostID   int
	ApplicantID int
	Status      string
}