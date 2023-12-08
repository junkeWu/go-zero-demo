package svc

import (
	sqlx2 "github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-demo/user-api/internal/config"
	"go-zero-demo/user-api/internal/middleware"
	genModel "go-zero-demo/user-api/model"
)

type ServiceContext struct {
	Config    config.Config
	Auth      rest.Middleware
	UserModel genModel.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlx := sqlx2.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:    c,
		Auth:      middleware.NewAuthMiddleware().Handle,
		UserModel: genModel.NewUserModel(sqlx, c.Cache),
	}
}
