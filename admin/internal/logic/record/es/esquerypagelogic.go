package es

import (
	"context"
	"fmt"
	"github.com/summer-gonner/traffica/admin/response"
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
	var datas []*types.EsQueryPageData
	if len(res.Data.Records) > 0 {
		for _, r := range res.Data.Records {

			data := &types.EsQueryPageData{
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
			datas = append(datas, data)
		}
	}
	success := response.SUCCESS

	return &types.EsQueryPageResp{
		Code:        success.Code,
		Message:     success.Message,
		CurrentPage: int(res.Data.CurrentPage),
		PageSize:    int(res.Data.PageSize),
		TotalSize:   res.Data.TotalPages,
		TotalPages:  res.Data.TotalPages,
		Data:        datas,
	}, nil

}
