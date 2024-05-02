package model

import (
	"fmt"

	"github.com/zilanlann/acmer-manage-system/server/global"
	"github.com/zilanlann/acmer-manage-system/server/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"size:255;not null;unique" json:"username" binding:"required"`
	RealName  string `gorm:"size:30" json:"realname" binding:"required"`
	Email     string `gorm:"size:255;not null;unique" json:"email" binding:"required"`
	Sex       bool   `json:"sex"`
	StudentID string `gorm:"size:30" json:"studentID"`
	Class     string `gorm:"size:30" json:"class"`
	Phone     string `gorm:"size:30" json:"phone"`
	CFHandle  string `gorm:"size:255" json:"cfHandle"`
	ATCHandle string `gorm:"size:255" json:"atcHandle"`
	Avatar    string `gorm:"size:255" json:"avatar"`
	Desc      string `gorm:"type:text" json:"desc"`
	Password  string `gorm:"size:255;not null" json:"password"`
	Role      string `gorm:"size:30;not null" json:"role"`
}

func AddUser(user User) error {
	fmt.Printf("user.Password: %v\n", user.Password)
	user.Password = utils.BcryptHash(user.Password)
	fmt.Printf("user.Password: %v\n", user.Password)
	return db.Create(&user).Error
}

func DeleteUser(id uint) error {
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
			return 0, nil // 用户不存在
		}
		return 0, err // 数据库错误
	}
	if utils.BcryptCheck(password, user.Password) {
		return int(user.ID), nil // 用户存在，密码正确
	}
	global.LOG.Warn("password is wrong")
	return 0, nil
}

func GetUserInfo(id int) (User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	return user, err
}

func UpdateUser(user User) error {
	return db.Select("username", "real_name", "email", "sex", "student_id", "class", "phone", "cf_handle", "atc_handle").Save(&user).Error
}

func UpdateUserRole(id uint, role string) error {
	return db.Model(&User{}).Where("id = ?", id).Update("role", role).Error
}

func UpdateUserAvatar(id uint, avatar string) error {
	return db.Model(&User{}).Where("id = ?", id).Update("avatar", avatar).Error
}

func UpdatePassword(id uint, password string) error {
	return db.Model(&User{}).Where("id = ?", id).Update("password", utils.BcryptHash(password)).Error
}

func GetACMersList() ([]User, error) {
	var users []User
	err := db.Model(&User{}).Where("role = ?", "acmer").Find(&users).Error
	return users, err
}

func GetAllUsersList() ([]User, error) {
	var users []User
	err := db.Model(&User{}).Find(&users).Error
	return users, err
}

func AddAdmin() error {
	hash := utils.BcryptHash("adminadmin123")
	return db.Create(&User{Username: "admin", Password: hash, Avatar: "https://userpic.codeforces.org/no-avatar.jpg", Role: "admin"}).Error
}
