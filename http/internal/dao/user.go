package dao

import "go-framework/http/internal/model"

type Iuserdao interface {
	GetUserByUid() *model.User
}

type userdao struct {
}

func NewUserDao() *userdao {
	return &userdao{}
}

func (u *userdao) GetUserByUid() *model.User {
	return &model.User{
		Name: "小明",
		Age:  13,
		Addr: "哈哈哈哈",
	}
}
