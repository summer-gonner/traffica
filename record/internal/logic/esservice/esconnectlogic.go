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

	client := &es.Client{
		Address:  "http://47.101.198.49:9200", // 设置为你的 Elasticsearch 地址
		Username: "",
		Password: "",
	}

	// 连接到 Elasticsearch
	err := client.Connect()
	if err != nil {
		return nil, err
	}
	logx.Infof("es连接成功")
	return &recordclient.EsResp{
		Result:  true,
		Message: "es连接成功",
	}, nil
}
