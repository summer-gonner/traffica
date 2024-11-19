package main

import (
	"flag"
	"fmt"
	"github.com/summer-gonner/traffica/sys/internal/config"
	deptserviceServer "github.com/summer-gonner/traffica/sys/internal/server/deptservice"
	dictitemserviceServer "github.com/summer-gonner/traffica/sys/internal/server/dictitemservice"
	dicttypeserviceServer "github.com/summer-gonner/traffica/sys/internal/server/dicttypeservice"
	loginlogserviceServer "github.com/summer-gonner/traffica/sys/internal/server/loginlogservice"
	menuserviceServer "github.com/summer-gonner/traffica/sys/internal/server/menuservice"
	operatelogServer "github.com/summer-gonner/traffica/sys/internal/server/operatelogservice"
	postserviceServer "github.com/summer-gonner/traffica/sys/internal/server/postservice"
	roleserviceServer "github.com/summer-gonner/traffica/sys/internal/server/roleservice"
	userserviceServer "github.com/summer-gonner/traffica/sys/internal/server/userservice"
	"github.com/summer-gonner/traffica/sys/internal/svc"
	"github.com/summer-gonner/traffica/sys/sysclient"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "sys/etc/sys.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		sysclient.RegisterUserServiceServer(grpcServer, userserviceServer.NewUserServiceServer(ctx))
		sysclient.RegisterRoleServiceServer(grpcServer, roleserviceServer.NewRoleServiceServer(ctx))
		sysclient.RegisterPostServiceServer(grpcServer, postserviceServer.NewPostServiceServer(ctx))
		sysclient.RegisterMenuServiceServer(grpcServer, menuserviceServer.NewMenuServiceServer(ctx))
		sysclient.RegisterDictTypeServiceServer(grpcServer, dicttypeserviceServer.NewDictTypeServiceServer(ctx))
		sysclient.RegisterDictItemServiceServer(grpcServer, dictitemserviceServer.NewDictItemServiceServer(ctx))
		sysclient.RegisterDeptServiceServer(grpcServer, deptserviceServer.NewDeptServiceServer(ctx))
		sysclient.RegisterLoginLogServiceServer(grpcServer, loginlogserviceServer.NewLoginLogServiceServer(ctx))
		sysclient.RegisterOperateLogServiceServer(grpcServer, operatelogServer.NewOperateLogServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
