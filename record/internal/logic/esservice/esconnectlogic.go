package esservicelogic

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/summer-gonner/traffica/record/gen/model"
	"github.com/summer-gonner/traffica/record/gen/query"
	"github.com/summer-gonner/traffica/record/http3"
	"github.com/summer-gonner/traffica/record/internal/svc"
	"github.com/summer-gonner/traffica/record/recordclient"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
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

func (l *EsConnectLogic) EsConnect(in *recordclient.EsConnectReq) (*recordclient.EsConnectResp, error) {
	// 1. 参数检查，确保输入的请求对象不为空
	if in == nil {
		return nil, fmt.Errorf("EsConnectReq is nil")
	}

	// 2. 根据 es id 从表里查找 es 的信息
	rei := query.RecEsInfo         // 假设这是查询结构或表的引用
	id, err := strconv.Atoi(in.Id) // 转换 ID
	if err != nil {
		return nil, fmt.Errorf("invalid ES ID format: %v", err)
	}

	// 查询数据库中的 Elasticsearch 信息
	es, _ := rei.WithContext(l.ctx).Where(rei.ID.Eq(int64(id))).First()
	// 确保查询结果不为空
	if es == nil {
		return nil, err
	}
	// 3. 根据查询的 es 信息，构建 Elasticsearch 客户端
	var client *elastic.Client
	if es.Name != "" && es.Password != "" {
		// 使用用户名和密码连接
		client, err = elastic.NewClient(
			elastic.SetURL(es.Address),
			elastic.SetBasicAuth(es.Username, es.Password),
			elastic.SetSniff(false), // 禁用嗅探（可选）
		)
	} else {
		// 没有用户名和密码的情况下连接
		client, err = elastic.NewClient(
			elastic.SetURL(es.Address),
			elastic.SetSniff(false), // 禁用嗅探（可选）
		)
	}
	// 4. 错误处理：如果连接失败，返回错误
	if err != nil {
		return nil, fmt.Errorf("es连接失败")
	}
	// 连接成功，更新数据库状态
	remark := http3.ES_CONNECT_SUCCESS
	_, _ = rei.WithContext(l.ctx).Where(rei.ID.Eq(int64(id))).Updates(model.RecEsInfo{
		Result: &http3.Y,
		Remark: &remark,
	})

	// 5. 连接成功，记录日志
	logx.Infof("成功连接 Elasticsearch: %v", client)

	// 6. 返回连接成功的响应
	return &recordclient.EsConnectResp{
		Code:    http3.SUCCESS,
		Message: http3.ES_CONNECT_SUCCESS,
	}, nil
}
