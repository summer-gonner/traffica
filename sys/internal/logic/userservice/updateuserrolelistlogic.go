package userservicelogic

import (
	"context"
	"errors"
	"github.com/summer-gonner/traffica/sys/gen/model"
	"github.com/summer-gonner/traffica/sys/gen/query"
	"github.com/summer-gonner/traffica/sys/internal/logic/common"
	"github.com/summer-gonner/traffica/sys/internal/svc"
	"github.com/summer-gonner/traffica/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateUserRoleListLogic 更新用户与角色的关联
/*
Author: LiuFeiHua
Date: 2024/5/23 17:38
*/
type UpdateUserRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleListLogic {
	return &UpdateUserRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateUserRoleList 更新用户与角色的关联(角色id为1的是系统预留超级管理员角色,不用关联)
// 1.判断是否为超级管理员
// 2.删除用户与角色的关联
// 3.添加用户与角色的关联
func (l *UpdateUserRoleListLogic) UpdateUserRoleList(in *sysclient.UpdateUserRoleListReq) (*sysclient.UpdateUserRoleListResp, error) {
	// 1.判断是否为超级管理员
	if common.IsAdmin(l.ctx, in.UserId, l.svcCtx.DB) {
		logc.Errorf(l.ctx, "系统预留超级管理员不用分配,参数:%+v", in)
		return nil, errors.New("系统预留超级管理员不用分配")
	}

	err := query.Q.Transaction(func(tx *query.Query) error {

		q := tx.SysUserRole
		userId := in.UserId

		// 2.删除用户与角色的关联
		if _, err := q.WithContext(l.ctx).Where(q.RoleID.Eq(userId)).Delete(); err != nil {
			logc.Errorf(l.ctx, "删除用户与角色的关联失败,参数:%+v,异常:%s", in, err.Error())
			return err
		}

		var userRoles []*model.SysUserRole
		for _, roleId := range in.RoleIds {
			userRoles = append(userRoles, &model.SysUserRole{
				RoleID: roleId,
				UserID: userId,
			})
		}

		// 3.添加用户与角色的关联
		if err := q.WithContext(l.ctx).CreateInBatches(userRoles, len(userRoles)); err != nil {
			logc.Errorf(l.ctx, "添加用户与角色的关联失败,参数:%+v,异常:%s", in, err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新用户与角色的关联失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新用户与角色的关联失败")
	}

	return &sysclient.UpdateUserRoleListResp{}, nil
}
