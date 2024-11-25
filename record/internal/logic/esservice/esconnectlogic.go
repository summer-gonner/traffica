package esservicelogic

import (
	"context"

	"github.com/summer-gonner/traffica/record/internal/svc"
	"github.com/summer-gonner/traffica/record/recordclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type EsConnectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEsConnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EsConnectLogic {
	return &EsConnectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EsConnectLogic) EsConnect(in *recordclient.EsReq) (*recordclient.EsResp, error) {
	// todo: add your logic here and delete this line

	return &recordclient.EsResp{}, nil
}
