package esservicelogic

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/summer-gonner/traffica/record/gen/query"
	"github.com/summer-gonner/traffica/record/internal/svc"
	"github.com/summer-gonner/traffica/record/recordclient"
	"github.com/zeromicro/go-zero/core/logc"
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

	// 2. 根据es id 从表里查找 es 的信息
	q := query.RecEsInfo
	id, err := strconv.Atoi(in.Id) // 转换 ID
	if err != nil {
		return nil, fmt.Errorf("invalid ES ID format: %v", err)
	}

	// 查询数据库中的 Elasticsearch 信息
	es, err := q.WithContext(l.ctx).Where(q.ID.Eq(int64(id))).First()
	if err != nil {
		logc.Errorf(l.ctx, "查询 Elasticsearch 信息失败, 异常: %s", err.Error())
		return nil, fmt.Errorf("查询 es 信息不存在")
	}

	// 确保查询结果不为空
	if es == nil {
		return nil, fmt.Errorf("未找到指定的 Elasticsearch 信息")
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
		return nil, fmt.Errorf("连接 Elasticsearch 失败: %v", err)
	}

	// 5. 连接成功，记录日志
	logx.Infof("成功连接 Elasticsearch: %v", client)

	// 6. 返回连接成功的响应
	return &recordclient.EsConnectResp{
		Result:  true,
		Message: "连接 Elasticsearch 成功",
	}, nil
}
