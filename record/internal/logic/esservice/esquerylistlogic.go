package esservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/summer-gonner/traffica/record/gen/query"
	"github.com/summer-gonner/traffica/record/http3"
	"github.com/zeromicro/go-zero/core/logc"
	"math"

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
	rei := query.RecEsInfo
	// 获取分页参数

	// 查询ES数据并进行分页
	logx.Infof("开始执行查询，ctx：%v", l.ctx)
	q := rei.WithContext(l.ctx)
	if in.Name == "" {
		q = q.Where(rei.Name.Like("%" + in.Name + "%"))
	}
	if in.Address == "" {
		q = q.Where(rei.Address.Like("%" + in.Address + "%"))
	}
	// 确保分页参数有效
	// 计算分页的偏移量
	offset := (in.CurrentPage - 1) * in.PageSize
	res, count, err := q.FindByPage(int(offset), int(in.PageSize))
	if err != nil {
		logx.Errorf("查询失败，错误：%v", err)
	}
	logx.Infof("查询完成")

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
		eqd.Remark = *r.Remark
		eqd.Result = *r.Result
		eqd.CreateTime = r.CreateTime.Format("2006-01-02 15:04:05")
		eqd.UpdateTime = r.UpdateTime.Format("2006-01-02 15:04:05")
		eqd.CreateBy = r.CreateBy
		eqd.UpdateBy = r.UpdateBy
		records = append(records, &eqd)
	}
	data := &recordclient.EsQueryListData{
		Records:     records,
		CurrentPage: in.CurrentPage,
		PageSize:    in.PageSize,
		TotalSize:   count,
		TotalPages:  int64(math.Ceil(float64(count) / float64(in.PageSize))),
	}
	// 构造响应
	resp := &recordclient.EsQueryListResp{
		Code:    http3.SUCCESS, // 总条数
		Data:    data,          // 当前页数据
		Message: http3.QUERY_SUCCESS,
	}

	return resp, nil
}
