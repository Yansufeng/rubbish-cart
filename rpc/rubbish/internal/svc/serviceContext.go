package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"rubbish-cart/rpc/rubbish/internal/config"
	"rubbish-cart/rpc/rubbish/model"
)

type ServiceContext struct {
	Config config.Config
	Model model.RubbishModel
	TypeModel model.RubbishtypeModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model: model.NewRubbishModel(sqlx.NewMysql(c.DataSource), c.Cache),
		TypeModel: model.NewRubbishtypeModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
