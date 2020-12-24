package service

import (
	"errors"
	"go-framework/http/api"
	"go-framework/http/internal/model"
	"go-framework/http/pkg/common"
)

type Iuserdao interface {
	GetUserByUid(uid int64) (*model.User, error)
}

type User struct {
	userDao Iuserdao
}

func NewUserService(userDao Iuserdao) *User {
	return &User{
		userDao: userDao,
	}
}

//获取用户信息
func (u *User) GetUser(uid int64) (*api.RespGetUsers, error) {
	modelUser, err := u.userDao.GetUserByUid(uid)
	if errors.Is(err, common.ERR_DAO_NOT_FOUND) {
		return &api.RespGetUsers{Name: "默认姓名", Age: 0}, nil
	}

	if err != nil {
		return nil, err
	}

	return &api.RespGetUsers{Name: modelUser.UserName, Age: modelUser.Aid}, nil
}
