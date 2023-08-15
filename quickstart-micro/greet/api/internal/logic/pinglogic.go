package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"google.com/uhziel/hello-go-zero/quickstart-micro/greet/api/internal/svc"
	"google.com/uhziel/hello-go-zero/quickstart-micro/greet/api/internal/types"
	"google.com/uhziel/hello-go-zero/quickstart-micro/greet/rpc/greet"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() (resp *types.Resp, err error) {
	if _, err = l.svcCtx.GreetRpc.Ping(l.ctx, new(greet.Placeholder)); err != nil {
		return
	}

	resp = new(types.Resp)
	resp.Msg = "pong"

	return
}
