package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"google.com/uhziel/hello-go-zero/quickstart-mono/greet/api/internal/svc"
	"google.com/uhziel/hello-go-zero/quickstart-mono/greet/api/internal/types"
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
	resp = new(types.Resp)
	resp.Msg = "pong"

	return
}
