package user

import (
	"context"

	"go-zero-demo/user-api/internal/svc"
	"go-zero-demo/user-api/internal/types"
	genModel "go-zero-demo/user-api/model"

	"github.com/zeromicro/go-zero/core/logx"
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
	user, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, req.UserId)
	if err != nil && err != genModel.ErrNotFound {
		resp.Code = 10000
		resp.Error = "内部错误"
		return &resp, nil
	}
	if user == nil {
		resp.Code = 10001
		resp.Error = "用户不存在"
		return &resp, nil
	}
	resp.Data = &user
	return &resp, nil
}
