package esservicelogic

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
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

type Client struct {
	Address  string
	Username string
	Password string
	client   *elastic.Client
}

// Connect 连接到 Elasticsearch
func (es *Client) connect() error {
	if es.Address == "" {
		return fmt.Errorf("Address is empty")
	}
	client, err := elastic.NewClient(
		elastic.SetURL(es.Address),
		elastic.SetSniff(false), // 如果需要禁用嗅探（可选）
	)
	if err != nil {
		return fmt.Errorf("connect es error: %s", err)
	}
	es.client = client
	return nil
}

func (l *EsConnectLogic) EsConnect(in *recordclient.EsReq) (*recordclient.EsResp, error) {
	if in == nil {
		return nil, fmt.Errorf("EsReq is nil")
	}

	client, err := elastic.NewClient(
		elastic.SetURL(in.Address),
		elastic.SetSniff(false), // 如果需要禁用嗅探（可选）
	)
	if err != nil {
		return nil, fmt.Errorf("connect es error: %s", err)
	}
	logx.Infof("es连接成功%v", client)
	return &recordclient.EsResp{
		Result:  true,
		Message: "es连接成功",
	}, nil
}
