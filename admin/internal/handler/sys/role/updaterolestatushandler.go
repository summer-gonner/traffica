package role

import (
	"github.com/summmer-gonner/traffica/admin/internal/logic/sys/role"
	"github.com/summmer-gonner/traffica/admin/internal/svc"
	"github.com/summmer-gonner/traffica/admin/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateRoleStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRoleStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewUpdateRoleStatusLogic(r.Context(), svcCtx)
		resp, err := l.UpdateRoleStatus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
