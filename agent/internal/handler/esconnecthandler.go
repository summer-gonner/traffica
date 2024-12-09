package handler

import (
	"net/http"

	"github.com/summer-gonner/traffica/agent/internal/logic"
	"github.com/summer-gonner/traffica/agent/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EsConnectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewEsConnectLogic(r.Context(), svcCtx)
		resp, err := l.EsConnect()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
