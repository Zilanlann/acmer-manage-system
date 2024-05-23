package cf_service

import (
	"github.com/zilanlann/acmer-manage-system/server/global"
	"github.com/zilanlann/acmer-manage-system/server/pkg/cf"
)

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
