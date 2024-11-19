package homerecommendproduct

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateRecommendProductSortLogic 人气推荐商品
/*
Author: LiuFeiHua
Date: 2024/5/14 9:41
*/
type UpdateRecommendProductSortLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRecommendProductSortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRecommendProductSortLogic {
	return &UpdateRecommendProductSortLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateRecommendProductSort 修改推荐排序
func (l *UpdateRecommendProductSortLogic) UpdateRecommendProductSort(req *types.UpdateRecommendProductSortReq) (resp *types.UpdateRecommendProductSortResp, err error) {
	_, err = l.svcCtx.HomeRecommendProductService.UpdateRecommendProductSort(l.ctx, &smsclient.UpdateRecommendProductSortReq{
		Id:   req.Id,
		Sort: req.Sort,
	})

	if err != nil {
		logc.Errorf(l.ctx, "修改推荐排序失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("修改推荐排序失败")
	}

	return &types.UpdateRecommendProductSortResp{
		Code:    "000000",
		Message: "修改推荐排序成功",
	}, nil
}
