package user

import "wangqingshui/models"

type Info interface {
	GetUserByUid(uid int) models.User
}
