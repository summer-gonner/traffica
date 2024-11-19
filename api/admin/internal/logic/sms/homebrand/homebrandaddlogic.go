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

// HomeBrandAddLogic 首页品牌信息
/*
Author: LiuFeiHua
Date: 2024/5/13 15:53
*/
type HomeBrandAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBrandAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeBrandAddLogic {
	return HomeBrandAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// HomeBrandAdd 添加首页品牌信息
// 1.根据brandIds查询品牌信息(pms-rpc)
// 2.添加首页品牌记录(sms-rpc)
// 3.修改品牌的推荐状态为推荐(pms-rpc)
func (l *HomeBrandAddLogic) HomeBrandAdd(req types.AddHomeBrandReq) (*types.AddHomeBrandResp, error) {
	// 1.根据brandIds查询品牌信息(pms-rpc)
	brandListResp, _ := l.svcCtx.BrandService.QueryBrandListByIds(l.ctx, &pmsclient.QueryBrandListByIdsReq{Ids: req.BrandIds})

	var list []*smsclient.HomeBrandAddData

	for _, item := range brandListResp.List {
		list = append(list, &smsclient.HomeBrandAddData{
			BrandId:         item.Id,
			BrandName:       item.Name,
			RecommendStatus: item.ShowStatus,
			Sort:            int32(item.Id),
		})
	}

	// 2.添加首页品牌记录(sms-rpc)
	_, err := l.svcCtx.HomeBrandService.AddHomeBrand(l.ctx, &smsclient.AddHomeBrandReq{
		BrandAddData: list,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加首页品牌信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加首页品牌失败")
	}

	// 3.修改品牌的推荐状态为推荐(pms-rpc)
	_, err = l.svcCtx.BrandService.UpdateBrandRecommendStatus(l.ctx, &pmsclient.UpdateBrandRecommendStatusReq{
		Ids:             req.BrandIds,
		RecommendStatus: 1,
	})
	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %+v,修改品牌的推荐状态异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加首页品牌信息失败")
	}

	return &types.AddHomeBrandResp{
		Code:    "000000",
		Message: "添加首页品牌成功",
	}, nil
}
