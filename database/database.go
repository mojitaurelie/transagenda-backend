package database

import (
	"fmt"
	"github.com/transagenda-back/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
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
	), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,  // Slow SQL threshold
				LogLevel:                  logger.Error, // Log level
				IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,         // Enable color
			},
		),
	})
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

func UserById(userId int) (*User, error) {
	var user *User
	err := db.Model(User{}).Where(User{ID: userId}).First(&user).Error
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

func AppointmentsByUserId(userId int) ([]*Appointment, error) {
	var appointments []*Appointment
	err := db.Preload("Contact").Model(Appointment{}).Where(Appointment{UserId: userId}).Find(&appointments).Error
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func PrescriptionsByUserId(userId int) ([]*Prescription, error) {
	var prescriptions []*Prescription
	err := db.Preload("Medecines").Model(Prescription{}).Where(Appointment{UserId: userId}).Find(&prescriptions).Error
	if err != nil {
		return nil, err
	}
	return prescriptions, nil
}

func ContactsByUserId(userId int) ([]*Contact, error) {
	var contacts []*Contact
	err := db.Model(Contact{}).Where(Contact{UserId: userId}).Find(&contacts).Error
	if err != nil {
		return nil, err
	}
	return contacts, nil
}
