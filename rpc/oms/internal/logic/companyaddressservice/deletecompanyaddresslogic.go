package companyaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteCompanyAddressLogic 删除公司收发货地址表
/*
Author: LiuFeiHua
Date: 2024/6/12 10:10
*/
type DeleteCompanyAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCompanyAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCompanyAddressLogic {
	return &DeleteCompanyAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteCompanyAddress 删除公司收发货地址表
func (l *DeleteCompanyAddressLogic) DeleteCompanyAddress(in *omsclient.DeleteCompanyAddressReq) (*omsclient.DeleteCompanyAddressResp, error) {
	q := query.OmsCompanyAddress
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &omsclient.DeleteCompanyAddressResp{}, nil
}
