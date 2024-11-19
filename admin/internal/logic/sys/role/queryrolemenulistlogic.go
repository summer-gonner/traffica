package role

import (
	"context"
	"github.com/summmer-gonner/traffica/admin/internal/common/errorx"
	"github.com/summmer-gonner/traffica/admin/internal/svc"
	"github.com/summmer-gonner/traffica/admin/internal/types"
	"github.com/summmer-gonner/traffica/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryRoleMenuListLogic 根据roleId查询角色的权限
/*
Author: LiuFeiHua
Date: 2024/5/23 17:25
*/
type QueryRoleMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryRoleMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryRoleMenuListLogic {
	return &QueryRoleMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryRoleMenuList 根据roleId查询角色的权限
func (l *QueryRoleMenuListLogic) QueryRoleMenuList(req *types.QueryRoleMenuListReq) (resp *types.QueryRoleMenuListResp, err error) {
	result, err := l.svcCtx.RoleService.QueryRoleMenuList(l.ctx, &sysclient.QueryRoleMenuListReq{
		RoleId: req.RoleId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询菜单信息失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var menuList []types.MenuTreeListData

	for _, menu := range result.List {
		menuList = append(menuList, types.MenuTreeListData{
			Key:      strconv.FormatInt(menu.Id, 10),
			Title:    menu.MenuName,
			ParentId: menu.ParentId,
			Id:       menu.Id,
			Label:    menu.MenuName,
		})
	}

	return &types.QueryRoleMenuListResp{
		Data: types.RoleMenuListData{
			MenuList: menuList,
			MenuIds:  result.MenuIds,
		},
		Code:    "000000",
		Message: "查询角色菜单成功",
	}, nil

}
