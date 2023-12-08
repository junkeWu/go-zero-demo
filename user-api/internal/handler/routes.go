// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "go-zero-demo/user-api/internal/handler/user"
	"go-zero-demo/user-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/info",
					Handler: user.UserInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/userapi/v1"),
	)
}
