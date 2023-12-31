package main

import (
	"context"
	"flag"
	"fmt"

	"go-zero-demo/user-rpc/internal/config"
	"go-zero-demo/user-rpc/internal/server"
	"go-zero-demo/user-rpc/internal/svc"
	"go-zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUserCenterServer(grpcServer, server.NewUserCenterServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	// 添加拦截器
	s.AddUnaryInterceptors(UserRpcServerIntercept)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func UserRpcServerIntercept(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	fmt.Println("handle before")
	resp, err = handler(ctx, req)
	fmt.Println("handle after")
	return resp, err
}
