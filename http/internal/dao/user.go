package dao

import (
	"github.com/google/wire"
	"go-framework/http/internal/model"
	"go-framework/http/internal/service"
)

var UserProvider = wire.NewSet(NewUserDao, wire.Bind(new(service.Iuserdao), new(*Userdao)))

type Userdao struct {
}

func NewUserDao() *Userdao {
	return &Userdao{}
}

func (u *Userdao) GetUserByUid(uid int64) *model.User {
	return &model.User{
		Name: "小明",
		Age:  13,
		Addr: "哈哈哈哈",
	}
}
