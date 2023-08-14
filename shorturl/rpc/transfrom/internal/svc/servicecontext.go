package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.com/uhziel/hello-go-zero/shorturl/rpc/transfrom/internal/config"
	"google.com/uhziel/hello-go-zero/shorturl/rpc/transfrom/model"
)

type ServiceContext struct {
	Config config.Config
	Model  model.ShorturlModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model: model.NewShorturlModel(
			sqlx.NewMysql(c.DataSource),
			c.Cache,
		),
	}
}
