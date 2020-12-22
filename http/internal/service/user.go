package service

import (
	"go-framework/http/internal/dao"
	"go-framework/http/internal/dto"
	"go-framework/http/internal/protocol"
)

type user struct {
	userDao dao.Iuserdao
}

func NewUserService(userDao dao.Iuserdao) *user {
	return &user{
		userDao: userDao,
	}
}

func (u *user) GetUser() *protocol.RespGetUsers {
	modelUser := u.userDao.GetUserByUid()

	//po转vo
	return dto.User.GetUser(modelUser)
}
