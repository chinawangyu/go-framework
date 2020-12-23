// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/google/wire"
	"go-framework/http/internal/dao"
	"go-framework/http/internal/service"
)

func NewUserService() *service.User {
	wire.Build(service.NewUserService, dao.UserProvider)
	return nil
}
