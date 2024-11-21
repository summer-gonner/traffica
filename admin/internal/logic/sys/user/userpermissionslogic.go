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

type UserPermissionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserPermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPermissionsLogic {
	return &UserPermissionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserPermissionsLogic) UserPermissions() (resp *types.UserPermissionsResp, err error) {
	// 这里的key和生成jwt token时传入的key一致
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()

	userMenusResp, err := l.svcCtx.UserService.UserPermissions(l.ctx, &sysclient.UserPermissionReq{
		UserId: userId,
	})
	if err != nil {
		logc.Errorf(l.ctx, "根据userId: %d,查询用户信息异常:%s", userId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}
	log.Printf("权限%v", userMenusResp.Permission)
	var permissions []string
	for _, u := range userMenusResp.Permission {
		permissions = append(permissions, u)
	}

	return &types.UserPermissionsResp{
		Code:    "000000",
		Message: "获取当前用户权限成功",
		Data:    permissions,
	}, nil
}
