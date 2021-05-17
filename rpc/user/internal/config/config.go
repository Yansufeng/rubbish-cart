package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	Cache      cache.CacheConf
}

const (
	Appid = "wx7229ff74e7759d3a"
	Appsecret = "eb5aa4eece4fc9f489434175e8b9f9ac"
)
