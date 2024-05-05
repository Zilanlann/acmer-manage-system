package contest_service

import "github.com/zilanlann/acmer-manage-system/server/model"

type ContestList struct {
	Contests []model.Contest `json:"list"`
	Total    int             `json:"total"`
}

func (c *ContestList) Get() (err error) {
	c.Contests, err = model.GetAllContestList()
	c.Total = len(c.Contests)
	return
}

type Contest struct {
	model.Contest
}

func (c *Contest) Get() (err error) {
	c.Contest, err = model.GetContestInfo(c.ID)
	return
}

func (c *Contest) Create() (err error) {
	err = model.CreateContest(&c.Contest)
	return
}

func (c *Contest) Update() (err error) {
	err = model.UpdateContest(&c.Contest)
	return
}

func (c *Contest) Delete() (err error) {
	err = model.DeleteContest(c.ID)
	return
}

type Team struct {
	model.Team
}

func (t *Team) Create() (err error) {
	err = model.CreateTeam(&t.Team)
	return
}

func (t *Team) Update() (err error) {
	err = model.UpdateTeam(&t.Team)
	return
}

func (t *Team) Delete() (err error) {
	err = model.DeleteTeam(t.ID)
	return
}

type Contestant struct {
	model.Contestant
}

func (c *Contestant) Create() (err error) {
	err = model.CreateContestant(&c.Contestant)
	return
}

func (c *Contestant) Update() (err error) {
	err = model.UpdateContestant(&c.Contestant)
	return
}

func (c *Contestant) Delete() (err error) {
	err = model.DeleteContestant(c.ID)
	return
}
