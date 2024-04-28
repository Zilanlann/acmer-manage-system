package model

import (
	"github.com/zilanlann/acmer-manage-system/server/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"size:255;not null;unique"`
	RealName  string `gorm:"size:30"`
	Email     string `gorm:"size:255;not null;unique"`
	StudentID string `gorm:"size:30"`
	Class     string `gorm:"size:30"`
	CFHandle  string `gorm:"size:255"`
	ATCHandle string `gorm:"size:255"`
	Avatar    string `gorm:"size:255"`
	Desc      string `gorm:"type:text"`
	Password  string `gorm:"size:255;not null"`
	Role      string `gorm:"size:30;not null"`
}

func AddUser(user User) error {
	user.Password = utils.BcryptHash(user.Password)
	return db.Create(&user).Error
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

func UpdateUser(user User) error {
	return db.Save(&user).Error
}

func UpdateUserRole(id int, role string) error {
	return db.Model(&User{}).Where("id = ?", id).Update("role", role).Error
}

func UpdateUserAvatar(id int, avatar string) error {
	return db.Model(&User{}).Where("id = ?", id).Update("avatar", avatar).Error
}

func GetAllNormalUsers() ([]User, error) {
	var users []User
	err := db.Model(&User{}).Where("role = ?", "acmer").Find(&users).Error
	return users, err
}

func AddAdmin() error {
	hash := utils.BcryptHash("adminadmin123")
	return db.Create(&User{Username: "admin", Password: hash, Avatar: "https://userpic.codeforces.org/no-avatar.jpg", Role: "admin"}).Error
}
