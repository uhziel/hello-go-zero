package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"google.com/uhziel/hello-go-zero/api/demo/internal/logic"
	"google.com/uhziel/hello-go-zero/api/demo/internal/svc"
	"google.com/uhziel/hello-go-zero/api/demo/internal/types"
)

func DemoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDemoLogic(r.Context(), svcCtx)
		resp, err := l.Demo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
