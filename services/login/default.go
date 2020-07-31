package login

import "wangqingshui/models"

const LOGIN_EXPIRE = 86400

var lo *login

func init() {
	lo = &login{}
}

type Loginor interface {
	SignIn(user models.User) (string, error)
	SignOut(token string) bool
	GetUserInfo(token string) *models.User
}
