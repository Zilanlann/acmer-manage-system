package contest_service

import "github.com/zilanlann/acmer-manage-system/server/model"

type ContestList struct {
	Contests []model.Contest `json:"list"`
	Total    int             `json:"total"`
}

type Contest struct {
	model.Contest
}

func (c *ContestList) Get() (err error) {
	c.Contests, err = model.GetAllContestList()
	c.Total = len(c.Contests)
	return
}

func (c *Contest) Get() (err error) {
	c.Contest, err = model.GetContestInfo(c.ID)
	return
}
