package config

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	CartRpc    zrpc.RpcClientConf
	RubbishRpc zrpc.RpcClientConf
	UserRpc    zrpc.RpcClientConf
}
