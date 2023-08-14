package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"google.com/uhziel/hello-go-zero/shorturl/api/internal/logic"
	"google.com/uhziel/hello-go-zero/shorturl/api/internal/svc"
	"google.com/uhziel/hello-go-zero/shorturl/api/internal/types"
)

func ShortenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShortenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewShortenLogic(r.Context(), svcCtx)
		resp, err := l.Shorten(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
