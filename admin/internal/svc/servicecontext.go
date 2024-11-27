package svc

import (
	"github.com/summer-gonner/traffica/admin/internal/config"
	middleware2 "github.com/summer-gonner/traffica/admin/internal/middleware"
	"github.com/summer-gonner/traffica/record/client/esservice"
	"github.com/summer-gonner/traffica/sys/client/deptservice"
	"github.com/summer-gonner/traffica/sys/client/dictitemservice"
	"github.com/summer-gonner/traffica/sys/client/dicttypeservice"
	"github.com/summer-gonner/traffica/sys/client/loginlogservice"
	"github.com/summer-gonner/traffica/sys/client/menuservice"
	"github.com/summer-gonner/traffica/sys/client/operatelogservice"
	"github.com/summer-gonner/traffica/sys/client/postservice"
	"github.com/summer-gonner/traffica/sys/client/roleservice"
	"github.com/summer-gonner/traffica/sys/client/userservice"

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

	//es
	EsService esservice.EsService

	Redis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.New(c.Redis.Address, redisConfig(c))
	sysClient := zrpc.MustNewClient(c.SysRpc)
	recordClient := zrpc.MustNewClient(c.RecordRpc)
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

		EsService: esservice.NewEsService(recordClient),

		CheckUrl: middleware2.NewCheckUrlMiddleware(newRedis).Handle,
		AddLog:   middleware2.NewAddLogMiddleware(operateLogService).Handle,

		Redis: newRedis,
	}
}

func redisConfig(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	}
}
