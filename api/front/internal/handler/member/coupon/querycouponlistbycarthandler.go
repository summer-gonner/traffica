package coupon

import (
	"net/http"

	"github.com/summmer-gonner/traffica/api/front/internal/logic/member/coupon"
	"github.com/summmer-gonner/traffica/api/front/internal/svc"
	"github.com/summmer-gonner/traffica/api/front/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryCouponListByCartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CouponListByCartReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := coupon.NewQueryCouponListByCartLogic(r.Context(), svcCtx)
		resp, err := l.QueryCouponListByCart(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
