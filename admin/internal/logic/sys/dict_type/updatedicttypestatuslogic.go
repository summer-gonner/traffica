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

// UpdateDictTypeStatusLogic 更新字典类型状态
/*
Author: LiuFeiHua
Date: 2024/5/29 16:56
*/
type UpdateDictTypeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDictTypeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictTypeStatusLogic {
	return &UpdateDictTypeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateDictTypeStatus 更新字典类型状态
func (l *UpdateDictTypeStatusLogic) UpdateDictTypeStatus(req *types.UpdateDictTypeStatusReq) (resp *types.UpdateDictTypeStatusResp, err error) {
	_, err = l.svcCtx.DictTypeService.UpdateDictTypeStatus(l.ctx, &sysclient.UpdateDictTypeStatusReq{
		Ids:        req.DictIds,
		DictStatus: req.DictStatus,
		UpdateBy:   l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新字典类型状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateDictTypeStatusResp{
		Code:    "000000",
		Message: "更新字典类型状态成功",
	}, nil
}
