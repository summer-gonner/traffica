package userservicelogic

import (
	"context"
	"github.com/summer-gonner/traffica/sys/gen/model"
	"github.com/summer-gonner/traffica/sys/gen/query"
	"github.com/summer-gonner/traffica/sys/internal/logic/common"

	"github.com/summer-gonner/traffica/sys/internal/svc"
	"github.com/summer-gonner/traffica/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPermissionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserPermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPermissionsLogic {
	return &UserPermissionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserPermissions 查询用户权限
func (l *UserPermissionsLogic) UserPermissions(in *sysclient.UserPermissionReq) (*sysclient.UserPermissionResp, error) {
	var result []*model.SysMenu
	if common.IsAdmin(l.ctx, in.UserId, l.svcCtx.DB) {
		result, _ = query.SysMenu.WithContext(l.ctx).Where(query.SysMenu.IsVisible.Eq(1)).Find()
	} else {
		sql := `
				select sm.*
				from sys_user_role sur
						 left join sys_role sr on sur.role_id = sr.id
						 left join sys_role_menu srm on sr.id = srm.role_id
						 left join sys_menu sm on srm.menu_id = sm.id
				where sur.user_id = ? and sm.is_visible=1
				order by sm.id
				`
		db := l.svcCtx.DB
		db.WithContext(l.ctx).Raw(sql, in.UserId).Scan(&result)
	}
	permissions := buildUserPermissions(result)
	return &sysclient.UserPermissionResp{
		Permission: permissions,
	}, nil

}

// 构建返回值
func buildUserPermissions(menus []*model.SysMenu) []string {
	var permissions []string
	for _, menu := range menus {
		permissions = append(permissions, menu.BackgroundURL)
	}
	return permissions
}
