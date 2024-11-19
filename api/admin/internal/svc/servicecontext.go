package svc

import (
	"github.com/feihua/zero-admin/api/admin/internal/config"
	"github.com/feihua/zero-admin/api/admin/internal/middleware"
	"github.com/feihua/zero-admin/rpc/sys/client/deptservice"
	"github.com/feihua/zero-admin/rpc/sys/client/dictitemservice"
	"github.com/feihua/zero-admin/rpc/sys/client/dicttypeservice"
	"github.com/feihua/zero-admin/rpc/sys/client/loginlogservice"
	"github.com/feihua/zero-admin/rpc/sys/client/menuservice"
	"github.com/feihua/zero-admin/rpc/sys/client/operatelogservice"
	"github.com/feihua/zero-admin/rpc/sys/client/postservice"
	"github.com/feihua/zero-admin/rpc/sys/client/roleservice"
	"github.com/feihua/zero-admin/rpc/sys/client/userservice"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	CheckUrl rest.Middleware
	AddLog   rest.Middleware

	//系统相关
	DeptService       deptservice.DeptService
	DictTypeService   dicttypeservice.DictTypeService
	DictItemService   dictitemservice.DictItemService
	PostService       postservice.PostService
	LoginLogService   loginlogservice.LoginLogService
	Operatelogservice operatelogservice.OperateLogService
	MenuService       menuservice.MenuService
	RoleService       roleservice.RoleService
	UserService       userservice.UserService

	Redis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.New(c.Redis.Address, redisConfig(c))
	sysClient := zrpc.MustNewClient(c.SysRpc)

	operateLogService := operatelogservice.NewOperateLogService(sysClient)
	return &ServiceContext{
		Config: c,

		DeptService:       deptservice.NewDeptService(sysClient),
		DictTypeService:   dicttypeservice.NewDictTypeService(sysClient),
		DictItemService:   dictitemservice.NewDictItemService(sysClient),
		PostService:       postservice.NewPostService(sysClient),
		LoginLogService:   loginlogservice.NewLoginLogService(sysClient),
		Operatelogservice: operateLogService,
		MenuService:       menuservice.NewMenuService(sysClient),
		RoleService:       roleservice.NewRoleService(sysClient),
		UserService:       userservice.NewUserService(sysClient),
		CheckUrl:          middleware.NewCheckUrlMiddleware(newRedis).Handle,
		AddLog:            middleware.NewAddLogMiddleware(operateLogService).Handle,

		Redis: newRedis,
	}
}

func redisConfig(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	}
}
