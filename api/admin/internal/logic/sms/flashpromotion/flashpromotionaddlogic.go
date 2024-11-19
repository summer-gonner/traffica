package flashpromotion

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// FlashPromotionAddLogic 秒杀活动
/*
Author: LiuFeiHua
Date: 2024/5/14 10:50
*/
type FlashPromotionAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionAddLogic {
	return FlashPromotionAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// FlashPromotionAdd 添加秒杀活动
func (l *FlashPromotionAddLogic) FlashPromotionAdd(req *types.AddFlashPromotionReq) (*types.AddFlashPromotionResp, error) {
	_, err := l.svcCtx.FlashPromotionService.AddFlashPromotion(l.ctx, &smsclient.AddFlashPromotionReq{
		Title:     req.Title,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Status:    req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加限时购物记录信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加限时购表失败")
	}

	return &types.AddFlashPromotionResp{
		Code:    "000000",
		Message: "添加限时购表成功",
	}, nil
}
