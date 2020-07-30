package login

import "wangqingshui/models"

var lo *login

func init() {
	lo = &login{}
}

type Loginor interface {
	SignIn(user models.User) (string, error)
	SignOut(token string) bool
	GetUserInfo(token string) *models.User
}
