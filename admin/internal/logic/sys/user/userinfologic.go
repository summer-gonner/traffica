package user

import (
	"context"
	"encoding/json"
	"github.com/summer-gonner/traffica/admin/internal/svc"
	"github.com/summer-gonner/traffica/admin/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

// UserInfoLogic 获取用户信息
/*
Author: LiuFeiHua
Date: 2023/12/18 14:01
*/
type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoLogic {
	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserInfo 获取用户信息
func (l *UserInfoLogic) UserInfo() (*types.UserInfoResp, error) {
	// 这里的key和生成jwt token时传入的key一致
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()
	logx.Infof("userid%v", userId)
	return &types.UserInfoResp{
		Code:    "000000",
		Message: "获取个人信息成功",
	}, nil
}
