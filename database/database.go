package database

import (
	"fmt"
	"github.com/transagenda-back/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func init() {
	dbConfig := config.Database()
	var err error
	db, err = gorm.Open(mysql.Open(
		fmt.Sprintf("%s:%s@tcp(%s:%d)/transagenda?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.Username,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port),
	), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func UserByUsername(username string) (*User, error) {
	var user *User
	err := db.Model(User{}).Where(User{Username: username}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AddUser(username, firstname, lastname string, password []byte, pronouns int) error {
	user := &User{
		Username:  username,
		Password:  password,
		Firstname: firstname,
		Lastname:  lastname,
		Pronouns:  pronouns,
	}
	return db.Save(user).Error
}
