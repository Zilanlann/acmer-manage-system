package models

import (
	"github.com/jinzhu/gorm"
	"github.com/zilanlann/acmer-manage-system/server/utils"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null"`
	RealName string `gorm:"size:255"`
	Avatar   string `gorm:"size:255"`
	Desc     string `gorm:"type:text"`
	Password string `gorm:"size:255;not null"`
	// HomePath string `gorm:"size:255"`
	Role string `gorm:"type:enum('Super Admin', 'Teacher', 'Student');not null"`
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
	return db.Create(&User{Username: username, Password: hash}).Error
}

func CheckUser(username, password string) (int, error) {
	if username == "" || password == "" {
		return 0, nil
	}
	var user User
	err := db.Select(user.ID, user.Password).Where("username = ?", username).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
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
