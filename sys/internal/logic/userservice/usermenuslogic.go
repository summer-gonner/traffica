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
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserMenusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserMenusLogic {
	return &UserMenusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取当前用户的菜单信息
func (l *UserMenusLogic) UserMenus(in *sysclient.UserMenusReq) (*sysclient.UserMenusResp, error) {
	//1.根据id查询用户信息
	q := query.SysUser
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.UserId)).First()

	// 2.判断用户是否存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logc.Errorf(l.ctx, "用户不存在,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("用户不存在")
	}

	if err != nil {
		logc.Errorf(l.ctx, "查询用户信息,参数：%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询用户信息异常")
	}

	//3.查询用户菜单和权限
	menuList := l.queryUserMenus(in.UserId)

	return &sysclient.UserMenusResp{
		UserMenuData: menuList,
	}, nil

}

// 查询用户菜单和权限
func (l *UserMenusLogic) queryUserMenus(userId int64) []*sysclient.UserMenusData {
	var result []*model.SysMenu
	if common.IsAdmin(l.ctx, userId, l.svcCtx.DB) {
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
		db.WithContext(l.ctx).Raw(sql, userId).Scan(&result)
	}
	return buildUserMenuTree(result)
}

// 构建返回值
func buildUserMenuTree(menus []*model.SysMenu) []*sysclient.UserMenusData {
	var menuListTrees []*sysclient.UserMenusData
	for _, menu := range menus {
		if menu.MenuType == 1 || menu.MenuType == 0 {
			menuListTrees = append(menuListTrees, &sysclient.UserMenusData{
				Id:        menu.ID,
				Path:      menu.MenuPath,
				Name:      menu.MenuName,
				Component: menu.VueComponent,
				Meta: &sysclient.Meta{
					Creator:     menu.CreateBy,
					Updater:     menu.UpdateBy,
					Title:       menu.MenuName,
					Permission:  menu.MenuPerms,
					Type:        menu.MenuType,
					Icon:        menu.MenuIcon,
					OrderNo:     menu.MenuSort,
					Component:   menu.VueComponent,
					IsExt:       true,
					ExtOpenMode: menu.MenuType,
					KeepAlive:   menu.MenuStatus,
					Show:        menu.IsVisible,
					ActiveMenu:  menu.BackgroundURL,
					Status:      menu.MenuStatus,
				},
			})
		}
	}
	return menuListTrees
}
