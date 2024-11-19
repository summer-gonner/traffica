package homebrandservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteHomeBrandLogic 删除首页推荐品牌表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:53
*/
type DeleteHomeBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHomeBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHomeBrandLogic {
	return &DeleteHomeBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteHomeBrand 删除首页推荐品牌表
func (l *DeleteHomeBrandLogic) DeleteHomeBrand(in *smsclient.DeleteHomeBrandReq) (*smsclient.DeleteHomeBrandResp, error) {
	q := query.SmsHomeBrand
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &smsclient.DeleteHomeBrandResp{}, nil
}
