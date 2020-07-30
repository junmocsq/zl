package login

import (
	"errors"
	"fmt"
	"wangqingshui/library"
	"wangqingshui/models"
)

type login struct {
}

func NewLogin() *login {
	return lo
}

func (l *login) SignIn(user models.User) (token string, err error) {
	if user.Uid == 0 || user.Name == "" {
		err = errors.New("test")
		return
	}
	prefix := fmt.Sprintf("user%07d", user.Uid)
	token = library.NewTools().CreateToken(prefix)
	return
}

func (l *login) SignOut(token string) bool {

}

func (l *login) GetUserInfo(token string) *models.User {

}
