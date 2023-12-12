package logic

import (
	"context"

	"go-zero-demo/pkg/result"
	"go-zero-demo/user-rpc/internal/genModel"
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
	result := result.Success(nil)
	resp := pb.GetUserInfoResp{
		Code: result.Code,
		Msg:  result.Msg,
	}
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.GetId())
	if err != nil && err != genModel.ErrNotFound {
		resp.Code = 10000
		resp.Msg = "内部错误"
		logx.Errorf("find one by user id :%v", err)
		return &resp, nil
	}
	if user == nil {
		resp.Code = 10001
		resp.Msg = "用户不存在"
		return &resp, nil
	}
	userInfo := pb.UserModel{
		Id:       in.GetId(),
		Nickname: user.Username,
	}
	resp = pb.GetUserInfoResp{
		Code: result.Code,
		Msg:  result.Msg,
		Data: &userInfo,
	}
	return &resp, nil
}
