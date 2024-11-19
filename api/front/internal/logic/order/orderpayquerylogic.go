package order

import (
	"context"

	"github.com/summmer-gonner/traffica/api/front/internal/svc"
	"github.com/summmer-gonner/traffica/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderPayQueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderPayQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderPayQueryLogic {
	return &OrderPayQueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderPayQueryLogic) OrderPayQuery(req *types.OrderPayQueryReq) (resp *types.OrderPayQueryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
