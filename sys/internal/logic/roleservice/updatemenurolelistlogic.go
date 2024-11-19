package roleservicelogic

import (
	"context"
	"errors"
	"github.com/summmer-gonner/traffica/sys/gen/model"
	"github.com/summmer-gonner/traffica/sys/gen/query"
	"github.com/summmer-gonner/traffica/sys/internal/svc"
	"github.com/summmer-gonner/traffica/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMenuRoleListLogic 更新角色与菜单的关联
/*
Author: LiuFeiHua
Date: 2024/5/24 15:31
*/
type UpdateMenuRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuRoleListLogic {
	return &UpdateMenuRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMenuRoleList 更新角色与菜单的关联
// 1.删除角色与菜单的关联
// 2.添加角色与菜单的关联
func (l *UpdateMenuRoleListLogic) UpdateMenuRoleList(in *sysclient.UpdateMenuRoleReq) (*sysclient.UpdateMenuRoleResp, error) {
	//sys_role表is_admin为1表示系统预留超级管理员角色,不用关联
	role := query.SysRole
	count, _ := role.WithContext(l.ctx).Where(role.ID.Eq(in.RoleId), role.IsAdmin.Eq(1)).Count()
	if count == 1 {
		return &sysclient.UpdateMenuRoleResp{}, nil
	}

	err := query.Q.Transaction(func(tx *query.Query) error {

		q := tx.SysRoleMenu
		// 1.删除角色与菜单的关联
		if _, err := q.WithContext(l.ctx).Where(q.RoleID.Eq(in.RoleId)).Delete(); err != nil {
			logc.Errorf(l.ctx, "删除角色与菜单的关联失败,参数:%+v,异常:%s", in, err.Error())
			return err
		}

		var roleMenus []*model.SysRoleMenu
		for _, menuId := range in.MenuIds {
			roleMenus = append(roleMenus, &model.SysRoleMenu{
				RoleID: in.RoleId,
				MenuID: menuId,
			})
		}

		// 2.添加角色与菜单的关联
		if err := q.WithContext(l.ctx).CreateInBatches(roleMenus, len(roleMenus)); err != nil {
			logc.Errorf(l.ctx, "添加角色与菜单的关联失败,参数:%+v,异常:%s", in, err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新角色与菜单的关联失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新角色与菜单的关联失败")
	}

	return &sysclient.UpdateMenuRoleResp{}, nil
}
