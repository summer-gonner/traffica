package es

import (
	"context"
	"fmt"
	"github.com/summer-gonner/traffica/record/recordclient"
	"strconv"

	"github.com/summer-gonner/traffica/admin/internal/svc"
	"github.com/summer-gonner/traffica/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EsQueryPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEsQueryPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EsQueryPageLogic {
	return &EsQueryPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EsQueryPageLogic) EsQueryPage(req *types.EsQueryPageReq) (resp *types.EsQueryPageResp, err error) {

	res, err := l.svcCtx.EsService.EsQueryList(l.ctx, &recordclient.EsQueryListReq{
		CurrentPage: int32(req.CurrentPage),
		PageSize:    int32(req.PageSize),
		Name:        req.Name,
		Address:     req.Address,
	})

	if err != nil {
		return nil, fmt.Errorf("查询es列表失败:%v", err.Error())
	}
	var records []*types.EsQueryPageRecordData
	if len(res.Data.Records) > 0 {
		for _, r := range res.Data.Records {

			record := &types.EsQueryPageRecordData{
				Id:         strconv.FormatInt(r.Id, 10),
				Name:       r.Name,
				Address:    r.Address,
				CreateTime: r.CreateTime,
				UpdateTime: r.UpdateTime,
				CreateBy:   r.CreateBy,
				UpdateBy:   r.UpdateBy,
				Result:     r.Result,
				Remark:     r.Remark,
			}
			records = append(records, record)
		}
	}

	return &types.EsQueryPageResp{
		Code:    "000000",
		Message: "查询成功",
		Data: types.EsQueryPageData{
			CurrentPage: int(res.Data.CurrentPage),
			PageSize:    int(res.Data.PageSize),
			TotalPages:  res.Data.TotalPages,
			TotalSize:   res.Data.TotalSize,
			Records:     records,
		},
	}, nil

	return
}
