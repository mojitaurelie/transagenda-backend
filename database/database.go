package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

const username string = "TRANS_DB_USERNAME"
const password string = "TRANS_DB_PASSWORD"
const host string = "TRANS_DB_HOST"
const port string = "TRANS_DB_PORT"

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.Open(
		fmt.Sprintf("%s:%s@tcp(%s:%s)/transagenda?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv(username),
			os.Getenv(password),
			os.Getenv(host),
			os.Getenv(port)),
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
