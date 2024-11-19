package member

import (
	"net/http"

	"github.com/summmer-gonner/traffica/api/web/internal/logic/member/member"
	"github.com/summmer-gonner/traffica/api/web/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func InfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := member.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
