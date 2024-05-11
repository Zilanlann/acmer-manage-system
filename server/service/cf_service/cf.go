package cf_service

import (
	"github.com/zilanlann/acmer-manage-system/server/global"
	"github.com/zilanlann/acmer-manage-system/server/pkg/cf"
)

type UserStatus struct {
	UserName      string `json:"userName"`
	RealName      string `json:"realName"`
	CFHandle      string `json:"cfHandle"`
	CfRating      int    `json:"cfRating"`
	WeeklyRating  int    `json:"weeklyRating"`
	MonthlyRating int    `json:"monthlyRating"`
	WeeklyActive  int    `json:"weeklyActive"`
	MonthlyActive int    `json:"monthlyActive"`
}

func (u *UserStatus) GetUserStatus() {
	cf.RefreshRatingChange(u.CFHandle)
	userInfo, err := cf.GetUserInfo(u.CFHandle)
	if err != nil {
		global.LOG.Error(err.Error())
	}
	u.CfRating = userInfo.Rating
	u.WeeklyRating, u.MonthlyRating, _ = cf.GetWMRating(u.CFHandle)
	u.WeeklyRating = u.CfRating - u.WeeklyRating
	u.MonthlyRating = u.CfRating - u.MonthlyRating
}

func GetAvatar(handle string) string {
	if err := cf.RefreshUserInfos([]string{handle}); err != nil {
		global.LOG.Error(err.Error())
	}
	var user cf.User
	user, err := cf.GetUserInfo(handle)
	if err != nil {
		global.LOG.Error(err.Error())
	}
	return user.Avatar
}
