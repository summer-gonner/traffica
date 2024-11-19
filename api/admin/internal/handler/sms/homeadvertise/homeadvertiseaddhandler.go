package homeadvertise

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/sms/homeadvertise"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HomeAdvertiseAddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddHomeAdvertiseReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := homeadvertise.NewHomeAdvertiseAddLogic(r.Context(), ctx)
		resp, err := l.HomeAdvertiseAdd(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
