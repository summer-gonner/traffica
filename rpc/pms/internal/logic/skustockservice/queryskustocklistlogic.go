package skustockservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QuerySkuStockListLogic 查询sku的库存列表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:12
*/
type QuerySkuStockListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySkuStockListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySkuStockListLogic {
	return &QuerySkuStockListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QuerySkuStockList 查询sku的库存列表
func (l *QuerySkuStockListLogic) QuerySkuStockList(in *pmsclient.QuerySkuStockListReq) (*pmsclient.QuerySkuStockListResp, error) {
	stock := query.PmsSkuStock
	q := stock.WithContext(l.ctx).Where(stock.ProductID.Eq(in.ProductId))
	if len(in.SkuCode) > 0 {
		q = q.Where(stock.SkuCode.Eq(in.SkuCode))
	}
	result, err := q.Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询库存列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.SkuStockListData
	for _, item := range result {

		list = append(list, &pmsclient.SkuStockListData{
			Id:             item.ID,
			ProductId:      item.ProductID,
			SkuCode:        item.SkuCode,
			Price:          item.Price,
			Stock:          item.Stock,
			LowStock:       item.LowStock,
			Pic:            item.Pic,
			Sale:           item.Sale,
			PromotionPrice: item.PromotionPrice,
			LockStock:      item.LockStock,
			SpData:         item.SpData,
		})
	}

	return &pmsclient.QuerySkuStockListResp{
		Total: 0,
		List:  list,
	}, nil

}
