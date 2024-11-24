package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/summer-gonner/traffica/admin/internal/common/errorx"
	"github.com/summer-gonner/traffica/admin/internal/svc"
	"github.com/summer-gonner/traffica/admin/internal/types"
	"github.com/summer-gonner/traffica/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
)

type UserMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserMenusLogic {
	return &UserMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserMenusLogic) UserMenus() (resp *types.UserMenusResp, err error) {
	// 这里的key和生成jwt token时传入的key一致
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()

	userMenusResp, err := l.svcCtx.UserService.UserMenus(l.ctx, &sysclient.UserMenusReq{
		UserId: userId,
	})
	if err != nil {
		logc.Errorf(l.ctx, "根据userId: %d,查询用户信息异常:%s", userId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}
	var root []*sysclient.UserMenusData
	var children []*sysclient.UserMenusData
	if len(userMenusResp.UserMenuData) > 0 {
		for _, userMenuData := range userMenusResp.UserMenuData {
			if userMenuData.MenuType == 0 { //抽出父级
				root = append(root, userMenuData)
			}
			if userMenuData.MenuType == 1 { //抽出子菜单
				children = append(children, userMenuData)
			}
		}
	} else {
		return nil, fmt.Errorf("菜单信息为空")
	}
	var buildMenuTrees []*types.UserMenusData
	if len(root) > 0 {
		for _, r := range root {
			userMenuTree := buildCurrentUserMenuTree(r, children)
			buildMenuTrees = append(buildMenuTrees, userMenuTree)
		}
	} else {
		return nil, fmt.Errorf("所有父级菜单为空")
	}
	log.Printf("菜单%v", buildMenuTrees)
	return &types.UserMenusResp{
		Code:    "000000",
		Message: "获取当前用户菜单成功",
		Data:    buildMenuTrees,
	}, nil

}

// 构建用户菜单树
func buildCurrentUserMenuTree(parentMenu *sysclient.UserMenusData, userMenusDatas []*sysclient.UserMenusData) *types.UserMenusData {
	// 初始化u为一个新的UserMenusData对象
	u := &types.UserMenusData{
		Id:        strconv.FormatInt(parentMenu.Id, 10),
		Name:      parentMenu.MenuName,
		Path:      parentMenu.MenuPath,
		Component: "",
		Redirect:  parentMenu.VueRedirect,
		Meta: types.Meta{
			Title:       parentMenu.MenuName,
			Icon:        parentMenu.MenuIcon,
			IsExt:       false,
			ExtOpenMode: 1,
			Type:        int(parentMenu.MenuType),
			OrderNo:     int(parentMenu.MenuSort),
			Show:        int(parentMenu.IsVisible),
			ActiveMenu:  "",
			Status:      int(parentMenu.MenuStatus),
			KeepAlive:   0,
		},
	}

	var children []*types.UserMenuChildData
	for _, userMenusData := range userMenusDatas {
		// 匹配父子菜单关系
		if userMenusData.ParentId == parentMenu.Id {
			child := &types.UserMenuChildData{
				Id:        strconv.FormatInt(userMenusData.Id, 10),
				Name:      userMenusData.MenuName,
				Path:      userMenusData.MenuPath,
				Component: userMenusData.VueComponent,
				Meta: types.Meta{
					Title:       userMenusData.MenuName,
					Icon:        userMenusData.VueIcon,
					IsExt:       false,
					ExtOpenMode: 1,
					Type:        int(userMenusData.MenuType),
					OrderNo:     int(userMenusData.MenuSort),
					Show:        int(userMenusData.IsVisible),
					ActiveMenu:  "",
					Status:      int(userMenusData.MenuStatus),
					KeepAlive:   0,
				},
			}
			children = append(children, child)
		}
	}

	// 如果没有子菜单，children会为空
	u.Children = children

	return u
}
