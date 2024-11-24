package menu

import (
	"context"
	"fmt"
	"github.com/summer-gonner/traffica/admin/internal/common/errorx"
	"github.com/summer-gonner/traffica/admin/internal/svc"
	"github.com/summer-gonner/traffica/admin/internal/types"
	"github.com/summer-gonner/traffica/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

// QueryMenuListLogic 查询菜单列表
/*
Author: LiuFeiHua
Date: 2023/12/18 15:27
*/
type QueryMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryMenuListLogic {
	return QueryMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMenuList 查询菜单列表
func (l *QueryMenuListLogic) QueryMenuList(req *types.QueryMenuListReq) (*types.QueryMenuListResp, error) {
	result, err := l.svcCtx.MenuService.QueryMenuList(l.ctx, &sysclient.QueryMenuListReq{})

	if err != nil {
		logc.Errorf(l.ctx, "查询菜单列表,参数: %+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var root []*sysclient.MenuListData
	var children []*sysclient.MenuListData
	if len(result.List) > 0 {
		for _, menu := range result.List {
			if menu.MenuType == 0 { //抽出父级
				root = append(root, menu)
			} else {
				children = append(children, menu)
			}

		}
	} else {
		return nil, fmt.Errorf("菜单信息为空")
	}
	var buildMenuTrees []*types.QueryMenuListData
	if len(root) > 0 {
		for _, r := range root {
			userMenuTree := buildMenuTree(r, children)
			buildMenuTrees = append(buildMenuTrees, userMenuTree)
		}
	} else {
		return nil, fmt.Errorf("所有父级菜单为空")
	}
	return &types.QueryMenuListResp{
		Code:    "000000",
		Message: "查询菜单成功",
		Data:    buildMenuTrees,
		Success: true,
		Total:   result.Total,
	}, nil
}

// 构建菜单树
func buildMenuTree(parentMenu *sysclient.MenuListData, userMenusDatas []*sysclient.MenuListData) *types.QueryMenuListData {
	// 初始化u为一个新的QueryMenuListData对象
	u := &types.QueryMenuListData{
		Id:           parentMenu.Id,
		MenuName:     parentMenu.MenuName,
		MenuIcon:     parentMenu.MenuIcon,
		MenuPath:     parentMenu.MenuPath,
		VueComponent: parentMenu.VueComponent,
		VueIcon:      parentMenu.VueIcon,
		MenuType:     parentMenu.MenuType,
		MenuSort:     parentMenu.MenuSort,
		IsVisible:    parentMenu.IsVisible,
		MenuStatus:   parentMenu.MenuStatus,
		MenuPerms:    parentMenu.MenuPerms,
	}

	// 查找所有直接子菜单，并递归构建它们的子菜单
	var children []*types.QueryMenuListData
	for _, userMenusData := range userMenusDatas {
		// 匹配父子菜单关系
		if userMenusData.ParentId == parentMenu.Id {
			// 递归调用buildMenuTree来构建子菜单
			child := buildMenuTree(userMenusData, userMenusDatas)
			children = append(children, child)
		}
	}
	// 设置当前菜单的子菜单列表
	u.Children = children

	return u
}
