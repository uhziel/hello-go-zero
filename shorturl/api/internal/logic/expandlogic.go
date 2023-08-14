package logic

import (
	"context"

	"google.com/uhziel/hello-go-zero/shorturl/api/internal/svc"
	"google.com/uhziel/hello-go-zero/shorturl/api/internal/types"
	"google.com/uhziel/hello-go-zero/shorturl/rpc/transfrom/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandLogic) Expand(req *types.ExpandReq) (*types.ExpandResp, error) {
	resp, err := l.svcCtx.Transform.Expand(l.ctx, &transform.ExpandReq{
		Shorten: req.Shorten,
	})
	if err != nil {
		return nil, err
	}

	return &types.ExpandResp{
		Url: resp.Url,
	}, nil
}
