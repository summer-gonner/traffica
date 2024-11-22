package user

import (
	"context"
	"encoding/json"
	"github.com/summer-gonner/traffica/admin/internal/common/errorx"
	"github.com/summer-gonner/traffica/admin/internal/svc"
	"github.com/summer-gonner/traffica/admin/internal/types"
	"github.com/summer-gonner/traffica/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
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
	var userMenusDatas []*types.UserMenusData
	//for _, u := range userMenusResp.UserMenuData {
	//	var userMenusData *types.UserMenusData
	//	meta := &types.Meta{
	//		Creator:     u.Meta.Creator,
	//		Updater:     u.Meta.Creator,
	//		Title:       u.Meta.Title,
	//		Permission:  u.Meta.Permission,
	//		Type:        int(u.Meta.Type),
	//		Icon:        u.Meta.Icon,
	//		OrderNo:     int(u.Meta.OrderNo),
	//		Component:   u.Meta.Component,
	//		IsExt:       u.Meta.IsExt,
	//		ExtOpenMode: int(u.Meta.ExtOpenMode),
	//		KeepAlive:   int(u.Meta.KeepAlive),
	//		Show:        int(u.Meta.Show),
	//		ActiveMenu:  u.Meta.ActiveMenu,
	//		Status:      int(u.Meta.Status),
	//	}
	//	userMenusData = &types.UserMenusData{
	//		Id:       strconv.FormatInt(u.Id, 10),
	//		Path:     u.Path,
	//		Name:     u.Name,
	//		Compnent: u.Component,
	//		Meta:     *meta,
	//	}
	//	userMenusDatas = append(userMenusDatas, userMenusData)
	//}
	log.Printf("目前当前用户菜单%v", userMenusResp.UserMenuData)
	var root []*types.UserMenusData
	if len(userMenusResp.UserMenuData) > 0 {
		for _, item := range userMenusResp.UserMenuData {
			if item. {

			}
		}
	}
	return &types.UserMenusResp{
		Code:    "000000",
		Message: "获取当前用户菜单成功",
		Data:    userMenusDatas,
	}, nil

}

//func  buildCurrentUserMenuTress(parentMenu types.UserMenusData, userMenusDatas []*types.UserMenusResp )  {
//	for _,userMenusData := range userMenusDatas {
//		if userMenusData. {
//
//		}
//	}
//}
