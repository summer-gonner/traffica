package homerecommendproduct

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateRecommendProductStatusLogic 人气推荐商品
type UpdateRecommendProductStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRecommendProductStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRecommendProductStatusLogic {
	return &UpdateRecommendProductStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateRecommendProductStatus 批量修改推荐状态
// 1.修改sms_home_recommend_product的记录(sms-rpc)
// 2.修改pms_product记录的状态
func (l *UpdateRecommendProductStatusLogic) UpdateRecommendProductStatus(req *types.UpdateRecommendProductStatusReq) (resp *types.UpdateRecommendProductStatusResp, err error) {
	// 1.修改sms_home_recommend_product的记录(sms-rpc)
	_, err = l.svcCtx.HomeRecommendProductService.UpdateHomeRecommendProductStatus(l.ctx, &smsclient.UpdateHomeRecommendProductStatusReq{
		Ids:             req.Ids,
		RecommendStatus: req.RecommendStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量修改人气推荐状态失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("批量修改人气推荐状态失败")
	}

	// 2.修改pms_product记录的状态
	_, err = l.svcCtx.ProductService.UpdateRecommendStatus(l.ctx, &pmsclient.UpdateProductStatusReq{
		Ids:    req.ProductIds,
		Status: req.RecommendStatus,
	})
	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %+v,修改人气推荐状态异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("批量修改人气推荐状态失败")
	}
	return &types.UpdateRecommendProductStatusResp{
		Code:    "000000",
		Message: "批量修改人气推荐状态成功",
	}, nil
}
