package login

import (
	"errors"
	"fmt"
	"wangqingshui/library"
	"wangqingshui/library/cache"
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
	// 设置session uid
	var c cache.Cache = cache.NewCache()
	if res, err := c.SetEx(token, "1", LOGIN_EXPIRE); err != nil && res {
		return
	}
	return "", err
}

func (l *login) SignOut(token string) bool {

}

func (l *login) GetUserInfo(token string) *models.User {

}
