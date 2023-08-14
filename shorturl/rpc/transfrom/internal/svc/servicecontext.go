package svc

import "google.com/uhziel/hello-go-zero/shorturl/rpc/transfrom/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
