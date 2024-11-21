package user

import (
	"context"
	"encoding/json"
	"github.com/summer-gonner/traffica/admin/internal/common/errorx"
	"github.com/summer-gonner/traffica/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

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

	userMenusResp, err := l.svcCtx.UserService.User(l.ctx, &sysclient.ProfileReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	if err != nil {
		logc.Errorf(l.ctx, "根据userId: %d,查询用户信息异常:%s", userId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}
	return &types.UserMenusResp{
		Code:    "000000",
		Message: "获取个人信息成功",
		Data:    []*types.UserMenusData{
			u
		},
	}, nil

}
