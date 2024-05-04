package model

import (
	"github.com/zilanlann/acmer-manage-system/server/global"
	"gorm.io/gorm"
)

type Contest struct {
	gorm.Model
	Name  string `gorm:"not null;unique" json:"name"`
	Desc  string `gorm:"type:text" json:"desc"`
	Teams []Team `json:"teams"`
}

type Team struct {
	gorm.Model
	ContestID   uint
	Name        string       `gorm:"not null;unique" json:"name"`
	CoachID     uint         `gorm:"foreignKey:UserID" json:"coach_id"`
	Coach       User         `json:"coach"`
	Contestants []Contestant `json:"contestants"`
}

type Contestant struct {
	gorm.Model
	TeamID uint
	UserID uint
	User   User `json:"user"`
}

func CreateContest(contest *Contest) error {
	return global.DB.Create(contest).Error
}

func CreateTeam(contest *Contest, team *Team) error {
	return global.DB.Create(team).Error
}

func CreateContestant(team *Team, contestant *Contestant) error {
	return global.DB.Create(contestant).Error
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
