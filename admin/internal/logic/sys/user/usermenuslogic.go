package user

import (
	"context"
	"encoding/json"
	"github.com/summer-gonner/traffica/admin/internal/common/errorx"
	"github.com/summer-gonner/traffica/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"log"
	"strconv"

	"github.com/summer-gonner/traffica/admin/internal/svc"
	"github.com/summer-gonner/traffica/admin/internal/types"

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
	for _, u := range userMenusResp.UserMenuData {
		var userMenusData *types.UserMenusData
		userMenusData = &types.UserMenusData{
			Id:   strconv.FormatInt(u.Id, 10),
			Path: u.Path,
			Name: u.Name,
		}
		userMenusDatas = append(userMenusDatas, userMenusData)
	}
	log.Printf("目前当前用户菜单%v", userMenusResp.UserMenuData)
	return &types.UserMenusResp{
		Code:    "000000",
		Message: "获取当前用户菜单成功",
		Data:    userMenusDatas,
	}, nil

}
