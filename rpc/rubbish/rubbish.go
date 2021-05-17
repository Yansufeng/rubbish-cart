package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"

	"rubbish-cart/rpc/rubbish/internal/config"
	"rubbish-cart/rpc/rubbish/internal/server"
	"rubbish-cart/rpc/rubbish/internal/svc"
	"rubbish-cart/rpc/rubbish/rubbish"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/rubbish.yaml", "the config file")

func init() {
	logx.Disable()
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewRubbishServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		rubbish.RegisterRubbishServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
