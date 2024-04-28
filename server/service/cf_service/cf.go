package cf_service

import "github.com/zilanlann/acmer-manage-system/server/pkg/cf"

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
	userInfo, _ := cf.GetUserInfo(u.CFHandle)
	u.CfRating = userInfo.Rating
	u.WeeklyRating, u.MonthlyRating, _ = cf.GetWMRating(u.CFHandle)
	u.WeeklyRating = u.CfRating - u.WeeklyRating
	u.MonthlyRating = u.CfRating - u.MonthlyRating
}

func GetAvatar(handle string) string {
	cf.RefreshUserInfos([]string{handle})
	user, _ := cf.GetUserInfo(handle)
	return user.Avatar
}