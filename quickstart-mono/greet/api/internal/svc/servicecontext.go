package svc

import (
	"google.com/uhziel/hello-go-zero/quickstart-mono/greet/api/internal/config"
)

type ServiceContext struct {
	Config   config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	
	return &ServiceContext{
		Config:   c,
		
	}
}
