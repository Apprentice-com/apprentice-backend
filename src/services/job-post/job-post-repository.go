package jobPostService

import (
	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	"github.com/KadirbekSharau/apprentice-backend/src/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

/* Create Education Details Repository Service */
func (r *repository) CreateJobPost(input *dto.CreateJobPost) (*models.JobPosts, string) {
	var post models.JobPosts
	db := r.db.Model(&post)
	errorCode := make(chan string, 1)

	post.UserID = input.UserID
	post.CompanyID = input.CompanyID
	post.LocationID = input.LocationID
	post.Name = input.Name
	post.Description = input.Description
	post.IsActive = true
	post.JobPostSkillSets = input.JobPostSkillSets
	
	addNew := r.db.Debug().Create(&post)

	db.Commit()

	if addNew.Error != nil {
		errorCode <- "CREATE_FAILED_403"
	} else {
		errorCode <- "nil"
	}

	return &post, <-errorCode
}

func (r *repository) GetAllJobPosts() (*[]models.JobPosts, string) {
	var posts []models.JobPosts
	errorCode := make(chan string, 1)

	db := r.db.Model(&posts)
	result := db.Debug().Select("*").Find(&posts)

	if result.Error != nil {
		errorCode <- "RESULTS_BOOKS_NOT_FOUND_404"
	} else {
		errorCode <- "nil"
	}

	return &posts, <- errorCode
}