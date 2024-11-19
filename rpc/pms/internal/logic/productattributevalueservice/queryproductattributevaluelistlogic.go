package productattributevalueservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductAttributeValueListLogic 查询存储产品参数信息的表列表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:50
*/
type QueryProductAttributeValueListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductAttributeValueListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductAttributeValueListLogic {
	return &QueryProductAttributeValueListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductAttributeValueList 查询存储产品参数信息的表列表
func (l *QueryProductAttributeValueListLogic) QueryProductAttributeValueList(in *pmsclient.QueryProductAttributeValueListReq) (*pmsclient.QueryProductAttributeValueListResp, error) {
	q := query.PmsProductAttributeValue
	result, err := q.WithContext(l.ctx).Where(q.ProductID.Eq(in.ProductId)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "查询产品参数列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*pmsclient.ProductAttributeValueListData
	for _, item := range result {

		list = append(list, &pmsclient.ProductAttributeValueListData{
			Id:                 item.ID,
			ProductId:          item.ProductID,
			ProductAttributeId: item.ProductAttributeID,
			Value:              *item.Value,
		})
	}

	return &pmsclient.QueryProductAttributeValueListResp{
		Total: 0,
		List:  list,
	}, nil

}
