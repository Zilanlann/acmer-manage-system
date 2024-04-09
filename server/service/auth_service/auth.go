package auth_service

import (
	"github.com/zilanlann/acmer-manage-system/server/model"
)

type Auth struct {
	UserId   int
	Username string
	Password string
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
	a.Role = user.Role
	return true, nil
}

func (a *Auth) Add() error {
	return model.AddUser(a.Username, a.Password)
}

func (a *Auth) Delete() error {
	return model.DeleteUser(a.UserId)
}

// func (a *Auth) Update() error {
// 	return model.UpdateUser(a.UserId, a.Username, a.Role)
// }