package logic

import (
	"context"

	"github.com/summer-gonner/traffica/agent/internal/svc"
	"github.com/summer-gonner/traffica/agent/internal/types"

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

func (l *EsConnectLogic) EsConnect() (resp *types.AgentHeartBeatResp, err error) {
	// todo: add your logic here and delete this line
	data := "张三"
	if err := l.svcCtx.KqPusherClient.Push(l.ctx, data); err != nil {
		logx.Errorf("KqPusherClient Push Error , err :%v", err)
	}
	return nil, nil
}
