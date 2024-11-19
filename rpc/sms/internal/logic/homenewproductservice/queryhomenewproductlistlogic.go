package homenewproductservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryHomeNewProductListLogic 查询新鲜好物表列表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:54
*/
type QueryHomeNewProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHomeNewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHomeNewProductListLogic {
	return &QueryHomeNewProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryHomeNewProductList 查询新鲜好物表列表
func (l *QueryHomeNewProductListLogic) QueryHomeNewProductList(in *smsclient.QueryHomeNewProductListReq) (*smsclient.QueryHomeNewProductListResp, error) {
	q := query.SmsHomeNewProduct.WithContext(l.ctx)
	if len(in.ProductName) > 0 {
		q = q.Where(query.SmsHomeNewProduct.ProductName.Like("%" + in.ProductName + "%"))
	}

	if in.RecommendStatus != 2 {
		q = q.Where(query.SmsHomeNewProduct.RecommendStatus.Eq(in.RecommendStatus))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询首页新鲜好物列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.HomeNewProductListData
	for _, item := range result {

		list = append(list, &smsclient.HomeNewProductListData{
			Id:              item.ID,
			ProductId:       item.ProductID,
			ProductName:     item.ProductName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	return &smsclient.QueryHomeNewProductListResp{
		Total: count,
		List:  list,
	}, nil

}
