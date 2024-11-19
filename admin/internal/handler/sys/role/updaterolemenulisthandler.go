package role

import (
	"github.com/summmer-gonner/traffica/admin/internal/logic/sys/role"
	"github.com/summmer-gonner/traffica/admin/internal/svc"
	"github.com/summmer-gonner/traffica/admin/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateRoleMenuListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRoleMenuListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewUpdateRoleMenuListLogic(r.Context(), svcCtx)
		resp, err := l.UpdateRoleMenuList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
