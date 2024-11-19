package menu

import (
	"context"
	"github.com/summmer-gonner/traffica/admin/internal/common/errorx"
	"github.com/summmer-gonner/traffica/admin/internal/svc"
	"github.com/summmer-gonner/traffica/admin/internal/types"
	"github.com/summmer-gonner/traffica/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMenuLogic 删除菜单
/*
Author: LiuFeiHua
Date: 2023/12/18 15:25
*/
type DeleteMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteMenuLogic {
	return DeleteMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteMenu 删除菜单
func (l *DeleteMenuLogic) DeleteMenu(req *types.DeleteMenuReq) (*types.DeleteMenuResp, error) {
	if _, err := l.svcCtx.MenuService.DeleteMenu(l.ctx, &sysclient.DeleteMenuReq{
		Ids: req.Ids,
	}); err != nil {
		logc.Errorf(l.ctx, "根据menuId: %+v,删除菜单异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.DeleteMenuResp{
		Code:    "000000",
		Message: "删除菜单成功",
	}, nil
}
