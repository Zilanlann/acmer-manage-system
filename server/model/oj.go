package model

import (
	"time"

	"github.com/zilanlann/acmer-manage-system/server/global"
	"gorm.io/gorm"
)

type OJContest struct {
	gorm.Model
	Name            string    `json:"name"`
	ContestID       int       `json:"contestID"`
	OJ              string    `json:"oj"`
	Type            string    `json:"type"`
	DurationSeconds int       `json:"durationSeconds"`
	StartTime       time.Time `json:"startTime"`
}

func CFCreateContest(contest *OJContest) error {
	return global.DB.Create(contest).Error
}

func CFGetContests() (contests []OJContest, err error) {
	err = global.DB.Find(&contests).Error
	return
}

func GetContestsWithTime(startTime, endTime time.Time) (contests []OJContest, err error) {
	err = global.DB.Where("start_time BETWEEN ? AND ?", startTime, endTime).Find(&contests).Error
	return
}
