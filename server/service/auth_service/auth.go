package auth_service

import (
	"github.com/zilanlann/acmer-manage-system/server/model"
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
	a.RealName = user.RealName
	a.Desc = user.Desc
	a.Role = user.Role
	return true, nil
}

func (a *Auth) Add() error {
	return model.AddUser(a.Username, a.Password)
}
