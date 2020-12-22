package dto

import (
	"go-framework/http/internal/model"
	"go-framework/http/internal/protocol"
)

var User = newUser()

type user struct {
}

func newUser() *user {
	return &user{}
}

//用户数据
func (u *user) GetUser(modelUser *model.User) *protocol.RespGetUsers {
	vo := &protocol.RespGetUsers{
		Name: modelUser.Name,
		Age:  modelUser.Age,
	}

	return vo
}
