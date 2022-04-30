package database

import "time"

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  []byte `json:"-"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Pronouns  int    `json:"pronouns"`
}

type Contact struct {
	ID          int    `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Information string `json:"information"`
	UserId      int    `json:"-"`
}

type Appointment struct {
	ID           int       `json:"id"`
	Date         time.Time `json:"date"`
	ContactRefer int       `json:"-"`
	Contact      Contact   `json:"contact" gorm:"foreignKey:ContactRefer"`
	Information  string    `json:"information"`
	Address      string    `json:"address"`
	UserId       int       `json:"-"`
}

type Medecine struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Dose         float32 `json:"dose"`
	Unit         string  `json:"unit"`
	Prescription int     `json:"-"`
}

type Prescription struct {
	ID          int        `json:"id"`
	Expire      time.Time  `json:"expire"`
	Information string     `json:"information"`
	UserId      int        `json:"-"`
	Medecines   []Medecine `json:"medecines" gorm:"foreignKey:Prescription"`
}
