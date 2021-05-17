package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"rubbish-cart/rpc/cart/internal/config"
	"rubbish-cart/rpc/cart/model"
)

type ServiceContext struct {
	Config config.Config
	Model model.CartModel
	Redis model.CartRedis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model: model.NewCartModel(sqlx.NewMysql(c.DataSource), c.Cache),
		Redis: model.NewCartRedis(),
	}
}
