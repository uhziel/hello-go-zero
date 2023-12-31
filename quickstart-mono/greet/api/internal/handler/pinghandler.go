package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"google.com/uhziel/hello-go-zero/quickstart-mono/greet/api/internal/logic"
	"google.com/uhziel/hello-go-zero/quickstart-mono/greet/api/internal/svc"
)

func pingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
