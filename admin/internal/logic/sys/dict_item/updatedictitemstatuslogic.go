package dict_item

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

// UpdateDictItemStatusLogic 更新字典数据状态
type UpdateDictItemStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDictItemStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDictItemStatusLogic {
	return &UpdateDictItemStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateDictItemStatus 更新字典数据状态
func (l *UpdateDictItemStatusLogic) UpdateDictItemStatus(req *types.UpdateDictItemStatusReq) (resp *types.UpdateDictItemStatusResp, err error) {
	_, err = l.svcCtx.DictItemService.UpdateDictItemStatus(l.ctx, &sysclient.UpdateDictItemStatusReq{
		Ids:        req.DictIds,
		DictStatus: req.DictStatus,
		UpdateBy:   l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新字典数据状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateDictItemStatusResp{
		Code:    "000000",
		Message: "更新字典数据状态成功",
	}, nil
}
