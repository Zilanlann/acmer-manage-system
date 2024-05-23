package user_service

import "github.com/zilanlann/acmer-manage-system/server/model"

type UserStatusList struct {
	List  []UserStatus `json:"list"`
	Total int          `json:"total"`
}

type UserStatus struct {
	model.UserStatus
	RealName     string     `json:"realname"`
	TagCountList []TagCount `json:"tag_count_list"`
}

type TagCount struct {
	Tag   string `json:"name"`
	Count int    `json:"value"`
}

func (ul *UserStatusList) Get() error {
	statusList, err := model.GetAllUserStatusList()
	if err != nil {
		return err
	}
	for _, status := range statusList {
		userInfo, err := model.GetUserInfo(status.UserID)
		if err != nil {
			return err
		}
		status.WeeklyActive, status.MonthlyActive, err = model.CalcUserActiveByUser(status.UserID)
		status.WeeklyActive /= status.CFRating / 10
		status.MonthlyActive /= status.CFRating / 10
		if err != nil {
			return err
		}
		ul.List = append(ul.List, UserStatus{
			UserStatus: status,
			RealName:   userInfo.RealName,
		})
		tagCountList, err := model.GetTagCountsByUser(status.UserID)
		if err != nil {
			return err
		}
		ul.List[len(ul.List)-1].TagCountList = make([]TagCount, 0, len(tagCountList))
		for tag, count := range tagCountList {
			ul.List[len(ul.List)-1].TagCountList = append(ul.List[len(ul.List)-1].TagCountList, TagCount{
				Tag:   tag,
				Count: count,
			})
		}
	}
	ul.Total = len(statusList)
	return nil
}
