package user

import (
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/user-api/internal/svc"
	"go-zero-demo/user-api/internal/types"
	"go-zero-demo/user-rpc/pb"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (rsp *types.ResponseBsee, err error) {
	resp := types.ResponseBsee{}
	// user, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, req.UserId)
	// if err != nil && err != genModel.ErrNotFound {
	// 	resp.Code = 10000
	// 	resp.Error = "内部错误"
	// 	logx.Errorf("find one by user id :%v", err)
	// 	return &resp, nil
	// }
	// if user == nil {
	// 	resp.Code = 10001
	// 	resp.Error = "用户不存在"
	// 	return &resp, nil
	// }
	// resp.Data = &user
	atoi, err := strconv.Atoi(req.UserId)
	if err != nil {
		resp.Code = 10000
		resp.Msg = "内部错误"
		logx.Errorf("atoi error:%v", err)
		return &resp, nil
	}
	info, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &pb.GetUserInfoReq{Id: int64(atoi)})
	if err != nil {
		resp.Code = 10000
		resp.Msg = "内部错误"
		logx.Errorf("find one by user id :%v", err)
		return &resp, nil
	}
	if info.Code != 0 {
		resp.Code = info.Code
		resp.Msg = info.Msg
		return &resp, nil
	}
	resp.Data = info.Data
	return &resp, nil
}
