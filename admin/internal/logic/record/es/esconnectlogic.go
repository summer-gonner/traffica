package es

import (
	"context"
	"github.com/summer-gonner/traffica/admin/internal/svc"
	"github.com/summer-gonner/traffica/admin/internal/types"
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
	connect, err := l.svcCtx.EsService.EsConnect(l.ctx, &recordclient.EsReq{
		Username: req.Username,
		Password: req.Password,
		Address:  req.Address,
	})
	if err != nil {
		return &types.EsConnectResp{
			Code:    "100001",
			Message: "Es连接失败",
		}, nil
	} else {
		return &types.EsConnectResp{
			Code:    "000000",
			Message: connect.Message,
		}, nil
	}

}
