package dict_type

import (
	"github.com/summmer-gonner/traffica/admin/internal/logic/sys/dict_type"
	"github.com/summmer-gonner/traffica/admin/internal/svc"
	"github.com/summmer-gonner/traffica/admin/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryDictTypeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryDictTypeListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := dict_type.NewQueryDictTypeListLogic(r.Context(), svcCtx)
		resp, err := l.QueryDictTypeList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
