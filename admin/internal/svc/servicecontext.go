package svc

import (
	"github.com/summmer-gonner/traffica/admin/internal/config"
	middleware2 "github.com/summmer-gonner/traffica/admin/internal/middleware"
	"github.com/summmer-gonner/traffica/sys/client/deptservice"
	"github.com/summmer-gonner/traffica/sys/client/dictitemservice"
	"github.com/summmer-gonner/traffica/sys/client/dicttypeservice"
	"github.com/summmer-gonner/traffica/sys/client/loginlogservice"
	"github.com/summmer-gonner/traffica/sys/client/menuservice"
	"github.com/summmer-gonner/traffica/sys/client/operatelogservice"
	"github.com/summmer-gonner/traffica/sys/client/postservice"
	"github.com/summmer-gonner/traffica/sys/client/roleservice"
	"github.com/summmer-gonner/traffica/sys/client/userservice"

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
		CheckUrl:          middleware2.NewCheckUrlMiddleware(newRedis).Handle,
		AddLog:            middleware2.NewAddLogMiddleware(operateLogService).Handle,

		Redis: newRedis,
	}
}

func redisConfig(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	}
}
