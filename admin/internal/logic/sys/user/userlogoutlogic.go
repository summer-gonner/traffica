package user

import (
	"context"
	"github.com/summer-gonner/traffica/admin/internal/svc"
	"github.com/summer-gonner/traffica/admin/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogoutLogic {
	return &UserLogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogoutLogic) UserLogout() (resp *types.LogoutResp, err error) {

	// 构造响应
	return &types.LogoutResp{
		Code:    "000000",
		Message: "退出登录成功",
	}, nil
}
