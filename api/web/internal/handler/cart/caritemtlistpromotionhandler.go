package cart

import (
	"net/http"

	"github.com/summmer-gonner/traffica/api/web/internal/logic/cart"
	"github.com/summmer-gonner/traffica/api/web/internal/svc"
	"github.com/summmer-gonner/traffica/api/web/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CarItemtListPromotionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CarItemListPromotionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cart.NewCarItemtListPromotionLogic(r.Context(), svcCtx)
		resp, err := l.CarItemtListPromotion(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
