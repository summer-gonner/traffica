package category

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductCategoryListLogic 商品分类
/*
Author: LiuFeiHua
Date: 2024/5/15 11:17
*/
type ProductCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCategoryListLogic {
	return ProductCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductCategoryList 查询商品分类
func (l *ProductCategoryListLogic) ProductCategoryList(req types.ListProductCategoryReq) (*types.ListProductCategoryResp, error) {
	resp, err := l.svcCtx.ProductCategoryService.QueryProductCategoryList(l.ctx, &pmsclient.QueryProductCategoryListReq{
		PageNum:    req.Current,
		PageSize:   req.PageSize,
		Name:       strings.TrimSpace(req.Name),
		ParentId:   req.ParentId,
		ShowStatus: req.ShowStatus, // 显示状态：0->不显示；1->显示
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询商品分类列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询商品分类失败")
	}

	var list []*types.ListProductCategoryData

	for _, item := range resp.List {
		list = append(list, &types.ListProductCategoryData{
			Id:           item.Id,
			ParentId:     item.ParentId,
			Name:         item.Name,
			Level:        item.Level,
			ProductCount: item.ProductCount,
			//ProductUnit:  strconv.FormatInt( item.ProductCount,10)+item.ProductUnit,
			ProductUnit: item.ProductUnit,
			NavStatus:   item.NavStatus,
			ShowStatus:  item.ShowStatus,
			Sort:        item.Sort,
			Icon:        item.Icon,
			Keywords:    item.Keywords,
			Description: item.Description,
		})
	}

	return &types.ListProductCategoryResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询商品分类成功",
	}, nil
}
