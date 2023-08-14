package svc

import (
	"google.com/uhziel/hello-go-zero/shorturl/api/internal/config"
	"google.com/uhziel/hello-go-zero/shorturl/rpc/transfrom/transformclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Transform transformclient.Transform
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Transform: transformclient.NewTransform(zrpc.MustNewClient(c.Transform)),
	}
}
