package auth_service

import (
	"github.com/zilanlann/acmer-manage-system/server/models"
)

type Auth struct {
	UserId   int
	Username string
	Password string
	RealName string
	Desc     string
	Role     string
}

func (a *Auth) Check() (bool, error) {
	id, err := models.CheckUser(a.Username, a.Password)
	if err != nil {
		return false, err
	}
	a.UserId = id
	user := models.GetUserInfo(id)
	a.Username = user.Username
	a.RealName = user.RealName
	a.Desc = user.Desc
	a.Role = user.Role
	return true, nil
}
