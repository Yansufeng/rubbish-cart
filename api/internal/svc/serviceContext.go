package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"rubbish-cart/api/internal/config"
	"rubbish-cart/rpc/cart/cartclient"
	"rubbish-cart/rpc/rubbish/rubbishclient"
	"rubbish-cart/rpc/user/userclient"
)

type ServiceContext struct {
	Config     config.Config
	CartRpc    cartclient.Cart
	RubbishRpc rubbishclient.Rubbish
	UserRpc    userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		CartRpc: cartclient.NewCart(zrpc.MustNewClient(c.CartRpc)),
		RubbishRpc: rubbishclient.NewRubbish(zrpc.MustNewClient(c.RubbishRpc)),
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
