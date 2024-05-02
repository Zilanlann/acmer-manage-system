package user_service

import "github.com/zilanlann/acmer-manage-system/server/model"

type UserList struct {
	Users []model.User `json:"list"`
	Total int          `json:"total"`
}

func (ul *UserList) Get() {
	ul.Users, _ = model.GetAllUsersList()
	ul.Total = len(ul.Users)
}

type User struct {
	model.User
}

func (u *User) Add() error {
	u.User.Role = "acmer"
	return model.AddUser(u.User)
}

func (u *User) Update() error {
	return model.UpdateUser(u.User)
}

func (u *User) Delete() error {
	return model.DeleteUser(u.ID)
}

func (u *User) UpdateRole() error {
	return model.UpdateUserRole(u.ID, u.Role)
}

func (u *User) UpdatePassword() error {
	return model.UpdatePassword(u.ID, u.Password)
}
