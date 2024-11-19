package post

import (
	"github.com/summer-gonner/traffica/admin/internal/logic/sys/post"
	"github.com/summer-gonner/traffica/admin/internal/svc"
	"github.com/summer-gonner/traffica/admin/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdatePostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdatePostReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := post.NewUpdatePostLogic(r.Context(), svcCtx)
		resp, err := l.UpdatePost(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
