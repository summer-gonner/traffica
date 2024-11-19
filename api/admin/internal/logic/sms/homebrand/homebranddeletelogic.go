package homebrand

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

// HomeBrandDeleteLogic 首页品牌信息
/*
Author: LiuFeiHua
Date: 2024/5/13 15:53
*/
type HomeBrandDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBrandDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeBrandDeleteLogic {
	return HomeBrandDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// HomeBrandDelete 删除首页品牌信息
// 1.删除sms_home_brand的记录(sms-rpc)
// 2.修改pms_brand记录的状态为不推荐(pms-rpc)
func (l *HomeBrandDeleteLogic) HomeBrandDelete(req types.DeleteHomeBrandReq) (*types.DeleteHomeBrandResp, error) {
	// 1.删除sms_home_brand的记录(sms-rpc)
	_, err := l.svcCtx.HomeBrandService.DeleteHomeBrand(l.ctx, &smsclient.DeleteHomeBrandReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除首页品牌异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除首页品牌失败")
	}

	// 2.修改品牌的推荐状态为不推荐(pms-rpc)
	_, err = l.svcCtx.BrandService.UpdateBrandRecommendStatus(l.ctx, &pmsclient.UpdateBrandRecommendStatusReq{
		Ids:             req.BrandIds,
		RecommendStatus: 0,
	})
	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %+v,修改品牌的推荐状态异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除首页品牌信息失败")
	}

	return &types.DeleteHomeBrandResp{
		Code:    "000000",
		Message: "删除首页品牌成功",
	}, nil
}
