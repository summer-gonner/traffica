package cart

import (
	"net/http"

	"github.com/summmer-gonner/traffica/api/front/internal/logic/cart"
	"github.com/summmer-gonner/traffica/api/front/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CarItemListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cart.NewCarItemListLogic(r.Context(), svcCtx)
		resp, err := l.CarItemList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
