package es

import (
	"context"
	"fmt"
	"github.com/summer-gonner/traffica/admin/internal/svc"
	"github.com/summer-gonner/traffica/admin/internal/types"
	"github.com/summer-gonner/traffica/admin/response"
	"github.com/summer-gonner/traffica/record/recordclient"
	"github.com/zeromicro/go-zero/core/logx"
)

type EsConnectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEsConnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EsConnectLogic {
	return &EsConnectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EsConnectLogic) EsConnect(req *types.EsConnectReq) (resp *types.EsConnectResp, err error) {
	if req.Id == "" {
		return nil, fmt.Errorf("elasticsearch id is empty")
	}
	_, err = l.svcCtx.EsService.EsConnect(l.ctx, &recordclient.EsConnectReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, fmt.Errorf("连接es失败:%v", err.Error())
	}
	success := response.ES_CONNECT_SUCCESS
	return &types.EsConnectResp{
		Code:    success.Code,
		Message: success.Message,
	}, nil

}
