package brand

import (
	"net/http"

	"github.com/summmer-gonner/traffica/api/front/internal/logic/brand"
	"github.com/summmer-gonner/traffica/api/front/internal/svc"
	"github.com/summmer-gonner/traffica/api/front/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryBrandListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BrandListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := brand.NewQueryBrandListLogic(r.Context(), svcCtx)
		resp, err := l.QueryBrandList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
