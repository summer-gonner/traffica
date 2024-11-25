package main

import (
	"flag"
	"fmt"

	"github.com/summer-gonner/traffica/record/internal/config"
	esserviceServer "github.com/summer-gonner/traffica/record/internal/server/esservice"
	"github.com/summer-gonner/traffica/record/internal/svc"
	"github.com/summer-gonner/traffica/record/recordclient"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/record.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		recordclient.RegisterEsServiceServer(grpcServer, esserviceServer.NewEsServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
