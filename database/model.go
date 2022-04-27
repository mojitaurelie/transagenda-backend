package database

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  []byte `json:"-"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Pronouns  int    `json:"pronouns"`
}
