package home

import (
	"net/http"

	"github.com/summmer-gonner/traffica/api/front/internal/logic/home"
	"github.com/summmer-gonner/traffica/api/front/internal/svc"
	"github.com/summmer-gonner/traffica/api/front/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RecommendHotProductListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecommendHotProductListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := home.NewRecommendHotProductListLogic(r.Context(), svcCtx)
		resp, err := l.RecommendHotProductList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
