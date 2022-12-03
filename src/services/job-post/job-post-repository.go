package jobPostService

import (
	"net/http"

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
func (r *repository) CreateJobPost(input *dto.CreateJobPost) (*models.JobPosts, int, string) {
	var post models.JobPosts
	db := r.db.Model(&post)

	post.UserID = input.UserID
	post.CompanyID = input.CompanyID
	post.LocationID = input.LocationID
	post.Name = input.Name
	post.Description = input.Description
	post.IsActive = true
	post.JobPostSkillSets = input.JobPostSkillSets
	
	if r.db.Debug().Create(&post).Error != nil {
		return nil, http.StatusForbidden, "Create new instance failed" 
	}
	db.Commit()
	return &post, http.StatusCreated, "nil"
}

func (r *repository) GetAllJobPosts() (*[]models.JobPosts, int, string) {
	var posts []models.JobPosts
	db := r.db.Model(&posts)

	if db.Debug().Select("*").Find(&posts).Error != nil {
		return &[]models.JobPosts{}, http.StatusNotFound, "Data do not exist"
	}
	return &posts, http.StatusOK, "nil"
}