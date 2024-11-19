package user

import (
	"context"
	"github.com/summer-gonner/traffica/admin/internal/common/errorx"
	"github.com/summer-gonner/traffica/admin/internal/svc"
	"github.com/summer-gonner/traffica/admin/internal/types"
	"github.com/summer-gonner/traffica/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// ReSetPasswordLogic 重置用户密码
/*
Author: LiuFeiHua
Date: 2023/12/18 13:54
*/
type ReSetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReSetPasswordLogic {
	return ReSetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ReSetPassword 重置用户密码
func (l *ReSetPasswordLogic) ReSetPassword(req *types.ReSetPasswordReq) (*types.ReSetPasswordResp, error) {

	_, err := l.svcCtx.UserService.ReSetPassword(l.ctx, &sysclient.ReSetPasswordReq{
		Id:       req.UserId,
		UpdateBy: l.ctx.Value("userName").(string),
	})

	if err != nil {
		return nil, errorx.NewDefaultError("重置用户密码异常")
	}

	return &types.ReSetPasswordResp{
		Code:    "000000",
		Message: "重置用户密码成功",
	}, nil
}
