package homerecommendproductservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateRecommendProductSortLogic 人气商品推荐
/*
Author: LiuFeiHua
Date: 2024/5/14 9:33
*/
type UpdateRecommendProductSortLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRecommendProductSortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRecommendProductSortLogic {
	return &UpdateRecommendProductSortLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateRecommendProductSort 修改推荐排序
func (l *UpdateRecommendProductSortLogic) UpdateRecommendProductSort(in *smsclient.UpdateRecommendProductSortReq) (*smsclient.UpdateRecommendProductSortResp, error) {
	q := query.SmsHomeRecommendProduct
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Update(q.Sort, in.Sort)

	if err != nil {
		return nil, err
	}

	return &smsclient.UpdateRecommendProductSortResp{}, nil
}
