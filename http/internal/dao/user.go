package dao

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	"go-framework/http/internal/model"
	"go-framework/http/internal/service"
	"go-framework/http/pkg/common"
	"go-framework/http/pkg/mysql"
)

var UserProvider = wire.NewSet(NewUserDao, wire.Bind(new(service.Iuserdao), new(*Userdao)))

type Userdao struct {
}

func NewUserDao() *Userdao {
	return &Userdao{}
}

func (u *Userdao) GetTableName() string {
	return "fudao_admin"
}

//获取用户信息
func (u *Userdao) GetUserByUid(uid int64) (*model.User, error) {
	modelUserData := &model.User{}
	db := mysql.MysqlSlavePool.Table(u.GetTableName()).Debug().Where("aid= ?", uid).Last(modelUserData)
	if db.Error != nil {
		if db.RecordNotFound() {
			return modelUserData, common.ERR_DAO_NOT_FOUND
		}
		return nil, errors.Wrapf(db.Error, "查询数据库失败 aid=%d", uid)
	}

	return modelUserData, nil
}
