package attention

import (
	"net/http"

	"github.com/feihua/zero-admin/api/front/internal/logic/member/attention"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryAttentionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := attention.NewQueryAttentionListLogic(r.Context(), svcCtx)
		resp, err := l.QueryAttentionList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
