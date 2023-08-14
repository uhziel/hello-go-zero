package logic

import (
	"context"

	"google.com/uhziel/hello-go-zero/shorturl/api/internal/svc"
	"google.com/uhziel/hello-go-zero/shorturl/api/internal/types"
	"google.com/uhziel/hello-go-zero/shorturl/rpc/transfrom/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req *types.ShortenReq) (resp *types.ShortenResp, err error) {
	respRpc, err := l.svcCtx.Transform.Shorten(l.ctx, &transform.ShortenReq{
		Url: req.Url,
	})
	if err != nil {
		return
	}

	resp = &types.ShortenResp{
		Shorten: respRpc.Shorten,
	}
	return
}
