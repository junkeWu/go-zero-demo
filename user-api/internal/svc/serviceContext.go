package svc

import (
	sqlx2 "github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-demo/user-api/internal/config"
	"go-zero-demo/user-api/internal/middleware"
	genModel "go-zero-demo/user-api/model"
	"go-zero-demo/user-rpc/usercenter"
)

type ServiceContext struct {
	Config        config.Config
	Auth          rest.Middleware
	UserModel     genModel.UserModel
	UserRpcClient usercenter.UserCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlx := sqlx2.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:        c,
		Auth:          middleware.NewAuthMiddleware().Handle,
		UserModel:     genModel.NewUserModel(sqlx, c.Cache),
		UserRpcClient: usercenter.NewUserCenter(zrpc.MustNewClient(c.RpcClientConfig)),
	}
}
