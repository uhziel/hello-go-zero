package svc

import (
	"google.com/uhziel/hello-go-zero/quickstart-micro/greet/api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
	"google.com/uhziel/hello-go-zero/quickstart-micro/greet/rpc/greet"
)

type ServiceContext struct {
	Config   config.Config
	GreetRpc greet.Greet
}

func NewServiceContext(c config.Config) *ServiceContext {
	client := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: "127.0.0.1:8080",
	})
	return &ServiceContext{
		Config:   c,
		GreetRpc: greet.NewGreet(client),
	}
}
