package homenewproductservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateHomeNewProductStatusLogic 更新新鲜好物表状态
/*
Author: LiuFeiHua
Date: 2024/6/12 17:55
*/
type UpdateHomeNewProductStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHomeNewProductStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomeNewProductStatusLogic {
	return &UpdateHomeNewProductStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateHomeNewProductStatus 更新新鲜好物表状态
func (l *UpdateHomeNewProductStatusLogic) UpdateHomeNewProductStatus(in *smsclient.UpdateHomeNewProductStatusReq) (*smsclient.UpdateHomeNewProductStatusResp, error) {
	q := query.SmsHomeNewProduct
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.RecommendStatus, in.RecommendStatus)

	if err != nil {
		return nil, err
	}

	return &smsclient.UpdateHomeNewProductStatusResp{}, nil
}
