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
	r.db.Model(&post).Commit()
	return &post, http.StatusCreated, "nil"
}

func (r *repository) GetAllJobPosts() (*[]models.JobPosts, int, string) {
	var posts []models.JobPosts
	if r.db.Model(&posts).Debug().Select("*").Find(&posts).Error != nil {
		return &[]models.JobPosts{}, http.StatusNotFound, "Data do not exist"
	}
	return &posts, http.StatusOK, "nil"
}

func (r *repository) GetJobPostByID(id string) (*models.JobPosts, int, string) {
	var jobpost models.JobPosts
	if r.db.Model(&jobpost).Debug().Select("*").Where("id = ?", id).Find(&jobpost).RowsAffected != 1 {
		return nil, http.StatusNotFound, "Data not found"
	}
	return &jobpost, http.StatusOK, "nil"
}