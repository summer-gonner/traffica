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
		elastic.SetBasicAuth(es.Username, es.Password), // 如果需要用户名和密码
		elastic.SetSniff(false),                        // 如果需要禁用嗅探（可选）
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

	if in.Address == "" {
		return nil, fmt.Errorf("elasticsearch address is empty")
	}
	if in.Username == "" {
		return nil, fmt.Errorf("elasticsearch username is empty")
	}
	if in.Password == "" {
		return nil, fmt.Errorf("elasticsearch password is empty")
	}

	client := &Client{
		Address:  in.Address, // 设置为你的 Elasticsearch 地址
		Username: in.Username,
		Password: in.Password,
	}

	// 连接到 Elasticsearch
	err := client.connect()
	if err != nil {
		return nil, err
	}
	logx.Infof("es连接成功")
	return &recordclient.EsResp{
		Result:  true,
		Message: "es连接成功",
	}, nil
}
