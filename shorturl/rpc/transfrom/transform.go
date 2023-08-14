package main

import (
	"flag"
	"fmt"

	"google.com/uhziel/hello-go-zero/shorturl/rpc/transfrom/internal/config"
	"google.com/uhziel/hello-go-zero/shorturl/rpc/transfrom/internal/server"
	"google.com/uhziel/hello-go-zero/shorturl/rpc/transfrom/internal/svc"
	"google.com/uhziel/hello-go-zero/shorturl/rpc/transfrom/transform"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/transform.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		transform.RegisterTransformServer(grpcServer, server.NewTransformServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
