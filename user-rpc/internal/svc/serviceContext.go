package svc

import (
	sqlx2 "github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/user-rpc/internal/config"
	"go-zero-demo/user-rpc/internal/genModel"
)

type ServiceContext struct {
	Config    config.Config
	UserModel genModel.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlx := sqlx2.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: genModel.NewUserModel(sqlx, c.Cache),
	}
}
