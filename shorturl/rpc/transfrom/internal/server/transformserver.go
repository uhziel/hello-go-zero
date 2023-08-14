// Code generated by goctl. DO NOT EDIT.
// Source: transform.proto

package server

import (
	"context"

	"google.com/uhziel/hello-go-zero/shorturl/rpc/transfrom/internal/logic"
	"google.com/uhziel/hello-go-zero/shorturl/rpc/transfrom/internal/svc"
	"google.com/uhziel/hello-go-zero/shorturl/rpc/transfrom/transform"
)

type TransformServer struct {
	svcCtx *svc.ServiceContext
	transform.UnimplementedTransformServer
}

func NewTransformServer(svcCtx *svc.ServiceContext) *TransformServer {
	return &TransformServer{
		svcCtx: svcCtx,
	}
}

func (s *TransformServer) Expand(ctx context.Context, in *transform.ExpandReq) (*transform.ExpandResp, error) {
	l := logic.NewExpandLogic(ctx, s.svcCtx)
	return l.Expand(in)
}

func (s *TransformServer) Shorten(ctx context.Context, in *transform.ShortenReq) (*transform.ShortenResp, error) {
	l := logic.NewShortenLogic(ctx, s.svcCtx)
	return l.Shorten(in)
}