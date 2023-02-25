package models

import (
	"fmt"
	"net/http"

	"github.com/KadirbekSharau/apprentice-backend/src/dto"
	"gorm.io/gorm"
)

// JobPost is database entity for each job post
type JobPost struct {
	gorm.Model
	UserID        uint
	CompanyID     uint
	LocationID    uint
	JobPostTypeID uint
	IsRemote      bool
	Salary        int
	Currency      string
	Name          string `gorm:"type:varchar(1000)"`
	Description   string `gorm:"type:varchar(130)"`
	Link          string `gorm:"type:varchar(1000)"`
	IsActive      bool   `gorm:"default:true"`
}

type Application struct {
	gorm.Model
	UserID uint
	JobID  uint
	Status string
	Resume string
}

// JobPostType is database entity for each job post type
type JobPostType struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(1000)"`
	JobPosts []JobPost `gorm:"foreignKey:JobPostTypeID"`
}

type JobPostRepository struct {
	db *gorm.DB
}

func NewJobPostRepository(db *gorm.DB) *JobPostRepository {
	return &JobPostRepository{db: db}
}

/* Create Education Details Repository Service */
func (r *JobPostRepository) CreateJobPost(input *dto.CreateJobPost) (*JobPost, int, string) {
	var post JobPost
	post.UserID = input.UserID
	post.CompanyID = input.CompanyID
	post.LocationID = input.LocationID
	post.Name = input.Name
	post.Description = input.Description
	post.IsActive = true

	if r.db.Debug().Create(&post).Error != nil {
		return nil, http.StatusForbidden, "Create new instance failed"
	}
	r.db.Model(&post).Commit()
	return &post, http.StatusCreated, "nil"
}

func (r *JobPostRepository) GetAllJobPosts() (*[]JobPost, int, string) {
	var posts []JobPost
	fmt.Println(r.db)
	if r.db.Model(&posts).Debug().Select("*").Find(&posts).Error != nil {
		fmt.Println(r.db)
		return &[]JobPost{}, http.StatusNotFound, "Data do not exist"
	}
	return &posts, http.StatusOK, "nil"
}

func (r *JobPostRepository) GetJobPostByID(id string) (*JobPost, int, string) {
	var jobpost JobPost
	if r.db.Model(&jobpost).Debug().Select("*").Where("id = ?", id).Find(&jobpost).RowsAffected != 1 {
		return nil, http.StatusNotFound, "Data not found"
	}
	return &jobpost, http.StatusOK, "nil"
}
