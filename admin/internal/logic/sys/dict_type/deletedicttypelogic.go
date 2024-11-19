package dict_type

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

// DeleteDictTypeLogic 删除字典类型信息
/*
Author: LiuFeiHua
Date: 2023/12/18 17:17
*/
type DeleteDictTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDictTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteDictTypeLogic {
	return DeleteDictTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteDictType 删除字典类型信息
func (l *DeleteDictTypeLogic) DeleteDictType(req *types.DeleteDictTypeReq) (*types.DeleteDictTypeResp, error) {
	_, err := l.svcCtx.DictTypeService.DeleteDictType(l.ctx, &sysclient.DeleteDictTypeReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据dictId: %+v,删除字典类型异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.DeleteDictTypeResp{
		Code:    "000000",
		Message: "删除字典类型成功",
	}, nil
}
