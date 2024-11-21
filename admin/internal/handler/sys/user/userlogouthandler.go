package user

import (
	"net/http"

	"github.com/summer-gonner/traffica/admin/internal/logic/sys/user"
	"github.com/summer-gonner/traffica/admin/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserLogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserLogoutLogic(r.Context(), svcCtx)
		resp, err := l.UserLogout()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
