package esservicelogic

import (
	"context"
	"github.com/summer-gonner/traffica/pkg/es"
	"github.com/summer-gonner/traffica/record/internal/svc"
	"github.com/summer-gonner/traffica/record/recordclient"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/appengine/log"
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

func (l *EsConnectLogic) EsConnect(in *recordclient.EsReq) (*recordclient.EsResp, error) {

	client := es.Client{
		Address:  in.Address,
		Password: in.Password,
		Username: in.Username,
	}
	err := client.Connect()
	if err != nil {
		log.Infof(l.ctx, "es连接失败%v", err)
		return nil, err
	}
	logx.Infof("es连接成功")
	return &recordclient.EsResp{
		Result:  true,
		Message: "es连接成功",
	}, nil
}
