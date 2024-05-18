package model

import (
	"time"

	"github.com/zilanlann/acmer-manage-system/server/global"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// OJContest struct definition
type OJContest struct {
	gorm.Model
	Name            string    `json:"name"`
	ContestID       int       `json:"contestID"`
	OJ              string    `json:"oj"`
	Type            string    `json:"type"`
	DurationSeconds int       `json:"durationSeconds"`
	StartTime       time.Time `json:"startTime"`
}

// OJCreateContest function to create a contest
func OJCreateContest(contest *OJContest) error {
	return global.DB.Create(contest).Error
}

// OJGetContests function to get all contests
func OJGetContests() (contests []OJContest, err error) {
	err = global.DB.Find(&contests).Error
	return
}

// GetContestsWithTime function to get contests within a time range
func GetContestsWithTime(startTime, endTime time.Time) (contests []OJContest, err error) {
	err = global.DB.Where("start_time BETWEEN ? AND ?", startTime, endTime).Find(&contests).Error
	return
}

// GetContestsByMonth function to get contests by month
func GetContestsByMonth(year int, month int) (contests []OJContest, err error) {
	startTime := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endTime := startTime.AddDate(0, 1, 0).Add(-time.Nanosecond)
	return GetContestsWithTime(startTime, endTime)
}

type OJSubmission struct {
	gorm.Model
	Name    string       `gorm:"not null" json:"name"`
	UserID  uint         `gorm:"not null" json:"userID"`
	Rating  int          `json:"rating"`
	Tags    []ProblemTag `gorm:"many2many:oj_submission_problem_tags;constraint:OnDelete:CASCADE;" json:"tags"`
	Verdict string       `json:"verdict"`
	OJ      string       `json:"oj"`
	Time    time.Time    `json:"time"`
}

type ProblemTag struct {
	gorm.Model
	Name string `gorm:"not null;unique" json:"name"`
}

// Create a new OJSubmission
func (submission *OJSubmission) Create() error {
	return global.DB.Create(&submission).Error
}

// Get an OJSubmission by ID
func (submission *OJSubmission) GetByID(id uint) error {
	return global.DB.First(submission, id).Error
}

// Update an existing OJSubmission
func (submission *OJSubmission) Update() error {
	return global.DB.Save(&submission).Error
}

// Delete an OJSubmission by ID (Unscoped and remove associations)
func (submission *OJSubmission) Delete() error {
	return global.DB.Unscoped().Select(clause.Associations).Delete(&submission).Error
}

// Create a new ProblemTag
func (tag *ProblemTag) Create() error {
	return global.DB.Create(&tag).Error
}

// Get a ProblemTag by ID
func (tag *ProblemTag) GetByID(id uint) error {
	return global.DB.First(tag, id).Error
}

// Update an existing ProblemTag
func (tag *ProblemTag) Update() error {
	return global.DB.Save(&tag).Error
}

// Delete a ProblemTag by ID (Unscoped and remove associations)
func (tag *ProblemTag) Delete() error {
	return global.DB.Unscoped().Select(clause.Associations).Delete(&tag).Error
}

func GetSubmissionByTime(startTime, endTime time.Time) (submissions []OJSubmission, err error) {
	err = global.DB.Where("time BETWEEN ? AND ?", startTime, endTime).Find(&submissions).Error
	return
}

func GetUserSubmissionByTime(userId uint, startTime, endTime time.Time) (submissions []OJSubmission, err error) {
	err = global.DB.Where("user_id = ? AND time BETWEEN ? AND ?", userId, startTime, endTime).Find(&submissions).Error
	return
}

func GetAllSubmissionsByUser(userId uint) (submissions []OJSubmission, err error) {
	err = global.DB.Preload("Tags").Where("user_id = ?", userId).Find(&submissions).Error
	return
}

func GetAllTags() (tags []ProblemTag, err error) {
	err = global.DB.Find(&tags).Error
	return
}

type UserStatus struct {
	gorm.Model
	UserID          uint `json:"userID"`
	CFRating        int  `json:"cfRating"`
	CFWeeklyRating  int  `json:"weeklyRating"`
	CFMonthlyRating int  `json:"monthlyRating"`
	WeeklyActive    int  `json:"weeklyActive"`
	MonthlyActive   int  `json:"monthlyActive"`
}

func (status *UserStatus) Create() error {
	return global.DB.Create(&status).Error
}

func (status *UserStatus) Update() error {
	return global.DB.Save(&status).Error
}

func (status *UserStatus) GetByUserID(userID uint) error {
	return global.DB.First(&status, userID).Error
}

func (status *UserStatus) Delete() error {
	return global.DB.Unscoped().Delete(&status).Error
}
