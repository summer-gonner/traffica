package log

import (
	"context"
	"github.com/summer-gonner/traffica/admin/internal/common/errorx"
	"github.com/summer-gonner/traffica/admin/internal/svc"
	"github.com/summer-gonner/traffica/admin/internal/types"
	"github.com/summer-gonner/traffica/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteLoginLogLogic 删除登录日志
/*
Author: LiuFeiHua
Date: 2023/12/18 17:19
*/
type DeleteLoginLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteLoginLogLogic {
	return DeleteLoginLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteLoginLog 删除登录日志
func (l *DeleteLoginLogLogic) DeleteLoginLog(req *types.DeleteLoginLogReq) (*types.DeleteLoginLogResp, error) {
	_, err := l.svcCtx.LoginLogService.DeleteLoginLog(l.ctx, &sysclient.DeleteLoginLogReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据LoginLog ids: %+v,删除登录日志异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.DeleteLoginLogResp{
		Code:    "000000",
		Message: "删除登录日志成功",
	}, nil
}
