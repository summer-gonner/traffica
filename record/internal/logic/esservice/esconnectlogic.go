package esservicelogic

import (
	"context"
	"github.com/summer-gonner/traffica/pkg/es"

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
	client := &es.Client{
		Address: "http://47.101.198.49:9200", // 设置为你的 Elasticsearch 地址
	}

	return &recordclient.EsResp{}, nil
}
