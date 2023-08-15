package svc

import "google.com/uhziel/hello-go-zero/quickstart-micro/greet/rpc/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
