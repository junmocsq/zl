package models

type User struct {
	Uid       int    `json:"uid"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Sex       int    `json:"sex"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Signature string `json:"signature"`
	CreateAt  int64  `json:"create_at"`
	LoginAt   int64  `json:"login_At"`
}
