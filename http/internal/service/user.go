package service

import (
	"go-framework/http/api"
	"go-framework/http/internal/model"
)

type Iuserdao interface {
	GetUserByUid(uid int64) *model.User
}

type User struct {
	userDao Iuserdao
}

func NewUserService(userDao Iuserdao) *User {
	return &User{
		userDao: userDao,
	}
}

func (u *User) GetUser(uid int64) *api.RespGetUsers {
	modelUser := u.userDao.GetUserByUid(uid)

	//po转vo
	return &api.RespGetUsers{
		Name: modelUser.Name,
		Age:  modelUser.Age,
	}
}
