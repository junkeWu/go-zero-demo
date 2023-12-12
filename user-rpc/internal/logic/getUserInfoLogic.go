package logic

import (
	"context"

	"go-zero-demo/pkg/result"
	"go-zero-demo/user-rpc/internal/svc"
	"go-zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {

	m := map[int64]string{
		1: "zhangSan",
		2: "wang",
	}
	nickname := "unknown"
	if val, ok := m[in.Id]; ok {
		nickname = val
	}
	result := result.Success(nil)
	userInfo := pb.UserModel{
		Id:       in.GetId(),
		Nickname: nickname,
	}
	resp := pb.GetUserInfoResp{
		Code: result.Code,
		Msg:  result.Msg,
		Data: &userInfo,
	}
	return &resp, nil
}
