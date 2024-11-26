package esservicelogic

import (
	"context"
	"fmt"
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
	if in == nil {
		return nil, fmt.Errorf("EsReq is nil")
	}

	if in.Address == "" {
		return nil, fmt.Errorf("elasticsearch address is empty")
	}
	if in.Username == "" {
		return nil, fmt.Errorf("elasticsearch username is empty")
	}
	if in.Password == "" {
		return nil, fmt.Errorf("elasticsearch password is empty")
	}

	client := &es.Client{
		Address:  in.Address, // 设置为你的 Elasticsearch 地址
		Username: in.Username,
		Password: in.Password,
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
