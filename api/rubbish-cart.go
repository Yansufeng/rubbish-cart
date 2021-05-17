package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	logic "rubbish-cart/api/internal/logic/ws"

	"rubbish-cart/api/internal/config"
	"rubbish-cart/api/internal/handler"
	"rubbish-cart/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/rubbish-cart-api.yaml", "the config file")

func init() {
	logx.Disable()
	hub := logic.NewHub()
	go hub.Run()
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
