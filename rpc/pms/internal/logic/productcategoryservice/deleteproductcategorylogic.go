package productcategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductCategoryLogic 删除产品分类
/*
Author: LiuFeiHua
Date: 2024/6/12 16:55
*/
type DeleteProductCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductCategoryLogic {
	return &DeleteProductCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductCategory 删除产品分类
func (l *DeleteProductCategoryLogic) DeleteProductCategory(in *pmsclient.DeleteProductCategoryReq) (*pmsclient.DeleteProductCategoryResp, error) {
	q := query.PmsProductCategory
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.DeleteProductCategoryResp{}, nil
}
