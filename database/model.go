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
	ID          int
	Firstname   string
	Lastname    string
	Information string
}

type Appointment struct {
	ID          int
	Date        time.Time
	Doctor      Contact
	Information string
	Address     string
}

type Medecine struct {
	ID   int
	Name string
	Dose float32
	Unit string
}

type Prescription struct {
	ID          int
	Expire      time.Time
	Information string
	Medecines   []Medecine
}
