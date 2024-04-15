package model

import (
	"github.com/zilanlann/acmer-manage-system/server/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"size:255;not null;unique"`
	RealName  string `gorm:"size:30"`
	StudentID string `gorm:"size:30"`
	Class     string `gorm:"size:30"`
	Avatar    string `gorm:"size:255"`
	Desc      string `gorm:"type:text"`
	Password  string `gorm:"size:255;not null"`
	Role      string `gorm:"type:enum('admin', 'teacher', 'stuCoach', 'acmer');not null"`
}

// AddUser adds a new user to the database with the given username and password.
//
// Parameters:
// - username: the username of the user to be added.
// - password: the password of the user to be added.
//
// Returns:
// - error: an error if there was a problem adding the user to the database.
func AddUser(username, password string) error {
	hash := utils.BcryptHash(password)
	return db.Create(&User{Username: username, Password: hash, Role: "acmer", Avatar: "https://userpic.codeforces.org/no-avatar.jpg"}).Error
}

func DeleteUser(id int) error {
	return db.Where("id = ?", id).Delete(&User{}).Error
}

func CheckUser(username, password string) (id int, err error) {
	if username == "" || password == "" {
		return 0, nil
	}
	var user User
	err = db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil // 用户存在，密码错误
		}
		return 0, err // 数据库错误
	}
	if utils.BcryptCheck(password, user.Password) {
		return int(user.ID), nil // 用户存在，密码正确
	}
	return 0, nil
}

func GetUserInfo(id int) (User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	return user, err
}

func UpdateUserRole(id int, role string) error {
	return db.Model(&User{}).Where("id = ?", id).Update("role", role).Error
}

func UpdateUserAvatar(id int, avatar string) error {
	return db.Model(&User{}).Where("id = ?", id).Update("avatar", avatar).Error
}

func GetAllUsers() ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	return users, err
}

func AddAdmin() error {
	hash := utils.BcryptHash("adminadmin123")
	return db.Create(&User{Username: "admin", Password: hash, Role: "admin"}).Error
}