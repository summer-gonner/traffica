package upload

import (
	"github.com/summmer-gonner/traffica/admin/internal/logic/sys/upload"
	"github.com/summmer-gonner/traffica/admin/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := upload.NewUploadLogic(r, r.Context(), svcCtx)
		resp, err := l.Upload()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
