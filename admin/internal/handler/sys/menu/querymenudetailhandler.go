package menu

import (
	"github.com/summer-gonner/traffica/admin/internal/logic/sys/menu"
	"github.com/summer-gonner/traffica/admin/internal/svc"
	"github.com/summer-gonner/traffica/admin/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryMenuDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryMenuDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewQueryMenuDetailLogic(r.Context(), svcCtx)
		resp, err := l.QueryMenuDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
