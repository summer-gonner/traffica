package es

import (
	"context"
	"fmt"
	"github.com/summer-gonner/traffica/admin/internal/svc"
	"github.com/summer-gonner/traffica/admin/internal/types"
	"github.com/summer-gonner/traffica/record/recordclient"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
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
	if req.Address == "" {
		return nil, fmt.Errorf("elasticsearch address is empty")
	}
	//if req.Username == "" {
	//	return nil, fmt.Errorf("elasticsearch username is empty")
	//}
	//if req.Password == "" {
	//	return nil, fmt.Errorf("elasticsearch password is empty")
	//}
	log.Printf("地址：%s 用户名：%s,密码：%s", req.Address, req.Username, req.Password)
	res, err := l.svcCtx.EsService.EsConnect(l.ctx, &recordclient.EsReq{
		Username: req.Username,
		Password: req.Password,
		Address:  req.Address,
	})

	if err != nil {
		return nil, fmt.Errorf("连接es失败")
	}
	return &types.EsConnectResp{
		Code:    "000000",
		Message: "连接es成功" + res.Message,
	}, nil

}
