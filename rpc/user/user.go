package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"

	"rubbish-cart/rpc/user/internal/config"
	"rubbish-cart/rpc/user/internal/server"
	"rubbish-cart/rpc/user/internal/svc"
	"rubbish-cart/rpc/user/user"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func init() {
	logx.Disable()
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewUserServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
