package oj_service

import (
	"time"

	"github.com/zilanlann/acmer-manage-system/server/model"
)

type OJContestList struct {
	Contests []model.OJContest `json:"list"`
	Total    int               `json:"total"`
}

// GetAllContests method to get all contests
func (o *OJContestList) GetAllContests() error {
	contests, err := model.OJGetContests()
	if err != nil {
		return err
	}
	o.Contests = contests
	o.Total = len(contests)
	return nil
}

// GetContests method to get contests within a specific time range
func (o *OJContestList) GetContests(startTime, endTime time.Time) error {
	contests, err := model.GetContestsWithTime(startTime, endTime)
	if err != nil {
		return err
	}
	o.Contests = contests
	o.Total = len(contests)
	return nil
}

// GetContestsByMonth method to get contests by month
func (o *OJContestList) GetContestsByMonth(year, month int) error {
	contests, err := model.GetContestsByMonth(year, month)
	if err != nil {
		return err
	}
	o.Contests = contests
	o.Total = len(contests)
	return nil
}

type OJSubmissionList struct {
	Submissions []model.OJSubmission `json:"list"`
	Total       int                  `json:"total"`
}

func (o *OJSubmissionList) GetByUser(userId uint) error {
	submissions, err := model.GetAllSubmissionsByUser(userId)
	if err != nil {
		return err
	}
	o.Submissions = submissions
	o.Total = len(submissions)
	return nil
}

func (o *OJSubmissionList) GetByUserAndTime(userId uint, lTime, rTime time.Time) error {
	submissions, err := model.GetUserSubmissionByTime(userId, lTime, rTime)
	if err != nil {
		return err
	}
	o.Submissions = submissions
	o.Total = len(submissions)
	return nil
}

type ProblemTagList struct {
	Tags  []model.ProblemTag `json:"list"`
	Total int                `json:"total"`
}

func (o *ProblemTagList) Get() error {
	tags, err := model.GetAllTags()
	if err != nil {
		return err
	}
	o.Tags = tags
	o.Total = len(tags)
	return nil
}
