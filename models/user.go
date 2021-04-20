package models

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Pass    string `json:"pass"`
	Created int64  `json:"created"`
	Updated int64  `json:"updated"`
	Deleted int64  `json:"deleted"`
}

func (u *User) TableName() string {
	return "user"
}
