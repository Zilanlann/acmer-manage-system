package model

import (
	"time"

	"github.com/zilanlann/acmer-manage-system/server/global"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Contest struct {
	gorm.Model
	Name      string    `gorm:"not null;unique" json:"name"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Desc      string    `gorm:"type:text" json:"desc"`
	Teams     []Team    `json:"teams"`
}

type Team struct {
	gorm.Model
	ContestID   uint
	ZhName      string       `gorm:"not null;unique" json:"zhName"`
	EnName      string       `json:"enName"`
	CoachID     uint         `gorm:"foreignKey:UserID" json:"coachID"`
	Coach       User         `json:"coach"`
	Contestants []Contestant `json:"contestants"`
	Desc        string       `gorm:"type:text" json:"desc"`
}

type Contestant struct {
	gorm.Model
	TeamID uint `gorm:"not null" json:"teamID"`
	UserID uint `gorm:"not null" json:"userID"`
	User   User `json:"user"`
}

func CreateContest(contest *Contest) error {
	return global.DB.Create(&contest).Error
}

func UpdateContest(contest *Contest) error {
	return global.DB.Updates(&contest).Error
}

func DeleteContest(id uint) error {
	return global.DB.Where("id = ?", id).Select(clause.Associations).Delete(&Contest{}).Error
}

func CreateTeam(team *Team) error {
	return global.DB.Create(&team).Error
}

func UpdateTeam(team *Team) error {
	return global.DB.Updates(&team).Error
}

func DeleteTeam(id uint) error {
	return global.DB.Select(clause.Associations).Where("id = ?", id).Unscoped().Delete(&Team{}).Error
}

func CreateContestant(contestant *Contestant) error {
	return global.DB.Create(&contestant).Error
}

func UpdateContestant(contestant *Contestant) error {
	return global.DB.Updates(&contestant).Error
}

func DeleteContestant(id uint) error {
	return global.DB.Select(clause.Associations).Where("id = ?", id).Unscoped().Delete(&Contestant{}).Error
}

func GetContestInfo(id uint) (Contest, error) {
	var contest Contest
	err := global.DB.Model(&Contest{}).Preload("Teams").Preload("Teams.Contestants").Preload("Teams.Coach").Preload("Teams.Contestants.User").Where("id = ?", id).First(&contest).Error
	return contest, err
}

func GetAllContestList() ([]Contest, error) {
	var contests []Contest
	err := global.DB.Model(&Contest{}).Preload("Teams").Preload("Teams.Coach").Preload("Teams.Contestants").Preload("Teams.Contestants.User").Find(&contests).Error
	return contests, err
}
