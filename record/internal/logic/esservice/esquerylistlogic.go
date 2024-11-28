package esservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/summer-gonner/traffica/record/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/summer-gonner/traffica/record/internal/svc"
	"github.com/summer-gonner/traffica/record/recordclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type EsQueryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEsQueryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EsQueryListLogic {
	return &EsQueryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EsQueryListLogic) EsQueryList(in *recordclient.EsQueryListReq) (*recordclient.EsQueryListResp, error) {
	q := query.RecEsInfo
	// 获取分页参数
	page := in.CurrentPage
	size := in.PageSize

	// 确保分页参数有效
	if page <= 0 {
		page = 1 // 默认从第1页开始
	}
	if size <= 0 {
		size = 10 // 默认每页返回10条数据
	}
	// 计算分页的偏移量
	offset := (page - 1) * size

	// 查询ES数据并进行分页
	res, err := q.WithContext(l.ctx).
		Where(q.Address.Eq(in.Address)).Or().Where(q.Name.Eq(in.Name)).Limit(int(size)).Offset(int(offset)).Find()

	if err != nil {
		logc.Errorf(l.ctx, "根据es地址：%s,查询es表失败,异常：%s", in.Name, err.Error())
		return nil, errors.New(fmt.Sprintf("查询es信息表失败"))
	}
	var records []*recordclient.EsQueryInfoData
	for _, r := range res {
		var eqd recordclient.EsQueryInfoData
		eqd.Id = r.ID
		eqd.Address = r.Address
		eqd.Username = r.Username
		eqd.Password = r.Password
		eqd.Name = r.Name
		eqd.Result = *r.Result
		records = append(records, &eqd)
	}
	var data *recordclient.EsQueryListData
	data.Records = records
	data.CurrentPage = page
	data.PageSize = size
	data.TotalPages = int64(size)
	data.TotalSize = int64(len(records))
	// 构造响应
	resp := &recordclient.EsQueryListResp{
		Code:    "0",  // 总条数
		Data:    data, // 当前页数据
		Message: "查询成功",
	}

	return resp, nil
}
