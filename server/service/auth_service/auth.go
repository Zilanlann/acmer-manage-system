package auth_service

import (
	"github.com/zilanlann/acmer-manage-system/server/model"
)

type Auth struct {
	UserId    int
	Username  string
	Password  string
	Realname  string
	Email     string
	CfHandle  string
	AtcHandle string
	Role      string
	Avatar    string
}

func (a *Auth) Check() (bool, error) {
	id, err := model.CheckUser(a.Username, a.Password)
	if err != nil {
		return false, err
	}
	a.UserId = id
	user, err := model.GetUserInfo(id)
	if err != nil {
		return false, err
	}
	a.Username = user.Username
	a.Role = user.Role
	a.Avatar = user.Avatar
	return true, nil
}

func (a *Auth) Add() error {
	u := model.User{
		Username:  a.Username,
		RealName:  a.Realname,
		Password:  a.Password,
		Email:     a.Email,
		CFHandle:  a.CfHandle,
		ATCHandle: a.AtcHandle,
		Avatar:    a.Avatar,
		Role:      "acmer",
	}

	return model.CreateUser(u)
}
