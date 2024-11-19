package flashpromotionsessionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddFlashPromotionSessionLogic 添加限时购场次表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:48
*/
type AddFlashPromotionSessionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFlashPromotionSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFlashPromotionSessionLogic {
	return &AddFlashPromotionSessionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddFlashPromotionSession 添加限时购场次表
func (l *AddFlashPromotionSessionLogic) AddFlashPromotionSession(in *smsclient.AddFlashPromotionSessionReq) (*smsclient.AddFlashPromotionSessionResp, error) {
	err := query.SmsFlashPromotionSession.WithContext(l.ctx).Create(&model.SmsFlashPromotionSession{
		Name:       in.Name,
		StartTime:  in.StartTime,
		EndTime:    in.EndTime,
		Status:     in.Status,
		CreateTime: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.AddFlashPromotionSessionResp{}, nil
}
