package companyaddress

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/order/companyaddress"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateCompanyAddressReceiveStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCompanyAddressReceiveStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := companyaddress.NewUpdateCompanyAddressReceiveStatusLogic(r.Context(), svcCtx)
		resp, err := l.UpdateCompanyAddressReceiveStatus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
